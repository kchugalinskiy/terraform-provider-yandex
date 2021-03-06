// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mongodb/v1/config/mongodb3_6.proto

package mongodb // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor int32

const (
	MongodConfig3_6_Storage_WiredTiger_CollectionConfig_COMPRESSOR_UNSPECIFIED MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor = 0
	// No compression.
	MongodConfig3_6_Storage_WiredTiger_CollectionConfig_NONE MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor = 1
	// The [Snappy](https://docs.mongodb.com/v3.6/reference/glossary/#term-snappy) compression.
	MongodConfig3_6_Storage_WiredTiger_CollectionConfig_SNAPPY MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor = 2
	// The [zlib](https://docs.mongodb.com/v3.6/reference/glossary/#term-zlib) compression.
	MongodConfig3_6_Storage_WiredTiger_CollectionConfig_ZLIB MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor = 3
)

var MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor_name = map[int32]string{
	0: "COMPRESSOR_UNSPECIFIED",
	1: "NONE",
	2: "SNAPPY",
	3: "ZLIB",
}
var MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor_value = map[string]int32{
	"COMPRESSOR_UNSPECIFIED": 0,
	"NONE":   1,
	"SNAPPY": 2,
	"ZLIB":   3,
}

func (x MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor) String() string {
	return proto.EnumName(MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor_name, int32(x))
}
func (MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0, 0, 1, 0}
}

type MongodConfig3_6_OperationProfiling_Mode int32

const (
	MongodConfig3_6_OperationProfiling_MODE_UNSPECIFIED MongodConfig3_6_OperationProfiling_Mode = 0
	// The profiler is off and does not collect any data.
	MongodConfig3_6_OperationProfiling_OFF MongodConfig3_6_OperationProfiling_Mode = 1
	// The profiler collects data for operations that take longer than the value of [slow_op_threshold].
	MongodConfig3_6_OperationProfiling_SLOW_OP MongodConfig3_6_OperationProfiling_Mode = 2
	// The profiler collects data for all operations.
	MongodConfig3_6_OperationProfiling_ALL MongodConfig3_6_OperationProfiling_Mode = 3
)

var MongodConfig3_6_OperationProfiling_Mode_name = map[int32]string{
	0: "MODE_UNSPECIFIED",
	1: "OFF",
	2: "SLOW_OP",
	3: "ALL",
}
var MongodConfig3_6_OperationProfiling_Mode_value = map[string]int32{
	"MODE_UNSPECIFIED": 0,
	"OFF":              1,
	"SLOW_OP":          2,
	"ALL":              3,
}

func (x MongodConfig3_6_OperationProfiling_Mode) String() string {
	return proto.EnumName(MongodConfig3_6_OperationProfiling_Mode_name, int32(x))
}
func (MongodConfig3_6_OperationProfiling_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 1, 0}
}

// Configuration of a mongod daemon. Supported options are a limited subset of all
// options described in [MongoDB documentation](https://docs.mongodb.com/v3.6/reference/configuration-options/).
type MongodConfig3_6 struct {
	// `storage` section of mongod configuration.
	Storage *MongodConfig3_6_Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	// `operationProfiling` section of mongod configuration.
	OperationProfiling *MongodConfig3_6_OperationProfiling `protobuf:"bytes,2,opt,name=operation_profiling,json=operationProfiling,proto3" json:"operation_profiling,omitempty"`
	// `net` section of mongod configuration.
	Net                  *MongodConfig3_6_Network `protobuf:"bytes,3,opt,name=net,proto3" json:"net,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MongodConfig3_6) Reset()         { *m = MongodConfig3_6{} }
func (m *MongodConfig3_6) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6) ProtoMessage()    {}
func (*MongodConfig3_6) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0}
}
func (m *MongodConfig3_6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6.Unmarshal(m, b)
}
func (m *MongodConfig3_6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6.Merge(dst, src)
}
func (m *MongodConfig3_6) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6.Size(m)
}
func (m *MongodConfig3_6) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6 proto.InternalMessageInfo

func (m *MongodConfig3_6) GetStorage() *MongodConfig3_6_Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *MongodConfig3_6) GetOperationProfiling() *MongodConfig3_6_OperationProfiling {
	if m != nil {
		return m.OperationProfiling
	}
	return nil
}

func (m *MongodConfig3_6) GetNet() *MongodConfig3_6_Network {
	if m != nil {
		return m.Net
	}
	return nil
}

type MongodConfig3_6_Storage struct {
	// Configuration of the WiredTiger storage engine.
	WiredTiger *MongodConfig3_6_Storage_WiredTiger `protobuf:"bytes,1,opt,name=wired_tiger,json=wiredTiger,proto3" json:"wired_tiger,omitempty"`
	// Configuration of the MongoDB [journal](https://docs.mongodb.com/v3.6/reference/glossary/#term-journal).
	Journal              *MongodConfig3_6_Storage_Journal `protobuf:"bytes,2,opt,name=journal,proto3" json:"journal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MongodConfig3_6_Storage) Reset()         { *m = MongodConfig3_6_Storage{} }
func (m *MongodConfig3_6_Storage) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6_Storage) ProtoMessage()    {}
func (*MongodConfig3_6_Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0}
}
func (m *MongodConfig3_6_Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Storage.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Storage.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Storage.Merge(dst, src)
}
func (m *MongodConfig3_6_Storage) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Storage.Size(m)
}
func (m *MongodConfig3_6_Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Storage proto.InternalMessageInfo

func (m *MongodConfig3_6_Storage) GetWiredTiger() *MongodConfig3_6_Storage_WiredTiger {
	if m != nil {
		return m.WiredTiger
	}
	return nil
}

func (m *MongodConfig3_6_Storage) GetJournal() *MongodConfig3_6_Storage_Journal {
	if m != nil {
		return m.Journal
	}
	return nil
}

// Configuration of WiredTiger storage engine.
type MongodConfig3_6_Storage_WiredTiger struct {
	// Engine configuration for WiredTiger.
	EngineConfig *MongodConfig3_6_Storage_WiredTiger_EngineConfig `protobuf:"bytes,1,opt,name=engine_config,json=engineConfig,proto3" json:"engine_config,omitempty"`
	// Collection configuration for WiredTiger.
	CollectionConfig     *MongodConfig3_6_Storage_WiredTiger_CollectionConfig `protobuf:"bytes,2,opt,name=collection_config,json=collectionConfig,proto3" json:"collection_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *MongodConfig3_6_Storage_WiredTiger) Reset()         { *m = MongodConfig3_6_Storage_WiredTiger{} }
func (m *MongodConfig3_6_Storage_WiredTiger) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6_Storage_WiredTiger) ProtoMessage()    {}
func (*MongodConfig3_6_Storage_WiredTiger) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0, 0}
}
func (m *MongodConfig3_6_Storage_WiredTiger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Storage_WiredTiger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Storage_WiredTiger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger.Merge(dst, src)
}
func (m *MongodConfig3_6_Storage_WiredTiger) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger.Size(m)
}
func (m *MongodConfig3_6_Storage_WiredTiger) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger proto.InternalMessageInfo

func (m *MongodConfig3_6_Storage_WiredTiger) GetEngineConfig() *MongodConfig3_6_Storage_WiredTiger_EngineConfig {
	if m != nil {
		return m.EngineConfig
	}
	return nil
}

func (m *MongodConfig3_6_Storage_WiredTiger) GetCollectionConfig() *MongodConfig3_6_Storage_WiredTiger_CollectionConfig {
	if m != nil {
		return m.CollectionConfig
	}
	return nil
}

type MongodConfig3_6_Storage_WiredTiger_EngineConfig struct {
	// The maximum size of the internal cache that WiredTiger will use for all data.
	CacheSizeGb          *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=cache_size_gb,json=cacheSizeGb,proto3" json:"cache_size_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) Reset() {
	*m = MongodConfig3_6_Storage_WiredTiger_EngineConfig{}
}
func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) String() string {
	return proto.CompactTextString(m)
}
func (*MongodConfig3_6_Storage_WiredTiger_EngineConfig) ProtoMessage() {}
func (*MongodConfig3_6_Storage_WiredTiger_EngineConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0, 0, 0}
}
func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Storage_WiredTiger_EngineConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig.Merge(dst, src)
}
func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig.Size(m)
}
func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_EngineConfig proto.InternalMessageInfo

func (m *MongodConfig3_6_Storage_WiredTiger_EngineConfig) GetCacheSizeGb() *wrappers.DoubleValue {
	if m != nil {
		return m.CacheSizeGb
	}
	return nil
}

type MongodConfig3_6_Storage_WiredTiger_CollectionConfig struct {
	// Default type of compression to use for collection data.
	BlockCompressor      MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor `protobuf:"varint,1,opt,name=block_compressor,json=blockCompressor,proto3,enum=yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor" json:"block_compressor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                       `json:"-"`
	XXX_unrecognized     []byte                                                         `json:"-"`
	XXX_sizecache        int32                                                          `json:"-"`
}

func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) Reset() {
	*m = MongodConfig3_6_Storage_WiredTiger_CollectionConfig{}
}
func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) String() string {
	return proto.CompactTextString(m)
}
func (*MongodConfig3_6_Storage_WiredTiger_CollectionConfig) ProtoMessage() {}
func (*MongodConfig3_6_Storage_WiredTiger_CollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0, 0, 1}
}
func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig.Merge(dst, src)
}
func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig.Size(m)
}
func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Storage_WiredTiger_CollectionConfig proto.InternalMessageInfo

func (m *MongodConfig3_6_Storage_WiredTiger_CollectionConfig) GetBlockCompressor() MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor {
	if m != nil {
		return m.BlockCompressor
	}
	return MongodConfig3_6_Storage_WiredTiger_CollectionConfig_COMPRESSOR_UNSPECIFIED
}

type MongodConfig3_6_Storage_Journal struct {
	// Whether the journal is enabled or disabled.
	// Possible values:
	// * true (default) — the journal is enabled.
	// * false — the journal is disabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Commit interval between journal operations, in milliseconds.
	// Default: 100.
	CommitInterval       *wrappers.Int64Value `protobuf:"bytes,2,opt,name=commit_interval,json=commitInterval,proto3" json:"commit_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MongodConfig3_6_Storage_Journal) Reset()         { *m = MongodConfig3_6_Storage_Journal{} }
func (m *MongodConfig3_6_Storage_Journal) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6_Storage_Journal) ProtoMessage()    {}
func (*MongodConfig3_6_Storage_Journal) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 0, 1}
}
func (m *MongodConfig3_6_Storage_Journal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Storage_Journal.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Storage_Journal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Storage_Journal.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Storage_Journal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Storage_Journal.Merge(dst, src)
}
func (m *MongodConfig3_6_Storage_Journal) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Storage_Journal.Size(m)
}
func (m *MongodConfig3_6_Storage_Journal) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Storage_Journal.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Storage_Journal proto.InternalMessageInfo

func (m *MongodConfig3_6_Storage_Journal) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *MongodConfig3_6_Storage_Journal) GetCommitInterval() *wrappers.Int64Value {
	if m != nil {
		return m.CommitInterval
	}
	return nil
}

type MongodConfig3_6_OperationProfiling struct {
	// Mode which specifies operations that should be profiled.
	Mode MongodConfig3_6_OperationProfiling_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6_OperationProfiling_Mode" json:"mode,omitempty"`
	// The slow operation time threshold, in milliseconds. Operations that run
	// for longer than this threshold are considered slow, and are processed by the profiler
	// running in the SLOW_OP mode.
	SlowOpThreshold      *wrappers.Int64Value `protobuf:"bytes,2,opt,name=slow_op_threshold,json=slowOpThreshold,proto3" json:"slow_op_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MongodConfig3_6_OperationProfiling) Reset()         { *m = MongodConfig3_6_OperationProfiling{} }
func (m *MongodConfig3_6_OperationProfiling) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6_OperationProfiling) ProtoMessage()    {}
func (*MongodConfig3_6_OperationProfiling) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 1}
}
func (m *MongodConfig3_6_OperationProfiling) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_OperationProfiling.Unmarshal(m, b)
}
func (m *MongodConfig3_6_OperationProfiling) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_OperationProfiling.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_OperationProfiling) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_OperationProfiling.Merge(dst, src)
}
func (m *MongodConfig3_6_OperationProfiling) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_OperationProfiling.Size(m)
}
func (m *MongodConfig3_6_OperationProfiling) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_OperationProfiling.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_OperationProfiling proto.InternalMessageInfo

func (m *MongodConfig3_6_OperationProfiling) GetMode() MongodConfig3_6_OperationProfiling_Mode {
	if m != nil {
		return m.Mode
	}
	return MongodConfig3_6_OperationProfiling_MODE_UNSPECIFIED
}

func (m *MongodConfig3_6_OperationProfiling) GetSlowOpThreshold() *wrappers.Int64Value {
	if m != nil {
		return m.SlowOpThreshold
	}
	return nil
}

type MongodConfig3_6_Network struct {
	// The maximum number of simultaneous connections that mongod will accept.
	MaxIncomingConnections *wrappers.Int64Value `protobuf:"bytes,1,opt,name=max_incoming_connections,json=maxIncomingConnections,proto3" json:"max_incoming_connections,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *MongodConfig3_6_Network) Reset()         { *m = MongodConfig3_6_Network{} }
func (m *MongodConfig3_6_Network) String() string { return proto.CompactTextString(m) }
func (*MongodConfig3_6_Network) ProtoMessage()    {}
func (*MongodConfig3_6_Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{0, 2}
}
func (m *MongodConfig3_6_Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfig3_6_Network.Unmarshal(m, b)
}
func (m *MongodConfig3_6_Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfig3_6_Network.Marshal(b, m, deterministic)
}
func (dst *MongodConfig3_6_Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfig3_6_Network.Merge(dst, src)
}
func (m *MongodConfig3_6_Network) XXX_Size() int {
	return xxx_messageInfo_MongodConfig3_6_Network.Size(m)
}
func (m *MongodConfig3_6_Network) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfig3_6_Network.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfig3_6_Network proto.InternalMessageInfo

func (m *MongodConfig3_6_Network) GetMaxIncomingConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxIncomingConnections
	}
	return nil
}

type MongodConfigSet3_6 struct {
	// Effective settings for a MongoDB 3.6 cluster (a combination of settings defined
	// in [user_config] and [default_config]).
	EffectiveConfig *MongodConfig3_6 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a MongoDB 3.6 cluster.
	UserConfig *MongodConfig3_6 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a MongoDB 3.6 cluster.
	DefaultConfig        *MongodConfig3_6 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MongodConfigSet3_6) Reset()         { *m = MongodConfigSet3_6{} }
func (m *MongodConfigSet3_6) String() string { return proto.CompactTextString(m) }
func (*MongodConfigSet3_6) ProtoMessage()    {}
func (*MongodConfigSet3_6) Descriptor() ([]byte, []int) {
	return fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2, []int{1}
}
func (m *MongodConfigSet3_6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongodConfigSet3_6.Unmarshal(m, b)
}
func (m *MongodConfigSet3_6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongodConfigSet3_6.Marshal(b, m, deterministic)
}
func (dst *MongodConfigSet3_6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongodConfigSet3_6.Merge(dst, src)
}
func (m *MongodConfigSet3_6) XXX_Size() int {
	return xxx_messageInfo_MongodConfigSet3_6.Size(m)
}
func (m *MongodConfigSet3_6) XXX_DiscardUnknown() {
	xxx_messageInfo_MongodConfigSet3_6.DiscardUnknown(m)
}

var xxx_messageInfo_MongodConfigSet3_6 proto.InternalMessageInfo

func (m *MongodConfigSet3_6) GetEffectiveConfig() *MongodConfig3_6 {
	if m != nil {
		return m.EffectiveConfig
	}
	return nil
}

func (m *MongodConfigSet3_6) GetUserConfig() *MongodConfig3_6 {
	if m != nil {
		return m.UserConfig
	}
	return nil
}

func (m *MongodConfigSet3_6) GetDefaultConfig() *MongodConfig3_6 {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*MongodConfig3_6)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6")
	proto.RegisterType((*MongodConfig3_6_Storage)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Storage")
	proto.RegisterType((*MongodConfig3_6_Storage_WiredTiger)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Storage.WiredTiger")
	proto.RegisterType((*MongodConfig3_6_Storage_WiredTiger_EngineConfig)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Storage.WiredTiger.EngineConfig")
	proto.RegisterType((*MongodConfig3_6_Storage_WiredTiger_CollectionConfig)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Storage.WiredTiger.CollectionConfig")
	proto.RegisterType((*MongodConfig3_6_Storage_Journal)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Storage.Journal")
	proto.RegisterType((*MongodConfig3_6_OperationProfiling)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.OperationProfiling")
	proto.RegisterType((*MongodConfig3_6_Network)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6.Network")
	proto.RegisterType((*MongodConfigSet3_6)(nil), "yandex.cloud.mdb.mongodb.v1.config.MongodConfigSet3_6")
	proto.RegisterEnum("yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor", MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor_name, MongodConfig3_6_Storage_WiredTiger_CollectionConfig_Compressor_value)
	proto.RegisterEnum("yandex.cloud.mdb.mongodb.v1.config.MongodConfig3_6_OperationProfiling_Mode", MongodConfig3_6_OperationProfiling_Mode_name, MongodConfig3_6_OperationProfiling_Mode_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mongodb/v1/config/mongodb3_6.proto", fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2)
}

var fileDescriptor_mongodb3_6_b6b4a9da8ede6ac2 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x71, 0x52, 0xea, 0xe5, 0xa4, 0x6d, 0xbc, 0x03, 0x5a, 0x55, 0xe6, 0x43, 0x28, 0x57,
	0xdc, 0x74, 0x5c, 0x6f, 0x4a, 0x85, 0x54, 0x09, 0xb1, 0x4d, 0x13, 0x08, 0x34, 0xb1, 0x65, 0x77,
	0xa9, 0xa8, 0x04, 0x96, 0x3f, 0x26, 0xae, 0x59, 0x7b, 0xc6, 0xf2, 0x47, 0x52, 0xf6, 0x16, 0x6e,
	0x91, 0x78, 0x0a, 0xde, 0x80, 0x67, 0xe8, 0x1d, 0x8f, 0xc0, 0x13, 0xf0, 0x04, 0xbd, 0x42, 0xf6,
	0x8c, 0xd3, 0x6e, 0x2a, 0xb1, 0x28, 0x5b, 0xee, 0xe2, 0x93, 0xf9, 0xff, 0xfe, 0x73, 0xce, 0x9c,
	0x39, 0x36, 0xf4, 0x7f, 0x72, 0x69, 0x40, 0xae, 0x34, 0x3f, 0x66, 0x65, 0xa0, 0x25, 0x81, 0xa7,
	0x25, 0x8c, 0x86, 0x2c, 0xf0, 0xb4, 0xb9, 0xae, 0xf9, 0x8c, 0xce, 0xa2, 0xb0, 0x89, 0xf4, 0x9d,
	0x43, 0x9c, 0x66, 0xac, 0x60, 0xa8, 0xc7, 0x45, 0xb8, 0x16, 0xe1, 0x24, 0xf0, 0xb0, 0x58, 0x82,
	0xe7, 0x3a, 0xe6, 0x22, 0xf5, 0xa3, 0x90, 0xb1, 0x30, 0x26, 0x5a, 0xad, 0xf0, 0xca, 0x99, 0xb6,
	0xc8, 0xdc, 0x34, 0x25, 0x59, 0xce, 0x19, 0xea, 0x87, 0xaf, 0x18, 0xcf, 0xdd, 0x38, 0x0a, 0xdc,
	0x22, 0x62, 0x94, 0xff, 0xdd, 0xfb, 0x6b, 0x0b, 0xba, 0x93, 0x1a, 0x3a, 0xa8, 0x79, 0x7d, 0xe7,
	0x10, 0x3d, 0x07, 0x39, 0x2f, 0x58, 0xe6, 0x86, 0x64, 0x57, 0xfa, 0x58, 0xfa, 0xa4, 0xf3, 0xf4,
	0x08, 0xbf, 0x7e, 0x23, 0x78, 0x85, 0x82, 0x6d, 0x8e, 0xb0, 0x1a, 0x16, 0x5a, 0xc0, 0xbb, 0x2c,
	0x25, 0x59, 0xed, 0xee, 0xa4, 0x19, 0x9b, 0x45, 0x71, 0x44, 0xc3, 0xdd, 0x56, 0x6d, 0x31, 0x5a,
	0xc7, 0xc2, 0x68, 0x70, 0x66, 0x43, 0xb3, 0x10, 0xbb, 0x17, 0x43, 0x13, 0x68, 0x53, 0x52, 0xec,
	0xb6, 0xd7, 0xcf, 0x65, 0x4a, 0x8a, 0x05, 0xcb, 0x5e, 0x58, 0x15, 0x47, 0xfd, 0x43, 0x06, 0x59,
	0x24, 0x87, 0x42, 0xe8, 0x2c, 0xa2, 0x8c, 0x04, 0x4e, 0x11, 0x85, 0x24, 0x13, 0xe5, 0x1a, 0xbd,
	0x41, 0xb9, 0xf0, 0x79, 0x85, 0x3b, 0xab, 0x68, 0x16, 0x2c, 0x96, 0xbf, 0xd1, 0xf7, 0x20, 0xff,
	0xc8, 0xca, 0x8c, 0xba, 0xb1, 0x28, 0xd8, 0xe0, 0x4d, 0x4c, 0xbe, 0xe6, 0x28, 0xab, 0x61, 0xaa,
	0x7f, 0x6e, 0x00, 0xdc, 0x3a, 0xa3, 0x2b, 0xd8, 0x26, 0x34, 0x8c, 0x28, 0x71, 0x38, 0x48, 0x24,
	0x66, 0x3f, 0x4c, 0x62, 0x78, 0x58, 0xb3, 0xf9, 0x12, 0x6b, 0x8b, 0xdc, 0x79, 0x42, 0xbf, 0x48,
	0xf0, 0xd8, 0x67, 0x71, 0x4c, 0xfc, 0xba, 0x4d, 0x84, 0x3d, 0x4f, 0xf9, 0xfc, 0x81, 0xec, 0x07,
	0x4b, 0xbe, 0xd8, 0x82, 0xe2, 0xaf, 0x44, 0x54, 0x13, 0xb6, 0xee, 0x6e, 0x12, 0x7d, 0x01, 0xdb,
	0xbe, 0xeb, 0x5f, 0x12, 0x27, 0x8f, 0x5e, 0x12, 0x27, 0xf4, 0x44, 0x41, 0x3e, 0xc0, 0xfc, 0xf6,
	0xe1, 0xe6, 0xf6, 0xe1, 0x13, 0x56, 0x7a, 0x31, 0xf9, 0xd6, 0x8d, 0x4b, 0x62, 0x75, 0x6a, 0x89,
	0x1d, 0xbd, 0x24, 0x5f, 0x7a, 0xea, 0xdf, 0x12, 0x28, 0xab, 0xc6, 0xe8, 0x57, 0x09, 0x14, 0x2f,
	0x66, 0xfe, 0x0b, 0xc7, 0x67, 0x49, 0x9a, 0x91, 0x3c, 0x67, 0xbc, 0x89, 0x76, 0x9e, 0x7a, 0xff,
	0x53, 0xb2, 0x78, 0xb0, 0x74, 0xb2, 0xba, 0xb5, 0xf7, 0x6d, 0xa0, 0xf7, 0x15, 0xc0, 0xed, 0x13,
	0x52, 0xe1, 0xc9, 0xc0, 0x98, 0x98, 0xd6, 0xd0, 0xb6, 0x0d, 0xcb, 0x79, 0x3e, 0xb5, 0xcd, 0xe1,
	0x60, 0x3c, 0x1a, 0x0f, 0x4f, 0x94, 0xb7, 0xd0, 0x23, 0xd8, 0x98, 0x1a, 0xd3, 0xa1, 0x22, 0x21,
	0x80, 0x4d, 0x7b, 0xfa, 0xcc, 0x34, 0xbf, 0x53, 0x5a, 0x55, 0xf4, 0xe2, 0x74, 0x7c, 0xac, 0xb4,
	0xd5, 0xdf, 0x24, 0x90, 0x45, 0x97, 0xa1, 0x03, 0x90, 0x09, 0x75, 0xbd, 0x98, 0x04, 0xa2, 0x6c,
	0xea, 0xbd, 0xb2, 0x1d, 0x33, 0x16, 0xf3, 0xa2, 0x35, 0x4b, 0x91, 0x01, 0x5d, 0x9f, 0x25, 0x49,
	0x54, 0x38, 0x11, 0x2d, 0x48, 0x36, 0x5f, 0x76, 0xfe, 0xfb, 0xf7, 0xd4, 0x63, 0x5a, 0x1c, 0x1e,
	0xd4, 0xf2, 0xe3, 0x77, 0x6e, 0xae, 0xf5, 0xb7, 0xf5, 0xbd, 0x4f, 0xf7, 0xf7, 0xad, 0x1d, 0x2e,
	0x1f, 0x0b, 0xb5, 0xfa, 0x73, 0x0b, 0xd0, 0xfd, 0x89, 0x81, 0x1c, 0xd8, 0x48, 0x58, 0x40, 0x44,
	0xd9, 0xbf, 0x79, 0x98, 0x39, 0x84, 0x27, 0x2c, 0x20, 0x56, 0x0d, 0x46, 0x06, 0x3c, 0xce, 0x63,
	0xb6, 0x70, 0x58, 0xea, 0x14, 0x97, 0x19, 0xc9, 0x2f, 0x59, 0x1c, 0xfc, 0x97, 0x54, 0x36, 0x6f,
	0xae, 0xf5, 0xd6, 0xe7, 0xfb, 0x56, 0xb7, 0x52, 0x1b, 0xe9, 0x59, 0xa3, 0xed, 0x1d, 0xc1, 0x46,
	0x85, 0x47, 0xef, 0x81, 0x32, 0x31, 0x4e, 0x86, 0x2b, 0x27, 0x23, 0x43, 0xdb, 0x18, 0x8d, 0x14,
	0x09, 0x75, 0x40, 0xb6, 0x4f, 0x8d, 0x73, 0xc7, 0x30, 0x95, 0x56, 0x15, 0x7d, 0x76, 0x7a, 0xaa,
	0xb4, 0x55, 0x0a, 0xb2, 0x98, 0x66, 0xc8, 0x87, 0xdd, 0xc4, 0xbd, 0x72, 0x22, 0xea, 0xb3, 0x24,
	0xa2, 0x61, 0x75, 0xd9, 0x28, 0x6f, 0x95, 0x5c, 0x1c, 0xd4, 0xbf, 0xee, 0x6f, 0xeb, 0xe6, 0x5a,
	0x7f, 0xa4, 0xef, 0xef, 0xe9, 0x87, 0xfd, 0xcf, 0x0e, 0xac, 0x27, 0x89, 0x7b, 0x35, 0x16, 0xa4,
	0xc1, 0x2d, 0xa8, 0xf7, 0x7b, 0x0b, 0xd0, 0xdd, 0x7a, 0xd9, 0xa4, 0xa8, 0xde, 0x31, 0x3f, 0x80,
	0x42, 0x66, 0xb3, 0x6a, 0xd1, 0x7c, 0x65, 0xc8, 0xf4, 0xd7, 0x38, 0x01, 0xab, 0xbb, 0x84, 0x89,
	0x9b, 0x75, 0x06, 0x9d, 0x32, 0x27, 0xd9, 0xab, 0x03, 0x64, 0x2d, 0x34, 0x54, 0x1c, 0x41, 0xbd,
	0x80, 0x9d, 0x80, 0xcc, 0xdc, 0x32, 0x2e, 0x1a, 0x70, 0x7b, 0x7d, 0xf0, 0xb6, 0x40, 0xf1, 0xc8,
	0xb1, 0x79, 0x31, 0x0d, 0xa3, 0xe2, 0xb2, 0xf4, 0xb0, 0xcf, 0x12, 0x8d, 0xf3, 0xf6, 0xf8, 0x5b,
	0x3b, 0x64, 0x7b, 0x21, 0xa1, 0xf5, 0x19, 0x68, 0xaf, 0xff, 0x8e, 0x38, 0x12, 0x11, 0x6f, 0xb3,
	0x56, 0xf4, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xb4, 0xd3, 0xe3, 0x7c, 0x08, 0x00, 0x00,
}
