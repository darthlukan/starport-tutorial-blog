// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: starporttutorialblog/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("starporttutorialblog/query.proto", fileDescriptor_95051b06caa61958) }

var fileDescriptor_95051b06caa61958 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x93, 0x02, 0x90, 0x52, 0x52, 0x46, 0xc8, 0xa2, 0x46, 0x4a, 0x46, 0x01, 0x4e, 0x40,
	0x47, 0x49, 0x4b, 0x37, 0x0e, 0x96, 0x63, 0xe1, 0x78, 0x8c, 0x3d, 0x46, 0xe4, 0x16, 0x1c, 0x6b,
	0xcb, 0x94, 0x5b, 0xae, 0x92, 0x8b, 0xac, 0x36, 0x51, 0xb4, 0x4d, 0xda, 0xd1, 0xd3, 0xfc, 0xf7,
	0x8a, 0xc7, 0xc8, 0x18, 0x3c, 0x05, 0xe6, 0xc4, 0x14, 0x0c, 0x5a, 0x69, 0x49, 0xc3, 0x4f, 0x52,
	0x61, 0xa8, 0x7d, 0x20, 0xa6, 0xfb, 0xd7, 0x2f, 0x0c, 0xdc, 0xd9, 0xf4, 0x8d, 0xae, 0xde, 0x83,
	0x77, 0x8f, 0xe5, 0x83, 0x26, 0xd2, 0x56, 0x01, 0x7a, 0x03, 0xe8, 0x1c, 0x31, 0xb2, 0x21, 0x17,
	0xd7, 0x9f, 0xe5, 0x53, 0x4b, 0xb1, 0xa7, 0x08, 0x12, 0xa3, 0x5a, 0xc7, 0xe0, 0xb7, 0x91, 0x8a,
	0xb1, 0x01, 0x8f, 0xda, 0xb8, 0x05, 0x5e, 0xd9, 0xe7, 0xbb, 0xe2, 0xe6, 0xe3, 0x42, 0xbc, 0xb5,
	0x87, 0x49, 0xe4, 0xe3, 0x24, 0xf2, 0xd3, 0x24, 0xf2, 0xff, 0x59, 0x64, 0xe3, 0x2c, 0xb2, 0xe3,
	0x2c, 0xb2, 0xcf, 0x77, 0x6d, 0xb8, 0x4b, 0xb2, 0x6e, 0xa9, 0x87, 0xab, 0x2d, 0x6c, 0x62, 0xd5,
	0x66, 0x56, 0x2d, 0x71, 0x7f, 0xb0, 0xdb, 0xcc, 0x83, 0x57, 0x51, 0xde, 0x2e, 0xa3, 0x2f, 0xe7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x45, 0x93, 0x44, 0x18, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "darthlukan.starporttutorialblog.starporttutorialblog.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "starporttutorialblog/query.proto",
}
