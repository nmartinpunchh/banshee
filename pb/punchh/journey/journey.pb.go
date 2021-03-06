// Code generated by protoc-gen-go. DO NOT EDIT.
// source: punchh/journey/journey.proto

package journeypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	workflow "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
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

// Status ..
type Status int32

const (
	Status_STATUS_INVALID  Status = 0
	Status_STATUS_DRAFT    Status = 1
	Status_STATUS_LIVE     Status = 2
	Status_STATUS_DISABLED Status = 3
)

var Status_name = map[int32]string{
	0: "STATUS_INVALID",
	1: "STATUS_DRAFT",
	2: "STATUS_LIVE",
	3: "STATUS_DISABLED",
}

var Status_value = map[string]int32{
	"STATUS_INVALID":  0,
	"STATUS_DRAFT":    1,
	"STATUS_LIVE":     2,
	"STATUS_DISABLED": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ba8e4128e438cf91, []int{0}
}

// Journey ...
type Journey struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status               Status               `protobuf:"varint,3,opt,name=status,proto3,enum=journey.Status" json:"status,omitempty"`
	SegmentId            string               `protobuf:"bytes,4,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	ControlGroupSize     int64                `protobuf:"varint,5,opt,name=control_group_size,json=controlGroupSize,proto3" json:"control_group_size,omitempty"`
	GuestEntryLimit      int64                `protobuf:"varint,6,opt,name=guest_entry_limit,json=guestEntryLimit,proto3" json:"guest_entry_limit,omitempty"`
	Workflow             *workflow.Workflow   `protobuf:"bytes,7,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Journey) Reset()         { *m = Journey{} }
func (m *Journey) String() string { return proto.CompactTextString(m) }
func (*Journey) ProtoMessage()    {}
func (*Journey) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8e4128e438cf91, []int{0}
}

func (m *Journey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Journey.Unmarshal(m, b)
}
func (m *Journey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Journey.Marshal(b, m, deterministic)
}
func (m *Journey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Journey.Merge(m, src)
}
func (m *Journey) XXX_Size() int {
	return xxx_messageInfo_Journey.Size(m)
}
func (m *Journey) XXX_DiscardUnknown() {
	xxx_messageInfo_Journey.DiscardUnknown(m)
}

var xxx_messageInfo_Journey proto.InternalMessageInfo

func (m *Journey) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Journey) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Journey) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_INVALID
}

func (m *Journey) GetSegmentId() string {
	if m != nil {
		return m.SegmentId
	}
	return ""
}

func (m *Journey) GetControlGroupSize() int64 {
	if m != nil {
		return m.ControlGroupSize
	}
	return 0
}

func (m *Journey) GetGuestEntryLimit() int64 {
	if m != nil {
		return m.GuestEntryLimit
	}
	return 0
}

func (m *Journey) GetWorkflow() *workflow.Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func init() {
	proto.RegisterEnum("journey.Status", Status_name, Status_value)
	proto.RegisterType((*Journey)(nil), "journey.Journey")
}

func init() { proto.RegisterFile("punchh/journey/journey.proto", fileDescriptor_ba8e4128e438cf91) }

var fileDescriptor_ba8e4128e438cf91 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xa5, 0x55, 0x28, 0x53, 0x53, 0x70, 0xbd, 0x90, 0xc6, 0x3f, 0xc4, 0x8b, 0xa4, 0x31,
	0x34, 0xa9, 0xf1, 0xe0, 0x91, 0xa6, 0xd5, 0x60, 0x88, 0x07, 0xc0, 0x9a, 0x78, 0x21, 0x6d, 0xd9,
	0x52, 0x14, 0x58, 0x02, 0x4b, 0x9a, 0xf6, 0x3b, 0xfa, 0x9d, 0x0c, 0xcb, 0xd2, 0xab, 0x27, 0x76,
	0x7e, 0xef, 0xcd, 0x0c, 0xf3, 0xe0, 0x2a, 0xaf, 0xb2, 0xf5, 0x76, 0x3b, 0xfe, 0x26, 0x55, 0x91,
	0xe1, 0x7d, 0xfb, 0x35, 0xf3, 0x82, 0x50, 0x82, 0x24, 0x5e, 0x0e, 0x6f, 0x23, 0x42, 0xa2, 0x04,
	0x8f, 0x19, 0x5e, 0x55, 0x9b, 0x31, 0x8d, 0x53, 0x5c, 0xd2, 0x65, 0x9a, 0x37, 0xce, 0xe1, 0x0d,
	0x9f, 0xb3, 0x23, 0xc5, 0xcf, 0x26, 0x21, 0xbb, 0xe3, 0xa3, 0xd1, 0xef, 0x7e, 0x3b, 0x20, 0xbd,
	0x35, 0xc3, 0xd0, 0x33, 0x40, 0x49, 0x97, 0x05, 0x0d, 0xea, 0x21, 0x9a, 0xa0, 0x0b, 0x46, 0x7f,
	0x32, 0x34, 0x9b, 0x0d, 0x66, 0xbb, 0xc1, 0xf4, 0xdb, 0x0d, 0xae, 0xcc, 0xdc, 0x75, 0x8d, 0x9e,
	0xa0, 0x87, 0xb3, 0xb0, 0x69, 0xec, 0xfc, 0xdb, 0x28, 0xe1, 0x2c, 0x64, 0x6d, 0xf7, 0x20, 0x96,
	0x74, 0x49, 0xab, 0x52, 0xeb, 0xea, 0x82, 0x31, 0x98, 0x28, 0x66, 0x7b, 0xa7, 0xc7, 0xb0, 0xcb,
	0x65, 0x74, 0x0d, 0x50, 0xe2, 0x28, 0xc5, 0x19, 0x0d, 0xe2, 0x50, 0x3b, 0xd5, 0x05, 0x43, 0x76,
	0x65, 0x4e, 0xec, 0x10, 0x3d, 0x00, 0x5a, 0x93, 0x8c, 0x16, 0x24, 0x09, 0xa2, 0x82, 0x54, 0x79,
	0x50, 0xc6, 0x07, 0xac, 0x9d, 0xe9, 0x82, 0xd1, 0x75, 0x55, 0xae, 0xbc, 0xd6, 0x82, 0x17, 0x1f,
	0x30, 0x1a, 0xc1, 0x45, 0x54, 0xe1, 0x92, 0x06, 0x38, 0xa3, 0xc5, 0x3e, 0x48, 0xe2, 0x34, 0xa6,
	0x9a, 0xc8, 0xcc, 0x0a, 0x13, 0xe6, 0x35, 0x77, 0x6a, 0x8c, 0x4c, 0xe8, 0xb5, 0x89, 0x69, 0x12,
	0x3b, 0x0c, 0x99, 0xc7, 0x08, 0x3f, 0xf9, 0xc3, 0x3d, 0x7a, 0x46, 0x3e, 0x88, 0xcd, 0xaf, 0x23,
	0x04, 0x03, 0xcf, 0xb7, 0xfc, 0x0f, 0x2f, 0xb0, 0xdf, 0x17, 0x96, 0x63, 0xcf, 0xd4, 0x13, 0xa4,
	0xc2, 0x39, 0x67, 0x33, 0xd7, 0x7a, 0xf1, 0x55, 0x01, 0x29, 0xd0, 0xe7, 0xc4, 0xb1, 0x17, 0x73,
	0xb5, 0x83, 0x2e, 0x41, 0x69, 0x2d, 0xb6, 0x67, 0x4d, 0x9d, 0xf9, 0x4c, 0xed, 0x4e, 0xfb, 0x5f,
	0x32, 0x0f, 0x26, 0x5f, 0xad, 0x44, 0x96, 0xe8, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0xca, 0x85, 0x06, 0x23, 0x02, 0x00, 0x00,
}
