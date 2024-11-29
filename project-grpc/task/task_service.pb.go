// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: task_service.proto

package task

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

type TaskReqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId    int64  `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	ProjectCode string `protobuf:"bytes,2,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	Page        int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *TaskReqMessage) Reset() {
	*x = TaskReqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReqMessage) ProtoMessage() {}

func (x *TaskReqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReqMessage.ProtoReflect.Descriptor instead.
func (*TaskReqMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{0}
}

func (x *TaskReqMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *TaskReqMessage) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskReqMessage) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TaskReqMessage) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type TaskStagesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectCode string `protobuf:"bytes,3,opt,name=projectCode,proto3" json:"projectCode,omitempty"`
	Sort        int32  `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime  string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Deleted     int32  `protobuf:"varint,7,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Id          int32  `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskStagesMessage) Reset() {
	*x = TaskStagesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStagesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStagesMessage) ProtoMessage() {}

func (x *TaskStagesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStagesMessage.ProtoReflect.Descriptor instead.
func (*TaskStagesMessage) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *TaskStagesMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TaskStagesMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskStagesMessage) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *TaskStagesMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TaskStagesMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskStagesMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *TaskStagesMessage) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *TaskStagesMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskStagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*TaskStagesMessage `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TaskStagesResponse) Reset() {
	*x = TaskStagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStagesResponse) ProtoMessage() {}

func (x *TaskStagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStagesResponse.ProtoReflect.Descriptor instead.
func (*TaskStagesResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{2}
}

func (x *TaskStagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskStagesResponse) GetList() []*TaskStagesMessage {
	if x != nil {
		return x.List
	}
	return nil
}

var File_task_service_proto protoreflect.FileDescriptor

var file_task_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x63, 0x0a, 0x0b, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d,
	0x5a, 0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_service_proto_rawDescOnce sync.Once
	file_task_service_proto_rawDescData = file_task_service_proto_rawDesc
)

func file_task_service_proto_rawDescGZIP() []byte {
	file_task_service_proto_rawDescOnce.Do(func() {
		file_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_service_proto_rawDescData)
	})
	return file_task_service_proto_rawDescData
}

var file_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_task_service_proto_goTypes = []interface{}{
	(*TaskReqMessage)(nil),     // 0: task.service.v1.TaskReqMessage
	(*TaskStagesMessage)(nil),  // 1: task.service.v1.TaskStagesMessage
	(*TaskStagesResponse)(nil), // 2: task.service.v1.TaskStagesResponse
}
var file_task_service_proto_depIdxs = []int32{
	1, // 0: task.service.v1.TaskStagesResponse.list:type_name -> task.service.v1.TaskStagesMessage
	0, // 1: task.service.v1.TaskService.TaskStages:input_type -> task.service.v1.TaskReqMessage
	2, // 2: task.service.v1.TaskService.TaskStages:output_type -> task.service.v1.TaskStagesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_service_proto_init() }
func file_task_service_proto_init() {
	if File_task_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReqMessage); i {
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
		file_task_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStagesMessage); i {
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
		file_task_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStagesResponse); i {
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
			RawDescriptor: file_task_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_service_proto_goTypes,
		DependencyIndexes: file_task_service_proto_depIdxs,
		MessageInfos:      file_task_service_proto_msgTypes,
	}.Build()
	File_task_service_proto = out.File
	file_task_service_proto_rawDesc = nil
	file_task_service_proto_goTypes = nil
	file_task_service_proto_depIdxs = nil
}
