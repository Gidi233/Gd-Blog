// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: GdBlog/v1/GdBlog.proto

package v1

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
	GdBlog_ListUser_FullMethodName = "/v1.GdBlog/ListUser"
)

// GdBlogClient is the client API for GdBlog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GdBlog 定义了一个 GdBlog RPC 服务.
type GdBlogClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type gdBlogClient struct {
	cc grpc.ClientConnInterface
}

func NewGdBlogClient(cc grpc.ClientConnInterface) GdBlogClient {
	return &gdBlogClient{cc}
}

func (c *gdBlogClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, GdBlog_ListUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GdBlogServer is the server API for GdBlog service.
// All implementations must embed UnimplementedGdBlogServer
// for forward compatibility.
//
// GdBlog 定义了一个 GdBlog RPC 服务.
type GdBlogServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedGdBlogServer()
}

// UnimplementedGdBlogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGdBlogServer struct{}

func (UnimplementedGdBlogServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedGdBlogServer) mustEmbedUnimplementedGdBlogServer() {}
func (UnimplementedGdBlogServer) testEmbeddedByValue()                {}

// UnsafeGdBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GdBlogServer will
// result in compilation errors.
type UnsafeGdBlogServer interface {
	mustEmbedUnimplementedGdBlogServer()
}

func RegisterGdBlogServer(s grpc.ServiceRegistrar, srv GdBlogServer) {
	// If the following call pancis, it indicates UnimplementedGdBlogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GdBlog_ServiceDesc, srv)
}

func _GdBlog_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GdBlogServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GdBlog_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GdBlogServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GdBlog_ServiceDesc is the grpc.ServiceDesc for GdBlog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GdBlog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GdBlog",
	HandlerType: (*GdBlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _GdBlog_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GdBlog/v1/GdBlog.proto",
}
