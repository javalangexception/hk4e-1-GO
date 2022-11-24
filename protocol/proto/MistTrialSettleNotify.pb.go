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
// source: MistTrialSettleNotify.proto

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

// CmdId: 8373
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type MistTrialSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExhibitionListDataRecordMap map[uint32]uint32 `protobuf:"bytes,15,rep,name=exhibition_list_data_record_map,json=exhibitionListDataRecordMap,proto3" json:"exhibition_list_data_record_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ExhibitionListDataResultMap map[uint32]uint32 `protobuf:"bytes,14,rep,name=exhibition_list_data_result_map,json=exhibitionListDataResultMap,proto3" json:"exhibition_list_data_result_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DungeonId                   uint32            `protobuf:"varint,5,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
}

func (x *MistTrialSettleNotify) Reset() {
	*x = MistTrialSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MistTrialSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MistTrialSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MistTrialSettleNotify) ProtoMessage() {}

func (x *MistTrialSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MistTrialSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MistTrialSettleNotify.ProtoReflect.Descriptor instead.
func (*MistTrialSettleNotify) Descriptor() ([]byte, []int) {
	return file_MistTrialSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MistTrialSettleNotify) GetExhibitionListDataRecordMap() map[uint32]uint32 {
	if x != nil {
		return x.ExhibitionListDataRecordMap
	}
	return nil
}

func (x *MistTrialSettleNotify) GetExhibitionListDataResultMap() map[uint32]uint32 {
	if x != nil {
		return x.ExhibitionListDataResultMap
	}
	return nil
}

func (x *MistTrialSettleNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

var File_MistTrialSettleNotify_proto protoreflect.FileDescriptor

var file_MistTrialSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x03, 0x0a, 0x15, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x83,
	0x01, 0x0a, 0x1f, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1b, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4d, 0x61, 0x70, 0x12, 0x83, 0x01, 0x0a, 0x1f, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x45, 0x78, 0x68,
	0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1b, 0x65,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x4e, 0x0a, 0x20, 0x45, 0x78, 0x68,
	0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x20, 0x45, 0x78, 0x68,
	0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MistTrialSettleNotify_proto_rawDescOnce sync.Once
	file_MistTrialSettleNotify_proto_rawDescData = file_MistTrialSettleNotify_proto_rawDesc
)

func file_MistTrialSettleNotify_proto_rawDescGZIP() []byte {
	file_MistTrialSettleNotify_proto_rawDescOnce.Do(func() {
		file_MistTrialSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MistTrialSettleNotify_proto_rawDescData)
	})
	return file_MistTrialSettleNotify_proto_rawDescData
}

var file_MistTrialSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_MistTrialSettleNotify_proto_goTypes = []interface{}{
	(*MistTrialSettleNotify)(nil), // 0: proto.MistTrialSettleNotify
	nil,                           // 1: proto.MistTrialSettleNotify.ExhibitionListDataRecordMapEntry
	nil,                           // 2: proto.MistTrialSettleNotify.ExhibitionListDataResultMapEntry
}
var file_MistTrialSettleNotify_proto_depIdxs = []int32{
	1, // 0: proto.MistTrialSettleNotify.exhibition_list_data_record_map:type_name -> proto.MistTrialSettleNotify.ExhibitionListDataRecordMapEntry
	2, // 1: proto.MistTrialSettleNotify.exhibition_list_data_result_map:type_name -> proto.MistTrialSettleNotify.ExhibitionListDataResultMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MistTrialSettleNotify_proto_init() }
func file_MistTrialSettleNotify_proto_init() {
	if File_MistTrialSettleNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MistTrialSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MistTrialSettleNotify); i {
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
			RawDescriptor: file_MistTrialSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MistTrialSettleNotify_proto_goTypes,
		DependencyIndexes: file_MistTrialSettleNotify_proto_depIdxs,
		MessageInfos:      file_MistTrialSettleNotify_proto_msgTypes,
	}.Build()
	File_MistTrialSettleNotify_proto = out.File
	file_MistTrialSettleNotify_proto_rawDesc = nil
	file_MistTrialSettleNotify_proto_goTypes = nil
	file_MistTrialSettleNotify_proto_depIdxs = nil
}