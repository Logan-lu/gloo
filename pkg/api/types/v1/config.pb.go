// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	config.proto
	metadata.proto
	role.proto
	status.proto
	upstream.proto
	virtualservice.proto

It has these top-level messages:
	Config
	Metadata
	Role
	Status
	Upstream
	ServiceInfo
	Function
	VirtualService
	Route
	RequestMatcher
	EventMatcher
	WeightedDestination
	Destination
	FunctionDestination
	UpstreamDestination
	SSLConfig
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Config is a top-level config object. It is used internally by gloo as a container for the entire set of config objects.
type Config struct {
	Upstreams       []*Upstream       `protobuf:"bytes,1,rep,name=upstreams" json:"upstreams,omitempty"`
	VirtualServices []*VirtualService `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services,omitempty"`
	Roles           []*Role           `protobuf:"bytes,3,rep,name=roles" json:"roles,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetUpstreams() []*Upstream {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Config) GetVirtualServices() []*VirtualService {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Config) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "gloo.api.v1.Config")
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
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
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(that1.VirtualServices[i]) {
			return false
		}
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if !this.Roles[i].Equal(that1.Roles[i]) {
			return false
		}
	}
	return true
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xcf, 0xc9, 0xcf, 0xd7, 0x4b,
	0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0xe2, 0x2b, 0x2d, 0x28, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x85,
	0x48, 0x4a, 0x89, 0x94, 0x65, 0x16, 0x95, 0x94, 0x26, 0xe6, 0x14, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x42, 0x45, 0xb9, 0x8a, 0xf2, 0x73, 0x60, 0x6c, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53,
	0x1f, 0xc4, 0x82, 0x88, 0x2a, 0xad, 0x63, 0xe4, 0x62, 0x73, 0x06, 0xdb, 0x22, 0x64, 0xcc, 0xc5,
	0x09, 0x33, 0xb4, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x54, 0x0f, 0xc9, 0x4e, 0xbd,
	0x50, 0xa8, 0x6c, 0x10, 0x42, 0x9d, 0x90, 0x1b, 0x97, 0x00, 0xd4, 0xe6, 0x78, 0xa8, 0xd5, 0xc5,
	0x12, 0x4c, 0x60, 0xbd, 0xd2, 0x28, 0x7a, 0xc3, 0x20, 0x8a, 0x82, 0x21, 0x6a, 0x82, 0xf8, 0xcb,
	0x50, 0xf8, 0xc5, 0x42, 0xea, 0x5c, 0xac, 0x20, 0xb7, 0x16, 0x4b, 0x30, 0x83, 0x35, 0x0b, 0xa2,
	0x68, 0x0e, 0xca, 0xcf, 0x49, 0x0d, 0x82, 0xc8, 0x3b, 0xe9, 0xad, 0x78, 0x24, 0xc7, 0x18, 0xa5,
	0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x9c, 0x9f, 0x93, 0xaf,
	0x9b, 0x99, 0xaf, 0x0f, 0xd2, 0xa1, 0x5f, 0x90, 0x9d, 0xae, 0x9f, 0x58, 0x90, 0xa9, 0x5f, 0x52,
	0x59, 0x90, 0x5a, 0xac, 0x5f, 0x66, 0x98, 0xc4, 0x06, 0xf6, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xad, 0x59, 0x13, 0x02, 0x4c, 0x01, 0x00, 0x00,
}
