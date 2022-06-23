// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: config/v1/config.proto

package configv1

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

// ConfigsServiceClient is the client API for ConfigsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigsServiceClient interface {
	GetConfigs(ctx context.Context, in *GetConfigsRequest, opts ...grpc.CallOption) (*GetConfigsResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error)
}

type configsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigsServiceClient(cc grpc.ClientConnInterface) ConfigsServiceClient {
	return &configsServiceClient{cc}
}

func (c *configsServiceClient) GetConfigs(ctx context.Context, in *GetConfigsRequest, opts ...grpc.CallOption) (*GetConfigsResponse, error) {
	out := new(GetConfigsResponse)
	err := c.cc.Invoke(ctx, "/config.v1.ConfigsService/GetConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/config.v1.ConfigsService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsServiceClient) CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error) {
	out := new(CreateConfigResponse)
	err := c.cc.Invoke(ctx, "/config.v1.ConfigsService/CreateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, "/config.v1.ConfigsService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsServiceClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error) {
	out := new(DeleteConfigResponse)
	err := c.cc.Invoke(ctx, "/config.v1.ConfigsService/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigsServiceServer is the server API for ConfigsService service.
// All implementations should embed UnimplementedConfigsServiceServer
// for forward compatibility
type ConfigsServiceServer interface {
	GetConfigs(context.Context, *GetConfigsRequest) (*GetConfigsResponse, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	CreateConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error)
}

// UnimplementedConfigsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConfigsServiceServer struct {
}

func (UnimplementedConfigsServiceServer) GetConfigs(context.Context, *GetConfigsRequest) (*GetConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigs not implemented")
}
func (UnimplementedConfigsServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigsServiceServer) CreateConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfig not implemented")
}
func (UnimplementedConfigsServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigsServiceServer) DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}

// UnsafeConfigsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigsServiceServer will
// result in compilation errors.
type UnsafeConfigsServiceServer interface {
	mustEmbedUnimplementedConfigsServiceServer()
}

func RegisterConfigsServiceServer(s grpc.ServiceRegistrar, srv ConfigsServiceServer) {
	s.RegisterService(&ConfigsService_ServiceDesc, srv)
}

func _ConfigsService_GetConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServiceServer).GetConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.v1.ConfigsService/GetConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServiceServer).GetConfigs(ctx, req.(*GetConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigsService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.v1.ConfigsService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigsService_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServiceServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.v1.ConfigsService/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServiceServer).CreateConfig(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigsService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.v1.ConfigsService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigsService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.v1.ConfigsService/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServiceServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigsService_ServiceDesc is the grpc.ServiceDesc for ConfigsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.v1.ConfigsService",
	HandlerType: (*ConfigsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigs",
			Handler:    _ConfigsService_GetConfigs_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _ConfigsService_GetConfig_Handler,
		},
		{
			MethodName: "CreateConfig",
			Handler:    _ConfigsService_CreateConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigsService_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _ConfigsService_DeleteConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/v1/config.proto",
}
