// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc/proto/helloworld.proto

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

// the request message containing the user's name.
type Hellorequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hellorequest) Reset() {
	*x = Hellorequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hellorequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hellorequest) ProtoMessage() {}

func (x *Hellorequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hellorequest.ProtoReflect.Descriptor instead.
func (*Hellorequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *Hellorequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// the response message containing the greetings
type Helloreply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Helloreply) Reset() {
	*x = Helloreply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helloreply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helloreply) ProtoMessage() {}

func (x *Helloreply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helloreply.ProtoReflect.Descriptor instead.
func (*Helloreply) Descriptor() ([]byte, []int) {
	return file_grpc_proto_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *Helloreply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto_helloworld_proto protoreflect.FileDescriptor

var file_grpc_proto_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x49, 0x0a, 0x07, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x3e, 0x0a, 0x08, 0x73, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_helloworld_proto_rawDescOnce sync.Once
	file_grpc_proto_helloworld_proto_rawDescData = file_grpc_proto_helloworld_proto_rawDesc
)

func file_grpc_proto_helloworld_proto_rawDescGZIP() []byte {
	file_grpc_proto_helloworld_proto_rawDescOnce.Do(func() {
		file_grpc_proto_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_helloworld_proto_rawDescData)
	})
	return file_grpc_proto_helloworld_proto_rawDescData
}

var file_grpc_proto_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_helloworld_proto_goTypes = []interface{}{
	(*Hellorequest)(nil), // 0: helloworld.hellorequest
	(*Helloreply)(nil),   // 1: helloworld.helloreply
}
var file_grpc_proto_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.greeter.sayhello:input_type -> helloworld.hellorequest
	1, // 1: helloworld.greeter.sayhello:output_type -> helloworld.helloreply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_helloworld_proto_init() }
func file_grpc_proto_helloworld_proto_init() {
	if File_grpc_proto_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hellorequest); i {
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
		file_grpc_proto_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helloreply); i {
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
			RawDescriptor: file_grpc_proto_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_helloworld_proto_goTypes,
		DependencyIndexes: file_grpc_proto_helloworld_proto_depIdxs,
		MessageInfos:      file_grpc_proto_helloworld_proto_msgTypes,
	}.Build()
	File_grpc_proto_helloworld_proto = out.File
	file_grpc_proto_helloworld_proto_rawDesc = nil
	file_grpc_proto_helloworld_proto_goTypes = nil
	file_grpc_proto_helloworld_proto_depIdxs = nil
}