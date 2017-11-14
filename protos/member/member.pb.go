// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member/member.proto

/*
Package member is a generated protocol buffer package.

It is generated from these files:
	member/member.proto

It has these top-level messages:
	Member
	MemberCreateRequest
	MemberQueryRequest
	MemberInfoResponse
*/
package member

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "diviner/protos/common"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Member struct {
	Id      string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string             `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Balance float64            `protobuf:"fixed64,3,opt,name=balance" json:"balance,omitempty"`
	Assets  map[string]float64 `protobuf:"bytes,4,rep,name=assets" json:"assets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Member) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Member) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Member) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Member) GetAssets() map[string]float64 {
	if m != nil {
		return m.Assets
	}
	return nil
}

type MemberCreateRequest struct {
	Member *Member              `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
	Check  *common.Verification `protobuf:"bytes,2,opt,name=check" json:"check,omitempty"`
}

func (m *MemberCreateRequest) Reset()                    { *m = MemberCreateRequest{} }
func (m *MemberCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*MemberCreateRequest) ProtoMessage()               {}
func (*MemberCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MemberCreateRequest) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *MemberCreateRequest) GetCheck() *common.Verification {
	if m != nil {
		return m.Check
	}
	return nil
}

type MemberQueryRequest struct {
	Id    string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Check *common.Verification `protobuf:"bytes,2,opt,name=check" json:"check,omitempty"`
}

func (m *MemberQueryRequest) Reset()                    { *m = MemberQueryRequest{} }
func (m *MemberQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*MemberQueryRequest) ProtoMessage()               {}
func (*MemberQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MemberQueryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MemberQueryRequest) GetCheck() *common.Verification {
	if m != nil {
		return m.Check
	}
	return nil
}

type MemberInfoResponse struct {
	Member *Member                    `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
	Time   *google_protobuf.Timestamp `protobuf:"bytes,999,opt,name=time" json:"time,omitempty"`
}

func (m *MemberInfoResponse) Reset()                    { *m = MemberInfoResponse{} }
func (m *MemberInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*MemberInfoResponse) ProtoMessage()               {}
func (*MemberInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MemberInfoResponse) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *MemberInfoResponse) GetTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*Member)(nil), "member.Member")
	proto.RegisterType((*MemberCreateRequest)(nil), "member.MemberCreateRequest")
	proto.RegisterType((*MemberQueryRequest)(nil), "member.MemberQueryRequest")
	proto.RegisterType((*MemberInfoResponse)(nil), "member.MemberInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MemberService service

type MemberServiceClient interface {
	Query(ctx context.Context, in *MemberQueryRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error)
	Create(ctx context.Context, in *MemberCreateRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error)
}

type memberServiceClient struct {
	cc *grpc.ClientConn
}

func NewMemberServiceClient(cc *grpc.ClientConn) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) Query(ctx context.Context, in *MemberQueryRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error) {
	out := new(MemberInfoResponse)
	err := grpc.Invoke(ctx, "/member.MemberService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) Create(ctx context.Context, in *MemberCreateRequest, opts ...grpc.CallOption) (*MemberInfoResponse, error) {
	out := new(MemberInfoResponse)
	err := grpc.Invoke(ctx, "/member.MemberService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MemberService service

type MemberServiceServer interface {
	Query(context.Context, *MemberQueryRequest) (*MemberInfoResponse, error)
	Create(context.Context, *MemberCreateRequest) (*MemberInfoResponse, error)
}

func RegisterMemberServiceServer(s *grpc.Server, srv MemberServiceServer) {
	s.RegisterService(&_MemberService_serviceDesc, srv)
}

func _MemberService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Query(ctx, req.(*MemberQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).Create(ctx, req.(*MemberCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _MemberService_Query_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MemberService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member/member.proto",
}

func init() { proto.RegisterFile("member/member.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xc1, 0x6e, 0xe2, 0x30,
	0x14, 0x5c, 0x07, 0x08, 0xda, 0x17, 0x2d, 0x5a, 0x19, 0x0e, 0x51, 0xf6, 0xb0, 0x28, 0x87, 0x15,
	0xda, 0x43, 0x22, 0x65, 0x2f, 0xdb, 0xde, 0x10, 0xea, 0xa1, 0x87, 0x4a, 0xad, 0x5b, 0xf5, 0xee,
	0x24, 0x0f, 0x6a, 0x41, 0x62, 0x6a, 0x3b, 0x48, 0x7c, 0x49, 0x3f, 0xa6, 0x1f, 0xd4, 0xdf, 0xa8,
	0xb0, 0x13, 0x09, 0x38, 0x54, 0xed, 0xc9, 0x19, 0xcf, 0x64, 0x3c, 0xef, 0x0d, 0x8c, 0x2b, 0xac,
	0x72, 0x54, 0xa9, 0x3b, 0x92, 0xad, 0x92, 0x46, 0x52, 0xdf, 0xa1, 0x68, 0x5c, 0xc8, 0xaa, 0x92,
	0x75, 0xea, 0x0e, 0x47, 0x46, 0xbf, 0x57, 0x52, 0xae, 0x36, 0x98, 0x5a, 0x94, 0x37, 0xcb, 0xd4,
	0x88, 0x0a, 0xb5, 0xe1, 0xd5, 0xd6, 0x09, 0xe2, 0x57, 0x02, 0xfe, 0x8d, 0x35, 0xa0, 0x23, 0xf0,
	0x44, 0x19, 0x92, 0x29, 0x99, 0x7d, 0x67, 0x9e, 0x28, 0x69, 0x08, 0x43, 0x5e, 0x96, 0x0a, 0xb5,
	0x0e, 0x3d, 0x7b, 0xd9, 0xc1, 0x03, 0x93, 0xf3, 0x0d, 0xaf, 0x0b, 0x0c, 0x7b, 0x53, 0x32, 0x23,
	0xac, 0x83, 0x34, 0x03, 0x9f, 0x6b, 0x8d, 0x46, 0x87, 0xfd, 0x69, 0x6f, 0x16, 0x64, 0x51, 0xd2,
	0x66, 0x75, 0x6f, 0x24, 0x73, 0x4b, 0x5e, 0xd5, 0x46, 0xed, 0x59, 0xab, 0x8c, 0x2e, 0x20, 0x38,
	0xba, 0xa6, 0x3f, 0xa1, 0xb7, 0xc6, 0x7d, 0x9b, 0xe3, 0xf0, 0x49, 0x27, 0x30, 0xd8, 0xf1, 0x4d,
	0x83, 0x36, 0x06, 0x61, 0x0e, 0x5c, 0x7a, 0xff, 0x49, 0x2c, 0x60, 0xec, 0x8c, 0x17, 0x0a, 0xb9,
	0x41, 0x86, 0xcf, 0x0d, 0x6a, 0x43, 0xff, 0x40, 0xbb, 0x14, 0xeb, 0x12, 0x64, 0xa3, 0xd3, 0x14,
	0xac, 0x65, 0xe9, 0x5f, 0x18, 0x14, 0x4f, 0x58, 0xac, 0xad, 0x71, 0x90, 0x4d, 0x92, 0x76, 0x77,
	0x8f, 0xa8, 0xc4, 0x52, 0x14, 0xdc, 0x08, 0x59, 0x33, 0x27, 0x89, 0x6f, 0x81, 0xba, 0xbf, 0xef,
	0x1a, 0x54, 0xfb, 0xee, 0xa5, 0xf3, 0x9d, 0x7d, 0xc5, 0xb1, 0xea, 0x1c, 0xaf, 0xeb, 0xa5, 0x64,
	0xa8, 0xb7, 0xb2, 0xd6, 0xf8, 0xe9, 0xec, 0x29, 0xf4, 0x0f, 0x5d, 0x86, 0x6f, 0x43, 0x2b, 0x8b,
	0x12, 0xd7, 0x74, 0xd2, 0x35, 0x9d, 0x3c, 0x74, 0x4d, 0x33, 0x2b, 0xcc, 0x5e, 0x08, 0xfc, 0x70,
	0x1e, 0xf7, 0xa8, 0x76, 0xa2, 0x40, 0x3a, 0x87, 0x81, 0x1d, 0x86, 0x9e, 0xb5, 0x74, 0x3c, 0x61,
	0x74, 0xc6, 0x1d, 0x67, 0x8d, 0xbf, 0xd1, 0x05, 0xf8, 0x6e, 0xf5, 0xf4, 0xd7, 0xa9, 0xee, 0xa4,
	0x90, 0x8f, 0x4d, 0x72, 0xdf, 0x66, 0xfe, 0xf7, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x09, 0x04,
	0x61, 0xdf, 0x02, 0x00, 0x00,
}
