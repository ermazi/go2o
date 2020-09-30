// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/shipment.proto

package proto // import "."

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

// 覆盖区域
type SCoverageValue struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Lng                  float64  `protobuf:"fixed64,3,opt,name=Lng,proto3" json:"Lng,omitempty"`
	Lat                  float64  `protobuf:"fixed64,4,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Radius               int32    `protobuf:"varint,5,opt,name=Radius,proto3" json:"Radius,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	AreaId               int64    `protobuf:"varint,7,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SCoverageValue) Reset()         { *m = SCoverageValue{} }
func (m *SCoverageValue) String() string { return proto.CompactTextString(m) }
func (*SCoverageValue) ProtoMessage()    {}
func (*SCoverageValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{0}
}
func (m *SCoverageValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCoverageValue.Unmarshal(m, b)
}
func (m *SCoverageValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCoverageValue.Marshal(b, m, deterministic)
}
func (dst *SCoverageValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCoverageValue.Merge(dst, src)
}
func (m *SCoverageValue) XXX_Size() int {
	return xxx_messageInfo_SCoverageValue.Size(m)
}
func (m *SCoverageValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SCoverageValue.DiscardUnknown(m)
}

var xxx_messageInfo_SCoverageValue proto.InternalMessageInfo

func (m *SCoverageValue) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SCoverageValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SCoverageValue) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *SCoverageValue) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *SCoverageValue) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *SCoverageValue) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SCoverageValue) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type ShipmentOrderListResponse struct {
	Value                []*SShipmentOrder `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ShipmentOrderListResponse) Reset()         { *m = ShipmentOrderListResponse{} }
func (m *ShipmentOrderListResponse) String() string { return proto.CompactTextString(m) }
func (*ShipmentOrderListResponse) ProtoMessage()    {}
func (*ShipmentOrderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{1}
}
func (m *ShipmentOrderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentOrderListResponse.Unmarshal(m, b)
}
func (m *ShipmentOrderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentOrderListResponse.Marshal(b, m, deterministic)
}
func (dst *ShipmentOrderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentOrderListResponse.Merge(dst, src)
}
func (m *ShipmentOrderListResponse) XXX_Size() int {
	return xxx_messageInfo_ShipmentOrderListResponse.Size(m)
}
func (m *ShipmentOrderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentOrderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentOrderListResponse proto.InternalMessageInfo

func (m *ShipmentOrderListResponse) GetValue() []*SShipmentOrder {
	if m != nil {
		return m.Value
	}
	return nil
}

// 发货单
type SShipmentOrder struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 订单编号
	OrderId int64 `protobuf:"varint,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	// 子订单编号
	SubOrderId int64 `protobuf:"varint,3,opt,name=SubOrderId,proto3" json:"SubOrderId,omitempty"`
	// 快递SP编号
	ExpressSpId int64 `protobuf:"varint,4,opt,name=ExpressSpId,proto3" json:"ExpressSpId,omitempty"`
	// 快递SP单号
	ShipOrderNo string `protobuf:"bytes,5,opt,name=ShipOrderNo,proto3" json:"ShipOrderNo,omitempty"`
	// 物流日志
	ShipmentLog string `protobuf:"bytes,6,opt,name=ShipmentLog,proto3" json:"ShipmentLog,omitempty"`
	// 运费
	Amount float64 `protobuf:"fixed64,7,opt,name=Amount,proto3" json:"Amount,omitempty"`
	// 实际运费
	FinalAmount float64 `protobuf:"fixed64,8,opt,name=FinalAmount,proto3" json:"FinalAmount,omitempty"`
	// 发货时间
	ShipTime int64 `protobuf:"varint,9,opt,name=ShipTime,proto3" json:"ShipTime,omitempty"`
	// 状态
	State int32 `protobuf:"varint,10,opt,name=State,proto3" json:"State,omitempty"`
	// 更新时间
	UpdateTime int64 `protobuf:"varint,11,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// 配送项目
	Items                []*SShipmentItem `protobuf:"bytes,12,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SShipmentOrder) Reset()         { *m = SShipmentOrder{} }
func (m *SShipmentOrder) String() string { return proto.CompactTextString(m) }
func (*SShipmentOrder) ProtoMessage()    {}
func (*SShipmentOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{2}
}
func (m *SShipmentOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipmentOrder.Unmarshal(m, b)
}
func (m *SShipmentOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipmentOrder.Marshal(b, m, deterministic)
}
func (dst *SShipmentOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipmentOrder.Merge(dst, src)
}
func (m *SShipmentOrder) XXX_Size() int {
	return xxx_messageInfo_SShipmentOrder.Size(m)
}
func (m *SShipmentOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipmentOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SShipmentOrder proto.InternalMessageInfo

func (m *SShipmentOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShipmentOrder) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *SShipmentOrder) GetSubOrderId() int64 {
	if m != nil {
		return m.SubOrderId
	}
	return 0
}

func (m *SShipmentOrder) GetExpressSpId() int64 {
	if m != nil {
		return m.ExpressSpId
	}
	return 0
}

func (m *SShipmentOrder) GetShipOrderNo() string {
	if m != nil {
		return m.ShipOrderNo
	}
	return ""
}

func (m *SShipmentOrder) GetShipmentLog() string {
	if m != nil {
		return m.ShipmentLog
	}
	return ""
}

func (m *SShipmentOrder) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SShipmentOrder) GetFinalAmount() float64 {
	if m != nil {
		return m.FinalAmount
	}
	return 0
}

func (m *SShipmentOrder) GetShipTime() int64 {
	if m != nil {
		return m.ShipTime
	}
	return 0
}

func (m *SShipmentOrder) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SShipmentOrder) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SShipmentOrder) GetItems() []*SShipmentItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 发货单详情
type SShipmentItem struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 商品交易快照编号
	SnapshotId int64 `protobuf:"varint,2,opt,name=SnapshotId,proto3" json:"SnapshotId,omitempty"`
	// 商品数量
	Quantity int32 `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	// 运费
	Amount float64 `protobuf:"fixed64,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	// 实际运费
	FinalAmount          float64  `protobuf:"fixed64,5,opt,name=FinalAmount,proto3" json:"FinalAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShipmentItem) Reset()         { *m = SShipmentItem{} }
func (m *SShipmentItem) String() string { return proto.CompactTextString(m) }
func (*SShipmentItem) ProtoMessage()    {}
func (*SShipmentItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{3}
}
func (m *SShipmentItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipmentItem.Unmarshal(m, b)
}
func (m *SShipmentItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipmentItem.Marshal(b, m, deterministic)
}
func (dst *SShipmentItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipmentItem.Merge(dst, src)
}
func (m *SShipmentItem) XXX_Size() int {
	return xxx_messageInfo_SShipmentItem.Size(m)
}
func (m *SShipmentItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipmentItem.DiscardUnknown(m)
}

var xxx_messageInfo_SShipmentItem proto.InternalMessageInfo

func (m *SShipmentItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShipmentItem) GetSnapshotId() int64 {
	if m != nil {
		return m.SnapshotId
	}
	return 0
}

func (m *SShipmentItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *SShipmentItem) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SShipmentItem) GetFinalAmount() float64 {
	if m != nil {
		return m.FinalAmount
	}
	return 0
}

type LogisticFlowTrackRequest struct {
	ShipperCode          string   `protobuf:"bytes,1,opt,name=shipperCode,proto3" json:"shipperCode,omitempty"`
	LogisticCode         string   `protobuf:"bytes,2,opt,name=logisticCode,proto3" json:"logisticCode,omitempty"`
	Invert               bool     `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogisticFlowTrackRequest) Reset()         { *m = LogisticFlowTrackRequest{} }
func (m *LogisticFlowTrackRequest) String() string { return proto.CompactTextString(m) }
func (*LogisticFlowTrackRequest) ProtoMessage()    {}
func (*LogisticFlowTrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{4}
}
func (m *LogisticFlowTrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogisticFlowTrackRequest.Unmarshal(m, b)
}
func (m *LogisticFlowTrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogisticFlowTrackRequest.Marshal(b, m, deterministic)
}
func (dst *LogisticFlowTrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogisticFlowTrackRequest.Merge(dst, src)
}
func (m *LogisticFlowTrackRequest) XXX_Size() int {
	return xxx_messageInfo_LogisticFlowTrackRequest.Size(m)
}
func (m *LogisticFlowTrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogisticFlowTrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogisticFlowTrackRequest proto.InternalMessageInfo

func (m *LogisticFlowTrackRequest) GetShipperCode() string {
	if m != nil {
		return m.ShipperCode
	}
	return ""
}

func (m *LogisticFlowTrackRequest) GetLogisticCode() string {
	if m != nil {
		return m.LogisticCode
	}
	return ""
}

func (m *LogisticFlowTrackRequest) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

type OrderLogisticTrackRequest struct {
	ShipOrderId          int64    `protobuf:"zigzag64,1,opt,name=shipOrderId,proto3" json:"shipOrderId,omitempty"`
	Invert               bool     `protobuf:"varint,2,opt,name=invert,proto3" json:"invert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderLogisticTrackRequest) Reset()         { *m = OrderLogisticTrackRequest{} }
func (m *OrderLogisticTrackRequest) String() string { return proto.CompactTextString(m) }
func (*OrderLogisticTrackRequest) ProtoMessage()    {}
func (*OrderLogisticTrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{5}
}
func (m *OrderLogisticTrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLogisticTrackRequest.Unmarshal(m, b)
}
func (m *OrderLogisticTrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLogisticTrackRequest.Marshal(b, m, deterministic)
}
func (dst *OrderLogisticTrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLogisticTrackRequest.Merge(dst, src)
}
func (m *OrderLogisticTrackRequest) XXX_Size() int {
	return xxx_messageInfo_OrderLogisticTrackRequest.Size(m)
}
func (m *OrderLogisticTrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLogisticTrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLogisticTrackRequest proto.InternalMessageInfo

func (m *OrderLogisticTrackRequest) GetShipOrderId() int64 {
	if m != nil {
		return m.ShipOrderId
	}
	return 0
}

func (m *OrderLogisticTrackRequest) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

// 发货单追踪
type SShipOrderTrack struct {
	// 返回状态码
	Code int32 `protobuf:"zigzag32,1,opt,name=Code,proto3" json:"Code,omitempty"`
	// 返回错误信息
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	// 物流单号
	LogisticCode string `protobuf:"bytes,3,opt,name=LogisticCode,proto3" json:"LogisticCode,omitempty"`
	// 承运商代码
	ShipperCode string `protobuf:"bytes,4,opt,name=ShipperCode,proto3" json:"ShipperCode,omitempty"`
	// * 承运商名称
	ShipperName string `protobuf:"bytes,5,opt,name=ShipperName,proto3" json:"ShipperName,omitempty"`
	// 发货状态
	ShipState string `protobuf:"bytes,6,opt,name=ShipState,proto3" json:"ShipState,omitempty"`
	// 更新时间
	UpdateTime int64 `protobuf:"zigzag64,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// 包含发货单流
	Flows                []*SShipFlow `protobuf:"bytes,8,rep,name=Flows,proto3" json:"Flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SShipOrderTrack) Reset()         { *m = SShipOrderTrack{} }
func (m *SShipOrderTrack) String() string { return proto.CompactTextString(m) }
func (*SShipOrderTrack) ProtoMessage()    {}
func (*SShipOrderTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{6}
}
func (m *SShipOrderTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipOrderTrack.Unmarshal(m, b)
}
func (m *SShipOrderTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipOrderTrack.Marshal(b, m, deterministic)
}
func (dst *SShipOrderTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipOrderTrack.Merge(dst, src)
}
func (m *SShipOrderTrack) XXX_Size() int {
	return xxx_messageInfo_SShipOrderTrack.Size(m)
}
func (m *SShipOrderTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipOrderTrack.DiscardUnknown(m)
}

var xxx_messageInfo_SShipOrderTrack proto.InternalMessageInfo

func (m *SShipOrderTrack) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SShipOrderTrack) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SShipOrderTrack) GetLogisticCode() string {
	if m != nil {
		return m.LogisticCode
	}
	return ""
}

func (m *SShipOrderTrack) GetShipperCode() string {
	if m != nil {
		return m.ShipperCode
	}
	return ""
}

func (m *SShipOrderTrack) GetShipperName() string {
	if m != nil {
		return m.ShipperName
	}
	return ""
}

func (m *SShipOrderTrack) GetShipState() string {
	if m != nil {
		return m.ShipState
	}
	return ""
}

func (m *SShipOrderTrack) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SShipOrderTrack) GetFlows() []*SShipFlow {
	if m != nil {
		return m.Flows
	}
	return nil
}

// 发货流
type SShipFlow struct {
	// 记录标题
	Subject string `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// 记录时间
	CreateTime string `protobuf:"bytes,2,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// 备注
	Remark               string   `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShipFlow) Reset()         { *m = SShipFlow{} }
func (m *SShipFlow) String() string { return proto.CompactTextString(m) }
func (*SShipFlow) ProtoMessage()    {}
func (*SShipFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_a83b16b3d7af732f, []int{7}
}
func (m *SShipFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipFlow.Unmarshal(m, b)
}
func (m *SShipFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipFlow.Marshal(b, m, deterministic)
}
func (dst *SShipFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipFlow.Merge(dst, src)
}
func (m *SShipFlow) XXX_Size() int {
	return xxx_messageInfo_SShipFlow.Size(m)
}
func (m *SShipFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipFlow.DiscardUnknown(m)
}

var xxx_messageInfo_SShipFlow proto.InternalMessageInfo

func (m *SShipFlow) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SShipFlow) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *SShipFlow) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*SCoverageValue)(nil), "SCoverageValue")
	proto.RegisterType((*ShipmentOrderListResponse)(nil), "ShipmentOrderListResponse")
	proto.RegisterType((*SShipmentOrder)(nil), "SShipmentOrder")
	proto.RegisterType((*SShipmentItem)(nil), "SShipmentItem")
	proto.RegisterType((*LogisticFlowTrackRequest)(nil), "LogisticFlowTrackRequest")
	proto.RegisterType((*OrderLogisticTrackRequest)(nil), "OrderLogisticTrackRequest")
	proto.RegisterType((*SShipOrderTrack)(nil), "SShipOrderTrack")
	proto.RegisterType((*SShipFlow)(nil), "SShipFlow")
}

func init() { proto.RegisterFile("message/shipment.proto", fileDescriptor_shipment_a83b16b3d7af732f) }

var fileDescriptor_shipment_a83b16b3d7af732f = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0xb8, 0x89, 0xa7, 0x25, 0x85, 0x15, 0xaa, 0xb6, 0x08, 0x21, 0xcb, 0x02, 0x29,
	0xa7, 0x22, 0xc1, 0x91, 0x53, 0x5b, 0x51, 0x29, 0x52, 0x28, 0x62, 0xdd, 0x72, 0x40, 0xe2, 0xb0,
	0xad, 0x47, 0xa9, 0x69, 0xfc, 0x83, 0x77, 0x5d, 0xca, 0x13, 0xf0, 0x04, 0xbc, 0x00, 0x8f, 0xc8,
	0x13, 0xa0, 0x9d, 0x5d, 0x27, 0x1b, 0x82, 0x38, 0xd9, 0xdf, 0x37, 0x33, 0x3b, 0x33, 0x9f, 0x3f,
	0x2f, 0x1c, 0x94, 0xa8, 0x94, 0x5c, 0xe0, 0x4b, 0x75, 0x53, 0x34, 0x25, 0x56, 0xfa, 0xa8, 0x69,
	0x6b, 0x5d, 0xa7, 0xbf, 0x02, 0x98, 0x64, 0xa7, 0xf5, 0x1d, 0xb6, 0x72, 0x81, 0x1f, 0xe5, 0xb2,
	0x43, 0x36, 0x81, 0xc1, 0x2c, 0xe7, 0x41, 0x12, 0x4c, 0x43, 0x31, 0x98, 0xe5, 0x8c, 0xc1, 0xf0,
	0x5c, 0x96, 0xc8, 0x07, 0x49, 0x30, 0x8d, 0x05, 0xbd, 0xb3, 0x87, 0x10, 0xce, 0xab, 0x05, 0x0f,
	0x93, 0x60, 0x1a, 0x08, 0xf3, 0x4a, 0x8c, 0xd4, 0x7c, 0xe8, 0x18, 0xa9, 0xd9, 0x01, 0xec, 0x08,
	0x99, 0x17, 0x9d, 0xe2, 0x51, 0x12, 0x4c, 0x23, 0xe1, 0x10, 0xe3, 0x30, 0x3a, 0xce, 0xf3, 0x16,
	0x95, 0xe2, 0x3b, 0x74, 0x64, 0x0f, 0x4d, 0xc5, 0x71, 0x8b, 0x72, 0x96, 0xf3, 0x11, 0x75, 0x77,
	0x28, 0x3d, 0x81, 0xc3, 0xcc, 0x8d, 0xfd, 0xbe, 0xcd, 0xb1, 0x9d, 0x17, 0x4a, 0x0b, 0x54, 0x4d,
	0x5d, 0x29, 0x64, 0x2f, 0x20, 0xa2, 0xb9, 0x79, 0x90, 0x84, 0xd3, 0xdd, 0x57, 0xfb, 0x47, 0xd9,
	0x46, 0xae, 0xb0, 0xd1, 0xf4, 0xf7, 0x00, 0x26, 0x9b, 0x91, 0xad, 0x45, 0x39, 0x8c, 0x28, 0x30,
	0xcb, 0x69, 0xd7, 0x50, 0xf4, 0x90, 0x3d, 0x03, 0xc8, 0xba, 0xab, 0x3e, 0x18, 0x52, 0xd0, 0x63,
	0x58, 0x02, 0xbb, 0x6f, 0xef, 0x1b, 0xb3, 0x43, 0xd6, 0xcc, 0x72, 0x12, 0x21, 0x14, 0x3e, 0x65,
	0x32, 0x4c, 0x73, 0x2a, 0x38, 0xaf, 0x49, 0x91, 0x58, 0xf8, 0x54, 0x9f, 0x61, 0xc6, 0x9b, 0xd7,
	0x0b, 0x27, 0x8d, 0x4f, 0x91, 0x3c, 0x65, 0xdd, 0x55, 0x9a, 0xe4, 0x09, 0x84, 0x43, 0xa6, 0xf2,
	0xac, 0xa8, 0xe4, 0xd2, 0x05, 0xc7, 0x14, 0xf4, 0x29, 0xf6, 0x04, 0xc6, 0xe6, 0xa0, 0x8b, 0xa2,
	0x44, 0x1e, 0xd3, 0x70, 0x2b, 0xcc, 0x1e, 0x43, 0x94, 0x69, 0xa9, 0x91, 0x03, 0x7d, 0x25, 0x0b,
	0xcc, 0xc6, 0x97, 0x4d, 0x2e, 0x35, 0x52, 0xcd, 0xae, 0xdd, 0x78, 0xcd, 0xb0, 0xe7, 0x10, 0xcd,
	0x34, 0x96, 0x8a, 0xef, 0x91, 0xea, 0x93, 0xb5, 0xea, 0x86, 0x16, 0x36, 0x98, 0xfe, 0x0c, 0xe0,
	0xc1, 0x46, 0x60, 0x4b, 0x73, 0xa3, 0x6c, 0x25, 0x1b, 0x75, 0x53, 0xeb, 0x95, 0xec, 0x1e, 0x63,
	0x26, 0xff, 0xd0, 0xc9, 0x4a, 0x17, 0xfa, 0x3b, 0xe9, 0x1e, 0x89, 0x15, 0xf6, 0xf4, 0x18, 0xfe,
	0x4f, 0x8f, 0x68, 0x4b, 0x8f, 0xf4, 0x1e, 0xf8, 0xbc, 0x5e, 0x14, 0x4a, 0x17, 0xd7, 0x67, 0xcb,
	0xfa, 0xdb, 0x45, 0x2b, 0xaf, 0x6f, 0x05, 0x7e, 0xed, 0x50, 0x51, 0xb5, 0xf9, 0x47, 0x1a, 0x6c,
	0x4f, 0xeb, 0x1c, 0x69, 0xd4, 0x58, 0xf8, 0x14, 0x4b, 0x61, 0x6f, 0xe9, 0xaa, 0x29, 0xc5, 0xfe,
	0x18, 0x1b, 0x9c, 0x99, 0xad, 0xa8, 0xee, 0xb0, 0xd5, 0x34, 0xf5, 0x58, 0x38, 0x94, 0x5e, 0xc2,
	0xa1, 0xb5, 0xb0, 0x4b, 0xfe, 0x57, 0xeb, 0xde, 0x67, 0xa6, 0x35, 0x13, 0x3e, 0xe5, 0x1d, 0x3b,
	0xd8, 0x38, 0xf6, 0xc7, 0x00, 0xf6, 0xb3, 0x95, 0x9b, 0xe8, 0x50, 0xf3, 0xdf, 0xae, 0x36, 0x78,
	0x24, 0xe8, 0xdd, 0x58, 0xfc, 0x9d, 0xbd, 0x08, 0xdc, 0xd4, 0x3d, 0x34, 0x4b, 0xcd, 0xfd, 0xa5,
	0x42, 0xbb, 0x94, 0xcf, 0xf5, 0x16, 0xed, 0xa5, 0x19, 0xae, 0x2d, 0xda, 0x4b, 0xb3, 0xce, 0xa0,
	0x2b, 0x23, 0xda, 0xc8, 0xa0, 0x9b, 0xe3, 0x29, 0xc4, 0x06, 0x5a, 0xcb, 0x59, 0x93, 0xaf, 0x89,
	0xbf, 0x6c, 0x37, 0x22, 0x01, 0x7c, 0xdb, 0x25, 0x10, 0x99, 0x0f, 0xa6, 0xf8, 0x98, 0x6c, 0x07,
	0xd6, 0x76, 0x86, 0x12, 0x36, 0x90, 0x7e, 0x86, 0x78, 0xc5, 0x99, 0x75, 0xb3, 0xee, 0xea, 0x0b,
	0x5e, 0x6b, 0xf7, 0x1d, 0x7b, 0x68, 0x1a, 0x9d, 0xb6, 0xd8, 0x37, 0xb2, 0x5a, 0x78, 0x0c, 0x5d,
	0x5e, 0x58, 0xca, 0xf6, 0xd6, 0x09, 0xe1, 0xd0, 0x49, 0xfc, 0x69, 0x74, 0xf4, 0x86, 0xae, 0xce,
	0xab, 0x1d, 0x7a, 0xbc, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xe8, 0x1d, 0x06, 0x5b, 0x05,
	0x00, 0x00,
}
