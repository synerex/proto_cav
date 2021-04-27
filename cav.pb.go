// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: cav.proto

package proto_cav

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ROS robot meta information message
type ROSMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotID   int64                  `protobuf:"varint,1,opt,name=robotID,proto3" json:"robotID,omitempty"`
	RobotName string                 `protobuf:"bytes,2,opt,name=robotName,proto3" json:"robotName,omitempty"`
	Origin    *Point                 `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	FrameID   string                 `protobuf:"bytes,4,opt,name=frameID,proto3" json:"frameID,omitempty"`
	Ts        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *ROSMeta) Reset() {
	*x = ROSMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ROSMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ROSMeta) ProtoMessage() {}

func (x *ROSMeta) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ROSMeta.ProtoReflect.Descriptor instead.
func (*ROSMeta) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{0}
}

func (x *ROSMeta) GetRobotID() int64 {
	if x != nil {
		return x.RobotID
	}
	return 0
}

func (x *ROSMeta) GetRobotName() string {
	if x != nil {
		return x.RobotName
	}
	return ""
}

func (x *ROSMeta) GetOrigin() *Point {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *ROSMeta) GetFrameID() string {
	if x != nil {
		return x.FrameID
	}
	return ""
}

func (x *ROSMeta) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type RobotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ros         *ROSMeta               `protobuf:"bytes,1,opt,name=ros,proto3" json:"ros,omitempty"`
	Velocity    float32                `protobuf:"fixed32,2,opt,name=velocity,proto3" json:"velocity,omitempty"`
	RotVelocity float32                `protobuf:"fixed32,3,opt,name=rotVelocity,proto3" json:"rotVelocity,omitempty"`
	Radius      float32                `protobuf:"fixed32,4,opt,name=radius,proto3" json:"radius,omitempty"`
	Ts          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *RobotInfo) Reset() {
	*x = RobotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotInfo) ProtoMessage() {}

func (x *RobotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotInfo.ProtoReflect.Descriptor instead.
func (*RobotInfo) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{1}
}

func (x *RobotInfo) GetRos() *ROSMeta {
	if x != nil {
		return x.Ros
	}
	return nil
}

func (x *RobotInfo) GetVelocity() float32 {
	if x != nil {
		return x.Velocity
	}
	return 0
}

func (x *RobotInfo) GetRotVelocity() float32 {
	if x != nil {
		return x.RotVelocity
	}
	return 0
}

func (x *RobotInfo) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *RobotInfo) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ros     *ROSMeta               `protobuf:"bytes,1,opt,name=ros,proto3" json:"ros,omitempty"`
	Current *Point                 `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"` //current position
	Ts      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{2}
}

func (x *Position) GetRos() *ROSMeta {
	if x != nil {
		return x.Ros
	}
	return nil
}

func (x *Position) GetCurrent() *Point {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Position) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

// destination request message
type DestinationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId     int64                  `protobuf:"varint,1,opt,name=robotId,proto3" json:"robotId,omitempty"`
	Seq         int64                  `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`                //sequence number
	Radius      float32                `protobuf:"fixed32,3,opt,name=radius,proto3" json:"radius,omitempty"`         //vehicle radius
	Origin      *Point                 `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`           //origin of coordinate axes
	Current     *Point                 `protobuf:"bytes,5,opt,name=current,proto3" json:"current,omitempty"`         //current position
	Destination *Point                 `protobuf:"bytes,6,opt,name=destination,proto3" json:"destination,omitempty"` //desitination position
	Ts          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *DestinationRequest) Reset() {
	*x = DestinationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationRequest) ProtoMessage() {}

func (x *DestinationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationRequest.ProtoReflect.Descriptor instead.
func (*DestinationRequest) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{3}
}

func (x *DestinationRequest) GetRobotId() int64 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

func (x *DestinationRequest) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *DestinationRequest) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *DestinationRequest) GetOrigin() *Point {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DestinationRequest) GetCurrent() *Point {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *DestinationRequest) GetDestination() *Point {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *DestinationRequest) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

// path message
type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId int64        `protobuf:"varint,1,opt,name=robotId,proto3" json:"robotId,omitempty"`
	Seq     int64        `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Origin  *Point       `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Path    []*PathPoint `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"` //path
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{4}
}

func (x *Path) GetRobotId() int64 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

func (x *Path) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Path) GetOrigin() *Point {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Path) GetPath() []*PathPoint {
	if x != nil {
		return x.Path
	}
	return nil
}

// each point message in a path
type PathPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq   int64                  `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Point *Point                 `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
	Ts    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *PathPoint) Reset() {
	*x = PathPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathPoint) ProtoMessage() {}

func (x *PathPoint) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathPoint.ProtoReflect.Descriptor instead.
func (*PathPoint) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{5}
}

func (x *PathPoint) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PathPoint) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *PathPoint) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

// map meta information message
type MapMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq        int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Origin     *Point   `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Height     int64    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Width      int64    `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Resolusion float32  `protobuf:"fixed32,5,opt,name=resolusion,proto3" json:"resolusion,omitempty"`
	Objects    []*Point `protobuf:"bytes,6,rep,name=objects,proto3" json:"objects,omitempty"` //object list in tha map
}

func (x *MapMeta) Reset() {
	*x = MapMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMeta) ProtoMessage() {}

func (x *MapMeta) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMeta.ProtoReflect.Descriptor instead.
func (*MapMeta) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{6}
}

func (x *MapMeta) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *MapMeta) GetOrigin() *Point {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *MapMeta) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *MapMeta) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *MapMeta) GetResolusion() float32 {
	if x != nil {
		return x.Resolusion
	}
	return 0
}

func (x *MapMeta) GetObjects() []*Point {
	if x != nil {
		return x.Objects
	}
	return nil
}

// grid map message
type GridMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq        int64   `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Origin     *Point  `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Resolution float32 `protobuf:"fixed32,3,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Width      uint64  `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`   //xwidth
	Height     uint64  `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"` //ywidth
	Gridmap    []*Grid `protobuf:"bytes,6,rep,name=gridmap,proto3" json:"gridmap,omitempty"`
}

func (x *GridMap) Reset() {
	*x = GridMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GridMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GridMap) ProtoMessage() {}

func (x *GridMap) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GridMap.ProtoReflect.Descriptor instead.
func (*GridMap) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{7}
}

func (x *GridMap) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *GridMap) GetOrigin() *Point {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *GridMap) GetResolution() float32 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

func (x *GridMap) GetWidth() uint64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GridMap) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GridMap) GetGridmap() []*Grid {
	if x != nil {
		return x.Gridmap
	}
	return nil
}

// a single grid in the grid map message
type Grid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Ix        uint64 `protobuf:"varint,2,opt,name=ix,proto3" json:"ix,omitempty"`               //index of x
	Iy        uint64 `protobuf:"varint,3,opt,name=iy,proto3" json:"iy,omitempty"`               // index of y
	Occupancy int64  `protobuf:"varint,4,opt,name=occupancy,proto3" json:"occupancy,omitempty"` //occupancy data
}

func (x *Grid) Reset() {
	*x = Grid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grid) ProtoMessage() {}

func (x *Grid) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grid.ProtoReflect.Descriptor instead.
func (*Grid) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{8}
}

func (x *Grid) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Grid) GetIx() uint64 {
	if x != nil {
		return x.Ix
	}
	return 0
}

func (x *Grid) GetIy() uint64 {
	if x != nil {
		return x.Iy
	}
	return 0
}

func (x *Grid) GetOccupancy() int64 {
	if x != nil {
		return x.Occupancy
	}
	return 0
}

// 2 axes point message
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cav_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_cav_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_cav_proto_rawDescGZIP(), []int{9}
}

func (x *Point) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_cav_proto protoreflect.FileDescriptor

var file_cav_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x76,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x07, 0x52, 0x4f, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22,
	0xad, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x03, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61, 0x76,
	0x2e, 0x52, 0x4f, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x03, 0x72, 0x6f, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x6f, 0x74,
	0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x72, 0x6f, 0x74, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x64,
	0x69, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22,
	0x7c, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x03, 0x72,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x52,
	0x4f, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x03, 0x72, 0x6f, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63,
	0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0xfc, 0x01,
	0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x7a, 0x0a, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x6b, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x68,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x07,
	0x47, 0x72, 0x69, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x76, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x67,
	0x72, 0x69, 0x64, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63,
	0x61, 0x76, 0x2e, 0x47, 0x72, 0x69, 0x64, 0x52, 0x07, 0x67, 0x72, 0x69, 0x64, 0x6d, 0x61, 0x70,
	0x22, 0x5a, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x78, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x22, 0x23, 0x0a, 0x05,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x79, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x61,
	0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cav_proto_rawDescOnce sync.Once
	file_cav_proto_rawDescData = file_cav_proto_rawDesc
)

func file_cav_proto_rawDescGZIP() []byte {
	file_cav_proto_rawDescOnce.Do(func() {
		file_cav_proto_rawDescData = protoimpl.X.CompressGZIP(file_cav_proto_rawDescData)
	})
	return file_cav_proto_rawDescData
}

var file_cav_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cav_proto_goTypes = []interface{}{
	(*ROSMeta)(nil),               // 0: cav.ROSMeta
	(*RobotInfo)(nil),             // 1: cav.RobotInfo
	(*Position)(nil),              // 2: cav.Position
	(*DestinationRequest)(nil),    // 3: cav.DestinationRequest
	(*Path)(nil),                  // 4: cav.Path
	(*PathPoint)(nil),             // 5: cav.PathPoint
	(*MapMeta)(nil),               // 6: cav.MapMeta
	(*GridMap)(nil),               // 7: cav.GridMap
	(*Grid)(nil),                  // 8: cav.Grid
	(*Point)(nil),                 // 9: cav.Point
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_cav_proto_depIdxs = []int32{
	9,  // 0: cav.ROSMeta.origin:type_name -> cav.Point
	10, // 1: cav.ROSMeta.ts:type_name -> google.protobuf.Timestamp
	0,  // 2: cav.RobotInfo.ros:type_name -> cav.ROSMeta
	10, // 3: cav.RobotInfo.ts:type_name -> google.protobuf.Timestamp
	0,  // 4: cav.Position.ros:type_name -> cav.ROSMeta
	9,  // 5: cav.Position.current:type_name -> cav.Point
	10, // 6: cav.Position.ts:type_name -> google.protobuf.Timestamp
	9,  // 7: cav.DestinationRequest.origin:type_name -> cav.Point
	9,  // 8: cav.DestinationRequest.current:type_name -> cav.Point
	9,  // 9: cav.DestinationRequest.destination:type_name -> cav.Point
	10, // 10: cav.DestinationRequest.ts:type_name -> google.protobuf.Timestamp
	9,  // 11: cav.Path.origin:type_name -> cav.Point
	5,  // 12: cav.Path.path:type_name -> cav.PathPoint
	9,  // 13: cav.PathPoint.point:type_name -> cav.Point
	10, // 14: cav.PathPoint.ts:type_name -> google.protobuf.Timestamp
	9,  // 15: cav.MapMeta.origin:type_name -> cav.Point
	9,  // 16: cav.MapMeta.objects:type_name -> cav.Point
	9,  // 17: cav.GridMap.origin:type_name -> cav.Point
	8,  // 18: cav.GridMap.gridmap:type_name -> cav.Grid
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_cav_proto_init() }
func file_cav_proto_init() {
	if File_cav_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cav_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ROSMeta); i {
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
		file_cav_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotInfo); i {
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
		file_cav_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_cav_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationRequest); i {
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
		file_cav_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_cav_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathPoint); i {
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
		file_cav_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMeta); i {
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
		file_cav_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GridMap); i {
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
		file_cav_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grid); i {
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
		file_cav_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_cav_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cav_proto_goTypes,
		DependencyIndexes: file_cav_proto_depIdxs,
		MessageInfos:      file_cav_proto_msgTypes,
	}.Build()
	File_cav_proto = out.File
	file_cav_proto_rawDesc = nil
	file_cav_proto_goTypes = nil
	file_cav_proto_depIdxs = nil
}
