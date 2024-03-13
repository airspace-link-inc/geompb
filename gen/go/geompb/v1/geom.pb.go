// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: geom/v1/geom.proto

package geompb

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

type Geometry_Type int32

const (
	Geometry_TYPE_UNSPECIFIED      Geometry_Type = 0
	Geometry_TYPE_POINT            Geometry_Type = 1  // coordinates[0], 2 coordinates: longitude, latitude
	Geometry_TYPE_POINTZ           Geometry_Type = 2  // coordinates[0], 3 coordinates: longitude, latitude, altitude
	Geometry_TYPE_MULTIPOINT       Geometry_Type = 3  // coordinates is empty, use geometries with POINT types
	Geometry_TYPE_MULTIPOINTZ      Geometry_Type = 4  // coordinates is empty, use geometries with POINTZ types
	Geometry_TYPE_POLYGON          Geometry_Type = 5  // coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW
	Geometry_TYPE_POLYGONZ         Geometry_Type = 6  // coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW
	Geometry_TYPE_MULTIPOLYGON     Geometry_Type = 7  // coordinates is empty, use geometries with POLYGON types
	Geometry_TYPE_MULTIPOLYGONZ    Geometry_Type = 8  // coordinates is empty, use geometries with POLYGONZ types
	Geometry_TYPE_LINESTRING       Geometry_Type = 9  // coordinates[0], list of longitude, latitude
	Geometry_TYPE_LINESTRINGZ      Geometry_Type = 10 // coordinates[0], list of longitude, latitude, altitude
	Geometry_TYPE_MULTILINESTRING  Geometry_Type = 11 // coordinates is empty, use geometries with LINESTRING types
	Geometry_TYPE_MULTILINESTRINGZ Geometry_Type = 12 // coordinates is empty, use geometries with LINESTRINGZ types
)

// Enum value maps for Geometry_Type.
var (
	Geometry_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "TYPE_POINT",
		2:  "TYPE_POINTZ",
		3:  "TYPE_MULTIPOINT",
		4:  "TYPE_MULTIPOINTZ",
		5:  "TYPE_POLYGON",
		6:  "TYPE_POLYGONZ",
		7:  "TYPE_MULTIPOLYGON",
		8:  "TYPE_MULTIPOLYGONZ",
		9:  "TYPE_LINESTRING",
		10: "TYPE_LINESTRINGZ",
		11: "TYPE_MULTILINESTRING",
		12: "TYPE_MULTILINESTRINGZ",
	}
	Geometry_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":      0,
		"TYPE_POINT":            1,
		"TYPE_POINTZ":           2,
		"TYPE_MULTIPOINT":       3,
		"TYPE_MULTIPOINTZ":      4,
		"TYPE_POLYGON":          5,
		"TYPE_POLYGONZ":         6,
		"TYPE_MULTIPOLYGON":     7,
		"TYPE_MULTIPOLYGONZ":    8,
		"TYPE_LINESTRING":       9,
		"TYPE_LINESTRINGZ":      10,
		"TYPE_MULTILINESTRING":  11,
		"TYPE_MULTILINESTRINGZ": 12,
	}
)

func (x Geometry_Type) Enum() *Geometry_Type {
	p := new(Geometry_Type)
	*p = x
	return p
}

func (x Geometry_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Geometry_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_geom_v1_geom_proto_enumTypes[0].Descriptor()
}

func (Geometry_Type) Type() protoreflect.EnumType {
	return &file_geom_v1_geom_proto_enumTypes[0]
}

func (x Geometry_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Geometry_Type.Descriptor instead.
func (Geometry_Type) EnumDescriptor() ([]byte, []int) {
	return file_geom_v1_geom_proto_rawDescGZIP(), []int{1, 0}
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequences []float64 `protobuf:"fixed64,1,rep,packed,name=sequences,proto3" json:"sequences,omitempty"` // lng, lat or lng lat altitude
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geom_v1_geom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_geom_v1_geom_proto_msgTypes[0]
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
	return file_geom_v1_geom_proto_rawDescGZIP(), []int{0}
}

func (x *Coordinates) GetSequences() []float64 {
	if x != nil {
		return x.Sequences
	}
	return nil
}

type Geometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Geometry_Type  `protobuf:"varint,1,opt,name=type,proto3,enum=geompb.v1.Geometry_Type" json:"type,omitempty"`
	Geometries  []*Geometry    `protobuf:"bytes,2,rep,name=geometries,proto3" json:"geometries,omitempty"`
	Coordinates []*Coordinates `protobuf:"bytes,3,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Geometry) Reset() {
	*x = Geometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geom_v1_geom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geometry) ProtoMessage() {}

func (x *Geometry) ProtoReflect() protoreflect.Message {
	mi := &file_geom_v1_geom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geometry.ProtoReflect.Descriptor instead.
func (*Geometry) Descriptor() ([]byte, []int) {
	return file_geom_v1_geom_proto_rawDescGZIP(), []int{1}
}

func (x *Geometry) GetType() Geometry_Type {
	if x != nil {
		return x.Type
	}
	return Geometry_TYPE_UNSPECIFIED
}

func (x *Geometry) GetGeometries() []*Geometry {
	if x != nil {
		return x.Geometries
	}
	return nil
}

func (x *Geometry) GetCoordinates() []*Coordinates {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_geom_v1_geom_proto protoreflect.FileDescriptor

var file_geom_v1_geom_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x65, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x65, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x22,
	0x2b, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xc6, 0x03, 0x0a,
	0x08, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65,
	0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x9c, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x5a, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5a, 0x10,
	0x04, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x59, 0x47, 0x4f,
	0x4e, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x59,
	0x47, 0x4f, 0x4e, 0x5a, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4f, 0x4c, 0x59, 0x47, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x16, 0x0a,
	0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4f, 0x4c, 0x59, 0x47,
	0x4f, 0x4e, 0x5a, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49,
	0x4e, 0x45, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5a, 0x10, 0x0a,
	0x12, 0x18, 0x0a, 0x14, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x4c, 0x49,
	0x4e, 0x45, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x4c, 0x49, 0x4e, 0x45, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x5a, 0x10, 0x0c, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x6e,
	0x6b, 0x2d, 0x69, 0x6e, 0x63, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x65,
	0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geom_v1_geom_proto_rawDescOnce sync.Once
	file_geom_v1_geom_proto_rawDescData = file_geom_v1_geom_proto_rawDesc
)

func file_geom_v1_geom_proto_rawDescGZIP() []byte {
	file_geom_v1_geom_proto_rawDescOnce.Do(func() {
		file_geom_v1_geom_proto_rawDescData = protoimpl.X.CompressGZIP(file_geom_v1_geom_proto_rawDescData)
	})
	return file_geom_v1_geom_proto_rawDescData
}

var file_geom_v1_geom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_geom_v1_geom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_geom_v1_geom_proto_goTypes = []interface{}{
	(Geometry_Type)(0),  // 0: geompb.v1.Geometry.Type
	(*Coordinates)(nil), // 1: geompb.v1.Coordinates
	(*Geometry)(nil),    // 2: geompb.v1.Geometry
}
var file_geom_v1_geom_proto_depIdxs = []int32{
	0, // 0: geompb.v1.Geometry.type:type_name -> geompb.v1.Geometry.Type
	2, // 1: geompb.v1.Geometry.geometries:type_name -> geompb.v1.Geometry
	1, // 2: geompb.v1.Geometry.coordinates:type_name -> geompb.v1.Coordinates
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_geom_v1_geom_proto_init() }
func file_geom_v1_geom_proto_init() {
	if File_geom_v1_geom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geom_v1_geom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_geom_v1_geom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geometry); i {
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
			RawDescriptor: file_geom_v1_geom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geom_v1_geom_proto_goTypes,
		DependencyIndexes: file_geom_v1_geom_proto_depIdxs,
		EnumInfos:         file_geom_v1_geom_proto_enumTypes,
		MessageInfos:      file_geom_v1_geom_proto_msgTypes,
	}.Build()
	File_geom_v1_geom_proto = out.File
	file_geom_v1_geom_proto_rawDesc = nil
	file_geom_v1_geom_proto_goTypes = nil
	file_geom_v1_geom_proto_depIdxs = nil
}