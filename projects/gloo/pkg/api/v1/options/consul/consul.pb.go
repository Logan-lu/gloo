// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto

package consul

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Upstream Spec for Consul Upstreams
// consul Upstreams represent a set of one or more addressable pods for a consul Service
// the Gloo consul Upstream maps to a single service port. Because consul Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// consul Upstreams are typically generated automatically by Gloo from the consul API
type UpstreamSpec struct {
	// The name of the Consul Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The list of service tags Gloo should search for on a service instance
	// before deciding whether or not to include the instance as part of this
	// upstream
	ServiceTags []string `protobuf:"bytes,2,rep,name=service_tags,json=serviceTags,proto3" json:"service_tags,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,3,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Is this consul service connect enabled.
	ConnectEnabled bool `protobuf:"varint,4,opt,name=connect_enabled,json=connectEnabled,proto3" json:"connect_enabled,omitempty"`
	// The data centers in which the service instance represented by this upstream is registered.
	DataCenters          []string `protobuf:"bytes,5,rep,name=data_centers,json=dataCenters,proto3" json:"data_centers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5077911f8bc0ad, []int{0}
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

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceTags() []string {
	if m != nil {
		return m.ServiceTags
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetConnectEnabled() bool {
	if m != nil {
		return m.ConnectEnabled
	}
	return false
}

func (m *UpstreamSpec) GetDataCenters() []string {
	if m != nil {
		return m.DataCenters
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "consul.options.gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto", fileDescriptor_3c5077911f8bc0ad)
}

var fileDescriptor_3c5077911f8bc0ad = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x65, 0x5a, 0x10, 0xa4, 0x15, 0x48, 0x11, 0x43, 0x54, 0x24, 0x94, 0xb2, 0xd0, 0x05,
	0x5b, 0xc0, 0x03, 0x20, 0xf1, 0x47, 0xc0, 0xc2, 0xd0, 0xc2, 0xc2, 0x12, 0x39, 0xee, 0xc9, 0x18,
	0x12, 0x9f, 0x15, 0xbb, 0x55, 0x1f, 0x89, 0x47, 0xe0, 0x79, 0x18, 0x78, 0x03, 0x76, 0xe4, 0xd8,
	0x40, 0x87, 0x0e, 0x88, 0x29, 0xb9, 0x4f, 0xdf, 0xdd, 0xfd, 0x72, 0x49, 0x6e, 0xa4, 0x72, 0x4f,
	0xb3, 0x92, 0x0a, 0xac, 0x99, 0xc5, 0x0a, 0x8f, 0x14, 0x32, 0x59, 0x21, 0x32, 0xd3, 0xe0, 0x33,
	0x08, 0x67, 0x43, 0xc5, 0x8d, 0x62, 0xf3, 0x63, 0x86, 0xc6, 0x29, 0xd4, 0x96, 0x09, 0xd4, 0x76,
	0x56, 0xc5, 0x07, 0x35, 0x0d, 0x3a, 0x4c, 0xf7, 0x62, 0x15, 0x1d, 0xea, 0xfb, 0xa8, 0x1f, 0x49,
	0x15, 0x0e, 0x76, 0x25, 0x4a, 0x6c, 0x3d, 0xe6, 0xdf, 0x42, 0xcb, 0x20, 0x85, 0x85, 0x0b, 0x10,
	0x16, 0x2e, 0xb2, 0xeb, 0x7f, 0x05, 0xb2, 0xd0, 0xcc, 0x95, 0x80, 0xc2, 0x1a, 0x10, 0x61, 0xd0,
	0xc1, 0x07, 0x49, 0xfa, 0x0f, 0xc6, 0xba, 0x06, 0x78, 0x3d, 0x31, 0x20, 0xd2, 0x61, 0xd2, 0xff,
	0xd6, 0x34, 0xaf, 0x21, 0x23, 0x39, 0x19, 0x6d, 0x8d, 0x7b, 0x91, 0xdd, 0xf1, 0x1a, 0x96, 0x15,
	0xc7, 0xa5, 0xcd, 0xd6, 0xf2, 0xce, 0x92, 0x72, 0xcf, 0xa5, 0x4d, 0x2f, 0x7f, 0x15, 0xbf, 0x2c,
	0xeb, 0xe4, 0x64, 0xd4, 0x3b, 0x19, 0xae, 0xfc, 0x6c, 0x3a, 0x09, 0xa6, 0x5f, 0xff, 0x33, 0xa5,
	0xcd, 0x72, 0x98, 0xec, 0x08, 0xd4, 0x1a, 0x84, 0x2b, 0x40, 0xf3, 0xb2, 0x82, 0x69, 0xd6, 0xcd,
	0xc9, 0x68, 0x73, 0xbc, 0x1d, 0xf1, 0x55, 0xa0, 0x3e, 0xd1, 0x94, 0x3b, 0x5e, 0x08, 0xd0, 0x0e,
	0x1a, 0x9b, 0xad, 0x87, 0x44, 0x9e, 0x5d, 0x04, 0x74, 0x7e, 0xfb, 0xf6, 0xd9, 0x25, 0xaf, 0xef,
	0xfb, 0xe4, 0xf1, 0xec, 0x6f, 0xb7, 0x33, 0x2f, 0x72, 0xf5, 0x0f, 0x2d, 0x37, 0xda, 0xd3, 0x9d,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x77, 0xcc, 0xaa, 0x1a, 0x16, 0x02, 0x00, 0x00,
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
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if len(this.ServiceTags) != len(that1.ServiceTags) {
		return false
	}
	for i := range this.ServiceTags {
		if this.ServiceTags[i] != that1.ServiceTags[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if this.ConnectEnabled != that1.ConnectEnabled {
		return false
	}
	if len(this.DataCenters) != len(that1.DataCenters) {
		return false
	}
	for i := range this.DataCenters {
		if this.DataCenters[i] != that1.DataCenters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
