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

// Request
// Build Job Request
type Request struct {
	// project name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// event branch
	Branch string `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	// repository url
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// commit id
	Id int32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	// build language
	Language string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	// default is "latest"
	Lanversion string `protobuf:"bytes,6,opt,name=lanversion,proto3" json:"lanversion,omitempty"`
	// is query the latest build job. Only use in query actions.
	Latest bool `protobuf:"varint,7,opt,name=latest,proto3" json:"latest,omitempty"`
	// build owner
	User string `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	// checkout_sha
	Sha string `protobuf:"bytes,9,opt,name=sha,proto3" json:"sha,omitempty"`
	// the latest commit message
	Message string `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	// build env
	Env                  map[string]string `protobuf:"bytes,11,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
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

func (m *Request) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Request) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Request) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
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

// JobDetail
// Job Detail Info
type JobDetail struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 1 - Git clone success
	//-1 - Git clone failed
	// 2 - Ntci parse success
	//-2 - Ntci parse failed
	// 3 - Building
	// 4 - Build success
	//-4 - Build failed
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

// JobRequest
// Query Job Info via this message
type JobRequest struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{4}
}

func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *JobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{5}
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

type Job struct {
	// Job Owner
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{6}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Log struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{7}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterMapType((map[string]string)(nil), "Request.EnvEntry")
	proto.RegisterType((*Reply)(nil), "Reply")
	proto.RegisterType((*JobDetail)(nil), "JobDetail")
	proto.RegisterType((*JobInfo)(nil), "JobInfo")
	proto.RegisterType((*JobRequest)(nil), "JobRequest")
	proto.RegisterType((*Builder)(nil), "Builder")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*Log)(nil), "Log")
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor_2e4aa7d76fd7ee8a) }

var fileDescriptor_2e4aa7d76fd7ee8a = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xff, 0xe9, 0xd9, 0x34, 0x81, 0x35, 0x21, 0x2f, 0x20, 0x56, 0xc2, 0x4d, 0xaf,
	0x22, 0x36, 0xc4, 0x84, 0xe0, 0x6e, 0x62, 0x42, 0x54, 0xbb, 0x40, 0xde, 0x1d, 0x12, 0x17, 0x4e,
	0x63, 0xba, 0x94, 0xd4, 0x2e, 0xb6, 0x13, 0xd4, 0x67, 0xe3, 0x05, 0xf6, 0x58, 0xc8, 0x8e, 0xdb,
	0xa6, 0xa8, 0xdc, 0x9d, 0xef, 0x38, 0x3e, 0xfe, 0xce, 0xef, 0x1c, 0x05, 0xd2, 0xee, 0xb2, 0x58,
	0x49, 0xa1, 0x45, 0xfe, 0xe8, 0x43, 0x42, 0xd8, 0xaf, 0x96, 0x29, 0x8d, 0x10, 0x84, 0x9c, 0x2e,
	0x19, 0xf6, 0xc6, 0xde, 0x64, 0x44, 0x6c, 0x8c, 0x9e, 0x41, 0x5c, 0x4a, 0xca, 0x67, 0x0f, 0xd8,
	0xb7, 0x59, 0xa7, 0xd0, 0x13, 0x08, 0x5a, 0xd9, 0xe0, 0xc0, 0x26, 0x4d, 0x88, 0x4e, 0xc1, 0xaf,
	0x2b, 0x1c, 0x8e, 0xbd, 0x49, 0x44, 0xfc, 0xba, 0x42, 0x19, 0xa4, 0x0d, 0xe5, 0xf3, 0x96, 0xce,
	0x19, 0x8e, 0xec, 0x67, 0x5b, 0x8d, 0x5e, 0x02, 0x34, 0x94, 0x77, 0x4c, 0xaa, 0x5a, 0x70, 0x1c,
	0xdb, 0xd3, 0x41, 0xc6, 0xbc, 0xda, 0x50, 0xcd, 0x94, 0xc6, 0xc9, 0xd8, 0x9b, 0xa4, 0xc4, 0x29,
	0xe3, 0xb0, 0x55, 0x4c, 0xe2, 0xb4, 0x77, 0x68, 0x62, 0xe3, 0x44, 0x3d, 0x50, 0x3c, 0xea, 0x9d,
	0xa8, 0x07, 0x8a, 0x30, 0x24, 0x4b, 0xa6, 0x94, 0x79, 0x18, 0x6c, 0x76, 0x23, 0xd1, 0x6b, 0x08,
	0x18, 0xef, 0xf0, 0xf1, 0x38, 0x98, 0x1c, 0x5f, 0x3d, 0x2d, 0x5c, 0xe3, 0xc5, 0x2d, 0xef, 0x6e,
	0xb9, 0x96, 0x6b, 0x62, 0x4e, 0xb3, 0x6b, 0x48, 0x37, 0x09, 0x53, 0xfc, 0x27, 0x5b, 0x3b, 0x22,
	0x26, 0x44, 0x67, 0x10, 0x75, 0xb4, 0x69, 0x99, 0xe3, 0xd1, 0x8b, 0x0f, 0xfe, 0x7b, 0x2f, 0x7f,
	0x07, 0x11, 0x61, 0xab, 0x66, 0x6d, 0x5c, 0xce, 0x44, 0xd5, 0x73, 0x8c, 0x88, 0x8d, 0x87, 0x9e,
	0xfc, 0x3d, 0x4f, 0xf9, 0x1f, 0x0f, 0x46, 0x53, 0x51, 0x7e, 0x62, 0x9a, 0xd6, 0xcd, 0xff, 0x66,
	0xa0, 0x34, 0xd5, 0xad, 0xb2, 0x57, 0x23, 0xe2, 0x14, 0x7a, 0x01, 0x23, 0x5d, 0x2f, 0x99, 0xd2,
	0x74, 0xb9, 0x72, 0x93, 0xd8, 0x25, 0x06, 0x93, 0x0b, 0x0f, 0x4d, 0x2e, 0xfa, 0x77, 0x72, 0xf1,
	0x76, 0x72, 0x8e, 0x68, 0x72, 0x90, 0x68, 0xba, 0xef, 0xfe, 0x23, 0x24, 0x53, 0x51, 0x7e, 0xe1,
	0x3f, 0x84, 0x21, 0x33, 0x13, 0x2d, 0xd7, 0xae, 0xef, 0x5e, 0xa0, 0x0c, 0xfc, 0x45, 0x85, 0x7d,
	0x4b, 0x1c, 0x8a, 0x6d, 0xa3, 0xc4, 0x5f, 0x54, 0xf9, 0x35, 0xc0, 0x54, 0x94, 0x9b, 0xf5, 0x3b,
	0x83, 0x48, 0xfc, 0xe6, 0x4c, 0xba, 0xde, 0x7b, 0xb1, 0x05, 0xe2, 0xef, 0x80, 0xe4, 0xdf, 0x21,
	0xb9, 0x69, 0xeb, 0xa6, 0x62, 0xd2, 0x5c, 0x5a, 0x0c, 0x80, 0xf5, 0xc2, 0x74, 0xb0, 0xa8, 0x2b,
	0x77, 0xc7, 0x84, 0x03, 0x86, 0xc1, 0x1e, 0xc3, 0xcd, 0x46, 0x85, 0xbb, 0x8d, 0xca, 0xcf, 0x21,
	0x98, 0x8a, 0xf2, 0xd0, 0x28, 0xf2, 0x0b, 0x08, 0xee, 0xc4, 0x7c, 0xc8, 0xc3, 0xdb, 0xe3, 0x71,
	0xf5, 0xe8, 0xc1, 0x89, 0xf5, 0x76, 0xcf, 0x64, 0x57, 0xcf, 0x18, 0x3a, 0x87, 0x80, 0xb4, 0x1c,
	0xa5, 0x9b, 0x65, 0xcb, 0xe2, 0xc2, 0x6e, 0x49, 0x7e, 0x84, 0x32, 0x08, 0xbf, 0xd6, 0x7c, 0x7e,
	0xf0, 0xec, 0x15, 0xc4, 0x9f, 0x99, 0x36, 0x36, 0x8e, 0x8b, 0x1d, 0xa3, 0x2c, 0x2d, 0x1c, 0xed,
	0xfc, 0x08, 0x5d, 0xd8, 0xbd, 0xb9, 0xef, 0xfb, 0x48, 0x0b, 0x47, 0x64, 0x50, 0xe3, 0x39, 0x8c,
	0xfa, 0x1a, 0xc6, 0x72, 0x68, 0x6e, 0x66, 0x61, 0x71, 0x27, 0xe6, 0xf9, 0xd1, 0x1b, 0x0f, 0x8d,
	0x01, 0x88, 0xd9, 0x14, 0x69, 0x1f, 0x39, 0x60, 0xe1, 0xe6, 0xf4, 0xdb, 0x49, 0x69, 0x6a, 0x16,
	0x72, 0x35, 0x2b, 0xba, 0xcb, 0x32, 0xb6, 0x7f, 0x8c, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x75, 0xe0, 0x2b, 0xa7, 0x3d, 0x04, 0x00, 0x00,
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
	GetJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// Builder update status via this rpc.
	JobStatus(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error)
	GetJobLog(ctx context.Context, in *Job, opts ...grpc.CallOption) (BuildService_GetJobLogClient, error)
	RestartJob(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
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

func (c *buildServiceClient) GetJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, "/BuildService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) JobStatus(ctx context.Context, in *Builder, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/BuildService/JobStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) GetJobLog(ctx context.Context, in *Job, opts ...grpc.CallOption) (BuildService_GetJobLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BuildService_serviceDesc.Streams[0], "/BuildService/GetJobLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildServiceGetJobLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BuildService_GetJobLogClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type buildServiceGetJobLogClient struct {
	grpc.ClientStream
}

func (x *buildServiceGetJobLogClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildServiceClient) RestartJob(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/BuildService/RestartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
type BuildServiceServer interface {
	Run(context.Context, *Request) (*Reply, error)
	Ping(context.Context, *Request) (*Reply, error)
	GetJob(context.Context, *JobRequest) (*JobInfo, error)
	// Builder update status via this rpc.
	JobStatus(context.Context, *Builder) (*Reply, error)
	GetJobLog(*Job, BuildService_GetJobLogServer) error
	RestartJob(context.Context, *Request) (*Reply, error)
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
func (*UnimplementedBuildServiceServer) GetJob(ctx context.Context, req *JobRequest) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedBuildServiceServer) JobStatus(ctx context.Context, req *Builder) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobStatus not implemented")
}
func (*UnimplementedBuildServiceServer) GetJobLog(req *Job, srv BuildService_GetJobLogServer) error {
	return status.Errorf(codes.Unimplemented, "method GetJobLog not implemented")
}
func (*UnimplementedBuildServiceServer) RestartJob(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartJob not implemented")
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
	in := new(JobRequest)
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
		return srv.(BuildServiceServer).GetJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_JobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Builder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).JobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BuildService/JobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).JobStatus(ctx, req.(*Builder))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_GetJobLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Job)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServiceServer).GetJobLog(m, &buildServiceGetJobLogServer{stream})
}

type BuildService_GetJobLogServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type buildServiceGetJobLogServer struct {
	grpc.ServerStream
}

func (x *buildServiceGetJobLogServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _BuildService_RestartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).RestartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BuildService/RestartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).RestartJob(ctx, req.(*Request))
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
		{
			MethodName: "JobStatus",
			Handler:    _BuildService_JobStatus_Handler,
		},
		{
			MethodName: "RestartJob",
			Handler:    _BuildService_RestartJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetJobLog",
			Handler:       _BuildService_GetJobLog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1.proto",
}
