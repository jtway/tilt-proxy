// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: tilt.proto

package tilt_proxy

import (
	_ "github.com/jtway/tilt-proxy/google/api"
	_ "github.com/jtway/tilt-proxy/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TiltRequest_Color int32

const (
	TiltRequest_Red    TiltRequest_Color = 0
	TiltRequest_Green  TiltRequest_Color = 1
	TiltRequest_Black  TiltRequest_Color = 2
	TiltRequest_Purple TiltRequest_Color = 3
	TiltRequest_Orange TiltRequest_Color = 4
	TiltRequest_Blue   TiltRequest_Color = 5
	TiltRequest_Yellow TiltRequest_Color = 6
	TiltRequest_Pink   TiltRequest_Color = 7
)

// Enum value maps for TiltRequest_Color.
var (
	TiltRequest_Color_name = map[int32]string{
		0: "Red",
		1: "Green",
		2: "Black",
		3: "Purple",
		4: "Orange",
		5: "Blue",
		6: "Yellow",
		7: "Pink",
	}
	TiltRequest_Color_value = map[string]int32{
		"Red":    0,
		"Green":  1,
		"Black":  2,
		"Purple": 3,
		"Orange": 4,
		"Blue":   5,
		"Yellow": 6,
		"Pink":   7,
	}
)

func (x TiltRequest_Color) Enum() *TiltRequest_Color {
	p := new(TiltRequest_Color)
	*p = x
	return p
}

func (x TiltRequest_Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TiltRequest_Color) Descriptor() protoreflect.EnumDescriptor {
	return file_tilt_proto_enumTypes[0].Descriptor()
}

func (TiltRequest_Color) Type() protoreflect.EnumType {
	return &file_tilt_proto_enumTypes[0]
}

func (x TiltRequest_Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TiltRequest_Color.Descriptor instead.
func (TiltRequest_Color) EnumDescriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{1, 0}
}

type TiltDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color           string                 `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Temperature     uint32                 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	SpecificGravity float32                `protobuf:"fixed32,3,opt,name=specificGravity,proto3" json:"specificGravity,omitempty"`
	LastReading     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lastReading,proto3" json:"lastReading,omitempty"`
}

func (x *TiltDevice) Reset() {
	*x = TiltDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tilt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiltDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiltDevice) ProtoMessage() {}

func (x *TiltDevice) ProtoReflect() protoreflect.Message {
	mi := &file_tilt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiltDevice.ProtoReflect.Descriptor instead.
func (*TiltDevice) Descriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{0}
}

func (x *TiltDevice) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TiltDevice) GetTemperature() uint32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *TiltDevice) GetSpecificGravity() float32 {
	if x != nil {
		return x.SpecificGravity
	}
	return 0
}

func (x *TiltDevice) GetLastReading() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReading
	}
	return nil
}

// The request message containing the user's name
type TiltRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color TiltRequest_Color `protobuf:"varint,1,opt,name=color,proto3,enum=tiltproxy.TiltRequest_Color" json:"color,omitempty"`
}

func (x *TiltRequest) Reset() {
	*x = TiltRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tilt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiltRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiltRequest) ProtoMessage() {}

func (x *TiltRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tilt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiltRequest.ProtoReflect.Descriptor instead.
func (*TiltRequest) Descriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{1}
}

func (x *TiltRequest) GetColor() TiltRequest_Color {
	if x != nil {
		return x.Color
	}
	return TiltRequest_Red
}

// The response message containing the greetings
type TiltResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tilt *TiltDevice `protobuf:"bytes,1,opt,name=tilt,proto3" json:"tilt,omitempty"`
}

func (x *TiltResponse) Reset() {
	*x = TiltResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tilt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiltResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiltResponse) ProtoMessage() {}

func (x *TiltResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tilt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiltResponse.ProtoReflect.Descriptor instead.
func (*TiltResponse) Descriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{2}
}

func (x *TiltResponse) GetTilt() *TiltDevice {
	if x != nil {
		return x.Tilt
	}
	return nil
}

type TiltsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TiltsRequest) Reset() {
	*x = TiltsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tilt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiltsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiltsRequest) ProtoMessage() {}

func (x *TiltsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tilt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiltsRequest.ProtoReflect.Descriptor instead.
func (*TiltsRequest) Descriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{3}
}

type TiltsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device []*TiltDevice `protobuf:"bytes,1,rep,name=device,proto3" json:"device,omitempty"`
}

func (x *TiltsResponse) Reset() {
	*x = TiltsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tilt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TiltsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TiltsResponse) ProtoMessage() {}

func (x *TiltsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tilt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TiltsResponse.ProtoReflect.Descriptor instead.
func (*TiltsResponse) Descriptor() ([]byte, []int) {
	return file_tilt_proto_rawDescGZIP(), []int{4}
}

func (x *TiltsResponse) GetDevice() []*TiltDevice {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_tilt_proto protoreflect.FileDescriptor

var file_tilt_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x69, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69,
	0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x54, 0x69, 0x6c, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xd7, 0x01, 0x0a, 0x0b, 0x54, 0x69, 0x6c, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x68, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x54, 0x69, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x42, 0x34, 0x92, 0x41, 0x31, 0x2a, 0x2f, 0x54, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6c, 0x74, 0x20, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22,
	0x5e, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x75, 0x72, 0x70, 0x6c,
	0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x42, 0x6c, 0x75, 0x65, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x6b, 0x10, 0x07, 0x22,
	0x39, 0x0a, 0x0c, 0x54, 0x69, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x74, 0x69, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x74, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54, 0x69, 0x6c, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x74, 0x69, 0x6c, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x69,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x0d, 0x54, 0x69,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x69,
	0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54, 0x69, 0x6c, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x32, 0xa7, 0x01, 0x0a, 0x0b, 0x54,
	0x69, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x54, 0x69,
	0x6c, 0x74, 0x12, 0x16, 0x2e, 0x74, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54,
	0x69, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x69, 0x6c,
	0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54, 0x69, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x69, 0x6c, 0x74, 0x12, 0x4d, 0x0a, 0x05, 0x54, 0x69, 0x6c, 0x74, 0x73, 0x12, 0x17,
	0x2e, 0x74, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54, 0x69, 0x6c, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x69, 0x6c, 0x74, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x54, 0x69, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x69, 0x6c, 0x74, 0x73, 0x42, 0x7b, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6c, 0x74,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x42, 0x09, 0x54, 0x69, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x74, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x69, 0x6c, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0xa2,
	0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x54, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0xca, 0x02, 0x09, 0x54, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0xe2, 0x02, 0x15,
	0x54, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x69, 0x6c, 0x74, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tilt_proto_rawDescOnce sync.Once
	file_tilt_proto_rawDescData = file_tilt_proto_rawDesc
)

func file_tilt_proto_rawDescGZIP() []byte {
	file_tilt_proto_rawDescOnce.Do(func() {
		file_tilt_proto_rawDescData = protoimpl.X.CompressGZIP(file_tilt_proto_rawDescData)
	})
	return file_tilt_proto_rawDescData
}

var file_tilt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tilt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tilt_proto_goTypes = []any{
	(TiltRequest_Color)(0),        // 0: tiltproxy.TiltRequest.Color
	(*TiltDevice)(nil),            // 1: tiltproxy.TiltDevice
	(*TiltRequest)(nil),           // 2: tiltproxy.TiltRequest
	(*TiltResponse)(nil),          // 3: tiltproxy.TiltResponse
	(*TiltsRequest)(nil),          // 4: tiltproxy.TiltsRequest
	(*TiltsResponse)(nil),         // 5: tiltproxy.TiltsResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_tilt_proto_depIdxs = []int32{
	6, // 0: tiltproxy.TiltDevice.lastReading:type_name -> google.protobuf.Timestamp
	0, // 1: tiltproxy.TiltRequest.color:type_name -> tiltproxy.TiltRequest.Color
	1, // 2: tiltproxy.TiltResponse.tilt:type_name -> tiltproxy.TiltDevice
	1, // 3: tiltproxy.TiltsResponse.device:type_name -> tiltproxy.TiltDevice
	2, // 4: tiltproxy.TiltService.Tilt:input_type -> tiltproxy.TiltRequest
	4, // 5: tiltproxy.TiltService.Tilts:input_type -> tiltproxy.TiltsRequest
	3, // 6: tiltproxy.TiltService.Tilt:output_type -> tiltproxy.TiltResponse
	5, // 7: tiltproxy.TiltService.Tilts:output_type -> tiltproxy.TiltsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tilt_proto_init() }
func file_tilt_proto_init() {
	if File_tilt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tilt_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TiltDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tilt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TiltRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tilt_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TiltResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tilt_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TiltsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tilt_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TiltsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tilt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tilt_proto_goTypes,
		DependencyIndexes: file_tilt_proto_depIdxs,
		EnumInfos:         file_tilt_proto_enumTypes,
		MessageInfos:      file_tilt_proto_msgTypes,
	}.Build()
	File_tilt_proto = out.File
	file_tilt_proto_rawDesc = nil
	file_tilt_proto_goTypes = nil
	file_tilt_proto_depIdxs = nil
}