// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enum.proto

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

type TaskStatus int32

const (
	TaskStatus_Task_Status_Unknown    TaskStatus = 0
	TaskStatus_Task_Status_Created    TaskStatus = 1
	TaskStatus_Task_Status_Processing TaskStatus = 2
	TaskStatus_Task_Status_Finished   TaskStatus = 3
	TaskStatus_Task_Status_Deleted    TaskStatus = 4
	TaskStatus_Task_Status_Outdate    TaskStatus = 5
	TaskStatus_Task_Status_Paused     TaskStatus = 6
	TaskStatus_Task_Status_Stopped    TaskStatus = 7
)

var TaskStatus_name = map[int32]string{
	0: "Task_Status_Unknown",
	1: "Task_Status_Created",
	2: "Task_Status_Processing",
	3: "Task_Status_Finished",
	4: "Task_Status_Deleted",
	5: "Task_Status_Outdate",
	6: "Task_Status_Paused",
	7: "Task_Status_Stopped",
}
var TaskStatus_value = map[string]int32{
	"Task_Status_Unknown":    0,
	"Task_Status_Created":    1,
	"Task_Status_Processing": 2,
	"Task_Status_Finished":   3,
	"Task_Status_Deleted":    4,
	"Task_Status_Outdate":    5,
	"Task_Status_Paused":     6,
	"Task_Status_Stopped":    7,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_bae4f56ceda19e2d, []int{0}
}

type SwitcherStatus int32

const (
	SwitcherStatus_Switcher_Status_Unknown SwitcherStatus = 0
	SwitcherStatus_Switcher_Status_Started SwitcherStatus = 1
	SwitcherStatus_Switcher_Status_Stopped SwitcherStatus = 2
	SwitcherStatus_Switcher_Status_Paused  SwitcherStatus = 3
)

var SwitcherStatus_name = map[int32]string{
	0: "Switcher_Status_Unknown",
	1: "Switcher_Status_Started",
	2: "Switcher_Status_Stopped",
	3: "Switcher_Status_Paused",
}
var SwitcherStatus_value = map[string]int32{
	"Switcher_Status_Unknown": 0,
	"Switcher_Status_Started": 1,
	"Switcher_Status_Stopped": 2,
	"Switcher_Status_Paused":  3,
}

func (x SwitcherStatus) String() string {
	return proto.EnumName(SwitcherStatus_name, int32(x))
}
func (SwitcherStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_bae4f56ceda19e2d, []int{1}
}

type CutboardType int32

const (
	CutboardType_CUTBOARD_TYPE_UNKNOWN CutboardType = 0
	CutboardType_CUTBOARD_TYPE_VEHICLE CutboardType = 1
	CutboardType_CUTBOARD_TYPE_SYMBOL  CutboardType = 2
	CutboardType_CUTBOARD_TYPE_PLATE   CutboardType = 3
)

var CutboardType_name = map[int32]string{
	0: "CUTBOARD_TYPE_UNKNOWN",
	1: "CUTBOARD_TYPE_VEHICLE",
	2: "CUTBOARD_TYPE_SYMBOL",
	3: "CUTBOARD_TYPE_PLATE",
}
var CutboardType_value = map[string]int32{
	"CUTBOARD_TYPE_UNKNOWN": 0,
	"CUTBOARD_TYPE_VEHICLE": 1,
	"CUTBOARD_TYPE_SYMBOL":  2,
	"CUTBOARD_TYPE_PLATE":   3,
}

func (x CutboardType) String() string {
	return proto.EnumName(CutboardType_name, int32(x))
}
func (CutboardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_bae4f56ceda19e2d, []int{2}
}

type TaskSourceType int32

const (
	TaskSourceType_TASK_TYPE_UNKNOWN TaskSourceType = 0
	// 静态图片任务
	TaskSourceType_TASK_TYPE_IMAGE TaskSourceType = 1
	// 视频文件任务
	TaskSourceType_TASK_TYPE_VIDEO TaskSourceType = 2
	// 视频流任务
	TaskSourceType_TASK_TYPE_STREAM TaskSourceType = 3
	// 动态图片任务
	TaskSourceType_TASK_TYPE_CAPTURE TaskSourceType = 4
	// 平台视频流
	TaskSourceType_TASK_TYPE_PLATFORM_STREAM TaskSourceType = 5
)

var TaskSourceType_name = map[int32]string{
	0: "TASK_TYPE_UNKNOWN",
	1: "TASK_TYPE_IMAGE",
	2: "TASK_TYPE_VIDEO",
	3: "TASK_TYPE_STREAM",
	4: "TASK_TYPE_CAPTURE",
	5: "TASK_TYPE_PLATFORM_STREAM",
}
var TaskSourceType_value = map[string]int32{
	"TASK_TYPE_UNKNOWN":         0,
	"TASK_TYPE_IMAGE":           1,
	"TASK_TYPE_VIDEO":           2,
	"TASK_TYPE_STREAM":          3,
	"TASK_TYPE_CAPTURE":         4,
	"TASK_TYPE_PLATFORM_STREAM": 5,
}

func (x TaskSourceType) String() string {
	return proto.EnumName(TaskSourceType_name, int32(x))
}
func (TaskSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_bae4f56ceda19e2d, []int{3}
}

// 图片类型
type ImageType int32

const (
	ImageType_ImageTypeUnknown ImageType = 0
	// 身份证图片
	ImageType_ImageTypeId ImageType = 1
	// 抓拍图片
	ImageType_ImageTypeCapture ImageType = 2
	// 用户上传图片
	ImageType_ImageTypeUpload ImageType = 3
	// 全景图
	ImageType_ImageTypePanorama ImageType = 100
	// 人脸大图
	ImageType_ImageTypeLargerFace ImageType = 101
	// 人脸小图
	ImageType_ImageTypeSmallerFace ImageType = 102
	// 车/人/非大图
	ImageType_ImageTypeLargerVehicle ImageType = 103
	// 车/人/非小图
	ImageType_ImageTypeSmallerVehicle ImageType = 104
)

var ImageType_name = map[int32]string{
	0:   "ImageTypeUnknown",
	1:   "ImageTypeId",
	2:   "ImageTypeCapture",
	3:   "ImageTypeUpload",
	100: "ImageTypePanorama",
	101: "ImageTypeLargerFace",
	102: "ImageTypeSmallerFace",
	103: "ImageTypeLargerVehicle",
	104: "ImageTypeSmallerVehicle",
}
var ImageType_value = map[string]int32{
	"ImageTypeUnknown":        0,
	"ImageTypeId":             1,
	"ImageTypeCapture":        2,
	"ImageTypeUpload":         3,
	"ImageTypePanorama":       100,
	"ImageTypeLargerFace":     101,
	"ImageTypeSmallerFace":    102,
	"ImageTypeLargerVehicle":  103,
	"ImageTypeSmallerVehicle": 104,
}

func (x ImageType) String() string {
	return proto.EnumName(ImageType_name, int32(x))
}
func (ImageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_bae4f56ceda19e2d, []int{4}
}

func init() {
	proto.RegisterEnum("TaskStatus", TaskStatus_name, TaskStatus_value)
	proto.RegisterEnum("SwitcherStatus", SwitcherStatus_name, SwitcherStatus_value)
	proto.RegisterEnum("CutboardType", CutboardType_name, CutboardType_value)
	proto.RegisterEnum("TaskSourceType", TaskSourceType_name, TaskSourceType_value)
	proto.RegisterEnum("ImageType", ImageType_name, ImageType_value)
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_enum_bae4f56ceda19e2d) }

var fileDescriptor_enum_bae4f56ceda19e2d = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0x83, 0x9d, 0xa4, 0xea, 0xb4, 0x4a, 0xa6, 0xc4, 0x49, 0x9a, 0x56, 0x55, 0xd5, 0xab,
	0x0f, 0xbe, 0xf4, 0x09, 0x30, 0xc6, 0xad, 0x15, 0xdb, 0x58, 0x06, 0xbb, 0xca, 0xc9, 0xda, 0xb0,
	0x53, 0x8c, 0x02, 0x2c, 0x5a, 0x16, 0x59, 0x3d, 0xf7, 0x39, 0xfa, 0x5c, 0xbd, 0xf6, 0x51, 0xaa,
	0xf5, 0x1f, 0x30, 0x24, 0x37, 0xe6, 0xfb, 0x31, 0x33, 0xdf, 0xcc, 0x00, 0x00, 0xa5, 0x45, 0xd2,
	0xcb, 0xa4, 0x50, 0xa2, 0xfb, 0xd7, 0x00, 0xf0, 0x59, 0xfe, 0xe4, 0x29, 0xa6, 0x8a, 0xdc, 0xbc,
	0x85, 0x2b, 0x1d, 0xad, 0x76, 0xe1, 0x6a, 0x91, 0x3e, 0xa5, 0x62, 0x93, 0xe2, 0x49, 0x13, 0xd8,
	0x92, 0x98, 0x22, 0x8e, 0x86, 0xf9, 0x01, 0x6e, 0x8e, 0xc1, 0x4c, 0x8a, 0x80, 0xf2, 0x3c, 0x4a,
	0x43, 0x6c, 0x99, 0xef, 0xa1, 0x73, 0xcc, 0x86, 0x51, 0x1a, 0xe5, 0x6b, 0xe2, 0xd8, 0x6e, 0x96,
	0x1b, 0x50, 0x4c, 0xba, 0xdc, 0x69, 0x13, 0xb8, 0x85, 0xe2, 0x4c, 0x11, 0x9e, 0x99, 0x37, 0x60,
	0xd6, 0xfa, 0xb0, 0x22, 0x27, 0x8e, 0xe7, 0xcd, 0x04, 0x4f, 0x89, 0x2c, 0x23, 0x8e, 0xaf, 0xba,
	0xbf, 0x0d, 0xb8, 0xf0, 0x36, 0x91, 0x0a, 0xd6, 0x24, 0xf7, 0xd3, 0x7d, 0x84, 0xdb, 0x83, 0xf2,
	0x7c, 0xc2, 0x17, 0xa0, 0xa7, 0x98, 0xdc, 0x4d, 0xf9, 0x22, 0xdc, 0x75, 0x6a, 0xe9, 0x15, 0x34,
	0xe1, 0xde, 0x5e, 0xbb, 0xbb, 0x81, 0xb7, 0x76, 0xa1, 0x1e, 0x05, 0x93, 0xdc, 0xff, 0x95, 0x91,
	0x79, 0x07, 0xd7, 0xf6, 0xc2, 0xef, 0xbb, 0xd6, 0x7c, 0xb0, 0xf2, 0x1f, 0x66, 0xce, 0x6a, 0x31,
	0xbd, 0x9f, 0xba, 0x3f, 0xa6, 0x78, 0xf2, 0x1c, 0x2d, 0x9d, 0xef, 0x23, 0x7b, 0xec, 0xa0, 0xa1,
	0x17, 0x59, 0x47, 0xde, 0xc3, 0xa4, 0xef, 0x8e, 0xb1, 0xa5, 0xc7, 0xaf, 0x93, 0xd9, 0xd8, 0xf2,
	0x1d, 0x6c, 0x77, 0xff, 0x18, 0x70, 0xb1, 0x3d, 0xac, 0x28, 0x64, 0x40, 0xdb, 0xde, 0xd7, 0xf0,
	0xce, 0xb7, 0xbc, 0xfb, 0x66, 0xdf, 0x2b, 0xb8, 0xac, 0xe4, 0xd1, 0xc4, 0xfa, 0xa6, 0x3b, 0xd6,
	0xc4, 0xe5, 0x68, 0xe0, 0xb8, 0xd8, 0x32, 0x3b, 0x80, 0x95, 0xe8, 0xf9, 0x73, 0xc7, 0x9a, 0x60,
	0xbb, 0x5e, 0xd6, 0xb6, 0x66, 0xfe, 0x62, 0xee, 0xe0, 0xa9, 0xf9, 0x09, 0xee, 0x2a, 0x59, 0xbb,
	0x1a, 0xba, 0xf3, 0xc9, 0x21, 0xeb, 0xac, 0xfb, 0xcf, 0x80, 0xd7, 0xa3, 0x84, 0x85, 0x3b, 0x6b,
	0x1d, 0xc0, 0x32, 0xa8, 0x4e, 0x72, 0x09, 0x6f, 0x4a, 0x75, 0xa4, 0xcf, 0x70, 0xfc, 0x9a, 0xcd,
	0x32, 0x55, 0x48, 0xc2, 0x96, 0xf6, 0x5a, 0x25, 0x67, 0xb1, 0x60, 0x7c, 0xe7, 0xaa, 0x14, 0x67,
	0x2c, 0x15, 0x92, 0x25, 0x0c, 0xb9, 0xde, 0x57, 0x29, 0x8f, 0x99, 0x0c, 0x49, 0x0e, 0x59, 0x40,
	0x48, 0x7a, 0xc5, 0x25, 0xf0, 0x12, 0x16, 0xc7, 0x7b, 0xf2, 0x53, 0x9f, 0xb7, 0x91, 0xb2, 0xa4,
	0x75, 0x14, 0xc4, 0x84, 0xa1, 0xfe, 0x2e, 0x9a, 0x59, 0x07, 0xb8, 0xee, 0x7f, 0x81, 0xcf, 0x81,
	0x48, 0x7a, 0x9c, 0x28, 0x0b, 0xe3, 0x28, 0x55, 0xdb, 0x27, 0xce, 0x14, 0xeb, 0xf1, 0xb0, 0x97,
	0x08, 0x4e, 0xf1, 0xe3, 0xf9, 0xf6, 0x2f, 0xfc, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x92,
	0xae, 0x82, 0x93, 0x03, 0x00, 0x00,
}
