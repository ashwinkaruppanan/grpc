// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/email.proto

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

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{0}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{1}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type NameList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *NameList) Reset() {
	*x = NameList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameList) ProtoMessage() {}

func (x *NameList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameList.ProtoReflect.Descriptor instead.
func (*NameList) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{2}
}

func (x *NameList) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type EmailList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailList []string `protobuf:"bytes,1,rep,name=EmailList,proto3" json:"EmailList,omitempty"`
}

func (x *EmailList) Reset() {
	*x = EmailList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailList) ProtoMessage() {}

func (x *EmailList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailList.ProtoReflect.Descriptor instead.
func (*EmailList) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{3}
}

func (x *EmailList) GetEmailList() []string {
	if x != nil {
		return x.EmailList
	}
	return nil
}

var File_proto_email_proto protoreflect.FileDescriptor

var file_proto_email_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x20, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0xf4, 0x01, 0x0a, 0x0e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x30, 0x01, 0x12, 0x39,
	0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x1d, 0x42, 0x69, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_email_proto_rawDescOnce sync.Once
	file_proto_email_proto_rawDescData = file_proto_email_proto_rawDesc
)

func file_proto_email_proto_rawDescGZIP() []byte {
	file_proto_email_proto_rawDescOnce.Do(func() {
		file_proto_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_email_proto_rawDescData)
	})
	return file_proto_email_proto_rawDescData
}

var file_proto_email_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_email_proto_goTypes = []interface{}{
	(*Name)(nil),      // 0: email.Name
	(*Email)(nil),     // 1: email.Email
	(*NameList)(nil),  // 2: email.NameList
	(*EmailList)(nil), // 3: email.EmailList
}
var file_proto_email_proto_depIdxs = []int32{
	0, // 0: email.EmailGenerator.UnaryConnection:input_type -> email.Name
	2, // 1: email.EmailGenerator.ServerStreamConnection:input_type -> email.NameList
	0, // 2: email.EmailGenerator.ClientStreamConnection:input_type -> email.Name
	0, // 3: email.EmailGenerator.BidirectionalStreamConnection:input_type -> email.Name
	1, // 4: email.EmailGenerator.UnaryConnection:output_type -> email.Email
	1, // 5: email.EmailGenerator.ServerStreamConnection:output_type -> email.Email
	3, // 6: email.EmailGenerator.ClientStreamConnection:output_type -> email.EmailList
	1, // 7: email.EmailGenerator.BidirectionalStreamConnection:output_type -> email.Email
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_email_proto_init() }
func file_proto_email_proto_init() {
	if File_proto_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_proto_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_proto_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameList); i {
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
		file_proto_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailList); i {
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
			RawDescriptor: file_proto_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_email_proto_goTypes,
		DependencyIndexes: file_proto_email_proto_depIdxs,
		MessageInfos:      file_proto_email_proto_msgTypes,
	}.Build()
	File_proto_email_proto = out.File
	file_proto_email_proto_rawDesc = nil
	file_proto_email_proto_goTypes = nil
	file_proto_email_proto_depIdxs = nil
}
