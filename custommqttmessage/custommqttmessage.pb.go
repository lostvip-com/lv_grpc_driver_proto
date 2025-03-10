// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: custommqttmessage/custommqttmessage.proto

package custommqttmessage

import (
	drivercommon "github.com/lostvip-com/lv_grpc_driver_proto/drivercommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlatformCustomPublishRequest struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	BaseRequest   *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId      int64                            `protobuf:"varint,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topic         string                           `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Qos           int32                            `protobuf:"varint,4,opt,name=qos,proto3" json:"qos,omitempty"`
	Retained      bool                             `protobuf:"varint,5,opt,name=retained,proto3" json:"retained,omitempty"`
	Payload       string                           `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlatformCustomPublishRequest) Reset() {
	*x = PlatformCustomPublishRequest{}
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlatformCustomPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformCustomPublishRequest) ProtoMessage() {}

func (x *PlatformCustomPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformCustomPublishRequest.ProtoReflect.Descriptor instead.
func (*PlatformCustomPublishRequest) Descriptor() ([]byte, []int) {
	return file_custommqttmessage_custommqttmessage_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformCustomPublishRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PlatformCustomPublishRequest) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *PlatformCustomPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PlatformCustomPublishRequest) GetQos() int32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *PlatformCustomPublishRequest) GetRetained() bool {
	if x != nil {
		return x.Retained
	}
	return false
}

func (x *PlatformCustomPublishRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type PlatformCustomSubscribeRequest struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	BaseRequest   *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId      int64                            `protobuf:"varint,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topic         string                           `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Qos           int32                            `protobuf:"varint,4,opt,name=qos,proto3" json:"qos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlatformCustomSubscribeRequest) Reset() {
	*x = PlatformCustomSubscribeRequest{}
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlatformCustomSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformCustomSubscribeRequest) ProtoMessage() {}

func (x *PlatformCustomSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformCustomSubscribeRequest.ProtoReflect.Descriptor instead.
func (*PlatformCustomSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_custommqttmessage_custommqttmessage_proto_rawDescGZIP(), []int{1}
}

func (x *PlatformCustomSubscribeRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PlatformCustomSubscribeRequest) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *PlatformCustomSubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PlatformCustomSubscribeRequest) GetQos() int32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

type PlatformCustomUnSubscribeRequest struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	BaseRequest   *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId      int64                            `protobuf:"varint,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topics        []string                         `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlatformCustomUnSubscribeRequest) Reset() {
	*x = PlatformCustomUnSubscribeRequest{}
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlatformCustomUnSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformCustomUnSubscribeRequest) ProtoMessage() {}

func (x *PlatformCustomUnSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_custommqttmessage_custommqttmessage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformCustomUnSubscribeRequest.ProtoReflect.Descriptor instead.
func (*PlatformCustomUnSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_custommqttmessage_custommqttmessage_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformCustomUnSubscribeRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PlatformCustomUnSubscribeRequest) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *PlatformCustomUnSubscribeRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

var File_custommqttmessage_custommqttmessage_proto protoreflect.FileDescriptor

var file_custommqttmessage_custommqttmessage_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x1c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x1e, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x6f, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x20, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x32, 0xe0, 0x02, 0x0a, 0x14, 0x52, 0x70, 0x63, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x71,
	0x74, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x68, 0x0a, 0x15, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x2f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x31,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x70, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x33,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x6f, 0x73, 0x74, 0x76, 0x69, 0x70, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6d, 0x71, 0x74, 0x74, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_custommqttmessage_custommqttmessage_proto_rawDescOnce sync.Once
	file_custommqttmessage_custommqttmessage_proto_rawDescData []byte
)

func file_custommqttmessage_custommqttmessage_proto_rawDescGZIP() []byte {
	file_custommqttmessage_custommqttmessage_proto_rawDescOnce.Do(func() {
		file_custommqttmessage_custommqttmessage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_custommqttmessage_custommqttmessage_proto_rawDesc), len(file_custommqttmessage_custommqttmessage_proto_rawDesc)))
	})
	return file_custommqttmessage_custommqttmessage_proto_rawDescData
}

var file_custommqttmessage_custommqttmessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_custommqttmessage_custommqttmessage_proto_goTypes = []any{
	(*PlatformCustomPublishRequest)(nil),     // 0: custommqttmessage.PlatformCustomPublishRequest
	(*PlatformCustomSubscribeRequest)(nil),   // 1: custommqttmessage.PlatformCustomSubscribeRequest
	(*PlatformCustomUnSubscribeRequest)(nil), // 2: custommqttmessage.PlatformCustomUnSubscribeRequest
	(*drivercommon.BaseRequestMessage)(nil),  // 3: drivercommon.BaseRequestMessage
	(*drivercommon.CommonResponse)(nil),      // 4: drivercommon.CommonResponse
}
var file_custommqttmessage_custommqttmessage_proto_depIdxs = []int32{
	3, // 0: custommqttmessage.PlatformCustomPublishRequest.baseRequest:type_name -> drivercommon.BaseRequestMessage
	3, // 1: custommqttmessage.PlatformCustomSubscribeRequest.baseRequest:type_name -> drivercommon.BaseRequestMessage
	3, // 2: custommqttmessage.PlatformCustomUnSubscribeRequest.baseRequest:type_name -> drivercommon.BaseRequestMessage
	0, // 3: custommqttmessage.RpcCustomMqttMessage.PlatformCustomPublish:input_type -> custommqttmessage.PlatformCustomPublishRequest
	1, // 4: custommqttmessage.RpcCustomMqttMessage.PlatformCustomSubscribe:input_type -> custommqttmessage.PlatformCustomSubscribeRequest
	2, // 5: custommqttmessage.RpcCustomMqttMessage.PlatformCustomUnSubscribe:input_type -> custommqttmessage.PlatformCustomUnSubscribeRequest
	4, // 6: custommqttmessage.RpcCustomMqttMessage.PlatformCustomPublish:output_type -> drivercommon.CommonResponse
	4, // 7: custommqttmessage.RpcCustomMqttMessage.PlatformCustomSubscribe:output_type -> drivercommon.CommonResponse
	4, // 8: custommqttmessage.RpcCustomMqttMessage.PlatformCustomUnSubscribe:output_type -> drivercommon.CommonResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_custommqttmessage_custommqttmessage_proto_init() }
func file_custommqttmessage_custommqttmessage_proto_init() {
	if File_custommqttmessage_custommqttmessage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_custommqttmessage_custommqttmessage_proto_rawDesc), len(file_custommqttmessage_custommqttmessage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_custommqttmessage_custommqttmessage_proto_goTypes,
		DependencyIndexes: file_custommqttmessage_custommqttmessage_proto_depIdxs,
		MessageInfos:      file_custommqttmessage_custommqttmessage_proto_msgTypes,
	}.Build()
	File_custommqttmessage_custommqttmessage_proto = out.File
	file_custommqttmessage_custommqttmessage_proto_goTypes = nil
	file_custommqttmessage_custommqttmessage_proto_depIdxs = nil
}
