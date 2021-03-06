// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DataRequest.proto

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

type DataRequest struct {
	Temperature          *bool    `protobuf:"varint,3,req,name=temperature" json:"temperature,omitempty"`
	Humanity             *bool    `protobuf:"varint,4,req,name=humanity" json:"humanity,omitempty"`
	AtmosphericPressure  *bool    `protobuf:"varint,5,req,name=atmosphericPressure" json:"atmosphericPressure,omitempty"`
	Pm2_5                *bool    `protobuf:"varint,6,req,name=pm2_5,json=pm25" json:"pm2_5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_DataRequest_0d473383d1f45601, []int{0}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (dst *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(dst, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetTemperature() bool {
	if m != nil && m.Temperature != nil {
		return *m.Temperature
	}
	return false
}

func (m *DataRequest) GetHumanity() bool {
	if m != nil && m.Humanity != nil {
		return *m.Humanity
	}
	return false
}

func (m *DataRequest) GetAtmosphericPressure() bool {
	if m != nil && m.AtmosphericPressure != nil {
		return *m.AtmosphericPressure
	}
	return false
}

func (m *DataRequest) GetPm2_5() bool {
	if m != nil && m.Pm2_5 != nil {
		return *m.Pm2_5
	}
	return false
}

func init() {
	proto.RegisterType((*DataRequest)(nil), "environmentSensors.DataRequest")
}

func init() { proto.RegisterFile("DataRequest.proto", fileDescriptor_DataRequest_0d473383d1f45601) }

var fileDescriptor_DataRequest_0d473383d1f45601 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x74, 0x49, 0x2c, 0x49,
	0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a,
	0xcd, 0x2b, 0xcb, 0x2c, 0xca, 0xcf, 0xcb, 0x4d, 0xcd, 0x2b, 0x09, 0x4e, 0xcd, 0x2b, 0xce, 0x2f,
	0x2a, 0x56, 0x9a, 0xc4, 0xc8, 0xc5, 0x8d, 0xa4, 0x52, 0x48, 0x81, 0x8b, 0xbb, 0x24, 0x35, 0xb7,
	0x20, 0xb5, 0x28, 0xb1, 0xa4, 0xb4, 0x28, 0x55, 0x82, 0x59, 0x81, 0x49, 0x83, 0x23, 0x08, 0x59,
	0x48, 0x48, 0x8a, 0x8b, 0x23, 0xa3, 0x34, 0x37, 0x31, 0x2f, 0xb3, 0xa4, 0x52, 0x82, 0x05, 0x2c,
	0x0d, 0xe7, 0x0b, 0x19, 0x70, 0x09, 0x27, 0x96, 0xe4, 0xe6, 0x17, 0x17, 0x64, 0xa4, 0x16, 0x65,
	0x26, 0x07, 0x14, 0xa5, 0x16, 0x17, 0x83, 0x4c, 0x61, 0x05, 0x2b, 0xc3, 0x26, 0x25, 0x24, 0xcc,
	0xc5, 0x5a, 0x90, 0x6b, 0x14, 0x6f, 0x2a, 0xc1, 0x06, 0x56, 0xc3, 0x52, 0x90, 0x6b, 0x64, 0x0a,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0x79, 0x20, 0x93, 0xbc, 0x00, 0x00, 0x00,
}
