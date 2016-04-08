// Code generated by protoc-gen-go.
// source: server/pps/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pps/persist/persist.proto

It has these top-level messages:
	JobInfo
	JobInfos
	JobOutput
	JobState
	PipelineInfo
	PipelineInfos
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type JobInfo struct {
	JobID           string                      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Transform       *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	PipelineName    string                      `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Shards          uint64                      `protobuf:"varint,4,opt,name=shards" json:"shards,omitempty"`
	Inputs          []*pachyderm_pps.JobInput   `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob       *pachyderm_pps.Job          `protobuf:"bytes,6,opt,name=parent_job,json=parentJob" json:"parent_job,omitempty"`
	CreatedAt       *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	OutputCommit    *pfs.Commit                 `protobuf:"bytes,8,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
	State           pachyderm_pps.JobState      `protobuf:"varint,9,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
	CommitIndex     string                      `protobuf:"bytes,10,opt,name=commit_index,json=commitIndex" json:"commit_index,omitempty"`
	ShardsStarted   uint64                      `protobuf:"varint,11,opt,name=shards_started,json=shardsStarted" json:"shards_started,omitempty"`
	ShardsSucceeded uint64                      `protobuf:"varint,12,opt,name=shards_succeeded,json=shardsSucceeded" json:"shards_succeeded,omitempty"`
	ShardsFailed    uint64                      `protobuf:"varint,13,opt,name=shards_failed,json=shardsFailed" json:"shards_failed,omitempty"`
}

func (m *JobInfo) Reset()                    { *m = JobInfo{} }
func (m *JobInfo) String() string            { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()               {}
func (*JobInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JobInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetInputs() []*pachyderm_pps.JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobInfo) GetParentJob() *pachyderm_pps.Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info,json=jobInfo" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()                    { *m = JobInfos{} }
func (m *JobInfos) String() string            { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()               {}
func (*JobInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type JobOutput struct {
	JobID        string      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	OutputCommit *pfs.Commit `protobuf:"bytes,2,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
}

func (m *JobOutput) Reset()                    { *m = JobOutput{} }
func (m *JobOutput) String() string            { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()               {}
func (*JobOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobOutput) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobState struct {
	JobID string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	State pachyderm_pps.JobState `protobuf:"varint,2,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobState) Reset()                    { *m = JobState{} }
func (m *JobState) String() string            { return proto.CompactTextString(m) }
func (*JobState) ProtoMessage()               {}
func (*JobState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PipelineInfo struct {
	PipelineName string                         `protobuf:"bytes,1,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Transform    *pachyderm_pps.Transform       `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards       uint64                         `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	Inputs       []*pachyderm_pps.PipelineInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	OutputRepo   *pfs.Repo                      `protobuf:"bytes,5,opt,name=output_repo,json=outputRepo" json:"output_repo,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp    `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *PipelineInfo) Reset()                    { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()               {}
func (*PipelineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PipelineInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputs() []*pachyderm_pps.PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info,json=pipelineInfo" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()                    { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()               {}
func (*PipelineInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.persist.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.persist.JobInfos")
	proto.RegisterType((*JobOutput)(nil), "pachyderm.pps.persist.JobOutput")
	proto.RegisterType((*JobState)(nil), "pachyderm.pps.persist.JobState")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.persist.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.persist.PipelineInfos")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.

// Client API for API service

type APIClient interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error)
	InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error)
	GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	SubscribePipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (API_SubscribePipelineInfosClient, error)
	// Shard rpcs
	// Returns the new job info
	StartShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	SucceedShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	FailShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListJobInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeleteJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreatePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetPipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListPipelineInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeletePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SubscribePipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (API_SubscribePipelineInfosClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/pachyderm.pps.persist.API/SubscribePipelineInfos", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISubscribePipelineInfosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_SubscribePipelineInfosClient interface {
	Recv() (*PipelineInfo, error)
	grpc.ClientStream
}

type aPISubscribePipelineInfosClient struct {
	grpc.ClientStream
}

func (x *aPISubscribePipelineInfosClient) Recv() (*PipelineInfo, error) {
	m := new(PipelineInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) StartShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/StartShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SucceedShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/SucceedShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) FailShard(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/FailShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(context.Context, *JobInfo) (*JobInfo, error)
	InspectJob(context.Context, *pachyderm_pps.InspectJobRequest) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(context.Context, *pachyderm_pps.ListJobRequest) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(context.Context, *pachyderm_pps.Job) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(context.Context, *JobOutput) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(context.Context, *JobState) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(context.Context, *PipelineInfo) (*PipelineInfo, error)
	GetPipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(context.Context, *google_protobuf.Empty) (*PipelineInfos, error)
	DeletePipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*google_protobuf.Empty, error)
	SubscribePipelineInfos(*google_protobuf.Empty, API_SubscribePipelineInfosServer) error
	// Shard rpcs
	// Returns the new job info
	StartShard(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	SucceedShard(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	FailShard(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJobInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJobInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobState)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobState(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreatePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PipelineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreatePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetPipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListPipelineInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListPipelineInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeletePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeletePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_SubscribePipelineInfos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).SubscribePipelineInfos(m, &aPISubscribePipelineInfosServer{stream})
}

type API_SubscribePipelineInfosServer interface {
	Send(*PipelineInfo) error
	grpc.ServerStream
}

type aPISubscribePipelineInfosServer struct {
	grpc.ServerStream
}

func (x *aPISubscribePipelineInfosServer) Send(m *PipelineInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _API_StartShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).StartShard(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_SucceedShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).SucceedShard(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_FailShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).FailShard(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.persist.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobInfo",
			Handler:    _API_CreateJobInfo_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _API_InspectJob_Handler,
		},
		{
			MethodName: "ListJobInfos",
			Handler:    _API_ListJobInfos_Handler,
		},
		{
			MethodName: "DeleteJobInfo",
			Handler:    _API_DeleteJobInfo_Handler,
		},
		{
			MethodName: "CreateJobOutput",
			Handler:    _API_CreateJobOutput_Handler,
		},
		{
			MethodName: "CreateJobState",
			Handler:    _API_CreateJobState_Handler,
		},
		{
			MethodName: "CreatePipelineInfo",
			Handler:    _API_CreatePipelineInfo_Handler,
		},
		{
			MethodName: "GetPipelineInfo",
			Handler:    _API_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ListPipelineInfos",
			Handler:    _API_ListPipelineInfos_Handler,
		},
		{
			MethodName: "DeletePipelineInfo",
			Handler:    _API_DeletePipelineInfo_Handler,
		},
		{
			MethodName: "StartShard",
			Handler:    _API_StartShard_Handler,
		},
		{
			MethodName: "SucceedShard",
			Handler:    _API_SucceedShard_Handler,
		},
		{
			MethodName: "FailShard",
			Handler:    _API_FailShard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePipelineInfos",
			Handler:       _API_SubscribePipelineInfos_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 801 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xc9, 0x6e, 0xdb, 0x48,
	0x10, 0x1d, 0x2d, 0x96, 0xc4, 0x12, 0x65, 0xcf, 0x34, 0x66, 0x6c, 0x42, 0xb3, 0x58, 0x23, 0xcf,
	0x00, 0x33, 0x01, 0x42, 0x39, 0x4e, 0x10, 0x20, 0x87, 0x1c, 0xbc, 0xc5, 0x51, 0x56, 0x99, 0xf2,
	0xc5, 0xb9, 0x28, 0x14, 0xd5, 0xb2, 0x69, 0x88, 0x4b, 0xd8, 0x54, 0x10, 0xff, 0x4d, 0xfe, 0x20,
	0xff, 0x92, 0x2f, 0x4a, 0xf5, 0x42, 0xed, 0x0b, 0x13, 0x1f, 0x04, 0xb1, 0x5f, 0xbd, 0x7a, 0x5d,
	0xfd, 0xba, 0x8a, 0x84, 0x1a, 0xa3, 0xd1, 0x47, 0x1a, 0x35, 0xc2, 0x90, 0x35, 0x42, 0x1a, 0x31,
	0x97, 0xc5, 0xc9, 0xbf, 0x19, 0x46, 0x41, 0x1c, 0x90, 0xdf, 0x42, 0xdb, 0xb9, 0xbe, 0xed, 0xd1,
	0xc8, 0x33, 0x91, 0x64, 0xaa, 0x60, 0xf5, 0xf7, 0xab, 0x20, 0xb8, 0x1a, 0xd0, 0x86, 0x20, 0x75,
	0x87, 0xfd, 0x06, 0xf5, 0xc2, 0xf8, 0x56, 0xe6, 0x54, 0x77, 0x67, 0x83, 0xb1, 0xeb, 0x51, 0x16,
	0xdb, 0x5e, 0xa8, 0x08, 0xbf, 0x3a, 0x03, 0x97, 0xfa, 0xb8, 0x55, 0x9f, 0xf1, 0xdf, 0x2c, 0xca,
	0x8b, 0x09, 0x15, 0x5a, 0xff, 0x92, 0x87, 0xe2, 0x8b, 0xa0, 0xdb, 0xf4, 0xfb, 0x58, 0x0c, 0x14,
	0x6e, 0x82, 0x6e, 0xc7, 0xed, 0x19, 0x99, 0x5a, 0xe6, 0x3f, 0xcd, 0xda, 0xc0, 0x55, 0xb3, 0x47,
	0x1e, 0x83, 0x16, 0x47, 0xb6, 0xcf, 0xfa, 0x41, 0xe4, 0x19, 0x59, 0x8c, 0x94, 0x0f, 0x0c, 0x73,
	0xba, 0xee, 0x8b, 0x24, 0x6e, 0x8d, 0xa9, 0x64, 0x0f, 0x2a, 0xa1, 0x1b, 0xd2, 0x81, 0xeb, 0xd3,
	0x8e, 0x6f, 0x7b, 0xd4, 0xc8, 0x09, 0x55, 0x3d, 0x01, 0xdf, 0x20, 0x46, 0xb6, 0xa1, 0xc0, 0xae,
	0xed, 0xa8, 0xc7, 0x8c, 0x3c, 0x46, 0xf3, 0x96, 0x5a, 0x91, 0x06, 0x14, 0x5c, 0x3f, 0x1c, 0xc6,
	0xcc, 0xd8, 0xa8, 0xe5, 0x70, 0xc7, 0x9d, 0x99, 0x1d, 0x45, 0xcd, 0x18, 0xb7, 0x14, 0x8d, 0x3c,
	0x00, 0x08, 0xed, 0x08, 0x0f, 0xd8, 0xc1, 0xaa, 0x8d, 0x82, 0x28, 0x93, 0xcc, 0x27, 0x59, 0x9a,
	0x64, 0xe1, 0x23, 0x79, 0x02, 0xe0, 0x44, 0xd4, 0x8e, 0x69, 0xaf, 0x63, 0xc7, 0x46, 0x51, 0xa4,
	0x54, 0x4d, 0xe9, 0xae, 0x99, 0xb8, 0x6b, 0x5e, 0x24, 0xee, 0x5a, 0x9a, 0x62, 0x1f, 0xc6, 0x64,
	0x1f, 0x2a, 0xc1, 0x30, 0xc6, 0x8d, 0x3b, 0x4e, 0xe0, 0x79, 0x6e, 0x6c, 0x94, 0x44, 0x76, 0xd9,
	0xe4, 0x7e, 0x1f, 0x0b, 0xc8, 0xd2, 0x25, 0x43, 0xae, 0xc8, 0x7d, 0xd8, 0x40, 0x95, 0x98, 0x1a,
	0x1a, 0x32, 0x37, 0x17, 0x9d, 0xa7, 0xcd, 0xc3, 0x96, 0x64, 0x91, 0xbf, 0x41, 0x97, 0xca, 0x1d,
	0xd7, 0xef, 0xd1, 0x4f, 0x06, 0x08, 0xef, 0xca, 0x12, 0x6b, 0x72, 0x88, 0xfc, 0x0b, 0x9b, 0xd2,
	0xac, 0x0e, 0xa6, 0x44, 0x58, 0x97, 0x51, 0x16, 0x16, 0x56, 0x24, 0xda, 0x96, 0x20, 0xf9, 0x1f,
	0x7e, 0x4e, 0x68, 0x43, 0xc7, 0xa1, 0xb4, 0x87, 0x44, 0x5d, 0x10, 0xb7, 0x14, 0x31, 0x81, 0xf9,
	0x8d, 0x29, 0x6a, 0xdf, 0x76, 0x07, 0xc8, 0xab, 0x08, 0x9e, 0x2e, 0xc1, 0x67, 0x02, 0xab, 0x9f,
	0x42, 0x49, 0x35, 0x0c, 0x43, 0x07, 0x4b, 0xa2, 0x63, 0x70, 0x81, 0x3d, 0xc3, 0xef, 0xe9, 0x2f,
	0x73, 0x61, 0x47, 0x9b, 0x2a, 0xc5, 0x2a, 0xde, 0xc8, 0x87, 0xfa, 0x05, 0x68, 0x88, 0xbd, 0x15,
	0x16, 0x2d, 0xeb, 0xbc, 0x39, 0x97, 0xb3, 0x6b, 0x5c, 0xae, 0xb7, 0x44, 0x71, 0xc2, 0xc9, 0x65,
	0xa2, 0xa3, 0x8b, 0xc8, 0xa6, 0xb9, 0x88, 0xfa, 0xe7, 0x2c, 0xe8, 0x2d, 0xd5, 0xb1, 0x62, 0x4a,
	0xe6, 0xda, 0x3a, 0xb3, 0xa0, 0xad, 0x7f, 0x74, 0x66, 0xc6, 0xe3, 0x90, 0x9b, 0x1a, 0x87, 0x47,
	0xa3, 0x71, 0xc8, 0x0b, 0x9b, 0xff, 0x98, 0x11, 0x1b, 0x57, 0x38, 0x39, 0x13, 0xf7, 0xa0, 0xac,
	0xfc, 0x8b, 0x68, 0x18, 0xe0, 0x24, 0xf1, 0x3a, 0x34, 0xe1, 0x9e, 0x85, 0x80, 0x05, 0x32, 0xca,
	0x9f, 0x67, 0x86, 0xa1, 0xf0, 0x1d, 0xc3, 0x50, 0xbf, 0x84, 0xca, 0xa4, 0x43, 0x8c, 0x3c, 0x9f,
	0xb0, 0x68, 0xa2, 0x37, 0xf6, 0x96, 0xf4, 0xc6, 0x64, 0xf2, 0xd8, 0x47, 0xbe, 0x3a, 0xf8, 0x5a,
	0x82, 0xdc, 0x61, 0xab, 0x49, 0xce, 0xa1, 0x72, 0x2c, 0xf6, 0x4b, 0xde, 0x55, 0x6b, 0xfa, 0xac,
	0xba, 0x26, 0x5e, 0xff, 0x89, 0xb4, 0x00, 0x9a, 0x3e, 0x0b, 0xa9, 0x23, 0xde, 0x05, 0xb5, 0x19,
	0xfe, 0x38, 0x64, 0xd1, 0x0f, 0x43, 0x3c, 0x73, 0x2a, 0x45, 0xfd, 0x15, 0x22, 0xa3, 0xe9, 0xf8,
	0x73, 0x26, 0x43, 0x05, 0x13, 0xc1, 0xdd, 0xd5, 0x82, 0x0c, 0x15, 0x9f, 0x42, 0xe5, 0x84, 0x0e,
	0xe8, 0xf8, 0xd8, 0x0b, 0xde, 0x68, 0xd5, 0xed, 0xb9, 0x5b, 0x3a, 0xe5, 0x5f, 0x0b, 0x4c, 0x7f,
	0x0d, 0x5b, 0x23, 0xd7, 0xd4, 0xa4, 0xd5, 0x96, 0x6f, 0x2a, 0x19, 0x2b, 0xe4, 0x5e, 0xc2, 0xe6,
	0x48, 0x4e, 0x8e, 0xd8, 0x8a, 0x23, 0x08, 0xc2, 0x0a, 0xb1, 0xf7, 0x40, 0xa4, 0xd8, 0xf4, 0x70,
	0xa5, 0x68, 0x91, 0x6a, 0x1a, 0x12, 0xee, 0x70, 0x0e, 0x5b, 0x67, 0x34, 0x9e, 0x92, 0xdf, 0x59,
	0x32, 0x36, 0x69, 0x25, 0xdb, 0xf0, 0x0b, 0xbf, 0xc4, 0xe9, 0x6e, 0x5f, 0x72, 0xc6, 0xea, 0x3f,
	0x29, 0x34, 0xf9, 0x25, 0x9f, 0x01, 0x91, 0x97, 0x9c, 0xae, 0xd4, 0xe5, 0x96, 0x5e, 0xc2, 0x76,
	0x7b, 0xd8, 0x65, 0x4e, 0xe4, 0x76, 0x69, 0xba, 0x12, 0xd3, 0x1d, 0x7b, 0x3f, 0x43, 0x8e, 0x00,
	0xc4, 0xf7, 0xa4, 0xcd, 0x5f, 0x47, 0x0b, 0xbb, 0x70, 0xfd, 0x78, 0x9c, 0x80, 0xae, 0x3e, 0x35,
	0x77, 0x51, 0x39, 0x04, 0x8d, 0x7f, 0x88, 0xee, 0x20, 0x71, 0xa4, 0xbd, 0x2b, 0x2a, 0xb0, 0x5b,
	0x10, 0x86, 0x3c, 0xfc, 0x16, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x77, 0xdf, 0xbe, 0xaa, 0x09, 0x00,
	0x00,
}
