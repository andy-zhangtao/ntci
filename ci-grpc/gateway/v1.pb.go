// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

package gateway_rpc_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuildRequest struct {
	//    Build User
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	//    Build Name
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildRequest) Reset()         { *m = BuildRequest{} }
func (m *BuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()    {}
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{0}
}

func (m *BuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRequest.Unmarshal(m, b)
}
func (m *BuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRequest.Merge(m, src)
}
func (m *BuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildRequest.Size(m)
}
func (m *BuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRequest proto.InternalMessageInfo

func (m *BuildRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BuildRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// JobDetail
// Job Detail Info
type JobDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Branch               string   `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Id                   int32    `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	Sha                  string   `protobuf:"bytes,7,opt,name=sha,proto3" json:"sha,omitempty"`
	Message              string   `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobDetail) Reset()         { *m = JobDetail{} }
func (m *JobDetail) String() string { return proto.CompactTextString(m) }
func (*JobDetail) ProtoMessage()    {}
func (*JobDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{1}
}

func (m *JobDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobDetail.Unmarshal(m, b)
}
func (m *JobDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobDetail.Marshal(b, m, deterministic)
}
func (m *JobDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDetail.Merge(m, src)
}
func (m *JobDetail) XXX_Size() int {
	return xxx_messageInfo_JobDetail.Size(m)
}
func (m *JobDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDetail.DiscardUnknown(m)
}

var xxx_messageInfo_JobDetail proto.InternalMessageInfo

func (m *JobDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobDetail) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *JobDetail) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *JobDetail) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *JobDetail) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *JobDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *JobDetail) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *JobDetail) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type JobInfo struct {
	Count                int32        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Jd                   []*JobDetail `protobuf:"bytes,2,rep,name=jd,proto3" json:"jd,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}
func (*JobInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{2}
}

func (m *JobInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobInfo.Unmarshal(m, b)
}
func (m *JobInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobInfo.Marshal(b, m, deterministic)
}
func (m *JobInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobInfo.Merge(m, src)
}
func (m *JobInfo) XXX_Size() int {
	return xxx_messageInfo_JobInfo.Size(m)
}
func (m *JobInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobInfo proto.InternalMessageInfo

func (m *JobInfo) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *JobInfo) GetJd() []*JobDetail {
	if m != nil {
		return m.Jd
	}
	return nil
}

// Builder will update status via this message.
type Builder struct {
	// The build job name
	Jname string `protobuf:"bytes,1,opt,name=jname,proto3" json:"jname,omitempty"`
	// The build job id
	Jid string `protobuf:"bytes,2,opt,name=jid,proto3" json:"jid,omitempty"`
	// Job Status
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// Job Owner
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder) Reset()         { *m = Builder{} }
func (m *Builder) String() string { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()    {}
func (*Builder) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{3}
}

func (m *Builder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Builder.Unmarshal(m, b)
}
func (m *Builder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Builder.Marshal(b, m, deterministic)
}
func (m *Builder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Builder.Merge(m, src)
}
func (m *Builder) XXX_Size() int {
	return xxx_messageInfo_Builder.Size(m)
}
func (m *Builder) XXX_DiscardUnknown() {
	xxx_messageInfo_Builder.DiscardUnknown(m)
}

var xxx_messageInfo_Builder proto.InternalMessageInfo

func (m *Builder) GetJname() string {
	if m != nil {
		return m.Jname
	}
	return ""
}

func (m *Builder) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *Builder) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Builder) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type Reply struct {
	// 0 - success
	// other - failed
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Empty when success, otherwise there will be error message
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{4}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "BuildRequest")
	proto.RegisterType((*JobDetail)(nil), "JobDetail")
	proto.RegisterType((*JobInfo)(nil), "JobInfo")
	proto.RegisterType((*Builder)(nil), "Builder")
	proto.RegisterType((*Reply)(nil), "Reply")
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor_2e4aa7d76fd7ee8a) }

var fileDescriptor_2e4aa7d76fd7ee8a = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x14, 0x8c, 0xed, 0xf8, 0x23, 0x6f, 0x77, 0x43, 0x10, 0x4b, 0x11, 0xa1, 0xd0, 0x60, 0x28, 0xe4,
	0x24, 0x48, 0x4a, 0x7b, 0xe9, 0x2d, 0x14, 0x42, 0x73, 0x54, 0x0f, 0x85, 0x42, 0x0f, 0xb2, 0xad,
	0x26, 0x36, 0xb6, 0xe5, 0x5a, 0x72, 0x4a, 0x7e, 0x5b, 0xff, 0x5c, 0x91, 0xec, 0x24, 0x0e, 0xf4,
	0xf6, 0x66, 0x34, 0xcf, 0xcc, 0xcc, 0x33, 0x04, 0xfb, 0x05, 0xa9, 0x6a, 0xa1, 0x44, 0xf8, 0x00,
	0x7f, 0x57, 0x4d, 0x9a, 0x27, 0x94, 0x7f, 0x36, 0x5c, 0x2a, 0x84, 0x60, 0xd8, 0x48, 0x5e, 0x63,
	0x6b, 0x66, 0xcd, 0x47, 0xd4, 0xcc, 0x9a, 0x2b, 0x59, 0xc1, 0xb1, 0xdd, 0x72, 0x7a, 0x0e, 0xbf,
	0x2d, 0x18, 0x6d, 0x44, 0xf4, 0xc4, 0x15, 0x4b, 0xf3, 0x93, 0xc2, 0x3a, 0x2b, 0xd0, 0x15, 0x78,
	0x52, 0x31, 0xd5, 0x48, 0xb3, 0xe7, 0xd2, 0x0e, 0xa1, 0x6b, 0x18, 0xa9, 0xb4, 0xe0, 0x52, 0xb1,
	0xa2, 0xc2, 0x8e, 0x59, 0x38, 0x13, 0x7a, 0x2b, 0xaa, 0x59, 0x19, 0xef, 0xf0, 0xd0, 0x3c, 0x75,
	0x08, 0x4d, 0xc0, 0x69, 0xea, 0x1c, 0xbb, 0x86, 0xd4, 0x23, 0x1a, 0x83, 0x9d, 0x26, 0xd8, 0x33,
	0xdf, 0xb6, 0xd3, 0x44, 0x2b, 0xe4, 0x8e, 0x61, 0xbf, 0x55, 0xc8, 0x1d, 0x43, 0x18, 0xfc, 0x82,
	0x4b, 0xc9, 0xb6, 0x1c, 0x07, 0x86, 0x3d, 0xc2, 0xf0, 0x11, 0xfc, 0x8d, 0x88, 0x9e, 0xcb, 0x0f,
	0x81, 0xfe, 0x83, 0x1b, 0x8b, 0xa6, 0x54, 0xc6, 0xbb, 0x4b, 0x5b, 0x80, 0xa6, 0x60, 0x67, 0x09,
	0xb6, 0x67, 0xce, 0xfc, 0xcf, 0x12, 0xc8, 0x29, 0x28, 0xb5, 0xb3, 0x24, 0x7c, 0x07, 0xdf, 0x54,
	0xc6, 0x6b, 0xbd, 0x9c, 0xf5, 0x82, 0xb7, 0x40, 0x3b, 0xc9, 0xd2, 0xa4, 0xab, 0x4b, 0x8f, 0xbd,
	0x2e, 0x9c, 0x8b, 0x2e, 0x8e, 0x6d, 0x0f, 0xcf, 0x6d, 0x87, 0xf7, 0xe0, 0x52, 0x5e, 0xe5, 0x07,
	0xfd, 0x18, 0x8b, 0x84, 0x77, 0xc6, 0xcc, 0xdc, 0x8f, 0x64, 0x5f, 0x44, 0x5a, 0xee, 0x01, 0xd6,
	0x4c, 0xf1, 0x57, 0x76, 0xa0, 0x55, 0x8c, 0x6e, 0x21, 0x58, 0x73, 0x65, 0x6c, 0xa2, 0x7f, 0xa4,
	0x7f, 0xe1, 0x69, 0x40, 0xba, 0xe8, 0xe1, 0x00, 0xdd, 0x98, 0x23, 0xbe, 0xb4, 0x66, 0x02, 0xd2,
	0xc5, 0x9a, 0x7a, 0xc4, 0x38, 0x08, 0x07, 0x68, 0x06, 0x40, 0xf5, 0x65, 0x6a, 0xb5, 0x11, 0xd1,
	0x6f, 0x8a, 0xd5, 0xe4, 0x6d, 0xbc, 0x65, 0x8a, 0x7f, 0xb1, 0x03, 0xa9, 0xab, 0x98, 0xec, 0x17,
	0x91, 0x67, 0xfe, 0xac, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x55, 0x3a, 0xfb, 0x65,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateWayRpcClient is the client API for GateWayRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateWayRpcClient interface {
	// GetBuild Query User's All Builds
	GetBuild(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// Builder update status via this rpc.
	JobStatus(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error)
	// Restart Specify Job
	RestartJob(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error)
}

type gateWayRpcClient struct {
	cc *grpc.ClientConn
}

func NewGateWayRpcClient(cc *grpc.ClientConn) GateWayRpcClient {
	return &gateWayRpcClient{cc}
}

func (c *gateWayRpcClient) GetBuild(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, "/GateWayRpc/GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateWayRpcClient) JobStatus(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/GateWayRpc/JobStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateWayRpcClient) RestartJob(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/GateWayRpc/RestartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateWayRpcServer is the server API for GateWayRpc service.
type GateWayRpcServer interface {
	// GetBuild Query User's All Builds
	GetBuild(context.Context, *BuildRequest) (*JobInfo, error)
	// Builder update status via this rpc.
	JobStatus(context.Context, *Builder) (*Reply, error)
	// Restart Specify Job
	RestartJob(context.Context, *Builder) (*Reply, error)
}

// UnimplementedGateWayRpcServer can be embedded to have forward compatible implementations.
type UnimplementedGateWayRpcServer struct {
}

func (*UnimplementedGateWayRpcServer) GetBuild(ctx context.Context, req *BuildRequest) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuild not implemented")
}
func (*UnimplementedGateWayRpcServer) JobStatus(ctx context.Context, req *Builder) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobStatus not implemented")
}
func (*UnimplementedGateWayRpcServer) RestartJob(ctx context.Context, req *Builder) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartJob not implemented")
}

func RegisterGateWayRpcServer(s *grpc.Server, srv GateWayRpcServer) {
	s.RegisterService(&_GateWayRpc_serviceDesc, srv)
}

func _GateWayRpc_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateWayRpcServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GateWayRpc/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateWayRpcServer).GetBuild(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateWayRpc_JobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Builder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateWayRpcServer).JobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GateWayRpc/JobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateWayRpcServer).JobStatus(ctx, req.(*Builder))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateWayRpc_RestartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Builder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateWayRpcServer).RestartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GateWayRpc/RestartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateWayRpcServer).RestartJob(ctx, req.(*Builder))
	}
	return interceptor(ctx, in, info, handler)
}

var _GateWayRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GateWayRpc",
	HandlerType: (*GateWayRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuild",
			Handler:    _GateWayRpc_GetBuild_Handler,
		},
		{
			MethodName: "JobStatus",
			Handler:    _GateWayRpc_JobStatus_Handler,
		},
		{
			MethodName: "RestartJob",
			Handler:    _GateWayRpc_RestartJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1.proto",
}
