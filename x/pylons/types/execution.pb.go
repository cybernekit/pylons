// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/pylons/execution.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type ItemRecord struct {
	Id      string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Doubles []DoubleKeyValue `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongKeyValue   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringKeyValue `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
}

func (m *ItemRecord) Reset()         { *m = ItemRecord{} }
func (m *ItemRecord) String() string { return proto.CompactTextString(m) }
func (*ItemRecord) ProtoMessage()    {}
func (*ItemRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4ba7e747c28b3a9, []int{0}
}
func (m *ItemRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ItemRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ItemRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRecord.Merge(m, src)
}
func (m *ItemRecord) XXX_Size() int {
	return m.Size()
}
func (m *ItemRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRecord proto.InternalMessageInfo

func (m *ItemRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ItemRecord) GetDoubles() []DoubleKeyValue {
	if m != nil {
		return m.Doubles
	}
	return nil
}

func (m *ItemRecord) GetLongs() []LongKeyValue {
	if m != nil {
		return m.Longs
	}
	return nil
}

func (m *ItemRecord) GetStrings() []StringKeyValue {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Execution struct {
	Creator             string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                  string                                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RecipeId            string                                   `protobuf:"bytes,3,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	CookbookId          string                                   `protobuf:"bytes,4,opt,name=cookbook_id,json=cookbookId,proto3" json:"cookbook_id,omitempty"`
	RecipeVersion       string                                   `protobuf:"bytes,5,opt,name=recipe_version,json=recipeVersion,proto3" json:"recipe_version,omitempty"`
	NodeVersion         uint64                                   `protobuf:"varint,6,opt,name=node_version,json=nodeVersion,proto3" json:"node_version,omitempty"`
	BlockHeight         int64                                    `protobuf:"varint,7,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	ItemInputs          []ItemRecord                             `protobuf:"bytes,8,rep,name=item_inputs,json=itemInputs,proto3" json:"item_inputs"`
	CoinInputs          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=coin_inputs,json=coinInputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coin_inputs"`
	CoinOutputs         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,10,rep,name=coin_outputs,json=coinOutputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coin_outputs"`
	ItemOutputIds       []string                                 `protobuf:"bytes,11,rep,name=item_output_ids,json=itemOutputIds,proto3" json:"item_output_ids,omitempty"`
	ItemModifyOutputIds []string                                 `protobuf:"bytes,12,rep,name=item_modify_output_ids,json=itemModifyOutputIds,proto3" json:"item_modify_output_ids,omitempty"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4ba7e747c28b3a9, []int{1}
}
func (m *Execution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return m.Size()
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Execution) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Execution) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

func (m *Execution) GetCookbookId() string {
	if m != nil {
		return m.CookbookId
	}
	return ""
}

func (m *Execution) GetRecipeVersion() string {
	if m != nil {
		return m.RecipeVersion
	}
	return ""
}

func (m *Execution) GetNodeVersion() uint64 {
	if m != nil {
		return m.NodeVersion
	}
	return 0
}

func (m *Execution) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Execution) GetItemInputs() []ItemRecord {
	if m != nil {
		return m.ItemInputs
	}
	return nil
}

func (m *Execution) GetCoinInputs() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CoinInputs
	}
	return nil
}

func (m *Execution) GetCoinOutputs() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CoinOutputs
	}
	return nil
}

func (m *Execution) GetItemOutputIds() []string {
	if m != nil {
		return m.ItemOutputIds
	}
	return nil
}

func (m *Execution) GetItemModifyOutputIds() []string {
	if m != nil {
		return m.ItemModifyOutputIds
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemRecord)(nil), "pylons.pylons.ItemRecord")
	proto.RegisterType((*Execution)(nil), "pylons.pylons.Execution")
}

func init() { proto.RegisterFile("pylons/pylons/execution.proto", fileDescriptor_a4ba7e747c28b3a9) }

var fileDescriptor_a4ba7e747c28b3a9 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x93, 0xf4, 0x91, 0x71, 0x52, 0xa4, 0x01, 0x21, 0x93, 0xaa, 0x4e, 0xa8, 0x84, 0xe4,
	0x45, 0x6b, 0x53, 0x58, 0x20, 0x16, 0x48, 0x28, 0x3c, 0x44, 0x04, 0x08, 0x14, 0xa4, 0x2e, 0xd8,
	0x44, 0xb1, 0x67, 0x70, 0x46, 0x71, 0xe6, 0x5a, 0x9e, 0x71, 0xd5, 0xfc, 0x05, 0xdf, 0xc1, 0x97,
	0x74, 0x59, 0x89, 0x0d, 0x2b, 0x40, 0xc9, 0x8a, 0xbf, 0x40, 0xf3, 0x70, 0x9b, 0x46, 0x2c, 0x59,
	0x8d, 0x7d, 0xee, 0x39, 0xf7, 0xf8, 0x9e, 0xb9, 0x46, 0x07, 0xf9, 0x22, 0x03, 0x2e, 0x22, 0x7b,
	0xd0, 0x73, 0x9a, 0x94, 0x92, 0x01, 0x0f, 0xf3, 0x02, 0x24, 0xe0, 0x8e, 0xc1, 0x43, 0x73, 0x74,
	0xef, 0xa4, 0x90, 0x82, 0xae, 0x44, 0xea, 0xc9, 0x90, 0xba, 0x7e, 0x02, 0x62, 0x0e, 0x22, 0x8a,
	0x27, 0x82, 0x46, 0x67, 0x27, 0x31, 0x95, 0x93, 0x93, 0x28, 0x01, 0x66, 0x9b, 0x74, 0xbd, 0x9b,
	0x1e, 0x4c, 0xd2, 0xb9, 0xa9, 0x1c, 0x7e, 0x77, 0x10, 0x1a, 0x4a, 0x3a, 0x1f, 0xd1, 0x04, 0x0a,
	0x82, 0xf7, 0x50, 0x9d, 0x11, 0xcf, 0xe9, 0x3b, 0x41, 0x6b, 0x54, 0x67, 0x04, 0x3f, 0x43, 0x3b,
	0x04, 0xca, 0x38, 0xa3, 0xc2, 0xab, 0xf7, 0x1b, 0x81, 0xfb, 0xe8, 0x20, 0xbc, 0xf1, 0x3d, 0xe1,
	0x4b, 0x5d, 0x7d, 0x4b, 0x17, 0xa7, 0x93, 0xac, 0xa4, 0x83, 0xe6, 0xc5, 0xcf, 0x5e, 0x6d, 0x54,
	0x69, 0xf0, 0x13, 0xb4, 0x95, 0x01, 0x4f, 0x85, 0xd7, 0xd0, 0xe2, 0xfd, 0x0d, 0xf1, 0x3b, 0xe0,
	0xe9, 0x86, 0xd4, 0xf0, 0x95, 0xaf, 0x90, 0x05, 0x53, 0xd2, 0xe6, 0x3f, 0x7d, 0x3f, 0xe9, 0xea,
	0xa6, 0xaf, 0xd5, 0x1c, 0xfe, 0x69, 0xa2, 0xd6, 0xab, 0x2a, 0x48, 0xec, 0xa1, 0x9d, 0xa4, 0xa0,
	0x13, 0x09, 0x85, 0x9d, 0xac, 0x7a, 0xb5, 0xe3, 0xd6, 0xaf, 0xc6, 0xdd, 0x47, 0xad, 0x82, 0x26,
	0x2c, 0xa7, 0x63, 0x46, 0xbc, 0x86, 0x86, 0x77, 0x0d, 0x30, 0x24, 0xb8, 0x87, 0xdc, 0x04, 0x60,
	0x16, 0x03, 0xcc, 0x54, 0xb9, 0xa9, 0xcb, 0xa8, 0x82, 0x86, 0x04, 0x3f, 0x40, 0x7b, 0x56, 0x7d,
	0x46, 0x0b, 0xc1, 0x80, 0x7b, 0x5b, 0x9a, 0xd3, 0x31, 0xe8, 0xa9, 0x01, 0xf1, 0x7d, 0xd4, 0xe6,
	0x40, 0xae, 0x49, 0xdb, 0x7d, 0x27, 0x68, 0x8e, 0x5c, 0x85, 0xad, 0x51, 0xe2, 0x0c, 0x92, 0xd9,
	0x78, 0x4a, 0x59, 0x3a, 0x95, 0xde, 0x4e, 0xdf, 0x09, 0x1a, 0x23, 0x57, 0x63, 0x6f, 0x34, 0x84,
	0x9f, 0x23, 0x57, 0x5d, 0xe3, 0x98, 0xf1, 0xbc, 0x94, 0xc2, 0xdb, 0xd5, 0x29, 0xdd, 0xdb, 0x48,
	0xe9, 0xfa, 0x66, 0x6d, 0x42, 0x48, 0x69, 0x86, 0x5a, 0x82, 0x33, 0x35, 0x0f, 0xe3, 0x55, 0x87,
	0x96, 0xed, 0x60, 0x56, 0x29, 0x54, 0xab, 0x14, 0xda, 0x55, 0x0a, 0x5f, 0x00, 0xe3, 0x83, 0x87,
	0xaa, 0xc3, 0xb7, 0x5f, 0xbd, 0x20, 0x65, 0x72, 0x5a, 0xc6, 0x61, 0x02, 0xf3, 0xc8, 0xee, 0x9d,
	0x39, 0x8e, 0x05, 0x99, 0x45, 0x72, 0x91, 0x53, 0xa1, 0x05, 0x42, 0x85, 0xc3, 0xb8, 0x75, 0xe3,
	0xa8, 0xad, 0xdd, 0xa0, 0x94, 0xda, 0x0e, 0xfd, 0x7f, 0x3b, 0x3d, 0xce, 0x07, 0xd3, 0x1f, 0x1f,
	0xa1, 0x5b, 0x3a, 0x1f, 0xe3, 0x37, 0x66, 0x44, 0x78, 0x6e, 0xbf, 0x11, 0xb4, 0x6c, 0x10, 0x1d,
	0x55, 0x34, 0xdc, 0x21, 0x11, 0xf8, 0x29, 0xba, 0xab, 0xd9, 0x73, 0x20, 0xec, 0xcb, 0x62, 0x5d,
	0xd4, 0x5e, 0x13, 0xdd, 0x56, 0x9c, 0xf7, 0x9a, 0x72, 0x25, 0x1d, 0xbc, 0xbe, 0x58, 0xfa, 0xce,
	0xe5, 0xd2, 0x77, 0x7e, 0x2f, 0x7d, 0xe7, 0xeb, 0xca, 0xaf, 0x5d, 0xae, 0xfc, 0xda, 0x8f, 0x95,
	0x5f, 0xfb, 0x7c, 0xb4, 0xf6, 0xe5, 0x1f, 0xf5, 0x85, 0x1c, 0x4b, 0x9a, 0x4c, 0xab, 0xbf, 0xf0,
	0xbc, 0x7a, 0xd0, 0x33, 0xc4, 0xdb, 0xfa, 0x87, 0x7c, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0x6e, 0x5b, 0xab, 0x10, 0x04, 0x00, 0x00,
}

func (m *ItemRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ItemRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Strings) > 0 {
		for iNdEx := len(m.Strings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Strings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Longs) > 0 {
		for iNdEx := len(m.Longs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Longs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Doubles) > 0 {
		for iNdEx := len(m.Doubles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Doubles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Execution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Execution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Execution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ItemModifyOutputIds) > 0 {
		for iNdEx := len(m.ItemModifyOutputIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ItemModifyOutputIds[iNdEx])
			copy(dAtA[i:], m.ItemModifyOutputIds[iNdEx])
			i = encodeVarintExecution(dAtA, i, uint64(len(m.ItemModifyOutputIds[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.ItemOutputIds) > 0 {
		for iNdEx := len(m.ItemOutputIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ItemOutputIds[iNdEx])
			copy(dAtA[i:], m.ItemOutputIds[iNdEx])
			i = encodeVarintExecution(dAtA, i, uint64(len(m.ItemOutputIds[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.CoinOutputs) > 0 {
		for iNdEx := len(m.CoinOutputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinOutputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.CoinInputs) > 0 {
		for iNdEx := len(m.CoinInputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinInputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ItemInputs) > 0 {
		for iNdEx := len(m.ItemInputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemInputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintExecution(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.NodeVersion != 0 {
		i = encodeVarintExecution(dAtA, i, uint64(m.NodeVersion))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RecipeVersion) > 0 {
		i -= len(m.RecipeVersion)
		copy(dAtA[i:], m.RecipeVersion)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.RecipeVersion)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CookbookId) > 0 {
		i -= len(m.CookbookId)
		copy(dAtA[i:], m.CookbookId)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.CookbookId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecipeId) > 0 {
		i -= len(m.RecipeId)
		copy(dAtA[i:], m.RecipeId)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.RecipeId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExecution(dAtA []byte, offset int, v uint64) int {
	offset -= sovExecution(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ItemRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	if len(m.Doubles) > 0 {
		for _, e := range m.Doubles {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.Longs) > 0 {
		for _, e := range m.Longs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.Strings) > 0 {
		for _, e := range m.Strings {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func (m *Execution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.RecipeId)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.CookbookId)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.RecipeVersion)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	if m.NodeVersion != 0 {
		n += 1 + sovExecution(uint64(m.NodeVersion))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovExecution(uint64(m.BlockHeight))
	}
	if len(m.ItemInputs) > 0 {
		for _, e := range m.ItemInputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.CoinInputs) > 0 {
		for _, e := range m.CoinInputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.CoinOutputs) > 0 {
		for _, e := range m.CoinOutputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.ItemOutputIds) > 0 {
		for _, s := range m.ItemOutputIds {
			l = len(s)
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.ItemModifyOutputIds) > 0 {
		for _, s := range m.ItemModifyOutputIds {
			l = len(s)
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func sovExecution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExecution(x uint64) (n int) {
	return sovExecution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ItemRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
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
			return fmt.Errorf("proto: ItemRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ItemRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Doubles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Doubles = append(m.Doubles, DoubleKeyValue{})
			if err := m.Doubles[len(m.Doubles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Longs = append(m.Longs, LongKeyValue{})
			if err := m.Longs[len(m.Longs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Strings = append(m.Strings, StringKeyValue{})
			if err := m.Strings[len(m.Strings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecution
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
func (m *Execution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
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
			return fmt.Errorf("proto: Execution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Execution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeVersion", wireType)
			}
			m.NodeVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemInputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemInputs = append(m.ItemInputs, ItemRecord{})
			if err := m.ItemInputs[len(m.ItemInputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinInputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinInputs = append(m.CoinInputs, types.Coin{})
			if err := m.CoinInputs[len(m.CoinInputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinOutputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinOutputs = append(m.CoinOutputs, types.Coin{})
			if err := m.CoinOutputs[len(m.CoinOutputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemOutputIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemOutputIds = append(m.ItemOutputIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemModifyOutputIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
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
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemModifyOutputIds = append(m.ItemModifyOutputIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecution
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
func skipExecution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecution
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
					return 0, ErrIntOverflowExecution
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
					return 0, ErrIntOverflowExecution
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
				return 0, ErrInvalidLengthExecution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExecution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExecution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExecution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExecution = fmt.Errorf("proto: unexpected end of group")
)
