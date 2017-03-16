// Code generated by protoc-gen-go.
// source: IngestConnection.proto
// DO NOT EDIT!

/*
Package Beam_Ftl_Ingest_Messages_Connection is a generated protocol buffer package.

It is generated from these files:
	IngestConnection.proto

It has these top-level messages:
	IngestMessage
	Connect
	Connect_Response
*/
package Beam_Ftl_Ingest_Messages_Connection

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolVersion int32

const (
	ProtocolVersion_V0 ProtocolVersion = 0
	ProtocolVersion_V1 ProtocolVersion = 1
)

var ProtocolVersion_name = map[int32]string{
	0: "V0",
	1: "V1",
}
var ProtocolVersion_value = map[string]int32{
	"V0": 0,
	"V1": 1,
}

func (x ProtocolVersion) String() string {
	return proto.EnumName(ProtocolVersion_name, int32(x))
}
func (ProtocolVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusCodes int32

const (
	StatusCodes_UNKNOWN                StatusCodes = 0
	StatusCodes_OK                     StatusCodes = 200
	StatusCodes_PING                   StatusCodes = 201
	StatusCodes_BAD_REQUEST            StatusCodes = 400
	StatusCodes_UNAUTHORIZED           StatusCodes = 401
	StatusCodes_OLD_VERSION            StatusCodes = 402
	StatusCodes_AUDIO_SSRC_COLLISION   StatusCodes = 403
	StatusCodes_VIDEO_SSRC_COLLISION   StatusCodes = 404
	StatusCodes_INVALID_STREAM_KEY     StatusCodes = 405
	StatusCodes_CHANNEL_IN_USE         StatusCodes = 406
	StatusCodes_REGION_UNSUPPORTED     StatusCodes = 407
	StatusCodes_NO_MEDIA_TIMEOUT       StatusCodes = 408
	StatusCodes_INTERNAL_SERVER_ERROR  StatusCodes = 500
	StatusCodes_INTERNAL_COMMAND_ERROR StatusCodes = 501
)

var StatusCodes_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "PING",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "OLD_VERSION",
	403: "AUDIO_SSRC_COLLISION",
	404: "VIDEO_SSRC_COLLISION",
	405: "INVALID_STREAM_KEY",
	406: "CHANNEL_IN_USE",
	407: "REGION_UNSUPPORTED",
	408: "NO_MEDIA_TIMEOUT",
	500: "INTERNAL_SERVER_ERROR",
	501: "INTERNAL_COMMAND_ERROR",
}
var StatusCodes_value = map[string]int32{
	"UNKNOWN":                0,
	"OK":                     200,
	"PING":                   201,
	"BAD_REQUEST":            400,
	"UNAUTHORIZED":           401,
	"OLD_VERSION":            402,
	"AUDIO_SSRC_COLLISION":   403,
	"VIDEO_SSRC_COLLISION":   404,
	"INVALID_STREAM_KEY":     405,
	"CHANNEL_IN_USE":         406,
	"REGION_UNSUPPORTED":     407,
	"NO_MEDIA_TIMEOUT":       408,
	"INTERNAL_SERVER_ERROR":  500,
	"INTERNAL_COMMAND_ERROR": 501,
}

func (x StatusCodes) String() string {
	return proto.EnumName(StatusCodes_name, int32(x))
}
func (StatusCodes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CommandType int32

const (
	CommandType_CONNECT          CommandType = 0
	CommandType_CONNECT_RESPONSE CommandType = 1
)

var CommandType_name = map[int32]string{
	0: "CONNECT",
	1: "CONNECT_RESPONSE",
}
var CommandType_value = map[string]int32{
	"CONNECT":          0,
	"CONNECT_RESPONSE": 1,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type IngestMessage struct {
	StatusCode  StatusCodes          `protobuf:"varint,1,opt,name=StatusCode,enum=Beam.Ftl.Ingest.Messages.Connection.StatusCodes" json:"StatusCode,omitempty"`
	CommandType CommandType          `protobuf:"varint,2,opt,name=CommandType,enum=Beam.Ftl.Ingest.Messages.Connection.CommandType" json:"CommandType,omitempty"`
	Command     *google_protobuf.Any `protobuf:"bytes,3,opt,name=Command" json:"Command,omitempty"`
}

func (m *IngestMessage) Reset()                    { *m = IngestMessage{} }
func (m *IngestMessage) String() string            { return proto.CompactTextString(m) }
func (*IngestMessage) ProtoMessage()               {}
func (*IngestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IngestMessage) GetStatusCode() StatusCodes {
	if m != nil {
		return m.StatusCode
	}
	return StatusCodes_UNKNOWN
}

func (m *IngestMessage) GetCommandType() CommandType {
	if m != nil {
		return m.CommandType
	}
	return CommandType_CONNECT
}

func (m *IngestMessage) GetCommand() *google_protobuf.Any {
	if m != nil {
		return m.Command
	}
	return nil
}

type Connect struct {
	ClientProtocolVersion ProtocolVersion `protobuf:"varint,1,opt,name=ClientProtocolVersion,enum=Beam.Ftl.Ingest.Messages.Connection.ProtocolVersion" json:"ClientProtocolVersion,omitempty"`
}

func (m *Connect) Reset()                    { *m = Connect{} }
func (m *Connect) String() string            { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()               {}
func (*Connect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Connect) GetClientProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.ClientProtocolVersion
	}
	return ProtocolVersion_V0
}

type Connect_Response struct {
	ServerProtocolVersion ProtocolVersion `protobuf:"varint,1,opt,name=ServerProtocolVersion,enum=Beam.Ftl.Ingest.Messages.Connection.ProtocolVersion" json:"ServerProtocolVersion,omitempty"`
	StatusCode            StatusCodes     `protobuf:"varint,2,opt,name=StatusCode,enum=Beam.Ftl.Ingest.Messages.Connection.StatusCodes" json:"StatusCode,omitempty"`
	HmacKey               string          `protobuf:"bytes,3,opt,name=HmacKey" json:"HmacKey,omitempty"`
}

func (m *Connect_Response) Reset()                    { *m = Connect_Response{} }
func (m *Connect_Response) String() string            { return proto.CompactTextString(m) }
func (*Connect_Response) ProtoMessage()               {}
func (*Connect_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Connect_Response) GetServerProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.ServerProtocolVersion
	}
	return ProtocolVersion_V0
}

func (m *Connect_Response) GetStatusCode() StatusCodes {
	if m != nil {
		return m.StatusCode
	}
	return StatusCodes_UNKNOWN
}

func (m *Connect_Response) GetHmacKey() string {
	if m != nil {
		return m.HmacKey
	}
	return ""
}

func init() {
	proto.RegisterType((*IngestMessage)(nil), "Beam.Ftl.Ingest.Messages.Connection.IngestMessage")
	proto.RegisterType((*Connect)(nil), "Beam.Ftl.Ingest.Messages.Connection.Connect")
	proto.RegisterType((*Connect_Response)(nil), "Beam.Ftl.Ingest.Messages.Connection.Connect_Response")
	proto.RegisterEnum("Beam.Ftl.Ingest.Messages.Connection.ProtocolVersion", ProtocolVersion_name, ProtocolVersion_value)
	proto.RegisterEnum("Beam.Ftl.Ingest.Messages.Connection.StatusCodes", StatusCodes_name, StatusCodes_value)
	proto.RegisterEnum("Beam.Ftl.Ingest.Messages.Connection.CommandType", CommandType_name, CommandType_value)
}

func init() { proto.RegisterFile("IngestConnection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0x06, 0xad, 0xda, 0x29, 0x0c, 0x63, 0xb6, 0xd1, 0x8d, 0x9b, 0x31, 0x6e, 0xa6,
	0x5d, 0x64, 0x63, 0xf0, 0x02, 0x59, 0x62, 0x36, 0xab, 0x89, 0x1d, 0xec, 0x24, 0x08, 0x6e, 0xac,
	0xac, 0x33, 0x55, 0x51, 0x9b, 0x54, 0x4d, 0x8a, 0xd4, 0xb7, 0xe0, 0x3f, 0x3c, 0x07, 0x4f, 0x01,
	0x0f, 0xc2, 0x0d, 0xd7, 0x70, 0x8f, 0x9a, 0xa6, 0xac, 0x94, 0x5d, 0x0c, 0x69, 0x57, 0xd6, 0xd1,
	0xf7, 0xfd, 0x3e, 0x9f, 0x63, 0xeb, 0xc0, 0x06, 0x4d, 0x3b, 0x3a, 0x2f, 0x9c, 0x2c, 0x4d, 0x75,
	0xbb, 0xe8, 0x66, 0xa9, 0x35, 0x18, 0x66, 0x45, 0x86, 0xef, 0x1f, 0xe9, 0xa4, 0x6f, 0x3d, 0x2e,
	0x7a, 0xd6, 0xd4, 0x60, 0xf9, 0x3a, 0xcf, 0x93, 0x8e, 0xce, 0xad, 0x73, 0xeb, 0xd6, 0x66, 0x27,
	0xcb, 0x3a, 0x3d, 0xbd, 0x5f, 0x22, 0xa7, 0xa3, 0x17, 0xfb, 0x49, 0x3a, 0x9e, 0xf2, 0x3b, 0xdf,
	0x0d, 0xb8, 0x31, 0x25, 0x2b, 0x10, 0x07, 0x00, 0xb2, 0x48, 0x8a, 0x51, 0xee, 0x64, 0x67, 0xba,
	0x69, 0x6c, 0x1b, 0xbb, 0xab, 0x87, 0x07, 0xd6, 0x25, 0xae, 0xb1, 0xce, 0xb1, 0x5c, 0xcc, 0x65,
	0x60, 0x01, 0x0d, 0x27, 0xeb, 0xf7, 0x93, 0xf4, 0x2c, 0x1c, 0x0f, 0x74, 0xb3, 0xf6, 0x1f, 0x91,
	0x73, 0x9c, 0x98, 0x0f, 0xc1, 0x16, 0xd4, 0xab, 0xb2, 0x69, 0x6e, 0x1b, 0xbb, 0x8d, 0xc3, 0x35,
	0x6b, 0x3a, 0xa4, 0x35, 0x1b, 0xd2, 0xb2, 0xd3, 0xb1, 0x98, 0x99, 0x76, 0x46, 0x13, 0x7f, 0x19,
	0x8b, 0x5f, 0xc2, 0xba, 0xd3, 0xeb, 0xea, 0xb4, 0x08, 0x26, 0xce, 0x76, 0xd6, 0x8b, 0xf5, 0x30,
	0xef, 0x66, 0x69, 0x35, 0xeb, 0xa3, 0x4b, 0x35, 0xb6, 0xc0, 0x8a, 0x8b, 0x23, 0x77, 0x7e, 0x18,
	0x80, 0x2a, 0x4a, 0x09, 0x9d, 0x0f, 0xb2, 0x34, 0xd7, 0x93, 0x06, 0xa4, 0x1e, 0xbe, 0xd2, 0xc3,
	0x2b, 0x6d, 0xe0, 0xc2, 0xc8, 0x85, 0xdf, 0xac, 0x5d, 0xc1, 0x6f, 0x36, 0xa1, 0x7e, 0xd2, 0x4f,
	0xda, 0x2d, 0x3d, 0x2e, 0x5f, 0x7e, 0x45, 0xcc, 0xca, 0xbd, 0x7b, 0x70, 0x73, 0xf1, 0xfa, 0x65,
	0xa8, 0xc5, 0x07, 0x68, 0xa9, 0x3c, 0x1f, 0x20, 0x63, 0xef, 0x4b, 0x0d, 0x1a, 0x73, 0xc1, 0xb8,
	0x01, 0xf5, 0x88, 0xb5, 0x18, 0x7f, 0xca, 0xd0, 0x12, 0xae, 0x43, 0x8d, 0xb7, 0xd0, 0x57, 0x03,
	0xaf, 0xc0, 0xb5, 0x80, 0xb2, 0x63, 0xf4, 0xcd, 0xc0, 0x08, 0x1a, 0x47, 0xb6, 0xab, 0x04, 0x79,
	0x12, 0x11, 0x19, 0xa2, 0xd7, 0x26, 0xbe, 0x05, 0xd7, 0x23, 0x66, 0x47, 0xe1, 0x09, 0x17, 0xf4,
	0x39, 0x71, 0xd1, 0x1b, 0x73, 0x62, 0xe2, 0x9e, 0xab, 0x62, 0x22, 0x24, 0xe5, 0x0c, 0xbd, 0x35,
	0xf1, 0x26, 0xac, 0xd9, 0x91, 0x4b, 0xb9, 0x92, 0x52, 0x38, 0xca, 0xe1, 0x9e, 0x47, 0x4b, 0xe9,
	0x5d, 0x29, 0xc5, 0xd4, 0x25, 0xff, 0x48, 0xef, 0x4d, 0x7c, 0x07, 0x30, 0x65, 0xb1, 0xed, 0x51,
	0x57, 0xc9, 0x50, 0x10, 0xdb, 0x57, 0x2d, 0xf2, 0x0c, 0x7d, 0x30, 0xf1, 0x6d, 0x58, 0x75, 0x4e,
	0x6c, 0xc6, 0x88, 0xa7, 0x28, 0x53, 0x91, 0x24, 0xe8, 0x63, 0xe9, 0x16, 0xe4, 0x98, 0x72, 0xa6,
	0x22, 0x26, 0xa3, 0x20, 0xe0, 0x22, 0x24, 0x2e, 0xfa, 0x64, 0xe2, 0x75, 0x40, 0x8c, 0x2b, 0x9f,
	0xb8, 0xd4, 0x56, 0x21, 0xf5, 0x09, 0x8f, 0x42, 0xf4, 0xd9, 0xc4, 0x5b, 0xb0, 0x4e, 0x59, 0x48,
	0x04, 0xb3, 0x3d, 0x25, 0x89, 0x88, 0x89, 0x50, 0x44, 0x08, 0x2e, 0xd0, 0x4f, 0x13, 0xdf, 0x85,
	0x8d, 0x3f, 0x9a, 0xc3, 0x7d, 0xdf, 0x66, 0x6e, 0x25, 0xfe, 0x32, 0xf7, 0x0e, 0xfe, 0xda, 0x9f,
	0xc9, 0x9b, 0x39, 0x9c, 0x31, 0xe2, 0x84, 0x68, 0x09, 0xaf, 0x01, 0xaa, 0x0a, 0x25, 0x88, 0x0c,
	0x38, 0x93, 0x04, 0x19, 0xa7, 0xcb, 0xe5, 0x12, 0x3c, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x65, 0xff, 0xfc, 0x36, 0x04, 0x00, 0x00,
}
