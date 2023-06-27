// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/bigquery/storage/v1/arrow.proto

package storagepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Compression codec's supported by Arrow.
type ArrowSerializationOptions_CompressionCodec int32

const (
	// If unspecified no compression will be used.
	ArrowSerializationOptions_COMPRESSION_UNSPECIFIED ArrowSerializationOptions_CompressionCodec = 0
	// LZ4 Frame (https://github.com/lz4/lz4/blob/dev/doc/lz4_Frame_format.md)
	ArrowSerializationOptions_LZ4_FRAME ArrowSerializationOptions_CompressionCodec = 1
	// Zstandard compression.
	ArrowSerializationOptions_ZSTD ArrowSerializationOptions_CompressionCodec = 2
)

// Enum value maps for ArrowSerializationOptions_CompressionCodec.
var (
	ArrowSerializationOptions_CompressionCodec_name = map[int32]string{
		0: "COMPRESSION_UNSPECIFIED",
		1: "LZ4_FRAME",
		2: "ZSTD",
	}
	ArrowSerializationOptions_CompressionCodec_value = map[string]int32{
		"COMPRESSION_UNSPECIFIED": 0,
		"LZ4_FRAME":               1,
		"ZSTD":                    2,
	}
)

func (x ArrowSerializationOptions_CompressionCodec) Enum() *ArrowSerializationOptions_CompressionCodec {
	p := new(ArrowSerializationOptions_CompressionCodec)
	*p = x
	return p
}

func (x ArrowSerializationOptions_CompressionCodec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArrowSerializationOptions_CompressionCodec) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_storage_v1_arrow_proto_enumTypes[0].Descriptor()
}

func (ArrowSerializationOptions_CompressionCodec) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_storage_v1_arrow_proto_enumTypes[0]
}

func (x ArrowSerializationOptions_CompressionCodec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArrowSerializationOptions_CompressionCodec.Descriptor instead.
func (ArrowSerializationOptions_CompressionCodec) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescGZIP(), []int{2, 0}
}

// Arrow schema as specified in
// https://arrow.apache.org/docs/python/api/datatypes.html
// and serialized to bytes using IPC:
// https://arrow.apache.org/docs/format/Columnar.html#serialization-and-interprocess-communication-ipc
//
// See code samples on how this message can be deserialized.
type ArrowSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IPC serialized Arrow schema.
	SerializedSchema []byte `protobuf:"bytes,1,opt,name=serialized_schema,json=serializedSchema,proto3" json:"serialized_schema,omitempty"`
}

func (x *ArrowSchema) Reset() {
	*x = ArrowSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowSchema) ProtoMessage() {}

func (x *ArrowSchema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowSchema.ProtoReflect.Descriptor instead.
func (*ArrowSchema) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescGZIP(), []int{0}
}

func (x *ArrowSchema) GetSerializedSchema() []byte {
	if x != nil {
		return x.SerializedSchema
	}
	return nil
}

// Arrow RecordBatch.
type ArrowRecordBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IPC-serialized Arrow RecordBatch.
	SerializedRecordBatch []byte `protobuf:"bytes,1,opt,name=serialized_record_batch,json=serializedRecordBatch,proto3" json:"serialized_record_batch,omitempty"`
	// [Deprecated] The count of rows in `serialized_record_batch`.
	// Please use the format-independent ReadRowsResponse.row_count instead.
	//
	// Deprecated: Marked as deprecated in google/cloud/bigquery/storage/v1/arrow.proto.
	RowCount int64 `protobuf:"varint,2,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
}

func (x *ArrowRecordBatch) Reset() {
	*x = ArrowRecordBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowRecordBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowRecordBatch) ProtoMessage() {}

func (x *ArrowRecordBatch) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowRecordBatch.ProtoReflect.Descriptor instead.
func (*ArrowRecordBatch) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescGZIP(), []int{1}
}

func (x *ArrowRecordBatch) GetSerializedRecordBatch() []byte {
	if x != nil {
		return x.SerializedRecordBatch
	}
	return nil
}

// Deprecated: Marked as deprecated in google/cloud/bigquery/storage/v1/arrow.proto.
func (x *ArrowRecordBatch) GetRowCount() int64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

// Contains options specific to Arrow Serialization.
type ArrowSerializationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The compression codec to use for Arrow buffers in serialized record
	// batches.
	BufferCompression ArrowSerializationOptions_CompressionCodec `protobuf:"varint,2,opt,name=buffer_compression,json=bufferCompression,proto3,enum=google.cloud.bigquery.storage.v1.ArrowSerializationOptions_CompressionCodec" json:"buffer_compression,omitempty"`
}

func (x *ArrowSerializationOptions) Reset() {
	*x = ArrowSerializationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowSerializationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowSerializationOptions) ProtoMessage() {}

func (x *ArrowSerializationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowSerializationOptions.ProtoReflect.Descriptor instead.
func (*ArrowSerializationOptions) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescGZIP(), []int{2}
}

func (x *ArrowSerializationOptions) GetBufferCompression() ArrowSerializationOptions_CompressionCodec {
	if x != nil {
		return x.BufferCompression
	}
	return ArrowSerializationOptions_COMPRESSION_UNSPECIFIED
}

var File_google_cloud_bigquery_storage_v1_arrow_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_storage_v1_arrow_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x3a, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x6b, 0x0a, 0x10,
	0x41, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x36, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x15, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe2, 0x01, 0x0a, 0x19, 0x41, 0x72,
	0x72, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7b, 0x0a, 0x12, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x63, 0x52, 0x11, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x50,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x5a, 0x34, 0x5f, 0x46, 0x52, 0x41,
	0x4d, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x53, 0x54, 0x44, 0x10, 0x02, 0x42, 0xba,
	0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescData = file_google_cloud_bigquery_storage_v1_arrow_proto_rawDesc
)

func file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescData)
	})
	return file_google_cloud_bigquery_storage_v1_arrow_proto_rawDescData
}

var file_google_cloud_bigquery_storage_v1_arrow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_bigquery_storage_v1_arrow_proto_goTypes = []interface{}{
	(ArrowSerializationOptions_CompressionCodec)(0), // 0: google.cloud.bigquery.storage.v1.ArrowSerializationOptions.CompressionCodec
	(*ArrowSchema)(nil),                             // 1: google.cloud.bigquery.storage.v1.ArrowSchema
	(*ArrowRecordBatch)(nil),                        // 2: google.cloud.bigquery.storage.v1.ArrowRecordBatch
	(*ArrowSerializationOptions)(nil),               // 3: google.cloud.bigquery.storage.v1.ArrowSerializationOptions
}
var file_google_cloud_bigquery_storage_v1_arrow_proto_depIdxs = []int32{
	0, // 0: google.cloud.bigquery.storage.v1.ArrowSerializationOptions.buffer_compression:type_name -> google.cloud.bigquery.storage.v1.ArrowSerializationOptions.CompressionCodec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_storage_v1_arrow_proto_init() }
func file_google_cloud_bigquery_storage_v1_arrow_proto_init() {
	if File_google_cloud_bigquery_storage_v1_arrow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowSchema); i {
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
		file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowRecordBatch); i {
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
		file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowSerializationOptions); i {
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
			RawDescriptor: file_google_cloud_bigquery_storage_v1_arrow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_storage_v1_arrow_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_storage_v1_arrow_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_storage_v1_arrow_proto_enumTypes,
		MessageInfos:      file_google_cloud_bigquery_storage_v1_arrow_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_storage_v1_arrow_proto = out.File
	file_google_cloud_bigquery_storage_v1_arrow_proto_rawDesc = nil
	file_google_cloud_bigquery_storage_v1_arrow_proto_goTypes = nil
	file_google_cloud_bigquery_storage_v1_arrow_proto_depIdxs = nil
}
