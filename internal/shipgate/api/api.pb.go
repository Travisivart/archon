// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.5
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShipList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ships []*ShipList_Ship `protobuf:"bytes,1,rep,name=ships,proto3" json:"ships,omitempty"`
}

func (x *ShipList) Reset() {
	*x = ShipList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipList) ProtoMessage() {}

func (x *ShipList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipList.ProtoReflect.Descriptor instead.
func (*ShipList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ShipList) GetShips() []*ShipList_Ship {
	if x != nil {
		return x.Ships
	}
	return nil
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port       string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	MaxPlayers int32  `protobuf:"varint,4,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegistrationRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegistrationRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *RegistrationRequest) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

type AccountAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AccountAuthRequest) Reset() {
	*x = AccountAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAuthRequest) ProtoMessage() {}

func (x *AccountAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAuthRequest.ProtoReflect.Descriptor instead.
func (*AccountAuthRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AccountAuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AccountAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username         string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email            string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	RegistrationDate string `protobuf:"bytes,4,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"`
	Guildcard        int64  `protobuf:"varint,5,opt,name=guildcard,proto3" json:"guildcard,omitempty"`
	GM               bool   `protobuf:"varint,6,opt,name=GM,proto3" json:"GM,omitempty"`
	Banned           bool   `protobuf:"varint,7,opt,name=banned,proto3" json:"banned,omitempty"`
	Active           bool   `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
	TeamId           int64  `protobuf:"varint,9,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PriviledgeLevel  []byte `protobuf:"bytes,10,opt,name=priviledge_level,json=priviledgeLevel,proto3" json:"priviledge_level,omitempty"`
}

func (x *AccountAuthResponse) Reset() {
	*x = AccountAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAuthResponse) ProtoMessage() {}

func (x *AccountAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAuthResponse.ProtoReflect.Descriptor instead.
func (*AccountAuthResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AccountAuthResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountAuthResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountAuthResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountAuthResponse) GetRegistrationDate() string {
	if x != nil {
		return x.RegistrationDate
	}
	return ""
}

func (x *AccountAuthResponse) GetGuildcard() int64 {
	if x != nil {
		return x.Guildcard
	}
	return 0
}

func (x *AccountAuthResponse) GetGM() bool {
	if x != nil {
		return x.GM
	}
	return false
}

func (x *AccountAuthResponse) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

func (x *AccountAuthResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AccountAuthResponse) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *AccountAuthResponse) GetPriviledgeLevel() []byte {
	if x != nil {
		return x.PriviledgeLevel
	}
	return nil
}

type ShipList_Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ip          string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port        string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	PlayerCount int32  `protobuf:"varint,5,opt,name=playerCount,proto3" json:"playerCount,omitempty"`
}

func (x *ShipList_Ship) Reset() {
	*x = ShipList_Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipList_Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipList_Ship) ProtoMessage() {}

func (x *ShipList_Ship) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipList_Ship.ProtoReflect.Descriptor instead.
func (*ShipList_Ship) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ShipList_Ship) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShipList_Ship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipList_Ship) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ShipList_Ship) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ShipList_Ship) GetPlayerCount() int32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01,
	0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x05, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x1a, 0x70, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x77, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22,
	0x30, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xa6, 0x02, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x47, 0x4d, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x47, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0xd6, 0x01, 0x0a, 0x0f, 0x53,
	0x68, 0x69, 0x70, 0x67, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x68, 0x69, 0x70, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x13, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(*ShipList)(nil),            // 0: api.ShipList
	(*RegistrationRequest)(nil), // 1: api.RegistrationRequest
	(*AccountAuthRequest)(nil),  // 2: api.AccountAuthRequest
	(*AccountAuthResponse)(nil), // 3: api.AccountAuthResponse
	(*ShipList_Ship)(nil),       // 4: api.ShipList.Ship
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	4, // 0: api.ShipList.ships:type_name -> api.ShipList.Ship
	5, // 1: api.ShipgateService.GetActiveShips:input_type -> google.protobuf.Empty
	1, // 2: api.ShipgateService.RegisterShip:input_type -> api.RegistrationRequest
	2, // 3: api.ShipgateService.AuthenticateAccount:input_type -> api.AccountAuthRequest
	0, // 4: api.ShipgateService.GetActiveShips:output_type -> api.ShipList
	5, // 5: api.ShipgateService.RegisterShip:output_type -> google.protobuf.Empty
	3, // 6: api.ShipgateService.AuthenticateAccount:output_type -> api.AccountAuthResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipList); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAuthRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAuthResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipList_Ship); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
