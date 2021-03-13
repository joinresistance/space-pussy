// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cyber/resources/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgConvert struct {
	Agent    string                                  `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent" yaml:"agent"`
	Amount   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount" yaml:"amount"`
	Resource string                                  `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource" yaml:"resource"`
	EndTime  int64                                   `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time" yaml:"end_time"`
}

func (m *MsgConvert) Reset()         { *m = MsgConvert{} }
func (m *MsgConvert) String() string { return proto.CompactTextString(m) }
func (*MsgConvert) ProtoMessage()    {}
func (*MsgConvert) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{0}
}
func (m *MsgConvert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvert.Merge(m, src)
}
func (m *MsgConvert) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvert) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvert.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvert proto.InternalMessageInfo

type MsgConvertResponse struct {
}

func (m *MsgConvertResponse) Reset()         { *m = MsgConvertResponse{} }
func (m *MsgConvertResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertResponse) ProtoMessage()    {}
func (*MsgConvertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{1}
}
func (m *MsgConvertResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertResponse.Merge(m, src)
}
func (m *MsgConvertResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertResponse proto.InternalMessageInfo

type MsgCreateResource struct {
	Sender   string                                  `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender" yaml:"sender"`
	Resource github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=resource,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"resource" yaml:"resource"`
	Receiver string                                  `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver" yaml:"receiver"`
}

func (m *MsgCreateResource) Reset()         { *m = MsgCreateResource{} }
func (m *MsgCreateResource) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResource) ProtoMessage()    {}
func (*MsgCreateResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{2}
}
func (m *MsgCreateResource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResource.Merge(m, src)
}
func (m *MsgCreateResource) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResource.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResource proto.InternalMessageInfo

type MsgCreateResourceResponse struct {
}

func (m *MsgCreateResourceResponse) Reset()         { *m = MsgCreateResourceResponse{} }
func (m *MsgCreateResourceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResourceResponse) ProtoMessage()    {}
func (*MsgCreateResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{3}
}
func (m *MsgCreateResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResourceResponse.Merge(m, src)
}
func (m *MsgCreateResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResourceResponse proto.InternalMessageInfo

type MsgRedeemResource struct {
	Sender   string                                  `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender" yaml:"sender"`
	Resource github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=resource,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"resource" yaml:"resource"`
}

func (m *MsgRedeemResource) Reset()         { *m = MsgRedeemResource{} }
func (m *MsgRedeemResource) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemResource) ProtoMessage()    {}
func (*MsgRedeemResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{4}
}
func (m *MsgRedeemResource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemResource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemResource.Merge(m, src)
}
func (m *MsgRedeemResource) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemResource.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemResource proto.InternalMessageInfo

type MsgRedeemResourceResponse struct {
}

func (m *MsgRedeemResourceResponse) Reset()         { *m = MsgRedeemResourceResponse{} }
func (m *MsgRedeemResourceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemResourceResponse) ProtoMessage()    {}
func (*MsgRedeemResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94512119bd1d1c33, []int{5}
}
func (m *MsgRedeemResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemResourceResponse.Merge(m, src)
}
func (m *MsgRedeemResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemResourceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvert)(nil), "cyber.resources.v1beta1.MsgConvert")
	proto.RegisterType((*MsgConvertResponse)(nil), "cyber.resources.v1beta1.MsgConvertResponse")
	proto.RegisterType((*MsgCreateResource)(nil), "cyber.resources.v1beta1.MsgCreateResource")
	proto.RegisterType((*MsgCreateResourceResponse)(nil), "cyber.resources.v1beta1.MsgCreateResourceResponse")
	proto.RegisterType((*MsgRedeemResource)(nil), "cyber.resources.v1beta1.MsgRedeemResource")
	proto.RegisterType((*MsgRedeemResourceResponse)(nil), "cyber.resources.v1beta1.MsgRedeemResourceResponse")
}

func init() { proto.RegisterFile("cyber/resources/v1beta1/tx.proto", fileDescriptor_94512119bd1d1c33) }

var fileDescriptor_94512119bd1d1c33 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xbd, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x93, 0x16, 0x7a, 0xc5, 0x7c, 0x89, 0xe8, 0xa4, 0x6b, 0x7b, 0x52, 0x5c, 0x85, 0x81,
	0x0a, 0x74, 0x89, 0xae, 0xdd, 0xca, 0xd6, 0xdb, 0x90, 0xca, 0x10, 0x31, 0xc1, 0x70, 0xca, 0xc7,
	0xab, 0x10, 0x41, 0xec, 0x2a, 0x76, 0xab, 0xeb, 0xc0, 0x88, 0xc4, 0xc8, 0xc6, 0x7a, 0xff, 0x09,
	0x1b, 0xba, 0xf1, 0x46, 0xc4, 0x10, 0xa1, 0x76, 0x41, 0x1d, 0x3b, 0x32, 0xa1, 0xda, 0x71, 0x93,
	0xde, 0x41, 0x29, 0x2b, 0x53, 0xe2, 0xc7, 0x8f, 0xdf, 0x27, 0xef, 0xcf, 0xb1, 0x51, 0x3b, 0x98,
	0xfa, 0x90, 0x3a, 0x29, 0x30, 0x3a, 0x4e, 0x03, 0x60, 0xce, 0xe4, 0xd8, 0x07, 0xee, 0x1d, 0x3b,
	0xfc, 0xcc, 0x1e, 0xa5, 0x94, 0x53, 0xe3, 0x40, 0x38, 0xec, 0xb5, 0xc3, 0xce, 0x1d, 0xad, 0xfd,
	0x88, 0x46, 0x54, 0x78, 0x9c, 0xd5, 0x9b, 0xb4, 0xb7, 0x0e, 0x02, 0xca, 0x12, 0xca, 0x4e, 0xe5,
	0x44, 0x40, 0x63, 0x22, 0x27, 0xac, 0xcf, 0x15, 0x84, 0x86, 0x2c, 0x3a, 0xa1, 0x64, 0x02, 0x29,
	0x37, 0x1c, 0x74, 0xd3, 0x8b, 0x80, 0xf0, 0x86, 0xde, 0xd6, 0x3b, 0xb7, 0x06, 0xcd, 0x45, 0x86,
	0xa5, 0xb0, 0xcc, 0xf0, 0x9d, 0xa9, 0x97, 0xbc, 0xed, 0x5b, 0x62, 0x68, 0xb9, 0x52, 0x36, 0xde,
	0xa1, 0x9a, 0x97, 0xd0, 0x31, 0xe1, 0x8d, 0x4a, 0x5b, 0xef, 0xdc, 0xee, 0x36, 0x6d, 0x99, 0x64,
	0xfb, 0x1e, 0x03, 0xf5, 0x51, 0xf6, 0x09, 0x8d, 0xc9, 0xe0, 0xd9, 0x45, 0x86, 0xb5, 0x45, 0x86,
	0xf3, 0x05, 0xcb, 0x0c, 0xdf, 0xcd, 0x2b, 0x8a, 0xb1, 0xf5, 0x33, 0xc3, 0x8f, 0xa2, 0x98, 0xbf,
	0x1e, 0xfb, 0x76, 0x40, 0x13, 0x47, 0xd6, 0xc9, 0x1f, 0x47, 0x2c, 0x7c, 0xe3, 0xf0, 0xe9, 0x08,
	0x98, 0xa8, 0xe5, 0xe6, 0x35, 0x8c, 0xa7, 0xa8, 0xae, 0x10, 0x34, 0xaa, 0xe2, 0x93, 0xf1, 0x22,
	0xc3, 0x6b, 0x6d, 0x99, 0xe1, 0xfb, 0x32, 0x43, 0x29, 0x96, 0xbb, 0x9e, 0x34, 0xfa, 0xa8, 0x0e,
	0x24, 0x3c, 0xe5, 0x71, 0x02, 0x8d, 0x1b, 0x6d, 0xbd, 0x53, 0x95, 0x8b, 0x95, 0x56, 0x2c, 0x56,
	0x8a, 0xe5, 0xee, 0x01, 0x09, 0x5f, 0xc4, 0x09, 0xf4, 0xeb, 0x1f, 0xce, 0xb1, 0xf6, 0xe3, 0x1c,
	0x6b, 0xd6, 0x3e, 0x32, 0x0a, 0x80, 0x2e, 0xb0, 0x11, 0x25, 0x0c, 0xac, 0x4f, 0x15, 0xf4, 0x60,
	0x25, 0xa7, 0xe0, 0x71, 0x70, 0x55, 0x62, 0x0f, 0xd5, 0x18, 0x90, 0x10, 0xd2, 0x9c, 0xef, 0xe1,
	0x0a, 0x87, 0x54, 0x0a, 0x1c, 0x72, 0x6c, 0xb9, 0xf9, 0x84, 0xf1, 0x5e, 0x2f, 0x35, 0xf9, 0x57,
	0xca, 0xcf, 0x73, 0xca, 0xdb, 0x18, 0xfc, 0x0b, 0xe9, 0x02, 0x97, 0x60, 0x1d, 0x40, 0x3c, 0x81,
	0x74, 0x93, 0xb5, 0xd4, 0xca, 0x39, 0x52, 0x11, 0xac, 0xe5, 0x6b, 0x89, 0xd7, 0x21, 0x6a, 0x5e,
	0x03, 0xb3, 0xc6, 0xf6, 0x4d, 0x17, 0xd8, 0x5c, 0x08, 0x01, 0x92, 0xff, 0x02, 0xdb, 0xb5, 0xce,
	0x37, 0x7b, 0x53, 0x9d, 0x77, 0xbf, 0x54, 0x50, 0x75, 0xc8, 0x22, 0xe3, 0x15, 0xda, 0x53, 0x87,
	0xf1, 0xa1, 0xfd, 0x87, 0x43, 0x6e, 0x17, 0x3f, 0x5c, 0xeb, 0xc9, 0x0e, 0x26, 0x15, 0x62, 0x8c,
	0xd0, 0xbd, 0x2b, 0x7f, 0xe4, 0xe3, 0xad, 0xcb, 0x37, 0xbc, 0xad, 0xee, 0xee, 0xde, 0x72, 0xe2,
	0x95, 0xcd, 0xdc, 0x9a, 0xb8, 0xe9, 0xdd, 0x9e, 0xf8, 0x7b, 0x90, 0x83, 0xe1, 0xc5, 0xcc, 0xd4,
	0x2f, 0x67, 0xa6, 0xfe, 0x7d, 0x66, 0xea, 0x1f, 0xe7, 0xa6, 0x76, 0x39, 0x37, 0xb5, 0xaf, 0x73,
	0x53, 0x7b, 0xd9, 0x2b, 0x6f, 0xde, 0xaa, 0x6e, 0x40, 0x49, 0x94, 0x02, 0x63, 0x4e, 0x44, 0x8f,
	0xe4, 0x8d, 0x7b, 0x56, 0xba, 0x73, 0xc5, 0x6e, 0xfa, 0x35, 0x71, 0x4f, 0xf6, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0xf9, 0xc1, 0xc1, 0x93, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Convert(ctx context.Context, in *MsgConvert, opts ...grpc.CallOption) (*MsgConvertResponse, error)
	CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error)
	RedeemResource(ctx context.Context, in *MsgRedeemResource, opts ...grpc.CallOption) (*MsgRedeemResourceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Convert(ctx context.Context, in *MsgConvert, opts ...grpc.CallOption) (*MsgConvertResponse, error) {
	out := new(MsgConvertResponse)
	err := c.cc.Invoke(ctx, "/cyber.resources.v1beta1.Msg/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error) {
	out := new(MsgCreateResourceResponse)
	err := c.cc.Invoke(ctx, "/cyber.resources.v1beta1.Msg/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RedeemResource(ctx context.Context, in *MsgRedeemResource, opts ...grpc.CallOption) (*MsgRedeemResourceResponse, error) {
	out := new(MsgRedeemResourceResponse)
	err := c.cc.Invoke(ctx, "/cyber.resources.v1beta1.Msg/RedeemResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Convert(context.Context, *MsgConvert) (*MsgConvertResponse, error)
	CreateResource(context.Context, *MsgCreateResource) (*MsgCreateResourceResponse, error)
	RedeemResource(context.Context, *MsgRedeemResource) (*MsgRedeemResourceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Convert(ctx context.Context, req *MsgConvert) (*MsgConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (*UnimplementedMsgServer) CreateResource(ctx context.Context, req *MsgCreateResource) (*MsgCreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (*UnimplementedMsgServer) RedeemResource(ctx context.Context, req *MsgRedeemResource) (*MsgRedeemResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemResource not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyber.resources.v1beta1.Msg/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Convert(ctx, req.(*MsgConvert))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyber.resources.v1beta1.Msg/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateResource(ctx, req.(*MsgCreateResource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RedeemResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRedeemResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RedeemResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyber.resources.v1beta1.Msg/RedeemResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RedeemResource(ctx, req.(*MsgRedeemResource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cyber.resources.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _Msg_Convert_Handler,
		},
		{
			MethodName: "CreateResource",
			Handler:    _Msg_CreateResource_Handler,
		},
		{
			MethodName: "RedeemResource",
			Handler:    _Msg_RedeemResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cyber/resources/v1beta1/tx.proto",
}

func (m *MsgConvert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Resource) > 0 {
		i -= len(m.Resource)
		copy(dAtA[i:], m.Resource)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Resource)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Agent) > 0 {
		i -= len(m.Agent)
		copy(dAtA[i:], m.Agent)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Agent)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRedeemResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemResource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemResource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRedeemResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Agent)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.EndTime != 0 {
		n += 1 + sovTx(uint64(m.EndTime))
	}
	return n
}

func (m *MsgConvertResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateResource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Resource.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRedeemResource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Resource.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRedeemResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Agent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConvertResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConvertResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRedeemResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRedeemResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRedeemResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRedeemResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)