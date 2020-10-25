// Code generated by protoc-gen-go. DO NOT EDIT.
// source: borrow.proto

package borrow

import (
	"context"
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type BorrowReq struct {
	// 书籍号
	BookNo string `protobuf:"bytes,1,opt,name=bookNo,proto3" json:"bookNo,omitempty"`
	// 借书人
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// 预计还书时间: 时间戳（秒）
	ReturnPlanDate       int64    `protobuf:"varint,3,opt,name=returnPlanDate,proto3" json:"returnPlanDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BorrowReq) Reset()         { *m = BorrowReq{} }
func (m *BorrowReq) String() string { return proto.CompactTextString(m) }
func (*BorrowReq) ProtoMessage()    {}
func (*BorrowReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95d669864314e4f0, []int{0}
}

func (m *BorrowReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BorrowReq.Unmarshal(m, b)
}
func (m *BorrowReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BorrowReq.Marshal(b, m, deterministic)
}
func (m *BorrowReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowReq.Merge(m, src)
}
func (m *BorrowReq) XXX_Size() int {
	return xxx_messageInfo_BorrowReq.Size(m)
}
func (m *BorrowReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowReq.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowReq proto.InternalMessageInfo

func (m *BorrowReq) GetBookNo() string {
	if m != nil {
		return m.BookNo
	}
	return ""
}

func (m *BorrowReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BorrowReq) GetReturnPlanDate() int64 {
	if m != nil {
		return m.ReturnPlanDate
	}
	return 0
}

type BaseReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseReply) Reset()         { *m = BaseReply{} }
func (m *BaseReply) String() string { return proto.CompactTextString(m) }
func (*BaseReply) ProtoMessage()    {}
func (*BaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_95d669864314e4f0, []int{1}
}

func (m *BaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseReply.Unmarshal(m, b)
}
func (m *BaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseReply.Marshal(b, m, deterministic)
}
func (m *BaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseReply.Merge(m, src)
}
func (m *BaseReply) XXX_Size() int {
	return xxx_messageInfo_BaseReply.Size(m)
}
func (m *BaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_BaseReply proto.InternalMessageInfo

type ReturnReq struct {
	// 书籍号
	BookNo string `protobuf:"bytes,1,opt,name=bookNo,proto3" json:"bookNo,omitempty"`
	// 借书人
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnReq) Reset()         { *m = ReturnReq{} }
func (m *ReturnReq) String() string { return proto.CompactTextString(m) }
func (*ReturnReq) ProtoMessage()    {}
func (*ReturnReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95d669864314e4f0, []int{2}
}

func (m *ReturnReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnReq.Unmarshal(m, b)
}
func (m *ReturnReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnReq.Marshal(b, m, deterministic)
}
func (m *ReturnReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnReq.Merge(m, src)
}
func (m *ReturnReq) XXX_Size() int {
	return xxx_messageInfo_ReturnReq.Size(m)
}
func (m *ReturnReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnReq proto.InternalMessageInfo

func (m *ReturnReq) GetBookNo() string {
	if m != nil {
		return m.BookNo
	}
	return ""
}

func (m *ReturnReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*BorrowReq)(nil), "borrow.BorrowReq")
	proto.RegisterType((*BaseReply)(nil), "borrow.BaseReply")
	proto.RegisterType((*ReturnReq)(nil), "borrow.ReturnReq")
}

func init() { proto.RegisterFile("borrow.proto", fileDescriptor_95d669864314e4f0) }

var fileDescriptor_95d669864314e4f0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xca, 0x2f, 0x2a,
	0xca, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x92, 0xb9, 0x38,
	0x9d, 0xc0, 0xac, 0xa0, 0xd4, 0x42, 0x21, 0x31, 0x2e, 0xb6, 0xa4, 0xfc, 0xfc, 0x6c, 0xbf, 0x7c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x0f, 0x24, 0x5e, 0x5a, 0x9c, 0x5a, 0xe4, 0x99,
	0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe5, 0x09, 0xa9, 0x71, 0xf1, 0x15, 0xa5, 0x96,
	0x94, 0x16, 0xe5, 0x05, 0xe4, 0x24, 0xe6, 0xb9, 0x24, 0x96, 0xa4, 0x4a, 0x30, 0x83, 0xe5, 0xd1,
	0x44, 0x95, 0xb8, 0xb9, 0x38, 0x9d, 0x12, 0x8b, 0x53, 0x83, 0x52, 0x0b, 0x72, 0x2a, 0x95, 0xac,
	0xb9, 0x38, 0x83, 0xc0, 0xd2, 0x64, 0xd8, 0x68, 0x94, 0xc1, 0xc5, 0x06, 0x71, 0xae, 0x90, 0x1e,
	0x9c, 0x25, 0xa8, 0x07, 0xf5, 0x19, 0xdc, 0x23, 0x52, 0x08, 0x21, 0x98, 0xb5, 0x20, 0xf5, 0x10,
	0x6b, 0x11, 0xea, 0xe1, 0xce, 0xc0, 0xa2, 0x3e, 0x89, 0x0d, 0x1c, 0x4e, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0x4f, 0xb1, 0xd5, 0x37, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BorrowClient is the client API for Borrow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BorrowClient interface {
	// 借书
	Borrow(ctx context.Context, in *BorrowReq, opts ...grpc.CallOption) (*BaseReply, error)
	// 还书
	Return(ctx context.Context, in *ReturnReq, opts ...grpc.CallOption) (*BaseReply, error)
}

type borrowClient struct {
	cc *grpc.ClientConn
}

func NewBorrowClient(cc *grpc.ClientConn) BorrowClient {
	return &borrowClient{cc}
}

func (c *borrowClient) Borrow(ctx context.Context, in *BorrowReq, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, "/borrow.Borrow/Borrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowClient) Return(ctx context.Context, in *ReturnReq, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, "/borrow.Borrow/Return", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorrowServer is the server API for Borrow service.
type BorrowServer interface {
	// 借书
	Borrow(context.Context, *BorrowReq) (*BaseReply, error)
	// 还书
	Return(context.Context, *ReturnReq) (*BaseReply, error)
}

// UnimplementedBorrowServer can be embedded to have forward compatible implementations.
type UnimplementedBorrowServer struct {
}

func (*UnimplementedBorrowServer) Borrow(ctx context.Context, req *BorrowReq) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrow not implemented")
}
func (*UnimplementedBorrowServer) Return(ctx context.Context, req *ReturnReq) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}

func RegisterBorrowServer(s *grpc.Server, srv BorrowServer) {
	s.RegisterService(&_Borrow_serviceDesc, srv)
}

func _Borrow_Borrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServer).Borrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/borrow.Borrow/Borrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServer).Borrow(ctx, req.(*BorrowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Borrow_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/borrow.Borrow/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowServer).Return(ctx, req.(*ReturnReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Borrow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "borrow.Borrow",
	HandlerType: (*BorrowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Borrow",
			Handler:    _Borrow_Borrow_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _Borrow_Return_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "borrow.proto",
}