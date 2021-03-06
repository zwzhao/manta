// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_modifiers.proto

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DOTA_MODIFIER_ENTRY_TYPE int32

const (
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE  DOTA_MODIFIER_ENTRY_TYPE = 1
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_REMOVED DOTA_MODIFIER_ENTRY_TYPE = 2
)

var DOTA_MODIFIER_ENTRY_TYPE_name = map[int32]string{
	1: "DOTA_MODIFIER_ENTRY_TYPE_ACTIVE",
	2: "DOTA_MODIFIER_ENTRY_TYPE_REMOVED",
}
var DOTA_MODIFIER_ENTRY_TYPE_value = map[string]int32{
	"DOTA_MODIFIER_ENTRY_TYPE_ACTIVE":  1,
	"DOTA_MODIFIER_ENTRY_TYPE_REMOVED": 2,
}

func (x DOTA_MODIFIER_ENTRY_TYPE) Enum() *DOTA_MODIFIER_ENTRY_TYPE {
	p := new(DOTA_MODIFIER_ENTRY_TYPE)
	*p = x
	return p
}
func (x DOTA_MODIFIER_ENTRY_TYPE) String() string {
	return proto.EnumName(DOTA_MODIFIER_ENTRY_TYPE_name, int32(x))
}
func (x *DOTA_MODIFIER_ENTRY_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DOTA_MODIFIER_ENTRY_TYPE_value, data, "DOTA_MODIFIER_ENTRY_TYPE")
	if err != nil {
		return err
	}
	*x = DOTA_MODIFIER_ENTRY_TYPE(value)
	return nil
}
func (DOTA_MODIFIER_ENTRY_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

type CDOTAModifierBuffTableEntry struct {
	EntryType           *DOTA_MODIFIER_ENTRY_TYPE `protobuf:"varint,1,req,name=entry_type,json=entryType,enum=dota.DOTA_MODIFIER_ENTRY_TYPE,def=1" json:"entry_type,omitempty"`
	Parent              *int32                    `protobuf:"varint,2,req,name=parent" json:"parent,omitempty"`
	Index               *int32                    `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	SerialNum           *int32                    `protobuf:"varint,4,req,name=serial_num,json=serialNum" json:"serial_num,omitempty"`
	ModifierClass       *int32                    `protobuf:"varint,5,opt,name=modifier_class,json=modifierClass" json:"modifier_class,omitempty"`
	AbilityLevel        *int32                    `protobuf:"varint,6,opt,name=ability_level,json=abilityLevel" json:"ability_level,omitempty"`
	StackCount          *int32                    `protobuf:"varint,7,opt,name=stack_count,json=stackCount" json:"stack_count,omitempty"`
	CreationTime        *float32                  `protobuf:"fixed32,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	Duration            *float32                  `protobuf:"fixed32,9,opt,name=duration,def=-1" json:"duration,omitempty"`
	Caster              *int32                    `protobuf:"varint,10,opt,name=caster" json:"caster,omitempty"`
	Ability             *int32                    `protobuf:"varint,11,opt,name=ability" json:"ability,omitempty"`
	Armor               *int32                    `protobuf:"varint,12,opt,name=armor" json:"armor,omitempty"`
	FadeTime            *float32                  `protobuf:"fixed32,13,opt,name=fade_time,json=fadeTime" json:"fade_time,omitempty"`
	Subtle              *bool                     `protobuf:"varint,14,opt,name=subtle" json:"subtle,omitempty"`
	ChannelTime         *float32                  `protobuf:"fixed32,15,opt,name=channel_time,json=channelTime" json:"channel_time,omitempty"`
	VStart              *CMsgVector               `protobuf:"bytes,16,opt,name=v_start,json=vStart" json:"v_start,omitempty"`
	VEnd                *CMsgVector               `protobuf:"bytes,17,opt,name=v_end,json=vEnd" json:"v_end,omitempty"`
	PortalLoopAppear    *string                   `protobuf:"bytes,18,opt,name=portal_loop_appear,json=portalLoopAppear" json:"portal_loop_appear,omitempty"`
	PortalLoopDisappear *string                   `protobuf:"bytes,19,opt,name=portal_loop_disappear,json=portalLoopDisappear" json:"portal_loop_disappear,omitempty"`
	HeroLoopAppear      *string                   `protobuf:"bytes,20,opt,name=hero_loop_appear,json=heroLoopAppear" json:"hero_loop_appear,omitempty"`
	HeroLoopDisappear   *string                   `protobuf:"bytes,21,opt,name=hero_loop_disappear,json=heroLoopDisappear" json:"hero_loop_disappear,omitempty"`
	MovementSpeed       *int32                    `protobuf:"varint,22,opt,name=movement_speed,json=movementSpeed" json:"movement_speed,omitempty"`
	Aura                *bool                     `protobuf:"varint,23,opt,name=aura" json:"aura,omitempty"`
	Activity            *int32                    `protobuf:"varint,24,opt,name=activity" json:"activity,omitempty"`
	Damage              *int32                    `protobuf:"varint,25,opt,name=damage" json:"damage,omitempty"`
	Range               *int32                    `protobuf:"varint,26,opt,name=range" json:"range,omitempty"`
	DdModifierIndex     *int32                    `protobuf:"varint,27,opt,name=dd_modifier_index,json=ddModifierIndex" json:"dd_modifier_index,omitempty"`
	DdAbilityId         *int32                    `protobuf:"varint,28,opt,name=dd_ability_id,json=ddAbilityId" json:"dd_ability_id,omitempty"`
	IllusionLabel       *string                   `protobuf:"bytes,29,opt,name=illusion_label,json=illusionLabel" json:"illusion_label,omitempty"`
	Active              *bool                     `protobuf:"varint,30,opt,name=active" json:"active,omitempty"`
	PlayerIds           *string                   `protobuf:"bytes,31,opt,name=player_ids,json=playerIds" json:"player_ids,omitempty"`
	LuaName             *string                   `protobuf:"bytes,32,opt,name=lua_name,json=luaName" json:"lua_name,omitempty"`
	AttackSpeed         *int32                    `protobuf:"varint,33,opt,name=attack_speed,json=attackSpeed" json:"attack_speed,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *CDOTAModifierBuffTableEntry) Reset()                    { *m = CDOTAModifierBuffTableEntry{} }
func (m *CDOTAModifierBuffTableEntry) String() string            { return proto.CompactTextString(m) }
func (*CDOTAModifierBuffTableEntry) ProtoMessage()               {}
func (*CDOTAModifierBuffTableEntry) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

const Default_CDOTAModifierBuffTableEntry_EntryType DOTA_MODIFIER_ENTRY_TYPE = DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE
const Default_CDOTAModifierBuffTableEntry_Duration float32 = -1

func (m *CDOTAModifierBuffTableEntry) GetEntryType() DOTA_MODIFIER_ENTRY_TYPE {
	if m != nil && m.EntryType != nil {
		return *m.EntryType
	}
	return Default_CDOTAModifierBuffTableEntry_EntryType
}

func (m *CDOTAModifierBuffTableEntry) GetParent() int32 {
	if m != nil && m.Parent != nil {
		return *m.Parent
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSerialNum() int32 {
	if m != nil && m.SerialNum != nil {
		return *m.SerialNum
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetModifierClass() int32 {
	if m != nil && m.ModifierClass != nil {
		return *m.ModifierClass
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbilityLevel() int32 {
	if m != nil && m.AbilityLevel != nil {
		return *m.AbilityLevel
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetStackCount() int32 {
	if m != nil && m.StackCount != nil {
		return *m.StackCount
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetCreationTime() float32 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return Default_CDOTAModifierBuffTableEntry_Duration
}

func (m *CDOTAModifierBuffTableEntry) GetCaster() int32 {
	if m != nil && m.Caster != nil {
		return *m.Caster
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbility() int32 {
	if m != nil && m.Ability != nil {
		return *m.Ability
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetArmor() int32 {
	if m != nil && m.Armor != nil {
		return *m.Armor
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetFadeTime() float32 {
	if m != nil && m.FadeTime != nil {
		return *m.FadeTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSubtle() bool {
	if m != nil && m.Subtle != nil {
		return *m.Subtle
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetChannelTime() float32 {
	if m != nil && m.ChannelTime != nil {
		return *m.ChannelTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetVStart() *CMsgVector {
	if m != nil {
		return m.VStart
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetVEnd() *CMsgVector {
	if m != nil {
		return m.VEnd
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopAppear() string {
	if m != nil && m.PortalLoopAppear != nil {
		return *m.PortalLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopDisappear() string {
	if m != nil && m.PortalLoopDisappear != nil {
		return *m.PortalLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopAppear() string {
	if m != nil && m.HeroLoopAppear != nil {
		return *m.HeroLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopDisappear() string {
	if m != nil && m.HeroLoopDisappear != nil {
		return *m.HeroLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetMovementSpeed() int32 {
	if m != nil && m.MovementSpeed != nil {
		return *m.MovementSpeed
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAura() bool {
	if m != nil && m.Aura != nil {
		return *m.Aura
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetActivity() int32 {
	if m != nil && m.Activity != nil {
		return *m.Activity
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetRange() int32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdModifierIndex() int32 {
	if m != nil && m.DdModifierIndex != nil {
		return *m.DdModifierIndex
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdAbilityId() int32 {
	if m != nil && m.DdAbilityId != nil {
		return *m.DdAbilityId
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetIllusionLabel() string {
	if m != nil && m.IllusionLabel != nil {
		return *m.IllusionLabel
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetPlayerIds() string {
	if m != nil && m.PlayerIds != nil {
		return *m.PlayerIds
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetLuaName() string {
	if m != nil && m.LuaName != nil {
		return *m.LuaName
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetAttackSpeed() int32 {
	if m != nil && m.AttackSpeed != nil {
		return *m.AttackSpeed
	}
	return 0
}

type CDOTALuaModifierEntry struct {
	ModifierType     *int32  `protobuf:"varint,1,req,name=modifier_type,json=modifierType" json:"modifier_type,omitempty"`
	ModifierFilename *string `protobuf:"bytes,2,req,name=modifier_filename,json=modifierFilename" json:"modifier_filename,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTALuaModifierEntry) Reset()                    { *m = CDOTALuaModifierEntry{} }
func (m *CDOTALuaModifierEntry) String() string            { return proto.CompactTextString(m) }
func (*CDOTALuaModifierEntry) ProtoMessage()               {}
func (*CDOTALuaModifierEntry) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *CDOTALuaModifierEntry) GetModifierType() int32 {
	if m != nil && m.ModifierType != nil {
		return *m.ModifierType
	}
	return 0
}

func (m *CDOTALuaModifierEntry) GetModifierFilename() string {
	if m != nil && m.ModifierFilename != nil {
		return *m.ModifierFilename
	}
	return ""
}

func init() {
	proto.RegisterType((*CDOTAModifierBuffTableEntry)(nil), "dota.CDOTAModifierBuffTableEntry")
	proto.RegisterType((*CDOTALuaModifierEntry)(nil), "dota.CDOTALuaModifierEntry")
	proto.RegisterEnum("dota.DOTA_MODIFIER_ENTRY_TYPE", DOTA_MODIFIER_ENTRY_TYPE_name, DOTA_MODIFIER_ENTRY_TYPE_value)
}

func init() { proto.RegisterFile("dota_modifiers.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4b, 0x6f, 0x1b, 0x37,
	0x10, 0xc7, 0xbb, 0xaa, 0x65, 0x4b, 0x23, 0xc9, 0x91, 0xe9, 0x47, 0x19, 0xbb, 0xb1, 0x37, 0x76,
	0x03, 0xa8, 0x69, 0x6b, 0xa0, 0x3e, 0xe6, 0xe6, 0xc8, 0x0a, 0x2a, 0xc0, 0x8f, 0x62, 0x23, 0x18,
	0xc8, 0xa5, 0x04, 0xb5, 0x1c, 0x39, 0x8b, 0x70, 0x1f, 0xe0, 0x72, 0xd5, 0xea, 0xd6, 0xef, 0xd1,
	0x2f, 0x5b, 0x70, 0xb8, 0x2b, 0xa5, 0x07, 0xa3, 0xb7, 0x9d, 0xdf, 0xfc, 0x67, 0x76, 0x48, 0xce,
	0x0c, 0x1c, 0xa8, 0xdc, 0x4a, 0x91, 0xe6, 0x2a, 0x59, 0x24, 0x68, 0xca, 0xcb, 0xc2, 0xe4, 0x36,
	0x67, 0x5b, 0x8e, 0x1e, 0x1f, 0x65, 0x68, 0xff, 0xcc, 0xcd, 0x97, 0xb9, 0x2c, 0xd1, 0xae, 0x0a,
	0xac, 0xbd, 0xe7, 0xff, 0x74, 0xe1, 0x64, 0x7c, 0xf3, 0x30, 0xbb, 0xbe, 0xab, 0xc3, 0xde, 0x57,
	0x8b, 0xc5, 0x4c, 0xce, 0x35, 0x4e, 0x32, 0x6b, 0x56, 0xec, 0x0f, 0x00, 0x74, 0x1f, 0xc2, 0x05,
	0xf1, 0x20, 0x6c, 0x8d, 0x76, 0xaf, 0x4e, 0x2f, 0x5d, 0xca, 0x4b, 0x17, 0x25, 0xee, 0x1e, 0x6e,
	0xa6, 0x1f, 0xa6, 0x93, 0x48, 0x4c, 0xee, 0x67, 0xd1, 0x27, 0x31, 0xfb, 0xf4, 0xfb, 0xe4, 0xdd,
	0xd9, 0x73, 0x1e, 0x71, 0x3d, 0x9e, 0x4d, 0x1f, 0x27, 0x51, 0x97, 0x52, 0xce, 0x56, 0x05, 0xb2,
	0x23, 0xd8, 0x2e, 0xa4, 0xc1, 0xcc, 0xf2, 0x56, 0xd8, 0x1a, 0xb5, 0xa3, 0xda, 0x62, 0x07, 0xd0,
	0x4e, 0x32, 0x85, 0x7f, 0xf1, 0x6f, 0x09, 0x7b, 0x83, 0xbd, 0x02, 0x28, 0xd1, 0x24, 0x52, 0x8b,
	0xac, 0x4a, 0xf9, 0x16, 0xb9, 0xba, 0x9e, 0xdc, 0x57, 0x29, 0x7b, 0x03, 0xbb, 0xcd, 0xe9, 0x45,
	0xac, 0x65, 0x59, 0xf2, 0x76, 0x18, 0x8c, 0xda, 0xd1, 0xa0, 0xa1, 0x63, 0x07, 0xd9, 0x05, 0x0c,
	0xe4, 0x3c, 0xd1, 0x89, 0x5d, 0x09, 0x8d, 0x4b, 0xd4, 0x7c, 0x9b, 0x54, 0xfd, 0x1a, 0xde, 0x3a,
	0xc6, 0xce, 0xa0, 0x57, 0x5a, 0x19, 0x7f, 0x11, 0x71, 0x5e, 0x65, 0x96, 0xef, 0x90, 0x04, 0x08,
	0x8d, 0x1d, 0x71, 0x59, 0x62, 0x83, 0xd2, 0x26, 0x79, 0x26, 0x6c, 0x92, 0x22, 0xef, 0x84, 0xc1,
	0xa8, 0x15, 0xf5, 0x1b, 0x38, 0x4b, 0x52, 0x64, 0xa7, 0xd0, 0x51, 0x95, 0x21, 0x9b, 0x77, 0x9d,
	0xff, 0x5d, 0xeb, 0x97, 0x5f, 0xa3, 0x35, 0x73, 0xc7, 0x8f, 0x65, 0x69, 0xd1, 0x70, 0xa0, 0x1f,
	0xd4, 0x16, 0xe3, 0xb0, 0x53, 0x57, 0xc3, 0x7b, 0xe4, 0x68, 0x4c, 0x77, 0x31, 0xd2, 0xa4, 0xb9,
	0xe1, 0x7d, 0xe2, 0xde, 0x60, 0x27, 0xd0, 0x5d, 0x48, 0x85, 0xbe, 0x90, 0x01, 0x15, 0xd2, 0x71,
	0x80, 0x8a, 0x38, 0x82, 0xed, 0xb2, 0x9a, 0x5b, 0x8d, 0x7c, 0x37, 0x0c, 0x46, 0x9d, 0xa8, 0xb6,
	0xd8, 0x6b, 0xe8, 0xc7, 0x9f, 0x65, 0x96, 0xa1, 0xf6, 0x71, 0x2f, 0x28, 0xae, 0x57, 0x33, 0x0a,
	0xfd, 0x11, 0x76, 0x96, 0xa2, 0xb4, 0xd2, 0x58, 0x3e, 0x0c, 0x83, 0x51, 0xef, 0x6a, 0xe8, 0xdf,
	0x7e, 0x7c, 0x57, 0x3e, 0x3d, 0x62, 0x6c, 0x73, 0x13, 0x6d, 0x2f, 0x3f, 0x3a, 0x3f, 0x7b, 0x03,
	0xed, 0xa5, 0xc0, 0x4c, 0xf1, 0xbd, 0x67, 0x84, 0x5b, 0xcb, 0x49, 0xa6, 0xd8, 0xcf, 0xc0, 0x8a,
	0xdc, 0x58, 0xa9, 0x85, 0xce, 0xf3, 0x42, 0xc8, 0xa2, 0x40, 0x69, 0x38, 0x0b, 0x83, 0x51, 0x37,
	0x1a, 0x7a, 0xcf, 0x6d, 0x9e, 0x17, 0xd7, 0xc4, 0xd9, 0x15, 0x1c, 0x7e, 0xad, 0x56, 0x49, 0x59,
	0x07, 0xec, 0x53, 0xc0, 0xfe, 0x26, 0xe0, 0xa6, 0x71, 0xb1, 0x11, 0x0c, 0x3f, 0xa3, 0xc9, 0xff,
	0x93, 0xff, 0x80, 0xe4, 0xbb, 0x8e, 0x7f, 0x95, 0xfd, 0x12, 0xf6, 0x37, 0xca, 0x4d, 0xee, 0x43,
	0x12, 0xef, 0x35, 0xe2, 0x4d, 0x66, 0xea, 0xaf, 0x25, 0xa6, 0x98, 0x59, 0x51, 0x16, 0x88, 0x8a,
	0x1f, 0x35, 0xfd, 0xe5, 0xe9, 0x47, 0x07, 0x19, 0x83, 0x2d, 0x59, 0x19, 0xc9, 0xbf, 0xa3, 0xdb,
	0xa6, 0x6f, 0x76, 0x0c, 0x1d, 0x19, 0xdb, 0x64, 0xe9, 0x5e, 0x94, 0x53, 0xd0, 0xda, 0x76, 0xef,
	0xa3, 0x64, 0x2a, 0x9f, 0x90, 0xbf, 0xf4, 0x4d, 0xe0, 0x2d, 0xf7, 0xd4, 0x46, 0x66, 0x4f, 0xc8,
	0x8f, 0xfd, 0x53, 0x93, 0xc1, 0xde, 0xc2, 0x9e, 0x52, 0xeb, 0x29, 0x17, 0x7e, 0x4a, 0x4e, 0x48,
	0xf1, 0x42, 0xa9, 0x66, 0x8c, 0xa7, 0x34, 0x2f, 0xe7, 0x30, 0x50, 0x4a, 0x34, 0xcd, 0x9e, 0x28,
	0xfe, 0x3d, 0xe9, 0x7a, 0x4a, 0x5d, 0x7b, 0x36, 0x55, 0xee, 0x50, 0x89, 0xd6, 0x55, 0xe9, 0xfa,
	0x58, 0xcb, 0x39, 0x6a, 0xfe, 0x8a, 0xce, 0x3f, 0x68, 0xe8, 0xad, 0x83, 0xae, 0x48, 0x2a, 0x18,
	0xf9, 0xa9, 0x6f, 0x22, 0x6f, 0xb9, 0x91, 0x2c, 0xb4, 0x5c, 0xb9, 0x4a, 0x54, 0xc9, 0xcf, 0x28,
	0xb4, 0xeb, 0xc9, 0x54, 0x95, 0xec, 0x25, 0x74, 0x74, 0x25, 0x45, 0x26, 0x53, 0xe4, 0x21, 0x39,
	0x77, 0x74, 0x25, 0xef, 0x65, 0x4a, 0xed, 0x27, 0x2d, 0x8d, 0x98, 0xbf, 0xcb, 0xd7, 0xbe, 0x36,
	0xcf, 0xe8, 0x26, 0xcf, 0x13, 0x38, 0xa4, 0xe5, 0x74, 0x5b, 0xc9, 0xe6, 0x60, 0x7e, 0x2d, 0x5d,
	0xc0, 0x7a, 0xa6, 0x37, 0x9b, 0xa9, 0x1d, 0xf5, 0x1b, 0x48, 0xbb, 0xe5, 0x27, 0xd8, 0x5b, 0x8b,
	0x16, 0x89, 0x46, 0x2a, 0xc2, 0xad, 0x99, 0x6e, 0x34, 0x6c, 0x1c, 0x1f, 0x6a, 0xfe, 0x16, 0x81,
	0x3f, 0xb7, 0xb6, 0xd8, 0x05, 0xfc, 0xdf, 0x4a, 0x1b, 0x06, 0xec, 0x07, 0x08, 0x9f, 0x15, 0x45,
	0x93, 0xbb, 0x87, 0xc7, 0xc9, 0xcd, 0xb0, 0xf5, 0xbe, 0xfd, 0x5b, 0xf0, 0x77, 0xf0, 0xcd, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0xdc, 0x7a, 0x30, 0xab, 0x05, 0x00, 0x00,
}
