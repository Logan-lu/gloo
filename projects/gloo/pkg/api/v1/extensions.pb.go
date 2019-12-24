// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/extensions.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Extensions struct {
	Configs              map[string]*types.Struct `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Extensions) Reset()         { *m = Extensions{} }
func (m *Extensions) String() string { return proto.CompactTextString(m) }
func (*Extensions) ProtoMessage()    {}
func (*Extensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c5093805e1927d5, []int{0}
}
func (m *Extensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extensions.Unmarshal(m, b)
}
func (m *Extensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extensions.Marshal(b, m, deterministic)
}
func (m *Extensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extensions.Merge(m, src)
}
func (m *Extensions) XXX_Size() int {
	return xxx_messageInfo_Extensions.Size(m)
}
func (m *Extensions) XXX_DiscardUnknown() {
	xxx_messageInfo_Extensions.DiscardUnknown(m)
}

var xxx_messageInfo_Extensions proto.InternalMessageInfo

func (m *Extensions) GetConfigs() map[string]*types.Struct {
	if m != nil {
		return m.Configs
	}
	return nil
}

type Extension struct {
	Config               *types.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Extension) Reset()         { *m = Extension{} }
func (m *Extension) String() string { return proto.CompactTextString(m) }
func (*Extension) ProtoMessage()    {}
func (*Extension) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c5093805e1927d5, []int{1}
}
func (m *Extension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extension.Unmarshal(m, b)
}
func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
}
func (m *Extension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extension.Merge(m, src)
}
func (m *Extension) XXX_Size() int {
	return xxx_messageInfo_Extension.Size(m)
}
func (m *Extension) XXX_DiscardUnknown() {
	xxx_messageInfo_Extension.DiscardUnknown(m)
}

var xxx_messageInfo_Extension proto.InternalMessageInfo

func (m *Extension) GetConfig() *types.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Extensions)(nil), "gloo.solo.io.Extensions")
	proto.RegisterMapType((map[string]*types.Struct)(nil), "gloo.solo.io.Extensions.ConfigsEntry")
	proto.RegisterType((*Extension)(nil), "gloo.solo.io.Extension")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/extensions.proto", fileDescriptor_4c5093805e1927d5)
}

var fileDescriptor_4c5093805e1927d5 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x56, 0x2b, 0xdd, 0xf6, 0x20, 0x8b, 0x60, 0x08, 0x22, 0xa1, 0x50, 0xc8, 0xa5,
	0x33, 0x5a, 0x2f, 0x52, 0x04, 0x41, 0xe9, 0x0b, 0xa4, 0x37, 0x6f, 0x4d, 0xd8, 0xae, 0x6b, 0x63,
	0x26, 0x64, 0x37, 0x25, 0x7d, 0x1d, 0x4f, 0x3e, 0x82, 0xcf, 0xe3, 0x3b, 0x78, 0x97, 0x64, 0xd3,
	0x5a, 0x2f, 0xbd, 0x0d, 0xf3, 0x7f, 0xfb, 0xff, 0xfc, 0x3b, 0x7c, 0x9c, 0x17, 0xf4, 0x26, 0x13,
	0x6b, 0x50, 0xa5, 0x44, 0xb8, 0xcc, 0x35, 0x6e, 0x6e, 0x51, 0x56, 0x56, 0x66, 0x46, 0x53, 0x66,
	0x20, 0x2f, 0xc8, 0x92, 0x18, 0xd6, 0x2a, 0x18, 0x4a, 0x09, 0x34, 0xf9, 0x57, 0x8a, 0x48, 0xa5,
	0x12, 0x1b, 0x2d, 0x2e, 0x57, 0x68, 0x6c, 0x51, 0x26, 0xd6, 0xb1, 0xfe, 0x85, 0x22, 0x45, 0xcd,
	0x88, 0xf5, 0xd4, 0x6e, 0x85, 0xac, 0xac, 0x5b, 0xca, 0xaa, 0x25, 0x47, 0x1f, 0x8c, 0xf3, 0xf9,
	0x3e, 0x4a, 0x3c, 0xf2, 0xb3, 0x84, 0xb2, 0x95, 0x56, 0xc6, 0x63, 0x41, 0x37, 0x1c, 0x4c, 0xc7,
	0x70, 0x18, 0x0b, 0x7f, 0x28, 0x3c, 0x3b, 0x6e, 0x9e, 0xd9, 0x62, 0x1b, 0xed, 0x5e, 0xf9, 0x0b,
	0x3e, 0x3c, 0x14, 0xc4, 0x39, 0xef, 0xae, 0xe5, 0xd6, 0x63, 0x01, 0x0b, 0xfb, 0x51, 0x3d, 0x8a,
	0x09, 0x3f, 0xdd, 0x2c, 0xd3, 0x52, 0x7a, 0x9d, 0x80, 0x85, 0x83, 0xe9, 0x25, 0xb8, 0x26, 0xb0,
	0x6b, 0x02, 0x8b, 0xa6, 0x49, 0xe4, 0xa8, 0x59, 0xe7, 0x9e, 0x8d, 0x1e, 0x78, 0x7f, 0x1f, 0x2c,
	0x90, 0xf7, 0x5c, 0x58, 0x63, 0x7a, 0xc4, 0xa0, 0xc5, 0x9e, 0x66, 0x5f, 0x3f, 0x27, 0xec, 0xf3,
	0xfb, 0x9a, 0xbd, 0xdc, 0x28, 0x6d, 0x5f, 0xcb, 0x18, 0x12, 0x7a, 0xc7, 0xba, 0xd1, 0x44, 0x93,
	0xfb, 0xf3, 0xff, 0x17, 0xc8, 0xd7, 0xaa, 0xbd, 0x42, 0xdc, 0x6b, 0x4c, 0xef, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x89, 0x27, 0xcf, 0x50, 0xa4, 0x01, 0x00, 0x00,
}

func (this *Extensions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extensions)
	if !ok {
		that2, ok := that.(Extensions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Configs) != len(that1.Configs) {
		return false
	}
	for i := range this.Configs {
		if !this.Configs[i].Equal(that1.Configs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Extension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extension)
	if !ok {
		that2, ok := that.(Extension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
