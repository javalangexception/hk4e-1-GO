// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: Material.proto

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

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      uint32              `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	DeleteInfo *MaterialDeleteInfo `protobuf:"bytes,2,opt,name=delete_info,json=deleteInfo,proto3" json:"delete_info,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Material_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_Material_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_Material_proto_rawDescGZIP(), []int{0}
}

func (x *Material) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Material) GetDeleteInfo() *MaterialDeleteInfo {
	if x != nil {
		return x.DeleteInfo
	}
	return nil
}

var File_Material_proto protoreflect.FileDescriptor

var file_Material_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5c, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Material_proto_rawDescOnce sync.Once
	file_Material_proto_rawDescData = file_Material_proto_rawDesc
)

func file_Material_proto_rawDescGZIP() []byte {
	file_Material_proto_rawDescOnce.Do(func() {
		file_Material_proto_rawDescData = protoimpl.X.CompressGZIP(file_Material_proto_rawDescData)
	})
	return file_Material_proto_rawDescData
}

var file_Material_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Material_proto_goTypes = []interface{}{
	(*Material)(nil),           // 0: proto.Material
	(*MaterialDeleteInfo)(nil), // 1: proto.MaterialDeleteInfo
}
var file_Material_proto_depIdxs = []int32{
	1, // 0: proto.Material.delete_info:type_name -> proto.MaterialDeleteInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Material_proto_init() }
func file_Material_proto_init() {
	if File_Material_proto != nil {
		return
	}
	file_MaterialDeleteInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Material_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
			RawDescriptor: file_Material_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Material_proto_goTypes,
		DependencyIndexes: file_Material_proto_depIdxs,
		MessageInfos:      file_Material_proto_msgTypes,
	}.Build()
	File_Material_proto = out.File
	file_Material_proto_rawDesc = nil
	file_Material_proto_goTypes = nil
	file_Material_proto_depIdxs = nil
}