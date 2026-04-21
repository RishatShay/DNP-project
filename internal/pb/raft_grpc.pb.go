package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type RaftServiceClient interface {
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	SubmitEntry(ctx context.Context, in *SubmitEntryRequest, opts ...grpc.CallOption) (*SubmitEntryResponse, error)
	QueryLog(ctx context.Context, in *QueryLogRequest, opts ...grpc.CallOption) (*QueryLogResponse, error)
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
}

type raftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftServiceClient(cc grpc.ClientConnInterface) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) SubmitEntry(ctx context.Context, in *SubmitEntryRequest, opts ...grpc.CallOption) (*SubmitEntryResponse, error) {
	out := new(SubmitEntryResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/SubmitEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) QueryLog(ctx context.Context, in *QueryLogRequest, opts ...grpc.CallOption) (*QueryLogResponse, error) {
	out := new(QueryLogResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/QueryLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type RaftServiceServer interface {
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	SubmitEntry(context.Context, *SubmitEntryRequest) (*SubmitEntryResponse, error)
	QueryLog(context.Context, *QueryLogRequest) (*QueryLogResponse, error)
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
}

type UnimplementedRaftServiceServer struct{}

func (*UnimplementedRaftServiceServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (*UnimplementedRaftServiceServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (*UnimplementedRaftServiceServer) SubmitEntry(context.Context, *SubmitEntryRequest) (*SubmitEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEntry not implemented")
}
func (*UnimplementedRaftServiceServer) QueryLog(context.Context, *QueryLogRequest) (*QueryLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLog not implemented")
}
func (*UnimplementedRaftServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func RegisterRaftServiceServer(s *grpc.Server, srv RaftServiceServer) {
	s.RegisterService(&_RaftService_serviceDesc, srv)
}

func _RaftService_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_SubmitEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).SubmitEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/SubmitEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).SubmitEntry(ctx, req.(*SubmitEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_QueryLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).QueryLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/QueryLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).QueryLog(ctx, req.(*QueryLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _RaftService_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _RaftService_AppendEntries_Handler,
		},
		{
			MethodName: "SubmitEntry",
			Handler:    _RaftService_SubmitEntry_Handler,
		},
		{
			MethodName: "QueryLog",
			Handler:    _RaftService_QueryLog_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _RaftService_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/raft.proto",
}
