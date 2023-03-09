// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: apicentral.proto

package proto

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

type Owner_Type int32

const (
	Owner_TEAM Owner_Type = 0
)

// Enum value maps for Owner_Type.
var (
	Owner_Type_name = map[int32]string{
		0: "TEAM",
	}
	Owner_Type_value = map[string]int32{
		"TEAM": 0,
	}
)

func (x Owner_Type) Enum() *Owner_Type {
	p := new(Owner_Type)
	*p = x
	return p
}

func (x Owner_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Owner_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apicentral_proto_enumTypes[0].Descriptor()
}

func (Owner_Type) Type() protoreflect.EnumType {
	return &file_apicentral_proto_enumTypes[0]
}

func (x Owner_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Owner_Type.Descriptor instead.
func (Owner_Type) EnumDescriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{2, 0}
}

type Reference_Type int32

const (
	Reference_SOFT Reference_Type = 0
	Reference_HARD Reference_Type = 1
)

// Enum value maps for Reference_Type.
var (
	Reference_Type_name = map[int32]string{
		0: "SOFT",
		1: "HARD",
	}
	Reference_Type_value = map[string]int32{
		"SOFT": 0,
		"HARD": 1,
	}
)

func (x Reference_Type) Enum() *Reference_Type {
	p := new(Reference_Type)
	*p = x
	return p
}

func (x Reference_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reference_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apicentral_proto_enumTypes[1].Descriptor()
}

func (Reference_Type) Type() protoreflect.EnumType {
	return &file_apicentral_proto_enumTypes[1]
}

func (x Reference_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reference_Type.Descriptor instead.
func (Reference_Type) EnumDescriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{3, 0}
}

// API Server generic resource structure.
type ResourceInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the group from which the resource belongs to. The server infers this from the endpoint the client submits the request to.
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// Resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of the resource
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata.
	Metadata   *Metadata         `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Attributes map[string]string `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceInstance) Reset() {
	*x = ResourceInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apicentral_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInstance) ProtoMessage() {}

func (x *ResourceInstance) ProtoReflect() protoreflect.Message {
	mi := &file_apicentral_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInstance.ProtoReflect.Descriptor instead.
func (*ResourceInstance) Descriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceInstance) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ResourceInstance) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceInstance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceInstance) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ResourceInstance) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Metadata that all server resources have. Data is generated by the server.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal id of the resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The scope where this resource was defined.
	Scope *Metadata_ScopeKind `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	// The URL representing this resource object.
	SelfLink string `protobuf:"bytes,6,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
	// resource references
	References []*Reference `protobuf:"bytes,7,rep,name=references,proto3" json:"references,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apicentral_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_apicentral_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetScope() *Metadata_ScopeKind {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Metadata) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

func (x *Metadata) GetReferences() []*Reference {
	if x != nil {
		return x.References
	}
	return nil
}

// Owner of the resource.
type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the owner of the resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the owner. Defaults to team if not present.
	Type Owner_Type `protobuf:"varint,2,opt,name=type,proto3,enum=central.events.v1.datamodel.Owner_Type" json:"type,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apicentral_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_apicentral_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{2}
}

func (x *Owner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Owner) GetType() Owner_Type {
	if x != nil {
		return x.Type
	}
	return Owner_TEAM
}

// Reference resource
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id generated by the server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The kind of the referenced resource.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// The name of the referenced resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The kind of the referenced resource scope.
	ScopeKind string `protobuf:"bytes,4,opt,name=scopeKind,proto3" json:"scopeKind,omitempty"`
	// The name of the referenced resource scope.
	ScopeName string `protobuf:"bytes,5,opt,name=scopeName,proto3" json:"scopeName,omitempty"`
	// The URL representing the referenced resource.
	SelfLink string `protobuf:"bytes,6,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
	// Defines the type of the reference: * soft - spec property that has this reference will get nulled out if the referenced resource gets removed. * hard - dictates that the current resource will get removed when the referenced resource gets removed.
	Type Reference_Type `protobuf:"varint,7,opt,name=type,proto3,enum=central.events.v1.datamodel.Reference_Type" json:"type,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apicentral_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_apicentral_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{3}
}

func (x *Reference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Reference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reference) GetScopeKind() string {
	if x != nil {
		return x.ScopeKind
	}
	return ""
}

func (x *Reference) GetScopeName() string {
	if x != nil {
		return x.ScopeName
	}
	return ""
}

func (x *Reference) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

func (x *Reference) GetType() Reference_Type {
	if x != nil {
		return x.Type
	}
	return Reference_SOFT
}

type Metadata_ScopeKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal id of the scope resource where the resource is defined.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The kind of the scope resource where the resource is defined.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// The name of the scope where the resource is defined.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The URL to access the scope resource.
	SelfLink string `protobuf:"bytes,4,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
}

func (x *Metadata_ScopeKind) Reset() {
	*x = Metadata_ScopeKind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apicentral_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata_ScopeKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata_ScopeKind) ProtoMessage() {}

func (x *Metadata_ScopeKind) ProtoReflect() protoreflect.Message {
	mi := &file_apicentral_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata_ScopeKind.ProtoReflect.Descriptor instead.
func (*Metadata_ScopeKind) Descriptor() ([]byte, []int) {
	return file_apicentral_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Metadata_ScopeKind) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata_ScopeKind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Metadata_ScopeKind) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata_ScopeKind) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

var File_apicentral_proto protoreflect.FileDescriptor

var file_apicentral_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0xb1, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xa6, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x45, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x46, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x5f, 0x0a, 0x09, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x66, 0x0a, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x10, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45,
	0x41, 0x4d, 0x10, 0x00, 0x22, 0xf8, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69,
	0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x4f, 0x46, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x52, 0x44, 0x10, 0x01, 0x42,
	0x18, 0x5a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apicentral_proto_rawDescOnce sync.Once
	file_apicentral_proto_rawDescData = file_apicentral_proto_rawDesc
)

func file_apicentral_proto_rawDescGZIP() []byte {
	file_apicentral_proto_rawDescOnce.Do(func() {
		file_apicentral_proto_rawDescData = protoimpl.X.CompressGZIP(file_apicentral_proto_rawDescData)
	})
	return file_apicentral_proto_rawDescData
}

var file_apicentral_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apicentral_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apicentral_proto_goTypes = []interface{}{
	(Owner_Type)(0),            // 0: central.events.v1.datamodel.Owner.Type
	(Reference_Type)(0),        // 1: central.events.v1.datamodel.Reference.Type
	(*ResourceInstance)(nil),   // 2: central.events.v1.datamodel.ResourceInstance
	(*Metadata)(nil),           // 3: central.events.v1.datamodel.Metadata
	(*Owner)(nil),              // 4: central.events.v1.datamodel.Owner
	(*Reference)(nil),          // 5: central.events.v1.datamodel.Reference
	nil,                        // 6: central.events.v1.datamodel.ResourceInstance.AttributesEntry
	(*Metadata_ScopeKind)(nil), // 7: central.events.v1.datamodel.Metadata.ScopeKind
}
var file_apicentral_proto_depIdxs = []int32{
	3, // 0: central.events.v1.datamodel.ResourceInstance.metadata:type_name -> central.events.v1.datamodel.Metadata
	6, // 1: central.events.v1.datamodel.ResourceInstance.attributes:type_name -> central.events.v1.datamodel.ResourceInstance.AttributesEntry
	7, // 2: central.events.v1.datamodel.Metadata.scope:type_name -> central.events.v1.datamodel.Metadata.ScopeKind
	5, // 3: central.events.v1.datamodel.Metadata.references:type_name -> central.events.v1.datamodel.Reference
	0, // 4: central.events.v1.datamodel.Owner.type:type_name -> central.events.v1.datamodel.Owner.Type
	1, // 5: central.events.v1.datamodel.Reference.type:type_name -> central.events.v1.datamodel.Reference.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apicentral_proto_init() }
func file_apicentral_proto_init() {
	if File_apicentral_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apicentral_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInstance); i {
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
		file_apicentral_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_apicentral_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
		file_apicentral_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_apicentral_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata_ScopeKind); i {
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
			RawDescriptor: file_apicentral_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apicentral_proto_goTypes,
		DependencyIndexes: file_apicentral_proto_depIdxs,
		EnumInfos:         file_apicentral_proto_enumTypes,
		MessageInfos:      file_apicentral_proto_msgTypes,
	}.Build()
	File_apicentral_proto = out.File
	file_apicentral_proto_rawDesc = nil
	file_apicentral_proto_goTypes = nil
	file_apicentral_proto_depIdxs = nil
}
