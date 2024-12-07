// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/lyrics.proto

package lyrics

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{0}
}

type LyricsBySongIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId string `protobuf:"bytes,1,opt,name=songId,proto3" json:"songId,omitempty"`
}

func (x *LyricsBySongIdRequest) Reset() {
	*x = LyricsBySongIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LyricsBySongIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LyricsBySongIdRequest) ProtoMessage() {}

func (x *LyricsBySongIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LyricsBySongIdRequest.ProtoReflect.Descriptor instead.
func (*LyricsBySongIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{1}
}

func (x *LyricsBySongIdRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

type CreateLyricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId  string `protobuf:"bytes,1,opt,name=songId,proto3" json:"songId,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateLyricsRequest) Reset() {
	*x = CreateLyricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLyricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLyricsRequest) ProtoMessage() {}

func (x *CreateLyricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLyricsRequest.ProtoReflect.Descriptor instead.
func (*CreateLyricsRequest) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLyricsRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *CreateLyricsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateLyricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId  string `protobuf:"bytes,1,opt,name=songId,proto3" json:"songId,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateLyricsRequest) Reset() {
	*x = UpdateLyricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLyricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLyricsRequest) ProtoMessage() {}

func (x *UpdateLyricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLyricsRequest.ProtoReflect.Descriptor instead.
func (*UpdateLyricsRequest) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateLyricsRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *UpdateLyricsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type LyricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SongId    string `protobuf:"bytes,2,opt,name=songId,proto3" json:"songId,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *LyricsResponse) Reset() {
	*x = LyricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LyricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LyricsResponse) ProtoMessage() {}

func (x *LyricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LyricsResponse.ProtoReflect.Descriptor instead.
func (*LyricsResponse) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{4}
}

func (x *LyricsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LyricsResponse) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *LyricsResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *LyricsResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LyricsResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type LyricsResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lyrics []*LyricsResponse `protobuf:"bytes,1,rep,name=lyrics,proto3" json:"lyrics,omitempty"`
}

func (x *LyricsResponseList) Reset() {
	*x = LyricsResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lyrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LyricsResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LyricsResponseList) ProtoMessage() {}

func (x *LyricsResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lyrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LyricsResponseList.ProtoReflect.Descriptor instead.
func (*LyricsResponseList) Descriptor() ([]byte, []int) {
	return file_protos_lyrics_proto_rawDescGZIP(), []int{5}
}

func (x *LyricsResponseList) GetLyrics() []*LyricsResponse {
	if x != nil {
		return x.Lyrics
	}
	return nil
}

var File_protos_lyrics_proto protoreflect.FileDescriptor

var file_protos_lyrics_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x15, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73,
	0x42, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x47, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x4c, 0x79,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x4c, 0x79,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73,
	0x32, 0xa0, 0x02, 0x0a, 0x0d, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x79, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x0d, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1a, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x79, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x53, 0x6f, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x79, 0x72, 0x69,
	0x63, 0x73, 0x42, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x79, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1b, 0x2e, 0x6c, 0x79, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1b,
	0x2e, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x79,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x79,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_lyrics_proto_rawDescOnce sync.Once
	file_protos_lyrics_proto_rawDescData = file_protos_lyrics_proto_rawDesc
)

func file_protos_lyrics_proto_rawDescGZIP() []byte {
	file_protos_lyrics_proto_rawDescOnce.Do(func() {
		file_protos_lyrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_lyrics_proto_rawDescData)
	})
	return file_protos_lyrics_proto_rawDescData
}

var file_protos_lyrics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_lyrics_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: lyrics.Empty
	(*LyricsBySongIdRequest)(nil), // 1: lyrics.LyricsBySongIdRequest
	(*CreateLyricsRequest)(nil),   // 2: lyrics.CreateLyricsRequest
	(*UpdateLyricsRequest)(nil),   // 3: lyrics.UpdateLyricsRequest
	(*LyricsResponse)(nil),        // 4: lyrics.LyricsResponse
	(*LyricsResponseList)(nil),    // 5: lyrics.LyricsResponseList
}
var file_protos_lyrics_proto_depIdxs = []int32{
	4, // 0: lyrics.LyricsResponseList.lyrics:type_name -> lyrics.LyricsResponse
	0, // 1: lyrics.LyricsService.GetAllLyrics:input_type -> lyrics.Empty
	1, // 2: lyrics.LyricsService.GetLyricsBySongId:input_type -> lyrics.LyricsBySongIdRequest
	2, // 3: lyrics.LyricsService.CreateLyrics:input_type -> lyrics.CreateLyricsRequest
	3, // 4: lyrics.LyricsService.UpdateLyrics:input_type -> lyrics.UpdateLyricsRequest
	5, // 5: lyrics.LyricsService.GetAllLyrics:output_type -> lyrics.LyricsResponseList
	4, // 6: lyrics.LyricsService.GetLyricsBySongId:output_type -> lyrics.LyricsResponse
	4, // 7: lyrics.LyricsService.CreateLyrics:output_type -> lyrics.LyricsResponse
	4, // 8: lyrics.LyricsService.UpdateLyrics:output_type -> lyrics.LyricsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_lyrics_proto_init() }
func file_protos_lyrics_proto_init() {
	if File_protos_lyrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_lyrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_lyrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LyricsBySongIdRequest); i {
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
		file_protos_lyrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLyricsRequest); i {
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
		file_protos_lyrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLyricsRequest); i {
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
		file_protos_lyrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LyricsResponse); i {
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
		file_protos_lyrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LyricsResponseList); i {
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
			RawDescriptor: file_protos_lyrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_lyrics_proto_goTypes,
		DependencyIndexes: file_protos_lyrics_proto_depIdxs,
		MessageInfos:      file_protos_lyrics_proto_msgTypes,
	}.Build()
	File_protos_lyrics_proto = out.File
	file_protos_lyrics_proto_rawDesc = nil
	file_protos_lyrics_proto_goTypes = nil
	file_protos_lyrics_proto_depIdxs = nil
}