// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userManage.proto

package userManage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserListReq struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	ManageId             int64    `protobuf:"varint,3,opt,name=manageId,proto3" json:"manageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListReq) Reset()         { *m = GetUserListReq{} }
func (m *GetUserListReq) String() string { return proto.CompactTextString(m) }
func (*GetUserListReq) ProtoMessage()    {}
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{0}
}

func (m *GetUserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListReq.Unmarshal(m, b)
}
func (m *GetUserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListReq.Marshal(b, m, deterministic)
}
func (m *GetUserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListReq.Merge(m, src)
}
func (m *GetUserListReq) XXX_Size() int {
	return xxx_messageInfo_GetUserListReq.Size(m)
}
func (m *GetUserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListReq proto.InternalMessageInfo

func (m *GetUserListReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetUserListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetUserListReq) GetManageId() int64 {
	if m != nil {
		return m.ManageId
	}
	return 0
}

type UserMesssage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	JobNumber            string   `protobuf:"bytes,6,opt,name=job_number,json=jobNumber,proto3" json:"job_number,omitempty"`
	Permission           int32    `protobuf:"varint,7,opt,name=permission,proto3" json:"permission,omitempty"`
	Gender               int32    `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Deleted              bool     `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMesssage) Reset()         { *m = UserMesssage{} }
func (m *UserMesssage) String() string { return proto.CompactTextString(m) }
func (*UserMesssage) ProtoMessage()    {}
func (*UserMesssage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{1}
}

func (m *UserMesssage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMesssage.Unmarshal(m, b)
}
func (m *UserMesssage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMesssage.Marshal(b, m, deterministic)
}
func (m *UserMesssage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMesssage.Merge(m, src)
}
func (m *UserMesssage) XXX_Size() int {
	return xxx_messageInfo_UserMesssage.Size(m)
}
func (m *UserMesssage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMesssage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMesssage proto.InternalMessageInfo

func (m *UserMesssage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserMesssage) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserMesssage) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *UserMesssage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserMesssage) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserMesssage) GetJobNumber() string {
	if m != nil {
		return m.JobNumber
	}
	return ""
}

func (m *UserMesssage) GetPermission() int32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func (m *UserMesssage) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserMesssage) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type UserListRsp struct {
	UserList             []*UserMesssage `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserListRsp) Reset()         { *m = UserListRsp{} }
func (m *UserListRsp) String() string { return proto.CompactTextString(m) }
func (*UserListRsp) ProtoMessage()    {}
func (*UserListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{2}
}

func (m *UserListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRsp.Unmarshal(m, b)
}
func (m *UserListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRsp.Marshal(b, m, deterministic)
}
func (m *UserListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRsp.Merge(m, src)
}
func (m *UserListRsp) XXX_Size() int {
	return xxx_messageInfo_UserListRsp.Size(m)
}
func (m *UserListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRsp proto.InternalMessageInfo

func (m *UserListRsp) GetUserList() []*UserMesssage {
	if m != nil {
		return m.UserList
	}
	return nil
}

type ChangeUserReq struct {
	ChangeId             int64    `protobuf:"varint,1,opt,name=changeId,proto3" json:"changeId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LoginName            string   `protobuf:"bytes,3,opt,name=loginName,proto3" json:"loginName,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	JobNumber            string   `protobuf:"bytes,5,opt,name=jobNumber,proto3" json:"jobNumber,omitempty"`
	Gender               int32    `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUserReq) Reset()         { *m = ChangeUserReq{} }
func (m *ChangeUserReq) String() string { return proto.CompactTextString(m) }
func (*ChangeUserReq) ProtoMessage()    {}
func (*ChangeUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{3}
}

func (m *ChangeUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUserReq.Unmarshal(m, b)
}
func (m *ChangeUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUserReq.Marshal(b, m, deterministic)
}
func (m *ChangeUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUserReq.Merge(m, src)
}
func (m *ChangeUserReq) XXX_Size() int {
	return xxx_messageInfo_ChangeUserReq.Size(m)
}
func (m *ChangeUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUserReq proto.InternalMessageInfo

func (m *ChangeUserReq) GetChangeId() int64 {
	if m != nil {
		return m.ChangeId
	}
	return 0
}

func (m *ChangeUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangeUserReq) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *ChangeUserReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *ChangeUserReq) GetJobNumber() string {
	if m != nil {
		return m.JobNumber
	}
	return ""
}

func (m *ChangeUserReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type ChangeUserRsp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUserRsp) Reset()         { *m = ChangeUserRsp{} }
func (m *ChangeUserRsp) String() string { return proto.CompactTextString(m) }
func (*ChangeUserRsp) ProtoMessage()    {}
func (*ChangeUserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{4}
}

func (m *ChangeUserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUserRsp.Unmarshal(m, b)
}
func (m *ChangeUserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUserRsp.Marshal(b, m, deterministic)
}
func (m *ChangeUserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUserRsp.Merge(m, src)
}
func (m *ChangeUserRsp) XXX_Size() int {
	return xxx_messageInfo_ChangeUserRsp.Size(m)
}
func (m *ChangeUserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUserRsp proto.InternalMessageInfo

func (m *ChangeUserRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChangeUserRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AddUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	JobNumber            string   `protobuf:"bytes,4,opt,name=jobNumber,proto3" json:"jobNumber,omitempty"`
	Gender               int32    `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserReq) Reset()         { *m = AddUserReq{} }
func (m *AddUserReq) String() string { return proto.CompactTextString(m) }
func (*AddUserReq) ProtoMessage()    {}
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{5}
}

func (m *AddUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserReq.Unmarshal(m, b)
}
func (m *AddUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserReq.Marshal(b, m, deterministic)
}
func (m *AddUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserReq.Merge(m, src)
}
func (m *AddUserReq) XXX_Size() int {
	return xxx_messageInfo_AddUserReq.Size(m)
}
func (m *AddUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserReq proto.InternalMessageInfo

func (m *AddUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddUserReq) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *AddUserReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *AddUserReq) GetJobNumber() string {
	if m != nil {
		return m.JobNumber
	}
	return ""
}

func (m *AddUserReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type AddUserRsp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRsp) Reset()         { *m = AddUserRsp{} }
func (m *AddUserRsp) String() string { return proto.CompactTextString(m) }
func (*AddUserRsp) ProtoMessage()    {}
func (*AddUserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{6}
}

func (m *AddUserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRsp.Unmarshal(m, b)
}
func (m *AddUserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRsp.Marshal(b, m, deterministic)
}
func (m *AddUserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRsp.Merge(m, src)
}
func (m *AddUserRsp) XXX_Size() int {
	return xxx_messageInfo_AddUserRsp.Size(m)
}
func (m *AddUserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRsp proto.InternalMessageInfo

func (m *AddUserRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddUserRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeleteUserReq struct {
	DeleteId             int64    `protobuf:"varint,1,opt,name=deleteId,proto3" json:"deleteId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserReq) Reset()         { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()    {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{7}
}

func (m *DeleteUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserReq.Unmarshal(m, b)
}
func (m *DeleteUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserReq.Marshal(b, m, deterministic)
}
func (m *DeleteUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserReq.Merge(m, src)
}
func (m *DeleteUserReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserReq.Size(m)
}
func (m *DeleteUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserReq proto.InternalMessageInfo

func (m *DeleteUserReq) GetDeleteId() int64 {
	if m != nil {
		return m.DeleteId
	}
	return 0
}

type DeleteUserRsp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRsp) Reset()         { *m = DeleteUserRsp{} }
func (m *DeleteUserRsp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRsp) ProtoMessage()    {}
func (*DeleteUserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22129ec94f3fdfd, []int{8}
}

func (m *DeleteUserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRsp.Unmarshal(m, b)
}
func (m *DeleteUserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRsp.Marshal(b, m, deterministic)
}
func (m *DeleteUserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRsp.Merge(m, src)
}
func (m *DeleteUserRsp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRsp.Size(m)
}
func (m *DeleteUserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRsp proto.InternalMessageInfo

func (m *DeleteUserRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteUserRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserListReq)(nil), "GetUserListReq")
	proto.RegisterType((*UserMesssage)(nil), "UserMesssage")
	proto.RegisterType((*UserListRsp)(nil), "UserListRsp")
	proto.RegisterType((*ChangeUserReq)(nil), "ChangeUserReq")
	proto.RegisterType((*ChangeUserRsp)(nil), "ChangeUserRsp")
	proto.RegisterType((*AddUserReq)(nil), "AddUserReq")
	proto.RegisterType((*AddUserRsp)(nil), "AddUserRsp")
	proto.RegisterType((*DeleteUserReq)(nil), "DeleteUserReq")
	proto.RegisterType((*DeleteUserRsp)(nil), "DeleteUserRsp")
}

func init() { proto.RegisterFile("userManage.proto", fileDescriptor_f22129ec94f3fdfd) }

var fileDescriptor_f22129ec94f3fdfd = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcb, 0x6a, 0xdc, 0x3e,
	0x14, 0xc6, 0xff, 0x1a, 0xcf, 0xcd, 0x67, 0x2e, 0xff, 0x20, 0x4a, 0x31, 0x26, 0x2d, 0xae, 0xa1,
	0x30, 0xa5, 0xe0, 0xc5, 0x74, 0xd3, 0x4d, 0x0b, 0x93, 0x16, 0xca, 0x40, 0x92, 0x82, 0x60, 0x36,
	0xdd, 0x04, 0x4f, 0x75, 0xc6, 0x71, 0x18, 0x5f, 0x6a, 0x79, 0x28, 0x79, 0x91, 0xbe, 0x49, 0x5f,
	0xa6, 0x2f, 0xd2, 0x6d, 0xd1, 0xf1, 0x4d, 0x5e, 0x84, 0x40, 0x76, 0xfe, 0x8e, 0xa4, 0x4f, 0xe7,
	0xfb, 0x49, 0x32, 0x9c, 0x9d, 0x14, 0x16, 0x57, 0x61, 0x1a, 0x46, 0x18, 0xe4, 0x45, 0x56, 0x66,
	0xfe, 0x37, 0x58, 0x7e, 0xc1, 0x72, 0xa7, 0xb0, 0xb8, 0x8c, 0x55, 0x29, 0xf0, 0x07, 0x7f, 0x0e,
	0xe3, 0xec, 0x70, 0x50, 0x58, 0x3a, 0xcc, 0x63, 0xab, 0x91, 0xa8, 0x15, 0x7f, 0x06, 0xa3, 0x63,
	0x9c, 0xc4, 0xa5, 0x33, 0xa0, 0x72, 0x25, 0xb8, 0x0b, 0xd3, 0x84, 0xfc, 0xb6, 0xd2, 0xb1, 0x3c,
	0xb6, 0xb2, 0x44, 0xab, 0xfd, 0xbf, 0x0c, 0xe6, 0xda, 0xf9, 0x0a, 0x95, 0x52, 0x61, 0x84, 0x7c,
	0x09, 0x83, 0x58, 0x92, 0xad, 0x25, 0x06, 0xb1, 0xe4, 0x2f, 0x00, 0x8e, 0x59, 0x14, 0xa7, 0x37,
	0x69, 0x98, 0x20, 0xf9, 0xda, 0xc2, 0xa6, 0xca, 0x75, 0x98, 0x20, 0x3f, 0x03, 0x2b, 0xff, 0x59,
	0xd9, 0xda, 0x42, 0x7f, 0x72, 0x0e, 0x43, 0x9a, 0x3a, 0xa4, 0x12, 0x7d, 0xf3, 0x57, 0x30, 0xcf,
	0x6f, 0xb3, 0x14, 0x6f, 0xd2, 0x53, 0xb2, 0xc7, 0xc2, 0x19, 0xd1, 0xd8, 0x8c, 0x6a, 0xd7, 0x54,
	0xd2, 0xfb, 0xdc, 0x65, 0xfb, 0x66, 0xc2, 0xb8, 0xda, 0xe7, 0x2e, 0xdb, 0xd7, 0xc3, 0x2f, 0x01,
	0x72, 0x2c, 0x92, 0x58, 0xa9, 0x38, 0x4b, 0x9d, 0x09, 0xc5, 0x33, 0x2a, 0x9a, 0x48, 0x84, 0xa9,
	0xc4, 0xc2, 0x99, 0x56, 0x44, 0x2a, 0xc5, 0x1d, 0x98, 0x48, 0x3c, 0x62, 0x89, 0xd2, 0xb1, 0x3d,
	0xb6, 0x9a, 0x8a, 0x46, 0xfa, 0xef, 0x61, 0xd6, 0x22, 0x55, 0x39, 0x7f, 0x03, 0xd3, 0x53, 0x2d,
	0x1d, 0xe6, 0x59, 0xab, 0xd9, 0x7a, 0x11, 0x98, 0x60, 0x44, 0x3b, 0xec, 0xff, 0x66, 0xb0, 0xf8,
	0x74, 0x1b, 0xa6, 0x11, 0xea, 0x09, 0xfa, 0x3c, 0x5c, 0x98, 0x7e, 0xa7, 0xc2, 0xb6, 0x41, 0xd7,
	0xea, 0x96, 0xc7, 0xc0, 0xe0, 0x71, 0x0e, 0x1d, 0xc2, 0x9a, 0x9d, 0xc1, 0xd4, 0x03, 0x93, 0x4c,
	0x0d, 0xb2, 0x07, 0xeb, 0x1c, 0x3a, 0x34, 0x35, 0x4c, 0x83, 0x55, 0xc7, 0x62, 0x6c, 0xb2, 0xf0,
	0x37, 0xbd, 0xb6, 0x55, 0xae, 0xe1, 0x24, 0x48, 0xe9, 0xa8, 0x6b, 0x5b, 0x34, 0x52, 0x5b, 0xe8,
	0xb8, 0x5b, 0x49, 0x6d, 0x5b, 0xa2, 0x56, 0xfe, 0x2f, 0x06, 0xb0, 0x91, 0xb2, 0xc9, 0xdd, 0x64,
	0x63, 0x0f, 0x65, 0x1b, 0x3c, 0x92, 0xcd, 0x7a, 0x24, 0xdb, 0xf0, 0xe1, 0x6c, 0xa3, 0x5e, 0xb6,
	0x8f, 0x5d, 0x5f, 0x4f, 0x0a, 0xf6, 0x16, 0x16, 0x9f, 0xe9, 0x62, 0x18, 0x47, 0x5a, 0xdd, 0x94,
	0xee, 0x48, 0x1b, 0xad, 0x41, 0x1a, 0x93, 0x9f, 0xb2, 0xdf, 0xfa, 0x0f, 0x03, 0xd8, 0xb5, 0x0f,
	0x9d, 0x7f, 0x00, 0xd7, 0x78, 0xe2, 0x17, 0xf7, 0x5f, 0x0f, 0x07, 0x55, 0x6e, 0x52, 0x79, 0x49,
	0x0f, 0xf8, 0xff, 0xa0, 0xff, 0xfe, 0xdd, 0x79, 0x60, 0x5c, 0x5d, 0xff, 0x3f, 0xbe, 0x86, 0xe5,
	0x2e, 0x97, 0x61, 0xd5, 0xd0, 0xc5, 0xfd, 0x56, 0xf2, 0x65, 0xd0, 0xbb, 0xa1, 0x6e, 0x4f, 0xd3,
	0x9a, 0xd7, 0x30, 0xa9, 0x89, 0xf1, 0x59, 0xd0, 0x9d, 0xa9, 0xdb, 0x89, 0xc6, 0xba, 0xcb, 0x5a,
	0x5b, 0xf7, 0x48, 0xb9, 0x3d, 0xad, 0xd7, 0xec, 0xc7, 0xf4, 0xdf, 0x7a, 0xf7, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0xf0, 0x93, 0xbf, 0xcb, 0x04, 0x00, 0x00,
}
