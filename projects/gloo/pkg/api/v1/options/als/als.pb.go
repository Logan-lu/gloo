// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/als/als.proto

package als

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/ext"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Contains various settings for Envoy's access logging service.
// See here for more information: https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/accesslog/v2/accesslog.proto#envoy-api-msg-config-filter-accesslog-v2-accesslog
type AccessLoggingService struct {
	AccessLog            []*AccessLog `protobuf:"bytes,1,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccessLoggingService) Reset()         { *m = AccessLoggingService{} }
func (m *AccessLoggingService) String() string { return proto.CompactTextString(m) }
func (*AccessLoggingService) ProtoMessage()    {}
func (*AccessLoggingService) Descriptor() ([]byte, []int) {
	return fileDescriptor_510ef0fc4b9989af, []int{0}
}
func (m *AccessLoggingService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLoggingService.Unmarshal(m, b)
}
func (m *AccessLoggingService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLoggingService.Marshal(b, m, deterministic)
}
func (m *AccessLoggingService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLoggingService.Merge(m, src)
}
func (m *AccessLoggingService) XXX_Size() int {
	return xxx_messageInfo_AccessLoggingService.Size(m)
}
func (m *AccessLoggingService) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLoggingService.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLoggingService proto.InternalMessageInfo

func (m *AccessLoggingService) GetAccessLog() []*AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

type AccessLog struct {
	// type of Access Logging service to implement
	//
	// Types that are valid to be assigned to OutputDestination:
	//	*AccessLog_FileSink
	//	*AccessLog_GrpcService
	OutputDestination    isAccessLog_OutputDestination `protobuf_oneof:"OutputDestination"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_510ef0fc4b9989af, []int{1}
}
func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

type isAccessLog_OutputDestination interface {
	isAccessLog_OutputDestination()
	Equal(interface{}) bool
}

type AccessLog_FileSink struct {
	FileSink *FileSink `protobuf:"bytes,2,opt,name=file_sink,json=fileSink,proto3,oneof" json:"file_sink,omitempty"`
}
type AccessLog_GrpcService struct {
	GrpcService *GrpcService `protobuf:"bytes,3,opt,name=grpc_service,json=grpcService,proto3,oneof" json:"grpc_service,omitempty"`
}

func (*AccessLog_FileSink) isAccessLog_OutputDestination()    {}
func (*AccessLog_GrpcService) isAccessLog_OutputDestination() {}

func (m *AccessLog) GetOutputDestination() isAccessLog_OutputDestination {
	if m != nil {
		return m.OutputDestination
	}
	return nil
}

func (m *AccessLog) GetFileSink() *FileSink {
	if x, ok := m.GetOutputDestination().(*AccessLog_FileSink); ok {
		return x.FileSink
	}
	return nil
}

func (m *AccessLog) GetGrpcService() *GrpcService {
	if x, ok := m.GetOutputDestination().(*AccessLog_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLog_FileSink)(nil),
		(*AccessLog_GrpcService)(nil),
	}
}

type FileSink struct {
	// the file path to which the file access logging service will sink
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// the format which the logs should be outputted by
	//
	// Types that are valid to be assigned to OutputFormat:
	//	*FileSink_StringFormat
	//	*FileSink_JsonFormat
	OutputFormat         isFileSink_OutputFormat `protobuf_oneof:"output_format"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FileSink) Reset()         { *m = FileSink{} }
func (m *FileSink) String() string { return proto.CompactTextString(m) }
func (*FileSink) ProtoMessage()    {}
func (*FileSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_510ef0fc4b9989af, []int{2}
}
func (m *FileSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSink.Unmarshal(m, b)
}
func (m *FileSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSink.Marshal(b, m, deterministic)
}
func (m *FileSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSink.Merge(m, src)
}
func (m *FileSink) XXX_Size() int {
	return xxx_messageInfo_FileSink.Size(m)
}
func (m *FileSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSink.DiscardUnknown(m)
}

var xxx_messageInfo_FileSink proto.InternalMessageInfo

type isFileSink_OutputFormat interface {
	isFileSink_OutputFormat()
	Equal(interface{}) bool
}

type FileSink_StringFormat struct {
	StringFormat string `protobuf:"bytes,2,opt,name=string_format,json=stringFormat,proto3,oneof" json:"string_format,omitempty"`
}
type FileSink_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof" json:"json_format,omitempty"`
}

func (*FileSink_StringFormat) isFileSink_OutputFormat() {}
func (*FileSink_JsonFormat) isFileSink_OutputFormat()   {}

func (m *FileSink) GetOutputFormat() isFileSink_OutputFormat {
	if m != nil {
		return m.OutputFormat
	}
	return nil
}

func (m *FileSink) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileSink) GetStringFormat() string {
	if x, ok := m.GetOutputFormat().(*FileSink_StringFormat); ok {
		return x.StringFormat
	}
	return ""
}

func (m *FileSink) GetJsonFormat() *types.Struct {
	if x, ok := m.GetOutputFormat().(*FileSink_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileSink_StringFormat)(nil),
		(*FileSink_JsonFormat)(nil),
	}
}

type GrpcService struct {
	// name of log stream
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The static cluster defined in bootstrap config to route to
	//
	// Types that are valid to be assigned to ServiceRef:
	//	*GrpcService_StaticClusterName
	ServiceRef                      isGrpcService_ServiceRef `protobuf_oneof:"service_ref"`
	AdditionalRequestHeadersToLog   []string                 `protobuf:"bytes,4,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                 `protobuf:"bytes,5,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                 `protobuf:"bytes,6,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                 `json:"-"`
	XXX_unrecognized                []byte                   `json:"-"`
	XXX_sizecache                   int32                    `json:"-"`
}

func (m *GrpcService) Reset()         { *m = GrpcService{} }
func (m *GrpcService) String() string { return proto.CompactTextString(m) }
func (*GrpcService) ProtoMessage()    {}
func (*GrpcService) Descriptor() ([]byte, []int) {
	return fileDescriptor_510ef0fc4b9989af, []int{3}
}
func (m *GrpcService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcService.Unmarshal(m, b)
}
func (m *GrpcService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcService.Marshal(b, m, deterministic)
}
func (m *GrpcService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcService.Merge(m, src)
}
func (m *GrpcService) XXX_Size() int {
	return xxx_messageInfo_GrpcService.Size(m)
}
func (m *GrpcService) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcService.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcService proto.InternalMessageInfo

type isGrpcService_ServiceRef interface {
	isGrpcService_ServiceRef()
	Equal(interface{}) bool
}

type GrpcService_StaticClusterName struct {
	StaticClusterName string `protobuf:"bytes,2,opt,name=static_cluster_name,json=staticClusterName,proto3,oneof" json:"static_cluster_name,omitempty"`
}

func (*GrpcService_StaticClusterName) isGrpcService_ServiceRef() {}

func (m *GrpcService) GetServiceRef() isGrpcService_ServiceRef {
	if m != nil {
		return m.ServiceRef
	}
	return nil
}

func (m *GrpcService) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *GrpcService) GetStaticClusterName() string {
	if x, ok := m.GetServiceRef().(*GrpcService_StaticClusterName); ok {
		return x.StaticClusterName
	}
	return ""
}

func (m *GrpcService) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcService_StaticClusterName)(nil),
	}
}

func init() {
	proto.RegisterType((*AccessLoggingService)(nil), "als.options.gloo.solo.io.AccessLoggingService")
	proto.RegisterType((*AccessLog)(nil), "als.options.gloo.solo.io.AccessLog")
	proto.RegisterType((*FileSink)(nil), "als.options.gloo.solo.io.FileSink")
	proto.RegisterType((*GrpcService)(nil), "als.options.gloo.solo.io.GrpcService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/als/als.proto", fileDescriptor_510ef0fc4b9989af)
}

var fileDescriptor_510ef0fc4b9989af = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x9b, 0x75, 0xdf, 0xbe, 0xc6, 0x59, 0x85, 0x96, 0x4d, 0xa2, 0x4c, 0x30, 0x4a, 0xa6,
	0x49, 0xbd, 0x80, 0x04, 0xc6, 0x1d, 0xe2, 0x66, 0x01, 0x46, 0x34, 0x4d, 0x20, 0xa5, 0xbb, 0xda,
	0x4d, 0xe4, 0xa6, 0x8e, 0xeb, 0xd5, 0xcd, 0x09, 0xb6, 0x33, 0xed, 0x39, 0x78, 0x0a, 0xee, 0x78,
	0x12, 0x5e, 0x82, 0xa7, 0xe0, 0x12, 0xc5, 0x76, 0xb6, 0x0e, 0x5a, 0x89, 0x8b, 0xaa, 0xe7, 0xf8,
	0xfc, 0xcf, 0xcf, 0xff, 0xa3, 0x13, 0xa3, 0x98, 0x32, 0x35, 0xab, 0x27, 0x61, 0x0e, 0x8b, 0x48,
	0x02, 0x87, 0x17, 0x0c, 0x22, 0xca, 0x01, 0xa2, 0x4a, 0xc0, 0x15, 0xc9, 0x95, 0x34, 0x19, 0xae,
	0x58, 0x74, 0xfd, 0x2a, 0x82, 0x4a, 0x31, 0x28, 0x65, 0x84, 0xb9, 0xfe, 0x85, 0x95, 0x00, 0x05,
	0xfe, 0xa0, 0x09, 0x6d, 0x29, 0x6c, 0xe4, 0x61, 0x43, 0x0a, 0x19, 0xec, 0xef, 0x51, 0xa0, 0xa0,
	0x45, 0x51, 0x13, 0x19, 0xfd, 0x7e, 0x9f, 0xdc, 0xa8, 0x88, 0xdc, 0x28, 0x9b, 0x3e, 0x5f, 0x61,
	0x41, 0xff, 0xcf, 0x99, 0x6a, 0x2f, 0x16, 0xa4, 0xb0, 0xea, 0xc7, 0x14, 0x80, 0x72, 0x12, 0xe9,
	0x6c, 0x52, 0x17, 0x91, 0x54, 0xa2, 0xce, 0x2d, 0x2b, 0xb8, 0x44, 0x7b, 0x27, 0x79, 0x4e, 0xa4,
	0x3c, 0x07, 0x4a, 0x59, 0x49, 0xc7, 0x44, 0x5c, 0xb3, 0x9c, 0xf8, 0x31, 0x42, 0x58, 0x9f, 0x67,
	0x1c, 0xe8, 0xc0, 0x19, 0x76, 0x47, 0xde, 0xf1, 0x61, 0xb8, 0xce, 0x77, 0x78, 0xcb, 0x48, 0x5d,
	0xdc, 0x86, 0xc1, 0x77, 0x07, 0xb9, 0xb7, 0x05, 0xff, 0x04, 0xb9, 0x05, 0xe3, 0x24, 0x93, 0xac,
	0x9c, 0x0f, 0x36, 0x86, 0xce, 0xc8, 0x3b, 0x0e, 0xd6, 0x03, 0x4f, 0x19, 0x27, 0x63, 0x56, 0xce,
	0x93, 0x4e, 0xda, 0x2b, 0x6c, 0xec, 0x9f, 0xa1, 0x6d, 0x2a, 0xaa, 0x3c, 0x93, 0xc6, 0xe4, 0xa0,
	0xab, 0x29, 0x47, 0xeb, 0x29, 0x1f, 0x45, 0x95, 0xdb, 0x89, 0x92, 0x4e, 0xea, 0xd1, 0xbb, 0x34,
	0xde, 0x45, 0x3b, 0x9f, 0x6b, 0x55, 0xd5, 0xea, 0x3d, 0x91, 0x8a, 0x95, 0xb8, 0xe9, 0x0e, 0xbe,
	0x3a, 0xa8, 0xd7, 0xde, 0xec, 0xfb, 0x68, 0xb3, 0xc2, 0x6a, 0x36, 0x70, 0x86, 0xce, 0xc8, 0x4d,
	0x75, 0xec, 0x1f, 0xa1, 0xbe, 0x54, 0x82, 0x95, 0x34, 0x2b, 0x40, 0x2c, 0xb0, 0xd2, 0x83, 0xb8,
	0x49, 0x27, 0xdd, 0x36, 0xc7, 0xa7, 0xfa, 0xd4, 0x7f, 0x83, 0xbc, 0x2b, 0x09, 0x65, 0x2b, 0x32,
	0x3e, 0x1f, 0x86, 0x66, 0x13, 0x61, 0xbb, 0x89, 0x70, 0xac, 0x37, 0x91, 0x74, 0x52, 0xd4, 0xa8,
	0x4d, 0x6f, 0xfc, 0x00, 0xf5, 0x41, 0x1b, 0xb3, 0xdd, 0xc1, 0x8f, 0x0d, 0xe4, 0x2d, 0x0d, 0xe2,
	0x3f, 0x42, 0x3d, 0x0e, 0x34, 0x2b, 0xf1, 0x82, 0x58, 0x6f, 0xff, 0x73, 0xa0, 0x9f, 0xf0, 0x82,
	0xf8, 0x2f, 0xd1, 0xae, 0x54, 0x58, 0xb1, 0x3c, 0xcb, 0x79, 0x2d, 0x15, 0x11, 0x46, 0xd5, 0x9a,
	0xdc, 0x31, 0xc5, 0x77, 0xa6, 0xa6, 0x3b, 0x12, 0xf4, 0x0c, 0x4f, 0xa7, 0xac, 0x99, 0x1e, 0xf3,
	0x4c, 0x90, 0x2f, 0x35, 0x91, 0x2a, 0x9b, 0x11, 0x3c, 0x25, 0x42, 0x66, 0x0a, 0xf4, 0xfa, 0x37,
	0x87, 0xdd, 0x91, 0x9b, 0x3e, 0xb9, 0x13, 0xa6, 0x46, 0x97, 0x18, 0xd9, 0x05, 0x34, 0xfb, 0x3d,
	0x43, 0xc1, 0x3d, 0x92, 0xac, 0xa0, 0x94, 0xe4, 0x4f, 0xd4, 0x7f, 0x1a, 0x75, 0xb0, 0x8c, 0x32,
	0xc2, 0x7b, 0xac, 0x73, 0x74, 0xb8, 0x8a, 0xa5, 0x04, 0x66, 0x7c, 0x09, 0xb6, 0xa5, 0x61, 0x4f,
	0xff, 0x86, 0x5d, 0x58, 0xa1, 0xa6, 0xc5, 0x7d, 0xe4, 0xd9, 0x2f, 0x26, 0x13, 0xa4, 0x88, 0x3f,
	0xfc, 0x8a, 0x9d, 0x6f, 0x3f, 0x0f, 0x9c, 0xcb, 0xb7, 0xff, 0xf6, 0x94, 0xab, 0x39, 0x5d, 0xf1,
	0x9c, 0x27, 0x5b, 0x7a, 0x8d, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x92, 0xbd, 0xdd, 0x37,
	0x11, 0x04, 0x00, 0x00,
}

func (this *AccessLoggingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLoggingService)
	if !ok {
		that2, ok := that.(AccessLoggingService)
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
	if len(this.AccessLog) != len(that1.AccessLog) {
		return false
	}
	for i := range this.AccessLog {
		if !this.AccessLog[i].Equal(that1.AccessLog[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog)
	if !ok {
		that2, ok := that.(AccessLog)
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
	if that1.OutputDestination == nil {
		if this.OutputDestination != nil {
			return false
		}
	} else if this.OutputDestination == nil {
		return false
	} else if !this.OutputDestination.Equal(that1.OutputDestination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog_FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_FileSink)
	if !ok {
		that2, ok := that.(AccessLog_FileSink)
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
	if !this.FileSink.Equal(that1.FileSink) {
		return false
	}
	return true
}
func (this *AccessLog_GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_GrpcService)
	if !ok {
		that2, ok := that.(AccessLog_GrpcService)
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
	if !this.GrpcService.Equal(that1.GrpcService) {
		return false
	}
	return true
}
func (this *FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink)
	if !ok {
		that2, ok := that.(FileSink)
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
	if that1.OutputFormat == nil {
		if this.OutputFormat != nil {
			return false
		}
	} else if this.OutputFormat == nil {
		return false
	} else if !this.OutputFormat.Equal(that1.OutputFormat) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FileSink_StringFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_StringFormat)
	if !ok {
		that2, ok := that.(FileSink_StringFormat)
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
	if this.StringFormat != that1.StringFormat {
		return false
	}
	return true
}
func (this *FileSink_JsonFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_JsonFormat)
	if !ok {
		that2, ok := that.(FileSink_JsonFormat)
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
	if !this.JsonFormat.Equal(that1.JsonFormat) {
		return false
	}
	return true
}
func (this *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
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
	if this.LogName != that1.LogName {
		return false
	}
	if that1.ServiceRef == nil {
		if this.ServiceRef != nil {
			return false
		}
	} else if this.ServiceRef == nil {
		return false
	} else if !this.ServiceRef.Equal(that1.ServiceRef) {
		return false
	}
	if len(this.AdditionalRequestHeadersToLog) != len(that1.AdditionalRequestHeadersToLog) {
		return false
	}
	for i := range this.AdditionalRequestHeadersToLog {
		if this.AdditionalRequestHeadersToLog[i] != that1.AdditionalRequestHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseHeadersToLog) != len(that1.AdditionalResponseHeadersToLog) {
		return false
	}
	for i := range this.AdditionalResponseHeadersToLog {
		if this.AdditionalResponseHeadersToLog[i] != that1.AdditionalResponseHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseTrailersToLog) != len(that1.AdditionalResponseTrailersToLog) {
		return false
	}
	for i := range this.AdditionalResponseTrailersToLog {
		if this.AdditionalResponseTrailersToLog[i] != that1.AdditionalResponseTrailersToLog[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GrpcService_StaticClusterName) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService_StaticClusterName)
	if !ok {
		that2, ok := that.(GrpcService_StaticClusterName)
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
	if this.StaticClusterName != that1.StaticClusterName {
		return false
	}
	return true
}
