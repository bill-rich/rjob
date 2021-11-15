// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: job_service.proto

package api

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

type StartJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Cpu     int32  `protobuf:"varint,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memor   int32  `protobuf:"varint,4,opt,name=memor,proto3" json:"memor,omitempty"`
	Blkio   int32  `protobuf:"varint,5,opt,name=blkio,proto3" json:"blkio,omitempty"`
}

func (x *StartJobInput) Reset() {
	*x = StartJobInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobInput) ProtoMessage() {}

func (x *StartJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobInput.ProtoReflect.Descriptor instead.
func (*StartJobInput) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartJobInput) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *StartJobInput) GetCpu() int32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *StartJobInput) GetMemor() int32 {
	if x != nil {
		return x.Memor
	}
	return 0
}

func (x *StartJobInput) GetBlkio() int32 {
	if x != nil {
		return x.Blkio
	}
	return 0
}

type StartJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobResponse.ProtoReflect.Descriptor instead.
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{1}
}

func (x *StartJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *StartJobResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type StopJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId   string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *StopJobInput) Reset() {
	*x = StopJobInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobInput) ProtoMessage() {}

func (x *StopJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobInput.ProtoReflect.Descriptor instead.
func (*StopJobInput) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{2}
}

func (x *StopJobInput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *StopJobInput) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type StopJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode int32 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *StopJobResponse) Reset() {
	*x = StopJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobResponse) ProtoMessage() {}

func (x *StopJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobResponse.ProtoReflect.Descriptor instead.
func (*StopJobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{3}
}

func (x *StopJobResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type StatusInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *StatusInput) Reset() {
	*x = StatusInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusInput) ProtoMessage() {}

func (x *StatusInput) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusInput.ProtoReflect.Descriptor instead.
func (*StatusInput) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{4}
}

func (x *StatusInput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ExitCode int32  `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StatusResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type MonitorJobInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *MonitorJobInput) Reset() {
	*x = MonitorJobInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorJobInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorJobInput) ProtoMessage() {}

func (x *MonitorJobInput) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorJobInput.ProtoReflect.Descriptor instead.
func (*MonitorJobInput) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{6}
}

func (x *MonitorJobInput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type MonitorJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk string `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *MonitorJobResponse) Reset() {
	*x = MonitorJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorJobResponse) ProtoMessage() {}

func (x *MonitorJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorJobResponse.ProtoReflect.Descriptor instead.
func (*MonitorJobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{7}
}

func (x *MonitorJobResponse) GetChunk() string {
	if x != nil {
		return x.Chunk
	}
	return ""
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info,json=jobInfo,proto3" json:"job_info,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListJobsResponse) GetJobInfo() []*JobInfo {
	if x != nil {
		return x.JobInfo
	}
	return nil
}

type JobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // Can be RUNNING, STOPPED, or KILLED
	ExitCode int32  `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{9}
}

func (x *JobInfo) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *JobInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobInfo) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{10}
}

var File_job_service_proto protoreflect.FileDescriptor

var file_job_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x70, 0x75,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6b, 0x69, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6c, 0x6b, 0x69, 0x6f, 0x22, 0x3f, 0x0a, 0x10,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a,
	0x0c, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x2e,
	0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x24,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x0f, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x37, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57, 0x0a, 0x07, 0x4a, 0x6f,
	0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xe5, 0x01, 0x0a,
	0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0e,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x0d, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4a, 0x6f,
	0x62, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x13, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x23, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2d, 0x72, 0x69, 0x63, 0x68, 0x2f, 0x72, 0x6a, 0x6f,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_service_proto_rawDescOnce sync.Once
	file_job_service_proto_rawDescData = file_job_service_proto_rawDesc
)

func file_job_service_proto_rawDescGZIP() []byte {
	file_job_service_proto_rawDescOnce.Do(func() {
		file_job_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_service_proto_rawDescData)
	})
	return file_job_service_proto_rawDescData
}

var file_job_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_job_service_proto_goTypes = []interface{}{
	(*StartJobInput)(nil),      // 0: StartJobInput
	(*StartJobResponse)(nil),   // 1: StartJobResponse
	(*StopJobInput)(nil),       // 2: StopJobInput
	(*StopJobResponse)(nil),    // 3: StopJobResponse
	(*StatusInput)(nil),        // 4: StatusInput
	(*StatusResponse)(nil),     // 5: StatusResponse
	(*MonitorJobInput)(nil),    // 6: MonitorJobInput
	(*MonitorJobResponse)(nil), // 7: MonitorJobResponse
	(*ListJobsResponse)(nil),   // 8: ListJobsResponse
	(*JobInfo)(nil),            // 9: JobInfo
	(*Empty)(nil),              // 10: Empty
}
var file_job_service_proto_depIdxs = []int32{
	9,  // 0: ListJobsResponse.job_info:type_name -> JobInfo
	0,  // 1: Jobs.Start:input_type -> StartJobInput
	2,  // 2: Jobs.Stop:input_type -> StopJobInput
	4,  // 3: Jobs.Status:input_type -> StatusInput
	6,  // 4: Jobs.Monitor:input_type -> MonitorJobInput
	10, // 5: Jobs.List:input_type -> Empty
	1,  // 6: Jobs.Start:output_type -> StartJobResponse
	3,  // 7: Jobs.Stop:output_type -> StopJobResponse
	5,  // 8: Jobs.Status:output_type -> StatusResponse
	7,  // 9: Jobs.Monitor:output_type -> MonitorJobResponse
	8,  // 10: Jobs.List:output_type -> ListJobsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_job_service_proto_init() }
func file_job_service_proto_init() {
	if File_job_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobInput); i {
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
		file_job_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobResponse); i {
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
		file_job_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopJobInput); i {
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
		file_job_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopJobResponse); i {
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
		file_job_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusInput); i {
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
		file_job_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_job_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorJobInput); i {
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
		file_job_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorJobResponse); i {
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
		file_job_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
		file_job_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobInfo); i {
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
		file_job_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_service_proto_goTypes,
		DependencyIndexes: file_job_service_proto_depIdxs,
		MessageInfos:      file_job_service_proto_msgTypes,
	}.Build()
	File_job_service_proto = out.File
	file_job_service_proto_rawDesc = nil
	file_job_service_proto_goTypes = nil
	file_job_service_proto_depIdxs = nil
}
