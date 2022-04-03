// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: proto/superhero/superhero.proto

package api

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

type Superhero_Home int32

const (
	Superhero_marvel Superhero_Home = 0
	Superhero_DC     Superhero_Home = 1
)

// Enum value maps for Superhero_Home.
var (
	Superhero_Home_name = map[int32]string{
		0: "marvel",
		1: "DC",
	}
	Superhero_Home_value = map[string]int32{
		"marvel": 0,
		"DC":     1,
	}
)

func (x Superhero_Home) Enum() *Superhero_Home {
	p := new(Superhero_Home)
	*p = x
	return p
}

func (x Superhero_Home) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Superhero_Home) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_superhero_superhero_proto_enumTypes[0].Descriptor()
}

func (Superhero_Home) Type() protoreflect.EnumType {
	return &file_proto_superhero_superhero_proto_enumTypes[0]
}

func (x Superhero_Home) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Superhero_Home.Descriptor instead.
func (Superhero_Home) EnumDescriptor() ([]byte, []int) {
	return file_proto_superhero_superhero_proto_rawDescGZIP(), []int{0, 0}
}

type Superhero_Attributes_Genre int32

const (
	Superhero_Attributes_male   Superhero_Attributes_Genre = 0
	Superhero_Attributes_female Superhero_Attributes_Genre = 1
)

// Enum value maps for Superhero_Attributes_Genre.
var (
	Superhero_Attributes_Genre_name = map[int32]string{
		0: "male",
		1: "female",
	}
	Superhero_Attributes_Genre_value = map[string]int32{
		"male":   0,
		"female": 1,
	}
)

func (x Superhero_Attributes_Genre) Enum() *Superhero_Attributes_Genre {
	p := new(Superhero_Attributes_Genre)
	*p = x
	return p
}

func (x Superhero_Attributes_Genre) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Superhero_Attributes_Genre) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_superhero_superhero_proto_enumTypes[1].Descriptor()
}

func (Superhero_Attributes_Genre) Type() protoreflect.EnumType {
	return &file_proto_superhero_superhero_proto_enumTypes[1]
}

func (x Superhero_Attributes_Genre) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Superhero_Attributes_Genre.Descriptor instead.
func (Superhero_Attributes_Genre) EnumDescriptor() ([]byte, []int) {
	return file_proto_superhero_superhero_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Superhero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Skill string         `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	Home  Superhero_Home `protobuf:"varint,4,opt,name=home,proto3,enum=superhero.v1.Superhero_Home" json:"home,omitempty"`
}

func (x *Superhero) Reset() {
	*x = Superhero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_superhero_superhero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Superhero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Superhero) ProtoMessage() {}

func (x *Superhero) ProtoReflect() protoreflect.Message {
	mi := &file_proto_superhero_superhero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Superhero.ProtoReflect.Descriptor instead.
func (*Superhero) Descriptor() ([]byte, []int) {
	return file_proto_superhero_superhero_proto_rawDescGZIP(), []int{0}
}

func (x *Superhero) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Superhero) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Superhero) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

func (x *Superhero) GetHome() Superhero_Home {
	if x != nil {
		return x.Home
	}
	return Superhero_marvel
}

// These are the attributes of the superhero
type Superhero_Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre    Superhero_Attributes_Genre `protobuf:"varint,1,opt,name=genre,proto3,enum=superhero.v1.Superhero_Attributes_Genre" json:"genre,omitempty"`
	Weakness string                     `protobuf:"bytes,2,opt,name=weakness,proto3" json:"weakness,omitempty"`
	Attack   int32                      `protobuf:"varint,3,opt,name=attack,proto3" json:"attack,omitempty"`
	Health   int32                      `protobuf:"varint,4,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *Superhero_Attributes) Reset() {
	*x = Superhero_Attributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_superhero_superhero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Superhero_Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Superhero_Attributes) ProtoMessage() {}

func (x *Superhero_Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_superhero_superhero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Superhero_Attributes.ProtoReflect.Descriptor instead.
func (*Superhero_Attributes) Descriptor() ([]byte, []int) {
	return file_proto_superhero_superhero_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Superhero_Attributes) GetGenre() Superhero_Attributes_Genre {
	if x != nil {
		return x.Genre
	}
	return Superhero_Attributes_male
}

func (x *Superhero_Attributes) GetWeakness() string {
	if x != nil {
		return x.Weakness
	}
	return ""
}

func (x *Superhero_Attributes) GetAttack() int32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *Superhero_Attributes) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

var File_proto_superhero_superhero_proto protoreflect.FileDescriptor

var file_proto_superhero_superhero_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72,
	0x6f, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x22,
	0xcd, 0x02, 0x0a, 0x09, 0x53, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x30, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x1a, 0xb7, 0x01, 0x0a, 0x0a, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68,
	0x65, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x65, 0x61, 0x6b,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x61, 0x6b,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x22, 0x1d, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x6d, 0x61, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x10, 0x01, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x6d,
	0x61, 0x72, 0x76, 0x65, 0x6c, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x43, 0x10, 0x01, 0x42,
	0x0f, 0x5a, 0x0d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_superhero_superhero_proto_rawDescOnce sync.Once
	file_proto_superhero_superhero_proto_rawDescData = file_proto_superhero_superhero_proto_rawDesc
)

func file_proto_superhero_superhero_proto_rawDescGZIP() []byte {
	file_proto_superhero_superhero_proto_rawDescOnce.Do(func() {
		file_proto_superhero_superhero_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_superhero_superhero_proto_rawDescData)
	})
	return file_proto_superhero_superhero_proto_rawDescData
}

var file_proto_superhero_superhero_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_superhero_superhero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_superhero_superhero_proto_goTypes = []interface{}{
	(Superhero_Home)(0),             // 0: superhero.v1.Superhero.Home
	(Superhero_Attributes_Genre)(0), // 1: superhero.v1.Superhero.Attributes.Genre
	(*Superhero)(nil),               // 2: superhero.v1.Superhero
	(*Superhero_Attributes)(nil),    // 3: superhero.v1.Superhero.Attributes
}
var file_proto_superhero_superhero_proto_depIdxs = []int32{
	0, // 0: superhero.v1.Superhero.home:type_name -> superhero.v1.Superhero.Home
	1, // 1: superhero.v1.Superhero.Attributes.genre:type_name -> superhero.v1.Superhero.Attributes.Genre
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_superhero_superhero_proto_init() }
func file_proto_superhero_superhero_proto_init() {
	if File_proto_superhero_superhero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_superhero_superhero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Superhero); i {
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
		file_proto_superhero_superhero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Superhero_Attributes); i {
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
			RawDescriptor: file_proto_superhero_superhero_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_superhero_superhero_proto_goTypes,
		DependencyIndexes: file_proto_superhero_superhero_proto_depIdxs,
		EnumInfos:         file_proto_superhero_superhero_proto_enumTypes,
		MessageInfos:      file_proto_superhero_superhero_proto_msgTypes,
	}.Build()
	File_proto_superhero_superhero_proto = out.File
	file_proto_superhero_superhero_proto_rawDesc = nil
	file_proto_superhero_superhero_proto_goTypes = nil
	file_proto_superhero_superhero_proto_depIdxs = nil
}
