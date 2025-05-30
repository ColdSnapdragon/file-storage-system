// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.15.5
// source: download.proto

package download

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReqEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReqEntry) Reset() {
	*x = ReqEntry{}
	mi := &file_download_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqEntry) ProtoMessage() {}

func (x *ReqEntry) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqEntry.ProtoReflect.Descriptor instead.
func (*ReqEntry) Descriptor() ([]byte, []int) {
	return file_download_proto_rawDescGZIP(), []int{0}
}

type RespEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Entry         string                 `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RespEntry) Reset() {
	*x = RespEntry{}
	mi := &file_download_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RespEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespEntry) ProtoMessage() {}

func (x *RespEntry) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespEntry.ProtoReflect.Descriptor instead.
func (*RespEntry) Descriptor() ([]byte, []int) {
	return file_download_proto_rawDescGZIP(), []int{1}
}

func (x *RespEntry) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespEntry) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RespEntry) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

var File_download_proto protoreflect.FileDescriptor

const file_download_proto_rawDesc = "" +
	"\n" +
	"\x0edownload.proto\x12\x19go.micro.service.download\"\n" +
	"\n" +
	"\bReqEntry\"O\n" +
	"\tRespEntry\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x14\n" +
	"\x05entry\x18\x03 \x01(\tR\x05entry2o\n" +
	"\x0fDownloadService\x12\\\n" +
	"\rDownloadEntry\x12#.go.micro.service.download.ReqEntry\x1a$.go.micro.service.download.RespEntry\"\x00B\vZ\t/downloadb\x06proto3"

var (
	file_download_proto_rawDescOnce sync.Once
	file_download_proto_rawDescData []byte
)

func file_download_proto_rawDescGZIP() []byte {
	file_download_proto_rawDescOnce.Do(func() {
		file_download_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_download_proto_rawDesc), len(file_download_proto_rawDesc)))
	})
	return file_download_proto_rawDescData
}

var file_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_download_proto_goTypes = []any{
	(*ReqEntry)(nil),  // 0: go.micro.service.download.ReqEntry
	(*RespEntry)(nil), // 1: go.micro.service.download.RespEntry
}
var file_download_proto_depIdxs = []int32{
	0, // 0: go.micro.service.download.DownloadService.DownloadEntry:input_type -> go.micro.service.download.ReqEntry
	1, // 1: go.micro.service.download.DownloadService.DownloadEntry:output_type -> go.micro.service.download.RespEntry
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_download_proto_init() }
func file_download_proto_init() {
	if File_download_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_download_proto_rawDesc), len(file_download_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_download_proto_goTypes,
		DependencyIndexes: file_download_proto_depIdxs,
		MessageInfos:      file_download_proto_msgTypes,
	}.Build()
	File_download_proto = out.File
	file_download_proto_goTypes = nil
	file_download_proto_depIdxs = nil
}
