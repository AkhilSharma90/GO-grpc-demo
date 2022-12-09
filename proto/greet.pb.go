// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greet.proto

// package name for our proto file

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// when you are not passing any params still you will need to add a type
// NoParam acts as a type for no parameters given.
type NoParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoParam) Reset()         { *m = NoParam{} }
func (m *NoParam) String() string { return proto.CompactTextString(m) }
func (*NoParam) ProtoMessage()    {}
func (*NoParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{0}
}

func (m *NoParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoParam.Unmarshal(m, b)
}
func (m *NoParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoParam.Marshal(b, m, deterministic)
}
func (m *NoParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoParam.Merge(m, src)
}
func (m *NoParam) XXX_Size() int {
	return xxx_messageInfo_NoParam.Size(m)
}
func (m *NoParam) XXX_DiscardUnknown() {
	xxx_messageInfo_NoParam.DiscardUnknown(m)
}

var xxx_messageInfo_NoParam proto.InternalMessageInfo

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NamesList struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamesList) Reset()         { *m = NamesList{} }
func (m *NamesList) String() string { return proto.CompactTextString(m) }
func (*NamesList) ProtoMessage()    {}
func (*NamesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{3}
}

func (m *NamesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamesList.Unmarshal(m, b)
}
func (m *NamesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamesList.Marshal(b, m, deterministic)
}
func (m *NamesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamesList.Merge(m, src)
}
func (m *NamesList) XXX_Size() int {
	return xxx_messageInfo_NamesList.Size(m)
}
func (m *NamesList) XXX_DiscardUnknown() {
	xxx_messageInfo_NamesList.DiscardUnknown(m)
}

var xxx_messageInfo_NamesList proto.InternalMessageInfo

func (m *NamesList) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type MessagesList struct {
	Messages             []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagesList) Reset()         { *m = MessagesList{} }
func (m *MessagesList) String() string { return proto.CompactTextString(m) }
func (*MessagesList) ProtoMessage()    {}
func (*MessagesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{4}
}

func (m *MessagesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagesList.Unmarshal(m, b)
}
func (m *MessagesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagesList.Marshal(b, m, deterministic)
}
func (m *MessagesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagesList.Merge(m, src)
}
func (m *MessagesList) XXX_Size() int {
	return xxx_messageInfo_MessagesList.Size(m)
}
func (m *MessagesList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagesList.DiscardUnknown(m)
}

var xxx_messageInfo_MessagesList proto.InternalMessageInfo

func (m *MessagesList) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*NoParam)(nil), "greet_service.NoParam")
	proto.RegisterType((*HelloRequest)(nil), "greet_service.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "greet_service.HelloResponse")
	proto.RegisterType((*NamesList)(nil), "greet_service.NamesList")
	proto.RegisterType((*MessagesList)(nil), "greet_service.MessagesList")
}

func init() { proto.RegisterFile("proto/greet.proto", fileDescriptor_95ca2ad3f55a58e3) }

var fileDescriptor_95ca2ad3f55a58e3 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x89, 0xff, 0xba, 0x1d, 0xba, 0x07, 0x83, 0x68, 0xa9, 0x22, 0x6b, 0x4e, 0xd5, 0x43,
	0x77, 0xd1, 0x17, 0x90, 0xf5, 0xa0, 0x07, 0x5d, 0xa4, 0xc5, 0x8b, 0x97, 0x25, 0xae, 0x43, 0x09,
	0xb4, 0xc9, 0x9a, 0x44, 0xc1, 0x87, 0xf5, 0x5d, 0xa4, 0x69, 0xe3, 0xba, 0x45, 0x16, 0x4f, 0x99,
	0x2f, 0xf9, 0xe6, 0xc7, 0x37, 0x43, 0x60, 0x7f, 0xa9, 0x95, 0x55, 0xe3, 0x52, 0x23, 0xda, 0xcc,
	0xd5, 0x74, 0xe8, 0xc4, 0xdc, 0xa0, 0xfe, 0x10, 0x0b, 0x64, 0x21, 0x04, 0x33, 0xf5, 0xc8, 0x35,
	0xaf, 0x19, 0x83, 0xe8, 0x0e, 0xab, 0x4a, 0xe5, 0xf8, 0xf6, 0x8e, 0xc6, 0x52, 0x0a, 0x3b, 0x92,
	0xd7, 0x18, 0x93, 0x11, 0x49, 0xc3, 0xdc, 0xd5, 0xec, 0x1c, 0x86, 0x9d, 0xc7, 0x2c, 0x95, 0x34,
	0x48, 0x63, 0x08, 0x6a, 0x34, 0x86, 0x97, 0xde, 0xe7, 0x25, 0x3b, 0x83, 0x70, 0xc6, 0x6b, 0x34,
	0xf7, 0xc2, 0x58, 0x7a, 0x00, 0xbb, 0x4d, 0xbf, 0x89, 0xc9, 0x68, 0x3b, 0x0d, 0xf3, 0x56, 0xb0,
	0x0b, 0x88, 0x1e, 0x5a, 0x77, 0xeb, 0x4a, 0x60, 0xd0, 0x75, 0x7b, 0xe3, 0x8f, 0xbe, 0xfc, 0xda,
	0x82, 0xe8, 0xb6, 0x89, 0x5e, 0xb4, 0xc9, 0xe9, 0x35, 0x0c, 0x0a, 0xfe, 0xe9, 0xd2, 0xd0, 0xc3,
	0x6c, 0x6d, 0xaa, 0xac, 0x1b, 0x29, 0x39, 0xe9, 0xdd, 0xaf, 0x67, 0x2f, 0xe0, 0xc8, 0x13, 0x1a,
	0x28, 0xea, 0xc2, 0x6a, 0xe4, 0xb5, 0x90, 0x25, 0x8d, 0xfb, 0x40, 0x3f, 0xc9, 0x66, 0xe4, 0x84,
	0xd0, 0xa7, 0x15, 0xf4, 0xa6, 0x12, 0x28, 0xed, 0x0a, 0x7a, 0xfc, 0x77, 0xab, 0xdb, 0x76, 0xd2,
	0x7f, 0xfc, 0xbd, 0x98, 0x94, 0xd0, 0x39, 0x9c, 0x7a, 0xec, 0x54, 0xbc, 0x0a, 0x8d, 0x0b, 0x2b,
	0x94, 0xe4, 0xd5, 0x3f, 0xe9, 0x1b, 0x53, 0xa7, 0x64, 0x42, 0xa6, 0xe1, 0x73, 0x90, 0x8d, 0xdd,
	0x17, 0x79, 0xd9, 0x73, 0xc7, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x21, 0x14, 0x38,
	0x3e, 0x02, 0x00, 0x00,
}
