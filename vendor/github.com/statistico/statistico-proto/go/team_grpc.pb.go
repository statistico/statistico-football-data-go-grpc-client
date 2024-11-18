// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: team.proto

package statistico

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TeamService_GetTeamByID_FullMethodName             = "/statistico.TeamService/GetTeamByID"
	TeamService_GetTeamsByCompetitionId_FullMethodName = "/statistico.TeamService/GetTeamsByCompetitionId"
	TeamService_GetTeamsBySeasonId_FullMethodName      = "/statistico.TeamService/GetTeamsBySeasonId"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	GetTeamByID(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error)
	GetTeamsByCompetitionId(ctx context.Context, in *CompetitionTeamsRequest, opts ...grpc.CallOption) (*TeamsResponse, error)
	GetTeamsBySeasonId(ctx context.Context, in *SeasonTeamsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Team], error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) GetTeamByID(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamsByCompetitionId(ctx context.Context, in *CompetitionTeamsRequest, opts ...grpc.CallOption) (*TeamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamsResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamsByCompetitionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamsBySeasonId(ctx context.Context, in *SeasonTeamsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Team], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TeamService_ServiceDesc.Streams[0], TeamService_GetTeamsBySeasonId_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SeasonTeamsRequest, Team]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TeamService_GetTeamsBySeasonIdClient = grpc.ServerStreamingClient[Team]

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility.
type TeamServiceServer interface {
	GetTeamByID(context.Context, *TeamRequest) (*Team, error)
	GetTeamsByCompetitionId(context.Context, *CompetitionTeamsRequest) (*TeamsResponse, error)
	GetTeamsBySeasonId(*SeasonTeamsRequest, grpc.ServerStreamingServer[Team]) error
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeamServiceServer struct{}

func (UnimplementedTeamServiceServer) GetTeamByID(context.Context, *TeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamByID not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamsByCompetitionId(context.Context, *CompetitionTeamsRequest) (*TeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsByCompetitionId not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamsBySeasonId(*SeasonTeamsRequest, grpc.ServerStreamingServer[Team]) error {
	return status.Errorf(codes.Unimplemented, "method GetTeamsBySeasonId not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}
func (UnimplementedTeamServiceServer) testEmbeddedByValue()                     {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	// If the following call pancis, it indicates UnimplementedTeamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_GetTeamByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamByID(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamsByCompetitionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompetitionTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamsByCompetitionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamsByCompetitionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamsByCompetitionId(ctx, req.(*CompetitionTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamsBySeasonId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SeasonTeamsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TeamServiceServer).GetTeamsBySeasonId(m, &grpc.GenericServerStream[SeasonTeamsRequest, Team]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TeamService_GetTeamsBySeasonIdServer = grpc.ServerStreamingServer[Team]

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statistico.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamByID",
			Handler:    _TeamService_GetTeamByID_Handler,
		},
		{
			MethodName: "GetTeamsByCompetitionId",
			Handler:    _TeamService_GetTeamsByCompetitionId_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTeamsBySeasonId",
			Handler:       _TeamService_GetTeamsBySeasonId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "team.proto",
}