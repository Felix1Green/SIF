// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/clients/profile.proto

package profile

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

type Errors int32

const (
	Errors_ProfileAlreadyExists   Errors = 0
	Errors_ProfileNotFound        Errors = 1
	Errors_ProfileDataNotProvided Errors = 2
	Errors_InternalServiceError   Errors = 3
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "ProfileAlreadyExists",
		1: "ProfileNotFound",
		2: "ProfileDataNotProvided",
		3: "InternalServiceError",
	}
	Errors_value = map[string]int32{
		"ProfileAlreadyExists":   0,
		"ProfileNotFound":        1,
		"ProfileDataNotProvided": 2,
		"InternalServiceError":   3,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_clients_profile_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_proto_clients_profile_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{0}
}

type GetProfileByUserIDIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetProfileByUserIDIn) Reset() {
	*x = GetProfileByUserIDIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileByUserIDIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByUserIDIn) ProtoMessage() {}

func (x *GetProfileByUserIDIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByUserIDIn.ProtoReflect.Descriptor instead.
func (*GetProfileByUserIDIn) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetProfileByUserIDIn) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName    *string `protobuf:"bytes,2,opt,name=UserName,proto3,oneof" json:"UserName,omitempty"`
	UserMail    string  `protobuf:"bytes,3,opt,name=UserMail,proto3" json:"UserMail,omitempty"`
	UserSurname *string `protobuf:"bytes,4,opt,name=UserSurname,proto3,oneof" json:"UserSurname,omitempty"`
	UserRole    *string `protobuf:"bytes,5,opt,name=UserRole,proto3,oneof" json:"UserRole,omitempty"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileData) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ProfileData) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *ProfileData) GetUserMail() string {
	if x != nil {
		return x.UserMail
	}
	return ""
}

func (x *ProfileData) GetUserSurname() string {
	if x != nil && x.UserSurname != nil {
		return *x.UserSurname
	}
	return ""
}

func (x *ProfileData) GetUserRole() string {
	if x != nil && x.UserRole != nil {
		return *x.UserRole
	}
	return ""
}

type GetProfileByUserIDOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ProfileData `protobuf:"bytes,1,opt,name=profile,proto3,oneof" json:"profile,omitempty"`
	Error   *Errors      `protobuf:"varint,2,opt,name=Error,proto3,enum=ProfileService.Errors,oneof" json:"Error,omitempty"`
	Success bool         `protobuf:"varint,3,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *GetProfileByUserIDOut) Reset() {
	*x = GetProfileByUserIDOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileByUserIDOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByUserIDOut) ProtoMessage() {}

func (x *GetProfileByUserIDOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByUserIDOut.ProtoReflect.Descriptor instead.
func (*GetProfileByUserIDOut) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{2}
}

func (x *GetProfileByUserIDOut) GetProfile() *ProfileData {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *GetProfileByUserIDOut) GetError() Errors {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Errors_ProfileAlreadyExists
}

func (x *GetProfileByUserIDOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateProfileIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ProfileData `protobuf:"bytes,1,opt,name=Profile,proto3" json:"Profile,omitempty"`
}

func (x *CreateProfileIn) Reset() {
	*x = CreateProfileIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileIn) ProtoMessage() {}

func (x *CreateProfileIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileIn.ProtoReflect.Descriptor instead.
func (*CreateProfileIn) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProfileIn) GetProfile() *ProfileData {
	if x != nil {
		return x.Profile
	}
	return nil
}

type CreateProfileOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ProfileData `protobuf:"bytes,1,opt,name=Profile,proto3,oneof" json:"Profile,omitempty"`
	Error   *Errors      `protobuf:"varint,2,opt,name=Error,proto3,enum=ProfileService.Errors,oneof" json:"Error,omitempty"`
	Success bool         `protobuf:"varint,3,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *CreateProfileOut) Reset() {
	*x = CreateProfileOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileOut) ProtoMessage() {}

func (x *CreateProfileOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileOut.ProtoReflect.Descriptor instead.
func (*CreateProfileOut) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProfileOut) GetProfile() *ProfileData {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *CreateProfileOut) GetError() Errors {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Errors_ProfileAlreadyExists
}

func (x *CreateProfileOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAllProfilesIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllProfilesIn) Reset() {
	*x = GetAllProfilesIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProfilesIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProfilesIn) ProtoMessage() {}

func (x *GetAllProfilesIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProfilesIn.ProtoReflect.Descriptor instead.
func (*GetAllProfilesIn) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{5}
}

type GetAllProfilesOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*ProfileData `protobuf:"bytes,1,rep,name=Profiles,proto3" json:"Profiles,omitempty"`
	Success  bool           `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	Error    *Errors        `protobuf:"varint,3,opt,name=Error,proto3,enum=ProfileService.Errors,oneof" json:"Error,omitempty"`
}

func (x *GetAllProfilesOut) Reset() {
	*x = GetAllProfilesOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProfilesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProfilesOut) ProtoMessage() {}

func (x *GetAllProfilesOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProfilesOut.ProtoReflect.Descriptor instead.
func (*GetAllProfilesOut) Descriptor() ([]byte, []int) {
	return file_proto_clients_profile_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllProfilesOut) GetProfiles() []*ProfileData {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *GetAllProfilesOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetAllProfilesOut) GetError() Errors {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Errors_ProfileAlreadyExists
}

var File_proto_clients_profile_proto protoreflect.FileDescriptor

var file_proto_clients_profile_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2e, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xd4, 0x01,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x12, 0x3a,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x48, 0x01, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x48, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x12, 0x35, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x48,
	0x01, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x12, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x22,
	0xa3, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x48,
	0x00, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x6d, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x03, 0x32, 0x97, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x49, 0x6e, 0x1a, 0x25, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x4f, 0x75, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x1a, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x1a, 0x21, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x42, 0x0a,
	0x5a, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_clients_profile_proto_rawDescOnce sync.Once
	file_proto_clients_profile_proto_rawDescData = file_proto_clients_profile_proto_rawDesc
)

func file_proto_clients_profile_proto_rawDescGZIP() []byte {
	file_proto_clients_profile_proto_rawDescOnce.Do(func() {
		file_proto_clients_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clients_profile_proto_rawDescData)
	})
	return file_proto_clients_profile_proto_rawDescData
}

var file_proto_clients_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_clients_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_clients_profile_proto_goTypes = []interface{}{
	(Errors)(0),                   // 0: ProfileService.Errors
	(*GetProfileByUserIDIn)(nil),  // 1: ProfileService.GetProfileByUserIDIn
	(*ProfileData)(nil),           // 2: ProfileService.ProfileData
	(*GetProfileByUserIDOut)(nil), // 3: ProfileService.GetProfileByUserIDOut
	(*CreateProfileIn)(nil),       // 4: ProfileService.CreateProfileIn
	(*CreateProfileOut)(nil),      // 5: ProfileService.CreateProfileOut
	(*GetAllProfilesIn)(nil),      // 6: ProfileService.GetAllProfilesIn
	(*GetAllProfilesOut)(nil),     // 7: ProfileService.GetAllProfilesOut
}
var file_proto_clients_profile_proto_depIdxs = []int32{
	2,  // 0: ProfileService.GetProfileByUserIDOut.profile:type_name -> ProfileService.ProfileData
	0,  // 1: ProfileService.GetProfileByUserIDOut.Error:type_name -> ProfileService.Errors
	2,  // 2: ProfileService.CreateProfileIn.Profile:type_name -> ProfileService.ProfileData
	2,  // 3: ProfileService.CreateProfileOut.Profile:type_name -> ProfileService.ProfileData
	0,  // 4: ProfileService.CreateProfileOut.Error:type_name -> ProfileService.Errors
	2,  // 5: ProfileService.GetAllProfilesOut.Profiles:type_name -> ProfileService.ProfileData
	0,  // 6: ProfileService.GetAllProfilesOut.Error:type_name -> ProfileService.Errors
	1,  // 7: ProfileService.Profile.GetProfileByUserID:input_type -> ProfileService.GetProfileByUserIDIn
	4,  // 8: ProfileService.Profile.CreateProfile:input_type -> ProfileService.CreateProfileIn
	6,  // 9: ProfileService.Profile.GetAllProfiles:input_type -> ProfileService.GetAllProfilesIn
	3,  // 10: ProfileService.Profile.GetProfileByUserID:output_type -> ProfileService.GetProfileByUserIDOut
	5,  // 11: ProfileService.Profile.CreateProfile:output_type -> ProfileService.CreateProfileOut
	7,  // 12: ProfileService.Profile.GetAllProfiles:output_type -> ProfileService.GetAllProfilesOut
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_clients_profile_proto_init() }
func file_proto_clients_profile_proto_init() {
	if File_proto_clients_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clients_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileByUserIDIn); i {
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
		file_proto_clients_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_proto_clients_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileByUserIDOut); i {
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
		file_proto_clients_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileIn); i {
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
		file_proto_clients_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileOut); i {
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
		file_proto_clients_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProfilesIn); i {
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
		file_proto_clients_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProfilesOut); i {
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
	file_proto_clients_profile_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_proto_clients_profile_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_clients_profile_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_proto_clients_profile_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clients_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clients_profile_proto_goTypes,
		DependencyIndexes: file_proto_clients_profile_proto_depIdxs,
		EnumInfos:         file_proto_clients_profile_proto_enumTypes,
		MessageInfos:      file_proto_clients_profile_proto_msgTypes,
	}.Build()
	File_proto_clients_profile_proto = out.File
	file_proto_clients_profile_proto_rawDesc = nil
	file_proto_clients_profile_proto_goTypes = nil
	file_proto_clients_profile_proto_depIdxs = nil
}
