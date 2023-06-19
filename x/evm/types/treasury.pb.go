// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/treasury.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type FundCollectedEvent struct {
	JobID       string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	BlockHeight uint64 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *FundCollectedEvent) Reset()         { *m = FundCollectedEvent{} }
func (m *FundCollectedEvent) String() string { return proto.CompactTextString(m) }
func (*FundCollectedEvent) ProtoMessage()    {}
func (*FundCollectedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3d48b3c1757ff6c, []int{0}
}
func (m *FundCollectedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundCollectedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundCollectedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundCollectedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundCollectedEvent.Merge(m, src)
}
func (m *FundCollectedEvent) XXX_Size() int {
	return m.Size()
}
func (m *FundCollectedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_FundCollectedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_FundCollectedEvent proto.InternalMessageInfo

func (m *FundCollectedEvent) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *FundCollectedEvent) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundCollectedEvent) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *FundCollectedEvent) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type CollectFunds struct {
	FromBlockHeight uint64 `protobuf:"varint,1,opt,name=fromBlockHeight,proto3" json:"fromBlockHeight,omitempty"`
	ToBlockHeight   uint64 `protobuf:"varint,2,opt,name=toBlockHeight,proto3" json:"toBlockHeight,omitempty"`
}

func (m *CollectFunds) Reset()         { *m = CollectFunds{} }
func (m *CollectFunds) String() string { return proto.CompactTextString(m) }
func (*CollectFunds) ProtoMessage()    {}
func (*CollectFunds) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3d48b3c1757ff6c, []int{1}
}
func (m *CollectFunds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectFunds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectFunds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectFunds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectFunds.Merge(m, src)
}
func (m *CollectFunds) XXX_Size() int {
	return m.Size()
}
func (m *CollectFunds) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectFunds.DiscardUnknown(m)
}

var xxx_messageInfo_CollectFunds proto.InternalMessageInfo

func (m *CollectFunds) GetFromBlockHeight() uint64 {
	if m != nil {
		return m.FromBlockHeight
	}
	return 0
}

func (m *CollectFunds) GetToBlockHeight() uint64 {
	if m != nil {
		return m.ToBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*FundCollectedEvent)(nil), "palomachain.paloma.evm.FundCollectedEvent")
	proto.RegisterType((*CollectFunds)(nil), "palomachain.paloma.evm.CollectFunds")
}

func init() {
	proto.RegisterFile("palomachain/paloma/evm/treasury.proto", fileDescriptor_a3d48b3c1757ff6c)
}

var fileDescriptor_a3d48b3c1757ff6c = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x9b, 0x39, 0x07, 0x46, 0x45, 0x08, 0x63, 0x14, 0x0f, 0xa1, 0x0c, 0x85, 0x7a, 0x69,
	0x0f, 0xfe, 0x83, 0x4d, 0x45, 0xaf, 0x3b, 0x7a, 0x10, 0xd2, 0xf6, 0xb3, 0xad, 0x36, 0xf9, 0x4a,
	0x9b, 0x16, 0xe7, 0xaf, 0xf0, 0x67, 0x79, 0xdc, 0xd1, 0xa3, 0xb4, 0x7f, 0x44, 0x9a, 0xec, 0x50,
	0xc5, 0xdb, 0xfb, 0x3e, 0x3c, 0x1f, 0x09, 0x2f, 0xbd, 0x2c, 0x45, 0x81, 0x52, 0xc4, 0x99, 0xc8,
	0x55, 0x68, 0x73, 0x08, 0xad, 0x0c, 0x75, 0x05, 0xa2, 0x6e, 0xaa, 0x6d, 0x50, 0x56, 0xa8, 0x91,
	0x2d, 0x46, 0x5a, 0x60, 0x73, 0x00, 0xad, 0x3c, 0x9f, 0xa7, 0x98, 0xa2, 0x51, 0xc2, 0x21, 0x59,
	0x7b, 0xf9, 0x4e, 0xd9, 0x5d, 0xa3, 0x92, 0x35, 0x16, 0x05, 0xc4, 0x1a, 0x92, 0xdb, 0x16, 0x94,
	0x66, 0x73, 0x7a, 0xf8, 0x82, 0xd1, 0xc3, 0x8d, 0x4b, 0x3c, 0xe2, 0x1f, 0x6d, 0x6c, 0x61, 0x0b,
	0x3a, 0x13, 0x12, 0x1b, 0xa5, 0xdd, 0x89, 0xc1, 0xfb, 0x36, 0xd8, 0x09, 0x28, 0x94, 0xee, 0x81,
	0xb5, 0x4d, 0x61, 0x1e, 0x3d, 0x8e, 0x0a, 0x8c, 0x5f, 0xef, 0x21, 0x4f, 0x33, 0xed, 0x4e, 0x3d,
	0xe2, 0x4f, 0x37, 0x63, 0xb4, 0x7c, 0xa2, 0x27, 0xfb, 0x77, 0x87, 0x2f, 0xd4, 0xcc, 0xa7, 0x67,
	0xcf, 0x15, 0xca, 0xd5, 0xe8, 0x8a, 0x98, 0xab, 0xbf, 0x98, 0x5d, 0xd0, 0x53, 0x8d, 0x63, 0x6f,
	0x62, 0xbc, 0xdf, 0x70, 0xb5, 0xfe, 0xec, 0x38, 0xd9, 0x75, 0x9c, 0x7c, 0x77, 0x9c, 0x7c, 0xf4,
	0xdc, 0xd9, 0xf5, 0xdc, 0xf9, 0xea, 0xb9, 0xf3, 0x78, 0x95, 0xe6, 0x3a, 0x6b, 0xa2, 0x20, 0x46,
	0x19, 0xfe, 0xb3, 0xea, 0x9b, 0xdd, 0x75, 0x5b, 0x42, 0x1d, 0xcd, 0xcc, 0x4e, 0xd7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6b, 0xc9, 0xbb, 0x6b, 0x7e, 0x01, 0x00, 0x00,
}

func (m *FundCollectedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundCollectedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundCollectedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectFunds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectFunds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectFunds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ToBlockHeight != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.ToBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.FromBlockHeight != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.FromBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreasury(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FundCollectedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTreasury(uint64(m.BlockHeight))
	}
	return n
}

func (m *CollectFunds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FromBlockHeight != 0 {
		n += 1 + sovTreasury(uint64(m.FromBlockHeight))
	}
	if m.ToBlockHeight != 0 {
		n += 1 + sovTreasury(uint64(m.ToBlockHeight))
	}
	return n
}

func sovTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreasury(x uint64) (n int) {
	return sovTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundCollectedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: FundCollectedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundCollectedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *CollectFunds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: CollectFunds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectFunds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlockHeight", wireType)
			}
			m.FromBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToBlockHeight", wireType)
			}
			m.ToBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func skipTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
				return 0, ErrInvalidLengthTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreasury = fmt.Errorf("proto: unexpected end of group")
)
