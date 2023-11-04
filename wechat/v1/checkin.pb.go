// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: wechat/v1/checkin.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求数据
type GetUserCheckInDayDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepartId   int32    `protobuf:"varint,1,opt,name=DepartId,proto3" json:"DepartId,omitempty"`    // 部门ID
	WorkNumber []string `protobuf:"bytes,2,rep,name=WorkNumber,proto3" json:"WorkNumber,omitempty"` // 员工工号
	StartTime  int32    `protobuf:"varint,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime    int32    `protobuf:"varint,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *GetUserCheckInDayDataReq) Reset() {
	*x = GetUserCheckInDayDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCheckInDayDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCheckInDayDataReq) ProtoMessage() {}

func (x *GetUserCheckInDayDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCheckInDayDataReq.ProtoReflect.Descriptor instead.
func (*GetUserCheckInDayDataReq) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserCheckInDayDataReq) GetDepartId() int32 {
	if x != nil {
		return x.DepartId
	}
	return 0
}

func (x *GetUserCheckInDayDataReq) GetWorkNumber() []string {
	if x != nil {
		return x.WorkNumber
	}
	return nil
}

func (x *GetUserCheckInDayDataReq) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetUserCheckInDayDataReq) GetEndTime() int32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

// 汇总数据
type SummaryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckinCount    int32  `protobuf:"varint,1,opt,name=CheckinCount,proto3" json:"CheckinCount,omitempty"`       // 当日打卡次数
	RegularWorkSec  int32  `protobuf:"varint,2,opt,name=RegularWorkSec,proto3" json:"RegularWorkSec,omitempty"`   // 当日实际工作时长，单位：秒
	StandardWorkSec int32  `protobuf:"varint,3,opt,name=StandardWorkSec,proto3" json:"StandardWorkSec,omitempty"` // 当日标准工作时长，单位：秒
	EarliestTime    string `protobuf:"bytes,4,opt,name=EarliestTime,proto3" json:"EarliestTime,omitempty"`        // 当日最早打卡时间
	LastestTime     string `protobuf:"bytes,5,opt,name=LastestTime,proto3" json:"LastestTime,omitempty"`          // 当日最晚打卡时间
	RecordType      int32  `protobuf:"varint,6,opt,name=RecordType,proto3" json:"RecordType,omitempty"`
}

func (x *SummaryInfo) Reset() {
	*x = SummaryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryInfo) ProtoMessage() {}

func (x *SummaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryInfo.ProtoReflect.Descriptor instead.
func (*SummaryInfo) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryInfo) GetCheckinCount() int32 {
	if x != nil {
		return x.CheckinCount
	}
	return 0
}

func (x *SummaryInfo) GetRegularWorkSec() int32 {
	if x != nil {
		return x.RegularWorkSec
	}
	return 0
}

func (x *SummaryInfo) GetStandardWorkSec() int32 {
	if x != nil {
		return x.StandardWorkSec
	}
	return 0
}

func (x *SummaryInfo) GetEarliestTime() string {
	if x != nil {
		return x.EarliestTime
	}
	return ""
}

func (x *SummaryInfo) GetLastestTime() string {
	if x != nil {
		return x.LastestTime
	}
	return ""
}

func (x *SummaryInfo) GetRecordType() int32 {
	if x != nil {
		return x.RecordType
	}
	return 0
}

// OtInfo 加班信息
type OtInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtStatus              int32   `protobuf:"varint,1,opt,name=OtStatus,proto3" json:"OtStatus,omitempty"`                           // 状态：0-无加班；1-正常；2-缺时长
	OtDuration            float32 `protobuf:"fixed32,2,opt,name=OtDuration,proto3" json:"OtDuration,omitempty"`                      // 加班时长
	ExceptionDuration     []int32 `protobuf:"varint,3,rep,packed,name=ExceptionDuration,proto3" json:"ExceptionDuration,omitempty"`  // ot_status为2下，加班不足的时长
	WorkdayOverAsVacation int32   `protobuf:"varint,4,opt,name=WorkdayOverAsVacation,proto3" json:"WorkdayOverAsVacation,omitempty"` // 工作日加班记为调休，单位秒
	WorkdayOverAsMoney    int32   `protobuf:"varint,5,opt,name=WorkdayOverAsMoney,proto3" json:"WorkdayOverAsMoney,omitempty"`       // 工作日加班记为加班费，单位秒
}

func (x *OtInfo) Reset() {
	*x = OtInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtInfo) ProtoMessage() {}

func (x *OtInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtInfo.ProtoReflect.Descriptor instead.
func (*OtInfo) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{2}
}

func (x *OtInfo) GetOtStatus() int32 {
	if x != nil {
		return x.OtStatus
	}
	return 0
}

func (x *OtInfo) GetOtDuration() float32 {
	if x != nil {
		return x.OtDuration
	}
	return 0
}

func (x *OtInfo) GetExceptionDuration() []int32 {
	if x != nil {
		return x.ExceptionDuration
	}
	return nil
}

func (x *OtInfo) GetWorkdayOverAsVacation() int32 {
	if x != nil {
		return x.WorkdayOverAsVacation
	}
	return 0
}

func (x *OtInfo) GetWorkdayOverAsMoney() int32 {
	if x != nil {
		return x.WorkdayOverAsMoney
	}
	return 0
}

// RuleInfo 打卡人员所属规则信息
type RuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupid      int32          `protobuf:"varint,1,opt,name=Groupid,proto3" json:"Groupid,omitempty"`
	Groupname    string         `protobuf:"bytes,2,opt,name=Groupname,proto3" json:"Groupname,omitempty"`
	Scheduleid   int32          `protobuf:"varint,3,opt,name=Scheduleid,proto3" json:"Scheduleid,omitempty"`
	Schedulename string         `protobuf:"bytes,4,opt,name=Schedulename,proto3" json:"Schedulename,omitempty"`
	Checkintime  []*CheckinTime `protobuf:"bytes,5,rep,name=Checkintime,proto3" json:"Checkintime,omitempty"`
}

func (x *RuleInfo) Reset() {
	*x = RuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleInfo) ProtoMessage() {}

func (x *RuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleInfo.ProtoReflect.Descriptor instead.
func (*RuleInfo) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{3}
}

func (x *RuleInfo) GetGroupid() int32 {
	if x != nil {
		return x.Groupid
	}
	return 0
}

func (x *RuleInfo) GetGroupname() string {
	if x != nil {
		return x.Groupname
	}
	return ""
}

func (x *RuleInfo) GetScheduleid() int32 {
	if x != nil {
		return x.Scheduleid
	}
	return 0
}

func (x *RuleInfo) GetSchedulename() string {
	if x != nil {
		return x.Schedulename
	}
	return ""
}

func (x *RuleInfo) GetCheckintime() []*CheckinTime {
	if x != nil {
		return x.Checkintime
	}
	return nil
}

type CheckinTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkSec    int32 `protobuf:"varint,1,opt,name=WorkSec,proto3" json:"WorkSec,omitempty"`       // 上班时间，为距离0点的时间差
	OffWorkSec int32 `protobuf:"varint,2,opt,name=OffWorkSec,proto3" json:"OffWorkSec,omitempty"` // 下班时间，为距离0点的时间差
}

func (x *CheckinTime) Reset() {
	*x = CheckinTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinTime) ProtoMessage() {}

func (x *CheckinTime) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinTime.ProtoReflect.Descriptor instead.
func (*CheckinTime) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{4}
}

func (x *CheckinTime) GetWorkSec() int32 {
	if x != nil {
		return x.WorkSec
	}
	return 0
}

func (x *CheckinTime) GetOffWorkSec() int32 {
	if x != nil {
		return x.OffWorkSec
	}
	return 0
}

type BaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        string    `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	RecordType  int32     `protobuf:"varint,2,opt,name=RecordType,proto3" json:"RecordType,omitempty"` // 记录类型：1-固定上下班；2-外出（此报表中不会出现外出打卡数据）；3-按班次上下班；4-自由签到；5-加班；7-无规则
	Name        string    `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	NameEx      string    `protobuf:"bytes,4,opt,name=NameEx,proto3" json:"NameEx,omitempty"`
	DepartsName string    `protobuf:"bytes,5,opt,name=DepartsName,proto3" json:"DepartsName,omitempty"`
	Acctid      string    `protobuf:"bytes,6,opt,name=Acctid,proto3" json:"Acctid,omitempty"`     // 打卡人员账号，即userid
	RuleInfo    *RuleInfo `protobuf:"bytes,7,opt,name=RuleInfo,proto3" json:"RuleInfo,omitempty"` // 打卡规则
	DayType     int32     `protobuf:"varint,8,opt,name=DayType,proto3" json:"DayType,omitempty"`  // 日报类型：0-工作日日报；1-休息日日报
}

func (x *BaseInfo) Reset() {
	*x = BaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseInfo) ProtoMessage() {}

func (x *BaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseInfo.ProtoReflect.Descriptor instead.
func (*BaseInfo) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{5}
}

func (x *BaseInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BaseInfo) GetRecordType() int32 {
	if x != nil {
		return x.RecordType
	}
	return 0
}

func (x *BaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseInfo) GetNameEx() string {
	if x != nil {
		return x.NameEx
	}
	return ""
}

func (x *BaseInfo) GetDepartsName() string {
	if x != nil {
		return x.DepartsName
	}
	return ""
}

func (x *BaseInfo) GetAcctid() string {
	if x != nil {
		return x.Acctid
	}
	return ""
}

func (x *BaseInfo) GetRuleInfo() *RuleInfo {
	if x != nil {
		return x.RuleInfo
	}
	return nil
}

func (x *BaseInfo) GetDayType() int32 {
	if x != nil {
		return x.DayType
	}
	return 0
}

type GetUserCheckInDayDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*GetUserCheckInDayData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetUserCheckInDayDataRes) Reset() {
	*x = GetUserCheckInDayDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCheckInDayDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCheckInDayDataRes) ProtoMessage() {}

func (x *GetUserCheckInDayDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCheckInDayDataRes.ProtoReflect.Descriptor instead.
func (*GetUserCheckInDayDataRes) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserCheckInDayDataRes) GetData() []*GetUserCheckInDayData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetUserCheckInDayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummaryInfo *SummaryInfo `protobuf:"bytes,1,opt,name=SummaryInfo,proto3" json:"SummaryInfo,omitempty"`
	OtInfo      *OtInfo      `protobuf:"bytes,2,opt,name=OtInfo,proto3" json:"OtInfo,omitempty"`
	BaseInfo    *BaseInfo    `protobuf:"bytes,3,opt,name=BaseInfo,proto3" json:"BaseInfo,omitempty"`
}

func (x *GetUserCheckInDayData) Reset() {
	*x = GetUserCheckInDayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_v1_checkin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCheckInDayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCheckInDayData) ProtoMessage() {}

func (x *GetUserCheckInDayData) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_v1_checkin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCheckInDayData.ProtoReflect.Descriptor instead.
func (*GetUserCheckInDayData) Descriptor() ([]byte, []int) {
	return file_wechat_v1_checkin_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserCheckInDayData) GetSummaryInfo() *SummaryInfo {
	if x != nil {
		return x.SummaryInfo
	}
	return nil
}

func (x *GetUserCheckInDayData) GetOtInfo() *OtInfo {
	if x != nil {
		return x.OtInfo
	}
	return nil
}

func (x *GetUserCheckInDayData) GetBaseInfo() *BaseInfo {
	if x != nil {
		return x.BaseInfo
	}
	return nil
}

var File_wechat_v1_checkin_proto protoreflect.FileDescriptor

var file_wechat_v1_checkin_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x6f,
	0x72, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x57, 0x6f, 0x72, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x12, 0x28,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x61, 0x72, 0x6c,
	0x69, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x45, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd8,
	0x01, 0x0a, 0x06, 0x4f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x4f, 0x74, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x11, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x64, 0x61, 0x79, 0x4f, 0x76,
	0x65, 0x72, 0x41, 0x73, 0x56, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x64, 0x61, 0x79, 0x4f, 0x76, 0x65, 0x72, 0x41,
	0x73, 0x56, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x57, 0x6f, 0x72,
	0x6b, 0x64, 0x61, 0x79, 0x4f, 0x76, 0x65, 0x72, 0x41, 0x73, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x64, 0x61, 0x79, 0x4f, 0x76,
	0x65, 0x72, 0x41, 0x73, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0xbd, 0x01, 0x0a, 0x08, 0x52, 0x75,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x53,
	0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x66, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4f, 0x66, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x53,
	0x65, 0x63, 0x22, 0xec, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x61, 0x6d, 0x65, 0x45,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x74, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x63, 0x63, 0x74, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x52, 0x75, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x52,
	0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x61, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x44, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x4d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0xa4, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0b, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x26, 0x0a, 0x06, 0x4f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x4f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x08, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x6e, 0x0a, 0x0d, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6a, 0x31, 0x35, 0x37, 0x30, 0x36, 0x39, 0x33, 0x36,
	0x35, 0x39, 0x2f, 0x67, 0x66, 0x63, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x77,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechat_v1_checkin_proto_rawDescOnce sync.Once
	file_wechat_v1_checkin_proto_rawDescData = file_wechat_v1_checkin_proto_rawDesc
)

func file_wechat_v1_checkin_proto_rawDescGZIP() []byte {
	file_wechat_v1_checkin_proto_rawDescOnce.Do(func() {
		file_wechat_v1_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechat_v1_checkin_proto_rawDescData)
	})
	return file_wechat_v1_checkin_proto_rawDescData
}

var file_wechat_v1_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wechat_v1_checkin_proto_goTypes = []interface{}{
	(*GetUserCheckInDayDataReq)(nil), // 0: common.GetUserCheckInDayDataReq
	(*SummaryInfo)(nil),              // 1: common.SummaryInfo
	(*OtInfo)(nil),                   // 2: common.OtInfo
	(*RuleInfo)(nil),                 // 3: common.RuleInfo
	(*CheckinTime)(nil),              // 4: common.CheckinTime
	(*BaseInfo)(nil),                 // 5: common.BaseInfo
	(*GetUserCheckInDayDataRes)(nil), // 6: common.GetUserCheckInDayDataRes
	(*GetUserCheckInDayData)(nil),    // 7: common.GetUserCheckInDayData
}
var file_wechat_v1_checkin_proto_depIdxs = []int32{
	4, // 0: common.RuleInfo.Checkintime:type_name -> common.CheckinTime
	3, // 1: common.BaseInfo.RuleInfo:type_name -> common.RuleInfo
	7, // 2: common.GetUserCheckInDayDataRes.Data:type_name -> common.GetUserCheckInDayData
	1, // 3: common.GetUserCheckInDayData.SummaryInfo:type_name -> common.SummaryInfo
	2, // 4: common.GetUserCheckInDayData.OtInfo:type_name -> common.OtInfo
	5, // 5: common.GetUserCheckInDayData.BaseInfo:type_name -> common.BaseInfo
	0, // 6: common.WechatCheckIn.GetUserCheckInDayData:input_type -> common.GetUserCheckInDayDataReq
	6, // 7: common.WechatCheckIn.GetUserCheckInDayData:output_type -> common.GetUserCheckInDayDataRes
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wechat_v1_checkin_proto_init() }
func file_wechat_v1_checkin_proto_init() {
	if File_wechat_v1_checkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechat_v1_checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCheckInDayDataReq); i {
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
		file_wechat_v1_checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryInfo); i {
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
		file_wechat_v1_checkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtInfo); i {
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
		file_wechat_v1_checkin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleInfo); i {
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
		file_wechat_v1_checkin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinTime); i {
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
		file_wechat_v1_checkin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseInfo); i {
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
		file_wechat_v1_checkin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCheckInDayDataRes); i {
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
		file_wechat_v1_checkin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCheckInDayData); i {
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
			RawDescriptor: file_wechat_v1_checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wechat_v1_checkin_proto_goTypes,
		DependencyIndexes: file_wechat_v1_checkin_proto_depIdxs,
		MessageInfos:      file_wechat_v1_checkin_proto_msgTypes,
	}.Build()
	File_wechat_v1_checkin_proto = out.File
	file_wechat_v1_checkin_proto_rawDesc = nil
	file_wechat_v1_checkin_proto_goTypes = nil
	file_wechat_v1_checkin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WechatCheckInClient is the client API for WechatCheckIn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WechatCheckInClient interface {
	GetUserCheckInDayData(ctx context.Context, in *GetUserCheckInDayDataReq, opts ...grpc.CallOption) (*GetUserCheckInDayDataRes, error)
}

type wechatCheckInClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatCheckInClient(cc grpc.ClientConnInterface) WechatCheckInClient {
	return &wechatCheckInClient{cc}
}

func (c *wechatCheckInClient) GetUserCheckInDayData(ctx context.Context, in *GetUserCheckInDayDataReq, opts ...grpc.CallOption) (*GetUserCheckInDayDataRes, error) {
	out := new(GetUserCheckInDayDataRes)
	err := c.cc.Invoke(ctx, "/common.WechatCheckIn/GetUserCheckInDayData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatCheckInServer is the server API for WechatCheckIn service.
type WechatCheckInServer interface {
	GetUserCheckInDayData(context.Context, *GetUserCheckInDayDataReq) (*GetUserCheckInDayDataRes, error)
}

// UnimplementedWechatCheckInServer can be embedded to have forward compatible implementations.
type UnimplementedWechatCheckInServer struct {
}

func (*UnimplementedWechatCheckInServer) GetUserCheckInDayData(context.Context, *GetUserCheckInDayDataReq) (*GetUserCheckInDayDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCheckInDayData not implemented")
}

func RegisterWechatCheckInServer(s *grpc.Server, srv WechatCheckInServer) {
	s.RegisterService(&_WechatCheckIn_serviceDesc, srv)
}

func _WechatCheckIn_GetUserCheckInDayData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCheckInDayDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatCheckInServer).GetUserCheckInDayData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.WechatCheckIn/GetUserCheckInDayData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatCheckInServer).GetUserCheckInDayData(ctx, req.(*GetUserCheckInDayDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WechatCheckIn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.WechatCheckIn",
	HandlerType: (*WechatCheckInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCheckInDayData",
			Handler:    _WechatCheckIn_GetUserCheckInDayData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wechat/v1/checkin.proto",
}
