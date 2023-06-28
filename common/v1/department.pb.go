// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: common/v1/department.proto

package department

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DepartmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`             // 部门名称
	Pid        int32  `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`              // 上级部门
	Remark     string `protobuf:"bytes,4,opt,name=Remark,proto3" json:"Remark,omitempty"`         // 预留备注信息
	CreateTime string `protobuf:"bytes,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"` // 数据新增时间
	UpdateTime string `protobuf:"bytes,6,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"` // 最后一次更新数据时间
}

func (x *DepartmentInfo) Reset() {
	*x = DepartmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentInfo) ProtoMessage() {}

func (x *DepartmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentInfo.ProtoReflect.Descriptor instead.
func (*DepartmentInfo) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{0}
}

func (x *DepartmentInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DepartmentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DepartmentInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *DepartmentInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *DepartmentInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *DepartmentInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`     // v: required
	Pid    string `protobuf:"bytes,2,opt,name=Pid,proto3" json:"Pid,omitempty"`       // v: required
	Remark string `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"` // v: required
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateReq) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *CreateReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{2}
}

type GetOneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` // v: required
}

func (x *GetOneReq) Reset() {
	*x = GetOneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneReq) ProtoMessage() {}

func (x *GetOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneReq.ProtoReflect.Descriptor instead.
func (*GetOneReq) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{3}
}

func (x *GetOneReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Department *DepartmentInfo `protobuf:"bytes,1,opt,name=Department,proto3" json:"Department,omitempty"`
}

func (x *GetOneRes) Reset() {
	*x = GetOneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRes) ProtoMessage() {}

func (x *GetOneRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRes.ProtoReflect.Descriptor instead.
func (*GetOneRes) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{4}
}

func (x *GetOneRes) GetDepartment() *DepartmentInfo {
	if x != nil {
		return x.Department
	}
	return nil
}

type GetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *GetListReq) Reset() {
	*x = GetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListReq) ProtoMessage() {}

func (x *GetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListReq.ProtoReflect.Descriptor instead.
func (*GetListReq) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{5}
}

func (x *GetListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type GetListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Department []*DepartmentInfo `protobuf:"bytes,1,rep,name=Department,proto3" json:"Department,omitempty"`
}

func (x *GetListRes) Reset() {
	*x = GetListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_department_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRes) ProtoMessage() {}

func (x *GetListRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_department_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRes.ProtoReflect.Descriptor instead.
func (*GetListRes) Descriptor() ([]byte, []int) {
	return file_common_v1_department_proto_rawDescGZIP(), []int{6}
}

func (x *GetListRes) GetDepartment() []*DepartmentInfo {
	if x != nil {
		return x.Department
	}
	return nil
}

var File_common_v1_department_proto protoreflect.FileDescriptor

var file_common_v1_department_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x22, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x1b, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xa5, 0x01, 0x0a, 0x0a,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x6a, 0x31, 0x35, 0x37, 0x30, 0x36, 0x39, 0x33, 0x36, 0x35, 0x39, 0x2f, 0x67,
	0x66, 0x63, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_department_proto_rawDescOnce sync.Once
	file_common_v1_department_proto_rawDescData = file_common_v1_department_proto_rawDesc
)

func file_common_v1_department_proto_rawDescGZIP() []byte {
	file_common_v1_department_proto_rawDescOnce.Do(func() {
		file_common_v1_department_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_department_proto_rawDescData)
	})
	return file_common_v1_department_proto_rawDescData
}

var file_common_v1_department_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_v1_department_proto_goTypes = []interface{}{
	(*DepartmentInfo)(nil), // 0: common.DepartmentInfo
	(*CreateReq)(nil),      // 1: common.CreateReq
	(*CreateRes)(nil),      // 2: common.CreateRes
	(*GetOneReq)(nil),      // 3: common.GetOneReq
	(*GetOneRes)(nil),      // 4: common.GetOneRes
	(*GetListReq)(nil),     // 5: common.GetListReq
	(*GetListRes)(nil),     // 6: common.GetListRes
}
var file_common_v1_department_proto_depIdxs = []int32{
	0, // 0: common.GetOneRes.Department:type_name -> common.DepartmentInfo
	0, // 1: common.GetListRes.Department:type_name -> common.DepartmentInfo
	1, // 2: common.Department.Create:input_type -> common.CreateReq
	3, // 3: common.Department.GetOne:input_type -> common.GetOneReq
	5, // 4: common.Department.GetList:input_type -> common.GetListReq
	2, // 5: common.Department.Create:output_type -> common.CreateRes
	4, // 6: common.Department.GetOne:output_type -> common.GetOneRes
	6, // 7: common.Department.GetList:output_type -> common.GetListRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_v1_department_proto_init() }
func file_common_v1_department_proto_init() {
	if File_common_v1_department_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_department_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentInfo); i {
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
		file_common_v1_department_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_common_v1_department_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
		file_common_v1_department_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneReq); i {
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
		file_common_v1_department_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRes); i {
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
		file_common_v1_department_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListReq); i {
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
		file_common_v1_department_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRes); i {
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
			RawDescriptor: file_common_v1_department_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_v1_department_proto_goTypes,
		DependencyIndexes: file_common_v1_department_proto_depIdxs,
		MessageInfos:      file_common_v1_department_proto_msgTypes,
	}.Build()
	File_common_v1_department_proto = out.File
	file_common_v1_department_proto_rawDesc = nil
	file_common_v1_department_proto_goTypes = nil
	file_common_v1_department_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DepartmentClient is the client API for Department service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DepartmentClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error)
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error)
}

type departmentClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentClient(cc grpc.ClientConnInterface) DepartmentClient {
	return &departmentClient{cc}
}

func (c *departmentClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, "/common.Department/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error) {
	out := new(GetOneRes)
	err := c.cc.Invoke(ctx, "/common.Department/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error) {
	out := new(GetListRes)
	err := c.cc.Invoke(ctx, "/common.Department/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServer is the server API for Department service.
type DepartmentServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetOne(context.Context, *GetOneReq) (*GetOneRes, error)
	GetList(context.Context, *GetListReq) (*GetListRes, error)
}

// UnimplementedDepartmentServer can be embedded to have forward compatible implementations.
type UnimplementedDepartmentServer struct {
}

func (*UnimplementedDepartmentServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDepartmentServer) GetOne(context.Context, *GetOneReq) (*GetOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (*UnimplementedDepartmentServer) GetList(context.Context, *GetListReq) (*GetListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterDepartmentServer(s *grpc.Server, srv DepartmentServer) {
	s.RegisterService(&_Department_serviceDesc, srv)
}

func _Department_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Department/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Department/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Department/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Department_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.Department",
	HandlerType: (*DepartmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Department_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Department_GetOne_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Department_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/v1/department.proto",
}
