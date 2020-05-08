// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Code int32

const (
	Code_FATAL Code = 0
)

var Code_name = map[int32]string{
	0: "FATAL",
}

var Code_value = map[string]int32{
	"FATAL": 0,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type ContainerDriver int32

const (
	ContainerDriver_CONTAINERD ContainerDriver = 0
	ContainerDriver_CRI        ContainerDriver = 1
)

var ContainerDriver_name = map[int32]string{
	0: "CONTAINERD",
	1: "CRI",
}

var ContainerDriver_value = map[string]int32{
	"CONTAINERD": 0,
	"CRI":        1,
}

func (x ContainerDriver) String() string {
	return proto.EnumName(ContainerDriver_name, int32(x))
}

func (ContainerDriver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type Error struct {
	Code                 Code       `protobuf:"varint,1,opt,name=code,proto3,enum=common.Code" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}

func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}

func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}

func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}

func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_FATAL
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Common metadata message nested in all reply message types
type Metadata struct {
	// hostname of the server response comes from (injected by proxy)
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// error is set if request failed to the upstream (rest of response is
	// undefined)
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}

func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}

func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}

func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}

func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Metadata) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Data struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Bytes                []byte    `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}

func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}

func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}

func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}

func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Data) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type DataResponse struct {
	Messages             []*Data  `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}

func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}

func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}

func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}

func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetMessages() []*Data {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Empty struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}

func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}

func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}

func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}

func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type EmptyResponse struct {
	Messages             []*Empty `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}

func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}

func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}

func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}

func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func (m *EmptyResponse) GetMessages() []*Empty {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Code", Code_name, Code_value)
	proto.RegisterEnum("common.ContainerDriver", ContainerDriver_name, ContainerDriver_value)
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*Data)(nil), "common.Data")
	proto.RegisterType((*DataResponse)(nil), "common.DataResponse")
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*EmptyResponse)(nil), "common.EmptyResponse")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x6f, 0xa3, 0x30,
	0x14, 0xc4, 0xc3, 0x26, 0xe4, 0xe3, 0xe5, 0x63, 0x59, 0x6f, 0x0e, 0x6c, 0x4e, 0x88, 0x13, 0xc9,
	0xee, 0x82, 0x94, 0xd5, 0x4a, 0x55, 0xd5, 0x0b, 0x85, 0x54, 0x4a, 0xd5, 0xa6, 0x92, 0x95, 0x53,
	0x6f, 0x26, 0xb8, 0x04, 0x29, 0x60, 0x84, 0x9d, 0x4a, 0xfc, 0xf7, 0x15, 0x36, 0x70, 0xe9, 0xa9,
	0x27, 0x34, 0xbc, 0x99, 0xdf, 0x3c, 0x63, 0xe0, 0xe7, 0x89, 0x65, 0x19, 0xcb, 0x3d, 0xf5, 0x70,
	0x8b, 0x92, 0x09, 0x86, 0x86, 0x4a, 0xad, 0x7e, 0x25, 0x8c, 0x25, 0x17, 0xea, 0xc9, 0xb7, 0xd1,
	0xf5, 0xcd, 0x23, 0x79, 0xa5, 0x2c, 0x36, 0x07, 0x7d, 0x57, 0x96, 0xac, 0x44, 0x16, 0x0c, 0x4e,
	0x2c, 0xa6, 0xa6, 0x66, 0x69, 0xce, 0x62, 0x3b, 0x73, 0x1b, 0x50, 0xc0, 0x62, 0x8a, 0xe5, 0x04,
	0x99, 0x30, 0xca, 0x28, 0xe7, 0x24, 0xa1, 0xe6, 0x37, 0x4b, 0x73, 0x26, 0xb8, 0x95, 0xc8, 0x85,
	0x51, 0x4c, 0x05, 0x49, 0x2f, 0xdc, 0xec, 0x5b, 0x7d, 0x67, 0xba, 0x5d, 0xba, 0xaa, 0xd1, 0x6d,
	0x1b, 0x5d, 0x3f, 0xaf, 0x70, 0x6b, 0xb2, 0xef, 0x60, 0xfc, 0x4c, 0x05, 0x89, 0x89, 0x20, 0x68,
	0x05, 0xe3, 0x33, 0xe3, 0x22, 0x27, 0x99, 0xea, 0x9e, 0xe0, 0x4e, 0xa3, 0x25, 0xe8, 0xb4, 0x5e,
	0xae, 0xe9, 0x53, 0xc2, 0x7e, 0x84, 0x41, 0x58, 0x27, 0xff, 0xc0, 0x38, 0x6b, 0x28, 0x32, 0x39,
	0xdd, 0x1a, 0xed, 0xd6, 0x2d, 0x1d, 0x77, 0x8e, 0x9a, 0x15, 0x55, 0x82, 0x72, 0xc9, 0x9a, 0x61,
	0x25, 0xec, 0x1b, 0x98, 0xd5, 0x2c, 0x4c, 0x79, 0xc1, 0x72, 0x4e, 0x91, 0x53, 0x33, 0xe5, 0xa1,
	0xb8, 0xa9, 0xc9, 0xa3, 0x74, 0x5f, 0x22, 0x6c, 0x78, 0x6a, 0x6a, 0xff, 0x07, 0x7d, 0x97, 0x15,
	0xa2, 0xfa, 0xda, 0x1a, 0xf6, 0x2d, 0xcc, 0x65, 0xac, 0x6b, 0x5c, 0x7f, 0x6a, 0x9c, 0xb7, 0x71,
	0x65, 0xec, 0xc6, 0x9b, 0x1f, 0x30, 0xa8, 0xaf, 0x03, 0x4d, 0x40, 0x7f, 0xf0, 0x8f, 0xfe, 0x93,
	0xd1, 0xdb, 0x6c, 0xe0, 0x7b, 0xc0, 0x72, 0x41, 0xd2, 0x9c, 0x96, 0x61, 0x99, 0xbe, 0xd3, 0x12,
	0x2d, 0x00, 0x82, 0x97, 0xc3, 0xd1, 0xdf, 0x1f, 0x76, 0x38, 0x34, 0x7a, 0x68, 0x04, 0xfd, 0x00,
	0xef, 0x0d, 0xed, 0xfe, 0xf7, 0xeb, 0x3a, 0x49, 0xc5, 0xf9, 0x1a, 0xd5, 0x7c, 0x4f, 0x90, 0x0b,
	0xe3, 0x7f, 0x79, 0xc5, 0x05, 0xcd, 0xb8, 0x52, 0x1e, 0x29, 0xd2, 0xe6, 0x07, 0x8a, 0x86, 0xf2,
	0xe6, 0xfe, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0x4d, 0xd1, 0x68, 0x58, 0x02, 0x00, 0x00,
}
