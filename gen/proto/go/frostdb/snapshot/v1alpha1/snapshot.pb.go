// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: frostdb/snapshot/v1alpha1/snapshot.proto

package snapshotv1alpha1

import (
	v1alpha1 "github.com/polarsignals/frostdb/gen/proto/go/frostdb/table/v1alpha1"
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

type Part_Encoding int32

const (
	Part_ENCODING_UNKNOWN Part_Encoding = 0
	Part_ENCODING_PARQUET Part_Encoding = 1
	Part_ENCODING_ARROW   Part_Encoding = 2
)

// Enum value maps for Part_Encoding.
var (
	Part_Encoding_name = map[int32]string{
		0: "ENCODING_UNKNOWN",
		1: "ENCODING_PARQUET",
		2: "ENCODING_ARROW",
	}
	Part_Encoding_value = map[string]int32{
		"ENCODING_UNKNOWN": 0,
		"ENCODING_PARQUET": 1,
		"ENCODING_ARROW":   2,
	}
)

func (x Part_Encoding) Enum() *Part_Encoding {
	p := new(Part_Encoding)
	*p = x
	return p
}

func (x Part_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Part_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_enumTypes[0].Descriptor()
}

func (Part_Encoding) Type() protoreflect.EnumType {
	return &file_frostdb_snapshot_v1alpha1_snapshot_proto_enumTypes[0]
}

func (x Part_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Part_Encoding.Descriptor instead.
func (Part_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{3, 0}
}

// FooterData is a message stored in the footer of a snapshot file that encodes
// data about the rest of the file.
type FooterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableMetadata []*Table `protobuf:"bytes,1,rep,name=table_metadata,json=tableMetadata,proto3" json:"table_metadata,omitempty"`
}

func (x *FooterData) Reset() {
	*x = FooterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooterData) ProtoMessage() {}

func (x *FooterData) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooterData.ProtoReflect.Descriptor instead.
func (*FooterData) Descriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *FooterData) GetTableMetadata() []*Table {
	if x != nil {
		return x.TableMetadata
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config          *v1alpha1.TableConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	ActiveBlock     *Table_TableBlock     `protobuf:"bytes,3,opt,name=active_block,json=activeBlock,proto3" json:"active_block,omitempty"`
	GranuleMetadata []*Granule            `protobuf:"bytes,4,rep,name=granule_metadata,json=granuleMetadata,proto3" json:"granule_metadata,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetConfig() *v1alpha1.TableConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Table) GetActiveBlock() *Table_TableBlock {
	if x != nil {
		return x.ActiveBlock
	}
	return nil
}

func (x *Table) GetGranuleMetadata() []*Granule {
	if x != nil {
		return x.GranuleMetadata
	}
	return nil
}

type Granule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartMetadata []*Part `protobuf:"bytes,1,rep,name=part_metadata,json=partMetadata,proto3" json:"part_metadata,omitempty"`
}

func (x *Granule) Reset() {
	*x = Granule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Granule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Granule) ProtoMessage() {}

func (x *Granule) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Granule.ProtoReflect.Descriptor instead.
func (*Granule) Descriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *Granule) GetPartMetadata() []*Part {
	if x != nil {
		return x.PartMetadata
	}
	return nil
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartOffset     int64         `protobuf:"varint,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	EndOffset       int64         `protobuf:"varint,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	Tx              uint64        `protobuf:"varint,3,opt,name=tx,proto3" json:"tx,omitempty"`
	CompactionLevel uint64        `protobuf:"varint,4,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	Encoding        Part_Encoding `protobuf:"varint,5,opt,name=encoding,proto3,enum=frostdb.snapshot.v1alpha1.Part_Encoding" json:"encoding,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *Part) GetStartOffset() int64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *Part) GetEndOffset() int64 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

func (x *Part) GetTx() uint64 {
	if x != nil {
		return x.Tx
	}
	return 0
}

func (x *Part) GetCompactionLevel() uint64 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *Part) GetEncoding() Part_Encoding {
	if x != nil {
		return x.Encoding
	}
	return Part_ENCODING_UNKNOWN
}

type Table_TableBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ulid   []byte `protobuf:"bytes,1,opt,name=ulid,proto3" json:"ulid,omitempty"`
	Size   int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	MinTx  uint64 `protobuf:"varint,3,opt,name=min_tx,json=minTx,proto3" json:"min_tx,omitempty"`
	PrevTx uint64 `protobuf:"varint,4,opt,name=prev_tx,json=prevTx,proto3" json:"prev_tx,omitempty"`
}

func (x *Table_TableBlock) Reset() {
	*x = Table_TableBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table_TableBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table_TableBlock) ProtoMessage() {}

func (x *Table_TableBlock) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table_TableBlock.ProtoReflect.Descriptor instead.
func (*Table_TableBlock) Descriptor() ([]byte, []int) {
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Table_TableBlock) GetUlid() []byte {
	if x != nil {
		return x.Ulid
	}
	return nil
}

func (x *Table_TableBlock) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Table_TableBlock) GetMinTx() uint64 {
	if x != nil {
		return x.MinTx
	}
	return 0
}

func (x *Table_TableBlock) GetPrevTx() uint64 {
	if x != nil {
		return x.PrevTx
	}
	return 0
}

var File_frostdb_snapshot_v1alpha1_snapshot_proto protoreflect.FileDescriptor

var file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDesc = []byte{
	0x0a, 0x28, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x72, 0x6f, 0x73,
	0x74, 0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x23, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0a, 0x46, 0x6f,
	0x6f, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x0d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xdd, 0x02, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x4e, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74,
	0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x4d, 0x0a, 0x10, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66,
	0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65,
	0x52, 0x0f, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x64, 0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75,
	0x6c, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x5f, 0x74,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x76, 0x54, 0x78, 0x22, 0x4f, 0x0a, 0x07, 0x47, 0x72, 0x61, 0x6e, 0x75,
	0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72, 0x6f, 0x73,
	0x74, 0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x02, 0x0a, 0x04, 0x50, 0x61, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x74, 0x78, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x44,
	0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x51, 0x55, 0x45, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52, 0x52, 0x4f, 0x57, 0x10, 0x02,
	0x42, 0x8d, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62,
	0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x66, 0x72, 0x6f,
	0x73, 0x74, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x46,
	0x53, 0x58, 0xaa, 0x02, 0x19, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x19, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x5c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x46, 0x72, 0x6f,
	0x73, 0x74, 0x64, 0x62, 0x5c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1b, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x3a, 0x3a, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescOnce sync.Once
	file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescData = file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDesc
)

func file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescGZIP() []byte {
	file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescOnce.Do(func() {
		file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescData)
	})
	return file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDescData
}

var file_frostdb_snapshot_v1alpha1_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_frostdb_snapshot_v1alpha1_snapshot_proto_goTypes = []interface{}{
	(Part_Encoding)(0),           // 0: frostdb.snapshot.v1alpha1.Part.Encoding
	(*FooterData)(nil),           // 1: frostdb.snapshot.v1alpha1.FooterData
	(*Table)(nil),                // 2: frostdb.snapshot.v1alpha1.Table
	(*Granule)(nil),              // 3: frostdb.snapshot.v1alpha1.Granule
	(*Part)(nil),                 // 4: frostdb.snapshot.v1alpha1.Part
	(*Table_TableBlock)(nil),     // 5: frostdb.snapshot.v1alpha1.Table.TableBlock
	(*v1alpha1.TableConfig)(nil), // 6: frostdb.table.v1alpha1.TableConfig
}
var file_frostdb_snapshot_v1alpha1_snapshot_proto_depIdxs = []int32{
	2, // 0: frostdb.snapshot.v1alpha1.FooterData.table_metadata:type_name -> frostdb.snapshot.v1alpha1.Table
	6, // 1: frostdb.snapshot.v1alpha1.Table.config:type_name -> frostdb.table.v1alpha1.TableConfig
	5, // 2: frostdb.snapshot.v1alpha1.Table.active_block:type_name -> frostdb.snapshot.v1alpha1.Table.TableBlock
	3, // 3: frostdb.snapshot.v1alpha1.Table.granule_metadata:type_name -> frostdb.snapshot.v1alpha1.Granule
	4, // 4: frostdb.snapshot.v1alpha1.Granule.part_metadata:type_name -> frostdb.snapshot.v1alpha1.Part
	0, // 5: frostdb.snapshot.v1alpha1.Part.encoding:type_name -> frostdb.snapshot.v1alpha1.Part.Encoding
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_frostdb_snapshot_v1alpha1_snapshot_proto_init() }
func file_frostdb_snapshot_v1alpha1_snapshot_proto_init() {
	if File_frostdb_snapshot_v1alpha1_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooterData); i {
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
		file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Granule); i {
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
		file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table_TableBlock); i {
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
			RawDescriptor: file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frostdb_snapshot_v1alpha1_snapshot_proto_goTypes,
		DependencyIndexes: file_frostdb_snapshot_v1alpha1_snapshot_proto_depIdxs,
		EnumInfos:         file_frostdb_snapshot_v1alpha1_snapshot_proto_enumTypes,
		MessageInfos:      file_frostdb_snapshot_v1alpha1_snapshot_proto_msgTypes,
	}.Build()
	File_frostdb_snapshot_v1alpha1_snapshot_proto = out.File
	file_frostdb_snapshot_v1alpha1_snapshot_proto_rawDesc = nil
	file_frostdb_snapshot_v1alpha1_snapshot_proto_goTypes = nil
	file_frostdb_snapshot_v1alpha1_snapshot_proto_depIdxs = nil
}
