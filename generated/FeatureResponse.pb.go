// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FeatureResponse.proto

package environmentSensors

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureResponse struct {
	HasTemperature         *bool    `protobuf:"varint,2,req,name=hasTemperature" json:"hasTemperature,omitempty"`
	HasHumanity            *bool    `protobuf:"varint,3,req,name=hasHumanity" json:"hasHumanity,omitempty"`
	HasAtmosphericPressure *bool    `protobuf:"varint,4,req,name=hasAtmosphericPressure" json:"hasAtmosphericPressure,omitempty"`
	HasPm2_5               *bool    `protobuf:"varint,5,req,name=hasPm2_5,json=hasPm25" json:"hasPm2_5,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *FeatureResponse) Reset()         { *m = FeatureResponse{} }
func (m *FeatureResponse) String() string { return proto.CompactTextString(m) }
func (*FeatureResponse) ProtoMessage()    {}
func (*FeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureResponse_2261acc9277b6bfc, []int{0}
}
func (m *FeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureResponse.Unmarshal(m, b)
}
func (m *FeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureResponse.Marshal(b, m, deterministic)
}
func (dst *FeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureResponse.Merge(dst, src)
}
func (m *FeatureResponse) XXX_Size() int {
	return xxx_messageInfo_FeatureResponse.Size(m)
}
func (m *FeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureResponse proto.InternalMessageInfo

func (m *FeatureResponse) GetHasTemperature() bool {
	if m != nil && m.HasTemperature != nil {
		return *m.HasTemperature
	}
	return false
}

func (m *FeatureResponse) GetHasHumanity() bool {
	if m != nil && m.HasHumanity != nil {
		return *m.HasHumanity
	}
	return false
}

func (m *FeatureResponse) GetHasAtmosphericPressure() bool {
	if m != nil && m.HasAtmosphericPressure != nil {
		return *m.HasAtmosphericPressure
	}
	return false
}

func (m *FeatureResponse) GetHasPm2_5() bool {
	if m != nil && m.HasPm2_5 != nil {
		return *m.HasPm2_5
	}
	return false
}

func init() {
	proto.RegisterType((*FeatureResponse)(nil), "environmentSensors.FeatureResponse")
}

func init() {
	proto.RegisterFile("FeatureResponse.proto", fileDescriptor_FeatureResponse_2261acc9277b6bfc)
}

var fileDescriptor_FeatureResponse_2261acc9277b6bfc = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcc, 0xb1, 0x0a, 0xc2, 0x40,
	0x0c, 0xc6, 0x71, 0xac, 0x8a, 0xe5, 0x04, 0x85, 0x03, 0xa5, 0x6e, 0xc5, 0x41, 0x9c, 0x1c, 0x84,
	0xba, 0xbb, 0x88, 0x63, 0xa9, 0xee, 0x12, 0x24, 0x70, 0x1d, 0x2e, 0x39, 0x92, 0xab, 0xe0, 0x4b,
	0xf9, 0x8c, 0xe2, 0xb9, 0x48, 0xc1, 0x31, 0xdf, 0xff, 0x47, 0xcc, 0xe2, 0x84, 0x10, 0x3b, 0xc1,
	0x06, 0x35, 0x30, 0x29, 0xee, 0x82, 0x70, 0x64, 0x6b, 0x91, 0x1e, 0xad, 0x30, 0x79, 0xa4, 0x78,
	0x41, 0x52, 0x16, 0x5d, 0xbf, 0x06, 0x66, 0xde, 0xd3, 0x76, 0x63, 0x66, 0x0e, 0xf4, 0x8a, 0x3e,
	0xa0, 0xa4, 0x52, 0x64, 0x65, 0xb6, 0xcd, 0x9b, 0xde, 0x6a, 0x4b, 0x33, 0x75, 0xa0, 0xe7, 0xce,
	0x03, 0xb5, 0xf1, 0x59, 0x0c, 0x13, 0xfa, 0x9d, 0xec, 0xc1, 0x2c, 0x1d, 0xe8, 0x31, 0x7a, 0xd6,
	0xe0, 0x50, 0xda, 0x7b, 0x2d, 0xa8, 0xfa, 0xf9, 0x38, 0x4a, 0xf8, 0x4f, 0xb5, 0x2b, 0x93, 0x3b,
	0xd0, 0xda, 0xef, 0x6f, 0x55, 0x31, 0x4e, 0x72, 0xf2, 0xbd, 0xab, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0x22, 0x51, 0x08, 0xdc, 0x00, 0x00, 0x00,
}
