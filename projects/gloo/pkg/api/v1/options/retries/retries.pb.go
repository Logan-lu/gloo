// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/ext"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Retry Policy applied at the Route and/or Virtual Hosts levels.
type RetryPolicy struct {
	// Specifies the conditions under which retry takes place. These are the same
	// conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	RetryOn string `protobuf:"bytes,1,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Specifies the allowed number of retries. This parameter is optional and
	// defaults to 1. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	NumRetries uint32 `protobuf:"varint,2,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"`
	// Specifies a non-zero upstream timeout per retry attempt. This parameter is optional.
	PerTryTimeout        *time.Duration `protobuf:"bytes,3,opt,name=per_try_timeout,json=perTryTimeout,proto3,stdduration" json:"per_try_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RetryPolicy) Reset()         { *m = RetryPolicy{} }
func (m *RetryPolicy) String() string { return proto.CompactTextString(m) }
func (*RetryPolicy) ProtoMessage()    {}
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c06018876f3ed3e, []int{0}
}
func (m *RetryPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryPolicy.Unmarshal(m, b)
}
func (m *RetryPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryPolicy.Marshal(b, m, deterministic)
}
func (m *RetryPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryPolicy.Merge(m, src)
}
func (m *RetryPolicy) XXX_Size() int {
	return xxx_messageInfo_RetryPolicy.Size(m)
}
func (m *RetryPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RetryPolicy proto.InternalMessageInfo

func (m *RetryPolicy) GetRetryOn() string {
	if m != nil {
		return m.RetryOn
	}
	return ""
}

func (m *RetryPolicy) GetNumRetries() uint32 {
	if m != nil {
		return m.NumRetries
	}
	return 0
}

func (m *RetryPolicy) GetPerTryTimeout() *time.Duration {
	if m != nil {
		return m.PerTryTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RetryPolicy)(nil), "retries.options.gloo.solo.io.RetryPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto", fileDescriptor_3c06018876f3ed3e)
}

var fileDescriptor_3c06018876f3ed3e = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x89, 0x8a, 0x3f, 0x29, 0x45, 0x28, 0x2e, 0x3a, 0x83, 0xcc, 0x14, 0x57, 0xdd, 0x98,
	0xa0, 0xbe, 0x80, 0x14, 0x41, 0x74, 0xa3, 0x94, 0x59, 0xb9, 0x29, 0xd3, 0x1a, 0x63, 0xb4, 0xed,
	0x0d, 0xe9, 0x8d, 0x4c, 0x5f, 0xc3, 0x95, 0x8f, 0xe0, 0x5b, 0x09, 0x3e, 0x85, 0x4b, 0x49, 0xd3,
	0x71, 0x27, 0xcc, 0x2a, 0xf7, 0x5c, 0xce, 0xc9, 0xfd, 0x38, 0xf4, 0x56, 0x2a, 0x7c, 0xb6, 0x25,
	0xab, 0xa0, 0xe1, 0x1d, 0xd4, 0x70, 0xaa, 0x80, 0xcb, 0x1a, 0x80, 0x6b, 0x03, 0x2f, 0xa2, 0xc2,
	0xce, 0xab, 0xa5, 0x56, 0xfc, 0xed, 0x8c, 0x83, 0x46, 0x05, 0x6d, 0xc7, 0x8d, 0x40, 0xa3, 0xc4,
	0xdf, 0xcb, 0xb4, 0x01, 0x84, 0xe8, 0x78, 0x2d, 0x47, 0x1b, 0x73, 0x51, 0xe6, 0x7e, 0x65, 0x0a,
	0xa6, 0x33, 0x09, 0x20, 0x6b, 0xc1, 0x07, 0x6f, 0x69, 0x9f, 0xf8, 0xa3, 0x35, 0x4b, 0xe7, 0xf3,
	0xe9, 0xe9, 0x91, 0x04, 0x09, 0xc3, 0xc8, 0xdd, 0x34, 0x6e, 0x43, 0xb1, 0x42, 0x2e, 0x56, 0xe8,
	0xe5, 0xc9, 0x3b, 0xa1, 0x41, 0x2e, 0xd0, 0xf4, 0xf7, 0x50, 0xab, 0xaa, 0x8f, 0x26, 0x74, 0xdf,
	0x1d, 0xed, 0x0b, 0x68, 0x63, 0x92, 0x90, 0xf4, 0x20, 0xdf, 0x1b, 0xf4, 0x5d, 0x1b, 0xcd, 0x69,
	0xd0, 0xda, 0xa6, 0x18, 0x99, 0xe2, 0xad, 0x84, 0xa4, 0x61, 0x4e, 0x5b, 0xdb, 0xe4, 0x7e, 0x13,
	0x5d, 0xd3, 0x43, 0x2d, 0x4c, 0xe1, 0xd2, 0xa8, 0x1a, 0x01, 0x16, 0xe3, 0xed, 0x84, 0xa4, 0xc1,
	0xf9, 0x84, 0x79, 0x54, 0xb6, 0x46, 0x65, 0x57, 0x23, 0x6a, 0xb6, 0xf3, 0xf1, 0x35, 0x27, 0x79,
	0xa8, 0x85, 0x59, 0x98, 0x7e, 0xe1, 0x53, 0xd9, 0xcd, 0x4f, 0x46, 0x3e, 0xbf, 0x67, 0xe4, 0xe1,
	0x72, 0xb3, 0x32, 0xf5, 0xab, 0xfc, 0xa7, 0xd0, 0x72, 0x77, 0x38, 0x79, 0xf1, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xe6, 0xf1, 0x26, 0xae, 0x97, 0x01, 0x00, 0x00,
}

func (this *RetryPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RetryPolicy)
	if !ok {
		that2, ok := that.(RetryPolicy)
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
	if this.RetryOn != that1.RetryOn {
		return false
	}
	if this.NumRetries != that1.NumRetries {
		return false
	}
	if this.PerTryTimeout != nil && that1.PerTryTimeout != nil {
		if *this.PerTryTimeout != *that1.PerTryTimeout {
			return false
		}
	} else if this.PerTryTimeout != nil {
		return false
	} else if that1.PerTryTimeout != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
