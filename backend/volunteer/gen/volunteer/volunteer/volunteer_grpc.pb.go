// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: volunteer/proto/volunteer.proto

package volunteer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VolunteerService_GetVolunteer_FullMethodName    = "/volunteer.VolunteerService/GetVolunteer"
	VolunteerService_ListVolunteers_FullMethodName  = "/volunteer.VolunteerService/ListVolunteers"
	VolunteerService_CreateVolunteer_FullMethodName = "/volunteer.VolunteerService/CreateVolunteer"
)

// VolunteerServiceClient is the client API for VolunteerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolunteerServiceClient interface {
	GetVolunteer(ctx context.Context, in *GetVolunteerRequest, opts ...grpc.CallOption) (*GetVolunteerResponse, error)
	ListVolunteers(ctx context.Context, in *ListVolunteersRequest, opts ...grpc.CallOption) (*ListVolunteersResponse, error)
	CreateVolunteer(ctx context.Context, in *CreateVolunteerRequest, opts ...grpc.CallOption) (*CreateVolunteerResponse, error)
}

type volunteerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVolunteerServiceClient(cc grpc.ClientConnInterface) VolunteerServiceClient {
	return &volunteerServiceClient{cc}
}

func (c *volunteerServiceClient) GetVolunteer(ctx context.Context, in *GetVolunteerRequest, opts ...grpc.CallOption) (*GetVolunteerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVolunteerResponse)
	err := c.cc.Invoke(ctx, VolunteerService_GetVolunteer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volunteerServiceClient) ListVolunteers(ctx context.Context, in *ListVolunteersRequest, opts ...grpc.CallOption) (*ListVolunteersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVolunteersResponse)
	err := c.cc.Invoke(ctx, VolunteerService_ListVolunteers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volunteerServiceClient) CreateVolunteer(ctx context.Context, in *CreateVolunteerRequest, opts ...grpc.CallOption) (*CreateVolunteerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVolunteerResponse)
	err := c.cc.Invoke(ctx, VolunteerService_CreateVolunteer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolunteerServiceServer is the server API for VolunteerService service.
// All implementations must embed UnimplementedVolunteerServiceServer
// for forward compatibility
type VolunteerServiceServer interface {
	GetVolunteer(context.Context, *GetVolunteerRequest) (*GetVolunteerResponse, error)
	ListVolunteers(context.Context, *ListVolunteersRequest) (*ListVolunteersResponse, error)
	CreateVolunteer(context.Context, *CreateVolunteerRequest) (*CreateVolunteerResponse, error)
	mustEmbedUnimplementedVolunteerServiceServer()
}

// UnimplementedVolunteerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVolunteerServiceServer struct {
}

func (UnimplementedVolunteerServiceServer) GetVolunteer(context.Context, *GetVolunteerRequest) (*GetVolunteerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolunteer not implemented")
}
func (UnimplementedVolunteerServiceServer) ListVolunteers(context.Context, *ListVolunteersRequest) (*ListVolunteersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolunteers not implemented")
}
func (UnimplementedVolunteerServiceServer) CreateVolunteer(context.Context, *CreateVolunteerRequest) (*CreateVolunteerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolunteer not implemented")
}
func (UnimplementedVolunteerServiceServer) mustEmbedUnimplementedVolunteerServiceServer() {}

// UnsafeVolunteerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolunteerServiceServer will
// result in compilation errors.
type UnsafeVolunteerServiceServer interface {
	mustEmbedUnimplementedVolunteerServiceServer()
}

func RegisterVolunteerServiceServer(s grpc.ServiceRegistrar, srv VolunteerServiceServer) {
	s.RegisterService(&VolunteerService_ServiceDesc, srv)
}

func _VolunteerService_GetVolunteer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolunteerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolunteerServiceServer).GetVolunteer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolunteerService_GetVolunteer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolunteerServiceServer).GetVolunteer(ctx, req.(*GetVolunteerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolunteerService_ListVolunteers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolunteersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolunteerServiceServer).ListVolunteers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolunteerService_ListVolunteers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolunteerServiceServer).ListVolunteers(ctx, req.(*ListVolunteersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolunteerService_CreateVolunteer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolunteerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolunteerServiceServer).CreateVolunteer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolunteerService_CreateVolunteer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolunteerServiceServer).CreateVolunteer(ctx, req.(*CreateVolunteerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VolunteerService_ServiceDesc is the grpc.ServiceDesc for VolunteerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolunteerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "volunteer.VolunteerService",
	HandlerType: (*VolunteerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVolunteer",
			Handler:    _VolunteerService_GetVolunteer_Handler,
		},
		{
			MethodName: "ListVolunteers",
			Handler:    _VolunteerService_ListVolunteers_Handler,
		},
		{
			MethodName: "CreateVolunteer",
			Handler:    _VolunteerService_CreateVolunteer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "volunteer/proto/volunteer.proto",
}
