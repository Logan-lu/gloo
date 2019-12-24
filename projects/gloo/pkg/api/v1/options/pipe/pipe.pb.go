// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/pipe/pipe.proto

package pipe

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
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

// Pipe upstreams are used to route request to services listening at a Unix Domain Socket.
// Pipe upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Pipe Upstreams must be created manually by users
type UpstreamSpec struct {
	// The Unix Domain Socket path.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *options.ServiceSpec `protobuf:"bytes,2,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb91cb4d828830e0, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "pipe.options.gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/pipe/pipe.proto", fileDescriptor_fb91cb4d828830e0)
}

var fileDescriptor_fb91cb4d828830e0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33,
	0xd4, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0xc8, 0x2c, 0x48, 0x05, 0x13, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x60, 0x36, 0x54, 0x56, 0x0f, 0xa4, 0x43, 0xaf, 0x38,
	0x3f, 0x27, 0x5f, 0x2f, 0x33, 0x5f, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x4a, 0x1f, 0xc4,
	0x82, 0x68, 0x90, 0x12, 0x4a, 0xad, 0x28, 0x81, 0x08, 0xa6, 0x56, 0x94, 0x40, 0xc5, 0x4c, 0xc0,
	0x36, 0xe1, 0xb5, 0xb7, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x35, 0xbe, 0xb8, 0x20, 0x35, 0x19,
	0xa2, 0x4b, 0x29, 0x83, 0x8b, 0x27, 0xb4, 0xa0, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x37, 0xb8, 0x20,
	0x35, 0x59, 0x48, 0x88, 0x8b, 0xa5, 0x20, 0xb1, 0x24, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xcc, 0x16, 0x72, 0xe1, 0xe2, 0x41, 0xd6, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4,
	0x88, 0xd5, 0xc1, 0x7a, 0xc1, 0x10, 0x95, 0x20, 0xc3, 0x82, 0xb8, 0x8b, 0x11, 0x1c, 0x27, 0xf7,
	0x1d, 0x5f, 0x59, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x8c, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x69, 0xd3, 0xcd, 0xcc, 0xd7, 0xc7, 0xe2, 0xf8, 0x82, 0xec,
	0x74, 0x6c, 0x01, 0x97, 0xc4, 0x06, 0x76, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x82, 0x75,
	0x93, 0xc6, 0x64, 0x01, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
