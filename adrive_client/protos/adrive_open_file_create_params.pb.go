// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: protos/adrive_client/adrive_open_file_create_params.proto

package protos

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

type AdriveOpenFileCreateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriveId         string                      `protobuf:"bytes,1,opt,name=drive_id,json=driveId,proto3" json:"drive_id,omitempty"`
	ParentFileId    string                      `protobuf:"bytes,2,opt,name=parent_file_id,json=parentFileId,proto3" json:"parent_file_id,omitempty"`
	Name            string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type            string                      `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	CheckNameMode   string                      `protobuf:"bytes,5,opt,name=check_name_mode,json=checkNameMode,proto3" json:"check_name_mode,omitempty"`
	PartInfoList    []*AdriveOpenFilePartInfo   `protobuf:"bytes,6,rep,name=part_info_list,json=partInfoList,proto3" json:"part_info_list,omitempty"`
	StreamsInfo     []*AdriveOpenFileStreamInfo `protobuf:"bytes,7,rep,name=streams_info,json=streamsInfo,proto3" json:"streams_info,omitempty"`
	PreHash         string                      `protobuf:"bytes,8,opt,name=pre_hash,json=preHash,proto3" json:"pre_hash,omitempty"`
	Size            int64                       `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	ContentHash     string                      `protobuf:"bytes,10,opt,name=content_hash,json=contentHash,proto3" json:"content_hash,omitempty"`
	ContentHashName string                      `protobuf:"bytes,11,opt,name=content_hash_name,json=contentHashName,proto3" json:"content_hash_name,omitempty"`
	ProofCode       string                      `protobuf:"bytes,12,opt,name=proof_code,json=proofCode,proto3" json:"proof_code,omitempty"`
	ProofVersion    string                      `protobuf:"bytes,13,opt,name=proof_version,json=proofVersion,proto3" json:"proof_version,omitempty"`
	LocalCreatedAt  string                      `protobuf:"bytes,14,opt,name=local_created_at,json=localCreatedAt,proto3" json:"local_created_at,omitempty"`
	LocalModifiedAt string                      `protobuf:"bytes,15,opt,name=local_modified_at,json=localModifiedAt,proto3" json:"local_modified_at,omitempty"`
}

func (x *AdriveOpenFileCreateParams) Reset() {
	*x = AdriveOpenFileCreateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_adrive_client_adrive_open_file_create_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdriveOpenFileCreateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdriveOpenFileCreateParams) ProtoMessage() {}

func (x *AdriveOpenFileCreateParams) ProtoReflect() protoreflect.Message {
	mi := &file_protos_adrive_client_adrive_open_file_create_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdriveOpenFileCreateParams.ProtoReflect.Descriptor instead.
func (*AdriveOpenFileCreateParams) Descriptor() ([]byte, []int) {
	return file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescGZIP(), []int{0}
}

func (x *AdriveOpenFileCreateParams) GetDriveId() string {
	if x != nil {
		return x.DriveId
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetParentFileId() string {
	if x != nil {
		return x.ParentFileId
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetCheckNameMode() string {
	if x != nil {
		return x.CheckNameMode
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetPartInfoList() []*AdriveOpenFilePartInfo {
	if x != nil {
		return x.PartInfoList
	}
	return nil
}

func (x *AdriveOpenFileCreateParams) GetStreamsInfo() []*AdriveOpenFileStreamInfo {
	if x != nil {
		return x.StreamsInfo
	}
	return nil
}

func (x *AdriveOpenFileCreateParams) GetPreHash() string {
	if x != nil {
		return x.PreHash
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AdriveOpenFileCreateParams) GetContentHash() string {
	if x != nil {
		return x.ContentHash
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetContentHashName() string {
	if x != nil {
		return x.ContentHashName
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetProofCode() string {
	if x != nil {
		return x.ProofCode
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetProofVersion() string {
	if x != nil {
		return x.ProofVersion
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetLocalCreatedAt() string {
	if x != nil {
		return x.LocalCreatedAt
	}
	return ""
}

func (x *AdriveOpenFileCreateParams) GetLocalModifiedAt() string {
	if x != nil {
		return x.LocalModifiedAt
	}
	return ""
}

var File_protos_adrive_client_adrive_open_file_create_params_proto protoreflect.FileDescriptor

var file_protos_adrive_client_adrive_open_file_create_params_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x1a, 0x41, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x41, 0x64, 0x72, 0x69, 0x76, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x41, 0x64, 0x72, 0x69, 0x76, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x42, 0x17, 0x5a,
	0x15, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescOnce sync.Once
	file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescData = file_protos_adrive_client_adrive_open_file_create_params_proto_rawDesc
)

func file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescGZIP() []byte {
	file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescOnce.Do(func() {
		file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescData)
	})
	return file_protos_adrive_client_adrive_open_file_create_params_proto_rawDescData
}

var file_protos_adrive_client_adrive_open_file_create_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_adrive_client_adrive_open_file_create_params_proto_goTypes = []interface{}{
	(*AdriveOpenFileCreateParams)(nil), // 0: AdriveOpenFileCreateParams
	(*AdriveOpenFilePartInfo)(nil),     // 1: AdriveOpenFilePartInfo
	(*AdriveOpenFileStreamInfo)(nil),   // 2: AdriveOpenFileStreamInfo
}
var file_protos_adrive_client_adrive_open_file_create_params_proto_depIdxs = []int32{
	1, // 0: AdriveOpenFileCreateParams.part_info_list:type_name -> AdriveOpenFilePartInfo
	2, // 1: AdriveOpenFileCreateParams.streams_info:type_name -> AdriveOpenFileStreamInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_adrive_client_adrive_open_file_create_params_proto_init() }
func file_protos_adrive_client_adrive_open_file_create_params_proto_init() {
	if File_protos_adrive_client_adrive_open_file_create_params_proto != nil {
		return
	}
	file_protos_adrive_client_adrive_open_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_adrive_client_adrive_open_file_create_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdriveOpenFileCreateParams); i {
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
			RawDescriptor: file_protos_adrive_client_adrive_open_file_create_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_adrive_client_adrive_open_file_create_params_proto_goTypes,
		DependencyIndexes: file_protos_adrive_client_adrive_open_file_create_params_proto_depIdxs,
		MessageInfos:      file_protos_adrive_client_adrive_open_file_create_params_proto_msgTypes,
	}.Build()
	File_protos_adrive_client_adrive_open_file_create_params_proto = out.File
	file_protos_adrive_client_adrive_open_file_create_params_proto_rawDesc = nil
	file_protos_adrive_client_adrive_open_file_create_params_proto_goTypes = nil
	file_protos_adrive_client_adrive_open_file_create_params_proto_depIdxs = nil
}
