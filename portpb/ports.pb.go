// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: portpb/ports.proto

package portpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string       `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string       `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string     `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string     `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates *Coordinates `protobuf:"bytes,6,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string       `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string       `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      *Unlocs      `protobuf:"bytes,9,opt,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string       `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portpb_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_portpb_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_portpb_ports_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() *Coordinates {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() *Unlocs {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat  float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float32 `protobuf:"fixed32,2,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portpb_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_portpb_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_portpb_ports_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinates) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

type Unlocs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unloc []string `protobuf:"bytes,1,rep,name=unloc,proto3" json:"unloc,omitempty"`
}

func (x *Unlocs) Reset() {
	*x = Unlocs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portpb_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unlocs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unlocs) ProtoMessage() {}

func (x *Unlocs) ProtoReflect() protoreflect.Message {
	mi := &file_portpb_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unlocs.ProtoReflect.Descriptor instead.
func (*Unlocs) Descriptor() ([]byte, []int) {
	return file_portpb_ports_proto_rawDescGZIP(), []int{2}
}

func (x *Unlocs) GetUnloc() []string {
	if x != nil {
		return x.Unloc
	}
	return nil
}

type PortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *PortRequest) Reset() {
	*x = PortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portpb_ports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRequest) ProtoMessage() {}

func (x *PortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portpb_ports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRequest.ProtoReflect.Descriptor instead.
func (*PortRequest) Descriptor() ([]byte, []int) {
	return file_portpb_ports_proto_rawDescGZIP(), []int{3}
}

func (x *PortRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type PortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PortResponse) Reset() {
	*x = PortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portpb_ports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortResponse) ProtoMessage() {}

func (x *PortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portpb_ports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortResponse.ProtoReflect.Descriptor instead.
func (*PortResponse) Descriptor() ([]byte, []int) {
	return file_portpb_ports_proto_rawDescGZIP(), []int{4}
}

func (x *PortResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_portpb_ports_proto protoreflect.FileDescriptor

var file_portpb_ports_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x73,
	0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x0b,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x22, 0x1e, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x22, 0x2d, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x26, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x47, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x64, 0x75, 0x61, 0x72, 0x64, 0x6f, 0x62, 0x63, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x6f, 0x2f,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portpb_ports_proto_rawDescOnce sync.Once
	file_portpb_ports_proto_rawDescData = file_portpb_ports_proto_rawDesc
)

func file_portpb_ports_proto_rawDescGZIP() []byte {
	file_portpb_ports_proto_rawDescOnce.Do(func() {
		file_portpb_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_portpb_ports_proto_rawDescData)
	})
	return file_portpb_ports_proto_rawDescData
}

var file_portpb_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_portpb_ports_proto_goTypes = []interface{}{
	(*Port)(nil),         // 0: port.Port
	(*Coordinates)(nil),  // 1: port.Coordinates
	(*Unlocs)(nil),       // 2: port.Unlocs
	(*PortRequest)(nil),  // 3: port.PortRequest
	(*PortResponse)(nil), // 4: port.PortResponse
}
var file_portpb_ports_proto_depIdxs = []int32{
	1, // 0: port.Port.coordinates:type_name -> port.Coordinates
	2, // 1: port.Port.unlocs:type_name -> port.Unlocs
	0, // 2: port.PortRequest.port:type_name -> port.Port
	3, // 3: port.PortService.PortsUpdate:input_type -> port.PortRequest
	4, // 4: port.PortService.PortsUpdate:output_type -> port.PortResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_portpb_ports_proto_init() }
func file_portpb_ports_proto_init() {
	if File_portpb_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portpb_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_portpb_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
		file_portpb_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unlocs); i {
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
		file_portpb_ports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRequest); i {
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
		file_portpb_ports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortResponse); i {
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
			RawDescriptor: file_portpb_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portpb_ports_proto_goTypes,
		DependencyIndexes: file_portpb_ports_proto_depIdxs,
		MessageInfos:      file_portpb_ports_proto_msgTypes,
	}.Build()
	File_portpb_ports_proto = out.File
	file_portpb_ports_proto_rawDesc = nil
	file_portpb_ports_proto_goTypes = nil
	file_portpb_ports_proto_depIdxs = nil
}
