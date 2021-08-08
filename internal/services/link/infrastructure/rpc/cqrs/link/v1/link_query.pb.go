// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: infrastructure/rpc/cqrs/link/v1/link_query.proto

package v1

import (
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *v1.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetLink() *v1.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links *v1.Links `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetLinks() *v1.Links {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_infrastructure_rpc_cqrs_link_v1_link_query_proto protoreflect.FileDescriptor

var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x3b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x32, 0x76, 0x0a,
	0x10, 0x4c, 0x69, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x62, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71,
	0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescOnce sync.Once
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData = file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc
)

func file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP() []byte {
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescOnce.Do(func() {
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData)
	})
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData
}

var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: infrastructure.rpc.cqrs.link.v1.GetRequest
	(*GetResponse)(nil),  // 1: infrastructure.rpc.cqrs.link.v1.GetResponse
	(*ListRequest)(nil),  // 2: infrastructure.rpc.cqrs.link.v1.ListRequest
	(*ListResponse)(nil), // 3: infrastructure.rpc.cqrs.link.v1.ListResponse
	(*v1.Link)(nil),      // 4: domain.link.v1.Link
	(*v1.Links)(nil),     // 5: domain.link.v1.Links
}
var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs = []int32{
	4, // 0: infrastructure.rpc.cqrs.link.v1.GetResponse.link:type_name -> domain.link.v1.Link
	5, // 1: infrastructure.rpc.cqrs.link.v1.ListResponse.links:type_name -> domain.link.v1.Links
	0, // 2: infrastructure.rpc.cqrs.link.v1.LinkQueryService.Get:input_type -> infrastructure.rpc.cqrs.link.v1.GetRequest
	1, // 3: infrastructure.rpc.cqrs.link.v1.LinkQueryService.Get:output_type -> infrastructure.rpc.cqrs.link.v1.GetResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infrastructure_rpc_cqrs_link_v1_link_query_proto_init() }
func file_infrastructure_rpc_cqrs_link_v1_link_query_proto_init() {
	if File_infrastructure_rpc_cqrs_link_v1_link_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes,
		DependencyIndexes: file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs,
		MessageInfos:      file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes,
	}.Build()
	File_infrastructure_rpc_cqrs_link_v1_link_query_proto = out.File
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc = nil
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes = nil
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs = nil
}
