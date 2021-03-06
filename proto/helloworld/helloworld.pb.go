// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetResultRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResultRequest) Reset()         { *m = GetResultRequest{} }
func (m *GetResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetResultRequest) ProtoMessage()    {}
func (*GetResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *GetResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResultRequest.Unmarshal(m, b)
}
func (m *GetResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResultRequest.Marshal(b, m, deterministic)
}
func (m *GetResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResultRequest.Merge(m, src)
}
func (m *GetResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetResultRequest.Size(m)
}
func (m *GetResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResultRequest proto.InternalMessageInfo

func (m *GetResultRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type GetResultReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResultReply) Reset()         { *m = GetResultReply{} }
func (m *GetResultReply) String() string { return proto.CompactTextString(m) }
func (*GetResultReply) ProtoMessage()    {}
func (*GetResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *GetResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResultReply.Unmarshal(m, b)
}
func (m *GetResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResultReply.Marshal(b, m, deterministic)
}
func (m *GetResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResultReply.Merge(m, src)
}
func (m *GetResultReply) XXX_Size() int {
	return xxx_messageInfo_GetResultReply.Size(m)
}
func (m *GetResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetResultReply proto.InternalMessageInfo

func (m *GetResultReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GetResultRequest)(nil), "helloworld.GetResultRequest")
	proto.RegisterType((*GetResultReply)(nil), "helloworld.GetResultReply")
}

func init() {
	proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2)
}

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5,
	0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x54, 0x2a, 0xa9, 0x71, 0x09, 0xb8, 0xa7,
	0x96, 0x04, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0x1a, 0x5c, 0x7c, 0x48, 0xea, 0x0a, 0x72, 0x2a, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0xc0, 0x5c, 0xa8,
	0x3a, 0x28, 0xcf, 0x28, 0x8f, 0x8b, 0xcb, 0x03, 0x64, 0x7b, 0x38, 0xc8, 0x76, 0xa1, 0x04, 0x2e,
	0x4e, 0xb8, 0x3e, 0x21, 0x19, 0x3d, 0x24, 0x97, 0xa2, 0x5b, 0x2b, 0x25, 0x85, 0x43, 0xb6, 0x20,
	0xa7, 0x52, 0x49, 0xb2, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xc2, 0x4a, 0x7c, 0xfa, 0x65, 0x86, 0xfa,
	0xe9, 0xa9, 0x25, 0xba, 0x10, 0xcb, 0xac, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0x1e, 0x31, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x75, 0x4e, 0x86, 0x63, 0x06, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldClient interface {
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultReply, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultReply, error) {
	out := new(GetResultReply)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorld/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
type HelloWorldServer interface {
	GetResult(context.Context, *GetResultRequest) (*GetResultReply, error)
}

// UnimplementedHelloWorldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (*UnimplementedHelloWorldServer) GetResult(ctx context.Context, req *GetResultRequest) (*GetResultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorld/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResult",
			Handler:    _HelloWorld_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
