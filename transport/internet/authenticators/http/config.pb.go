// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/authenticators/http/config.proto
// DO NOT EDIT!

/*
Package http is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/authenticators/http/config.proto

It has these top-level messages:
	Header
	Version
	Method
	RequestConfig
	Status
	ResponseConfig
	Config
*/
package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Header struct {
	// "Accept", "Cookie", etc
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Each entry must be valid in one piece. Random entry will be chosen if multiple entries present.
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HTTP version. Default value "1.1".
type Version struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// HTTP method. Default value "GET".
type Method struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RequestConfig struct {
	// Full HTTP version like "1.1".
	Version *Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// GET, POST, CONNECT etc
	Method *Method `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	// URI like "/login.php"
	Uri    []string  `protobuf:"bytes,3,rep,name=uri" json:"uri,omitempty"`
	Header []*Header `protobuf:"bytes,4,rep,name=header" json:"header,omitempty"`
}

func (m *RequestConfig) Reset()                    { *m = RequestConfig{} }
func (m *RequestConfig) String() string            { return proto.CompactTextString(m) }
func (*RequestConfig) ProtoMessage()               {}
func (*RequestConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestConfig) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *RequestConfig) GetMethod() *Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RequestConfig) GetHeader() []*Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type Status struct {
	// Status code. Default "200".
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// Statue reason. Default "OK".
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ResponseConfig struct {
	Version *Version  `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Status  *Status   `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Header  []*Header `protobuf:"bytes,3,rep,name=header" json:"header,omitempty"`
}

func (m *ResponseConfig) Reset()                    { *m = ResponseConfig{} }
func (m *ResponseConfig) String() string            { return proto.CompactTextString(m) }
func (*ResponseConfig) ProtoMessage()               {}
func (*ResponseConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResponseConfig) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ResponseConfig) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ResponseConfig) GetHeader() []*Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type Config struct {
	Request  *RequestConfig  `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Response *ResponseConfig `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Config) GetRequest() *RequestConfig {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Config) GetResponse() *ResponseConfig {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "v2ray.core.transport.internet.authenticators.http.Header")
	proto.RegisterType((*Version)(nil), "v2ray.core.transport.internet.authenticators.http.Version")
	proto.RegisterType((*Method)(nil), "v2ray.core.transport.internet.authenticators.http.Method")
	proto.RegisterType((*RequestConfig)(nil), "v2ray.core.transport.internet.authenticators.http.RequestConfig")
	proto.RegisterType((*Status)(nil), "v2ray.core.transport.internet.authenticators.http.Status")
	proto.RegisterType((*ResponseConfig)(nil), "v2ray.core.transport.internet.authenticators.http.ResponseConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.authenticators.http.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/authenticators/http/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x40, 0xd5, 0xa4, 0x9b, 0x6e, 0xa7, 0xda, 0xd5, 0xca, 0x5a, 0xad, 0x72, 0xda, 0xad, 0x72,
	0xea, 0xc9, 0xd1, 0x76, 0x97, 0x03, 0x9c, 0xa0, 0x5c, 0x10, 0x12, 0x12, 0x18, 0xc4, 0xa1, 0x12,
	0x07, 0x93, 0x0e, 0x24, 0x12, 0xb1, 0x83, 0xed, 0x54, 0xe2, 0x1f, 0xf8, 0x05, 0xbe, 0x86, 0x1f,
	0x43, 0xb1, 0x9d, 0xd2, 0x1e, 0x38, 0x34, 0xc0, 0xcd, 0x4e, 0x66, 0x9e, 0x67, 0xde, 0x58, 0x86,
	0xd9, 0x72, 0xaa, 0xf8, 0x03, 0xcd, 0x64, 0x99, 0x66, 0x52, 0x61, 0x6a, 0x14, 0x17, 0xba, 0x92,
	0xca, 0xa4, 0x85, 0x30, 0xa8, 0x04, 0x9a, 0x94, 0xd7, 0x26, 0x47, 0x61, 0x8a, 0x8c, 0x1b, 0xa9,
	0x74, 0x9a, 0x1b, 0x53, 0xa5, 0x99, 0x14, 0x37, 0xc5, 0x2d, 0xad, 0x94, 0x34, 0x92, 0xfc, 0x6d,
	0x19, 0x0a, 0xe9, 0x2a, 0x9f, 0xb6, 0xf9, 0x74, 0x33, 0x9f, 0x36, 0xf9, 0xc9, 0x14, 0xa2, 0x23,
	0xe4, 0x0b, 0x54, 0x84, 0x40, 0x5f, 0xf0, 0x12, 0xe3, 0xde, 0xb8, 0x37, 0x19, 0x32, 0xbb, 0x26,
	0x3f, 0xe1, 0xcb, 0x92, 0xdf, 0xd5, 0x18, 0x07, 0xe3, 0x70, 0x32, 0x64, 0x6e, 0x93, 0xfc, 0x81,
	0xc1, 0x25, 0x2a, 0x5d, 0x48, 0xf1, 0x1a, 0xe0, 0xb2, 0x7c, 0xc0, 0x6f, 0x88, 0x4e, 0xd0, 0xe4,
	0x72, 0xf1, 0xc6, 0xff, 0xa7, 0x00, 0xbe, 0x31, 0xbc, 0xaf, 0x51, 0x9b, 0x43, 0x5b, 0x3f, 0xb9,
	0x80, 0xc1, 0xd2, 0x21, 0x6d, 0xe4, 0x68, 0xba, 0x47, 0xb7, 0xee, 0x85, 0xfa, 0xa2, 0x58, 0x8b,
	0x22, 0x67, 0x10, 0x95, 0xb6, 0x8e, 0x38, 0xb0, 0xd0, 0xdd, 0x0e, 0x50, 0xd7, 0x08, 0xf3, 0x20,
	0xf2, 0x03, 0xc2, 0x5a, 0x15, 0x71, 0x68, 0x7d, 0x34, 0xcb, 0xe6, 0x90, 0xdc, 0x1a, 0x8c, 0xfb,
	0xe3, 0xb0, 0xe3, 0x21, 0x6e, 0x04, 0xcc, 0x83, 0x92, 0xff, 0x10, 0x9d, 0x1b, 0x6e, 0x6a, 0xdd,
	0x0c, 0x25, 0x93, 0x8b, 0xd5, 0x50, 0x9a, 0x35, 0xf9, 0x05, 0x91, 0x42, 0xae, 0xa5, 0xb0, 0x5d,
	0x0d, 0x99, 0xdf, 0x25, 0x8f, 0x01, 0x7c, 0x67, 0xa8, 0x2b, 0x29, 0x34, 0x7e, 0xb6, 0x56, 0x6d,
	0xcb, 0x7b, 0x87, 0x56, 0xd7, 0x1f, 0xf3, 0xa0, 0x35, 0x89, 0xe1, 0x47, 0x49, 0x7c, 0xee, 0x41,
	0xe4, 0x35, 0xcc, 0x61, 0xa0, 0xdc, 0x75, 0xf3, 0x1a, 0xf6, 0x3b, 0xe0, 0x37, 0x2e, 0x2c, 0x6b,
	0x81, 0xe4, 0x0a, 0xbe, 0x2a, 0x2f, 0xdd, 0xeb, 0x38, 0xe8, 0x04, 0x5f, 0x9f, 0x1b, 0x5b, 0x21,
	0x67, 0xc7, 0xb0, 0x93, 0xc9, 0x72, 0x7b, 0xe2, 0x6c, 0xe4, 0x50, 0xa7, 0xcd, 0xc3, 0x30, 0xef,
	0x37, 0x9f, 0xae, 0x23, 0xfb, 0x4a, 0xfc, 0x7b, 0x09, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x1d, 0x02,
	0x42, 0x6b, 0x04, 0x00, 0x00,
}