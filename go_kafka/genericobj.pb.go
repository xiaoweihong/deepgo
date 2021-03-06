// Code generated by protoc-gen-go. DO NOT EDIT.
// source: genericobj.proto

package go_kafka

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

type NonMotorVehicleGesture int32

const (
	NonMotorVehicleGesture_ATTITUDE_POSITIVE NonMotorVehicleGesture = 0
	NonMotorVehicleGesture_ATTITUDE_RIGHT    NonMotorVehicleGesture = 1
	NonMotorVehicleGesture_ATTITUDE_LEFT     NonMotorVehicleGesture = 2
	NonMotorVehicleGesture_ATTITUDE_BACK     NonMotorVehicleGesture = 3
)

var NonMotorVehicleGesture_name = map[int32]string{
	0: "ATTITUDE_POSITIVE",
	1: "ATTITUDE_RIGHT",
	2: "ATTITUDE_LEFT",
	3: "ATTITUDE_BACK",
}
var NonMotorVehicleGesture_value = map[string]int32{
	"ATTITUDE_POSITIVE": 0,
	"ATTITUDE_RIGHT":    1,
	"ATTITUDE_LEFT":     2,
	"ATTITUDE_BACK":     3,
}

func (x NonMotorVehicleGesture) String() string {
	return proto.EnumName(NonMotorVehicleGesture_name, int32(x))
}
func (NonMotorVehicleGesture) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{0}
}

type NonMotorVehicleType int32

const (
	NonMotorVehicleType_NONMOTOR_TYPE_UNKNOWN    NonMotorVehicleType = 0
	NonMotorVehicleType_NONMOTOR_TYPE_VEHICLE    NonMotorVehicleType = 1
	NonMotorVehicleType_NONMOTOR_TYPE_PEDESTRIAN NonMotorVehicleType = 2
	NonMotorVehicleType_NONMOTOR_TYPE_VEHICLE2   NonMotorVehicleType = 3
	NonMotorVehicleType_NONMOTOR_TYPE_VEHICLE3   NonMotorVehicleType = 4
	NonMotorVehicleType_NONMOTOR_TYPE_ROOF       NonMotorVehicleType = 5
	NonMotorVehicleType_NONMOTOR_TYPE_SEAL       NonMotorVehicleType = 6
	NonMotorVehicleType_NONMOTOR_TYPE_BICYCLE    NonMotorVehicleType = 7
)

var NonMotorVehicleType_name = map[int32]string{
	0: "NONMOTOR_TYPE_UNKNOWN",
	1: "NONMOTOR_TYPE_VEHICLE",
	2: "NONMOTOR_TYPE_PEDESTRIAN",
	3: "NONMOTOR_TYPE_VEHICLE2",
	4: "NONMOTOR_TYPE_VEHICLE3",
	5: "NONMOTOR_TYPE_ROOF",
	6: "NONMOTOR_TYPE_SEAL",
	7: "NONMOTOR_TYPE_BICYCLE",
}
var NonMotorVehicleType_value = map[string]int32{
	"NONMOTOR_TYPE_UNKNOWN":    0,
	"NONMOTOR_TYPE_VEHICLE":    1,
	"NONMOTOR_TYPE_PEDESTRIAN": 2,
	"NONMOTOR_TYPE_VEHICLE2":   3,
	"NONMOTOR_TYPE_VEHICLE3":   4,
	"NONMOTOR_TYPE_ROOF":       5,
	"NONMOTOR_TYPE_SEAL":       6,
	"NONMOTOR_TYPE_BICYCLE":    7,
}

func (x NonMotorVehicleType) String() string {
	return proto.EnumName(NonMotorVehicleType_name, int32(x))
}
func (NonMotorVehicleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{1}
}

type GenericObj struct {
	ObjType              ObjType     `protobuf:"varint,1,opt,name=ObjType,proto3,enum=ObjType" json:"ObjType,omitempty"`
	FmtType              DataFmtType `protobuf:"varint,2,opt,name=FmtType,proto3,enum=DataFmtType" json:"FmtType,omitempty"`
	BinData              []byte      `protobuf:"bytes,3,opt,name=BinData,proto3" json:"BinData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenericObj) Reset()         { *m = GenericObj{} }
func (m *GenericObj) String() string { return proto.CompactTextString(m) }
func (*GenericObj) ProtoMessage()    {}
func (*GenericObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{0}
}
func (m *GenericObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericObj.Unmarshal(m, b)
}
func (m *GenericObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericObj.Marshal(b, m, deterministic)
}
func (dst *GenericObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericObj.Merge(dst, src)
}
func (m *GenericObj) XXX_Size() int {
	return xxx_messageInfo_GenericObj.Size(m)
}
func (m *GenericObj) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericObj.DiscardUnknown(m)
}

var xxx_messageInfo_GenericObj proto.InternalMessageInfo

func (m *GenericObj) GetObjType() ObjType {
	if m != nil {
		return m.ObjType
	}
	return ObjType_OBJ_TYPE_UNKNOWN
}

func (m *GenericObj) GetFmtType() DataFmtType {
	if m != nil {
		return m.FmtType
	}
	return DataFmtType_UNKNOWNFMT
}

func (m *GenericObj) GetBinData() []byte {
	if m != nil {
		return m.BinData
	}
	return nil
}

type VehicleObj struct {
	Metadata             *SrcMetadata  `protobuf:"bytes,1,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	Img                  *Image        `protobuf:"bytes,2,opt,name=Img,proto3" json:"Img,omitempty"`
	Vehicle              []*RecVehicle `protobuf:"bytes,3,rep,name=Vehicle,proto3" json:"Vehicle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VehicleObj) Reset()         { *m = VehicleObj{} }
func (m *VehicleObj) String() string { return proto.CompactTextString(m) }
func (*VehicleObj) ProtoMessage()    {}
func (*VehicleObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{1}
}
func (m *VehicleObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleObj.Unmarshal(m, b)
}
func (m *VehicleObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleObj.Marshal(b, m, deterministic)
}
func (dst *VehicleObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleObj.Merge(dst, src)
}
func (m *VehicleObj) XXX_Size() int {
	return xxx_messageInfo_VehicleObj.Size(m)
}
func (m *VehicleObj) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleObj.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleObj proto.InternalMessageInfo

func (m *VehicleObj) GetMetadata() *SrcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VehicleObj) GetImg() *Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *VehicleObj) GetVehicle() []*RecVehicle {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

type PedestrianObj struct {
	Metadata             *SrcMetadata     `protobuf:"bytes,1,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	Img                  *Image           `protobuf:"bytes,2,opt,name=Img,proto3" json:"Img,omitempty"`
	Pedestrian           []*RecPedestrian `protobuf:"bytes,3,rep,name=Pedestrian,proto3" json:"Pedestrian,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PedestrianObj) Reset()         { *m = PedestrianObj{} }
func (m *PedestrianObj) String() string { return proto.CompactTextString(m) }
func (*PedestrianObj) ProtoMessage()    {}
func (*PedestrianObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{2}
}
func (m *PedestrianObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PedestrianObj.Unmarshal(m, b)
}
func (m *PedestrianObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PedestrianObj.Marshal(b, m, deterministic)
}
func (dst *PedestrianObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PedestrianObj.Merge(dst, src)
}
func (m *PedestrianObj) XXX_Size() int {
	return xxx_messageInfo_PedestrianObj.Size(m)
}
func (m *PedestrianObj) XXX_DiscardUnknown() {
	xxx_messageInfo_PedestrianObj.DiscardUnknown(m)
}

var xxx_messageInfo_PedestrianObj proto.InternalMessageInfo

func (m *PedestrianObj) GetMetadata() *SrcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PedestrianObj) GetImg() *Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *PedestrianObj) GetPedestrian() []*RecPedestrian {
	if m != nil {
		return m.Pedestrian
	}
	return nil
}

type NonMotorVehicleObj struct {
	Metadata             *SrcMetadata          `protobuf:"bytes,1,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	Img                  *Image                `protobuf:"bytes,2,opt,name=Img,proto3" json:"Img,omitempty"`
	NonMotorVehicles     []*RecNonMotorVehicle `protobuf:"bytes,3,rep,name=NonMotorVehicles,proto3" json:"NonMotorVehicles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NonMotorVehicleObj) Reset()         { *m = NonMotorVehicleObj{} }
func (m *NonMotorVehicleObj) String() string { return proto.CompactTextString(m) }
func (*NonMotorVehicleObj) ProtoMessage()    {}
func (*NonMotorVehicleObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{3}
}
func (m *NonMotorVehicleObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMotorVehicleObj.Unmarshal(m, b)
}
func (m *NonMotorVehicleObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMotorVehicleObj.Marshal(b, m, deterministic)
}
func (dst *NonMotorVehicleObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMotorVehicleObj.Merge(dst, src)
}
func (m *NonMotorVehicleObj) XXX_Size() int {
	return xxx_messageInfo_NonMotorVehicleObj.Size(m)
}
func (m *NonMotorVehicleObj) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMotorVehicleObj.DiscardUnknown(m)
}

var xxx_messageInfo_NonMotorVehicleObj proto.InternalMessageInfo

func (m *NonMotorVehicleObj) GetMetadata() *SrcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NonMotorVehicleObj) GetImg() *Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *NonMotorVehicleObj) GetNonMotorVehicles() []*RecNonMotorVehicle {
	if m != nil {
		return m.NonMotorVehicles
	}
	return nil
}

type FaceObj struct {
	Metadata             *SrcMetadata `protobuf:"bytes,1,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	Img                  *Image       `protobuf:"bytes,2,opt,name=Img,proto3" json:"Img,omitempty"`
	Faces                []*RecFace   `protobuf:"bytes,3,rep,name=Faces,proto3" json:"Faces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FaceObj) Reset()         { *m = FaceObj{} }
func (m *FaceObj) String() string { return proto.CompactTextString(m) }
func (*FaceObj) ProtoMessage()    {}
func (*FaceObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{4}
}
func (m *FaceObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceObj.Unmarshal(m, b)
}
func (m *FaceObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceObj.Marshal(b, m, deterministic)
}
func (dst *FaceObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceObj.Merge(dst, src)
}
func (m *FaceObj) XXX_Size() int {
	return xxx_messageInfo_FaceObj.Size(m)
}
func (m *FaceObj) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceObj.DiscardUnknown(m)
}

var xxx_messageInfo_FaceObj proto.InternalMessageInfo

func (m *FaceObj) GetMetadata() *SrcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FaceObj) GetImg() *Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *FaceObj) GetFaces() []*RecFace {
	if m != nil {
		return m.Faces
	}
	return nil
}

type RecObjs struct {
	Metadata             *SrcMetadata          `protobuf:"bytes,1,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	Img                  *Image                `protobuf:"bytes,2,opt,name=Img,proto3" json:"Img,omitempty"`
	Vehicle              []*RecVehicle         `protobuf:"bytes,3,rep,name=Vehicle,proto3" json:"Vehicle,omitempty"`
	Pedestrian           []*RecPedestrian      `protobuf:"bytes,4,rep,name=Pedestrian,proto3" json:"Pedestrian,omitempty"`
	NonMotorVehicles     []*RecNonMotorVehicle `protobuf:"bytes,5,rep,name=NonMotorVehicles,proto3" json:"NonMotorVehicles,omitempty"`
	RecFaces             []*RecFace            `protobuf:"bytes,6,rep,name=RecFaces,proto3" json:"RecFaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RecObjs) Reset()         { *m = RecObjs{} }
func (m *RecObjs) String() string { return proto.CompactTextString(m) }
func (*RecObjs) ProtoMessage()    {}
func (*RecObjs) Descriptor() ([]byte, []int) {
	return fileDescriptor_genericobj_d4cf9678f6c07910, []int{5}
}
func (m *RecObjs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecObjs.Unmarshal(m, b)
}
func (m *RecObjs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecObjs.Marshal(b, m, deterministic)
}
func (dst *RecObjs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecObjs.Merge(dst, src)
}
func (m *RecObjs) XXX_Size() int {
	return xxx_messageInfo_RecObjs.Size(m)
}
func (m *RecObjs) XXX_DiscardUnknown() {
	xxx_messageInfo_RecObjs.DiscardUnknown(m)
}

var xxx_messageInfo_RecObjs proto.InternalMessageInfo

func (m *RecObjs) GetMetadata() *SrcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RecObjs) GetImg() *Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *RecObjs) GetVehicle() []*RecVehicle {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

func (m *RecObjs) GetPedestrian() []*RecPedestrian {
	if m != nil {
		return m.Pedestrian
	}
	return nil
}

func (m *RecObjs) GetNonMotorVehicles() []*RecNonMotorVehicle {
	if m != nil {
		return m.NonMotorVehicles
	}
	return nil
}

func (m *RecObjs) GetRecFaces() []*RecFace {
	if m != nil {
		return m.RecFaces
	}
	return nil
}

func init() {
	proto.RegisterType((*GenericObj)(nil), "GenericObj")
	proto.RegisterType((*VehicleObj)(nil), "VehicleObj")
	proto.RegisterType((*PedestrianObj)(nil), "PedestrianObj")
	proto.RegisterType((*NonMotorVehicleObj)(nil), "NonMotorVehicleObj")
	proto.RegisterType((*FaceObj)(nil), "FaceObj")
	proto.RegisterType((*RecObjs)(nil), "RecObjs")
	proto.RegisterEnum("NonMotorVehicleGesture", NonMotorVehicleGesture_name, NonMotorVehicleGesture_value)
	proto.RegisterEnum("NonMotorVehicleType", NonMotorVehicleType_name, NonMotorVehicleType_value)
}

func init() { proto.RegisterFile("genericobj.proto", fileDescriptor_genericobj_d4cf9678f6c07910) }

var fileDescriptor_genericobj_d4cf9678f6c07910 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x8d, 0x71, 0x82, 0xd1, 0x40, 0xd0, 0x66, 0xa3, 0x20, 0x37, 0xaa, 0x5a, 0x6a, 0xb5, 0x15,
	0xca, 0x61, 0x0f, 0xe4, 0x07, 0x54, 0x7c, 0x18, 0x62, 0x05, 0x6c, 0xb4, 0x6c, 0xa8, 0x72, 0x42,
	0xc6, 0xac, 0xa8, 0x11, 0xf6, 0x22, 0xe3, 0xaa, 0xea, 0xb9, 0xc7, 0xfe, 0x82, 0xfe, 0xc7, 0xfe,
	0x88, 0x6a, 0x8d, 0x4d, 0x8a, 0x49, 0x55, 0x55, 0x42, 0xbd, 0x79, 0xde, 0x7b, 0x3b, 0xef, 0x79,
	0x76, 0x6c, 0x40, 0x0b, 0x1e, 0xf2, 0xc8, 0xf7, 0xc4, 0x6c, 0x49, 0xd6, 0x91, 0x88, 0xc5, 0x75,
	0xc5, 0x13, 0x41, 0x20, 0xc2, 0x6d, 0x65, 0x44, 0x00, 0xfd, 0xad, 0xc2, 0x99, 0x2d, 0xb1, 0x01,
	0x9a, 0x33, 0x5b, 0xb2, 0xaf, 0x6b, 0xae, 0x2b, 0x75, 0xa5, 0x51, 0x6d, 0x96, 0x48, 0x5a, 0xd3,
	0x8c, 0xc0, 0xef, 0x41, 0xeb, 0x05, 0x71, 0xa2, 0x29, 0x24, 0x9a, 0x0a, 0xe9, 0xba, 0xb1, 0x9b,
	0x62, 0x34, 0x23, 0xb1, 0x0e, 0x5a, 0xdb, 0x0f, 0x25, 0xa5, 0xab, 0x75, 0xa5, 0x51, 0xa1, 0x59,
	0x69, 0x7c, 0x01, 0x98, 0xf0, 0x4f, 0xbe, 0xb7, 0xe2, 0xd2, 0xb3, 0x01, 0xa5, 0x21, 0x8f, 0xdd,
	0xb9, 0x14, 0x4a, 0xd3, 0x72, 0xb3, 0x42, 0xc6, 0x91, 0x97, 0x61, 0x74, 0xc7, 0x62, 0x1d, 0x54,
	0x2b, 0x58, 0x24, 0xae, 0xe5, 0x66, 0x91, 0x58, 0x81, 0xbb, 0xe0, 0x54, 0x42, 0xf8, 0x1d, 0x68,
	0x69, 0x47, 0x5d, 0xad, 0xab, 0x8d, 0x72, 0xb3, 0x4c, 0x28, 0xf7, 0x52, 0x88, 0x66, 0x9c, 0xf1,
	0x4d, 0x81, 0xf3, 0x11, 0x9f, 0xf3, 0x4d, 0x1c, 0xf9, 0x6e, 0x78, 0x2c, 0x73, 0x02, 0xf0, 0xd4,
	0x34, 0xf5, 0xaf, 0x4a, 0xff, 0x27, 0x94, 0xfe, 0xa6, 0x30, 0x7e, 0x28, 0x80, 0x6d, 0x11, 0x0e,
	0x45, 0x2c, 0xa2, 0x23, 0xcf, 0xe1, 0x03, 0xa0, 0x5c, 0xe7, 0x4d, 0x1a, 0xe8, 0x52, 0x06, 0xca,
	0x71, 0xf4, 0x40, 0x6c, 0x04, 0xa0, 0xf5, 0x5c, 0xef, 0x68, 0x79, 0x5e, 0xc1, 0x99, 0x6c, 0x97,
	0x85, 0x28, 0xc9, 0x10, 0x12, 0xa0, 0x5b, 0xd8, 0xf8, 0x5e, 0x00, 0x8d, 0x72, 0xb9, 0x7a, 0x9b,
	0xff, 0xb8, 0x07, 0xb9, 0x1b, 0x3b, 0xfd, 0xdb, 0x8d, 0x3d, 0x3b, 0xd6, 0xb3, 0x7f, 0x18, 0x2b,
	0x7e, 0x0b, 0xa5, 0xf4, 0xcd, 0x37, 0x7a, 0x31, 0x37, 0x8a, 0x1d, 0x73, 0xe3, 0x43, 0x2d, 0x77,
	0xb2, 0xcf, 0x37, 0xf1, 0xe7, 0x88, 0xe3, 0x2b, 0xb8, 0x68, 0x31, 0x66, 0xb1, 0x87, 0xae, 0x39,
	0x1d, 0x39, 0x63, 0x8b, 0x59, 0x13, 0x13, 0x9d, 0x60, 0x0c, 0xd5, 0x1d, 0x4c, 0xad, 0xfe, 0x1d,
	0x43, 0x0a, 0xbe, 0x80, 0xf3, 0x1d, 0x36, 0x30, 0x7b, 0x0c, 0x15, 0xf6, 0xa0, 0x76, 0xab, 0x73,
	0x8f, 0xd4, 0x9b, 0x9f, 0x0a, 0x5c, 0xe6, 0xbc, 0x92, 0x8f, 0xf6, 0x05, 0x5c, 0xd9, 0x8e, 0x3d,
	0x74, 0x98, 0x43, 0xa7, 0xec, 0x71, 0x64, 0x4e, 0x1f, 0xec, 0x7b, 0xdb, 0xf9, 0x68, 0xa3, 0x93,
	0x43, 0x6a, 0x62, 0xde, 0x59, 0x9d, 0x81, 0x89, 0x14, 0xfc, 0x12, 0xf4, 0x7d, 0x6a, 0x64, 0x76,
	0xcd, 0x31, 0xa3, 0x56, 0xcb, 0x46, 0x05, 0x7c, 0x0d, 0xb5, 0x67, 0x0f, 0x36, 0x91, 0xfa, 0x47,
	0xee, 0x16, 0x9d, 0xe2, 0x1a, 0xe0, 0x7d, 0x8e, 0x3a, 0x4e, 0x0f, 0x9d, 0x1d, 0xe2, 0x63, 0xb3,
	0x35, 0x40, 0xc5, 0xc3, 0x80, 0x6d, 0xab, 0xf3, 0x28, 0x03, 0x6a, 0xed, 0x37, 0xf0, 0xda, 0x13,
	0x01, 0x99, 0x73, 0xbe, 0x5e, 0xac, 0xfc, 0x30, 0x4e, 0x9e, 0xe4, 0x2e, 0x91, 0xf9, 0x82, 0x04,
	0x62, 0xce, 0x57, 0xb3, 0x62, 0xf2, 0x3f, 0xbc, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x8a,
	0x80, 0xec, 0x31, 0x05, 0x00, 0x00,
}
