// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: delivery.proto

package pb

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

type DeliveryCostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginCityId      string  `protobuf:"bytes,1,opt,name=originCityId,proto3" json:"originCityId,omitempty"`
	DestinationCityId string  `protobuf:"bytes,2,opt,name=destinationCityId,proto3" json:"destinationCityId,omitempty"`
	CartIds           []int32 `protobuf:"varint,3,rep,packed,name=cartIds,proto3" json:"cartIds,omitempty"`
	Courier           string  `protobuf:"bytes,4,opt,name=courier,proto3" json:"courier,omitempty"`
}

func (x *DeliveryCostRequest) Reset() {
	*x = DeliveryCostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryCostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryCostRequest) ProtoMessage() {}

func (x *DeliveryCostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryCostRequest.ProtoReflect.Descriptor instead.
func (*DeliveryCostRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryCostRequest) GetOriginCityId() string {
	if x != nil {
		return x.OriginCityId
	}
	return ""
}

func (x *DeliveryCostRequest) GetDestinationCityId() string {
	if x != nil {
		return x.DestinationCityId
	}
	return ""
}

func (x *DeliveryCostRequest) GetCartIds() []int32 {
	if x != nil {
		return x.CartIds
	}
	return nil
}

func (x *DeliveryCostRequest) GetCourier() string {
	if x != nil {
		return x.Courier
	}
	return ""
}

type DeliveryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId       string `protobuf:"bytes,1,opt,name=cityId,proto3" json:"cityId,omitempty"`
	CityName     string `protobuf:"bytes,2,opt,name=cityName,proto3" json:"cityName,omitempty"`
	ProvinceId   string `protobuf:"bytes,3,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
	ProvinceName string `protobuf:"bytes,4,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	Type         string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	PostalCode   string `protobuf:"bytes,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
}

func (x *DeliveryItem) Reset() {
	*x = DeliveryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryItem) ProtoMessage() {}

func (x *DeliveryItem) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryItem.ProtoReflect.Descriptor instead.
func (*DeliveryItem) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryItem) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *DeliveryItem) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *DeliveryItem) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *DeliveryItem) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *DeliveryItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeliveryItem) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

type ServiceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Cost        int32  `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	Etd         string `protobuf:"bytes,4,opt,name=etd,proto3" json:"etd,omitempty"`
}

func (x *ServiceItem) Reset() {
	*x = ServiceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceItem) ProtoMessage() {}

func (x *ServiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceItem.ProtoReflect.Descriptor instead.
func (*ServiceItem) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceItem) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceItem) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *ServiceItem) GetEtd() string {
	if x != nil {
		return x.Etd
	}
	return ""
}

type DeliveryCostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      *DeliveryItem  `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination *DeliveryItem  `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Service     []*ServiceItem `protobuf:"bytes,3,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *DeliveryCostResponse) Reset() {
	*x = DeliveryCostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryCostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryCostResponse) ProtoMessage() {}

func (x *DeliveryCostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryCostResponse.ProtoReflect.Descriptor instead.
func (*DeliveryCostResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryCostResponse) GetOrigin() *DeliveryItem {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DeliveryCostResponse) GetDestination() *DeliveryItem {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *DeliveryCostResponse) GetService() []*ServiceItem {
	if x != nil {
		return x.Service
	}
	return nil
}

type Province struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId   string `protobuf:"bytes,1,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
	ProvinceName string `protobuf:"bytes,2,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
}

func (x *Province) Reset() {
	*x = Province{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Province) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Province) ProtoMessage() {}

func (x *Province) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Province.ProtoReflect.Descriptor instead.
func (*Province) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{4}
}

func (x *Province) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *Province) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId       string `protobuf:"bytes,1,opt,name=cityId,proto3" json:"cityId,omitempty"`
	CityName     string `protobuf:"bytes,2,opt,name=cityName,proto3" json:"cityName,omitempty"`
	ProvinceId   string `protobuf:"bytes,3,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
	ProvinceName string `protobuf:"bytes,4,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	Type         string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	PostalCode   string `protobuf:"bytes,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{5}
}

func (x *City) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *City) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *City) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *City) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *City) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *City) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

type GetProvinceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provinces []*Province `protobuf:"bytes,1,rep,name=provinces,proto3" json:"provinces,omitempty"`
}

func (x *GetProvinceResponse) Reset() {
	*x = GetProvinceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvinceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvinceResponse) ProtoMessage() {}

func (x *GetProvinceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvinceResponse.ProtoReflect.Descriptor instead.
func (*GetProvinceResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{6}
}

func (x *GetProvinceResponse) GetProvinces() []*Province {
	if x != nil {
		return x.Provinces
	}
	return nil
}

type GetCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId string `protobuf:"bytes,1,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
}

func (x *GetCityRequest) Reset() {
	*x = GetCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityRequest) ProtoMessage() {}

func (x *GetCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityRequest.ProtoReflect.Descriptor instead.
func (*GetCityRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{7}
}

func (x *GetCityRequest) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

type GetCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cities []*City `protobuf:"bytes,1,rep,name=cities,proto3" json:"cities,omitempty"`
}

func (x *GetCityResponse) Reset() {
	*x = GetCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityResponse) ProtoMessage() {}

func (x *GetCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityResponse.ProtoReflect.Descriptor instead.
func (*GetCityResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{8}
}

func (x *GetCityResponse) GetCities() []*City {
	if x != nil {
		return x.Cities
	}
	return nil
}

var File_delivery_proto protoreflect.FileDescriptor

var file_delivery_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x77, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x74, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x74,
	0x64, 0x22, 0x96, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x32, 0xb9, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delivery_proto_rawDescOnce sync.Once
	file_delivery_proto_rawDescData = file_delivery_proto_rawDesc
)

func file_delivery_proto_rawDescGZIP() []byte {
	file_delivery_proto_rawDescOnce.Do(func() {
		file_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_proto_rawDescData)
	})
	return file_delivery_proto_rawDescData
}

var file_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_delivery_proto_goTypes = []any{
	(*DeliveryCostRequest)(nil),  // 0: DeliveryCostRequest
	(*DeliveryItem)(nil),         // 1: DeliveryItem
	(*ServiceItem)(nil),          // 2: ServiceItem
	(*DeliveryCostResponse)(nil), // 3: DeliveryCostResponse
	(*Province)(nil),             // 4: Province
	(*City)(nil),                 // 5: City
	(*GetProvinceResponse)(nil),  // 6: GetProvinceResponse
	(*GetCityRequest)(nil),       // 7: GetCityRequest
	(*GetCityResponse)(nil),      // 8: GetCityResponse
	(*emptypb.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_delivery_proto_depIdxs = []int32{
	1, // 0: DeliveryCostResponse.origin:type_name -> DeliveryItem
	1, // 1: DeliveryCostResponse.destination:type_name -> DeliveryItem
	2, // 2: DeliveryCostResponse.service:type_name -> ServiceItem
	4, // 3: GetProvinceResponse.provinces:type_name -> Province
	5, // 4: GetCityResponse.cities:type_name -> City
	0, // 5: DeliveryService.DeliveryCost:input_type -> DeliveryCostRequest
	9, // 6: DeliveryService.GetProvince:input_type -> google.protobuf.Empty
	7, // 7: DeliveryService.GetCity:input_type -> GetCityRequest
	3, // 8: DeliveryService.DeliveryCost:output_type -> DeliveryCostResponse
	6, // 9: DeliveryService.GetProvince:output_type -> GetProvinceResponse
	8, // 10: DeliveryService.GetCity:output_type -> GetCityResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_delivery_proto_init() }
func file_delivery_proto_init() {
	if File_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeliveryCostRequest); i {
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
		file_delivery_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeliveryItem); i {
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
		file_delivery_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceItem); i {
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
		file_delivery_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeliveryCostResponse); i {
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
		file_delivery_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Province); i {
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
		file_delivery_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*City); i {
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
		file_delivery_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetProvinceResponse); i {
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
		file_delivery_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityRequest); i {
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
		file_delivery_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityResponse); i {
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
			RawDescriptor: file_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_proto_goTypes,
		DependencyIndexes: file_delivery_proto_depIdxs,
		MessageInfos:      file_delivery_proto_msgTypes,
	}.Build()
	File_delivery_proto = out.File
	file_delivery_proto_rawDesc = nil
	file_delivery_proto_goTypes = nil
	file_delivery_proto_depIdxs = nil
}
