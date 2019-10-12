// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

package build_rpc_v1

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

type Request struct {
	// project name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// event branch
	Branch string `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	// repository url
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// commit id
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// build language
	Language string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	// default is "latest"
	Lanversion string `protobuf:"bytes,6,opt,name=lanversion,proto3" json:"lanversion,omitempty"`
	// is query the latest build job. Only use in query actions.
	Latest               bool     `protobuf:"varint,7,opt,name=latest,proto3" json:"latest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Request) GetLanversion() string {
	if m != nil {
		return m.Lanversion
	}
	return ""
}

func (m *Request) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
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
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{1}
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

type JobDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Branch               string   `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobDetail) Reset()         { *m = JobDetail{} }
func (m *JobDetail) String() string { return proto.CompactTextString(m) }
func (*JobDetail) ProtoMessage()    {}
func (*JobDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{2}
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
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{3}
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

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Reply)(nil), "Reply")
	proto.RegisterType((*JobDetail)(nil), "JobDetail")
	proto.RegisterType((*JobInfo)(nil), "JobInfo")
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor_2e4aa7d76fd7ee8a) }

var fileDescriptor_2e4aa7d76fd7ee8a = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xe3, 0xff, 0xce, 0x6c, 0x08, 0x8b, 0x58, 0x16, 0xad, 0x59, 0x42, 0xf0, 0x29, 0x27,
	0x43, 0x52, 0x7a, 0xea, 0x2d, 0x14, 0x4a, 0x73, 0x2a, 0xea, 0xad, 0x37, 0xd9, 0x9e, 0xa6, 0x0a,
	0xb6, 0xe4, 0x5a, 0xb2, 0xa1, 0xa7, 0x7e, 0xa0, 0x7e, 0xc9, 0x62, 0xc5, 0x4d, 0x73, 0x48, 0x6f,
	0xf3, 0xde, 0x60, 0xde, 0x6f, 0x9e, 0x05, 0x71, 0xbf, 0xce, 0x9a, 0x56, 0x19, 0x95, 0x7e, 0x38,
	0x10, 0x31, 0x7c, 0xed, 0x50, 0x1b, 0x42, 0xc0, 0x97, 0xbc, 0x46, 0xea, 0x2c, 0x9d, 0xd5, 0x94,
	0xd9, 0x99, 0xfc, 0x85, 0x30, 0x6f, 0xb9, 0x2c, 0x5e, 0xa8, 0x6b, 0xdd, 0x51, 0x91, 0xdf, 0xe0,
	0x75, 0x6d, 0x45, 0x3d, 0x6b, 0x0e, 0x23, 0x99, 0x83, 0x2b, 0x4a, 0xea, 0x5b, 0xc3, 0x15, 0x25,
	0x49, 0x20, 0xae, 0xb8, 0xdc, 0x77, 0x7c, 0x8f, 0x34, 0xb0, 0xee, 0x49, 0x93, 0x05, 0x40, 0xc5,
	0x65, 0x8f, 0xad, 0x16, 0x4a, 0xd2, 0xd0, 0x6e, 0xcf, 0x9c, 0x21, 0xb5, 0xe2, 0x06, 0xb5, 0xa1,
	0xd1, 0xd2, 0x59, 0xc5, 0x6c, 0x54, 0xe9, 0x35, 0x04, 0x0c, 0x9b, 0xea, 0x6d, 0x40, 0x2d, 0x54,
	0x79, 0x44, 0x0d, 0x98, 0x9d, 0x09, 0x85, 0xa8, 0x46, 0xad, 0x87, 0xbc, 0x23, 0xeb, 0x97, 0x4c,
	0xdf, 0x61, 0xba, 0x53, 0xf9, 0x2d, 0x1a, 0x2e, 0xaa, 0x9f, 0xae, 0xd4, 0x86, 0x9b, 0x4e, 0xdb,
	0x2f, 0x03, 0x36, 0x2a, 0xf2, 0x1f, 0xa6, 0x46, 0xd4, 0xa8, 0x0d, 0xaf, 0x9b, 0xf1, 0xd6, 0x6f,
	0xe3, 0xac, 0x1b, 0xff, 0x52, 0x37, 0xc1, 0xa9, 0x9b, 0xf4, 0x06, 0xa2, 0x9d, 0xca, 0xef, 0xe5,
	0xb3, 0x22, 0x7f, 0x20, 0x28, 0x54, 0x27, 0xcd, 0x88, 0x7e, 0x14, 0x24, 0x01, 0xf7, 0x50, 0x52,
	0x77, 0xe9, 0xad, 0x7e, 0x6d, 0x20, 0x3b, 0xc1, 0x32, 0xf7, 0x50, 0x6e, 0x10, 0x66, 0xdb, 0x4e,
	0x54, 0xe5, 0x23, 0xb6, 0xbd, 0x28, 0x90, 0xfc, 0x03, 0x8f, 0x75, 0x92, 0xc4, 0xd9, 0xf8, 0xdf,
	0x92, 0x30, 0xb3, 0xa5, 0xa4, 0x13, 0x92, 0x80, 0xff, 0x20, 0xe4, 0xfe, 0xe2, 0x6e, 0x01, 0xe1,
	0x1d, 0x9a, 0x9d, 0xca, 0xcf, 0xb6, 0x71, 0x36, 0x62, 0xa5, 0x93, 0xed, 0xfc, 0x69, 0x96, 0x0f,
	0x31, 0x59, 0xdb, 0x14, 0x59, 0xbf, 0xce, 0x43, 0xfb, 0x40, 0xae, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0xd0, 0x26, 0x99, 0x2c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildServiceClient is the client API for BuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildServiceClient interface {
	Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	GetJob(ctx context.Context, in *Request, opts ...grpc.CallOption) (*JobInfo, error)
}

type buildServiceClient struct {
	cc *grpc.ClientConn
}

func NewBuildServiceClient(cc *grpc.ClientConn) BuildServiceClient {
	return &buildServiceClient{cc}
}

func (c *buildServiceClient) Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/BuildService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/BuildService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) GetJob(ctx context.Context, in *Request, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, "/BuildService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
type BuildServiceServer interface {
	Run(context.Context, *Request) (*Reply, error)
	Ping(context.Context, *Request) (*Reply, error)
	GetJob(context.Context, *Request) (*JobInfo, error)
}

// UnimplementedBuildServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBuildServiceServer struct {
}

func (*UnimplementedBuildServiceServer) Run(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedBuildServiceServer) Ping(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBuildServiceServer) GetJob(ctx context.Context, req *Request) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}

func RegisterBuildServiceServer(s *grpc.Server, srv BuildServiceServer) {
	s.RegisterService(&_BuildService_serviceDesc, srv)
}

func _BuildService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BuildService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).Run(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BuildService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BuildService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).GetJob(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BuildService",
	HandlerType: (*BuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _BuildService_Run_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BuildService_Ping_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _BuildService_GetJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1.proto",
}
