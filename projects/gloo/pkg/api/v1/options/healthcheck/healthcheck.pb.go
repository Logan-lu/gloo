// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/healthcheck/healthcheck.proto

package healthcheck

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Add this config to a Listener/Gateway to Enable Envoy Health Checks on that port
type HealthCheck struct {
	// match health check requests using this exact path
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee028ea519a37577, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "healthcheck.options.gloo.solo.io.HealthCheck")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/healthcheck/healthcheck.proto", fileDescriptor_ee028ea519a37577)
}

var fileDescriptor_ee028ea519a37577 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x54, 0xa4, 0x13, 0x12, 0xa1, 0x8b, 0x28, 0xe0, 0x8a, 0xd3, 0x41, 0x45, 0x83, 0x57, 0x88,
	0x8e, 0x12, 0x28, 0xa8, 0x28, 0x28, 0xe9, 0x92, 0x68, 0xcf, 0x36, 0x09, 0xcc, 0xca, 0xde, 0x83,
	0x7b, 0x12, 0x4f, 0xe0, 0x3d, 0xfc, 0x81, 0x1e, 0xd9, 0x0e, 0x52, 0x68, 0xae, 0x1b, 0xef, 0xec,
	0x78, 0x66, 0xa7, 0xbe, 0x91, 0x80, 0x17, 0xee, 0x35, 0x92, 0x1d, 0x01, 0x6a, 0xc5, 0xd3, 0xfb,
	0x15, 0x41, 0xd4, 0xe3, 0x2d, 0x92, 0xe3, 0x76, 0x54, 0xd7, 0x3b, 0xee, 0x87, 0x39, 0x36, 0x12,
	0xa0, 0x68, 0xd6, 0xf3, 0xd1, 0x24, 0x31, 0xe9, 0x1b, 0x13, 0x31, 0xc2, 0x78, 0x2c, 0x8f, 0x2d,
	0x2c, 0xf2, 0x32, 0x25, 0x54, 0x74, 0xcb, 0xd3, 0x44, 0x5f, 0x0e, 0x5e, 0xff, 0xec, 0x02, 0x6f,
	0x26, 0xaa, 0xe1, 0x9d, 0x96, 0x7d, 0xde, 0xe9, 0x34, 0x5b, 0x59, 0xc0, 0x8e, 0x4c, 0xf9, 0xd5,
	0x6d, 0x37, 0xf4, 0x11, 0x5a, 0x11, 0x0e, 0xb1, 0xf0, 0xe7, 0x67, 0xf5, 0xd1, 0x43, 0x0e, 0x72,
	0x97, 0x82, 0x34, 0x4d, 0xbd, 0x90, 0x56, 0xdd, 0x49, 0xb5, 0xae, 0x2e, 0x0e, 0x9f, 0x32, 0xbe,
	0x7d, 0xfc, 0xfa, 0x59, 0x54, 0x9f, 0xdf, 0xab, 0xea, 0xf9, 0xde, 0x7a, 0x75, 0xdb, 0xce, 0xf4,
	0x78, 0xa5, 0x1c, 0xc3, 0xa3, 0x5c, 0xfe, 0xbf, 0x07, 0x19, 0xec, 0x9e, 0x2e, 0xba, 0x83, 0xec,
	0x7c, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xca, 0x9d, 0x11, 0xa6, 0x3e, 0x01, 0x00, 0x00,
}

func (this *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
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
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
