// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/volunteer.proto

package volunteer

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

type GetVolunteerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVolunteerRequest) Reset() {
	*x = GetVolunteerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVolunteerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVolunteerRequest) ProtoMessage() {}

func (x *GetVolunteerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVolunteerRequest.ProtoReflect.Descriptor instead.
func (*GetVolunteerRequest) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{0}
}

func (x *GetVolunteerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Volunteer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *int32   `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname         string   `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Email           string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	TelephoneNumber string   `protobuf:"bytes,5,opt,name=telephone_number,json=telephoneNumber,proto3" json:"telephone_number,omitempty"`
	Photo           *string  `protobuf:"bytes,6,opt,name=photo,proto3,oneof" json:"photo,omitempty"`
	Tags            []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Voivodeship     string   `protobuf:"bytes,8,opt,name=voivodeship,proto3" json:"voivodeship,omitempty"`
}

func (x *Volunteer) Reset() {
	*x = Volunteer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volunteer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volunteer) ProtoMessage() {}

func (x *Volunteer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volunteer.ProtoReflect.Descriptor instead.
func (*Volunteer) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{1}
}

func (x *Volunteer) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Volunteer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Volunteer) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Volunteer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Volunteer) GetTelephoneNumber() string {
	if x != nil {
		return x.TelephoneNumber
	}
	return ""
}

func (x *Volunteer) GetPhoto() string {
	if x != nil && x.Photo != nil {
		return *x.Photo
	}
	return ""
}

func (x *Volunteer) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Volunteer) GetVoivodeship() string {
	if x != nil {
		return x.Voivodeship
	}
	return ""
}

type GetVolunteerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volunteer *Volunteer `protobuf:"bytes,1,opt,name=volunteer,proto3" json:"volunteer,omitempty"`
}

func (x *GetVolunteerResponse) Reset() {
	*x = GetVolunteerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVolunteerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVolunteerResponse) ProtoMessage() {}

func (x *GetVolunteerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVolunteerResponse.ProtoReflect.Descriptor instead.
func (*GetVolunteerResponse) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{2}
}

func (x *GetVolunteerResponse) GetVolunteer() *Volunteer {
	if x != nil {
		return x.Volunteer
	}
	return nil
}

type ListVolunteersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags        []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	Voivodeship string   `protobuf:"bytes,2,opt,name=voivodeship,proto3" json:"voivodeship,omitempty"`
}

func (x *ListVolunteersRequest) Reset() {
	*x = ListVolunteersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolunteersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolunteersRequest) ProtoMessage() {}

func (x *ListVolunteersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolunteersRequest.ProtoReflect.Descriptor instead.
func (*ListVolunteersRequest) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{3}
}

func (x *ListVolunteersRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListVolunteersRequest) GetVoivodeship() string {
	if x != nil {
		return x.Voivodeship
	}
	return ""
}

type ListVolunteersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volunteers []*Volunteer `protobuf:"bytes,1,rep,name=volunteers,proto3" json:"volunteers,omitempty"`
}

func (x *ListVolunteersResponse) Reset() {
	*x = ListVolunteersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolunteersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolunteersResponse) ProtoMessage() {}

func (x *ListVolunteersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolunteersResponse.ProtoReflect.Descriptor instead.
func (*ListVolunteersResponse) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{4}
}

func (x *ListVolunteersResponse) GetVolunteers() []*Volunteer {
	if x != nil {
		return x.Volunteers
	}
	return nil
}

type CreateVolunteerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volunteer *Volunteer `protobuf:"bytes,1,opt,name=volunteer,proto3" json:"volunteer,omitempty"`
}

func (x *CreateVolunteerRequest) Reset() {
	*x = CreateVolunteerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolunteerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolunteerRequest) ProtoMessage() {}

func (x *CreateVolunteerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolunteerRequest.ProtoReflect.Descriptor instead.
func (*CreateVolunteerRequest) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{5}
}

func (x *CreateVolunteerRequest) GetVolunteer() *Volunteer {
	if x != nil {
		return x.Volunteer
	}
	return nil
}

type CreateVolunteerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volunteer *Volunteer `protobuf:"bytes,1,opt,name=volunteer,proto3" json:"volunteer,omitempty"`
}

func (x *CreateVolunteerResponse) Reset() {
	*x = CreateVolunteerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volunteer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolunteerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolunteerResponse) ProtoMessage() {}

func (x *CreateVolunteerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volunteer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolunteerResponse.ProtoReflect.Descriptor instead.
func (*CreateVolunteerResponse) Descriptor() ([]byte, []int) {
	return file_proto_volunteer_proto_rawDescGZIP(), []int{6}
}

func (x *CreateVolunteerResponse) GetVolunteer() *Volunteer {
	if x != nil {
		return x.Volunteer
	}
	return nil
}

var File_proto_volunteer_proto protoreflect.FileDescriptor

var file_proto_volunteer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65,
	0x65, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x09, 0x56, 0x6f,
	0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f,
	0x69, 0x76, 0x6f, 0x64, 0x65, 0x73, 0x68, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x76, 0x6f, 0x69, 0x76, 0x6f, 0x64, 0x65, 0x73, 0x68, 0x69, 0x70, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e,
	0x74, 0x65, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x09,
	0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x76, 0x6f, 0x64,
	0x65, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x69,
	0x76, 0x6f, 0x64, 0x65, 0x73, 0x68, 0x69, 0x70, 0x22, 0x4e, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65,
	0x65, 0x72, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x0a, 0x76, 0x6f,
	0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x09, 0x76, 0x6f, 0x6c,
	0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72,
	0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x52, 0x09, 0x76, 0x6f, 0x6c, 0x75,
	0x6e, 0x74, 0x65, 0x65, 0x72, 0x32, 0x94, 0x02, 0x0a, 0x10, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74,
	0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e,
	0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75,
	0x6e, 0x74, 0x65, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6e,
	0x74, 0x65, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6e,
	0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6e, 0x74, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_volunteer_proto_rawDescOnce sync.Once
	file_proto_volunteer_proto_rawDescData = file_proto_volunteer_proto_rawDesc
)

func file_proto_volunteer_proto_rawDescGZIP() []byte {
	file_proto_volunteer_proto_rawDescOnce.Do(func() {
		file_proto_volunteer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_volunteer_proto_rawDescData)
	})
	return file_proto_volunteer_proto_rawDescData
}

var file_proto_volunteer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_volunteer_proto_goTypes = []any{
	(*GetVolunteerRequest)(nil),     // 0: volunteer.GetVolunteerRequest
	(*Volunteer)(nil),               // 1: volunteer.Volunteer
	(*GetVolunteerResponse)(nil),    // 2: volunteer.GetVolunteerResponse
	(*ListVolunteersRequest)(nil),   // 3: volunteer.ListVolunteersRequest
	(*ListVolunteersResponse)(nil),  // 4: volunteer.ListVolunteersResponse
	(*CreateVolunteerRequest)(nil),  // 5: volunteer.CreateVolunteerRequest
	(*CreateVolunteerResponse)(nil), // 6: volunteer.CreateVolunteerResponse
}
var file_proto_volunteer_proto_depIdxs = []int32{
	1, // 0: volunteer.GetVolunteerResponse.volunteer:type_name -> volunteer.Volunteer
	1, // 1: volunteer.ListVolunteersResponse.volunteers:type_name -> volunteer.Volunteer
	1, // 2: volunteer.CreateVolunteerRequest.volunteer:type_name -> volunteer.Volunteer
	1, // 3: volunteer.CreateVolunteerResponse.volunteer:type_name -> volunteer.Volunteer
	0, // 4: volunteer.VolunteerService.GetVolunteer:input_type -> volunteer.GetVolunteerRequest
	3, // 5: volunteer.VolunteerService.ListVolunteers:input_type -> volunteer.ListVolunteersRequest
	5, // 6: volunteer.VolunteerService.CreateVolunteer:input_type -> volunteer.CreateVolunteerRequest
	2, // 7: volunteer.VolunteerService.GetVolunteer:output_type -> volunteer.GetVolunteerResponse
	4, // 8: volunteer.VolunteerService.ListVolunteers:output_type -> volunteer.ListVolunteersResponse
	6, // 9: volunteer.VolunteerService.CreateVolunteer:output_type -> volunteer.CreateVolunteerResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_volunteer_proto_init() }
func file_proto_volunteer_proto_init() {
	if File_proto_volunteer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_volunteer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetVolunteerRequest); i {
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
		file_proto_volunteer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Volunteer); i {
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
		file_proto_volunteer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetVolunteerResponse); i {
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
		file_proto_volunteer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListVolunteersRequest); i {
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
		file_proto_volunteer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListVolunteersResponse); i {
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
		file_proto_volunteer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVolunteerRequest); i {
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
		file_proto_volunteer_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVolunteerResponse); i {
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
	file_proto_volunteer_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_volunteer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_volunteer_proto_goTypes,
		DependencyIndexes: file_proto_volunteer_proto_depIdxs,
		MessageInfos:      file_proto_volunteer_proto_msgTypes,
	}.Build()
	File_proto_volunteer_proto = out.File
	file_proto_volunteer_proto_rawDesc = nil
	file_proto_volunteer_proto_goTypes = nil
	file_proto_volunteer_proto_depIdxs = nil
}
