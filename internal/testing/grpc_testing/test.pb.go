// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: test.proto

package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A protobuf representation for grpc status. This is used by test
// clients to specify a status that the server should attempt to return.
type EchoStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoStatus) Reset() {
	*x = EchoStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoStatus) ProtoMessage() {}

func (x *EchoStatus) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoStatus.ProtoReflect.Descriptor instead.
func (*EchoStatus) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *EchoStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EchoStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Unary request.
type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseStatus *EchoStatus `protobuf:"bytes,1,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	User           *User       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleRequest) GetResponseStatus() *EchoStatus {
	if x != nil {
		return x.ResponseStatus
	}
	return nil
}

func (x *SimpleRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *StreamingInputCallRequest) Reset() {
	*x = StreamingInputCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingInputCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingInputCallRequest) ProtoMessage() {}

func (x *StreamingInputCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingInputCallRequest.ProtoReflect.Descriptor instead.
func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingInputCallRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize int32 `protobuf:"varint,1,opt,name=aggregated_payload_size,json=aggregatedPayloadSize,proto3" json:"aggregated_payload_size,omitempty"`
	User                  *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *StreamingInputCallResponse) Reset() {
	*x = StreamingInputCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingInputCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingInputCallResponse) ProtoMessage() {}

func (x *StreamingInputCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingInputCallResponse.ProtoReflect.Descriptor instead.
func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if x != nil {
		return x.AggregatedPayloadSize
	}
	return 0
}

func (x *StreamingInputCallResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Configuration for a particular response.
type ResponseParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs int32 `protobuf:"varint,2,opt,name=interval_us,json=intervalUs,proto3" json:"interval_us,omitempty"`
}

func (x *ResponseParameters) Reset() {
	*x = ResponseParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseParameters) ProtoMessage() {}

func (x *ResponseParameters) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseParameters.ProtoReflect.Descriptor instead.
func (*ResponseParameters) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseParameters) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ResponseParameters) GetIntervalUs() int32 {
	if x != nil {
		return x.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,1,rep,name=response_parameters,json=responseParameters,proto3" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Whether server should return a given status
	ResponseStatus *EchoStatus `protobuf:"bytes,3,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
}

func (x *StreamingOutputCallRequest) Reset() {
	*x = StreamingOutputCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOutputCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOutputCallRequest) ProtoMessage() {}

func (x *StreamingOutputCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOutputCallRequest.ProtoReflect.Descriptor instead.
func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{7}
}

func (x *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if x != nil {
		return x.ResponseParameters
	}
	return nil
}

func (x *StreamingOutputCallRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *StreamingOutputCallRequest) GetResponseStatus() *EchoStatus {
	if x != nil {
		return x.ResponseStatus
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload to increase response size.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *StreamingOutputCallResponse) Reset() {
	*x = StreamingOutputCallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingOutputCallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingOutputCallResponse) ProtoMessage() {}

func (x *StreamingOutputCallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingOutputCallResponse.ProtoReflect.Descriptor instead.
func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{8}
}

func (x *StreamingOutputCallResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0d, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x43, 0x0a,
	0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x4e, 0x0a, 0x19, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x55, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xb0, 0x05, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x5c, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x33, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c,
	0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x7f, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x32, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x7f, 0x0a, 0x0e, 0x46, 0x75, 0x6c, 0x6c,
	0x44, 0x75, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x33, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c,
	0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7f, 0x0a, 0x0e, 0x48, 0x61, 0x6c,
	0x66, 0x44, 0x75, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x33, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6c, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x3b,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_test_proto_goTypes = []interface{}{
	(*User)(nil),                        // 0: grpc_client_cli.testing.User
	(*EchoStatus)(nil),                  // 1: grpc_client_cli.testing.EchoStatus
	(*SimpleRequest)(nil),               // 2: grpc_client_cli.testing.SimpleRequest
	(*SimpleResponse)(nil),              // 3: grpc_client_cli.testing.SimpleResponse
	(*StreamingInputCallRequest)(nil),   // 4: grpc_client_cli.testing.StreamingInputCallRequest
	(*StreamingInputCallResponse)(nil),  // 5: grpc_client_cli.testing.StreamingInputCallResponse
	(*ResponseParameters)(nil),          // 6: grpc_client_cli.testing.ResponseParameters
	(*StreamingOutputCallRequest)(nil),  // 7: grpc_client_cli.testing.StreamingOutputCallRequest
	(*StreamingOutputCallResponse)(nil), // 8: grpc_client_cli.testing.StreamingOutputCallResponse
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_test_proto_depIdxs = []int32{
	1,  // 0: grpc_client_cli.testing.SimpleRequest.response_status:type_name -> grpc_client_cli.testing.EchoStatus
	0,  // 1: grpc_client_cli.testing.SimpleRequest.user:type_name -> grpc_client_cli.testing.User
	0,  // 2: grpc_client_cli.testing.SimpleResponse.user:type_name -> grpc_client_cli.testing.User
	0,  // 3: grpc_client_cli.testing.StreamingInputCallRequest.user:type_name -> grpc_client_cli.testing.User
	0,  // 4: grpc_client_cli.testing.StreamingInputCallResponse.user:type_name -> grpc_client_cli.testing.User
	6,  // 5: grpc_client_cli.testing.StreamingOutputCallRequest.response_parameters:type_name -> grpc_client_cli.testing.ResponseParameters
	0,  // 6: grpc_client_cli.testing.StreamingOutputCallRequest.user:type_name -> grpc_client_cli.testing.User
	1,  // 7: grpc_client_cli.testing.StreamingOutputCallRequest.response_status:type_name -> grpc_client_cli.testing.EchoStatus
	0,  // 8: grpc_client_cli.testing.StreamingOutputCallResponse.user:type_name -> grpc_client_cli.testing.User
	9,  // 9: grpc_client_cli.testing.TestService.EmptyCall:input_type -> google.protobuf.Empty
	2,  // 10: grpc_client_cli.testing.TestService.UnaryCall:input_type -> grpc_client_cli.testing.SimpleRequest
	7,  // 11: grpc_client_cli.testing.TestService.StreamingOutputCall:input_type -> grpc_client_cli.testing.StreamingOutputCallRequest
	4,  // 12: grpc_client_cli.testing.TestService.StreamingInputCall:input_type -> grpc_client_cli.testing.StreamingInputCallRequest
	7,  // 13: grpc_client_cli.testing.TestService.FullDuplexCall:input_type -> grpc_client_cli.testing.StreamingOutputCallRequest
	7,  // 14: grpc_client_cli.testing.TestService.HalfDuplexCall:input_type -> grpc_client_cli.testing.StreamingOutputCallRequest
	9,  // 15: grpc_client_cli.testing.TestService.EmptyCall:output_type -> google.protobuf.Empty
	3,  // 16: grpc_client_cli.testing.TestService.UnaryCall:output_type -> grpc_client_cli.testing.SimpleResponse
	8,  // 17: grpc_client_cli.testing.TestService.StreamingOutputCall:output_type -> grpc_client_cli.testing.StreamingOutputCallResponse
	5,  // 18: grpc_client_cli.testing.TestService.StreamingInputCall:output_type -> grpc_client_cli.testing.StreamingInputCallResponse
	8,  // 19: grpc_client_cli.testing.TestService.FullDuplexCall:output_type -> grpc_client_cli.testing.StreamingOutputCallResponse
	8,  // 20: grpc_client_cli.testing.TestService.HalfDuplexCall:output_type -> grpc_client_cli.testing.StreamingOutputCallResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoStatus); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingInputCallRequest); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingInputCallResponse); i {
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
		file_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseParameters); i {
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
		file_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOutputCallRequest); i {
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
		file_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingOutputCallResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
