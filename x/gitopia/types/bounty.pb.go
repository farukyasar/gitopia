// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/gitopia/bounty.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type BountyState int32

const (
	BountyStateSRCDEBITTED  BountyState = 0
	BountyStateDESTCREDITED BountyState = 1
	BountyStateREVERTEDBACK BountyState = 2
)

var BountyState_name = map[int32]string{
	0: "BOUNTY_STATE_SRCDEBITTED",
	1: "BOUNTY_STATE_DESTCREDITED",
	2: "BOUNTY_STATE_REVERTEDBACK",
}

var BountyState_value = map[string]int32{
	"BOUNTY_STATE_SRCDEBITTED":  0,
	"BOUNTY_STATE_DESTCREDITED": 1,
	"BOUNTY_STATE_REVERTEDBACK": 2,
}

func (x BountyState) String() string {
	return proto.EnumName(BountyState_name, int32(x))
}

func (BountyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b40b67d80582e8a3, []int{0}
}

type BountyParent int32

const (
	BountyParentIssue BountyParent = 0
)

var BountyParent_name = map[int32]string{
	0: "BOUNTY_PARENT_ISSUE",
}

var BountyParent_value = map[string]int32{
	"BOUNTY_PARENT_ISSUE": 0,
}

func (x BountyParent) String() string {
	return proto.EnumName(BountyParent_name, int32(x))
}

func (BountyParent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b40b67d80582e8a3, []int{1}
}

type Bounty struct {
	Id           uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	State        BountyState                              `protobuf:"varint,3,opt,name=state,proto3,enum=gitopia.gitopia.gitopia.BountyState" json:"state,omitempty"`
	RepositoryId uint64                                   `protobuf:"varint,4,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
	ParentIid    uint64                                   `protobuf:"varint,5,opt,name=parentIid,proto3" json:"parentIid,omitempty"`
	Parent       BountyParent                             `protobuf:"varint,6,opt,name=parent,proto3,enum=gitopia.gitopia.gitopia.BountyParent" json:"parent,omitempty"`
	ExpireAt     int64                                    `protobuf:"varint,7,opt,name=expireAt,proto3" json:"expireAt,omitempty"`
	RewardedTo   string                                   `protobuf:"bytes,8,opt,name=rewardedTo,proto3" json:"rewardedTo,omitempty"`
	CreatedAt    int64                                    `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64                                    `protobuf:"varint,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Creator      string                                   `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Bounty) Reset()         { *m = Bounty{} }
func (m *Bounty) String() string { return proto.CompactTextString(m) }
func (*Bounty) ProtoMessage()    {}
func (*Bounty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40b67d80582e8a3, []int{0}
}
func (m *Bounty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bounty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bounty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bounty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bounty.Merge(m, src)
}
func (m *Bounty) XXX_Size() int {
	return m.Size()
}
func (m *Bounty) XXX_DiscardUnknown() {
	xxx_messageInfo_Bounty.DiscardUnknown(m)
}

var xxx_messageInfo_Bounty proto.InternalMessageInfo

func (m *Bounty) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bounty) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Bounty) GetState() BountyState {
	if m != nil {
		return m.State
	}
	return BountyStateSRCDEBITTED
}

func (m *Bounty) GetRepositoryId() uint64 {
	if m != nil {
		return m.RepositoryId
	}
	return 0
}

func (m *Bounty) GetParentIid() uint64 {
	if m != nil {
		return m.ParentIid
	}
	return 0
}

func (m *Bounty) GetParent() BountyParent {
	if m != nil {
		return m.Parent
	}
	return BountyParentIssue
}

func (m *Bounty) GetExpireAt() int64 {
	if m != nil {
		return m.ExpireAt
	}
	return 0
}

func (m *Bounty) GetRewardedTo() string {
	if m != nil {
		return m.RewardedTo
	}
	return ""
}

func (m *Bounty) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Bounty) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Bounty) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.BountyState", BountyState_name, BountyState_value)
	proto.RegisterEnum("gitopia.gitopia.gitopia.BountyParent", BountyParent_name, BountyParent_value)
	proto.RegisterType((*Bounty)(nil), "gitopia.gitopia.gitopia.Bounty")
}

func init() { proto.RegisterFile("gitopia/gitopia/bounty.proto", fileDescriptor_b40b67d80582e8a3) }

var fileDescriptor_b40b67d80582e8a3 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xb7, 0x93, 0x34, 0x6d, 0x94, 0x52, 0x32, 0xed, 0x4f, 0x55, 0xaf, 0xb8, 0xa6, 0x6c, 0x60,
	0x02, 0xb3, 0xd7, 0xee, 0x32, 0x0a, 0x3b, 0xc4, 0xb1, 0x0f, 0x66, 0xd0, 0x15, 0xd9, 0x1d, 0x6c,
	0x97, 0xe0, 0xc4, 0x22, 0x33, 0x23, 0x91, 0xb1, 0x94, 0xad, 0x79, 0x83, 0x91, 0xd3, 0x5e, 0x20,
	0xa7, 0xdd, 0xf6, 0x1e, 0x83, 0x1e, 0x76, 0xe8, 0x71, 0xa7, 0x6d, 0x24, 0x2f, 0x32, 0x2c, 0xbb,
	0xa9, 0xda, 0x32, 0x7a, 0xfa, 0xfc, 0xfd, 0xfe, 0x7d, 0x9f, 0x85, 0x04, 0x76, 0x87, 0x09, 0xa7,
	0x69, 0x12, 0xd9, 0x97, 0xb5, 0x4f, 0x27, 0x63, 0x3e, 0xb5, 0xd2, 0x8c, 0x72, 0x0a, 0xb7, 0x4b,
	0xd4, 0xba, 0x51, 0xb5, 0x07, 0x43, 0x3a, 0xa4, 0x42, 0x63, 0xe7, 0x5f, 0x85, 0x5c, 0xd3, 0x07,
	0x94, 0x8d, 0x28, 0xb3, 0xfb, 0x11, 0x23, 0xf6, 0xa7, 0x83, 0x3e, 0xe1, 0xd1, 0x81, 0x3d, 0xa0,
	0xc9, 0xb8, 0xe0, 0xf7, 0x7f, 0x56, 0x41, 0xdd, 0x11, 0xf9, 0x70, 0x0b, 0x54, 0x92, 0x18, 0xa9,
	0x86, 0x6a, 0xd6, 0x70, 0x25, 0x89, 0xe1, 0x00, 0xd4, 0xa3, 0x51, 0x4e, 0xa1, 0x8a, 0x51, 0x35,
	0x9b, 0x87, 0x3b, 0x56, 0x91, 0x65, 0xe5, 0x59, 0x56, 0x99, 0x65, 0x75, 0x69, 0x32, 0x76, 0x9e,
	0x9f, 0xff, 0xde, 0x53, 0xbe, 0xff, 0xd9, 0x33, 0x87, 0x09, 0xff, 0x30, 0xe9, 0x5b, 0x03, 0x3a,
	0xb2, 0xcb, 0xc1, 0x45, 0x79, 0xc6, 0xe2, 0x8f, 0x36, 0x9f, 0xa6, 0x84, 0x09, 0x03, 0xc3, 0x65,
	0x34, 0x3c, 0x02, 0x6b, 0x8c, 0x47, 0x9c, 0xa0, 0xaa, 0xa1, 0x9a, 0x5b, 0x87, 0x4f, 0xac, 0xff,
	0xfc, 0x9e, 0x55, 0x2c, 0x19, 0xe4, 0x5a, 0x5c, 0x58, 0xe0, 0x3e, 0xd8, 0xcc, 0x48, 0x4a, 0x59,
	0xc2, 0x69, 0x36, 0xf5, 0x63, 0x54, 0x13, 0xab, 0x5f, 0xc3, 0xe0, 0x2e, 0x68, 0xa4, 0x51, 0x46,
	0xc6, 0xdc, 0x4f, 0x62, 0xb4, 0x26, 0x04, 0x57, 0x00, 0x7c, 0x05, 0xea, 0x45, 0x83, 0xea, 0x62,
	0xfc, 0xd3, 0x3b, 0xc6, 0x9f, 0x08, 0x31, 0x2e, 0x4d, 0x50, 0x03, 0x1b, 0xe4, 0x2c, 0x4d, 0x32,
	0xd2, 0xe1, 0x68, 0xdd, 0x50, 0xcd, 0x2a, 0x5e, 0xf5, 0x50, 0x07, 0x20, 0x23, 0x9f, 0xa3, 0x2c,
	0x26, 0x71, 0x48, 0xd1, 0x86, 0xa1, 0x9a, 0x0d, 0x2c, 0x21, 0xf9, 0x62, 0x83, 0x8c, 0x44, 0x9c,
	0xc4, 0x1d, 0x8e, 0x1a, 0xc2, 0x7c, 0x05, 0xe4, 0xec, 0x24, 0x8d, 0x4b, 0x16, 0x14, 0xec, 0x0a,
	0x80, 0x08, 0xac, 0x0b, 0x29, 0xcd, 0x50, 0x53, 0x04, 0x5f, 0xb6, 0xed, 0x1f, 0x2a, 0x68, 0x4a,
	0x27, 0x05, 0x5f, 0x02, 0xe4, 0xbc, 0x39, 0x3d, 0x0e, 0xdf, 0xf5, 0x82, 0xb0, 0x13, 0x7a, 0xbd,
	0x00, 0x77, 0x5d, 0xcf, 0xf1, 0xc3, 0xd0, 0x73, 0x5b, 0x8a, 0xa6, 0xcd, 0xe6, 0xc6, 0x23, 0x49,
	0x2e, 0xb1, 0xf0, 0x08, 0xec, 0x5c, 0x73, 0xba, 0x5e, 0x10, 0x76, 0xb1, 0xe7, 0xfa, 0xb9, 0x55,
	0xd5, 0x1e, 0xcf, 0xe6, 0xc6, 0xb6, 0x64, 0x95, 0xe9, 0x5b, 0x5e, 0xec, 0xbd, 0xf5, 0x70, 0xe8,
	0xb9, 0x4e, 0xa7, 0xfb, 0xba, 0x55, 0xb9, 0xe5, 0x95, 0x69, 0xad, 0xf6, 0xe5, 0x9b, 0xae, 0xb4,
	0x5d, 0xb0, 0x29, 0x9f, 0x38, 0xb4, 0xc0, 0xfd, 0x32, 0xf1, 0xa4, 0x83, 0xbd, 0xe3, 0xb0, 0xe7,
	0x07, 0xc1, 0xa9, 0xd7, 0x52, 0xb4, 0x87, 0xb3, 0xb9, 0x71, 0x4f, 0x96, 0xfa, 0x8c, 0x4d, 0x48,
	0x91, 0xe2, 0xb8, 0xe7, 0x0b, 0x5d, 0xbd, 0x58, 0xe8, 0xea, 0xdf, 0x85, 0xae, 0x7e, 0x5d, 0xea,
	0xca, 0xc5, 0x52, 0x57, 0x7e, 0x2d, 0x75, 0xe5, 0x7d, 0x5b, 0xba, 0xa8, 0x37, 0x9f, 0xdb, 0xd9,
	0xea, 0x4b, 0x5c, 0xd8, 0x7e, 0x5d, 0xbc, 0x94, 0x17, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x81,
	0x44, 0x2a, 0xca, 0x98, 0x03, 0x00, 0x00,
}

func (m *Bounty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bounty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bounty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBounty(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x50
	}
	if m.CreatedAt != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RewardedTo) > 0 {
		i -= len(m.RewardedTo)
		copy(dAtA[i:], m.RewardedTo)
		i = encodeVarintBounty(dAtA, i, uint64(len(m.RewardedTo)))
		i--
		dAtA[i] = 0x42
	}
	if m.ExpireAt != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.ExpireAt))
		i--
		dAtA[i] = 0x38
	}
	if m.Parent != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.Parent))
		i--
		dAtA[i] = 0x30
	}
	if m.ParentIid != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.ParentIid))
		i--
		dAtA[i] = 0x28
	}
	if m.RepositoryId != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.RepositoryId))
		i--
		dAtA[i] = 0x20
	}
	if m.State != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBounty(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBounty(dAtA []byte, offset int, v uint64) int {
	offset -= sovBounty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bounty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBounty(uint64(m.Id))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovBounty(uint64(l))
		}
	}
	if m.State != 0 {
		n += 1 + sovBounty(uint64(m.State))
	}
	if m.RepositoryId != 0 {
		n += 1 + sovBounty(uint64(m.RepositoryId))
	}
	if m.ParentIid != 0 {
		n += 1 + sovBounty(uint64(m.ParentIid))
	}
	if m.Parent != 0 {
		n += 1 + sovBounty(uint64(m.Parent))
	}
	if m.ExpireAt != 0 {
		n += 1 + sovBounty(uint64(m.ExpireAt))
	}
	l = len(m.RewardedTo)
	if l > 0 {
		n += 1 + l + sovBounty(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovBounty(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovBounty(uint64(m.UpdatedAt))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBounty(uint64(l))
	}
	return n
}

func sovBounty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBounty(x uint64) (n int) {
	return sovBounty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bounty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBounty
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
			return fmt.Errorf("proto: Bounty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bounty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= BountyState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryId", wireType)
			}
			m.RepositoryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepositoryId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentIid", wireType)
			}
			m.ParentIid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParentIid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			m.Parent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Parent |= BountyParent(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			m.ExpireAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardedTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardedTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBounty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBounty
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
func skipBounty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBounty
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
					return 0, ErrIntOverflowBounty
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
					return 0, ErrIntOverflowBounty
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
				return 0, ErrInvalidLengthBounty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBounty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBounty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBounty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBounty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBounty = fmt.Errorf("proto: unexpected end of group")
)
