// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gateway/api/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//
// The **RouteTable** is a child Routing object for the Gloo Gateway.
//
// A **RouteTable** gets built into the complete routing configuration
// when it is referenced by a `delegateAction`, either
// in a parent VirtualService or another RouteTable.
//
// Routes specified in a RouteTable must have their paths start with the prefix provided in the
// parent's matcher.
//
// For example, the following configuration:
//
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /1
//
// ```
//
// would *not be valid*, while
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /a/1
//
// ```
//
// *would* be valid.
//
//
// A complete configuration might look as follows:
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'any'
//   namespace: 'any'
// spec:
//   virtualHost:
//     domains:
//     - 'any.com'
//     routes:
//     - matcher:
//         prefix: '/a' # delegate ownership of routes for `any.com/a`
//       delegateAction:
//         name: 'a-routes'
//         namespace: 'a'
//     - matcher:
//         prefix: '/b' # delegate ownership of routes for `any.com/b`
//       delegateAction:
//         name: 'b-routes'
//         namespace: 'b'
// ```
//
// * A root-level **VirtualService** which delegates routing to to the `a-routes` and `b-routes` **RouteTables**.
// * Routes with `delegateActions` can only use a `prefix` matcher.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'a-routes'
//   namespace: 'a'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/a/`
//         prefix: '/a/1'
//       routeAction:
//         single:
//           upstream:
//             name: 'foo-upstream'
//
//     - matcher:
//         prefix: '/a/2'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
// ```
//
// * A **RouteTable** which defines two routes.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'b-routes'
//   namespace: 'b'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/b/`
//         regex: '/b/3'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
//     - matcher:
//         prefix: '/b/c/'
//       # routes in the RouteTable can perform any action, including a delegateAction
//       delegateAction:
//         name: 'c-routes'
//         namespace: 'c'
//
// ```
//
// * A **RouteTable** which both *defines a route* and *delegates to* another **RouteTable**.
//
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'c-routes'
//   namespace: 'c'
// spec:
//   routes:
//     - matcher:
//         exact: '/b/c/4'
//       routeAction:
//         single:
//           upstream:
//             name: 'qux-upstream'
// ```
//
// * A RouteTable which is a child of another route table.
//
//
// Would produce the following route config for `mydomain.com`:
//
// ```
// /a/1 -> foo-upstream
// /a/2 -> bar-upstream
// /b/3 -> baz-upstream
// /b/c/4 -> qux-upstream
// ```
//
type RouteTable struct {
	// the list of routes for the route table
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteTable) Reset()         { *m = RouteTable{} }
func (m *RouteTable) String() string { return proto.CompactTextString(m) }
func (*RouteTable) ProtoMessage()    {}
func (*RouteTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_5624f75fbb628821, []int{0}
}
func (m *RouteTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTable.Unmarshal(m, b)
}
func (m *RouteTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTable.Marshal(b, m, deterministic)
}
func (m *RouteTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTable.Merge(m, src)
}
func (m *RouteTable) XXX_Size() int {
	return xxx_messageInfo_RouteTable.Size(m)
}
func (m *RouteTable) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTable.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTable proto.InternalMessageInfo

func (m *RouteTable) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RouteTable) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *RouteTable) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*RouteTable)(nil), "gateway.solo.io.RouteTable")
}

func init() {
	proto.RegisterFile("projects/gateway/api/v1/route_table.proto", fileDescriptor_5624f75fbb628821)
}

var fileDescriptor_5624f75fbb628821 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0xff, 0xfe, 0x36, 0xd5, 0x0c, 0x46, 0x63, 0x43, 0x08, 0x21, 0x0a, 0x84, 0x15, 0x2e,
	0x9c, 0x89, 0xb0, 0x31, 0x24, 0x6e, 0xd8, 0x1a, 0x37, 0xd5, 0x95, 0x1b, 0x32, 0x94, 0x61, 0x1c,
	0x29, 0xde, 0x66, 0xe6, 0x16, 0x71, 0xcb, 0xd3, 0xf8, 0x08, 0x3e, 0x82, 0x4f, 0x81, 0x89, 0x6f,
	0x80, 0x89, 0x7b, 0xd3, 0x76, 0x4a, 0x94, 0xc8, 0xae, 0xbd, 0xe7, 0xdc, 0x33, 0xf3, 0x9d, 0x21,
	0xa7, 0xb1, 0x86, 0x07, 0x11, 0xa2, 0x61, 0x92, 0xa3, 0x78, 0xe2, 0xcf, 0x8c, 0xc7, 0x8a, 0xcd,
	0xce, 0x99, 0x86, 0x04, 0xc5, 0x00, 0xf9, 0x30, 0x12, 0x34, 0xd6, 0x80, 0xe0, 0x1f, 0x5a, 0x07,
	0x35, 0x10, 0x01, 0x55, 0x50, 0x2b, 0x4b, 0x90, 0x90, 0x69, 0x2c, 0xfd, 0xca, 0x6d, 0x35, 0x5f,
	0xcc, 0x31, 0x1f, 0x8a, 0x39, 0xda, 0x59, 0x3d, 0x5d, 0x39, 0x9b, 0x28, 0x2c, 0xd2, 0xa7, 0x02,
	0xf9, 0x88, 0x23, 0xb7, 0xfa, 0xf1, 0xa6, 0x6e, 0x90, 0x63, 0x62, 0xb6, 0x6d, 0x17, 0xff, 0x56,
	0xef, 0xc8, 0x08, 0x80, 0x6d, 0x03, 0x99, 0x29, 0x8d, 0x09, 0x8f, 0x06, 0x46, 0xe8, 0x99, 0x0a,
	0x2d, 0x4c, 0xeb, 0xdd, 0x21, 0x24, 0x48, 0x11, 0x6f, 0x53, 0x42, 0x9f, 0x12, 0x2f, 0x03, 0x36,
	0x55, 0xa7, 0xb9, 0xd3, 0x2e, 0x75, 0x2a, 0x74, 0x03, 0x96, 0x66, 0xe6, 0xc0, 0xba, 0xfc, 0x2b,
	0xe2, 0xe5, 0x57, 0xac, 0x7a, 0x4d, 0xa7, 0x5d, 0xea, 0x94, 0x69, 0x08, 0x5a, 0xac, 0xcd, 0x37,
	0x99, 0xd6, 0x3f, 0x79, 0xfd, 0x72, 0x9d, 0xb7, 0x65, 0xe3, 0xdf, 0xe7, 0xb2, 0x71, 0x84, 0xc2,
	0xe0, 0x48, 0x8d, 0xc7, 0xbd, 0x96, 0x92, 0x8f, 0xa0, 0x45, 0x2b, 0xb0, 0x11, 0xfe, 0x05, 0xd9,
	0x2b, 0xfa, 0xa8, 0xee, 0x66, 0x71, 0x95, 0xdf, 0x71, 0xd7, 0x56, 0xed, 0xbb, 0x69, 0x58, 0xb0,
	0x76, 0xf7, 0x6a, 0x8b, 0x95, 0xeb, 0x92, 0xff, 0x1a, 0x17, 0x2b, 0xf7, 0xc0, 0xdf, 0xff, 0xf1,
	0x66, 0xa6, 0x7f, 0x99, 0x1e, 0xfe, 0xf2, 0x51, 0x77, 0xee, 0xba, 0x52, 0xe1, 0x7d, 0x32, 0xa4,
	0x21, 0x4c, 0xf3, 0xea, 0x14, 0xb0, 0xbf, 0x2b, 0x8b, 0x27, 0xd2, 0xd6, 0x36, 0xf4, 0xb2, 0x9e,
	0xba, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x8d, 0xdb, 0xa8, 0x21, 0x02, 0x00, 0x00,
}

func (this *RouteTable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTable)
	if !ok {
		that2, ok := that.(RouteTable)
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
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
