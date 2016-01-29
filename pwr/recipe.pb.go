// Code generated by protoc-gen-go.
// source: pwr/recipe.proto
// DO NOT EDIT!

/*
Package pwr is a generated protocol buffer package.

It is generated from these files:
	pwr/recipe.proto

It has these top-level messages:
	RecipeHeader
	Container
	SignatureHeader
	BlockHash
	SyncHeader
	SyncOp
	Hash
*/
package pwr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HashType int32

const (
	// librsync default
	HashType_MD5 HashType = 0
	// cf. https://godoc.org/golang.org/x/crypto/sha3
	HashType_SHAKESUM128 HashType = 1
)

var HashType_name = map[int32]string{
	0: "MD5",
	1: "SHAKESUM128",
}
var HashType_value = map[string]int32{
	"MD5":         0,
	"SHAKESUM128": 1,
}

func (x HashType) String() string {
	return proto.EnumName(HashType_name, int32(x))
}
func (HashType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RecipeHeader_Version int32

const (
	RecipeHeader_V1 RecipeHeader_Version = 0
)

var RecipeHeader_Version_name = map[int32]string{
	0: "V1",
}
var RecipeHeader_Version_value = map[string]int32{
	"V1": 0,
}

func (x RecipeHeader_Version) String() string {
	return proto.EnumName(RecipeHeader_Version_name, int32(x))
}
func (RecipeHeader_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type RecipeHeader_Compression int32

const (
	RecipeHeader_UNCOMPRESSED RecipeHeader_Compression = 0
	RecipeHeader_BROTLI       RecipeHeader_Compression = 1
)

var RecipeHeader_Compression_name = map[int32]string{
	0: "UNCOMPRESSED",
	1: "BROTLI",
}
var RecipeHeader_Compression_value = map[string]int32{
	"UNCOMPRESSED": 0,
	"BROTLI":       1,
}

func (x RecipeHeader_Compression) String() string {
	return proto.EnumName(RecipeHeader_Compression_name, int32(x))
}
func (RecipeHeader_Compression) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type SyncOp_Type int32

const (
	SyncOp_BLOCK          SyncOp_Type = 0
	SyncOp_BLOCK_RANGE    SyncOp_Type = 1
	SyncOp_DATA           SyncOp_Type = 2
	SyncOp_HEY_YOU_DID_IT SyncOp_Type = 2049
)

var SyncOp_Type_name = map[int32]string{
	0:    "BLOCK",
	1:    "BLOCK_RANGE",
	2:    "DATA",
	2049: "HEY_YOU_DID_IT",
}
var SyncOp_Type_value = map[string]int32{
	"BLOCK":          0,
	"BLOCK_RANGE":    1,
	"DATA":           2,
	"HEY_YOU_DID_IT": 2049,
}

func (x SyncOp_Type) String() string {
	return proto.EnumName(SyncOp_Type_name, int32(x))
}
func (SyncOp_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type RecipeHeader struct {
	Version          RecipeHeader_Version     `protobuf:"varint,1,opt,name=version,enum=io.itch.wharf.pwr.RecipeHeader_Version" json:"version,omitempty"`
	Compression      RecipeHeader_Compression `protobuf:"varint,2,opt,name=compression,enum=io.itch.wharf.pwr.RecipeHeader_Compression" json:"compression,omitempty"`
	CompressionLevel int32                    `protobuf:"varint,3,opt,name=compressionLevel" json:"compressionLevel,omitempty"`
}

func (m *RecipeHeader) Reset()                    { *m = RecipeHeader{} }
func (m *RecipeHeader) String() string            { return proto.CompactTextString(m) }
func (*RecipeHeader) ProtoMessage()               {}
func (*RecipeHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Container struct {
	Size     int64                `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Dirs     []*Container_Dir     `protobuf:"bytes,2,rep,name=dirs" json:"dirs,omitempty"`
	Files    []*Container_File    `protobuf:"bytes,3,rep,name=files" json:"files,omitempty"`
	Symlinks []*Container_Symlink `protobuf:"bytes,4,rep,name=symlinks" json:"symlinks,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetDirs() []*Container_Dir {
	if m != nil {
		return m.Dirs
	}
	return nil
}

func (m *Container) GetFiles() []*Container_File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Container) GetSymlinks() []*Container_Symlink {
	if m != nil {
		return m.Symlinks
	}
	return nil
}

type Container_Dir struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Mode uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
}

func (m *Container_Dir) Reset()                    { *m = Container_Dir{} }
func (m *Container_Dir) String() string            { return proto.CompactTextString(m) }
func (*Container_Dir) ProtoMessage()               {}
func (*Container_Dir) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Container_File struct {
	Path       string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Mode       uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	Size       int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	BlockIndex int64  `protobuf:"varint,4,opt,name=blockIndex" json:"blockIndex,omitempty"`
	BlockSpan  int64  `protobuf:"varint,5,opt,name=blockSpan" json:"blockSpan,omitempty"`
}

func (m *Container_File) Reset()                    { *m = Container_File{} }
func (m *Container_File) String() string            { return proto.CompactTextString(m) }
func (*Container_File) ProtoMessage()               {}
func (*Container_File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type Container_Symlink struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Mode uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	Dest string `protobuf:"bytes,3,opt,name=dest" json:"dest,omitempty"`
}

func (m *Container_Symlink) Reset()                    { *m = Container_Symlink{} }
func (m *Container_Symlink) String() string            { return proto.CompactTextString(m) }
func (*Container_Symlink) ProtoMessage()               {}
func (*Container_Symlink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

type SignatureHeader struct {
	BlockCount int64    `protobuf:"varint,1,opt,name=blockCount" json:"blockCount,omitempty"`
	HashType   HashType `protobuf:"varint,2,opt,name=hashType,enum=io.itch.wharf.pwr.HashType" json:"hashType,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BlockHash struct {
	FileIndex  int64  `protobuf:"varint,1,opt,name=fileIndex" json:"fileIndex,omitempty"`
	BlockIndex int64  `protobuf:"varint,2,opt,name=blockIndex" json:"blockIndex,omitempty"`
	WeakHash   int32  `protobuf:"varint,3,opt,name=weakHash" json:"weakHash,omitempty"`
	StrongHash []byte `protobuf:"bytes,4,opt,name=strongHash,proto3" json:"strongHash,omitempty"`
}

func (m *BlockHash) Reset()                    { *m = BlockHash{} }
func (m *BlockHash) String() string            { return proto.CompactTextString(m) }
func (*BlockHash) ProtoMessage()               {}
func (*BlockHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SyncHeader struct {
	FileIndex int64 `protobuf:"varint,1,opt,name=fileIndex" json:"fileIndex,omitempty"`
}

func (m *SyncHeader) Reset()                    { *m = SyncHeader{} }
func (m *SyncHeader) String() string            { return proto.CompactTextString(m) }
func (*SyncHeader) ProtoMessage()               {}
func (*SyncHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SyncOp struct {
	Type       SyncOp_Type `protobuf:"varint,1,opt,name=type,enum=io.itch.wharf.pwr.SyncOp_Type" json:"type,omitempty"`
	FileIndex  int64       `protobuf:"varint,2,opt,name=fileIndex" json:"fileIndex,omitempty"`
	BlockIndex int64       `protobuf:"varint,3,opt,name=blockIndex" json:"blockIndex,omitempty"`
	BlockSpan  int64       `protobuf:"varint,4,opt,name=blockSpan" json:"blockSpan,omitempty"`
	Data       []byte      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SyncOp) Reset()                    { *m = SyncOp{} }
func (m *SyncOp) String() string            { return proto.CompactTextString(m) }
func (*SyncOp) ProtoMessage()               {}
func (*SyncOp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Hash struct {
	Type     HashType `protobuf:"varint,1,opt,name=type,enum=io.itch.wharf.pwr.HashType" json:"type,omitempty"`
	Contents []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*RecipeHeader)(nil), "io.itch.wharf.pwr.RecipeHeader")
	proto.RegisterType((*Container)(nil), "io.itch.wharf.pwr.Container")
	proto.RegisterType((*Container_Dir)(nil), "io.itch.wharf.pwr.Container.Dir")
	proto.RegisterType((*Container_File)(nil), "io.itch.wharf.pwr.Container.File")
	proto.RegisterType((*Container_Symlink)(nil), "io.itch.wharf.pwr.Container.Symlink")
	proto.RegisterType((*SignatureHeader)(nil), "io.itch.wharf.pwr.SignatureHeader")
	proto.RegisterType((*BlockHash)(nil), "io.itch.wharf.pwr.BlockHash")
	proto.RegisterType((*SyncHeader)(nil), "io.itch.wharf.pwr.SyncHeader")
	proto.RegisterType((*SyncOp)(nil), "io.itch.wharf.pwr.SyncOp")
	proto.RegisterType((*Hash)(nil), "io.itch.wharf.pwr.Hash")
	proto.RegisterEnum("io.itch.wharf.pwr.HashType", HashType_name, HashType_value)
	proto.RegisterEnum("io.itch.wharf.pwr.RecipeHeader_Version", RecipeHeader_Version_name, RecipeHeader_Version_value)
	proto.RegisterEnum("io.itch.wharf.pwr.RecipeHeader_Compression", RecipeHeader_Compression_name, RecipeHeader_Compression_value)
	proto.RegisterEnum("io.itch.wharf.pwr.SyncOp_Type", SyncOp_Type_name, SyncOp_Type_value)
}

var fileDescriptor0 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0xc5, 0xb1, 0xf3, 0x35, 0x49, 0xc1, 0x6c, 0x2f, 0x51, 0x5a, 0x21, 0x6a, 0x21, 0x15, 0x81,
	0x6a, 0x44, 0xda, 0xaa, 0x3d, 0x36, 0x5f, 0x6d, 0x22, 0x08, 0xa9, 0xd6, 0x01, 0x89, 0x5e, 0x22,
	0xe3, 0x2c, 0xc4, 0x25, 0xd8, 0x96, 0x6d, 0xa0, 0xf4, 0xd4, 0x1e, 0x7a, 0xeb, 0x1f, 0xeb, 0xaf,
	0xe9, 0x5f, 0xe8, 0xee, 0xd8, 0x31, 0x06, 0x22, 0xb7, 0xbd, 0xcd, 0xce, 0xbe, 0x37, 0x7e, 0x6f,
	0x66, 0xd6, 0xa0, 0x7a, 0xd7, 0xfe, 0x8e, 0xcf, 0x2c, 0xdb, 0x63, 0xba, 0xe7, 0xbb, 0xa1, 0x4b,
	0x56, 0x6d, 0x57, 0xb7, 0x43, 0x6b, 0xaa, 0x5f, 0x4f, 0x4d, 0xff, 0x54, 0xe7, 0xf7, 0xda, 0xcf,
	0x1c, 0x54, 0x29, 0x62, 0x7a, 0xcc, 0x9c, 0x30, 0x9f, 0x34, 0xa1, 0x78, 0xc5, 0xfc, 0xc0, 0x76,
	0x9d, 0x9a, 0xb4, 0x2e, 0x6d, 0x2e, 0x37, 0x9e, 0xeb, 0x0f, 0x58, 0x7a, 0x9a, 0xa1, 0x1f, 0x45,
	0x70, 0x3a, 0xe7, 0x91, 0x01, 0x54, 0x2c, 0xf7, 0xc2, 0xf3, 0x59, 0x80, 0x65, 0x72, 0x58, 0x66,
	0xfb, 0x6f, 0x65, 0xda, 0xb7, 0x14, 0x9a, 0xe6, 0x93, 0x2d, 0x50, 0x53, 0xc7, 0x7d, 0x76, 0xc5,
	0x66, 0x35, 0x99, 0xd7, 0xcc, 0xd3, 0x07, 0x79, 0x6d, 0x15, 0x8a, 0xb1, 0x1c, 0x52, 0x80, 0xdc,
	0xd1, 0xae, 0xba, 0xa4, 0x6d, 0x43, 0x25, 0x55, 0x9a, 0xa8, 0x50, 0x3d, 0x3c, 0x68, 0x0f, 0x07,
	0x1f, 0x69, 0xd7, 0x30, 0xba, 0x1d, 0x75, 0x89, 0x00, 0x14, 0x5a, 0x74, 0x38, 0xda, 0xef, 0xab,
	0x92, 0xf6, 0x4b, 0x86, 0x72, 0xdb, 0x75, 0x42, 0xd3, 0x76, 0x78, 0x2f, 0x08, 0x28, 0x81, 0xfd,
	0x95, 0x61, 0x23, 0x64, 0x8a, 0x31, 0x79, 0x05, 0xca, 0xc4, 0xf6, 0x03, 0xee, 0x4a, 0xde, 0xac,
	0x34, 0xd6, 0x17, 0xb8, 0x4a, 0xf8, 0x7a, 0xc7, 0xf6, 0x29, 0xa2, 0xc9, 0x1b, 0xc8, 0x9f, 0xda,
	0x33, 0x16, 0x70, 0xe1, 0x82, 0xf6, 0x2c, 0x93, 0xf6, 0x9e, 0x23, 0x69, 0x84, 0x27, 0xef, 0xa0,
	0x14, 0xdc, 0x5c, 0xcc, 0x6c, 0xe7, 0x3c, 0xa8, 0x29, 0xc8, 0xdd, 0xc8, 0xe4, 0x1a, 0x11, 0x98,
	0x26, 0xac, 0xfa, 0x0b, 0x90, 0xb9, 0x0e, 0xe1, 0xc5, 0x33, 0xc3, 0x29, 0x7a, 0x29, 0x53, 0x8c,
	0x45, 0xee, 0xc2, 0x9d, 0x30, 0x9c, 0xd0, 0x23, 0x8a, 0x71, 0xfd, 0x9b, 0x04, 0x8a, 0x10, 0xf0,
	0xaf, 0x84, 0xa4, 0x49, 0x72, 0xaa, 0x49, 0x6b, 0x00, 0x27, 0x33, 0xd7, 0x3a, 0xef, 0x3b, 0x13,
	0xf6, 0x85, 0xeb, 0x16, 0x37, 0xa9, 0x0c, 0x79, 0x0a, 0x65, 0x3c, 0x19, 0x9e, 0xe9, 0xd4, 0xf2,
	0x78, 0x7d, 0x9b, 0xa8, 0x77, 0xa1, 0x18, 0xdb, 0xf8, 0x1f, 0x11, 0x13, 0x16, 0x84, 0x28, 0x82,
	0xe3, 0x44, 0xac, 0x7d, 0x86, 0x15, 0xc3, 0x3e, 0x73, 0xcc, 0xf0, 0xd2, 0x9f, 0x2f, 0xf7, 0x5c,
	0x57, 0xdb, 0xbd, 0x74, 0xc2, 0x78, 0xac, 0xa9, 0x0c, 0x1f, 0x53, 0x69, 0x6a, 0x06, 0xd3, 0xd1,
	0x8d, 0xc7, 0xe2, 0xb5, 0x7d, 0xb2, 0xa0, 0xdb, 0xbd, 0x18, 0x42, 0x13, 0xb0, 0xf6, 0x43, 0x82,
	0x72, 0x4b, 0xd4, 0x11, 0x77, 0xc2, 0x9e, 0x98, 0x5e, 0xe4, 0x3e, 0xfa, 0xca, 0x6d, 0xe2, 0x5e,
	0x73, 0x72, 0x0f, 0x9a, 0x53, 0x87, 0xd2, 0x35, 0x33, 0xb1, 0x52, 0xbc, 0xe7, 0xc9, 0x59, 0x70,
	0x83, 0xd0, 0x77, 0x9d, 0x33, 0xbc, 0x15, 0x8d, 0xad, 0xd2, 0x54, 0x46, 0xdb, 0x02, 0x30, 0x6e,
	0x1c, 0x2b, 0xb6, 0x9b, 0xa9, 0x43, 0xfb, 0x2d, 0x41, 0x41, 0x80, 0x87, 0x1e, 0x69, 0x80, 0x12,
	0x0a, 0xcf, 0xd1, 0x8b, 0x5f, 0x5b, 0xe0, 0x39, 0x02, 0xea, 0x68, 0x1b, 0xb1, 0x77, 0x8b, 0xe7,
	0xb2, 0x4d, 0xca, 0xd9, 0x1b, 0xa0, 0xdc, 0xdb, 0x00, 0x1c, 0xa7, 0x19, 0x9a, 0xb8, 0x1a, 0x55,
	0x8a, 0xb1, 0xd6, 0x04, 0x45, 0x7c, 0x9d, 0x94, 0x21, 0xdf, 0xda, 0x1f, 0xb6, 0xf7, 0xf8, 0xcb,
	0x5d, 0x81, 0x0a, 0x86, 0x63, 0xda, 0x3c, 0xf8, 0xd0, 0x55, 0x25, 0x52, 0x02, 0xa5, 0xd3, 0x1c,
	0x35, 0xd5, 0x1c, 0x79, 0x0c, 0xcb, 0xbd, 0xee, 0xf1, 0xf8, 0x78, 0x78, 0x38, 0xee, 0xf4, 0x3b,
	0xe3, 0xfe, 0x48, 0xfd, 0xae, 0x6a, 0x06, 0x28, 0xd8, 0xc5, 0x9d, 0x3b, 0x76, 0x33, 0x47, 0x1c,
	0x79, 0xe5, 0x23, 0xb1, 0xf8, 0x13, 0x63, 0x4e, 0x18, 0xa0, 0xd5, 0x2a, 0x4d, 0xce, 0x5b, 0x1b,
	0x50, 0x9a, 0xa3, 0x49, 0x11, 0xe4, 0x41, 0xe7, 0x75, 0xa4, 0xcc, 0xe8, 0x35, 0xf7, 0xba, 0xc6,
	0xe1, 0x60, 0xb7, 0xf1, 0x56, 0x95, 0x5a, 0xf9, 0x4f, 0x32, 0xaf, 0x7b, 0x52, 0xc0, 0x1f, 0xf1,
	0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x7e, 0x15, 0x7f, 0x9c, 0x05, 0x00, 0x00,
}
