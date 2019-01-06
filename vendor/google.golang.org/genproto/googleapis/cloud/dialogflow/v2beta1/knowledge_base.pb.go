// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/knowledge_base.proto

package dialogflow // import "google.golang.org/genproto/googleapis/cloud/dialogflow/v2beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents knowledge base resource.
type KnowledgeBase struct {
	// The knowledge base resource name.
	// The name must be empty when creating a knowledge base.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the knowledge base. The name must be 1024
	// bytes or less; otherwise, the creation request fails.
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KnowledgeBase) Reset()         { *m = KnowledgeBase{} }
func (m *KnowledgeBase) String() string { return proto.CompactTextString(m) }
func (*KnowledgeBase) ProtoMessage()    {}
func (*KnowledgeBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{0}
}
func (m *KnowledgeBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnowledgeBase.Unmarshal(m, b)
}
func (m *KnowledgeBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnowledgeBase.Marshal(b, m, deterministic)
}
func (dst *KnowledgeBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnowledgeBase.Merge(dst, src)
}
func (m *KnowledgeBase) XXX_Size() int {
	return xxx_messageInfo_KnowledgeBase.Size(m)
}
func (m *KnowledgeBase) XXX_DiscardUnknown() {
	xxx_messageInfo_KnowledgeBase.DiscardUnknown(m)
}

var xxx_messageInfo_KnowledgeBase proto.InternalMessageInfo

func (m *KnowledgeBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KnowledgeBase) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// Request message for [KnowledgeBases.ListKnowledgeBases][google.cloud.dialogflow.v2beta1.KnowledgeBases.ListKnowledgeBases].
type ListKnowledgeBasesRequest struct {
	// Required. The agent to list of knowledge bases for.
	// Format: `projects/<Project ID>/agent`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 10 and at most 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKnowledgeBasesRequest) Reset()         { *m = ListKnowledgeBasesRequest{} }
func (m *ListKnowledgeBasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListKnowledgeBasesRequest) ProtoMessage()    {}
func (*ListKnowledgeBasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{1}
}
func (m *ListKnowledgeBasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Unmarshal(m, b)
}
func (m *ListKnowledgeBasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Marshal(b, m, deterministic)
}
func (dst *ListKnowledgeBasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKnowledgeBasesRequest.Merge(dst, src)
}
func (m *ListKnowledgeBasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Size(m)
}
func (m *ListKnowledgeBasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKnowledgeBasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListKnowledgeBasesRequest proto.InternalMessageInfo

func (m *ListKnowledgeBasesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListKnowledgeBasesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListKnowledgeBasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for [KnowledgeBases.ListKnowledgeBases][google.cloud.dialogflow.v2beta1.KnowledgeBases.ListKnowledgeBases].
type ListKnowledgeBasesResponse struct {
	// The list of knowledge bases.
	KnowledgeBases []*KnowledgeBase `protobuf:"bytes,1,rep,name=knowledge_bases,json=knowledgeBases,proto3" json:"knowledge_bases,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKnowledgeBasesResponse) Reset()         { *m = ListKnowledgeBasesResponse{} }
func (m *ListKnowledgeBasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListKnowledgeBasesResponse) ProtoMessage()    {}
func (*ListKnowledgeBasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{2}
}
func (m *ListKnowledgeBasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Unmarshal(m, b)
}
func (m *ListKnowledgeBasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Marshal(b, m, deterministic)
}
func (dst *ListKnowledgeBasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKnowledgeBasesResponse.Merge(dst, src)
}
func (m *ListKnowledgeBasesResponse) XXX_Size() int {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Size(m)
}
func (m *ListKnowledgeBasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKnowledgeBasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListKnowledgeBasesResponse proto.InternalMessageInfo

func (m *ListKnowledgeBasesResponse) GetKnowledgeBases() []*KnowledgeBase {
	if m != nil {
		return m.KnowledgeBases
	}
	return nil
}

func (m *ListKnowledgeBasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for [KnowledgeBase.GetDocument][].
type GetKnowledgeBaseRequest struct {
	// Required. The name of the knowledge base to retrieve.
	// Format `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKnowledgeBaseRequest) Reset()         { *m = GetKnowledgeBaseRequest{} }
func (m *GetKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetKnowledgeBaseRequest) ProtoMessage()    {}
func (*GetKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{3}
}
func (m *GetKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *GetKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (dst *GetKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKnowledgeBaseRequest.Merge(dst, src)
}
func (m *GetKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Size(m)
}
func (m *GetKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKnowledgeBaseRequest proto.InternalMessageInfo

func (m *GetKnowledgeBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for [KnowledgeBases.CreateKnowledgeBase][google.cloud.dialogflow.v2beta1.KnowledgeBases.CreateKnowledgeBase].
type CreateKnowledgeBaseRequest struct {
	// Required. The agent to create a knowledge base for.
	// Format: `projects/<Project ID>/agent`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The knowledge base to create.
	KnowledgeBase        *KnowledgeBase `protobuf:"bytes,2,opt,name=knowledge_base,json=knowledgeBase,proto3" json:"knowledge_base,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateKnowledgeBaseRequest) Reset()         { *m = CreateKnowledgeBaseRequest{} }
func (m *CreateKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateKnowledgeBaseRequest) ProtoMessage()    {}
func (*CreateKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{4}
}
func (m *CreateKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *CreateKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (dst *CreateKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateKnowledgeBaseRequest.Merge(dst, src)
}
func (m *CreateKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Size(m)
}
func (m *CreateKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateKnowledgeBaseRequest proto.InternalMessageInfo

func (m *CreateKnowledgeBaseRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateKnowledgeBaseRequest) GetKnowledgeBase() *KnowledgeBase {
	if m != nil {
		return m.KnowledgeBase
	}
	return nil
}

// Request message for [KnowledgeBases.DeleteKnowledgeBase][google.cloud.dialogflow.v2beta1.KnowledgeBases.DeleteKnowledgeBase].
type DeleteKnowledgeBaseRequest struct {
	// Required. The name of the knowledge base to delete.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Force deletes the knowledge base. When set to true, any documents
	// in the knowledge base are also deleted.
	Force                bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKnowledgeBaseRequest) Reset()         { *m = DeleteKnowledgeBaseRequest{} }
func (m *DeleteKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteKnowledgeBaseRequest) ProtoMessage()    {}
func (*DeleteKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowledge_base_5078794a2722dd36, []int{5}
}
func (m *DeleteKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *DeleteKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKnowledgeBaseRequest.Merge(dst, src)
}
func (m *DeleteKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Size(m)
}
func (m *DeleteKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKnowledgeBaseRequest proto.InternalMessageInfo

func (m *DeleteKnowledgeBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteKnowledgeBaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func init() {
	proto.RegisterType((*KnowledgeBase)(nil), "google.cloud.dialogflow.v2beta1.KnowledgeBase")
	proto.RegisterType((*ListKnowledgeBasesRequest)(nil), "google.cloud.dialogflow.v2beta1.ListKnowledgeBasesRequest")
	proto.RegisterType((*ListKnowledgeBasesResponse)(nil), "google.cloud.dialogflow.v2beta1.ListKnowledgeBasesResponse")
	proto.RegisterType((*GetKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.GetKnowledgeBaseRequest")
	proto.RegisterType((*CreateKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.CreateKnowledgeBaseRequest")
	proto.RegisterType((*DeleteKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.DeleteKnowledgeBaseRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KnowledgeBasesClient is the client API for KnowledgeBases service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KnowledgeBasesClient interface {
	// Returns the list of all knowledge bases of the specified agent.
	ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error)
	// Retrieves the specified knowledge base.
	GetKnowledgeBase(ctx context.Context, in *GetKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error)
	// Creates a knowledge base.
	CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error)
	// Deletes the specified knowledge base.
	DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type knowledgeBasesClient struct {
	cc *grpc.ClientConn
}

func NewKnowledgeBasesClient(cc *grpc.ClientConn) KnowledgeBasesClient {
	return &knowledgeBasesClient{cc}
}

func (c *knowledgeBasesClient) ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error) {
	out := new(ListKnowledgeBasesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/ListKnowledgeBases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) GetKnowledgeBase(ctx context.Context, in *GetKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error) {
	out := new(KnowledgeBase)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/GetKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error) {
	out := new(KnowledgeBase)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/CreateKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/DeleteKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnowledgeBasesServer is the server API for KnowledgeBases service.
type KnowledgeBasesServer interface {
	// Returns the list of all knowledge bases of the specified agent.
	ListKnowledgeBases(context.Context, *ListKnowledgeBasesRequest) (*ListKnowledgeBasesResponse, error)
	// Retrieves the specified knowledge base.
	GetKnowledgeBase(context.Context, *GetKnowledgeBaseRequest) (*KnowledgeBase, error)
	// Creates a knowledge base.
	CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*KnowledgeBase, error)
	// Deletes the specified knowledge base.
	DeleteKnowledgeBase(context.Context, *DeleteKnowledgeBaseRequest) (*empty.Empty, error)
}

func RegisterKnowledgeBasesServer(s *grpc.Server, srv KnowledgeBasesServer) {
	s.RegisterService(&_KnowledgeBases_serviceDesc, srv)
}

func _KnowledgeBases_ListKnowledgeBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKnowledgeBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).ListKnowledgeBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/ListKnowledgeBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).ListKnowledgeBases(ctx, req.(*ListKnowledgeBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_GetKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).GetKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/GetKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).GetKnowledgeBase(ctx, req.(*GetKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_CreateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).CreateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/CreateKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).CreateKnowledgeBase(ctx, req.(*CreateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_DeleteKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).DeleteKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/DeleteKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).DeleteKnowledgeBase(ctx, req.(*DeleteKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KnowledgeBases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.KnowledgeBases",
	HandlerType: (*KnowledgeBasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKnowledgeBases",
			Handler:    _KnowledgeBases_ListKnowledgeBases_Handler,
		},
		{
			MethodName: "GetKnowledgeBase",
			Handler:    _KnowledgeBases_GetKnowledgeBase_Handler,
		},
		{
			MethodName: "CreateKnowledgeBase",
			Handler:    _KnowledgeBases_CreateKnowledgeBase_Handler,
		},
		{
			MethodName: "DeleteKnowledgeBase",
			Handler:    _KnowledgeBases_DeleteKnowledgeBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/knowledge_base.proto",
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/knowledge_base.proto", fileDescriptor_knowledge_base_5078794a2722dd36)
}

var fileDescriptor_knowledge_base_5078794a2722dd36 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0xce, 0x2c, 0x82, 0x30, 0x08, 0x98, 0xc1, 0x20, 0x16, 0x0d, 0x58, 0x13, 0x43, 0xd6, 0xd8,
	0x09, 0xe0, 0xc1, 0x40, 0xbc, 0x00, 0x42, 0x8c, 0xc6, 0x90, 0xf5, 0x57, 0xc2, 0x65, 0x33, 0xbb,
	0xfb, 0xb6, 0xd6, 0xed, 0xce, 0xd4, 0xce, 0x20, 0x82, 0xe1, 0xc2, 0xcd, 0x98, 0x78, 0xf1, 0xec,
	0xc5, 0xa3, 0x27, 0xff, 0x17, 0xbd, 0x79, 0x35, 0xfe, 0x0b, 0x7a, 0x34, 0x9d, 0xe9, 0x02, 0x53,
	0x5a, 0xd7, 0x7a, 0xdb, 0x79, 0x6f, 0xbe, 0x37, 0xdf, 0xf7, 0xbe, 0xf7, 0xba, 0xf8, 0x96, 0x2f,
	0x84, 0x1f, 0x02, 0x6d, 0x86, 0x62, 0xa7, 0x45, 0x5b, 0x01, 0x0b, 0x85, 0xdf, 0x0e, 0xc5, 0x2e,
	0x7d, 0xb5, 0xd8, 0x00, 0xc5, 0x16, 0x68, 0x87, 0x8b, 0xdd, 0x10, 0x5a, 0x3e, 0xd4, 0x1b, 0x4c,
	0x82, 0x17, 0xc5, 0x42, 0x09, 0x32, 0x6b, 0x50, 0x9e, 0x46, 0x79, 0xc7, 0x28, 0x2f, 0x45, 0x39,
	0x97, 0xd3, 0xb2, 0x2c, 0x0a, 0x28, 0xe3, 0x5c, 0x28, 0xa6, 0x02, 0xc1, 0xa5, 0x81, 0x3b, 0x33,
	0x69, 0x56, 0x9f, 0x1a, 0x3b, 0x6d, 0x0a, 0xdd, 0x48, 0xed, 0xa5, 0xc9, 0xb9, 0x6c, 0xb2, 0x1d,
	0x40, 0xd8, 0xaa, 0x77, 0x99, 0xec, 0x98, 0x1b, 0xee, 0x06, 0x1e, 0xbb, 0xdf, 0x63, 0xb5, 0xca,
	0x24, 0x10, 0x82, 0xcf, 0x70, 0xd6, 0x85, 0x69, 0x34, 0x87, 0xe6, 0x47, 0x6a, 0xfa, 0x37, 0xb9,
	0x8a, 0xcf, 0xb5, 0x02, 0x19, 0x85, 0x6c, 0xaf, 0xae, 0x73, 0x15, 0x9d, 0x1b, 0x4d, 0x63, 0x0f,
	0x59, 0x17, 0x5c, 0x81, 0x2f, 0x3d, 0x08, 0xa4, 0xb2, 0x6a, 0xc9, 0x1a, 0xbc, 0xdc, 0x01, 0xa9,
	0xc8, 0x14, 0x1e, 0x8a, 0x58, 0x0c, 0x5c, 0xa5, 0x55, 0xd3, 0x13, 0x99, 0xc1, 0x23, 0x11, 0xf3,
	0xa1, 0x2e, 0x83, 0x7d, 0x53, 0x74, 0xb0, 0x36, 0x9c, 0x04, 0x1e, 0x05, 0xfb, 0x40, 0xae, 0x60,
	0xac, 0x93, 0x4a, 0x74, 0x80, 0x4f, 0x0f, 0x68, 0xa0, 0xbe, 0xfe, 0x38, 0x09, 0xb8, 0x1f, 0x11,
	0x76, 0xf2, 0x5e, 0x94, 0x91, 0xe0, 0x12, 0xc8, 0x33, 0x3c, 0x61, 0x77, 0x5b, 0x4e, 0xa3, 0xb9,
	0x81, 0xf9, 0xd1, 0x45, 0xcf, 0xeb, 0xd3, 0x6f, 0xcf, 0xaa, 0x58, 0x1b, 0xef, 0x58, 0x0f, 0x90,
	0xeb, 0x78, 0x82, 0xc3, 0x6b, 0x55, 0x3f, 0xc1, 0xcd, 0xb4, 0x63, 0x2c, 0x09, 0x6f, 0x1d, 0xf1,
	0xbb, 0x89, 0x2f, 0x6e, 0x82, 0xcd, 0xae, 0xd7, 0x8e, 0x9c, 0x16, 0xbb, 0xef, 0x10, 0x76, 0xd6,
	0x62, 0x60, 0x0a, 0x72, 0x21, 0x45, 0x1d, 0x7c, 0x82, 0xc7, 0x6d, 0x99, 0x9a, 0x4c, 0x79, 0x95,
	0x63, 0x96, 0x4a, 0x77, 0x03, 0x3b, 0xeb, 0x10, 0x42, 0x01, 0x99, 0xbc, 0x11, 0xb9, 0x80, 0x07,
	0xdb, 0x22, 0x6e, 0x9a, 0xf7, 0x87, 0x6b, 0xe6, 0xb0, 0xf8, 0xfd, 0x2c, 0x1e, 0xb7, 0x0d, 0x22,
	0xbf, 0x10, 0x26, 0xa7, 0x7d, 0x23, 0xcb, 0x7d, 0x09, 0x17, 0x8e, 0x97, 0xb3, 0xf2, 0x5f, 0x58,
	0x33, 0x28, 0xee, 0xf3, 0xc3, 0xaf, 0x3f, 0x3e, 0x54, 0x1a, 0xe4, 0xc6, 0xd1, 0x96, 0xbe, 0x31,
	0xbd, 0xbd, 0x13, 0xc5, 0xe2, 0x05, 0x34, 0x95, 0xa4, 0xd5, 0x03, 0x6a, 0x0f, 0xc1, 0xf6, 0x12,
	0x59, 0xf8, 0xcb, 0x75, 0xca, 0x7c, 0xe0, 0x2a, 0x0b, 0x22, 0x3f, 0x11, 0x3e, 0x9f, 0x1d, 0x09,
	0x72, 0xbb, 0x2f, 0xf7, 0x82, 0x29, 0x72, 0x4a, 0x5a, 0x9c, 0x27, 0x34, 0x71, 0xee, 0x24, 0x6f,
	0x9b, 0x30, 0xad, 0x1e, 0xd8, 0x42, 0xb3, 0xd7, 0xb5, 0xcc, 0xd3, 0x20, 0xf2, 0xbe, 0x82, 0x27,
	0x73, 0x66, 0x99, 0xf4, 0xf7, 0xa9, 0x78, 0x03, 0x4a, 0xcb, 0x7d, 0x8b, 0xb4, 0xde, 0x43, 0xe4,
	0x96, 0x71, 0x76, 0x39, 0xb3, 0x4e, 0xdb, 0x6b, 0x6e, 0x79, 0xa7, 0xb3, 0x45, 0xc8, 0x37, 0x84,
	0x27, 0x73, 0xf6, 0xe9, 0x1f, 0x1a, 0x52, 0xbc, 0x85, 0xce, 0x54, 0x0f, 0xdc, 0xfb, 0xb8, 0x7b,
	0x77, 0x93, 0x2f, 0x7f, 0xcf, 0xe7, 0x6a, 0x39, 0x9f, 0xab, 0xe5, 0x7d, 0x5e, 0xfd, 0x82, 0xf0,
	0xb5, 0xa6, 0xe8, 0xf6, 0x13, 0xb1, 0x4a, 0x2c, 0xfe, 0x5b, 0x09, 0xdd, 0x2d, 0xb4, 0x7d, 0x2f,
	0x85, 0xf9, 0x22, 0x64, 0xdc, 0xf7, 0x44, 0xec, 0x53, 0x1f, 0xb8, 0x16, 0x43, 0x4d, 0x8a, 0x45,
	0x81, 0x2c, 0xfc, 0x33, 0x5d, 0x39, 0x0e, 0xfd, 0x46, 0xe8, 0x53, 0xa5, 0xb2, 0xbe, 0xf1, 0xb9,
	0x32, 0xbb, 0x69, 0x6a, 0xae, 0x69, 0x2a, 0xeb, 0xc7, 0x54, 0x9e, 0x1a, 0x50, 0x63, 0x48, 0xd7,
	0x5f, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xad, 0x8b, 0xb2, 0x9f, 0xab, 0x07, 0x00, 0x00,
}