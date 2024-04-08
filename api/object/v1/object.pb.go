// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: api/object/v1/object.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	BucketId int64  `protobuf:"varint,3,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Replicas int32  `protobuf:"varint,5,opt,name=replicas,proto3" json:"replicas,omitempty"`
}

func (x *CreateObjectRequest) Reset() {
	*x = CreateObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateObjectRequest) ProtoMessage() {}

func (x *CreateObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateObjectRequest.ProtoReflect.Descriptor instead.
func (*CreateObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{0}
}

func (x *CreateObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateObjectRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateObjectRequest) GetBucketId() int64 {
	if x != nil {
		return x.BucketId
	}
	return 0
}

func (x *CreateObjectRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateObjectRequest) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

type CreateObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64  `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Status   int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg      string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateObjectReply) Reset() {
	*x = CreateObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateObjectReply) ProtoMessage() {}

func (x *CreateObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateObjectReply.ProtoReflect.Descriptor instead.
func (*CreateObjectReply) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{1}
}

func (x *CreateObjectReply) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *CreateObjectReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64  `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BucketId int64  `protobuf:"varint,4,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UpdateObjectRequest) Reset() {
	*x = UpdateObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObjectRequest) ProtoMessage() {}

func (x *UpdateObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateObjectRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *UpdateObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateObjectRequest) GetBucketId() int64 {
	if x != nil {
		return x.BucketId
	}
	return 0
}

func (x *UpdateObjectRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UpdateObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateObjectReply) Reset() {
	*x = UpdateObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObjectReply) ProtoMessage() {}

func (x *UpdateObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObjectReply.ProtoReflect.Descriptor instead.
func (*UpdateObjectReply) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateObjectReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *DeleteObjectRequest) Reset() {
	*x = DeleteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectRequest) ProtoMessage() {}

func (x *DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteObjectRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type DeleteObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteObjectReply) Reset() {
	*x = DeleteObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectReply) ProtoMessage() {}

func (x *DeleteObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectReply.ProtoReflect.Descriptor instead.
func (*DeleteObjectReply) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteObjectReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *GetObjectRequest) Reset() {
	*x = GetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectRequest) ProtoMessage() {}

func (x *GetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectRequest.ProtoReflect.Descriptor instead.
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{6}
}

func (x *GetObjectRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type GetObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *ObjectInfo `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Status int32       `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetObjectReply) Reset() {
	*x = GetObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectReply) ProtoMessage() {}

func (x *GetObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectReply.ProtoReflect.Descriptor instead.
func (*GetObjectReply) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{7}
}

func (x *GetObjectReply) GetObject() *ObjectInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *GetObjectReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketId int64 `protobuf:"varint,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListObjectRequest) Reset() {
	*x = ListObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectRequest) ProtoMessage() {}

func (x *ListObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectRequest.ProtoReflect.Descriptor instead.
func (*ListObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{8}
}

func (x *ListObjectRequest) GetBucketId() int64 {
	if x != nil {
		return x.BucketId
	}
	return 0
}

func (x *ListObjectRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListObjectRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*ObjectInfo `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	Total   int64         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Status  int32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ListObjectReply) Reset() {
	*x = ListObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectReply) ProtoMessage() {}

func (x *ListObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectReply.ProtoReflect.Descriptor instead.
func (*ListObjectReply) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{9}
}

func (x *ListObjectReply) GetObjects() []*ObjectInfo {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *ListObjectReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListObjectReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId  int64  `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	BucketId  int64  `protobuf:"varint,4,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Replicas  int32  `protobuf:"varint,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
	CreatedAt int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ObjectInfo) Reset() {
	*x = ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_object_v1_object_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectInfo) ProtoMessage() {}

func (x *ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_object_v1_object_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectInfo.ProtoReflect.Descriptor instead.
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return file_api_object_v1_object_proto_rawDescGZIP(), []int{10}
}

func (x *ObjectInfo) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *ObjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ObjectInfo) GetBucketId() int64 {
	if x != nil {
		return x.BucketId
	}
	return 0
}

func (x *ObjectInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ObjectInfo) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *ObjectInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ObjectInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_api_object_v1_object_proto protoreflect.FileDescriptor

var file_api_object_v1_object_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x8a, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x77, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x32, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x6d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x61, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xdc, 0x01, 0x0a, 0x0a,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xa7, 0x03, 0x0a, 0x06, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x54, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x2a, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_object_v1_object_proto_rawDescOnce sync.Once
	file_api_object_v1_object_proto_rawDescData = file_api_object_v1_object_proto_rawDesc
)

func file_api_object_v1_object_proto_rawDescGZIP() []byte {
	file_api_object_v1_object_proto_rawDescOnce.Do(func() {
		file_api_object_v1_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_object_v1_object_proto_rawDescData)
	})
	return file_api_object_v1_object_proto_rawDescData
}

var file_api_object_v1_object_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_object_v1_object_proto_goTypes = []interface{}{
	(*CreateObjectRequest)(nil), // 0: api.object.v1.CreateObjectRequest
	(*CreateObjectReply)(nil),   // 1: api.object.v1.CreateObjectReply
	(*UpdateObjectRequest)(nil), // 2: api.object.v1.UpdateObjectRequest
	(*UpdateObjectReply)(nil),   // 3: api.object.v1.UpdateObjectReply
	(*DeleteObjectRequest)(nil), // 4: api.object.v1.DeleteObjectRequest
	(*DeleteObjectReply)(nil),   // 5: api.object.v1.DeleteObjectReply
	(*GetObjectRequest)(nil),    // 6: api.object.v1.GetObjectRequest
	(*GetObjectReply)(nil),      // 7: api.object.v1.GetObjectReply
	(*ListObjectRequest)(nil),   // 8: api.object.v1.ListObjectRequest
	(*ListObjectReply)(nil),     // 9: api.object.v1.ListObjectReply
	(*ObjectInfo)(nil),          // 10: api.object.v1.ObjectInfo
}
var file_api_object_v1_object_proto_depIdxs = []int32{
	10, // 0: api.object.v1.GetObjectReply.object:type_name -> api.object.v1.ObjectInfo
	10, // 1: api.object.v1.ListObjectReply.objects:type_name -> api.object.v1.ObjectInfo
	0,  // 2: api.object.v1.Object.CreateObject:input_type -> api.object.v1.CreateObjectRequest
	2,  // 3: api.object.v1.Object.UpdateObject:input_type -> api.object.v1.UpdateObjectRequest
	4,  // 4: api.object.v1.Object.DeleteObject:input_type -> api.object.v1.DeleteObjectRequest
	6,  // 5: api.object.v1.Object.GetObject:input_type -> api.object.v1.GetObjectRequest
	8,  // 6: api.object.v1.Object.ListObject:input_type -> api.object.v1.ListObjectRequest
	1,  // 7: api.object.v1.Object.CreateObject:output_type -> api.object.v1.CreateObjectReply
	3,  // 8: api.object.v1.Object.UpdateObject:output_type -> api.object.v1.UpdateObjectReply
	5,  // 9: api.object.v1.Object.DeleteObject:output_type -> api.object.v1.DeleteObjectReply
	7,  // 10: api.object.v1.Object.GetObject:output_type -> api.object.v1.GetObjectReply
	9,  // 11: api.object.v1.Object.ListObject:output_type -> api.object.v1.ListObjectReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_object_v1_object_proto_init() }
func file_api_object_v1_object_proto_init() {
	if File_api_object_v1_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_object_v1_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateObjectRequest); i {
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
		file_api_object_v1_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateObjectReply); i {
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
		file_api_object_v1_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObjectRequest); i {
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
		file_api_object_v1_object_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObjectReply); i {
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
		file_api_object_v1_object_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectRequest); i {
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
		file_api_object_v1_object_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectReply); i {
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
		file_api_object_v1_object_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectRequest); i {
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
		file_api_object_v1_object_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectReply); i {
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
		file_api_object_v1_object_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectRequest); i {
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
		file_api_object_v1_object_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectReply); i {
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
		file_api_object_v1_object_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectInfo); i {
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
			RawDescriptor: file_api_object_v1_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_object_v1_object_proto_goTypes,
		DependencyIndexes: file_api_object_v1_object_proto_depIdxs,
		MessageInfos:      file_api_object_v1_object_proto_msgTypes,
	}.Build()
	File_api_object_v1_object_proto = out.File
	file_api_object_v1_object_proto_rawDesc = nil
	file_api_object_v1_object_proto_goTypes = nil
	file_api_object_v1_object_proto_depIdxs = nil
}
