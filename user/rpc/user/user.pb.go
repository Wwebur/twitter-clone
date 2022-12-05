// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc/user/user.proto

package user

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// FindUserByIDRequest request body for FindUserByID
type FindUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FindUserByIDRequest) Reset() {
	*x = FindUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByIDRequest) ProtoMessage() {}

func (x *FindUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByIDRequest.ProtoReflect.Descriptor instead.
func (*FindUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *FindUserByIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// FindUserByIDResponse response body for FindUserByID
type FindUserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindUserByIDResponse) Reset() {
	*x = FindUserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByIDResponse) ProtoMessage() {}

func (x *FindUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByIDResponse.ProtoReflect.Descriptor instead.
func (*FindUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserByIDResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// FindUserByEmailRequest request body for FindUserByEmail
type FindUserByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindUserByEmailRequest) Reset() {
	*x = FindUserByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailRequest) ProtoMessage() {}

func (x *FindUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*FindUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *FindUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// FindUserByEmailResponse response body for FindUserByEmail
type FindUserByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindUserByEmailResponse) Reset() {
	*x = FindUserByEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailResponse) ProtoMessage() {}

func (x *FindUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*FindUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *FindUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// CreateUserRequest request body for CreateUser
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ScreenName       string `protobuf:"bytes,2,opt,name=screen_name,json=screenName,proto3" json:"screen_name,omitempty"`
	PasswordHash     string `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Email            string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Bio              string `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Location         string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Website          string `protobuf:"bytes,7,opt,name=website,proto3" json:"website,omitempty"`
	ProfileImageUrl  string `protobuf:"bytes,8,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	ProfileBannerUrl string `protobuf:"bytes,9,opt,name=profile_banner_url,json=profileBannerUrl,proto3" json:"profile_banner_url,omitempty"`
	BirthDate        string `protobuf:"bytes,10,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetScreenName() string {
	if x != nil {
		return x.ScreenName
	}
	return ""
}

func (x *CreateUserRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *CreateUserRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateUserRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CreateUserRequest) GetProfileImageUrl() string {
	if x != nil {
		return x.ProfileImageUrl
	}
	return ""
}

func (x *CreateUserRequest) GetProfileBannerUrl() string {
	if x != nil {
		return x.ProfileBannerUrl
	}
	return ""
}

func (x *CreateUserRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

// CreateUserResponse response body for CreateUser
type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// DeleteUserRequest request body for DeleteUser
type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// DeleteUserResponse response body for DeleteUser
type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// User represents the user model
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           int64                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name             string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ScreenName       string               `protobuf:"bytes,3,opt,name=screen_name,json=screenName,proto3" json:"screen_name,omitempty"`
	PasswordHash     string               `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Email            string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Bio              string               `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio,omitempty"`
	Location         string               `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Website          string               `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	ProfileImageUrl  string               `protobuf:"bytes,9,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	ProfileBannerUrl string               `protobuf:"bytes,10,opt,name=profile_banner_url,json=profileBannerUrl,proto3" json:"profile_banner_url,omitempty"`
	BirthDate        *timestamp.Timestamp `protobuf:"bytes,11,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	FollowersCount   int32                `protobuf:"varint,12,opt,name=followers_count,json=followersCount,proto3" json:"followers_count,omitempty"`
	FollowingsCount  int32                `protobuf:"varint,13,opt,name=followings_count,json=followingsCount,proto3" json:"followings_count,omitempty"`
	CreatedAt        *timestamp.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        *timestamp.Timestamp `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rpc_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetScreenName() string {
	if x != nil {
		return x.ScreenName
	}
	return ""
}

func (x *User) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *User) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *User) GetProfileImageUrl() string {
	if x != nil {
		return x.ProfileImageUrl
	}
	return ""
}

func (x *User) GetProfileBannerUrl() string {
	if x != nil {
		return x.ProfileBannerUrl
	}
	return ""
}

func (x *User) GetBirthDate() *timestamp.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *User) GetFollowersCount() int32 {
	if x != nil {
		return x.FollowersCount
	}
	return 0
}

func (x *User) GetFollowingsCount() int32 {
	if x != nil {
		return x.FollowingsCount
	}
	return 0
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_rpc_user_user_proto protoreflect.FileDescriptor

var file_rpc_user_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f,
	0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x6f,
	0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xc4, 0x02, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74,
	0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xb6,
	0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xef, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x32, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74,
	0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f,
	0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x68, 0x6f,
	0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x80, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f,
	0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x68, 0x6f,
	0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x30, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63, 0x2e, 0x74,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x63,
	0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f,
	0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x68, 0x6f, 0x74, 0x70, 0x6f, 0x74, 0x61,
	0x74, 0x6f, 0x63, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x6e,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x72, 0x70, 0x63,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_user_user_proto_rawDescOnce sync.Once
	file_rpc_user_user_proto_rawDescData = file_rpc_user_user_proto_rawDesc
)

func file_rpc_user_user_proto_rawDescGZIP() []byte {
	file_rpc_user_user_proto_rawDescOnce.Do(func() {
		file_rpc_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_user_user_proto_rawDescData)
	})
	return file_rpc_user_user_proto_rawDescData
}

var file_rpc_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_user_user_proto_goTypes = []interface{}{
	(*FindUserByIDRequest)(nil),     // 0: hotpotatoc.twitter_clone.user.FindUserByIDRequest
	(*FindUserByIDResponse)(nil),    // 1: hotpotatoc.twitter_clone.user.FindUserByIDResponse
	(*FindUserByEmailRequest)(nil),  // 2: hotpotatoc.twitter_clone.user.FindUserByEmailRequest
	(*FindUserByEmailResponse)(nil), // 3: hotpotatoc.twitter_clone.user.FindUserByEmailResponse
	(*CreateUserRequest)(nil),       // 4: hotpotatoc.twitter_clone.user.CreateUserRequest
	(*CreateUserResponse)(nil),      // 5: hotpotatoc.twitter_clone.user.CreateUserResponse
	(*DeleteUserRequest)(nil),       // 6: hotpotatoc.twitter_clone.user.DeleteUserRequest
	(*DeleteUserResponse)(nil),      // 7: hotpotatoc.twitter_clone.user.DeleteUserResponse
	(*User)(nil),                    // 8: hotpotatoc.twitter_clone.user.User
	(*timestamp.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_rpc_user_user_proto_depIdxs = []int32{
	8,  // 0: hotpotatoc.twitter_clone.user.FindUserByIDResponse.user:type_name -> hotpotatoc.twitter_clone.user.User
	8,  // 1: hotpotatoc.twitter_clone.user.FindUserByEmailResponse.user:type_name -> hotpotatoc.twitter_clone.user.User
	8,  // 2: hotpotatoc.twitter_clone.user.CreateUserResponse.user:type_name -> hotpotatoc.twitter_clone.user.User
	9,  // 3: hotpotatoc.twitter_clone.user.User.birth_date:type_name -> google.protobuf.Timestamp
	9,  // 4: hotpotatoc.twitter_clone.user.User.created_at:type_name -> google.protobuf.Timestamp
	9,  // 5: hotpotatoc.twitter_clone.user.User.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 6: hotpotatoc.twitter_clone.user.UserService.FindUserByID:input_type -> hotpotatoc.twitter_clone.user.FindUserByIDRequest
	2,  // 7: hotpotatoc.twitter_clone.user.UserService.FindUserByEmail:input_type -> hotpotatoc.twitter_clone.user.FindUserByEmailRequest
	4,  // 8: hotpotatoc.twitter_clone.user.UserService.CreateUser:input_type -> hotpotatoc.twitter_clone.user.CreateUserRequest
	6,  // 9: hotpotatoc.twitter_clone.user.UserService.DeleteUser:input_type -> hotpotatoc.twitter_clone.user.DeleteUserRequest
	1,  // 10: hotpotatoc.twitter_clone.user.UserService.FindUserByID:output_type -> hotpotatoc.twitter_clone.user.FindUserByIDResponse
	3,  // 11: hotpotatoc.twitter_clone.user.UserService.FindUserByEmail:output_type -> hotpotatoc.twitter_clone.user.FindUserByEmailResponse
	5,  // 12: hotpotatoc.twitter_clone.user.UserService.CreateUser:output_type -> hotpotatoc.twitter_clone.user.CreateUserResponse
	7,  // 13: hotpotatoc.twitter_clone.user.UserService.DeleteUser:output_type -> hotpotatoc.twitter_clone.user.DeleteUserResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rpc_user_user_proto_init() }
func file_rpc_user_user_proto_init() {
	if File_rpc_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByIDRequest); i {
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
		file_rpc_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByIDResponse); i {
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
		file_rpc_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailRequest); i {
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
		file_rpc_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailResponse); i {
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
		file_rpc_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_rpc_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_rpc_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_rpc_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_rpc_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_rpc_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_user_user_proto_goTypes,
		DependencyIndexes: file_rpc_user_user_proto_depIdxs,
		MessageInfos:      file_rpc_user_user_proto_msgTypes,
	}.Build()
	File_rpc_user_user_proto = out.File
	file_rpc_user_user_proto_rawDesc = nil
	file_rpc_user_user_proto_goTypes = nil
	file_rpc_user_user_proto_depIdxs = nil
}