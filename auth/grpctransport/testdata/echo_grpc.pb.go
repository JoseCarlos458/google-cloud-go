// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: echo.proto

package echo

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

const (
	Echoer_Echo_FullMethodName = "/echo.Echoer/Echo"
)

// EchoerClient is the client API for Echoer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoerClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type echoerClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoerClient(cc grpc.ClientConnInterface) EchoerClient {
	return &echoerClient{cc}
}

func (c *echoerClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, Echoer_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoerServer is the server API for Echoer service.
// All implementations must embed UnimplementedEchoerServer
// for forward compatibility
type EchoerServer interface {
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
	mustEmbedUnimplementedEchoerServer()
}

// UnimplementedEchoerServer must be embedded to have forward compatible implementations.
type UnimplementedEchoerServer struct {
}

func (UnimplementedEchoerServer) Echo(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoerServer) mustEmbedUnimplementedEchoerServer() {}

// UnsafeEchoerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoerServer will
// result in compilation errors.
type UnsafeEchoerServer interface {
	mustEmbedUnimplementedEchoerServer()
}

func RegisterEchoerServer(s grpc.ServiceRegistrar, srv EchoerServer) {
	s.RegisterService(&Echoer_ServiceDesc, srv)
}

func _Echoer_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Echoer_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoerServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Echoer_ServiceDesc is the grpc.ServiceDesc for Echoer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Echoer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echoer",
	HandlerType: (*EchoerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echoer_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
