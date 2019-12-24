// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/kubernetes/kubernetes.proto

package kubernetes

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

// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The access port of the kubernetes service is listening. This port is used by Gloo to look up the corresponding
	// port on the pod for routing.
	ServicePort uint32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Allows finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Subset configuration. For discovery sources that has labels (like kubernetes). this
	// configuration allows you to partition the upstream to a set of subsets.
	// for each unique set of keys and values, a subset will be created.
	SubsetSpec           *options.SubsetSpec `protobuf:"bytes,6,opt,name=subset_spec,json=subsetSpec,proto3" json:"subset_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d279bec831b81d5, []int{0}
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

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetSubsetSpec() *options.SubsetSpec {
	if m != nil {
		return m.SubsetSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "kubernetes.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterMapType((map[string]string)(nil), "kubernetes.options.gloo.solo.io.UpstreamSpec.SelectorEntry")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/kubernetes/kubernetes.proto", fileDescriptor_1d279bec831b81d5)
}

var fileDescriptor_1d279bec831b81d5 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xaa, 0xda, 0x40,
	0x14, 0xc6, 0x89, 0x51, 0x69, 0x27, 0x0a, 0x76, 0x70, 0x11, 0x5c, 0xb4, 0xb1, 0xab, 0x40, 0xe9,
	0x0c, 0xd5, 0x2e, 0xa4, 0xae, 0xfa, 0x6f, 0xd9, 0x52, 0x94, 0x52, 0xe8, 0xa6, 0x24, 0xe1, 0x90,
	0xa6, 0x26, 0x39, 0xc3, 0xcc, 0x44, 0xf4, 0x6d, 0xba, 0xec, 0x23, 0xf4, 0x79, 0xee, 0x3b, 0xdc,
	0xfd, 0x25, 0x93, 0x98, 0x3b, 0x17, 0xf4, 0xde, 0xbb, 0x3b, 0x7e, 0x7c, 0xdf, 0xef, 0x1c, 0xbf,
	0x09, 0x59, 0x09, 0x89, 0x7f, 0x20, 0xd1, 0x8a, 0xa7, 0x39, 0x22, 0x8f, 0x44, 0xc6, 0xf7, 0x6f,
	0x38, 0x0a, 0x9d, 0x61, 0xa9, 0xf8, 0xae, 0x8a, 0x41, 0x96, 0xa0, 0xc1, 0x1e, 0x99, 0x90, 0xa8,
	0x91, 0xbe, 0xb0, 0x94, 0xd6, 0xcf, 0x6a, 0x06, 0x53, 0x98, 0x23, 0xcb, 0x70, 0x36, 0x4d, 0x31,
	0x45, 0xe3, 0xe5, 0xf5, 0xd4, 0xc4, 0x66, 0x14, 0x0e, 0xba, 0x11, 0xe1, 0xa0, 0x5b, 0xed, 0xad,
	0xd9, 0x7d, 0xef, 0x25, 0x0a, 0xe4, 0x3e, 0x4b, 0xe0, 0x97, 0x12, 0x90, 0xb4, 0xa9, 0xe5, 0x23,
	0x52, 0x55, 0xac, 0x40, 0x5b, 0xa1, 0x97, 0x7f, 0x5d, 0x32, 0xfa, 0x2e, 0x94, 0x96, 0x10, 0x15,
	0x5b, 0x01, 0x09, 0x9d, 0x93, 0xd1, 0x89, 0x5d, 0x46, 0x05, 0xf8, 0x4e, 0xe0, 0x84, 0x4f, 0x37,
	0x5e, 0xab, 0x7d, 0x8d, 0x0a, 0xa0, 0xaf, 0xc8, 0x33, 0xdb, 0xa2, 0x44, 0x94, 0x80, 0xdf, 0x33,
	0xbe, 0x89, 0xe5, 0x33, 0xba, 0xcd, 0x13, 0x28, 0xb5, 0xef, 0x06, 0x4e, 0x38, 0xee, 0x78, 0xdf,
	0x50, 0x6a, 0xfa, 0x83, 0x3c, 0x51, 0x90, 0x43, 0xa2, 0x51, 0xfa, 0xfd, 0xc0, 0x0d, 0xbd, 0xc5,
	0x9a, 0x3d, 0x50, 0x26, 0xb3, 0x6f, 0x66, 0xdb, 0x36, 0xfd, 0xb9, 0xd4, 0xf2, 0xb8, 0xe9, 0x60,
	0xf4, 0xd3, 0xed, 0xee, 0xfa, 0x2f, 0xfb, 0x83, 0xc0, 0x09, 0xbd, 0xc5, 0xfc, 0x3c, 0x71, 0xdb,
	0x38, 0x6b, 0x60, 0x77, 0x9e, 0x69, 0xe4, 0x3d, 0xf1, 0xac, 0xde, 0xfc, 0xa1, 0x81, 0x04, 0x17,
	0x20, 0xc6, 0x68, 0x18, 0x44, 0x75, 0xf3, 0x6c, 0x4d, 0xc6, 0x77, 0x6e, 0xa4, 0x13, 0xe2, 0xee,
	0xe0, 0xd8, 0x96, 0x5b, 0x8f, 0x74, 0x4a, 0x06, 0xfb, 0x28, 0xaf, 0x4e, 0x45, 0x36, 0x3f, 0xde,
	0xf5, 0x56, 0xce, 0x87, 0x2f, 0xff, 0xaf, 0xfb, 0xce, 0xbf, 0xab, 0xe7, 0xce, 0xcf, 0x8f, 0x69,
	0xa6, 0x7f, 0x57, 0x31, 0x4b, 0xb0, 0xe0, 0xf5, 0xc6, 0xd7, 0x19, 0xf2, 0x33, 0x8f, 0x2e, 0x76,
	0xe9, 0xe5, 0x0f, 0x37, 0x1e, 0x9a, 0x87, 0x5f, 0xde, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x8c,
	0x32, 0x2c, 0xea, 0x02, 0x00, 0x00,
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
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !this.SubsetSpec.Equal(that1.SubsetSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
