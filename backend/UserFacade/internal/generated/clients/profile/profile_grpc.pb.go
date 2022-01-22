// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package profile

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	GetProfileByUserID(ctx context.Context, in *GetProfileByUserIDIn, opts ...grpc.CallOption) (*GetProfileByUserIDOut, error)
	CreateProfile(ctx context.Context, in *CreateProfileIn, opts ...grpc.CallOption) (*CreateProfileOut, error)
	GetAllProfiles(ctx context.Context, in *GetAllProfilesIn, opts ...grpc.CallOption) (*GetAllProfilesOut, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfileByUserID(ctx context.Context, in *GetProfileByUserIDIn, opts ...grpc.CallOption) (*GetProfileByUserIDOut, error) {
	out := new(GetProfileByUserIDOut)
	err := c.cc.Invoke(ctx, "/ProfileService.Profile/GetProfileByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CreateProfile(ctx context.Context, in *CreateProfileIn, opts ...grpc.CallOption) (*CreateProfileOut, error) {
	out := new(CreateProfileOut)
	err := c.cc.Invoke(ctx, "/ProfileService.Profile/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetAllProfiles(ctx context.Context, in *GetAllProfilesIn, opts ...grpc.CallOption) (*GetAllProfilesOut, error) {
	out := new(GetAllProfilesOut)
	err := c.cc.Invoke(ctx, "/ProfileService.Profile/GetAllProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	GetProfileByUserID(context.Context, *GetProfileByUserIDIn) (*GetProfileByUserIDOut, error)
	CreateProfile(context.Context, *CreateProfileIn) (*CreateProfileOut, error)
	GetAllProfiles(context.Context, *GetAllProfilesIn) (*GetAllProfilesOut, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) GetProfileByUserID(context.Context, *GetProfileByUserIDIn) (*GetProfileByUserIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileByUserID not implemented")
}
func (UnimplementedProfileServer) CreateProfile(context.Context, *CreateProfileIn) (*CreateProfileOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServer) GetAllProfiles(context.Context, *GetAllProfilesIn) (*GetAllProfilesOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProfiles not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_GetProfileByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByUserIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService.Profile/GetProfileByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileByUserID(ctx, req.(*GetProfileByUserIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService.Profile/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CreateProfile(ctx, req.(*CreateProfileIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetAllProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProfilesIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetAllProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService.Profile/GetAllProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetAllProfiles(ctx, req.(*GetAllProfilesIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileService.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileByUserID",
			Handler:    _Profile_GetProfileByUserID_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _Profile_CreateProfile_Handler,
		},
		{
			MethodName: "GetAllProfiles",
			Handler:    _Profile_GetAllProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/clients/profile.proto",
}
