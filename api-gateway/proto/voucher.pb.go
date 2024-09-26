// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: proto/voucher.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherId  string  `protobuf:"bytes,1,opt,name=voucherId,proto3" json:"voucherId,omitempty"`
	Discount   float64 `protobuf:"fixed64,2,opt,name=discount,proto3" json:"discount,omitempty"`   // Use double instead of dedcimal(10,2)
	ValidUntil string  `protobuf:"bytes,3,opt,name=validUntil,proto3" json:"validUntil,omitempty"` // Use string for date or handle with Timestamp in your code
}

func (x *AddVoucherRequest) Reset() {
	*x = AddVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVoucherRequest) ProtoMessage() {}

func (x *AddVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVoucherRequest.ProtoReflect.Descriptor instead.
func (*AddVoucherRequest) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{0}
}

func (x *AddVoucherRequest) GetVoucherId() string {
	if x != nil {
		return x.VoucherId
	}
	return ""
}

func (x *AddVoucherRequest) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *AddVoucherRequest) GetValidUntil() string {
	if x != nil {
		return x.ValidUntil
	}
	return ""
}

type AddVoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddVoucherResponse) Reset() {
	*x = AddVoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVoucherResponse) ProtoMessage() {}

func (x *AddVoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVoucherResponse.ProtoReflect.Descriptor instead.
func (*AddVoucherResponse) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{1}
}

func (x *AddVoucherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherId string `protobuf:"bytes,1,opt,name=voucherId,proto3" json:"voucherId,omitempty"`
}

func (x *GetVoucherRequest) Reset() {
	*x = GetVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoucherRequest) ProtoMessage() {}

func (x *GetVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoucherRequest.ProtoReflect.Descriptor instead.
func (*GetVoucherRequest) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{2}
}

func (x *GetVoucherRequest) GetVoucherId() string {
	if x != nil {
		return x.VoucherId
	}
	return ""
}

type GetVoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherId  string  `protobuf:"bytes,1,opt,name=voucherId,proto3" json:"voucherId,omitempty"`
	Discount   float64 `protobuf:"fixed64,2,opt,name=discount,proto3" json:"discount,omitempty"`
	ValidUntil string  `protobuf:"bytes,3,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	Used       bool    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *GetVoucherResponse) Reset() {
	*x = GetVoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoucherResponse) ProtoMessage() {}

func (x *GetVoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoucherResponse.ProtoReflect.Descriptor instead.
func (*GetVoucherResponse) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{3}
}

func (x *GetVoucherResponse) GetVoucherId() string {
	if x != nil {
		return x.VoucherId
	}
	return ""
}

func (x *GetVoucherResponse) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *GetVoucherResponse) GetValidUntil() string {
	if x != nil {
		return x.ValidUntil
	}
	return ""
}

func (x *GetVoucherResponse) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

type GetVoucherListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vouchers []*Voucher `protobuf:"bytes,1,rep,name=vouchers,proto3" json:"vouchers,omitempty"` // List of vouchers
}

func (x *GetVoucherListResponse) Reset() {
	*x = GetVoucherListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoucherListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoucherListResponse) ProtoMessage() {}

func (x *GetVoucherListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoucherListResponse.ProtoReflect.Descriptor instead.
func (*GetVoucherListResponse) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{4}
}

func (x *GetVoucherListResponse) GetVouchers() []*Voucher {
	if x != nil {
		return x.Vouchers
	}
	return nil
}

type Voucher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherId  string  `protobuf:"bytes,1,opt,name=voucherId,proto3" json:"voucherId,omitempty"`
	Discount   float64 `protobuf:"fixed64,2,opt,name=discount,proto3" json:"discount,omitempty"`
	ValidUntil string  `protobuf:"bytes,3,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	Used       bool    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *Voucher) Reset() {
	*x = Voucher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_voucher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voucher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voucher) ProtoMessage() {}

func (x *Voucher) ProtoReflect() protoreflect.Message {
	mi := &file_proto_voucher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voucher.ProtoReflect.Descriptor instead.
func (*Voucher) Descriptor() ([]byte, []int) {
	return file_proto_voucher_proto_rawDescGZIP(), []int{5}
}

func (x *Voucher) GetVoucherId() string {
	if x != nil {
		return x.VoucherId
	}
	return ""
}

func (x *Voucher) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Voucher) GetValidUntil() string {
	if x != nil {
		return x.ValidUntil
	}
	return ""
}

func (x *Voucher) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

var File_proto_voucher_proto protoreflect.FileDescriptor

var file_proto_voucher_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69,
	0x6c, 0x22, 0x2e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e,
	0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x08, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x77, 0x0a, 0x07, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x64, 0x32, 0xc1, 0x01, 0x0a, 0x0e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_voucher_proto_rawDescOnce sync.Once
	file_proto_voucher_proto_rawDescData = file_proto_voucher_proto_rawDesc
)

func file_proto_voucher_proto_rawDescGZIP() []byte {
	file_proto_voucher_proto_rawDescOnce.Do(func() {
		file_proto_voucher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_voucher_proto_rawDescData)
	})
	return file_proto_voucher_proto_rawDescData
}

var file_proto_voucher_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_voucher_proto_goTypes = []any{
	(*AddVoucherRequest)(nil),      // 0: AddVoucherRequest
	(*AddVoucherResponse)(nil),     // 1: AddVoucherResponse
	(*GetVoucherRequest)(nil),      // 2: GetVoucherRequest
	(*GetVoucherResponse)(nil),     // 3: GetVoucherResponse
	(*GetVoucherListResponse)(nil), // 4: GetVoucherListResponse
	(*Voucher)(nil),                // 5: Voucher
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_proto_voucher_proto_depIdxs = []int32{
	5, // 0: GetVoucherListResponse.vouchers:type_name -> Voucher
	0, // 1: VoucherService.AddVoucher:input_type -> AddVoucherRequest
	2, // 2: VoucherService.GetVoucher:input_type -> GetVoucherRequest
	6, // 3: VoucherService.GetVoucherList:input_type -> google.protobuf.Empty
	1, // 4: VoucherService.AddVoucher:output_type -> AddVoucherResponse
	3, // 5: VoucherService.GetVoucher:output_type -> GetVoucherResponse
	4, // 6: VoucherService.GetVoucherList:output_type -> GetVoucherListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_voucher_proto_init() }
func file_proto_voucher_proto_init() {
	if File_proto_voucher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_voucher_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddVoucherRequest); i {
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
		file_proto_voucher_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddVoucherResponse); i {
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
		file_proto_voucher_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetVoucherRequest); i {
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
		file_proto_voucher_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetVoucherResponse); i {
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
		file_proto_voucher_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetVoucherListResponse); i {
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
		file_proto_voucher_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Voucher); i {
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
			RawDescriptor: file_proto_voucher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_voucher_proto_goTypes,
		DependencyIndexes: file_proto_voucher_proto_depIdxs,
		MessageInfos:      file_proto_voucher_proto_msgTypes,
	}.Build()
	File_proto_voucher_proto = out.File
	file_proto_voucher_proto_rawDesc = nil
	file_proto_voucher_proto_goTypes = nil
	file_proto_voucher_proto_depIdxs = nil
}
