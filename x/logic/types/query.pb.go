// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logic/v1beta/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryServiceParamsRequest is request type for the QueryService/Params RPC method.
type QueryServiceParamsRequest struct {
}

func (m *QueryServiceParamsRequest) Reset()         { *m = QueryServiceParamsRequest{} }
func (m *QueryServiceParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryServiceParamsRequest) ProtoMessage()    {}
func (*QueryServiceParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{0}
}
func (m *QueryServiceParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceParamsRequest.Merge(m, src)
}
func (m *QueryServiceParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceParamsRequest proto.InternalMessageInfo

// QueryServiceParamsResponse is response type for the QueryService/Params RPC method.
type QueryServiceParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params",omitempty`
}

func (m *QueryServiceParamsResponse) Reset()         { *m = QueryServiceParamsResponse{} }
func (m *QueryServiceParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryServiceParamsResponse) ProtoMessage()    {}
func (*QueryServiceParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{1}
}
func (m *QueryServiceParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceParamsResponse.Merge(m, src)
}
func (m *QueryServiceParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceParamsResponse proto.InternalMessageInfo

func (m *QueryServiceParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryServiceAskRequest is request type for the QueryService/Ask RPC method.
type QueryServiceAskRequest struct {
	// program is the logic program to be queried.
	Program string `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty" yaml:"program",omitempty`
	// query is the query string to be executed.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty" yaml:"query",omitempty`
}

func (m *QueryServiceAskRequest) Reset()         { *m = QueryServiceAskRequest{} }
func (m *QueryServiceAskRequest) String() string { return proto.CompactTextString(m) }
func (*QueryServiceAskRequest) ProtoMessage()    {}
func (*QueryServiceAskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{2}
}
func (m *QueryServiceAskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceAskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceAskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceAskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceAskRequest.Merge(m, src)
}
func (m *QueryServiceAskRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceAskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceAskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceAskRequest proto.InternalMessageInfo

func (m *QueryServiceAskRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *QueryServiceAskRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// QueryServiceAskResponse is response type for the QueryService/Ask RPC method.
type QueryServiceAskResponse struct {
	// height is the block height at which the query was executed.
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty" yaml:"height",omitempty`
	// gas_used is the amount of gas used to execute the query.
	GasUsed uint64 `protobuf:"varint,2,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty" yaml:"gas_used",omitempty`
	// answer is the answer to the query.
	Answer *Answer `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty" yaml:"answer",omitempty`
}

func (m *QueryServiceAskResponse) Reset()         { *m = QueryServiceAskResponse{} }
func (m *QueryServiceAskResponse) String() string { return proto.CompactTextString(m) }
func (*QueryServiceAskResponse) ProtoMessage()    {}
func (*QueryServiceAskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{3}
}
func (m *QueryServiceAskResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceAskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceAskResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceAskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceAskResponse.Merge(m, src)
}
func (m *QueryServiceAskResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceAskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceAskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceAskResponse proto.InternalMessageInfo

func (m *QueryServiceAskResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *QueryServiceAskResponse) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *QueryServiceAskResponse) GetAnswer() *Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryServiceParamsRequest)(nil), "logic.v1beta.QueryServiceParamsRequest")
	proto.RegisterType((*QueryServiceParamsResponse)(nil), "logic.v1beta.QueryServiceParamsResponse")
	proto.RegisterType((*QueryServiceAskRequest)(nil), "logic.v1beta.QueryServiceAskRequest")
	proto.RegisterType((*QueryServiceAskResponse)(nil), "logic.v1beta.QueryServiceAskResponse")
}

func init() { proto.RegisterFile("logic/v1beta/query.proto", fileDescriptor_2a2e6b98fc71dc0f) }

var fileDescriptor_2a2e6b98fc71dc0f = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x8d, 0xd3, 0x60, 0xe0, 0xe8, 0x74, 0x2a, 0x8d, 0xe3, 0xb6, 0x76, 0x65, 0x40, 0x74, 0x40,
	0x36, 0xb4, 0x08, 0xa1, 0x4a, 0x0c, 0xf1, 0xc6, 0x80, 0x04, 0x45, 0x2c, 0x2c, 0xe8, 0x92, 0x9c,
	0x2e, 0x96, 0x63, 0xdf, 0xc5, 0x77, 0x29, 0x64, 0x40, 0x42, 0x4c, 0x8c, 0x48, 0x2c, 0x8c, 0xfc,
	0x9c, 0x8e, 0x95, 0x58, 0x60, 0xb1, 0x50, 0xc2, 0x2f, 0xc8, 0x2f, 0xa8, 0x7a, 0xdf, 0x45, 0xb5,
	0xd5, 0x56, 0x59, 0xa2, 0xe8, 0x7b, 0xef, 0x7d, 0xef, 0xbd, 0xf3, 0x1d, 0x72, 0x46, 0x9c, 0x25,
	0xfd, 0xe8, 0xf8, 0x49, 0x8f, 0x2a, 0x12, 0x8d, 0x27, 0xb4, 0x98, 0x86, 0xa2, 0xe0, 0x8a, 0xe3,
	0x75, 0x8d, 0x84, 0x80, 0xb8, 0x1b, 0x8c, 0x33, 0xae, 0x81, 0xe8, 0xfc, 0x1f, 0x70, 0xdc, 0x6d,
	0xc6, 0x39, 0x1b, 0xd1, 0x88, 0x88, 0x24, 0x22, 0x79, 0xce, 0x15, 0x51, 0x09, 0xcf, 0xa5, 0x41,
	0x3b, 0xb5, 0xdd, 0x82, 0x14, 0x24, 0x5b, 0x42, 0x75, 0x5b, 0x35, 0x15, 0xd4, 0x20, 0xc1, 0x16,
	0xea, 0xbc, 0x39, 0x4f, 0xf1, 0x96, 0x16, 0xc7, 0x49, 0x9f, 0xbe, 0xd6, 0xaa, 0x23, 0x3a, 0x9e,
	0x50, 0xa9, 0x82, 0x14, 0xb9, 0x57, 0x81, 0x52, 0xf0, 0x5c, 0x52, 0xfc, 0x0a, 0xd9, 0x60, 0xe2,
	0x58, 0xbb, 0xd6, 0xde, 0x9d, 0xfd, 0x8d, 0xb0, 0x5a, 0x21, 0x04, 0x76, 0xec, 0x9f, 0x94, 0x7e,
	0x63, 0x51, 0xfa, 0xed, 0x29, 0xc9, 0x46, 0x87, 0x01, 0x28, 0x82, 0x47, 0x3c, 0x4b, 0x14, 0xcd,
	0x84, 0x9a, 0x1e, 0x99, 0x25, 0xc1, 0x37, 0x0b, 0x6d, 0x56, 0xdd, 0xba, 0x32, 0x35, 0x39, 0xf0,
	0x33, 0x74, 0x53, 0x14, 0x9c, 0x15, 0x24, 0xd3, 0x56, 0xb7, 0xe3, 0xed, 0x45, 0xe9, 0x3b, 0x66,
	0x21, 0x00, 0xd5, 0x8d, 0x4b, 0x32, 0x7e, 0x8c, 0x6e, 0xe8, 0x23, 0x76, 0x9a, 0x5a, 0xe5, 0x2e,
	0x4a, 0x7f, 0x13, 0x54, 0x7a, 0x5c, 0xd5, 0x00, 0xf1, 0xb0, 0xf5, 0xf3, 0x97, 0x6f, 0x05, 0x7f,
	0x2d, 0xd4, 0xbe, 0x14, 0xc5, 0xb4, 0x3e, 0x40, 0xf6, 0x90, 0x26, 0x6c, 0xa8, 0x74, 0x94, 0x56,
	0xbc, 0x75, 0xd1, 0x0d, 0xe6, 0xb5, 0x6e, 0x30, 0xc2, 0xcf, 0xd1, 0x2d, 0x46, 0xe4, 0x87, 0x89,
	0xa4, 0x03, 0x9d, 0xa5, 0x15, 0xef, 0x2c, 0x4a, 0xbf, 0x03, 0xb2, 0x25, 0x52, 0xab, 0xc0, 0x88,
	0x7c, 0x27, 0xe9, 0x00, 0xbf, 0x44, 0x36, 0xc9, 0xe5, 0x47, 0x5a, 0x38, 0x6b, 0x57, 0x1d, 0x72,
	0x57, 0x63, 0xd5, 0x10, 0xc0, 0xae, 0x85, 0x80, 0x11, 0x74, 0xdb, 0xff, 0xd2, 0x44, 0xeb, 0xd5,
	0x6e, 0xf8, 0x33, 0xb2, 0xe1, 0x53, 0xe1, 0x87, 0xf5, 0xdd, 0xd7, 0xde, 0x0b, 0x77, 0x6f, 0x35,
	0x11, 0x4e, 0x2b, 0xd8, 0xfd, 0xfa, 0xfb, 0xff, 0x8f, 0xa6, 0x8b, 0x9d, 0x88, 0xa7, 0xe2, 0xa9,
	0xfe, 0x19, 0x44, 0x70, 0x19, 0xe1, 0xb3, 0xe3, 0x31, 0x5a, 0xeb, 0xca, 0x14, 0xdf, 0xbf, 0x7e,
	0xe5, 0xc5, 0x45, 0x70, 0x1f, 0xac, 0x60, 0x19, 0xd7, 0x1d, 0xed, 0xda, 0xc6, 0x77, 0x2f, 0xbb,
	0x12, 0x99, 0xc6, 0x2f, 0x4e, 0x66, 0x9e, 0x75, 0x3a, 0xf3, 0xac, 0x7f, 0x33, 0xcf, 0xfa, 0x3e,
	0xf7, 0x1a, 0xa7, 0x73, 0xaf, 0xf1, 0x67, 0xee, 0x35, 0xde, 0xdf, 0x63, 0x89, 0x1a, 0x4e, 0x7a,
	0x61, 0x9f, 0x67, 0x55, 0xe9, 0x27, 0x23, 0xd6, 0x0f, 0xa7, 0x67, 0xeb, 0x97, 0x73, 0x70, 0x16,
	0x00, 0x00, 0xff, 0xff, 0x85, 0x0a, 0x35, 0xed, 0xcc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// Params queries all parameters for the logic module.
	Params(ctx context.Context, in *QueryServiceParamsRequest, opts ...grpc.CallOption) (*QueryServiceParamsResponse, error)
	// Ask executes a logic query and returns the solutions found.
	// Since the query is without any side-effect, the query is not executed in the context of a transaction and no fee
	// is charged for this, but the execution is constrained by the current limits configured in the module.
	Ask(ctx context.Context, in *QueryServiceAskRequest, opts ...grpc.CallOption) (*QueryServiceAskResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Params(ctx context.Context, in *QueryServiceParamsRequest, opts ...grpc.CallOption) (*QueryServiceParamsResponse, error) {
	out := new(QueryServiceParamsResponse)
	err := c.cc.Invoke(ctx, "/logic.v1beta.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Ask(ctx context.Context, in *QueryServiceAskRequest, opts ...grpc.CallOption) (*QueryServiceAskResponse, error) {
	out := new(QueryServiceAskResponse)
	err := c.cc.Invoke(ctx, "/logic.v1beta.QueryService/Ask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// Params queries all parameters for the logic module.
	Params(context.Context, *QueryServiceParamsRequest) (*QueryServiceParamsResponse, error)
	// Ask executes a logic query and returns the solutions found.
	// Since the query is without any side-effect, the query is not executed in the context of a transaction and no fee
	// is charged for this, but the execution is constrained by the current limits configured in the module.
	Ask(context.Context, *QueryServiceAskRequest) (*QueryServiceAskResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *QueryServiceParamsRequest) (*QueryServiceParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServiceServer) Ask(ctx context.Context, req *QueryServiceAskRequest) (*QueryServiceAskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ask not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.v1beta.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*QueryServiceParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Ask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceAskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Ask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.v1beta.QueryService/Ask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Ask(ctx, req.(*QueryServiceAskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logic.v1beta.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
		{
			MethodName: "Ask",
			Handler:    _QueryService_Ask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic/v1beta/query.proto",
}

func (m *QueryServiceParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryServiceParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryServiceAskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceAskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceAskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Program) > 0 {
		i -= len(m.Program)
		copy(dAtA[i:], m.Program)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Program)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryServiceAskResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceAskResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceAskResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Answer != nil {
		{
			size, err := m.Answer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GasUsed != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
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
func (m *QueryServiceParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryServiceParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryServiceAskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Program)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryServiceAskResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if m.GasUsed != 0 {
		n += 1 + sovQuery(uint64(m.GasUsed))
	}
	if m.Answer != nil {
		l = m.Answer.Size()
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
func (m *QueryServiceParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryServiceParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryServiceParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryServiceParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryServiceAskRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryServiceAskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceAskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Program", wireType)
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
			m.Program = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
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
			m.Query = string(dAtA[iNdEx:postIndex])
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
func (m *QueryServiceAskResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryServiceAskResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceAskResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
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
			if m.Answer == nil {
				m.Answer = &Answer{}
			}
			if err := m.Answer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
