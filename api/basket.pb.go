// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basket.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AddProductRequest struct {
	Basketid             string   `protobuf:"bytes,1,opt,name=basketid,proto3" json:"basketid,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductRequest) Reset()         { *m = AddProductRequest{} }
func (m *AddProductRequest) String() string { return proto.CompactTextString(m) }
func (*AddProductRequest) ProtoMessage()    {}
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{0}
}

func (m *AddProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductRequest.Unmarshal(m, b)
}
func (m *AddProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductRequest.Marshal(b, m, deterministic)
}
func (m *AddProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductRequest.Merge(m, src)
}
func (m *AddProductRequest) XXX_Size() int {
	return xxx_messageInfo_AddProductRequest.Size(m)
}
func (m *AddProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductRequest proto.InternalMessageInfo

func (m *AddProductRequest) GetBasketid() string {
	if m != nil {
		return m.Basketid
	}
	return ""
}

func (m *AddProductRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

type CreateBasketRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBasketRequest) Reset()         { *m = CreateBasketRequest{} }
func (m *CreateBasketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBasketRequest) ProtoMessage()    {}
func (*CreateBasketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{1}
}

func (m *CreateBasketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBasketRequest.Unmarshal(m, b)
}
func (m *CreateBasketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBasketRequest.Marshal(b, m, deterministic)
}
func (m *CreateBasketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBasketRequest.Merge(m, src)
}
func (m *CreateBasketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBasketRequest.Size(m)
}
func (m *CreateBasketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBasketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBasketRequest proto.InternalMessageInfo

type GetAmountRequest struct {
	Basketid             string   `protobuf:"bytes,1,opt,name=basketid,proto3" json:"basketid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAmountRequest) Reset()         { *m = GetAmountRequest{} }
func (m *GetAmountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAmountRequest) ProtoMessage()    {}
func (*GetAmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{2}
}

func (m *GetAmountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAmountRequest.Unmarshal(m, b)
}
func (m *GetAmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAmountRequest.Marshal(b, m, deterministic)
}
func (m *GetAmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAmountRequest.Merge(m, src)
}
func (m *GetAmountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAmountRequest.Size(m)
}
func (m *GetAmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAmountRequest proto.InternalMessageInfo

func (m *GetAmountRequest) GetBasketid() string {
	if m != nil {
		return m.Basketid
	}
	return ""
}

type GetAmountResponse struct {
	Total                string   `protobuf:"bytes,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAmountResponse) Reset()         { *m = GetAmountResponse{} }
func (m *GetAmountResponse) String() string { return proto.CompactTextString(m) }
func (*GetAmountResponse) ProtoMessage()    {}
func (*GetAmountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{3}
}

func (m *GetAmountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAmountResponse.Unmarshal(m, b)
}
func (m *GetAmountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAmountResponse.Marshal(b, m, deterministic)
}
func (m *GetAmountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAmountResponse.Merge(m, src)
}
func (m *GetAmountResponse) XXX_Size() int {
	return xxx_messageInfo_GetAmountResponse.Size(m)
}
func (m *GetAmountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAmountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAmountResponse proto.InternalMessageInfo

func (m *GetAmountResponse) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

type RemoveBasketRequest struct {
	Basketid             string   `protobuf:"bytes,1,opt,name=basketid,proto3" json:"basketid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBasketRequest) Reset()         { *m = RemoveBasketRequest{} }
func (m *RemoveBasketRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveBasketRequest) ProtoMessage()    {}
func (*RemoveBasketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{4}
}

func (m *RemoveBasketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBasketRequest.Unmarshal(m, b)
}
func (m *RemoveBasketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBasketRequest.Marshal(b, m, deterministic)
}
func (m *RemoveBasketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBasketRequest.Merge(m, src)
}
func (m *RemoveBasketRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveBasketRequest.Size(m)
}
func (m *RemoveBasketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBasketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBasketRequest proto.InternalMessageInfo

func (m *RemoveBasketRequest) GetBasketid() string {
	if m != nil {
		return m.Basketid
	}
	return ""
}

type Basket struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items                []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Total                string   `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Basket) Reset()         { *m = Basket{} }
func (m *Basket) String() string { return proto.CompactTextString(m) }
func (*Basket) ProtoMessage()    {}
func (*Basket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{5}
}

func (m *Basket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basket.Unmarshal(m, b)
}
func (m *Basket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basket.Marshal(b, m, deterministic)
}
func (m *Basket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basket.Merge(m, src)
}
func (m *Basket) XXX_Size() int {
	return xxx_messageInfo_Basket.Size(m)
}
func (m *Basket) XXX_DiscardUnknown() {
	xxx_messageInfo_Basket.DiscardUnknown(m)
}

var xxx_messageInfo_Basket proto.InternalMessageInfo

func (m *Basket) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Basket) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Basket) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

type BasketResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasketResponse) Reset()         { *m = BasketResponse{} }
func (m *BasketResponse) String() string { return proto.CompactTextString(m) }
func (*BasketResponse) ProtoMessage()    {}
func (*BasketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e143bae2179266f, []int{6}
}

func (m *BasketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasketResponse.Unmarshal(m, b)
}
func (m *BasketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasketResponse.Marshal(b, m, deterministic)
}
func (m *BasketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketResponse.Merge(m, src)
}
func (m *BasketResponse) XXX_Size() int {
	return xxx_messageInfo_BasketResponse.Size(m)
}
func (m *BasketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BasketResponse proto.InternalMessageInfo

func (m *BasketResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BasketResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*AddProductRequest)(nil), "api.AddProductRequest")
	proto.RegisterType((*CreateBasketRequest)(nil), "api.CreateBasketRequest")
	proto.RegisterType((*GetAmountRequest)(nil), "api.GetAmountRequest")
	proto.RegisterType((*GetAmountResponse)(nil), "api.GetAmountResponse")
	proto.RegisterType((*RemoveBasketRequest)(nil), "api.RemoveBasketRequest")
	proto.RegisterType((*Basket)(nil), "api.Basket")
	proto.RegisterType((*BasketResponse)(nil), "api.BasketResponse")
}

func init() { proto.RegisterFile("basket.proto", fileDescriptor_6e143bae2179266f) }

var fileDescriptor_6e143bae2179266f = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0xb5, 0x60, 0xab, 0x4e, 0xb1, 0xda, 0xc5, 0x36, 0x84, 0x93, 0xd9, 0x93, 0x5e, 0x48, 0xd4,
	0x83, 0x89, 0x31, 0x31, 0xad, 0x26, 0xc6, 0x9b, 0x41, 0x7f, 0x80, 0xc2, 0x1c, 0x88, 0xd2, 0xc5,
	0x65, 0xe9, 0x4f, 0xf8, 0xd3, 0x2e, 0xbb, 0x0b, 0xae, 0xb5, 0x9a, 0xde, 0x76, 0x66, 0xde, 0x9b,
	0x37, 0xbc, 0x07, 0x78, 0x8b, 0xa4, 0x7a, 0x43, 0x11, 0x95, 0x9c, 0x09, 0x46, 0xdc, 0xa4, 0xcc,
	0xe9, 0x13, 0x8c, 0x67, 0x59, 0xf6, 0xcc, 0x59, 0x56, 0xa7, 0x22, 0xc6, 0x8f, 0x1a, 0x2b, 0x41,
	0x42, 0xd8, 0xd7, 0xc8, 0x3c, 0x0b, 0x7a, 0xa7, 0xbd, 0xb3, 0x83, 0xb8, 0xab, 0x49, 0x00, 0x7b,
	0xa5, 0x46, 0x07, 0x8e, 0x1a, 0xb5, 0x25, 0x9d, 0x80, 0x7f, 0xcf, 0x31, 0x11, 0x38, 0x57, 0x58,
	0xb3, 0x8c, 0x46, 0x70, 0xfc, 0x88, 0x62, 0x56, 0xb0, 0x7a, 0xb9, 0x8d, 0x00, 0x3d, 0x87, 0xb1,
	0x85, 0xaf, 0x4a, 0xb6, 0xac, 0x90, 0x9c, 0x40, 0x5f, 0x30, 0x91, 0xbc, 0x1b, 0xb4, 0x2e, 0xe8,
	0x05, 0xf8, 0x31, 0x16, 0x6c, 0xf5, 0x53, 0xf1, 0xdf, 0xed, 0x0f, 0x30, 0xd0, 0x60, 0x32, 0x02,
	0xa7, 0x9b, 0xcb, 0x57, 0x23, 0x91, 0x0b, 0x2c, 0x2a, 0xf9, 0x59, 0x6e, 0x23, 0xa1, 0x8a, 0x6f,
	0x61, 0xd7, 0x16, 0xbe, 0x81, 0x51, 0x2b, 0x69, 0x0e, 0x24, 0xb0, 0x9b, 0xb2, 0x0c, 0xd5, 0xbe,
	0x7e, 0xac, 0xde, 0x0d, 0x17, 0x39, 0x67, 0xdc, 0x18, 0xa5, 0x8b, 0xcb, 0x4f, 0x07, 0x0e, 0x35,
	0xf9, 0x05, 0xf9, 0x2a, 0x4f, 0x91, 0x5c, 0x83, 0x67, 0x1b, 0x47, 0x82, 0x48, 0x26, 0x13, 0x6d,
	0xf0, 0x32, 0x1c, 0xaa, 0x89, 0xee, 0xd1, 0x1d, 0x72, 0x0b, 0x43, 0x19, 0xde, 0x2b, 0x33, 0xbc,
	0xa9, 0x9a, 0xfe, 0x8a, 0x33, 0xf4, 0x2d, 0x56, 0x7b, 0xb0, 0x64, 0xcf, 0xe1, 0x48, 0x1a, 0xad,
	0xdb, 0xda, 0x6e, 0x32, 0x51, 0xc8, 0xf5, 0xb8, 0xc2, 0xe9, 0x7a, 0xbb, 0xdb, 0x71, 0x07, 0x9e,
	0x9d, 0x80, 0x39, 0x7d, 0x43, 0x28, 0x7f, 0x1c, 0xb1, 0x18, 0xa8, 0x7f, 0xf1, 0xea, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x11, 0xd4, 0x84, 0x35, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasketServiceClient interface {
	CreateBasket(ctx context.Context, in *CreateBasketRequest, opts ...grpc.CallOption) (*Basket, error)
	AddToBasket(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*BasketResponse, error)
	GetBasketAmount(ctx context.Context, in *GetAmountRequest, opts ...grpc.CallOption) (*GetAmountResponse, error)
	RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*BasketResponse, error)
}

type basketServiceClient struct {
	cc *grpc.ClientConn
}

func NewBasketServiceClient(cc *grpc.ClientConn) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) CreateBasket(ctx context.Context, in *CreateBasketRequest, opts ...grpc.CallOption) (*Basket, error) {
	out := new(Basket)
	err := c.cc.Invoke(ctx, "/api.BasketService/CreateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) AddToBasket(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*BasketResponse, error) {
	out := new(BasketResponse)
	err := c.cc.Invoke(ctx, "/api.BasketService/AddToBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetBasketAmount(ctx context.Context, in *GetAmountRequest, opts ...grpc.CallOption) (*GetAmountResponse, error) {
	out := new(GetAmountResponse)
	err := c.cc.Invoke(ctx, "/api.BasketService/GetBasketAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*BasketResponse, error) {
	out := new(BasketResponse)
	err := c.cc.Invoke(ctx, "/api.BasketService/RemoveBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
type BasketServiceServer interface {
	CreateBasket(context.Context, *CreateBasketRequest) (*Basket, error)
	AddToBasket(context.Context, *AddProductRequest) (*BasketResponse, error)
	GetBasketAmount(context.Context, *GetAmountRequest) (*GetAmountResponse, error)
	RemoveBasket(context.Context, *RemoveBasketRequest) (*BasketResponse, error)
}

// UnimplementedBasketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBasketServiceServer struct {
}

func (*UnimplementedBasketServiceServer) CreateBasket(ctx context.Context, req *CreateBasketRequest) (*Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBasket not implemented")
}
func (*UnimplementedBasketServiceServer) AddToBasket(ctx context.Context, req *AddProductRequest) (*BasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBasket not implemented")
}
func (*UnimplementedBasketServiceServer) GetBasketAmount(ctx context.Context, req *GetAmountRequest) (*GetAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasketAmount not implemented")
}
func (*UnimplementedBasketServiceServer) RemoveBasket(ctx context.Context, req *RemoveBasketRequest) (*BasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBasket not implemented")
}

func RegisterBasketServiceServer(s *grpc.Server, srv BasketServiceServer) {
	s.RegisterService(&_BasketService_serviceDesc, srv)
}

func _BasketService_CreateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).CreateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BasketService/CreateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).CreateBasket(ctx, req.(*CreateBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_AddToBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).AddToBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BasketService/AddToBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).AddToBasket(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetBasketAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetBasketAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BasketService/GetBasketAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetBasketAmount(ctx, req.(*GetAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_RemoveBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).RemoveBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BasketService/RemoveBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).RemoveBasket(ctx, req.(*RemoveBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBasket",
			Handler:    _BasketService_CreateBasket_Handler,
		},
		{
			MethodName: "AddToBasket",
			Handler:    _BasketService_AddToBasket_Handler,
		},
		{
			MethodName: "GetBasketAmount",
			Handler:    _BasketService_GetBasketAmount_Handler,
		},
		{
			MethodName: "RemoveBasket",
			Handler:    _BasketService_RemoveBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basket.proto",
}