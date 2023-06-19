// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgSubmitNewJob struct {
	Creator                 string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	HexSmartContractAddress string `protobuf:"bytes,2,opt,name=hexSmartContractAddress,proto3" json:"hexSmartContractAddress,omitempty"`
	HexPayload              string `protobuf:"bytes,3,opt,name=hexPayload,proto3" json:"hexPayload,omitempty"`
	Abi                     string `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty"`
	Method                  string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	ChainType               string `protobuf:"bytes,6,opt,name=chainType,proto3" json:"chainType,omitempty"`
	ChainReferenceID        string `protobuf:"bytes,7,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
}

func (m *MsgSubmitNewJob) Reset()         { *m = MsgSubmitNewJob{} }
func (m *MsgSubmitNewJob) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitNewJob) ProtoMessage()    {}
func (*MsgSubmitNewJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{0}
}
func (m *MsgSubmitNewJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitNewJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitNewJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitNewJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitNewJob.Merge(m, src)
}
func (m *MsgSubmitNewJob) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitNewJob) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitNewJob.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitNewJob proto.InternalMessageInfo

func (m *MsgSubmitNewJob) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitNewJob) GetHexSmartContractAddress() string {
	if m != nil {
		return m.HexSmartContractAddress
	}
	return ""
}

func (m *MsgSubmitNewJob) GetHexPayload() string {
	if m != nil {
		return m.HexPayload
	}
	return ""
}

func (m *MsgSubmitNewJob) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *MsgSubmitNewJob) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *MsgSubmitNewJob) GetChainType() string {
	if m != nil {
		return m.ChainType
	}
	return ""
}

func (m *MsgSubmitNewJob) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

type MsgSubmitNewJobResponse struct {
}

func (m *MsgSubmitNewJobResponse) Reset()         { *m = MsgSubmitNewJobResponse{} }
func (m *MsgSubmitNewJobResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitNewJobResponse) ProtoMessage()    {}
func (*MsgSubmitNewJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{1}
}
func (m *MsgSubmitNewJobResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitNewJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitNewJobResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitNewJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitNewJobResponse.Merge(m, src)
}
func (m *MsgSubmitNewJobResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitNewJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitNewJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitNewJobResponse proto.InternalMessageInfo

type MsgUploadNewSmartContractTemp struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Abi              string `protobuf:"bytes,2,opt,name=abi,proto3" json:"abi,omitempty"`
	Bytecode         string `protobuf:"bytes,3,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	ConstructorInput string `protobuf:"bytes,4,opt,name=constructorInput,proto3" json:"constructorInput,omitempty"`
	ChainReferenceID string `protobuf:"bytes,5,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
}

func (m *MsgUploadNewSmartContractTemp) Reset()         { *m = MsgUploadNewSmartContractTemp{} }
func (m *MsgUploadNewSmartContractTemp) String() string { return proto.CompactTextString(m) }
func (*MsgUploadNewSmartContractTemp) ProtoMessage()    {}
func (*MsgUploadNewSmartContractTemp) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{2}
}
func (m *MsgUploadNewSmartContractTemp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUploadNewSmartContractTemp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUploadNewSmartContractTemp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUploadNewSmartContractTemp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUploadNewSmartContractTemp.Merge(m, src)
}
func (m *MsgUploadNewSmartContractTemp) XXX_Size() int {
	return m.Size()
}
func (m *MsgUploadNewSmartContractTemp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUploadNewSmartContractTemp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUploadNewSmartContractTemp proto.InternalMessageInfo

func (m *MsgUploadNewSmartContractTemp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUploadNewSmartContractTemp) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *MsgUploadNewSmartContractTemp) GetBytecode() string {
	if m != nil {
		return m.Bytecode
	}
	return ""
}

func (m *MsgUploadNewSmartContractTemp) GetConstructorInput() string {
	if m != nil {
		return m.ConstructorInput
	}
	return ""
}

func (m *MsgUploadNewSmartContractTemp) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

type MsgUploadNewSmartContractTempResponse struct {
}

func (m *MsgUploadNewSmartContractTempResponse) Reset()         { *m = MsgUploadNewSmartContractTempResponse{} }
func (m *MsgUploadNewSmartContractTempResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUploadNewSmartContractTempResponse) ProtoMessage()    {}
func (*MsgUploadNewSmartContractTempResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{3}
}
func (m *MsgUploadNewSmartContractTempResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUploadNewSmartContractTempResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUploadNewSmartContractTempResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUploadNewSmartContractTempResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUploadNewSmartContractTempResponse.Merge(m, src)
}
func (m *MsgUploadNewSmartContractTempResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUploadNewSmartContractTempResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUploadNewSmartContractTempResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUploadNewSmartContractTempResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitNewJob)(nil), "palomachain.paloma.evm.MsgSubmitNewJob")
	proto.RegisterType((*MsgSubmitNewJobResponse)(nil), "palomachain.paloma.evm.MsgSubmitNewJobResponse")
	proto.RegisterType((*MsgUploadNewSmartContractTemp)(nil), "palomachain.paloma.evm.MsgUploadNewSmartContractTemp")
	proto.RegisterType((*MsgUploadNewSmartContractTempResponse)(nil), "palomachain.paloma.evm.MsgUploadNewSmartContractTempResponse")
}

func init() { proto.RegisterFile("palomachain/paloma/evm/tx.proto", fileDescriptor_631cfc68eb1fd278) }

var fileDescriptor_631cfc68eb1fd278 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0x6f, 0xef, 0xc8, 0x85, 0x8c, 0x90, 0x88, 0xb6, 0x48, 0x16, 0x0b, 0x0c, 0x3a, 0x09,
	0x05, 0x28, 0x6c, 0x09, 0x84, 0x44, 0x43, 0x01, 0xa1, 0x09, 0xd2, 0x45, 0xe8, 0x12, 0x1a, 0xba,
	0xb5, 0x3d, 0xd8, 0x96, 0xb2, 0xde, 0xd5, 0xee, 0x3a, 0xb1, 0x7b, 0x1e, 0x00, 0xf1, 0x42, 0xb4,
	0x94, 0x29, 0x29, 0xd1, 0xdd, 0x53, 0xd0, 0xa1, 0xec, 0xd9, 0x24, 0xf7, 0xcf, 0x42, 0x74, 0x33,
	0xdf, 0x7c, 0x6b, 0xcf, 0x6f, 0x34, 0x03, 0x0f, 0x15, 0x3f, 0x93, 0x82, 0xc7, 0x19, 0xcf, 0x8b,
	0x70, 0x1e, 0x87, 0x78, 0x2e, 0x42, 0x5b, 0x05, 0x4a, 0x4b, 0x2b, 0xe9, 0xde, 0x0d, 0x43, 0x30,
	0x8f, 0x03, 0x3c, 0x17, 0xa3, 0xdf, 0x04, 0xee, 0x8e, 0x4d, 0x7a, 0x52, 0x46, 0x22, 0xb7, 0xc7,
	0x78, 0xf1, 0x5e, 0x46, 0x94, 0xc1, 0x76, 0xac, 0x91, 0x5b, 0xa9, 0x19, 0x79, 0x44, 0x9e, 0xec,
	0x4c, 0xda, 0x94, 0xbe, 0x82, 0xfd, 0x0c, 0xab, 0x13, 0xc1, 0xb5, 0x3d, 0x94, 0x85, 0xd5, 0x3c,
	0xb6, 0x6f, 0x92, 0x44, 0xa3, 0x31, 0xac, 0xef, 0x9c, 0x9b, 0xca, 0xd4, 0x07, 0xc8, 0xb0, 0xfa,
	0xc0, 0xeb, 0x33, 0xc9, 0x13, 0x36, 0x70, 0xe6, 0x1b, 0x0a, 0xdd, 0x85, 0x01, 0x8f, 0x72, 0x76,
	0xcb, 0x15, 0xae, 0x42, 0xba, 0x07, 0x43, 0x81, 0x36, 0x93, 0x09, 0xdb, 0x72, 0x62, 0x93, 0xd1,
	0xfb, 0xb0, 0xe3, 0x28, 0x4e, 0x6b, 0x85, 0x6c, 0xe8, 0x4a, 0xd7, 0x02, 0x7d, 0x06, 0xbb, 0x2e,
	0x99, 0xe0, 0x67, 0xd4, 0x58, 0xc4, 0x78, 0xf4, 0x8e, 0x6d, 0x3b, 0xd3, 0x8a, 0x3e, 0xba, 0x07,
	0xfb, 0x4b, 0xe8, 0x13, 0x34, 0x4a, 0x16, 0x06, 0x47, 0xdf, 0x09, 0x3c, 0x18, 0x9b, 0xf4, 0xa3,
	0xba, 0x6a, 0xee, 0x18, 0x2f, 0x16, 0x98, 0x4e, 0x51, 0xa8, 0x8e, 0x21, 0x35, 0x28, 0xfd, 0x6b,
	0x14, 0x0f, 0x6e, 0x47, 0xb5, 0xc5, 0x58, 0x26, 0xd8, 0xa0, 0xff, 0xcd, 0x5d, 0xc3, 0xb2, 0x30,
	0x56, 0x97, 0xb1, 0x95, 0xfa, 0xa8, 0x50, 0xa5, 0x6d, 0xa6, 0xb0, 0xa2, 0xaf, 0x85, 0xdb, 0xda,
	0x00, 0x77, 0x00, 0x8f, 0x3b, 0x01, 0x5a, 0xd4, 0xe7, 0x5f, 0xfa, 0x30, 0x18, 0x9b, 0x94, 0x66,
	0x70, 0x67, 0x61, 0x0b, 0x0e, 0x82, 0xf5, 0x2b, 0x13, 0x2c, 0xcd, 0xcc, 0x0b, 0xff, 0xd1, 0xd8,
	0xfe, 0x91, 0x7e, 0x23, 0xe0, 0x75, 0x4c, 0xf6, 0x65, 0xc7, 0xf7, 0x36, 0x3f, 0xf3, 0x5e, 0xff,
	0xd7, 0xb3, 0xb6, 0xa9, 0xb7, 0x87, 0x3f, 0xa6, 0x3e, 0xb9, 0x9c, 0xfa, 0xe4, 0xd7, 0xd4, 0x27,
	0x5f, 0x67, 0x7e, 0xef, 0x72, 0xe6, 0xf7, 0x7e, 0xce, 0xfc, 0xde, 0xa7, 0xa7, 0x69, 0x6e, 0xb3,
	0x32, 0x0a, 0x62, 0x29, 0xc2, 0x35, 0x67, 0x56, 0xcd, 0x0f, 0xad, 0x56, 0x68, 0xa2, 0xa1, 0x3b,
	0xb6, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xb8, 0x85, 0x30, 0x8f, 0x03, 0x00, 0x00,
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
	SubmitNewJob(ctx context.Context, in *MsgSubmitNewJob, opts ...grpc.CallOption) (*MsgSubmitNewJobResponse, error)
	UploadNewSmartContractTemp(ctx context.Context, in *MsgUploadNewSmartContractTemp, opts ...grpc.CallOption) (*MsgUploadNewSmartContractTempResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitNewJob(ctx context.Context, in *MsgSubmitNewJob, opts ...grpc.CallOption) (*MsgSubmitNewJobResponse, error) {
	out := new(MsgSubmitNewJobResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/SubmitNewJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UploadNewSmartContractTemp(ctx context.Context, in *MsgUploadNewSmartContractTemp, opts ...grpc.CallOption) (*MsgUploadNewSmartContractTempResponse, error) {
	out := new(MsgUploadNewSmartContractTempResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/UploadNewSmartContractTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitNewJob(context.Context, *MsgSubmitNewJob) (*MsgSubmitNewJobResponse, error)
	UploadNewSmartContractTemp(context.Context, *MsgUploadNewSmartContractTemp) (*MsgUploadNewSmartContractTempResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitNewJob(ctx context.Context, req *MsgSubmitNewJob) (*MsgSubmitNewJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitNewJob not implemented")
}
func (*UnimplementedMsgServer) UploadNewSmartContractTemp(ctx context.Context, req *MsgUploadNewSmartContractTemp) (*MsgUploadNewSmartContractTempResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadNewSmartContractTemp not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitNewJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitNewJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitNewJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/SubmitNewJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitNewJob(ctx, req.(*MsgSubmitNewJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UploadNewSmartContractTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUploadNewSmartContractTemp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UploadNewSmartContractTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/UploadNewSmartContractTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UploadNewSmartContractTemp(ctx, req.(*MsgUploadNewSmartContractTemp))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "palomachain.paloma.evm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitNewJob",
			Handler:    _Msg_SubmitNewJob_Handler,
		},
		{
			MethodName: "UploadNewSmartContractTemp",
			Handler:    _Msg_UploadNewSmartContractTemp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "palomachain/paloma/evm/tx.proto",
}

func (m *MsgSubmitNewJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitNewJob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitNewJob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainType) > 0 {
		i -= len(m.ChainType)
		copy(dAtA[i:], m.ChainType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HexPayload) > 0 {
		i -= len(m.HexPayload)
		copy(dAtA[i:], m.HexPayload)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HexPayload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HexSmartContractAddress) > 0 {
		i -= len(m.HexSmartContractAddress)
		copy(dAtA[i:], m.HexSmartContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HexSmartContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitNewJobResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitNewJobResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitNewJobResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUploadNewSmartContractTemp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUploadNewSmartContractTemp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUploadNewSmartContractTemp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ConstructorInput) > 0 {
		i -= len(m.ConstructorInput)
		copy(dAtA[i:], m.ConstructorInput)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConstructorInput)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUploadNewSmartContractTempResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUploadNewSmartContractTempResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUploadNewSmartContractTempResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSubmitNewJob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HexSmartContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HexPayload)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitNewJobResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUploadNewSmartContractTemp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConstructorInput)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUploadNewSmartContractTempResponse) Size() (n int) {
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
func (m *MsgSubmitNewJob) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitNewJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitNewJob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexSmartContractAddress", wireType)
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
			m.HexSmartContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexPayload", wireType)
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
			m.HexPayload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
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
			m.Abi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
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
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
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
			m.ChainType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
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
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgSubmitNewJobResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitNewJobResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitNewJobResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgUploadNewSmartContractTemp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUploadNewSmartContractTemp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUploadNewSmartContractTemp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
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
			m.Abi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
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
			m.Bytecode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstructorInput", wireType)
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
			m.ConstructorInput = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
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
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgUploadNewSmartContractTempResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUploadNewSmartContractTempResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUploadNewSmartContractTempResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
