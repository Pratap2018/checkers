// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// this line is used by starport scaffolding # 3
type QueryGetStoredGameRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetStoredGameRequest) Reset()         { *m = QueryGetStoredGameRequest{} }
func (m *QueryGetStoredGameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetStoredGameRequest) ProtoMessage()    {}
func (*QueryGetStoredGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{0}
}
func (m *QueryGetStoredGameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetStoredGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetStoredGameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetStoredGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetStoredGameRequest.Merge(m, src)
}
func (m *QueryGetStoredGameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetStoredGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetStoredGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetStoredGameRequest proto.InternalMessageInfo

func (m *QueryGetStoredGameRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetStoredGameResponse struct {
	StoredGame *StoredGame `protobuf:"bytes,1,opt,name=StoredGame,proto3" json:"StoredGame,omitempty"`
}

func (m *QueryGetStoredGameResponse) Reset()         { *m = QueryGetStoredGameResponse{} }
func (m *QueryGetStoredGameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetStoredGameResponse) ProtoMessage()    {}
func (*QueryGetStoredGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{1}
}
func (m *QueryGetStoredGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetStoredGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetStoredGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetStoredGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetStoredGameResponse.Merge(m, src)
}
func (m *QueryGetStoredGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetStoredGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetStoredGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetStoredGameResponse proto.InternalMessageInfo

func (m *QueryGetStoredGameResponse) GetStoredGame() *StoredGame {
	if m != nil {
		return m.StoredGame
	}
	return nil
}

type QueryAllStoredGameRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllStoredGameRequest) Reset()         { *m = QueryAllStoredGameRequest{} }
func (m *QueryAllStoredGameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllStoredGameRequest) ProtoMessage()    {}
func (*QueryAllStoredGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{2}
}
func (m *QueryAllStoredGameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllStoredGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllStoredGameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllStoredGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllStoredGameRequest.Merge(m, src)
}
func (m *QueryAllStoredGameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllStoredGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllStoredGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllStoredGameRequest proto.InternalMessageInfo

func (m *QueryAllStoredGameRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllStoredGameResponse struct {
	StoredGame []*StoredGame       `protobuf:"bytes,1,rep,name=StoredGame,proto3" json:"StoredGame,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllStoredGameResponse) Reset()         { *m = QueryAllStoredGameResponse{} }
func (m *QueryAllStoredGameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllStoredGameResponse) ProtoMessage()    {}
func (*QueryAllStoredGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{3}
}
func (m *QueryAllStoredGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllStoredGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllStoredGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllStoredGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllStoredGameResponse.Merge(m, src)
}
func (m *QueryAllStoredGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllStoredGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllStoredGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllStoredGameResponse proto.InternalMessageInfo

func (m *QueryAllStoredGameResponse) GetStoredGame() []*StoredGame {
	if m != nil {
		return m.StoredGame
	}
	return nil
}

func (m *QueryAllStoredGameResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGetNextGameRequest struct {
}

func (m *QueryGetNextGameRequest) Reset()         { *m = QueryGetNextGameRequest{} }
func (m *QueryGetNextGameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetNextGameRequest) ProtoMessage()    {}
func (*QueryGetNextGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{4}
}
func (m *QueryGetNextGameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNextGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNextGameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNextGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNextGameRequest.Merge(m, src)
}
func (m *QueryGetNextGameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNextGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNextGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNextGameRequest proto.InternalMessageInfo

type QueryGetNextGameResponse struct {
	NextGame *NextGame `protobuf:"bytes,1,opt,name=NextGame,proto3" json:"NextGame,omitempty"`
}

func (m *QueryGetNextGameResponse) Reset()         { *m = QueryGetNextGameResponse{} }
func (m *QueryGetNextGameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetNextGameResponse) ProtoMessage()    {}
func (*QueryGetNextGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c482788bba85e7a, []int{5}
}
func (m *QueryGetNextGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNextGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNextGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNextGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNextGameResponse.Merge(m, src)
}
func (m *QueryGetNextGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNextGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNextGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNextGameResponse proto.InternalMessageInfo

func (m *QueryGetNextGameResponse) GetNextGame() *NextGame {
	if m != nil {
		return m.NextGame
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetStoredGameRequest)(nil), "Pratap2018.checkers.checkers.QueryGetStoredGameRequest")
	proto.RegisterType((*QueryGetStoredGameResponse)(nil), "Pratap2018.checkers.checkers.QueryGetStoredGameResponse")
	proto.RegisterType((*QueryAllStoredGameRequest)(nil), "Pratap2018.checkers.checkers.QueryAllStoredGameRequest")
	proto.RegisterType((*QueryAllStoredGameResponse)(nil), "Pratap2018.checkers.checkers.QueryAllStoredGameResponse")
	proto.RegisterType((*QueryGetNextGameRequest)(nil), "Pratap2018.checkers.checkers.QueryGetNextGameRequest")
	proto.RegisterType((*QueryGetNextGameResponse)(nil), "Pratap2018.checkers.checkers.QueryGetNextGameResponse")
}

func init() { proto.RegisterFile("checkers/query.proto", fileDescriptor_3c482788bba85e7a) }

var fileDescriptor_3c482788bba85e7a = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0x85, 0x22, 0x38, 0xc4, 0x72, 0xaa, 0x44, 0x6a, 0x55, 0x16, 0xf2, 0x50, 0xa2,
	0x0e, 0x77, 0x49, 0x10, 0x90, 0xb5, 0x1d, 0x08, 0x2c, 0xa8, 0x84, 0x8d, 0x01, 0x74, 0x71, 0x5f,
	0x5c, 0x0b, 0xe7, 0xce, 0xf5, 0x5d, 0x50, 0x2a, 0xc4, 0xc2, 0x27, 0x40, 0xe2, 0x3b, 0x20, 0x31,
	0xb1, 0xb0, 0xf0, 0x0d, 0x18, 0x2b, 0xb1, 0x30, 0xa2, 0x84, 0x0f, 0x82, 0x72, 0x3e, 0xff, 0x69,
	0xea, 0x1a, 0x87, 0xcd, 0xf1, 0x7b, 0xcf, 0xf3, 0xfe, 0xde, 0xe7, 0xde, 0x18, 0x6f, 0xf9, 0xc7,
	0xe0, 0xbf, 0x81, 0x44, 0xb1, 0x93, 0x29, 0x24, 0xa7, 0x34, 0x4e, 0xa4, 0x96, 0x64, 0xe7, 0x30,
	0xe1, 0x9a, 0xc7, 0xfd, 0x6e, 0x6f, 0x40, 0xb3, 0x03, 0xf9, 0x83, 0xb3, 0x13, 0x48, 0x19, 0x44,
	0xc0, 0x78, 0x1c, 0x32, 0x2e, 0x84, 0xd4, 0x5c, 0x87, 0x52, 0xa8, 0x54, 0xeb, 0xec, 0xf9, 0x52,
	0x4d, 0xa4, 0x62, 0x63, 0xae, 0x20, 0x35, 0x65, 0x6f, 0x7b, 0x63, 0xd0, 0xbc, 0xc7, 0x62, 0x1e,
	0x84, 0xc2, 0x1c, 0xb6, 0x67, 0x9d, 0xbc, 0xbb, 0xd2, 0x32, 0x81, 0xa3, 0x57, 0x01, 0x9f, 0x80,
	0xad, 0xb5, 0xf3, 0x9a, 0x80, 0x99, 0x2e, 0x55, 0xbc, 0x1e, 0xde, 0x7e, 0xb6, 0xf4, 0x1d, 0x82,
	0x7e, 0x6e, 0x64, 0x43, 0x3e, 0x81, 0x11, 0x9c, 0x4c, 0x41, 0x69, 0xb2, 0x85, 0x37, 0x43, 0x71,
	0x04, 0xb3, 0x36, 0xba, 0x83, 0x3a, 0x37, 0x46, 0xe9, 0x0f, 0xef, 0x35, 0x76, 0xaa, 0x24, 0x2a,
	0x96, 0x42, 0x01, 0x79, 0x8c, 0x71, 0xf1, 0xd6, 0x08, 0x6f, 0xf6, 0x3b, 0xb4, 0x2e, 0x03, 0x5a,
	0x72, 0x29, 0x69, 0x3d, 0xdf, 0xa2, 0xed, 0x47, 0xd1, 0x45, 0xb4, 0x47, 0x18, 0x17, 0x09, 0xd8,
	0x36, 0xbb, 0x34, 0x8d, 0x8b, 0x2e, 0xe3, 0xa2, 0xe9, 0x1d, 0xd8, 0xb8, 0xe8, 0x21, 0x0f, 0x32,
	0xed, 0xa8, 0xa4, 0xf4, 0xbe, 0x22, 0x3b, 0xcd, 0x4a, 0x97, 0x4b, 0xa6, 0xb9, 0xf2, 0xbf, 0xd3,
	0x90, 0xe1, 0x39, 0xe0, 0x0d, 0x03, 0x7c, 0xf7, 0x9f, 0xc0, 0x29, 0xc6, 0x39, 0xe2, 0x6d, 0x7c,
	0x3b, 0x8b, 0xff, 0x29, 0xcc, 0x74, 0x29, 0x14, 0xef, 0x25, 0x6e, 0x5f, 0x2c, 0xd9, 0x49, 0x0e,
	0xf0, 0xf5, 0xec, 0x5d, 0x1e, 0x57, 0xed, 0x1c, 0xb9, 0x43, 0xae, 0xeb, 0x7f, 0xbe, 0x8a, 0x37,
	0x4d, 0x03, 0xf2, 0x1d, 0x95, 0x83, 0x21, 0x0f, 0xeb, 0xad, 0x2e, 0xdd, 0x30, 0x67, 0xb0, 0xbe,
	0x30, 0x9d, 0xc7, 0x1b, 0x7c, 0xf8, 0xf9, 0xe7, 0xd3, 0x46, 0x9f, 0x74, 0x59, 0xe1, 0xc0, 0xf2,
	0x35, 0x5f, 0xf9, 0x2f, 0x2c, 0x95, 0xec, 0x9d, 0x59, 0xdf, 0xf7, 0xe4, 0x1b, 0xc2, 0xb7, 0x0a,
	0xc3, 0xfd, 0x28, 0x6a, 0x84, 0x5f, 0xb5, 0x85, 0x8d, 0xf0, 0x2b, 0x17, 0xcb, 0xeb, 0x1a, 0xfc,
	0x3d, 0xd2, 0x69, 0x8a, 0x4f, 0xbe, 0xa0, 0xe2, 0x06, 0xc9, 0xfd, 0x66, 0xb9, 0xad, 0x2c, 0x88,
	0xf3, 0x60, 0x5d, 0x99, 0xa5, 0xa5, 0x86, 0xb6, 0x43, 0x76, 0xeb, 0x69, 0x85, 0xd5, 0x1d, 0x3c,
	0xf9, 0x31, 0x77, 0xd1, 0xd9, 0xdc, 0x45, 0xbf, 0xe7, 0x2e, 0xfa, 0xb8, 0x70, 0x5b, 0x67, 0x0b,
	0xb7, 0xf5, 0x6b, 0xe1, 0xb6, 0x5e, 0xb0, 0x20, 0xd4, 0xc7, 0xd3, 0x31, 0xf5, 0xe5, 0xa4, 0xd2,
	0x6b, 0x56, 0x3c, 0xea, 0xd3, 0x18, 0xd4, 0xf8, 0x9a, 0xf9, 0x4e, 0xdd, 0xfb, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0x48, 0x02, 0x3f, 0x5d, 0x05, 0x00, 0x00,
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
	// Queries a storedGame by index.
	StoredGame(ctx context.Context, in *QueryGetStoredGameRequest, opts ...grpc.CallOption) (*QueryGetStoredGameResponse, error)
	// Queries a list of storedGame items.
	StoredGameAll(ctx context.Context, in *QueryAllStoredGameRequest, opts ...grpc.CallOption) (*QueryAllStoredGameResponse, error)
	// Queries a nextGame by index.
	NextGame(ctx context.Context, in *QueryGetNextGameRequest, opts ...grpc.CallOption) (*QueryGetNextGameResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) StoredGame(ctx context.Context, in *QueryGetStoredGameRequest, opts ...grpc.CallOption) (*QueryGetStoredGameResponse, error) {
	out := new(QueryGetStoredGameResponse)
	err := c.cc.Invoke(ctx, "/Pratap2018.checkers.checkers.Query/StoredGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StoredGameAll(ctx context.Context, in *QueryAllStoredGameRequest, opts ...grpc.CallOption) (*QueryAllStoredGameResponse, error) {
	out := new(QueryAllStoredGameResponse)
	err := c.cc.Invoke(ctx, "/Pratap2018.checkers.checkers.Query/StoredGameAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NextGame(ctx context.Context, in *QueryGetNextGameRequest, opts ...grpc.CallOption) (*QueryGetNextGameResponse, error) {
	out := new(QueryGetNextGameResponse)
	err := c.cc.Invoke(ctx, "/Pratap2018.checkers.checkers.Query/NextGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a storedGame by index.
	StoredGame(context.Context, *QueryGetStoredGameRequest) (*QueryGetStoredGameResponse, error)
	// Queries a list of storedGame items.
	StoredGameAll(context.Context, *QueryAllStoredGameRequest) (*QueryAllStoredGameResponse, error)
	// Queries a nextGame by index.
	NextGame(context.Context, *QueryGetNextGameRequest) (*QueryGetNextGameResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) StoredGame(ctx context.Context, req *QueryGetStoredGameRequest) (*QueryGetStoredGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoredGame not implemented")
}
func (*UnimplementedQueryServer) StoredGameAll(ctx context.Context, req *QueryAllStoredGameRequest) (*QueryAllStoredGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoredGameAll not implemented")
}
func (*UnimplementedQueryServer) NextGame(ctx context.Context, req *QueryGetNextGameRequest) (*QueryGetNextGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextGame not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_StoredGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStoredGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StoredGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pratap2018.checkers.checkers.Query/StoredGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StoredGame(ctx, req.(*QueryGetStoredGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StoredGameAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllStoredGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StoredGameAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pratap2018.checkers.checkers.Query/StoredGameAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StoredGameAll(ctx, req.(*QueryAllStoredGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NextGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetNextGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NextGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pratap2018.checkers.checkers.Query/NextGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NextGame(ctx, req.(*QueryGetNextGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Pratap2018.checkers.checkers.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoredGame",
			Handler:    _Query_StoredGame_Handler,
		},
		{
			MethodName: "StoredGameAll",
			Handler:    _Query_StoredGameAll_Handler,
		},
		{
			MethodName: "NextGame",
			Handler:    _Query_NextGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkers/query.proto",
}

func (m *QueryGetStoredGameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetStoredGameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetStoredGameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetStoredGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetStoredGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetStoredGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StoredGame != nil {
		{
			size, err := m.StoredGame.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllStoredGameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllStoredGameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllStoredGameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllStoredGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllStoredGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllStoredGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StoredGame) > 0 {
		for iNdEx := len(m.StoredGame) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredGame[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetNextGameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNextGameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNextGameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetNextGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNextGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNextGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextGame != nil {
		{
			size, err := m.NextGame.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetStoredGameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetStoredGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoredGame != nil {
		l = m.StoredGame.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllStoredGameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllStoredGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StoredGame) > 0 {
		for _, e := range m.StoredGame {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetNextGameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetNextGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextGame != nil {
		l = m.NextGame.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetStoredGameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetStoredGameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetStoredGameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetStoredGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetStoredGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetStoredGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredGame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoredGame == nil {
				m.StoredGame = &StoredGame{}
			}
			if err := m.StoredGame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllStoredGameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllStoredGameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllStoredGameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllStoredGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllStoredGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllStoredGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredGame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredGame = append(m.StoredGame, &StoredGame{})
			if err := m.StoredGame[len(m.StoredGame)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetNextGameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetNextGameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNextGameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetNextGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetNextGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNextGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextGame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextGame == nil {
				m.NextGame = &NextGame{}
			}
			if err := m.NextGame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
