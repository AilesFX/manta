// Code generated by protoc-gen-go.
// source: steammessages_cloud.steamworkssdk.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CCloud_GetUploadServerInfo_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Request) Reset()         { *m = CCloud_GetUploadServerInfo_Request{} }
func (m *CCloud_GetUploadServerInfo_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Request) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor37, []int{0}
}

func (m *CCloud_GetUploadServerInfo_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_GetUploadServerInfo_Response struct {
	ServerUrl        *string `protobuf:"bytes,1,opt,name=server_url" json:"server_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Response) Reset()         { *m = CCloud_GetUploadServerInfo_Response{} }
func (m *CCloud_GetUploadServerInfo_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Response) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor37, []int{1}
}

func (m *CCloud_GetUploadServerInfo_Response) GetServerUrl() string {
	if m != nil && m.ServerUrl != nil {
		return *m.ServerUrl
	}
	return ""
}

type CCloud_GetFileDetails_Request struct {
	Ugcid            *uint64 `protobuf:"varint,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetFileDetails_Request) Reset()                    { *m = CCloud_GetFileDetails_Request{} }
func (m *CCloud_GetFileDetails_Request) String() string            { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Request) ProtoMessage()               {}
func (*CCloud_GetFileDetails_Request) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{2} }

func (m *CCloud_GetFileDetails_Request) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_GetFileDetails_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_UserFile struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Ugcid            *uint64 `protobuf:"varint,2,opt,name=ugcid" json:"ugcid,omitempty"`
	Filename         *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	FileSize         *uint32 `protobuf:"varint,5,opt,name=file_size" json:"file_size,omitempty"`
	Url              *string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	SteamidCreator   *uint64 `protobuf:"fixed64,7,opt,name=steamid_creator" json:"steamid_creator,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_UserFile) Reset()                    { *m = CCloud_UserFile{} }
func (m *CCloud_UserFile) String() string            { return proto.CompactTextString(m) }
func (*CCloud_UserFile) ProtoMessage()               {}
func (*CCloud_UserFile) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{3} }

func (m *CCloud_UserFile) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_UserFile) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_UserFile) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_UserFile) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CCloud_UserFile) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_UserFile) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *CCloud_UserFile) GetSteamidCreator() uint64 {
	if m != nil && m.SteamidCreator != nil {
		return *m.SteamidCreator
	}
	return 0
}

type CCloud_GetFileDetails_Response struct {
	Details          *CCloud_UserFile `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CCloud_GetFileDetails_Response) Reset()                    { *m = CCloud_GetFileDetails_Response{} }
func (m *CCloud_GetFileDetails_Response) String() string            { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Response) ProtoMessage()               {}
func (*CCloud_GetFileDetails_Response) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{4} }

func (m *CCloud_GetFileDetails_Response) GetDetails() *CCloud_UserFile {
	if m != nil {
		return m.Details
	}
	return nil
}

type CCloud_EnumerateUserFiles_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ExtendedDetails  *bool   `protobuf:"varint,2,opt,name=extended_details" json:"extended_details,omitempty"`
	Count            *uint32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	StartIndex       *uint32 `protobuf:"varint,4,opt,name=start_index" json:"start_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Request) Reset()         { *m = CCloud_EnumerateUserFiles_Request{} }
func (m *CCloud_EnumerateUserFiles_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Request) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor37, []int{5}
}

func (m *CCloud_EnumerateUserFiles_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetExtendedDetails() bool {
	if m != nil && m.ExtendedDetails != nil {
		return *m.ExtendedDetails
	}
	return false
}

func (m *CCloud_EnumerateUserFiles_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

type CCloud_EnumerateUserFiles_Response struct {
	Files            []*CCloud_UserFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	TotalFiles       *uint32            `protobuf:"varint,2,opt,name=total_files" json:"total_files,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Response) Reset()         { *m = CCloud_EnumerateUserFiles_Response{} }
func (m *CCloud_EnumerateUserFiles_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Response) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor37, []int{6}
}

func (m *CCloud_EnumerateUserFiles_Response) GetFiles() []*CCloud_UserFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CCloud_EnumerateUserFiles_Response) GetTotalFiles() uint32 {
	if m != nil && m.TotalFiles != nil {
		return *m.TotalFiles
	}
	return 0
}

type CCloud_Delete_Request struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_Delete_Request) Reset()                    { *m = CCloud_Delete_Request{} }
func (m *CCloud_Delete_Request) String() string            { return proto.CompactTextString(m) }
func (*CCloud_Delete_Request) ProtoMessage()               {}
func (*CCloud_Delete_Request) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{7} }

func (m *CCloud_Delete_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_Delete_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_Delete_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_Delete_Response) Reset()                    { *m = CCloud_Delete_Response{} }
func (m *CCloud_Delete_Response) String() string            { return proto.CompactTextString(m) }
func (*CCloud_Delete_Response) ProtoMessage()               {}
func (*CCloud_Delete_Response) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{8} }

func init() {
	proto.RegisterType((*CCloud_GetUploadServerInfo_Request)(nil), "dota.CCloud_GetUploadServerInfo_Request")
	proto.RegisterType((*CCloud_GetUploadServerInfo_Response)(nil), "dota.CCloud_GetUploadServerInfo_Response")
	proto.RegisterType((*CCloud_GetFileDetails_Request)(nil), "dota.CCloud_GetFileDetails_Request")
	proto.RegisterType((*CCloud_UserFile)(nil), "dota.CCloud_UserFile")
	proto.RegisterType((*CCloud_GetFileDetails_Response)(nil), "dota.CCloud_GetFileDetails_Response")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Request)(nil), "dota.CCloud_EnumerateUserFiles_Request")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Response)(nil), "dota.CCloud_EnumerateUserFiles_Response")
	proto.RegisterType((*CCloud_Delete_Request)(nil), "dota.CCloud_Delete_Request")
	proto.RegisterType((*CCloud_Delete_Response)(nil), "dota.CCloud_Delete_Response")
}

func init() { proto.RegisterFile("steammessages_cloud.steamworkssdk.proto", fileDescriptor37) }

var fileDescriptor37 = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0xc7, 0x35, 0x24, 0x9b, 0xb6, 0x8e, 0x42, 0x8b, 0xab, 0xc2, 0x6a, 0xcb, 0x87, 0x3b, 0x4d,
	0xc9, 0x46, 0x20, 0x53, 0x55, 0x80, 0x04, 0x48, 0x48, 0xdd, 0x2c, 0x94, 0x4a, 0x20, 0xa4, 0x44,
	0x11, 0xaa, 0x10, 0x1a, 0x79, 0x77, 0xce, 0x6e, 0xac, 0x78, 0xec, 0xc1, 0xf6, 0x34, 0x01, 0xe5,
	0x02, 0xb8, 0x86, 0x5b, 0x6e, 0xb8, 0xe7, 0x0d, 0x78, 0x03, 0x5e, 0x86, 0xb7, 0xe0, 0xd8, 0x33,
	0xb3, 0xd9, 0x09, 0x9b, 0x00, 0x77, 0x59, 0xcf, 0xf9, 0xf8, 0x9d, 0x73, 0xfe, 0xe7, 0x84, 0xec,
	0x38, 0x0f, 0xa2, 0x28, 0xc0, 0x39, 0x31, 0x07, 0x97, 0x4d, 0x95, 0xa9, 0x72, 0x1e, 0xdf, 0x4e,
	0x8c, 0x3d, 0x76, 0x2e, 0x3f, 0xe6, 0xa5, 0x35, 0xde, 0xd0, 0xf5, 0xdc, 0x78, 0x31, 0xe0, 0x5d,
	0xf3, 0x4a, 0xcb, 0x99, 0x84, 0x3c, 0x9b, 0x08, 0x07, 0xab, 0xbc, 0xd2, 0x9c, 0xa4, 0x7b, 0x7b,
	0x21, 0x66, 0xf6, 0x04, 0xfc, 0x61, 0xa9, 0x8c, 0xc8, 0x0f, 0xc0, 0x3e, 0x07, 0xfb, 0x54, 0xcf,
	0x4c, 0xb6, 0x0f, 0xdf, 0x56, 0xe0, 0x3c, 0xfd, 0x98, 0xf4, 0x44, 0x59, 0xca, 0xbc, 0x9f, 0xb0,
	0x64, 0xb8, 0x35, 0x7a, 0xe7, 0xa7, 0x3f, 0xfa, 0x6f, 0x3d, 0x2e, 0x4b, 0xf6, 0x74, 0xcc, 0xbc,
	0x61, 0x27, 0x47, 0x72, 0x7a, 0xc4, 0x04, 0x9b, 0x49, 0x05, 0xec, 0x44, 0x2a, 0xc5, 0x26, 0xc0,
	0xaa, 0x18, 0x0b, 0x72, 0x34, 0xe0, 0xe9, 0x07, 0xe4, 0xfe, 0x95, 0x59, 0x5c, 0x69, 0xb4, 0x03,
	0x4a, 0x09, 0x71, 0xf1, 0x39, 0xab, 0xac, 0x8a, 0xb9, 0x6e, 0xa4, 0x3f, 0x27, 0xe4, 0xb5, 0x73,
	0xdf, 0x4f, 0x31, 0xc5, 0x18, 0xbc, 0x90, 0xca, 0x2d, 0xe0, 0x3e, 0x22, 0xbd, 0x6a, 0x3e, 0x6d,
	0xe0, 0xd6, 0x47, 0x6f, 0x23, 0xdc, 0x10, 0xc1, 0xcc, 0x8c, 0xf9, 0x23, 0x60, 0xd1, 0xb5, 0x46,
	0x43, 0xd4, 0x39, 0x78, 0x96, 0xd7, 0xfe, 0x6c, 0x66, 0x2c, 0xa7, 0xbc, 0xad, 0xec, 0x85, 0x58,
	0xd9, 0x1b, 0xe8, 0x7c, 0xb7, 0xad, 0x0c, 0xbd, 0xa3, 0xdf, 0x04, 0x94, 0xd1, 0x73, 0x17, 0x2b,
	0xf9, 0x25, 0x21, 0x37, 0x1b, 0x9c, 0x43, 0x64, 0x0d, 0x3c, 0x74, 0xab, 0xd3, 0x9d, 0xf0, 0xb3,
	0xe6, 0x09, 0x21, 0xd7, 0xe9, 0x2d, 0x72, 0x3d, 0x44, 0xd1, 0xa2, 0x80, 0xfe, 0x5a, 0x28, 0x89,
	0xbe, 0x44, 0x6e, 0x78, 0x89, 0x23, 0xf2, 0xa2, 0x28, 0xfb, 0xeb, 0xd1, 0x08, 0x9f, 0x82, 0x51,
	0xe6, 0xe4, 0xf7, 0xd0, 0xef, 0xc5, 0x30, 0x9b, 0x64, 0x2d, 0x74, 0x61, 0x23, 0xba, 0xbc, 0x42,
	0x6e, 0xc6, 0xe9, 0xc9, 0x3c, 0x9b, 0x5a, 0x10, 0xde, 0xd8, 0xfe, 0x35, 0xfc, 0xb0, 0x91, 0x7e,
	0x46, 0x5e, 0xbf, 0xac, 0x3b, 0x4d, 0x53, 0xdf, 0x24, 0xd7, 0x9a, 0x8a, 0x23, 0xdf, 0xe6, 0xa3,
	0x3b, 0x3c, 0x28, 0x85, 0x5f, 0xa8, 0x22, 0xfd, 0x6b, 0x8d, 0xdc, 0x6b, 0xde, 0x3e, 0xd1, 0x55,
	0x01, 0x56, 0x78, 0x68, 0x3f, 0x9e, 0x37, 0xfb, 0xdd, 0xae, 0x12, 0x1e, 0x60, 0xbf, 0xee, 0x9d,
	0x2b, 0x01, 0x5a, 0xc7, 0x45, 0xf3, 0x1c, 0xce, 0x81, 0xd3, 0xdf, 0x12, 0x72, 0x0b, 0x4e, 0x3d,
	0x68, 0x14, 0x44, 0xd6, 0xd2, 0x84, 0xf6, 0x5c, 0x1f, 0xfd, 0x90, 0x60, 0x88, 0xb3, 0xe1, 0x97,
	0xa5, 0x97, 0x46, 0x0b, 0xb5, 0xcb, 0xb0, 0x0e, 0xd6, 0xda, 0x2e, 0x66, 0x35, 0x11, 0xd3, 0x63,
	0x66, 0xf4, 0x52, 0xd8, 0x99, 0xa9, 0x74, 0xce, 0xd9, 0x18, 0x66, 0xa2, 0x52, 0x3e, 0x0c, 0x07,
	0xbf, 0xab, 0xef, 0x98, 0x05, 0x5f, 0x59, 0x1d, 0x74, 0x87, 0xa6, 0x22, 0x90, 0xe5, 0x4c, 0xe8,
	0x9c, 0x1d, 0x3e, 0xd9, 0x0b, 0x7f, 0x36, 0xaa, 0x58, 0x8e, 0x41, 0xcf, 0x48, 0x6f, 0x8a, 0x7f,
	0xf9, 0x38, 0x9e, 0xad, 0x51, 0x81, 0x40, 0x72, 0x09, 0xe8, 0x0b, 0x71, 0x2a, 0x8b, 0xaa, 0x60,
	0x58, 0xdb, 0x04, 0x6c, 0x88, 0x60, 0xc1, 0xb5, 0x39, 0xeb, 0x74, 0x35, 0x9a, 0x74, 0x6c, 0x2a,
	0x94, 0xea, 0x42, 0x09, 0x56, 0x34, 0xfe, 0xe8, 0xf8, 0xde, 0xc3, 0x87, 0x4d, 0xea, 0x16, 0x93,
	0x53, 0x4b, 0x36, 0x51, 0x09, 0xd6, 0x67, 0x12, 0x2b, 0x3e, 0x8d, 0x7a, 0xd8, 0x1a, 0x7d, 0x83,
	0x0c, 0xcf, 0x96, 0x18, 0x0e, 0x82, 0x85, 0xd4, 0x73, 0x16, 0x8d, 0x42, 0xdc, 0x09, 0xcc, 0xa5,
	0x5e, 0x34, 0x1c, 0xcd, 0x98, 0xf0, 0xdd, 0xcc, 0xa1, 0xce, 0x68, 0xa5, 0x83, 0x63, 0x53, 0xb8,
	0x92, 0xce, 0xf3, 0x34, 0x5b, 0x6c, 0xfd, 0xca, 0x51, 0x37, 0xca, 0xd9, 0x26, 0xbd, 0x08, 0x8b,
	0xb3, 0x5e, 0xbb, 0x54, 0x37, 0xf4, 0x36, 0xd9, 0xf4, 0xf8, 0xae, 0xb2, 0xda, 0x36, 0xee, 0x51,
	0xfa, 0x8c, 0xdc, 0x69, 0xec, 0xc6, 0xa0, 0xc0, 0xc3, 0x42, 0x3f, 0xcb, 0xdb, 0x10, 0x17, 0xfc,
	0x7f, 0x6f, 0x60, 0x9f, 0xbc, 0x7c, 0x31, 0x74, 0xcd, 0xfb, 0xe8, 0xd7, 0x1e, 0xe9, 0xc5, 0x2f,
	0xf4, 0xf7, 0x84, 0xdc, 0x5e, 0x71, 0x69, 0xe8, 0xb0, 0x53, 0xc2, 0x15, 0x17, 0x6f, 0xb0, 0xfb,
	0x1f, 0x2c, 0xeb, 0xb4, 0xe9, 0x87, 0x48, 0xfc, 0xfe, 0x7e, 0x9c, 0xa7, 0x8b, 0xc8, 0x87, 0xfb,
	0x9f, 0xb7, 0xed, 0xc6, 0x23, 0x5b, 0xa2, 0x66, 0xe2, 0xd5, 0x66, 0xf5, 0x6d, 0x0b, 0x77, 0x07,
	0x55, 0x51, 0xe1, 0x2f, 0x4e, 0x7f, 0x4c, 0xc8, 0x8b, 0xdd, 0xc5, 0xa5, 0xf7, 0x2f, 0x66, 0x5e,
	0x71, 0xf3, 0x06, 0xdb, 0x57, 0x1b, 0x35, 0x64, 0xdb, 0x48, 0xc6, 0x5a, 0xb2, 0x76, 0x95, 0x82,
	0x52, 0x96, 0xae, 0x22, 0xa7, 0x7f, 0x26, 0x84, 0xfe, 0x53, 0x06, 0x74, 0xa7, 0x93, 0xe2, 0xf2,
	0x93, 0x30, 0x18, 0xfe, 0xbb, 0x61, 0xc3, 0xf3, 0x35, 0xf2, 0x7c, 0xb5, 0x30, 0x70, 0x4b, 0x14,
	0x6e, 0xa9, 0x2b, 0xa1, 0x77, 0x82, 0xcd, 0xe5, 0x73, 0xd0, 0xf5, 0x02, 0x8f, 0xf9, 0xa2, 0x88,
	0xaa, 0x0c, 0xca, 0x3e, 0x5f, 0x23, 0xe1, 0xd1, 0x34, 0x1c, 0x53, 0x4e, 0x35, 0xd9, 0xa8, 0x05,
	0x41, 0xef, 0x76, 0x80, 0xba, 0x02, 0x1c, 0xbc, 0xba, 0xfa, 0x63, 0x43, 0xb8, 0x8b, 0x84, 0x0f,
	0xea, 0x47, 0xd7, 0xfe, 0x47, 0x9b, 0x59, 0x53, 0xc4, 0x61, 0x06, 0xb8, 0x1d, 0x57, 0x0f, 0x93,
	0x0f, 0xa2, 0xe9, 0xe3, 0x38, 0x54, 0x39, 0x85, 0xc8, 0x7f, 0x10, 0x4e, 0x74, 0x53, 0x55, 0x98,
	0x7b, 0xdc, 0x4b, 0xc7, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xea, 0x71, 0x8f, 0x34, 0xc4, 0x07,
	0x00, 0x00,
}
