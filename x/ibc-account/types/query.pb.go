// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/account/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type QueryIBCAccountRequest struct {
	// address is the address to query.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryIBCAccountRequest) Reset()         { *m = QueryIBCAccountRequest{} }
func (m *QueryIBCAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountRequest) ProtoMessage()    {}
func (*QueryIBCAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{0}
}
func (m *QueryIBCAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountRequest.Merge(m, src)
}
func (m *QueryIBCAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountRequest proto.InternalMessageInfo

type QueryIBCAccountFromDataRequest struct {
	Port    string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *QueryIBCAccountFromDataRequest) Reset()         { *m = QueryIBCAccountFromDataRequest{} }
func (m *QueryIBCAccountFromDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromDataRequest) ProtoMessage()    {}
func (*QueryIBCAccountFromDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{1}
}
func (m *QueryIBCAccountFromDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromDataRequest.Merge(m, src)
}
func (m *QueryIBCAccountFromDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromDataRequest proto.InternalMessageInfo

type QueryIBCAccountResponse struct {
	// account defines the account of the corresponding address.
	Account *IBCAccount `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *QueryIBCAccountResponse) Reset()         { *m = QueryIBCAccountResponse{} }
func (m *QueryIBCAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountResponse) ProtoMessage()    {}
func (*QueryIBCAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{2}
}
func (m *QueryIBCAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountResponse.Merge(m, src)
}
func (m *QueryIBCAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountResponse proto.InternalMessageInfo

func (m *QueryIBCAccountResponse) GetAccount() *IBCAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIBCAccountRequest)(nil), "ibc.account.QueryIBCAccountRequest")
	proto.RegisterType((*QueryIBCAccountFromDataRequest)(nil), "ibc.account.QueryIBCAccountFromDataRequest")
	proto.RegisterType((*QueryIBCAccountResponse)(nil), "ibc.account.QueryIBCAccountResponse")
}

func init() { proto.RegisterFile("ibc/account/query.proto", fileDescriptor_5738b2fac406b004) }

var fileDescriptor_5738b2fac406b004 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xb5, 0xc3, 0xc7, 0xc1, 0x5e, 0xb7, 0x42, 0xc4, 0x44, 0xc8, 0x87, 0x02, 0x05, 0x12, 0x8a,
	0x57, 0xbe, 0x54, 0x20, 0x1a, 0x0e, 0x84, 0x74, 0x02, 0x8a, 0xbb, 0x92, 0x6e, 0xbd, 0x59, 0x9c,
	0x95, 0xec, 0x1d, 0xdf, 0xee, 0x1a, 0x71, 0xb2, 0xdc, 0x50, 0x51, 0x22, 0xf8, 0x03, 0xf7, 0x3b,
	0xf8, 0x05, 0x94, 0x27, 0xd1, 0x40, 0x87, 0x12, 0x0a, 0x7e, 0x06, 0x5a, 0x7b, 0x2d, 0x9c, 0xcb,
	0x29, 0x4a, 0x95, 0xd9, 0x99, 0x37, 0xef, 0xbd, 0xbc, 0x31, 0x1a, 0x8a, 0x84, 0x11, 0xca, 0x18,
	0x94, 0xd2, 0x90, 0x93, 0x92, 0xab, 0xd3, 0xa8, 0x50, 0x60, 0x00, 0xef, 0x8a, 0x84, 0x45, 0x6e,
	0x30, 0xba, 0x95, 0x42, 0x0a, 0x4d, 0x9f, 0xd8, 0xaa, 0x85, 0x8c, 0xee, 0xa6, 0x00, 0x69, 0xc6,
	0x09, 0x2d, 0x04, 0xa1, 0x52, 0x82, 0xa1, 0x46, 0x80, 0xd4, 0x6e, 0x7a, 0xa7, 0xcf, 0xec, 0x7e,
	0xdb, 0xd1, 0xf8, 0x29, 0xba, 0x7d, 0x64, 0xa5, 0x0e, 0x0f, 0x9e, 0x3f, 0x6b, 0x07, 0xc7, 0xfc,
	0xa4, 0xe4, 0xda, 0xe0, 0x00, 0xed, 0xd0, 0xd9, 0x4c, 0x71, 0xad, 0x03, 0xff, 0x9e, 0xff, 0xf0,
	0xe6, 0x71, 0xf7, 0x7c, 0x72, 0xe3, 0xd3, 0xd9, 0x9e, 0xf7, 0xf7, 0x6c, 0xcf, 0x1b, 0x67, 0x28,
	0xbc, 0xb0, 0xfd, 0x52, 0x41, 0xfe, 0x82, 0x1a, 0xda, 0xb1, 0x60, 0x74, 0xb5, 0x00, 0x65, 0x1c,
	0x45, 0x53, 0x5b, 0x66, 0x36, 0xa7, 0x52, 0xf2, 0x2c, 0x18, 0xb4, 0xcc, 0xee, 0x69, 0xd1, 0x33,
	0x6a, 0x68, 0x70, 0xa5, 0x45, 0xdb, 0xba, 0xa7, 0xf6, 0x1a, 0x0d, 0xd7, 0xbc, 0xea, 0x02, 0xa4,
	0xe6, 0x38, 0x46, 0x3b, 0xee, 0x7f, 0x35, 0x4a, 0xbb, 0xfb, 0xc3, 0xa8, 0x17, 0x5a, 0xd4, 0xdb,
	0xe8, 0x70, 0xfb, 0xbf, 0x06, 0xe8, 0x5a, 0x43, 0x87, 0xbf, 0xf8, 0x08, 0xfd, 0x47, 0xe0, 0xfb,
	0x2b, 0xab, 0x97, 0xa7, 0x33, 0x7a, 0xb0, 0x19, 0xd4, 0xda, 0x1a, 0x3f, 0xfe, 0xf8, 0xe3, 0xcf,
	0xd7, 0xc1, 0x14, 0xc7, 0x84, 0x81, 0xce, 0x41, 0x13, 0x91, 0xb0, 0x49, 0x77, 0x88, 0xf7, 0x71,
	0xc2, 0x0d, 0x8d, 0x57, 0x7a, 0x95, 0xcb, 0xb8, 0xc6, 0xdf, 0x7c, 0x84, 0xd7, 0x63, 0xc5, 0x8f,
	0x36, 0xe9, 0x5e, 0x08, 0x7f, 0x4b, 0x93, 0x47, 0x8d, 0xc9, 0x57, 0xf8, 0x70, 0x4b, 0x93, 0x93,
	0x77, 0x0a, 0xf2, 0x89, 0xbd, 0x0d, 0xa9, 0xec, 0x3d, 0x6b, 0x52, 0xb9, 0xfb, 0xd5, 0xa4, 0xb2,
	0xed, 0xfa, 0xe0, 0xcd, 0xf7, 0x45, 0xe8, 0x9f, 0x2f, 0x42, 0xff, 0xf7, 0x22, 0xf4, 0x3f, 0x2f,
	0x43, 0xef, 0x7c, 0x19, 0x7a, 0x3f, 0x97, 0xa1, 0xf7, 0x76, 0x9a, 0x0a, 0x33, 0x2f, 0x93, 0x88,
	0x41, 0x4e, 0x84, 0x34, 0x5c, 0xb1, 0x39, 0x15, 0x32, 0xe1, 0x2a, 0x13, 0x92, 0x08, 0x46, 0xc9,
	0x87, 0x15, 0x79, 0x73, 0x5a, 0x70, 0x9d, 0x5c, 0x6f, 0xbe, 0xd5, 0xe9, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x62, 0x70, 0x7a, 0xcd, 0x22, 0x03, 0x00, 0x00,
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
	IBCAccount(ctx context.Context, in *QueryIBCAccountRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error)
	IBCAccountFromData(ctx context.Context, in *QueryIBCAccountFromDataRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IBCAccount(ctx context.Context, in *QueryIBCAccountRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error) {
	out := new(QueryIBCAccountResponse)
	err := c.cc.Invoke(ctx, "/ibc.account.Query/IBCAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IBCAccountFromData(ctx context.Context, in *QueryIBCAccountFromDataRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error) {
	out := new(QueryIBCAccountResponse)
	err := c.cc.Invoke(ctx, "/ibc.account.Query/IBCAccountFromData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	IBCAccount(context.Context, *QueryIBCAccountRequest) (*QueryIBCAccountResponse, error)
	IBCAccountFromData(context.Context, *QueryIBCAccountFromDataRequest) (*QueryIBCAccountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IBCAccount(ctx context.Context, req *QueryIBCAccountRequest) (*QueryIBCAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccount not implemented")
}
func (*UnimplementedQueryServer) IBCAccountFromData(ctx context.Context, req *QueryIBCAccountFromDataRequest) (*QueryIBCAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccountFromData not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IBCAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.account.Query/IBCAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccount(ctx, req.(*QueryIBCAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IBCAccountFromData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountFromDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccountFromData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.account.Query/IBCAccountFromData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccountFromData(ctx, req.(*QueryIBCAccountFromDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.account.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IBCAccount",
			Handler:    _Query_IBCAccount_Handler,
		},
		{
			MethodName: "IBCAccountFromData",
			Handler:    _Query_IBCAccountFromData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/account/query.proto",
}

func (m *QueryIBCAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountFromDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryIBCAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountFromDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
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
func (m *QueryIBCAccountRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIBCAccountFromDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountFromDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
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
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIBCAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			if m.Account == nil {
				m.Account = &IBCAccount{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
