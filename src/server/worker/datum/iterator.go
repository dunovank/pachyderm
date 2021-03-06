package datum

import (
	"archive/tar"
	"io"
	"path"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
	"github.com/pachyderm/pachyderm/src/client/pps"
	"github.com/pachyderm/pachyderm/src/server/pkg/stream"
	"github.com/pachyderm/pachyderm/src/server/worker/common"
)

// Iterator is the standard interface for a datum iterator.
type Iterator interface {
	// Iterate iterates over a set of datums.
	Iterate(func(*Meta) error) error
}

type pfsIterator struct {
	pachClient *client.APIClient
	input      *pps.PFSInput
}

func newPFSIterator(pachClient *client.APIClient, input *pps.PFSInput) Iterator {
	if input.Commit == "" {
		// this can happen if a pipeline with multiple inputs has been triggered
		// before all commits have inputs
		return &pfsIterator{}
	}
	return &pfsIterator{
		pachClient: pachClient,
		input:      input,
	}
}

func (pi *pfsIterator) Iterate(cb func(*Meta) error) error {
	if pi.input == nil {
		return nil
	}
	repo := pi.input.Repo
	commit := pi.input.Commit
	pattern := pi.input.Glob
	return pi.pachClient.GlobFile(repo, commit, pattern, func(fi *pfs.FileInfo) error {
		// TODO: Implement joins.
		//g := glob.MustCompile(pi.input.Glob, '/')
		//joinOn := g.Replace(fi.File.Path, pi.input.JoinOn)
		return cb(&Meta{
			Inputs: []*common.Input{
				&common.Input{
					FileInfo: fi,
					//JoinOn:     joinOn,
					Name:       pi.input.Name,
					Lazy:       pi.input.Lazy,
					Branch:     pi.input.Branch,
					EmptyFiles: pi.input.EmptyFiles,
					S3:         pi.input.S3,
				},
			},
		})
	})
}

type unionIterator struct {
	iterators []Iterator
}

func newUnionIterator(pachClient *client.APIClient, inputs []*pps.Input) (Iterator, error) {
	ui := &unionIterator{}
	for _, input := range inputs {
		di, err := NewIterator(pachClient, input)
		if err != nil {
			return nil, err
		}
		ui.iterators = append(ui.iterators, di)
	}
	return ui, nil
}

// TODO: It might make sense to check if duplicate datums show up in the merge?
func (ui *unionIterator) Iterate(cb func(*Meta) error) error {
	return Merge(ui.iterators, func(metas []*Meta) error {
		return cb(metas[0])
	})
}

type crossIterator struct {
	iterators []Iterator
}

func newCrossIterator(pachClient *client.APIClient, inputs []*pps.Input) (Iterator, error) {
	ci := &crossIterator{}
	for _, input := range inputs {
		di, err := NewIterator(pachClient, input)
		if err != nil {
			return nil, err
		}
		ci.iterators = append(ci.iterators, di)
	}
	return ci, nil
}

func (ci *crossIterator) Iterate(cb func(*Meta) error) error {
	return iterate(nil, ci.iterators, cb)
}

func iterate(crossInputs []*common.Input, iterators []Iterator, cb func(*Meta) error) error {
	if len(iterators) == 0 {
		return cb(&Meta{Inputs: crossInputs})
	}
	// TODO: Might want to exit fast for the zero datums case.
	return iterators[0].Iterate(func(meta *Meta) error {
		return iterate(append(crossInputs, meta.Inputs...), iterators[1:], cb)
	})
}

// TODO: Need inspect file.
//type gitIterator struct {
//	inputs []*common.Input
//}
//
//func newGitIterator(pachClient *client.APIClient, input *pps.GitInput) (Iterator, error) {
//	if input.Commit == "" {
//		// this can happen if a pipeline with multiple inputs has been triggered
//		// before all commits have inputs
//		return &gitIterator{}, nil
//	}
//	fi, err := pachClient.InspectFile(input.Name, input.Commit, "/commit.json")
//	if err != nil {
//		return nil, err
//	}
//	return &gitIterator{
//		inputs: []*common.Input{
//			&common.Input{
//				FileInfo: fi,
//				Name:     input.Name,
//				Branch:   input.Branch,
//				GitURL:   input.URL,
//			},
//		},
//	}, nil
//}
//
//func (gi *gitIterator) Iterate(cb func([]*common.Input) error) error {
//	if len(gi.inputs) == 0 {
//		return nil
//	}
//	return cb(gi.inputs)
//}

func newCronIterator(pachClient *client.APIClient, input *pps.CronInput) Iterator {
	return newPFSIterator(pachClient, &pps.PFSInput{
		Name:   input.Name,
		Repo:   input.Repo,
		Branch: "master",
		Commit: input.Commit,
		Glob:   "/*",
	})
}

// Hasher is the standard interface for a datum hasher.
type Hasher interface {
	// Hash computes the datum hash based on the inputs.
	Hash([]*common.Input) string
}

type jobIterator struct {
	iterator Iterator
	jobID    string
	hasher   Hasher
}

// NewJobIterator creates a new job iterator.
func NewJobIterator(iterator Iterator, jobID string, hasher Hasher) Iterator {
	return &jobIterator{
		iterator: iterator,
		jobID:    jobID,
		hasher:   hasher,
	}
}

func (ji *jobIterator) Iterate(cb func(*Meta) error) error {
	return ji.iterator.Iterate(func(meta *Meta) error {
		meta.JobID = ji.jobID
		meta.Hash = ji.hasher.Hash(meta.Inputs)
		return cb(meta)
	})
}

type fileSetIterator struct {
	pachClient   *client.APIClient
	repo, commit string
}

// NewFileSetIterator creates a new fileset iterator.
func NewFileSetIterator(pachClient *client.APIClient, repo, commit string) Iterator {
	return &fileSetIterator{
		pachClient: pachClient,
		repo:       repo,
		commit:     commit,
	}
}

func (fsi *fileSetIterator) Iterate(cb func(*Meta) error) error {
	r, err := fsi.pachClient.GetTarFile(fsi.repo, fsi.commit, path.Join("/", MetaPrefix, "*", MetaFileName))
	if err != nil {
		return err
	}
	tr := tar.NewReader(r)
	for {
		_, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		meta := &Meta{}
		if err := jsonpb.Unmarshal(tr, meta); err != nil {
			return err
		}
		if err := cb(meta); err != nil {
			return err
		}
	}
}

// Merge merges multiple datum iterators (key is datum ID).
func Merge(dits []Iterator, cb func([]*Meta) error) error {
	var ss []stream.Stream
	for _, dit := range dits {
		ss = append(ss, newDatumStream(dit, len(ss)))
	}
	pq := stream.NewPriorityQueue(ss)
	return pq.Iterate(func(ss []stream.Stream, _ ...string) error {
		var metas []*Meta
		for _, s := range ss {
			metas = append(metas, s.(*datumStream).meta)
		}
		return cb(metas)
	})
}

type datumStream struct {
	meta     *Meta
	metaChan chan *Meta
	errChan  chan error
	priority int
}

func newDatumStream(dit Iterator, priority int) *datumStream {
	metaChan := make(chan *Meta)
	errChan := make(chan error, 1)
	go func() {
		if err := dit.Iterate(func(meta *Meta) error {
			metaChan <- meta
			return nil
		}); err != nil {
			errChan <- err
			return
		}
		close(metaChan)
	}()
	return &datumStream{
		metaChan: metaChan,
		errChan:  errChan,
		priority: priority,
	}
}

func (ds *datumStream) Next() error {
	select {
	case meta, more := <-ds.metaChan:
		if !more {
			return io.EOF
		}
		ds.meta = meta
		return nil
	case err := <-ds.errChan:
		return err
	}
}

func (ds *datumStream) Key() string {
	return common.DatumID(ds.meta.Inputs)
}

func (ds *datumStream) Priority() int {
	return ds.priority
}

// NewIterator creates a new datum iterator.
func NewIterator(pachClient *client.APIClient, input *pps.Input) (Iterator, error) {
	switch {
	case input.Pfs != nil:
		return newPFSIterator(pachClient, input.Pfs), nil
	case input.Union != nil:
		return newUnionIterator(pachClient, input.Union)
	case input.Cross != nil:
		return newCrossIterator(pachClient, input.Cross)
	//case input.Join != nil:
	//	return newJoinIterator(pachClient, input.Join)
	case input.Cron != nil:
		return newCronIterator(pachClient, input.Cron), nil
		//case input.Git != nil:
		//	return newGitIterator(pachClient, input.Git)
	}
	return nil, errors.Errorf("unrecognized input type: %v", input)
}

// TODO: Implement joins.
//type joinIterator struct {
//	datums   [][]*common.Input
//	location int
//}
//
//func newJoinIterator(pachClient *client.APIClient, join []*pps.Input) (Iterator, error) {
//	result := &joinIterator{}
//	om := ordered_map.NewOrderedMap()
//
//	for i, input := range join {
//		datumIterator, err := NewIterator(pachClient, input)
//		if err != nil {
//			return nil, err
//		}
//		for datumIterator.Next() {
//			x := datumIterator.Datum()
//			for _, k := range x {
//				tupleI, ok := om.Get(k.JoinOn)
//				var tuple [][]*common.Input
//				if !ok {
//					tuple = make([][]*common.Input, len(join))
//				} else {
//					tuple = tupleI.([][]*common.Input)
//				}
//				tuple[i] = append(tuple[i], k)
//				om.Set(k.JoinOn, tuple)
//			}
//		}
//	}
//
//	iter := om.IterFunc()
//	for kv, ok := iter(); ok; kv, ok = iter() {
//		tuple := kv.Value.([][]*common.Input)
//		cross, err := newCrossListIterator(pachClient, tuple)
//		if err != nil {
//			return nil, err
//		}
//		for cross.Next() {
//			result.datums = append(result.datums, cross.Datum())
//		}
//	}
//	result.location = -1
//	return result, nil
//}
//
//func (d *joinIterator) Reset() {
//	d.location = -1
//}
//
//func (d *joinIterator) Len() int {
//	return len(d.datums)
//}
//
//func (d *joinIterator) Next() bool {
//	if d.location < len(d.datums) {
//		d.location++
//	}
//	return d.location < len(d.datums)
//}
//
//func (d *joinIterator) Datum() []*common.Input {
//	var result []*common.Input
//	result = append(result, d.datums[d.location]...)
//	sortInputs(result)
//	return result
//}
//
//func (d *joinIterator) DatumN(n int) []*common.Input {
//	d.location = n
//	return d.Datum()
//}
