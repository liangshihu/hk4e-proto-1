// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_routine.proto

package pb

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

type PlayerRoutineDataNotify_CmdId int32

const (
	PlayerRoutineDataNotify_NONE             PlayerRoutineDataNotify_CmdId = 0
	PlayerRoutineDataNotify_CMD_ID           PlayerRoutineDataNotify_CmdId = 3526
	PlayerRoutineDataNotify_ENET_CHANNEL_ID  PlayerRoutineDataNotify_CmdId = 0
	PlayerRoutineDataNotify_ENET_IS_RELIABLE PlayerRoutineDataNotify_CmdId = 1
)

// Enum value maps for PlayerRoutineDataNotify_CmdId.
var (
	PlayerRoutineDataNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3526: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerRoutineDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3526,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerRoutineDataNotify_CmdId) Enum() *PlayerRoutineDataNotify_CmdId {
	p := new(PlayerRoutineDataNotify_CmdId)
	*p = x
	return p
}

func (x PlayerRoutineDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerRoutineDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_routine_proto_enumTypes[0].Descriptor()
}

func (PlayerRoutineDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_routine_proto_enumTypes[0]
}

func (x PlayerRoutineDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerRoutineDataNotify_CmdId.Descriptor instead.
func (PlayerRoutineDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{1, 0}
}

type WorldAllRoutineTypeNotify_CmdId int32

const (
	WorldAllRoutineTypeNotify_NONE             WorldAllRoutineTypeNotify_CmdId = 0
	WorldAllRoutineTypeNotify_CMD_ID           WorldAllRoutineTypeNotify_CmdId = 3518
	WorldAllRoutineTypeNotify_ENET_CHANNEL_ID  WorldAllRoutineTypeNotify_CmdId = 0
	WorldAllRoutineTypeNotify_ENET_IS_RELIABLE WorldAllRoutineTypeNotify_CmdId = 1
)

// Enum value maps for WorldAllRoutineTypeNotify_CmdId.
var (
	WorldAllRoutineTypeNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3518: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	WorldAllRoutineTypeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3518,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x WorldAllRoutineTypeNotify_CmdId) Enum() *WorldAllRoutineTypeNotify_CmdId {
	p := new(WorldAllRoutineTypeNotify_CmdId)
	*p = x
	return p
}

func (x WorldAllRoutineTypeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldAllRoutineTypeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_routine_proto_enumTypes[1].Descriptor()
}

func (WorldAllRoutineTypeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_routine_proto_enumTypes[1]
}

func (x WorldAllRoutineTypeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldAllRoutineTypeNotify_CmdId.Descriptor instead.
func (WorldAllRoutineTypeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{4, 0}
}

type WorldRoutineTypeRefreshNotify_CmdId int32

const (
	WorldRoutineTypeRefreshNotify_NONE             WorldRoutineTypeRefreshNotify_CmdId = 0
	WorldRoutineTypeRefreshNotify_CMD_ID           WorldRoutineTypeRefreshNotify_CmdId = 3525
	WorldRoutineTypeRefreshNotify_ENET_CHANNEL_ID  WorldRoutineTypeRefreshNotify_CmdId = 0
	WorldRoutineTypeRefreshNotify_ENET_IS_RELIABLE WorldRoutineTypeRefreshNotify_CmdId = 1
)

// Enum value maps for WorldRoutineTypeRefreshNotify_CmdId.
var (
	WorldRoutineTypeRefreshNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3525: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	WorldRoutineTypeRefreshNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3525,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x WorldRoutineTypeRefreshNotify_CmdId) Enum() *WorldRoutineTypeRefreshNotify_CmdId {
	p := new(WorldRoutineTypeRefreshNotify_CmdId)
	*p = x
	return p
}

func (x WorldRoutineTypeRefreshNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldRoutineTypeRefreshNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_routine_proto_enumTypes[2].Descriptor()
}

func (WorldRoutineTypeRefreshNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_routine_proto_enumTypes[2]
}

func (x WorldRoutineTypeRefreshNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldRoutineTypeRefreshNotify_CmdId.Descriptor instead.
func (WorldRoutineTypeRefreshNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{5, 0}
}

type WorldRoutineChangeNotify_CmdId int32

const (
	WorldRoutineChangeNotify_NONE             WorldRoutineChangeNotify_CmdId = 0
	WorldRoutineChangeNotify_CMD_ID           WorldRoutineChangeNotify_CmdId = 3507
	WorldRoutineChangeNotify_ENET_CHANNEL_ID  WorldRoutineChangeNotify_CmdId = 0
	WorldRoutineChangeNotify_ENET_IS_RELIABLE WorldRoutineChangeNotify_CmdId = 1
)

// Enum value maps for WorldRoutineChangeNotify_CmdId.
var (
	WorldRoutineChangeNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3507: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	WorldRoutineChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3507,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x WorldRoutineChangeNotify_CmdId) Enum() *WorldRoutineChangeNotify_CmdId {
	p := new(WorldRoutineChangeNotify_CmdId)
	*p = x
	return p
}

func (x WorldRoutineChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldRoutineChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_routine_proto_enumTypes[3].Descriptor()
}

func (WorldRoutineChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_routine_proto_enumTypes[3]
}

func (x WorldRoutineChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldRoutineChangeNotify_CmdId.Descriptor instead.
func (WorldRoutineChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{6, 0}
}

type WorldRoutineTypeCloseNotify_CmdId int32

const (
	WorldRoutineTypeCloseNotify_NONE             WorldRoutineTypeCloseNotify_CmdId = 0
	WorldRoutineTypeCloseNotify_CMD_ID           WorldRoutineTypeCloseNotify_CmdId = 3502
	WorldRoutineTypeCloseNotify_ENET_CHANNEL_ID  WorldRoutineTypeCloseNotify_CmdId = 0
	WorldRoutineTypeCloseNotify_ENET_IS_RELIABLE WorldRoutineTypeCloseNotify_CmdId = 1
)

// Enum value maps for WorldRoutineTypeCloseNotify_CmdId.
var (
	WorldRoutineTypeCloseNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3502: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	WorldRoutineTypeCloseNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3502,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x WorldRoutineTypeCloseNotify_CmdId) Enum() *WorldRoutineTypeCloseNotify_CmdId {
	p := new(WorldRoutineTypeCloseNotify_CmdId)
	*p = x
	return p
}

func (x WorldRoutineTypeCloseNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldRoutineTypeCloseNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_routine_proto_enumTypes[4].Descriptor()
}

func (WorldRoutineTypeCloseNotify_CmdId) Type() protoreflect.EnumType {
	return &file_cmd_cmd_routine_proto_enumTypes[4]
}

func (x WorldRoutineTypeCloseNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldRoutineTypeCloseNotify_CmdId.Descriptor instead.
func (WorldRoutineTypeCloseNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{7, 0}
}

type PlayerRoutineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType uint32 `protobuf:"varint,8,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
	FinishedNum uint32 `protobuf:"varint,15,opt,name=finished_num,json=finishedNum,proto3" json:"finished_num,omitempty"`
}

func (x *PlayerRoutineInfo) Reset() {
	*x = PlayerRoutineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRoutineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRoutineInfo) ProtoMessage() {}

func (x *PlayerRoutineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRoutineInfo.ProtoReflect.Descriptor instead.
func (*PlayerRoutineInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRoutineInfo) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

func (x *PlayerRoutineInfo) GetFinishedNum() uint32 {
	if x != nil {
		return x.FinishedNum
	}
	return 0
}

type PlayerRoutineDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineInfoList []*PlayerRoutineInfo `protobuf:"bytes,11,rep,name=routine_info_list,json=routineInfoList,proto3" json:"routine_info_list,omitempty"`
}

func (x *PlayerRoutineDataNotify) Reset() {
	*x = PlayerRoutineDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRoutineDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRoutineDataNotify) ProtoMessage() {}

func (x *PlayerRoutineDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRoutineDataNotify.ProtoReflect.Descriptor instead.
func (*PlayerRoutineDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerRoutineDataNotify) GetRoutineInfoList() []*PlayerRoutineInfo {
	if x != nil {
		return x.RoutineInfoList
	}
	return nil
}

type WorldRoutineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress       uint32 `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`
	IsFinished     bool   `protobuf:"varint,14,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	FinishProgress uint32 `protobuf:"varint,3,opt,name=finish_progress,json=finishProgress,proto3" json:"finish_progress,omitempty"`
	RoutineId      uint32 `protobuf:"varint,11,opt,name=routine_id,json=routineId,proto3" json:"routine_id,omitempty"`
}

func (x *WorldRoutineInfo) Reset() {
	*x = WorldRoutineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineInfo) ProtoMessage() {}

func (x *WorldRoutineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineInfo.ProtoReflect.Descriptor instead.
func (*WorldRoutineInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{2}
}

func (x *WorldRoutineInfo) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *WorldRoutineInfo) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *WorldRoutineInfo) GetFinishProgress() uint32 {
	if x != nil {
		return x.FinishProgress
	}
	return 0
}

func (x *WorldRoutineInfo) GetRoutineId() uint32 {
	if x != nil {
		return x.RoutineId
	}
	return 0
}

type WorldRoutineTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType          uint32              `protobuf:"varint,13,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
	NextRefreshTime      uint32              `protobuf:"varint,12,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	WorldRoutineInfoList []*WorldRoutineInfo `protobuf:"bytes,3,rep,name=world_routine_info_list,json=worldRoutineInfoList,proto3" json:"world_routine_info_list,omitempty"`
}

func (x *WorldRoutineTypeInfo) Reset() {
	*x = WorldRoutineTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeInfo) ProtoMessage() {}

func (x *WorldRoutineTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeInfo.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{3}
}

func (x *WorldRoutineTypeInfo) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetWorldRoutineInfoList() []*WorldRoutineInfo {
	if x != nil {
		return x.WorldRoutineInfoList
	}
	return nil
}

type WorldAllRoutineTypeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldRoutineTypeList []*WorldRoutineTypeInfo `protobuf:"bytes,12,rep,name=world_routine_type_list,json=worldRoutineTypeList,proto3" json:"world_routine_type_list,omitempty"`
}

func (x *WorldAllRoutineTypeNotify) Reset() {
	*x = WorldAllRoutineTypeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldAllRoutineTypeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldAllRoutineTypeNotify) ProtoMessage() {}

func (x *WorldAllRoutineTypeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldAllRoutineTypeNotify.ProtoReflect.Descriptor instead.
func (*WorldAllRoutineTypeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{4}
}

func (x *WorldAllRoutineTypeNotify) GetWorldRoutineTypeList() []*WorldRoutineTypeInfo {
	if x != nil {
		return x.WorldRoutineTypeList
	}
	return nil
}

type WorldRoutineTypeRefreshNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldRoutineType *WorldRoutineTypeInfo `protobuf:"bytes,7,opt,name=world_routine_type,json=worldRoutineType,proto3" json:"world_routine_type,omitempty"`
}

func (x *WorldRoutineTypeRefreshNotify) Reset() {
	*x = WorldRoutineTypeRefreshNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeRefreshNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeRefreshNotify) ProtoMessage() {}

func (x *WorldRoutineTypeRefreshNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeRefreshNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeRefreshNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{5}
}

func (x *WorldRoutineTypeRefreshNotify) GetWorldRoutineType() *WorldRoutineTypeInfo {
	if x != nil {
		return x.WorldRoutineType
	}
	return nil
}

type WorldRoutineChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineInfo *WorldRoutineInfo `protobuf:"bytes,3,opt,name=routine_info,json=routineInfo,proto3" json:"routine_info,omitempty"`
	RoutineType uint32            `protobuf:"varint,11,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
}

func (x *WorldRoutineChangeNotify) Reset() {
	*x = WorldRoutineChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineChangeNotify) ProtoMessage() {}

func (x *WorldRoutineChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineChangeNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineChangeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{6}
}

func (x *WorldRoutineChangeNotify) GetRoutineInfo() *WorldRoutineInfo {
	if x != nil {
		return x.RoutineInfo
	}
	return nil
}

func (x *WorldRoutineChangeNotify) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

type WorldRoutineTypeCloseNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType uint32 `protobuf:"varint,7,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
}

func (x *WorldRoutineTypeCloseNotify) Reset() {
	*x = WorldRoutineTypeCloseNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeCloseNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeCloseNotify) ProtoMessage() {}

func (x *WorldRoutineTypeCloseNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeCloseNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeCloseNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{7}
}

func (x *WorldRoutineTypeCloseNotify) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

var File_cmd_cmd_routine_proto protoreflect.FileDescriptor

var file_cmd_cmd_routine_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0xae, 0x01, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x44, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc6, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x17, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x52, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xbe, 0x1b, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xb9, 0x01, 0x0a, 0x1d,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x49, 0x0a,
	0x12, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc5, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xc8, 0x01, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xb3, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xae, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_routine_proto_rawDescOnce sync.Once
	file_cmd_cmd_routine_proto_rawDescData = file_cmd_cmd_routine_proto_rawDesc
)

func file_cmd_cmd_routine_proto_rawDescGZIP() []byte {
	file_cmd_cmd_routine_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_routine_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_routine_proto_rawDescData)
	})
	return file_cmd_cmd_routine_proto_rawDescData
}

var file_cmd_cmd_routine_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_cmd_cmd_routine_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_cmd_routine_proto_goTypes = []interface{}{
	(PlayerRoutineDataNotify_CmdId)(0),       // 0: proto.PlayerRoutineDataNotify.CmdId
	(WorldAllRoutineTypeNotify_CmdId)(0),     // 1: proto.WorldAllRoutineTypeNotify.CmdId
	(WorldRoutineTypeRefreshNotify_CmdId)(0), // 2: proto.WorldRoutineTypeRefreshNotify.CmdId
	(WorldRoutineChangeNotify_CmdId)(0),      // 3: proto.WorldRoutineChangeNotify.CmdId
	(WorldRoutineTypeCloseNotify_CmdId)(0),   // 4: proto.WorldRoutineTypeCloseNotify.CmdId
	(*PlayerRoutineInfo)(nil),                // 5: proto.PlayerRoutineInfo
	(*PlayerRoutineDataNotify)(nil),          // 6: proto.PlayerRoutineDataNotify
	(*WorldRoutineInfo)(nil),                 // 7: proto.WorldRoutineInfo
	(*WorldRoutineTypeInfo)(nil),             // 8: proto.WorldRoutineTypeInfo
	(*WorldAllRoutineTypeNotify)(nil),        // 9: proto.WorldAllRoutineTypeNotify
	(*WorldRoutineTypeRefreshNotify)(nil),    // 10: proto.WorldRoutineTypeRefreshNotify
	(*WorldRoutineChangeNotify)(nil),         // 11: proto.WorldRoutineChangeNotify
	(*WorldRoutineTypeCloseNotify)(nil),      // 12: proto.WorldRoutineTypeCloseNotify
}
var file_cmd_cmd_routine_proto_depIdxs = []int32{
	5, // 0: proto.PlayerRoutineDataNotify.routine_info_list:type_name -> proto.PlayerRoutineInfo
	7, // 1: proto.WorldRoutineTypeInfo.world_routine_info_list:type_name -> proto.WorldRoutineInfo
	8, // 2: proto.WorldAllRoutineTypeNotify.world_routine_type_list:type_name -> proto.WorldRoutineTypeInfo
	8, // 3: proto.WorldRoutineTypeRefreshNotify.world_routine_type:type_name -> proto.WorldRoutineTypeInfo
	7, // 4: proto.WorldRoutineChangeNotify.routine_info:type_name -> proto.WorldRoutineInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cmd_cmd_routine_proto_init() }
func file_cmd_cmd_routine_proto_init() {
	if File_cmd_cmd_routine_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_routine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRoutineInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRoutineDataNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldAllRoutineTypeNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeRefreshNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineChangeNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeCloseNotify); i {
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
			RawDescriptor: file_cmd_cmd_routine_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_routine_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_routine_proto_depIdxs,
		EnumInfos:         file_cmd_cmd_routine_proto_enumTypes,
		MessageInfos:      file_cmd_cmd_routine_proto_msgTypes,
	}.Build()
	File_cmd_cmd_routine_proto = out.File
	file_cmd_cmd_routine_proto_rawDesc = nil
	file_cmd_cmd_routine_proto_goTypes = nil
	file_cmd_cmd_routine_proto_depIdxs = nil
}