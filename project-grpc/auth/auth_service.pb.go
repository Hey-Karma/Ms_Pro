// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: auth_service.proto

package auth

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

type AuthReqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId         int64    `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	OrganizationCode string   `protobuf:"bytes,2,opt,name=organizationCode,proto3" json:"organizationCode,omitempty"`
	Page             int64    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize         int64    `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Action           string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	AuthId           int64    `protobuf:"varint,6,opt,name=authId,proto3" json:"authId,omitempty"`
	Nodes            []string `protobuf:"bytes,7,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *AuthReqMessage) Reset() {
	*x = AuthReqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReqMessage) ProtoMessage() {}

func (x *AuthReqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReqMessage.ProtoReflect.Descriptor instead.
func (*AuthReqMessage) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReqMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *AuthReqMessage) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *AuthReqMessage) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AuthReqMessage) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AuthReqMessage) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuthReqMessage) GetAuthId() int64 {
	if x != nil {
		return x.AuthId
	}
	return 0
}

func (x *AuthReqMessage) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ProjectAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationCode string `protobuf:"bytes,2,opt,name=OrganizationCode,proto3" json:"OrganizationCode,omitempty"`
	Title            string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	CreateAt         string `protobuf:"bytes,4,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	Sort             int32  `protobuf:"varint,5,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Status           int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Desc             string `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
	CreateBy         int64  `protobuf:"varint,8,opt,name=CreateBy,proto3" json:"CreateBy,omitempty"`
	IsDefault        int32  `protobuf:"varint,9,opt,name=IsDefault,proto3" json:"IsDefault,omitempty"`
	Type             string `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`
	CanDelete        int32  `protobuf:"varint,11,opt,name=CanDelete,proto3" json:"CanDelete,omitempty"`
}

func (x *ProjectAuth) Reset() {
	*x = ProjectAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAuth) ProtoMessage() {}

func (x *ProjectAuth) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAuth.ProtoReflect.Descriptor instead.
func (*ProjectAuth) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectAuth) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectAuth) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *ProjectAuth) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProjectAuth) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *ProjectAuth) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ProjectAuth) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProjectAuth) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ProjectAuth) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *ProjectAuth) GetIsDefault() int32 {
	if x != nil {
		return x.IsDefault
	}
	return 0
}

func (x *ProjectAuth) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProjectAuth) GetCanDelete() int32 {
	if x != nil {
		return x.CanDelete
	}
	return 0
}

type ProjectNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Node     string                `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Title    string                `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Pnode    string                `protobuf:"bytes,4,opt,name=pnode,proto3" json:"pnode,omitempty"`
	IsLogin  int32                 `protobuf:"varint,5,opt,name=isLogin,proto3" json:"isLogin,omitempty"`
	IsMenu   int32                 `protobuf:"varint,6,opt,name=isMenu,proto3" json:"isMenu,omitempty"`
	IsAuth   int32                 `protobuf:"varint,7,opt,name=isAuth,proto3" json:"isAuth,omitempty"`
	Checked  bool                  `protobuf:"varint,8,opt,name=checked,proto3" json:"checked,omitempty"`
	Key      string                `protobuf:"bytes,9,opt,name=key,proto3" json:"key,omitempty"`
	Children []*ProjectNodeMessage `protobuf:"bytes,10,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ProjectNodeMessage) Reset() {
	*x = ProjectNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectNodeMessage) ProtoMessage() {}

func (x *ProjectNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectNodeMessage.ProtoReflect.Descriptor instead.
func (*ProjectNodeMessage) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectNodeMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectNodeMessage) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *ProjectNodeMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProjectNodeMessage) GetPnode() string {
	if x != nil {
		return x.Pnode
	}
	return ""
}

func (x *ProjectNodeMessage) GetIsLogin() int32 {
	if x != nil {
		return x.IsLogin
	}
	return 0
}

func (x *ProjectNodeMessage) GetIsMenu() int32 {
	if x != nil {
		return x.IsMenu
	}
	return 0
}

func (x *ProjectNodeMessage) GetIsAuth() int32 {
	if x != nil {
		return x.IsAuth
	}
	return 0
}

func (x *ProjectNodeMessage) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *ProjectNodeMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ProjectNodeMessage) GetChildren() []*ProjectNodeMessage {
	if x != nil {
		return x.Children
	}
	return nil
}

type ApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List        []*ProjectNodeMessage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	CheckedList []string              `protobuf:"bytes,2,rep,name=checkedList,proto3" json:"checkedList,omitempty"`
}

func (x *ApplyResponse) Reset() {
	*x = ApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyResponse) ProtoMessage() {}

func (x *ApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyResponse.ProtoReflect.Descriptor instead.
func (*ApplyResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *ApplyResponse) GetList() []*ProjectNodeMessage {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ApplyResponse) GetCheckedList() []string {
	if x != nil {
		return x.CheckedList
	}
	return nil
}

type ListAuthMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ProjectAuth `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListAuthMessage) Reset() {
	*x = ListAuthMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthMessage) ProtoMessage() {}

func (x *ListAuthMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthMessage.ProtoReflect.Descriptor instead.
func (*ListAuthMessage) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAuthMessage) GetList() []*ProjectAuth {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListAuthMessage) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AuthNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AuthNodesResponse) Reset() {
	*x = AuthNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthNodesResponse) ProtoMessage() {}

func (x *AuthNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthNodesResponse.ProtoReflect.Descriptor instead.
func (*AuthNodesResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *AuthNodesResponse) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x49, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x22, 0x9b, 0x02, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x6a,
	0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x27, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x88,
	0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f,
	0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x13, 0x41,
	0x75, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_service_proto_rawDescOnce sync.Once
	file_auth_service_proto_rawDescData = file_auth_service_proto_rawDesc
)

func file_auth_service_proto_rawDescGZIP() []byte {
	file_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_proto_rawDescData)
	})
	return file_auth_service_proto_rawDescData
}

var file_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_service_proto_goTypes = []interface{}{
	(*AuthReqMessage)(nil),     // 0: auth.service.v1.AuthReqMessage
	(*ProjectAuth)(nil),        // 1: auth.service.v1.ProjectAuth
	(*ProjectNodeMessage)(nil), // 2: auth.service.v1.ProjectNodeMessage
	(*ApplyResponse)(nil),      // 3: auth.service.v1.ApplyResponse
	(*ListAuthMessage)(nil),    // 4: auth.service.v1.ListAuthMessage
	(*AuthNodesResponse)(nil),  // 5: auth.service.v1.AuthNodesResponse
}
var file_auth_service_proto_depIdxs = []int32{
	2, // 0: auth.service.v1.ProjectNodeMessage.children:type_name -> auth.service.v1.ProjectNodeMessage
	2, // 1: auth.service.v1.ApplyResponse.list:type_name -> auth.service.v1.ProjectNodeMessage
	1, // 2: auth.service.v1.ListAuthMessage.list:type_name -> auth.service.v1.ProjectAuth
	0, // 3: auth.service.v1.AuthService.AuthList:input_type -> auth.service.v1.AuthReqMessage
	0, // 4: auth.service.v1.AuthService.Apply:input_type -> auth.service.v1.AuthReqMessage
	0, // 5: auth.service.v1.AuthService.AuthNodesByMemberId:input_type -> auth.service.v1.AuthReqMessage
	4, // 6: auth.service.v1.AuthService.AuthList:output_type -> auth.service.v1.ListAuthMessage
	3, // 7: auth.service.v1.AuthService.Apply:output_type -> auth.service.v1.ApplyResponse
	5, // 8: auth.service.v1.AuthService.AuthNodesByMemberId:output_type -> auth.service.v1.AuthNodesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReqMessage); i {
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
		file_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAuth); i {
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
		file_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectNodeMessage); i {
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
		file_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyResponse); i {
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
		file_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthMessage); i {
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
		file_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthNodesResponse); i {
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
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_service_proto_msgTypes,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}
