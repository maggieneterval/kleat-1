// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: canary/google.proto

package canary

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

// Configuration for the Google canary integration.
type Google struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the integration is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*GoogleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// Whether GCS is enabled as a persistent store.
	GcsEnabled bool `protobuf:"varint,3,opt,name=gcsEnabled,proto3" json:"gcsEnabled,omitempty"`
	// Whether Stackdriver is enabled as a metrics source.
	StackdriverEnabled bool `protobuf:"varint,4,opt,name=stackdriverEnabled,proto3" json:"stackdriverEnabled,omitempty"`
	// Number of milliseconds to wait in between caching the names of available
	// Stackdriver metric types (used when building canary configs). Defaults to
	// 60000.
	MetadataCachingIntervalMS int32 `protobuf:"varint,5,opt,name=metadataCachingIntervalMS,proto3" json:"metadataCachingIntervalMS,omitempty"`
}

func (x *Google) Reset() {
	*x = Google{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_google_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Google) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Google) ProtoMessage() {}

func (x *Google) ProtoReflect() protoreflect.Message {
	mi := &file_canary_google_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Google.ProtoReflect.Descriptor instead.
func (*Google) Descriptor() ([]byte, []int) {
	return file_canary_google_proto_rawDescGZIP(), []int{0}
}

func (x *Google) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Google) GetAccounts() []*GoogleAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Google) GetGcsEnabled() bool {
	if x != nil {
		return x.GcsEnabled
	}
	return false
}

func (x *Google) GetStackdriverEnabled() bool {
	if x != nil {
		return x.StackdriverEnabled
	}
	return false
}

func (x *Google) GetMetadataCachingIntervalMS() int32 {
	if x != nil {
		return x.MetadataCachingIntervalMS
	}
	return 0
}

// Configuration for a Google account.
type GoogleAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The path to a JSON service account that Spinnaker will use as credentials.
	// This is only needed if Spinnaker is not deployed on a Google Compute Engine
	// VM, or needs permissions not afforded to the VM it is running on. See
	// https://cloud.google.com/compute/docs/access/service-accounts for more information.
	JsonPath string `protobuf:"bytes,2,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// The name of a storage bucket that this account has access to. If you
	// specify a globally unique bucket name that doesn't exist yet, Kayenta will
	// create that bucket for you.
	Bucket string `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// This is only required if the bucket you specify doesn’t exist yet. In that
	// case, the bucket will be created in that location. See
	// https://cloud.google.com/storage/docs/managing-buckets#manage-class-location.
	BucketLocation string `protobuf:"bytes,4,opt,name=bucketLocation,proto3" json:"bucketLocation,omitempty"`
	// The root folder in the chosen bucket in which to store all of the canary
	// service's persistent data in. Defaults to `kayenta`.
	RootFolder string `protobuf:"bytes,5,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	// (Required) The Google Cloud Platform project the canary service will use to
	// consume GCS and Stackdriver.
	Project string `protobuf:"bytes,6,opt,name=project,proto3" json:"project,omitempty"`
	// If enabling Stackdriver, include METRICS_STORE in this list.
	// If enabling GCS, include CONFIGURATION_STORE and/or OBJECT_STORE in this
	// list.
	SupportedTypes []SupportedType `protobuf:"varint,7,rep,packed,name=supportedTypes,proto3,enum=proto.canary.SupportedType" json:"supportedTypes,omitempty"`
}

func (x *GoogleAccount) Reset() {
	*x = GoogleAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_google_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAccount) ProtoMessage() {}

func (x *GoogleAccount) ProtoReflect() protoreflect.Message {
	mi := &file_canary_google_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAccount.ProtoReflect.Descriptor instead.
func (*GoogleAccount) Descriptor() ([]byte, []int) {
	return file_canary_google_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleAccount) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *GoogleAccount) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GoogleAccount) GetBucketLocation() string {
	if x != nil {
		return x.BucketLocation
	}
	return ""
}

func (x *GoogleAccount) GetRootFolder() string {
	if x != nil {
		return x.RootFolder
	}
	return ""
}

func (x *GoogleAccount) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GoogleAccount) GetSupportedTypes() []SupportedType {
	if x != nil {
		return x.SupportedTypes
	}
	return nil
}

var File_canary_google_proto protoreflect.FileDescriptor

var file_canary_google_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x1a, 0x1b, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x01, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x63, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x67, 0x63, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3c,
	0x0a, 0x19, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x19, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x53, 0x22, 0xfe, 0x01, 0x0a,
	0x0d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_canary_google_proto_rawDescOnce sync.Once
	file_canary_google_proto_rawDescData = file_canary_google_proto_rawDesc
)

func file_canary_google_proto_rawDescGZIP() []byte {
	file_canary_google_proto_rawDescOnce.Do(func() {
		file_canary_google_proto_rawDescData = protoimpl.X.CompressGZIP(file_canary_google_proto_rawDescData)
	})
	return file_canary_google_proto_rawDescData
}

var file_canary_google_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_canary_google_proto_goTypes = []interface{}{
	(*Google)(nil),        // 0: proto.canary.Google
	(*GoogleAccount)(nil), // 1: proto.canary.GoogleAccount
	(SupportedType)(0),    // 2: proto.canary.SupportedType
}
var file_canary_google_proto_depIdxs = []int32{
	1, // 0: proto.canary.Google.accounts:type_name -> proto.canary.GoogleAccount
	2, // 1: proto.canary.GoogleAccount.supportedTypes:type_name -> proto.canary.SupportedType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_canary_google_proto_init() }
func file_canary_google_proto_init() {
	if File_canary_google_proto != nil {
		return
	}
	file_canary_supported_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_canary_google_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Google); i {
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
		file_canary_google_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAccount); i {
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
			RawDescriptor: file_canary_google_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_canary_google_proto_goTypes,
		DependencyIndexes: file_canary_google_proto_depIdxs,
		MessageInfos:      file_canary_google_proto_msgTypes,
	}.Build()
	File_canary_google_proto = out.File
	file_canary_google_proto_rawDesc = nil
	file_canary_google_proto_goTypes = nil
	file_canary_google_proto_depIdxs = nil
}
