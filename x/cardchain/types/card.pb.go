// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/card.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Card struct {
	Owner              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner"`
	Content            []byte                                        `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Image              []byte                                        `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	FullArt            string                                        `protobuf:"bytes,4,opt,name=fullArt,proto3" json:"fullArt,omitempty"`
	Notes              string                                        `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Status             string                                        `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	VotePool           github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,7,opt,name=votePool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"votePool"`
	FairEnoughVotes    uint64                                        `protobuf:"varint,8,opt,name=fairEnoughVotes,proto3" json:"fairEnoughVotes,omitempty"`
	OverpoweredVotes   uint64                                        `protobuf:"varint,9,opt,name=overpoweredVotes,proto3" json:"overpoweredVotes,omitempty"`
	UnderpoweredVotes  uint64                                        `protobuf:"varint,10,opt,name=underpoweredVotes,proto3" json:"underpoweredVotes,omitempty"`
	InappropriateVotes uint64                                        `protobuf:"varint,11,opt,name=inappropriateVotes,proto3" json:"inappropriateVotes,omitempty"`
	Nerflevel          int64                                         `protobuf:"varint,12,opt,name=nerflevel,proto3" json:"nerflevel,omitempty"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc719406a70201e7, []int{0}
}
func (m *Card) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Card.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return m.Size()
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Card) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Card) GetFullArt() string {
	if m != nil {
		return m.FullArt
	}
	return ""
}

func (m *Card) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Card) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Card) GetFairEnoughVotes() uint64 {
	if m != nil {
		return m.FairEnoughVotes
	}
	return 0
}

func (m *Card) GetOverpoweredVotes() uint64 {
	if m != nil {
		return m.OverpoweredVotes
	}
	return 0
}

func (m *Card) GetUnderpoweredVotes() uint64 {
	if m != nil {
		return m.UnderpoweredVotes
	}
	return 0
}

func (m *Card) GetInappropriateVotes() uint64 {
	if m != nil {
		return m.InappropriateVotes
	}
	return 0
}

func (m *Card) GetNerflevel() int64 {
	if m != nil {
		return m.Nerflevel
	}
	return 0
}

func init() {
	proto.RegisterType((*Card)(nil), "DecentralCardGame.cardchain.cardchain.Card")
}

func init() { proto.RegisterFile("cardchain/card.proto", fileDescriptor_cc719406a70201e7) }

var fileDescriptor_cc719406a70201e7 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0xf6, 0xcf, 0xbd, 0x1d, 0x2f, 0xa8, 0x43, 0x91, 0x41, 0x24, 0x37, 0x08, 0x62,
	0x10, 0x6f, 0xb2, 0x10, 0xc1, 0x6d, 0x6f, 0x15, 0x17, 0x77, 0x23, 0x59, 0xb8, 0x70, 0x37, 0x9d,
	0x9c, 0xa6, 0xc1, 0x64, 0x4e, 0x98, 0x99, 0xb4, 0xfa, 0x16, 0x3e, 0x56, 0x97, 0x5d, 0x8a, 0x8b,
	0x22, 0xed, 0xce, 0xa7, 0x90, 0x4c, 0xfa, 0x47, 0x5a, 0x17, 0xae, 0x72, 0xbe, 0x6f, 0x7e, 0x67,
	0x32, 0xe7, 0xf0, 0xd1, 0xa1, 0x14, 0x3a, 0x95, 0x33, 0x91, 0xab, 0xb8, 0xa9, 0xa2, 0x4a, 0xa3,
	0x45, 0xf6, 0xfc, 0x1d, 0x48, 0x50, 0x56, 0x8b, 0x62, 0x2c, 0x74, 0xfa, 0x41, 0x94, 0x10, 0x1d,
	0xb8, 0x63, 0xf5, 0x64, 0x98, 0x61, 0x86, 0xae, 0x23, 0x6e, 0xaa, 0xb6, 0xf9, 0xd9, 0xef, 0x0e,
	0xed, 0x36, 0x6d, 0xec, 0x8e, 0xf6, 0x70, 0xa1, 0x40, 0x73, 0x12, 0x90, 0x70, 0x70, 0xfb, 0x66,
	0xb9, 0xbe, 0xf6, 0x7e, 0xae, 0xaf, 0x6f, 0xb2, 0xdc, 0xce, 0xea, 0x49, 0x24, 0xb1, 0x8c, 0x25,
	0x9a, 0x12, 0xcd, 0xee, 0x73, 0x63, 0xd2, 0x2f, 0xb1, 0xfd, 0x56, 0x81, 0x89, 0x46, 0x52, 0x8e,
	0xd2, 0x54, 0x83, 0x31, 0x49, 0x7b, 0x07, 0xe3, 0xf4, 0x42, 0xa2, 0xb2, 0xa0, 0x2c, 0xbf, 0x17,
	0x90, 0xf0, 0x2a, 0xd9, 0x4b, 0x36, 0xa4, 0xbd, 0xbc, 0x14, 0x19, 0xf0, 0x8e, 0xf3, 0x5b, 0xd1,
	0xf0, 0xd3, 0xba, 0x28, 0x46, 0xda, 0xf2, 0x6e, 0xf3, 0xfb, 0x64, 0x2f, 0x1b, 0x5e, 0xa1, 0x05,
	0xc3, 0x7b, 0xce, 0x6f, 0x05, 0x7b, 0x4c, 0xfb, 0xc6, 0x0a, 0x5b, 0x1b, 0xde, 0x77, 0xf6, 0x4e,
	0xb1, 0x3b, 0x7a, 0x39, 0x47, 0x0b, 0x1f, 0x11, 0x0b, 0x7e, 0xe1, 0xe6, 0x88, 0x77, 0x73, 0xbc,
	0xf8, 0x8f, 0x39, 0xc6, 0x98, 0xab, 0xe4, 0x70, 0x01, 0x0b, 0xe9, 0x83, 0xa9, 0xc8, 0xf5, 0x7b,
	0x85, 0x75, 0x36, 0xfb, 0xe4, 0x1e, 0x71, 0x19, 0x90, 0xb0, 0x9b, 0x9c, 0xda, 0xec, 0x25, 0x7d,
	0x88, 0x73, 0xd0, 0x15, 0x2e, 0x40, 0x43, 0xda, 0xa2, 0x03, 0x87, 0x9e, 0xf9, 0xec, 0x15, 0x7d,
	0x54, 0xab, 0xf4, 0x04, 0xa6, 0x0e, 0x3e, 0x3f, 0x60, 0x11, 0x65, 0xb9, 0x12, 0x55, 0xa5, 0xb1,
	0xd2, 0xb9, 0xb0, 0xd0, 0xe2, 0xf7, 0x1d, 0xfe, 0x8f, 0x13, 0xf6, 0x94, 0x0e, 0x14, 0xe8, 0x69,
	0x01, 0x73, 0x28, 0xf8, 0x55, 0x40, 0xc2, 0x4e, 0x72, 0x34, 0x6e, 0x93, 0xe5, 0xc6, 0x27, 0xab,
	0x8d, 0x4f, 0x7e, 0x6d, 0x7c, 0xf2, 0x7d, 0xeb, 0x7b, 0xab, 0xad, 0xef, 0xfd, 0xd8, 0xfa, 0xde,
	0xe7, 0xb7, 0x7f, 0xad, 0xe7, 0x2c, 0x4e, 0xf1, 0xf8, 0x10, 0xbb, 0xaf, 0xf1, 0x31, 0x82, 0x6e,
	0x69, 0x93, 0xbe, 0xcb, 0xd1, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x05, 0x0d, 0xc2,
	0x9c, 0x02, 0x00, 0x00,
}

func (m *Card) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Card) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Card) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nerflevel != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.Nerflevel))
		i--
		dAtA[i] = 0x60
	}
	if m.InappropriateVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.InappropriateVotes))
		i--
		dAtA[i] = 0x58
	}
	if m.UnderpoweredVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.UnderpoweredVotes))
		i--
		dAtA[i] = 0x50
	}
	if m.OverpoweredVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.OverpoweredVotes))
		i--
		dAtA[i] = 0x48
	}
	if m.FairEnoughVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.FairEnoughVotes))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.VotePool.Size()
		i -= size
		if _, err := m.VotePool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCard(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Notes) > 0 {
		i -= len(m.Notes)
		copy(dAtA[i:], m.Notes)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Notes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FullArt) > 0 {
		i -= len(m.FullArt)
		copy(dAtA[i:], m.FullArt)
		i = encodeVarintCard(dAtA, i, uint64(len(m.FullArt)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Owner.Size()
		i -= size
		if _, err := m.Owner.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCard(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCard(dAtA []byte, offset int, v uint64) int {
	offset -= sovCard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Card) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovCard(uint64(l))
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.FullArt)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.Notes)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = m.VotePool.Size()
	n += 1 + l + sovCard(uint64(l))
	if m.FairEnoughVotes != 0 {
		n += 1 + sovCard(uint64(m.FairEnoughVotes))
	}
	if m.OverpoweredVotes != 0 {
		n += 1 + sovCard(uint64(m.OverpoweredVotes))
	}
	if m.UnderpoweredVotes != 0 {
		n += 1 + sovCard(uint64(m.UnderpoweredVotes))
	}
	if m.InappropriateVotes != 0 {
		n += 1 + sovCard(uint64(m.InappropriateVotes))
	}
	if m.Nerflevel != 0 {
		n += 1 + sovCard(uint64(m.Nerflevel))
	}
	return n
}

func sovCard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCard(x uint64) (n int) {
	return sovCard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Card) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCard
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
			return fmt.Errorf("proto: Card: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Card: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = append(m.Image[:0], dAtA[iNdEx:postIndex]...)
			if m.Image == nil {
				m.Image = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullArt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullArt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FairEnoughVotes", wireType)
			}
			m.FairEnoughVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FairEnoughVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverpoweredVotes", wireType)
			}
			m.OverpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OverpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderpoweredVotes", wireType)
			}
			m.UnderpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnderpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InappropriateVotes", wireType)
			}
			m.InappropriateVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InappropriateVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nerflevel", wireType)
			}
			m.Nerflevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nerflevel |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCard
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
func skipCard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCard
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
					return 0, ErrIntOverflowCard
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
					return 0, ErrIntOverflowCard
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
				return 0, ErrInvalidLengthCard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCard = fmt.Errorf("proto: unexpected end of group")
)
