// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.3
// source: imgurl.proto

package grpc

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

type Format int32

const (
	Format_JPG  Format = 0
	Format_PNG  Format = 1
	Format_BMP  Format = 2
	Format_WEBP Format = 3
	Format_GIF  Format = 4
	Format_ICO  Format = 5
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "JPG",
		1: "PNG",
		2: "BMP",
		3: "WEBP",
		4: "GIF",
		5: "ICO",
	}
	Format_value = map[string]int32{
		"JPG":  0,
		"PNG":  1,
		"BMP":  2,
		"WEBP": 3,
		"GIF":  4,
		"ICO":  5,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_imgurl_proto_enumTypes[0].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_imgurl_proto_enumTypes[0]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_imgurl_proto_rawDescGZIP(), []int{0}
}

type UrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image  string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Params []string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
	Format *Format  `protobuf:"varint,3,opt,name=format,proto3,enum=Format,oneof" json:"format,omitempty"`
}

func (x *UrlRequest) Reset() {
	*x = UrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imgurl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imgurl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRequest.ProtoReflect.Descriptor instead.
func (*UrlRequest) Descriptor() ([]byte, []int) {
	return file_imgurl_proto_rawDescGZIP(), []int{0}
}

func (x *UrlRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UrlRequest) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *UrlRequest) GetFormat() Format {
	if x != nil && x.Format != nil {
		return *x.Format
	}
	return Format_JPG
}

type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imgurl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imgurl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_imgurl_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_imgurl_proto protoreflect.FileDescriptor

var file_imgurl_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6d, 0x67, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b,
	0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x48, 0x00, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x1f, 0x0a, 0x0b, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x3f, 0x0a, 0x06,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x47, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x50, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4d, 0x50, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x42, 0x50, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x49, 0x46, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x43, 0x4f, 0x10, 0x05, 0x32, 0x34, 0x0a,
	0x09, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x08, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x90, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x78, 0x70, 0x72, 0x6f, 0x2f, 0x69,
	0x6d, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x75, 0x72, 0x6c, 0x2d, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0xca, 0x02, 0x26, 0x53, 0x69,
	0x74, 0x65, 0x77, 0x6f, 0x72, 0x78, 0x5c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x49, 0x6d, 0x67, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0xe2, 0x02, 0x2f, 0x53, 0x69, 0x74, 0x65, 0x77, 0x6f, 0x72, 0x78, 0x5c,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x49, 0x6d, 0x67, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x5c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imgurl_proto_rawDescOnce sync.Once
	file_imgurl_proto_rawDescData = file_imgurl_proto_rawDesc
)

func file_imgurl_proto_rawDescGZIP() []byte {
	file_imgurl_proto_rawDescOnce.Do(func() {
		file_imgurl_proto_rawDescData = protoimpl.X.CompressGZIP(file_imgurl_proto_rawDescData)
	})
	return file_imgurl_proto_rawDescData
}

var file_imgurl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_imgurl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_imgurl_proto_goTypes = []interface{}{
	(Format)(0),         // 0: Format
	(*UrlRequest)(nil),  // 1: UrlRequest
	(*UrlResponse)(nil), // 2: UrlResponse
}
var file_imgurl_proto_depIdxs = []int32{
	0, // 0: UrlRequest.format:type_name -> Format
	1, // 1: Generator.Generate:input_type -> UrlRequest
	2, // 2: Generator.Generate:output_type -> UrlResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_imgurl_proto_init() }
func file_imgurl_proto_init() {
	if File_imgurl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imgurl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlRequest); i {
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
		file_imgurl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
	file_imgurl_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_imgurl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imgurl_proto_goTypes,
		DependencyIndexes: file_imgurl_proto_depIdxs,
		EnumInfos:         file_imgurl_proto_enumTypes,
		MessageInfos:      file_imgurl_proto_msgTypes,
	}.Build()
	File_imgurl_proto = out.File
	file_imgurl_proto_rawDesc = nil
	file_imgurl_proto_goTypes = nil
	file_imgurl_proto_depIdxs = nil
}
