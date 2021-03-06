// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: common/common.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ResponseStatus_StatusCode int32

const (
	ResponseStatus_SUCCESS               ResponseStatus_StatusCode = 0
	ResponseStatus_UNAUTHORIZED          ResponseStatus_StatusCode = 401
	ResponseStatus_INTERNAL_SERVER_ERROR ResponseStatus_StatusCode = 500
	ResponseStatus_UNKNOWN               ResponseStatus_StatusCode = 1400
	ResponseStatus_OBJECT_NOT_FOUND      ResponseStatus_StatusCode = 1401
	ResponseStatus_VALUE_NOT_FOUND       ResponseStatus_StatusCode = 1402
	ResponseStatus_VALUE_ALREADY_EXISTS  ResponseStatus_StatusCode = 1403
	ResponseStatus_VALUE_OUT_OF_RANGE    ResponseStatus_StatusCode = 1404
	ResponseStatus_INVALID_VALUE         ResponseStatus_StatusCode = 1405
	ResponseStatus_FAILED_DEPENDENCY     ResponseStatus_StatusCode = 1406
	ResponseStatus_FORBIDDEN             ResponseStatus_StatusCode = 1407
	ResponseStatus_USERNAME_TAKEN        ResponseStatus_StatusCode = 1408
	ResponseStatus_EMAIL_TAKEN           ResponseStatus_StatusCode = 1409
	ResponseStatus_REPO_NOT_EXISTS       ResponseStatus_StatusCode = 1410
	ResponseStatus_GET_BRANCH_FAILED     ResponseStatus_StatusCode = 1411
	ResponseStatus_GET_COMMIT_FAILED     ResponseStatus_StatusCode = 1412
	ResponseStatus_INVALID_GIT_URL       ResponseStatus_StatusCode = 1413
	ResponseStatus_EMPTY_REPO            ResponseStatus_StatusCode = 1414
)

// Enum value maps for ResponseStatus_StatusCode.
var (
	ResponseStatus_StatusCode_name = map[int32]string{
		0:    "SUCCESS",
		401:  "UNAUTHORIZED",
		500:  "INTERNAL_SERVER_ERROR",
		1400: "UNKNOWN",
		1401: "OBJECT_NOT_FOUND",
		1402: "VALUE_NOT_FOUND",
		1403: "VALUE_ALREADY_EXISTS",
		1404: "VALUE_OUT_OF_RANGE",
		1405: "INVALID_VALUE",
		1406: "FAILED_DEPENDENCY",
		1407: "FORBIDDEN",
		1408: "USERNAME_TAKEN",
		1409: "EMAIL_TAKEN",
		1410: "REPO_NOT_EXISTS",
		1411: "GET_BRANCH_FAILED",
		1412: "GET_COMMIT_FAILED",
		1413: "INVALID_GIT_URL",
		1414: "EMPTY_REPO",
	}
	ResponseStatus_StatusCode_value = map[string]int32{
		"SUCCESS":               0,
		"UNAUTHORIZED":          401,
		"INTERNAL_SERVER_ERROR": 500,
		"UNKNOWN":               1400,
		"OBJECT_NOT_FOUND":      1401,
		"VALUE_NOT_FOUND":       1402,
		"VALUE_ALREADY_EXISTS":  1403,
		"VALUE_OUT_OF_RANGE":    1404,
		"INVALID_VALUE":         1405,
		"FAILED_DEPENDENCY":     1406,
		"FORBIDDEN":             1407,
		"USERNAME_TAKEN":        1408,
		"EMAIL_TAKEN":           1409,
		"REPO_NOT_EXISTS":       1410,
		"GET_BRANCH_FAILED":     1411,
		"GET_COMMIT_FAILED":     1412,
		"INVALID_GIT_URL":       1413,
		"EMPTY_REPO":            1414,
	}
)

func (x ResponseStatus_StatusCode) Enum() *ResponseStatus_StatusCode {
	p := new(ResponseStatus_StatusCode)
	*p = x
	return p
}

func (x ResponseStatus_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (ResponseStatus_StatusCode) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x ResponseStatus_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus_StatusCode.Descriptor instead.
func (ResponseStatus_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ResponseStatus_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.ResponseStatus_StatusCode" json:"code,omitempty"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseStatus) GetCode() ResponseStatus_StatusCode {
	if x != nil {
		return x.Code
	}
	return ResponseStatus_SUCCESS
}

func (x *ResponseStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xeb, 0x03,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x35, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x87, 0x03, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x91, 0x03,
	0x12, 0x1a, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x12, 0x0c, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xf8, 0x0a, 0x12, 0x15, 0x0a, 0x10, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xf9,
	0x0a, 0x12, 0x14, 0x0a, 0x0f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0xfa, 0x0a, 0x12, 0x19, 0x0a, 0x14, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0xfb, 0x0a, 0x12, 0x17, 0x0a, 0x12, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0xfc, 0x0a, 0x12, 0x12, 0x0a, 0x0d, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0xfd, 0x0a, 0x12,
	0x16, 0x0a, 0x11, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x50, 0x45, 0x4e, 0x44,
	0x45, 0x4e, 0x43, 0x59, 0x10, 0xfe, 0x0a, 0x12, 0x0e, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49,
	0x44, 0x44, 0x45, 0x4e, 0x10, 0xff, 0x0a, 0x12, 0x13, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x80, 0x0b, 0x12, 0x10, 0x0a, 0x0b,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x81, 0x0b, 0x12, 0x14,
	0x0a, 0x0f, 0x52, 0x45, 0x50, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x53, 0x10, 0x82, 0x0b, 0x12, 0x16, 0x0a, 0x11, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x52, 0x41, 0x4e,
	0x43, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x83, 0x0b, 0x12, 0x16, 0x0a, 0x11,
	0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x84, 0x0b, 0x12, 0x14, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x47, 0x49, 0x54, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x85, 0x0b, 0x12, 0x0f, 0x0a, 0x0a, 0x45, 0x4d,
	0x50, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x10, 0x86, 0x0b, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x74, 0x65, 0x6c, 0x65,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x76, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_common_proto_goTypes = []interface{}{
	(ResponseStatus_StatusCode)(0), // 0: common.ResponseStatus.StatusCode
	(*ResponseStatus)(nil),         // 1: common.ResponseStatus
}
var file_common_common_proto_depIdxs = []int32{
	0, // 0: common.ResponseStatus.code:type_name -> common.ResponseStatus.StatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
