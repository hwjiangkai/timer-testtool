// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: raft.proto

package raft

import (
	context "context"
	raftpb "github.com/linkall-labs/vanus/raft/raftpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_raft_proto protoreflect.FileDescriptor

var file_raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x69,
	0x6e, 0x6b, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x61, 0x6e, 0x75, 0x73, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72,
	0x61, 0x66, 0x74, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x46, 0x0a, 0x0a, 0x52, 0x61, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x6c, 0x6c, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x76, 0x61, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_raft_proto_goTypes = []interface{}{
	(*raftpb.Message)(nil), // 0: raftpb.Message
	(*emptypb.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_raft_proto_depIdxs = []int32{
	0, // 0: linkall.vanus.raft.RaftServer.SendMessage:input_type -> raftpb.Message
	1, // 1: linkall.vanus.raft.RaftServer.SendMessage:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_raft_proto_init() }
func file_raft_proto_init() {
	if File_raft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_proto_goTypes,
		DependencyIndexes: file_raft_proto_depIdxs,
	}.Build()
	File_raft_proto = out.File
	file_raft_proto_rawDesc = nil
	file_raft_proto_goTypes = nil
	file_raft_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RaftServerClient is the client API for RaftServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftServerClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (RaftServer_SendMessageClient, error)
}

type raftServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftServerClient(cc grpc.ClientConnInterface) RaftServerClient {
	return &raftServerClient{cc}
}

func (c *raftServerClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (RaftServer_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RaftServer_serviceDesc.Streams[0], "/linkall.vanus.raft.RaftServer/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftServerSendMessageClient{stream}
	return x, nil
}

type RaftServer_SendMessageClient interface {
	Send(*raftpb.Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type raftServerSendMessageClient struct {
	grpc.ClientStream
}

func (x *raftServerSendMessageClient) Send(m *raftpb.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftServerSendMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftServerServer is the server API for RaftServer service.
type RaftServerServer interface {
	SendMessage(RaftServer_SendMessageServer) error
}

// UnimplementedRaftServerServer can be embedded to have forward compatible implementations.
type UnimplementedRaftServerServer struct {
}

func (*UnimplementedRaftServerServer) SendMessage(RaftServer_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterRaftServerServer(s *grpc.Server, srv RaftServerServer) {
	s.RegisterService(&_RaftServer_serviceDesc, srv)
}

func _RaftServer_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftServerServer).SendMessage(&raftServerSendMessageServer{stream})
}

type RaftServer_SendMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*raftpb.Message, error)
	grpc.ServerStream
}

type raftServerSendMessageServer struct {
	grpc.ServerStream
}

func (x *raftServerSendMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftServerSendMessageServer) Recv() (*raftpb.Message, error) {
	m := new(raftpb.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RaftServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linkall.vanus.raft.RaftServer",
	HandlerType: (*RaftServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _RaftServer_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "raft.proto",
}
