// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/ext"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
	return fileDescriptor_4d1ea5a66e7f9a13, []int{0}
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
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto", fileDescriptor_4d1ea5a66e7f9a13)
}

var fileDescriptor_4d1ea5a66e7f9a13 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4e, 0x02, 0x41,
	0x10, 0x87, 0x3d, 0xbd, 0x9c, 0xe6, 0xf0, 0x4f, 0xbc, 0x10, 0x82, 0x14, 0x42, 0xa8, 0x68, 0xdc,
	0x0d, 0xd0, 0x18, 0x62, 0xe3, 0x25, 0xc6, 0xca, 0xe6, 0xb4, 0xb2, 0x21, 0xcb, 0xb1, 0xac, 0x2b,
	0x87, 0x73, 0xd9, 0x9d, 0x43, 0x6c, 0x79, 0x1a, 0x1f, 0xc5, 0x77, 0x30, 0xa1, 0xf0, 0x0d, 0xb0,
	0xb2, 0x34, 0xb7, 0x2c, 0x44, 0x8d, 0x05, 0x54, 0xbb, 0x93, 0xf9, 0x7e, 0x33, 0xd9, 0x6f, 0xfd,
	0x4b, 0x21, 0xf1, 0x21, 0xeb, 0x91, 0x18, 0x46, 0x54, 0x43, 0x02, 0x67, 0x12, 0xa8, 0x48, 0x00,
	0x68, 0xaa, 0xe0, 0x91, 0xc7, 0xa8, 0xa9, 0x60, 0xc8, 0x9f, 0xd9, 0x0b, 0x65, 0xa9, 0xa4, 0xe3,
	0x26, 0x55, 0x90, 0x21, 0xef, 0x22, 0xeb, 0x25, 0x9c, 0xa4, 0x0a, 0x10, 0x82, 0x23, 0x4b, 0x90,
	0x3c, 0x4f, 0x24, 0x54, 0x8a, 0x02, 0x04, 0x98, 0x1e, 0xcd, 0x6f, 0x0b, 0xac, 0x72, 0xc0, 0x27,
	0x48, 0xf9, 0x04, 0x6d, 0xd9, 0xfc, 0x67, 0xb1, 0x39, 0x87, 0x12, 0x97, 0xbb, 0x46, 0x1c, 0x59,
	0x9f, 0x21, 0xb3, 0x11, 0xba, 0x46, 0x44, 0x23, 0xc3, 0x4c, 0x6f, 0xb0, 0x63, 0x59, 0xdb, 0xc8,
	0xd5, 0xa6, 0x3e, 0xc6, 0x52, 0x61, 0xc6, 0x92, 0xae, 0xe6, 0x6a, 0x2c, 0x63, 0xeb, 0xa4, 0xfe,
	0xee, 0xf8, 0x7e, 0x94, 0x9b, 0xba, 0xcb, 0x45, 0x05, 0xc4, 0xf7, 0x8c, 0x37, 0x5d, 0x76, 0x6a,
	0x3b, 0x8d, 0x42, 0xab, 0x44, 0xfe, 0x38, 0x23, 0x06, 0x8e, 0x2c, 0x15, 0x5c, 0xfb, 0xde, 0xe2,
	0x21, 0x65, 0xaf, 0xe6, 0x34, 0x0a, 0xad, 0x22, 0x89, 0x41, 0xf1, 0x15, 0x7c, 0x6b, 0x7a, 0xe1,
	0xc9, 0xdb, 0xac, 0xba, 0xf5, 0x39, 0xab, 0x1e, 0x23, 0xd7, 0xd8, 0x97, 0x83, 0x41, 0xa7, 0x2e,
	0xc5, 0x13, 0x28, 0x5e, 0x8f, 0x6c, 0x3c, 0x38, 0xf7, 0xf7, 0x96, 0x12, 0xcb, 0xbb, 0x66, 0x54,
	0xe9, 0xf7, 0xa8, 0x1b, 0xdb, 0x0d, 0xdd, 0x7c, 0x58, 0xb4, 0xa2, 0x3b, 0x95, 0xe9, 0xdc, 0x75,
	0xfd, 0x6d, 0x85, 0xd3, 0xb9, 0x7b, 0x18, 0xec, 0xff, 0xf8, 0x76, 0x1d, 0x5e, 0x7c, 0x85, 0xce,
	0xeb, 0xc7, 0xa9, 0x73, 0xdf, 0x5e, 0xdb, 0x56, 0x3a, 0x14, 0xd6, 0x58, 0xcf, 0x33, 0x8a, 0xda,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x59, 0x79, 0x97, 0x7b, 0x02, 0x00, 0x00,
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
