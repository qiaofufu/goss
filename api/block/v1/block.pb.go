// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: api/block/v1/block.proto

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

type CreateBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Size     int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Num      int32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Checksum int32 `protobuf:"varint,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *CreateBlockRequest) Reset() {
	*x = CreateBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlockRequest) ProtoMessage() {}

func (x *CreateBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlockRequest.ProtoReflect.Descriptor instead.
func (*CreateBlockRequest) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlockRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *CreateBlockRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateBlockRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CreateBlockRequest) GetChecksum() int32 {
	if x != nil {
		return x.Checksum
	}
	return 0
}

type CreateBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId int64  `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Status  int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateBlockReply) Reset() {
	*x = CreateBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlockReply) ProtoMessage() {}

func (x *CreateBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlockReply.ProtoReflect.Descriptor instead.
func (*CreateBlockReply) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlockReply) GetBlockId() int64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *CreateBlockReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateBlockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *DeleteBlockRequest) Reset() {
	*x = DeleteBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlockRequest) ProtoMessage() {}

func (x *DeleteBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlockRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlockRequest) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBlockRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type DeleteBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteBlockReply) Reset() {
	*x = DeleteBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlockReply) ProtoMessage() {}

func (x *DeleteBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlockReply.ProtoReflect.Descriptor instead.
func (*DeleteBlockReply) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBlockReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteBlockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *ListBlockRequest) Reset() {
	*x = ListBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockRequest) ProtoMessage() {}

func (x *ListBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockRequest.ProtoReflect.Descriptor instead.
func (*ListBlockRequest) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{4}
}

func (x *ListBlockRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type ListBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*BlockInfo `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Status int32        `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string       `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Total  int64        `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListBlockReply) Reset() {
	*x = ListBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockReply) ProtoMessage() {}

func (x *ListBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockReply.ProtoReflect.Descriptor instead.
func (*ListBlockReply) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{5}
}

func (x *ListBlockReply) GetBlocks() []*BlockInfo {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *ListBlockReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListBlockReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListBlockReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId   int64 `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	ObjectId  int64 `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Size      int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Num       int32 `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	Checksum  int32 `protobuf:"varint,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
	CreatedAt int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64 `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_block_v1_block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_block_v1_block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_api_block_v1_block_proto_rawDescGZIP(), []int{6}
}

func (x *BlockInfo) GetBlockId() int64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *BlockInfo) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *BlockInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlockInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BlockInfo) GetChecksum() int32 {
	if x != nil {
		return x.Checksum
	}
	return 0
}

func (x *BlockInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *BlockInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_api_block_v1_block_proto protoreflect.FileDescriptor

var file_api_block_v1_block_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x73, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x57, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x31, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc3, 0x01, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x32, 0xf4, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x4f, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x28, 0x0a, 0x0c, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_block_v1_block_proto_rawDescOnce sync.Once
	file_api_block_v1_block_proto_rawDescData = file_api_block_v1_block_proto_rawDesc
)

func file_api_block_v1_block_proto_rawDescGZIP() []byte {
	file_api_block_v1_block_proto_rawDescOnce.Do(func() {
		file_api_block_v1_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_block_v1_block_proto_rawDescData)
	})
	return file_api_block_v1_block_proto_rawDescData
}

var file_api_block_v1_block_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_block_v1_block_proto_goTypes = []interface{}{
	(*CreateBlockRequest)(nil), // 0: api.block.v1.CreateBlockRequest
	(*CreateBlockReply)(nil),   // 1: api.block.v1.CreateBlockReply
	(*DeleteBlockRequest)(nil), // 2: api.block.v1.DeleteBlockRequest
	(*DeleteBlockReply)(nil),   // 3: api.block.v1.DeleteBlockReply
	(*ListBlockRequest)(nil),   // 4: api.block.v1.ListBlockRequest
	(*ListBlockReply)(nil),     // 5: api.block.v1.ListBlockReply
	(*BlockInfo)(nil),          // 6: api.block.v1.BlockInfo
}
var file_api_block_v1_block_proto_depIdxs = []int32{
	6, // 0: api.block.v1.ListBlockReply.blocks:type_name -> api.block.v1.BlockInfo
	0, // 1: api.block.v1.Block.CreateBlock:input_type -> api.block.v1.CreateBlockRequest
	2, // 2: api.block.v1.Block.DeleteBlock:input_type -> api.block.v1.DeleteBlockRequest
	4, // 3: api.block.v1.Block.ListBlock:input_type -> api.block.v1.ListBlockRequest
	1, // 4: api.block.v1.Block.CreateBlock:output_type -> api.block.v1.CreateBlockReply
	3, // 5: api.block.v1.Block.DeleteBlock:output_type -> api.block.v1.DeleteBlockReply
	5, // 6: api.block.v1.Block.ListBlock:output_type -> api.block.v1.ListBlockReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_block_v1_block_proto_init() }
func file_api_block_v1_block_proto_init() {
	if File_api_block_v1_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_block_v1_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlockRequest); i {
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
		file_api_block_v1_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlockReply); i {
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
		file_api_block_v1_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlockRequest); i {
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
		file_api_block_v1_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlockReply); i {
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
		file_api_block_v1_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlockRequest); i {
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
		file_api_block_v1_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlockReply); i {
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
		file_api_block_v1_block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
			RawDescriptor: file_api_block_v1_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_block_v1_block_proto_goTypes,
		DependencyIndexes: file_api_block_v1_block_proto_depIdxs,
		MessageInfos:      file_api_block_v1_block_proto_msgTypes,
	}.Build()
	File_api_block_v1_block_proto = out.File
	file_api_block_v1_block_proto_rawDesc = nil
	file_api_block_v1_block_proto_goTypes = nil
	file_api_block_v1_block_proto_depIdxs = nil
}
