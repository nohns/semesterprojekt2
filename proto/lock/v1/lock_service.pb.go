// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/lock_service.proto

package lockv1

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

type GetLockStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLockStateRequest) Reset() {
	*x = GetLockStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_lock_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockStateRequest) ProtoMessage() {}

func (x *GetLockStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_lock_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockStateRequest.ProtoReflect.Descriptor instead.
func (*GetLockStateRequest) Descriptor() ([]byte, []int) {
	return file_v1_lock_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLockStateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLockStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engaged bool `protobuf:"varint,1,opt,name=engaged,proto3" json:"engaged,omitempty"`
}

func (x *GetLockStateResponse) Reset() {
	*x = GetLockStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_lock_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockStateResponse) ProtoMessage() {}

func (x *GetLockStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_lock_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockStateResponse.ProtoReflect.Descriptor instead.
func (*GetLockStateResponse) Descriptor() ([]byte, []int) {
	return file_v1_lock_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLockStateResponse) GetEngaged() bool {
	if x != nil {
		return x.Engaged
	}
	return false
}

type SetLockStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Engaged bool   `protobuf:"varint,2,opt,name=engaged,proto3" json:"engaged,omitempty"`
}

func (x *SetLockStateRequest) Reset() {
	*x = SetLockStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_lock_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLockStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLockStateRequest) ProtoMessage() {}

func (x *SetLockStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_lock_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLockStateRequest.ProtoReflect.Descriptor instead.
func (*SetLockStateRequest) Descriptor() ([]byte, []int) {
	return file_v1_lock_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetLockStateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetLockStateRequest) GetEngaged() bool {
	if x != nil {
		return x.Engaged
	}
	return false
}

type SetLockStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engaged bool `protobuf:"varint,1,opt,name=engaged,proto3" json:"engaged,omitempty"`
}

func (x *SetLockStateResponse) Reset() {
	*x = SetLockStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_lock_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLockStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLockStateResponse) ProtoMessage() {}

func (x *SetLockStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_lock_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLockStateResponse.ProtoReflect.Descriptor instead.
func (*SetLockStateResponse) Descriptor() ([]byte, []int) {
	return file_v1_lock_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetLockStateResponse) GetEngaged() bool {
	if x != nil {
		return x.Engaged
	}
	return false
}

var File_v1_lock_service_proto protoreflect.FileDescriptor

var file_v1_lock_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65,
	0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x64, 0x32, 0xab, 0x01, 0x0a,
	0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x53,
	0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x68, 0x6e, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63,
	0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_lock_service_proto_rawDescOnce sync.Once
	file_v1_lock_service_proto_rawDescData = file_v1_lock_service_proto_rawDesc
)

func file_v1_lock_service_proto_rawDescGZIP() []byte {
	file_v1_lock_service_proto_rawDescOnce.Do(func() {
		file_v1_lock_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_lock_service_proto_rawDescData)
	})
	return file_v1_lock_service_proto_rawDescData
}

var file_v1_lock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_lock_service_proto_goTypes = []interface{}{
	(*GetLockStateRequest)(nil),  // 0: lock.v1.GetLockStateRequest
	(*GetLockStateResponse)(nil), // 1: lock.v1.GetLockStateResponse
	(*SetLockStateRequest)(nil),  // 2: lock.v1.SetLockStateRequest
	(*SetLockStateResponse)(nil), // 3: lock.v1.SetLockStateResponse
}
var file_v1_lock_service_proto_depIdxs = []int32{
	0, // 0: lock.v1.LockService.GetLockState:input_type -> lock.v1.GetLockStateRequest
	2, // 1: lock.v1.LockService.SetLockState:input_type -> lock.v1.SetLockStateRequest
	1, // 2: lock.v1.LockService.GetLockState:output_type -> lock.v1.GetLockStateResponse
	3, // 3: lock.v1.LockService.SetLockState:output_type -> lock.v1.SetLockStateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_lock_service_proto_init() }
func file_v1_lock_service_proto_init() {
	if File_v1_lock_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_lock_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockStateRequest); i {
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
		file_v1_lock_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockStateResponse); i {
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
		file_v1_lock_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLockStateRequest); i {
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
		file_v1_lock_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLockStateResponse); i {
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
			RawDescriptor: file_v1_lock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_lock_service_proto_goTypes,
		DependencyIndexes: file_v1_lock_service_proto_depIdxs,
		MessageInfos:      file_v1_lock_service_proto_msgTypes,
	}.Build()
	File_v1_lock_service_proto = out.File
	file_v1_lock_service_proto_rawDesc = nil
	file_v1_lock_service_proto_goTypes = nil
	file_v1_lock_service_proto_depIdxs = nil
}
