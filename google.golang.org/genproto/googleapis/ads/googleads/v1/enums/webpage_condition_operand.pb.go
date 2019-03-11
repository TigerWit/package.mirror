// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/webpage_condition_operand.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The webpage condition operand in webpage criterion.
type WebpageConditionOperandEnum_WebpageConditionOperand int32

const (
	// Not specified.
	WebpageConditionOperandEnum_UNSPECIFIED WebpageConditionOperandEnum_WebpageConditionOperand = 0
	// Used for return value only. Represents value unknown in this version.
	WebpageConditionOperandEnum_UNKNOWN WebpageConditionOperandEnum_WebpageConditionOperand = 1
	// Operand denoting a webpage URL targeting condition.
	WebpageConditionOperandEnum_URL WebpageConditionOperandEnum_WebpageConditionOperand = 2
	// Operand denoting a webpage category targeting condition.
	WebpageConditionOperandEnum_CATEGORY WebpageConditionOperandEnum_WebpageConditionOperand = 3
	// Operand denoting a webpage title targeting condition.
	WebpageConditionOperandEnum_PAGE_TITLE WebpageConditionOperandEnum_WebpageConditionOperand = 4
	// Operand denoting a webpage content targeting condition.
	WebpageConditionOperandEnum_PAGE_CONTENT WebpageConditionOperandEnum_WebpageConditionOperand = 5
	// Operand denoting a webpage custom label targeting condition.
	WebpageConditionOperandEnum_CUSTOM_LABEL WebpageConditionOperandEnum_WebpageConditionOperand = 6
)

var WebpageConditionOperandEnum_WebpageConditionOperand_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "URL",
	3: "CATEGORY",
	4: "PAGE_TITLE",
	5: "PAGE_CONTENT",
	6: "CUSTOM_LABEL",
}
var WebpageConditionOperandEnum_WebpageConditionOperand_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"URL":          2,
	"CATEGORY":     3,
	"PAGE_TITLE":   4,
	"PAGE_CONTENT": 5,
	"CUSTOM_LABEL": 6,
}

func (x WebpageConditionOperandEnum_WebpageConditionOperand) String() string {
	return proto.EnumName(WebpageConditionOperandEnum_WebpageConditionOperand_name, int32(x))
}
func (WebpageConditionOperandEnum_WebpageConditionOperand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_webpage_condition_operand_8ff719365c7e2862, []int{0, 0}
}

// Container for enum describing webpage condition operand in webpage criterion.
type WebpageConditionOperandEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebpageConditionOperandEnum) Reset()         { *m = WebpageConditionOperandEnum{} }
func (m *WebpageConditionOperandEnum) String() string { return proto.CompactTextString(m) }
func (*WebpageConditionOperandEnum) ProtoMessage()    {}
func (*WebpageConditionOperandEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_webpage_condition_operand_8ff719365c7e2862, []int{0}
}
func (m *WebpageConditionOperandEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebpageConditionOperandEnum.Unmarshal(m, b)
}
func (m *WebpageConditionOperandEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebpageConditionOperandEnum.Marshal(b, m, deterministic)
}
func (dst *WebpageConditionOperandEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebpageConditionOperandEnum.Merge(dst, src)
}
func (m *WebpageConditionOperandEnum) XXX_Size() int {
	return xxx_messageInfo_WebpageConditionOperandEnum.Size(m)
}
func (m *WebpageConditionOperandEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_WebpageConditionOperandEnum.DiscardUnknown(m)
}

var xxx_messageInfo_WebpageConditionOperandEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WebpageConditionOperandEnum)(nil), "google.ads.googleads.v1.enums.WebpageConditionOperandEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.WebpageConditionOperandEnum_WebpageConditionOperand", WebpageConditionOperandEnum_WebpageConditionOperand_name, WebpageConditionOperandEnum_WebpageConditionOperand_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/webpage_condition_operand.proto", fileDescriptor_webpage_condition_operand_8ff719365c7e2862)
}

var fileDescriptor_webpage_condition_operand_8ff719365c7e2862 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x7d, 0x2d, 0x4f, 0x30, 0x03, 0xd1, 0x49, 0x37, 0x26, 0x0a, 0x0b, 0xf8, 0x80, 0x69, 0x1a,
	0x77, 0x63, 0x5c, 0x4c, 0xeb, 0xd8, 0x10, 0x6b, 0xdb, 0x40, 0x0b, 0xd1, 0x34, 0x69, 0x0a, 0x6d,
	0x9a, 0x26, 0x30, 0xd3, 0x30, 0x80, 0x7b, 0x3f, 0x85, 0xa5, 0x9f, 0xe2, 0xa7, 0xb8, 0xf3, 0x0f,
	0x4c, 0x3b, 0x96, 0x1d, 0x6e, 0x26, 0x67, 0xee, 0xb9, 0xe7, 0xcc, 0xbd, 0x67, 0xc0, 0x7d, 0xce,
	0x79, 0xbe, 0xca, 0xf4, 0x24, 0x15, 0xba, 0x84, 0x15, 0xda, 0x1b, 0x7a, 0xc6, 0x76, 0x6b, 0xa1,
	0xbf, 0x65, 0x8b, 0x32, 0xc9, 0xb3, 0x78, 0xc9, 0x59, 0x5a, 0x6c, 0x0b, 0xce, 0x62, 0x5e, 0x66,
	0x9b, 0x84, 0xa5, 0xa8, 0xdc, 0xf0, 0x2d, 0xd7, 0x06, 0x52, 0x83, 0x92, 0x54, 0xa0, 0xa3, 0x1c,
	0xed, 0x0d, 0x54, 0xcb, 0xaf, 0xfb, 0x8d, 0x7b, 0x59, 0xe8, 0x09, 0x63, 0x7c, 0x9b, 0x54, 0x1e,
	0x42, 0x8a, 0x47, 0x07, 0x05, 0xdc, 0xcc, 0xe5, 0x03, 0x56, 0xe3, 0xef, 0x49, 0x7b, 0xca, 0x76,
	0xeb, 0xd1, 0xbb, 0x02, 0xae, 0x4e, 0xf0, 0xda, 0x25, 0xe8, 0x86, 0xee, 0xd4, 0xa7, 0xd6, 0xf8,
	0x71, 0x4c, 0x1f, 0xe0, 0x3f, 0xad, 0x0b, 0x3a, 0xa1, 0xfb, 0xe4, 0x7a, 0x73, 0x17, 0x2a, 0x5a,
	0x07, 0xb4, 0xc2, 0x89, 0x03, 0x55, 0xad, 0x07, 0xce, 0x2d, 0x12, 0x50, 0xdb, 0x9b, 0xbc, 0xc0,
	0x96, 0x76, 0x01, 0x80, 0x4f, 0x6c, 0x1a, 0x07, 0xe3, 0xc0, 0xa1, 0xf0, 0xbf, 0x06, 0x41, 0xaf,
	0xbe, 0x5b, 0x9e, 0x1b, 0x50, 0x37, 0x80, 0x67, 0x55, 0xc5, 0x0a, 0xa7, 0x81, 0xf7, 0x1c, 0x3b,
	0xc4, 0xa4, 0x0e, 0x6c, 0x9b, 0xdf, 0x0a, 0x18, 0x2e, 0xf9, 0x1a, 0xfd, 0xb9, 0xa8, 0xd9, 0x3f,
	0x31, 0xa7, 0x5f, 0x2d, 0xea, 0x2b, 0xaf, 0xe6, 0xaf, 0x3c, 0xe7, 0xab, 0x84, 0xe5, 0x88, 0x6f,
	0x72, 0x3d, 0xcf, 0x58, 0x1d, 0x43, 0x13, 0x7b, 0x59, 0x88, 0x13, 0xbf, 0x70, 0x57, 0x9f, 0x07,
	0xb5, 0x65, 0x13, 0xf2, 0xa1, 0x0e, 0x6c, 0x69, 0x45, 0x52, 0x81, 0x24, 0xac, 0xd0, 0xcc, 0x40,
	0x55, 0x66, 0xe2, 0xb3, 0xe1, 0x23, 0x92, 0x8a, 0xe8, 0xc8, 0x47, 0x33, 0x23, 0xaa, 0xf9, 0x2f,
	0x75, 0x28, 0x8b, 0x18, 0x93, 0x54, 0x60, 0x7c, 0xec, 0xc0, 0x78, 0x66, 0x60, 0x5c, 0xf7, 0x2c,
	0xda, 0xf5, 0x60, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x47, 0x65, 0x8f, 0x1d, 0x02,
	0x00, 0x00,
}
