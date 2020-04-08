// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_common.proto

package refproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Team is either blue or yellow
type Team int32

const (
	// team not set
	Team_UNKNOWN Team = 0
	// yellow team
	Team_YELLOW Team = 1
	// blue team
	Team_BLUE Team = 2
)

var Team_name = map[int32]string{
	0: "UNKNOWN",
	1: "YELLOW",
	2: "BLUE",
}

var Team_value = map[string]int32{
	"UNKNOWN": 0,
	"YELLOW":  1,
	"BLUE":    2,
}

func (x Team) Enum() *Team {
	p := new(Team)
	*p = x
	return p
}

func (x Team) String() string {
	return proto.EnumName(Team_name, int32(x))
}

func (x *Team) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Team_value, data, "Team")
	if err != nil {
		return err
	}
	*x = Team(value)
	return nil
}

func (Team) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{0}
}

type ControllerReply_StatusCode int32

const (
	ControllerReply_UNKNOWN_STATUS_CODE ControllerReply_StatusCode = 0
	ControllerReply_OK                  ControllerReply_StatusCode = 1
	ControllerReply_REJECTED            ControllerReply_StatusCode = 2
)

var ControllerReply_StatusCode_name = map[int32]string{
	0: "UNKNOWN_STATUS_CODE",
	1: "OK",
	2: "REJECTED",
}

var ControllerReply_StatusCode_value = map[string]int32{
	"UNKNOWN_STATUS_CODE": 0,
	"OK":                  1,
	"REJECTED":            2,
}

func (x ControllerReply_StatusCode) Enum() *ControllerReply_StatusCode {
	p := new(ControllerReply_StatusCode)
	*p = x
	return p
}

func (x ControllerReply_StatusCode) String() string {
	return proto.EnumName(ControllerReply_StatusCode_name, int32(x))
}

func (x *ControllerReply_StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ControllerReply_StatusCode_value, data, "ControllerReply_StatusCode")
	if err != nil {
		return err
	}
	*x = ControllerReply_StatusCode(value)
	return nil
}

func (ControllerReply_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{2, 0}
}

type ControllerReply_Verification int32

const (
	ControllerReply_UNKNOWN_VERIFICATION ControllerReply_Verification = 0
	ControllerReply_VERIFIED             ControllerReply_Verification = 1
	ControllerReply_UNVERIFIED           ControllerReply_Verification = 2
)

var ControllerReply_Verification_name = map[int32]string{
	0: "UNKNOWN_VERIFICATION",
	1: "VERIFIED",
	2: "UNVERIFIED",
}

var ControllerReply_Verification_value = map[string]int32{
	"UNKNOWN_VERIFICATION": 0,
	"VERIFIED":             1,
	"UNVERIFIED":           2,
}

func (x ControllerReply_Verification) Enum() *ControllerReply_Verification {
	p := new(ControllerReply_Verification)
	*p = x
	return p
}

func (x ControllerReply_Verification) String() string {
	return proto.EnumName(ControllerReply_Verification_name, int32(x))
}

func (x *ControllerReply_Verification) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ControllerReply_Verification_value, data, "ControllerReply_Verification")
	if err != nil {
		return err
	}
	*x = ControllerReply_Verification(value)
	return nil
}

func (ControllerReply_Verification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{2, 1}
}

// BotId is the combination of a team and a robot id
type BotId struct {
	// the robot id - a negative value indicates that the id is not set
	Id *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// the team that the robot belongs to
	Team                 *Team    `protobuf:"varint,2,opt,name=team,enum=Team" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotId) Reset()         { *m = BotId{} }
func (m *BotId) String() string { return proto.CompactTextString(m) }
func (*BotId) ProtoMessage()    {}
func (*BotId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{0}
}

func (m *BotId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotId.Unmarshal(m, b)
}
func (m *BotId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotId.Marshal(b, m, deterministic)
}
func (m *BotId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotId.Merge(m, src)
}
func (m *BotId) XXX_Size() int {
	return xxx_messageInfo_BotId.Size(m)
}
func (m *BotId) XXX_DiscardUnknown() {
	xxx_messageInfo_BotId.DiscardUnknown(m)
}

var xxx_messageInfo_BotId proto.InternalMessageInfo

func (m *BotId) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *BotId) GetTeam() Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return Team_UNKNOWN
}

// Location is a 2d-coordinate on the field in ssl-vision coordinate system. Units are in meters.
type Location struct {
	// the x-coordinate in [m] in the ssl-vision coordinate system
	X *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	// the y-coordinate in [m] in the ssl-vision coordinate system
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Location) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

// a reply that is sent by the controller for each request from teams or autoRefs
type ControllerReply struct {
	// status_code is an optional code that indicates the result of the last request
	StatusCode *ControllerReply_StatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,enum=ControllerReply_StatusCode" json:"status_code,omitempty"`
	// reason is an optional explanation of the status code
	Reason *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	// next_token must be send with the next request, if secure communication is used
	// the token is used to avoid replay attacks
	// the token is always present in the very first message before the registration starts
	// the token is not present, if secure communication is not used
	NextToken *string `protobuf:"bytes,3,opt,name=next_token,json=nextToken" json:"next_token,omitempty"`
	// verification indicates if the last request could be verified (secure communication)
	Verification         *ControllerReply_Verification `protobuf:"varint,4,opt,name=verification,enum=ControllerReply_Verification" json:"verification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ControllerReply) Reset()         { *m = ControllerReply{} }
func (m *ControllerReply) String() string { return proto.CompactTextString(m) }
func (*ControllerReply) ProtoMessage()    {}
func (*ControllerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{2}
}

func (m *ControllerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerReply.Unmarshal(m, b)
}
func (m *ControllerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerReply.Marshal(b, m, deterministic)
}
func (m *ControllerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerReply.Merge(m, src)
}
func (m *ControllerReply) XXX_Size() int {
	return xxx_messageInfo_ControllerReply.Size(m)
}
func (m *ControllerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerReply.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerReply proto.InternalMessageInfo

func (m *ControllerReply) GetStatusCode() ControllerReply_StatusCode {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return ControllerReply_UNKNOWN_STATUS_CODE
}

func (m *ControllerReply) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func (m *ControllerReply) GetNextToken() string {
	if m != nil && m.NextToken != nil {
		return *m.NextToken
	}
	return ""
}

func (m *ControllerReply) GetVerification() ControllerReply_Verification {
	if m != nil && m.Verification != nil {
		return *m.Verification
	}
	return ControllerReply_UNKNOWN_VERIFICATION
}

// Signature can be added to a request to let it be verfied by the controller
type Signature struct {
	// the token that was received with the last controller reply
	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	// the PKCS1v15 signature of this message
	Pkcs1V15             []byte   `protobuf:"bytes,2,req,name=pkcs1v15" json:"pkcs1v15,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{3}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Signature) GetPkcs1V15() []byte {
	if m != nil {
		return m.Pkcs1V15
	}
	return nil
}

// BallSpeedMeasurement is a single measurement sample of the ball speed
type BallSpeedMeasurement struct {
	// The UNIX timestamp [μs] when the measurement was taken.
	Timestamp *uint64 `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	// the ball speed measurement [m/s]
	BallSpeed *float32 `protobuf:"fixed32,2,req,name=ball_speed,json=ballSpeed" json:"ball_speed,omitempty"`
	// the estimated initial ball speed (kick speed) [m/s]
	InitialBallSpeed     *float32 `protobuf:"fixed32,3,opt,name=initial_ball_speed,json=initialBallSpeed" json:"initial_ball_speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BallSpeedMeasurement) Reset()         { *m = BallSpeedMeasurement{} }
func (m *BallSpeedMeasurement) String() string { return proto.CompactTextString(m) }
func (*BallSpeedMeasurement) ProtoMessage()    {}
func (*BallSpeedMeasurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d09d87af007c4fe, []int{4}
}

func (m *BallSpeedMeasurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BallSpeedMeasurement.Unmarshal(m, b)
}
func (m *BallSpeedMeasurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BallSpeedMeasurement.Marshal(b, m, deterministic)
}
func (m *BallSpeedMeasurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BallSpeedMeasurement.Merge(m, src)
}
func (m *BallSpeedMeasurement) XXX_Size() int {
	return xxx_messageInfo_BallSpeedMeasurement.Size(m)
}
func (m *BallSpeedMeasurement) XXX_DiscardUnknown() {
	xxx_messageInfo_BallSpeedMeasurement.DiscardUnknown(m)
}

var xxx_messageInfo_BallSpeedMeasurement proto.InternalMessageInfo

func (m *BallSpeedMeasurement) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *BallSpeedMeasurement) GetBallSpeed() float32 {
	if m != nil && m.BallSpeed != nil {
		return *m.BallSpeed
	}
	return 0
}

func (m *BallSpeedMeasurement) GetInitialBallSpeed() float32 {
	if m != nil && m.InitialBallSpeed != nil {
		return *m.InitialBallSpeed
	}
	return 0
}

func init() {
	proto.RegisterEnum("Team", Team_name, Team_value)
	proto.RegisterEnum("ControllerReply_StatusCode", ControllerReply_StatusCode_name, ControllerReply_StatusCode_value)
	proto.RegisterEnum("ControllerReply_Verification", ControllerReply_Verification_name, ControllerReply_Verification_value)
	proto.RegisterType((*BotId)(nil), "BotId")
	proto.RegisterType((*Location)(nil), "Location")
	proto.RegisterType((*ControllerReply)(nil), "ControllerReply")
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*BallSpeedMeasurement)(nil), "BallSpeedMeasurement")
}

func init() { proto.RegisterFile("ssl_game_controller_common.proto", fileDescriptor_3d09d87af007c4fe) }

var fileDescriptor_3d09d87af007c4fe = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x2e, 0x49, 0x88, 0xa7, 0x51, 0xb0, 0x86, 0x08, 0xcc, 0x47, 0xa5, 0xc8, 0x07, 0x14,
	0x10, 0x8a, 0xd4, 0x48, 0x9c, 0x80, 0x43, 0x3e, 0x5c, 0x29, 0x34, 0x24, 0xd2, 0xc6, 0x69, 0xc5,
	0xc9, 0xda, 0xda, 0x4b, 0xb5, 0xaa, 0xed, 0xb5, 0xbc, 0x9b, 0x2a, 0x39, 0xf2, 0x07, 0xf9, 0x4d,
	0xc8, 0x76, 0xea, 0x06, 0xb8, 0xf9, 0xbd, 0x79, 0x6f, 0xde, 0x78, 0x66, 0xa1, 0xaf, 0x75, 0x1c,
	0xdc, 0xf2, 0x44, 0x04, 0xa1, 0x4a, 0x4d, 0xae, 0xe2, 0x58, 0xe4, 0x41, 0xa8, 0x92, 0x44, 0xa5,
	0xc3, 0x2c, 0x57, 0x46, 0xb9, 0x23, 0x68, 0x4e, 0x94, 0x99, 0x47, 0xd8, 0x05, 0x2a, 0x23, 0x87,
	0xf4, 0xc9, 0xa0, 0xc9, 0xa8, 0x8c, 0xf0, 0x15, 0x34, 0x8c, 0xe0, 0x89, 0x43, 0xfb, 0x64, 0xd0,
	0x1d, 0x35, 0x87, 0xbe, 0xe0, 0x09, 0x2b, 0x29, 0xf7, 0x1d, 0xb4, 0x17, 0x2a, 0xe4, 0x46, 0xaa,
	0x14, 0x3b, 0x40, 0x76, 0x0e, 0xe9, 0xd3, 0x01, 0x65, 0x64, 0x57, 0xa0, 0xbd, 0x43, 0x2b, 0xb4,
	0x77, 0x7f, 0x53, 0x78, 0x36, 0xad, 0x73, 0x99, 0xc8, 0xe2, 0x3d, 0x7e, 0x81, 0x53, 0x6d, 0xb8,
	0xd9, 0xea, 0x20, 0x54, 0x91, 0x28, 0xf3, 0xba, 0xa3, 0x37, 0xc3, 0x7f, 0x64, 0xc3, 0x75, 0xa9,
	0x99, 0xaa, 0x48, 0x30, 0xd0, 0xf5, 0x37, 0xbe, 0x80, 0x56, 0x2e, 0xb8, 0x56, 0x69, 0x39, 0x96,
	0xc5, 0x0e, 0x08, 0xcf, 0x00, 0x52, 0xb1, 0x33, 0x81, 0x51, 0x77, 0x22, 0x75, 0x9e, 0x94, 0x35,
	0xab, 0x60, 0xfc, 0x82, 0xc0, 0x31, 0x74, 0xee, 0x45, 0x2e, 0x7f, 0xca, 0x6a, 0x68, 0xa7, 0x51,
	0xa6, 0x9e, 0xfd, 0x97, 0x7a, 0x75, 0x24, 0x62, 0x7f, 0x59, 0xdc, 0xcf, 0x00, 0x8f, 0x33, 0xe1,
	0x4b, 0x78, 0xbe, 0x59, 0x5e, 0x2e, 0x57, 0xd7, 0xcb, 0x60, 0xed, 0x8f, 0xfd, 0xcd, 0x3a, 0x98,
	0xae, 0x66, 0x9e, 0x7d, 0x82, 0x2d, 0xa0, 0xab, 0x4b, 0x9b, 0x60, 0x07, 0xda, 0xcc, 0xfb, 0xe6,
	0x4d, 0x7d, 0x6f, 0x66, 0x53, 0xf7, 0x02, 0x3a, 0xc7, 0xad, 0xd1, 0x81, 0xde, 0x83, 0xfd, 0xca,
	0x63, 0xf3, 0x8b, 0xf9, 0x74, 0xec, 0xcf, 0x57, 0x4b, 0xfb, 0xa4, 0xf0, 0x55, 0x8c, 0x37, 0xb3,
	0x09, 0x76, 0x01, 0x36, 0xcb, 0x1a, 0x53, 0xf7, 0x2b, 0x58, 0x6b, 0x79, 0x9b, 0x72, 0xb3, 0xcd,
	0x05, 0xf6, 0xa0, 0x59, 0xfd, 0x6e, 0xb1, 0x7d, 0x8b, 0x55, 0x00, 0x5f, 0x43, 0x3b, 0xbb, 0x0b,
	0xf5, 0xf9, 0xfd, 0xf9, 0xa7, 0xf2, 0x10, 0x1d, 0x56, 0x63, 0xf7, 0x17, 0x81, 0xde, 0x84, 0xc7,
	0xf1, 0x3a, 0x13, 0x22, 0xfa, 0x2e, 0xb8, 0xde, 0xe6, 0x22, 0x11, 0xa9, 0xc1, 0xb7, 0x60, 0x19,
	0x99, 0x08, 0x6d, 0x78, 0x92, 0x95, 0xed, 0x1a, 0xec, 0x91, 0x28, 0x96, 0x7b, 0xc3, 0xe3, 0x38,
	0xd0, 0x85, 0xed, 0x70, 0x5d, 0xeb, 0xe6, 0xa1, 0x0f, 0x7e, 0x04, 0x94, 0xa9, 0x34, 0x92, 0xc7,
	0xc1, 0x91, 0xac, 0xb8, 0x01, 0x65, 0xf6, 0xa1, 0x52, 0xa7, 0x7e, 0x78, 0x0f, 0x8d, 0xe2, 0x25,
	0xe1, 0x29, 0x3c, 0x3d, 0xac, 0xc0, 0x3e, 0x41, 0x80, 0xd6, 0x0f, 0x6f, 0xb1, 0x58, 0x5d, 0xdb,
	0x04, 0xdb, 0xd0, 0x98, 0x2c, 0x36, 0x9e, 0x4d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xdb,
	0x45, 0xd3, 0xbd, 0x02, 0x00, 0x00,
}