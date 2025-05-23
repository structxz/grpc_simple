// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.3
// source: api/notes/messages.proto

package notes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NoteInfo - information of the note
type NoteInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// title - name of the note
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// content - content of the note
	Content       string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoteInfo) Reset() {
	*x = NoteInfo{}
	mi := &file_api_notes_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteInfo) ProtoMessage() {}

func (x *NoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteInfo.ProtoReflect.Descriptor instead.
func (*NoteInfo) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{0}
}

func (x *NoteInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NoteInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Note - full note model
type Note struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id - unique id of the note
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// info - information about note
	Info          *NoteInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Note) Reset() {
	*x = Note{}
	mi := &file_api_notes_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Note) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetInfo() *NoteInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveNoteRequest - request to save the note
type SaveNoteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// info - request's content
	Info          *NoteInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveNoteRequest) Reset() {
	*x = SaveNoteRequest{}
	mi := &file_api_notes_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNoteRequest) ProtoMessage() {}

func (x *SaveNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNoteRequest.ProtoReflect.Descriptor instead.
func (*SaveNoteRequest) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{2}
}

func (x *SaveNoteRequest) GetInfo() *NoteInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveNoteResponse - response to save the note
type SaveNoteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id - unique id of the note
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveNoteResponse) Reset() {
	*x = SaveNoteResponse{}
	mi := &file_api_notes_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNoteResponse) ProtoMessage() {}

func (x *SaveNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNoteResponse.ProtoReflect.Descriptor instead.
func (*SaveNoteResponse) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{3}
}

func (x *SaveNoteResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ListNoteRequest - request to get all notes
type ListNotesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesRequest) Reset() {
	*x = ListNotesRequest{}
	mi := &file_api_notes_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesRequest) ProtoMessage() {}

func (x *ListNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesRequest.ProtoReflect.Descriptor instead.
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{4}
}

// ListNoteResponse - response to get all notes
type ListNotesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Notes []*Note                `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	// Types that are valid to be assigned to Meta:
	//
	//	*ListNotesResponse_Foo_
	//	*ListNotesResponse_Bar_
	Meta          isListNotesResponse_Meta `protobuf_oneof:"meta"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesResponse) Reset() {
	*x = ListNotesResponse{}
	mi := &file_api_notes_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse) ProtoMessage() {}

func (x *ListNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse.ProtoReflect.Descriptor instead.
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *ListNotesResponse) GetMeta() isListNotesResponse_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListNotesResponse) GetFoo() *ListNotesResponse_Foo {
	if x != nil {
		if x, ok := x.Meta.(*ListNotesResponse_Foo_); ok {
			return x.Foo
		}
	}
	return nil
}

func (x *ListNotesResponse) GetBar() *ListNotesResponse_Bar {
	if x != nil {
		if x, ok := x.Meta.(*ListNotesResponse_Bar_); ok {
			return x.Bar
		}
	}
	return nil
}

type isListNotesResponse_Meta interface {
	isListNotesResponse_Meta()
}

type ListNotesResponse_Foo_ struct {
	Foo *ListNotesResponse_Foo `protobuf:"bytes,2,opt,name=foo,proto3,oneof"`
}

type ListNotesResponse_Bar_ struct {
	Bar *ListNotesResponse_Bar `protobuf:"bytes,3,opt,name=bar,proto3,oneof"`
}

func (*ListNotesResponse_Foo_) isListNotesResponse_Meta() {}

func (*ListNotesResponse_Bar_) isListNotesResponse_Meta() {}

type ListNotesResponse_Foo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Int           int64                  `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesResponse_Foo) Reset() {
	*x = ListNotesResponse_Foo{}
	mi := &file_api_notes_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesResponse_Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse_Foo) ProtoMessage() {}

func (x *ListNotesResponse_Foo) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse_Foo.ProtoReflect.Descriptor instead.
func (*ListNotesResponse_Foo) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListNotesResponse_Foo) GetInt() int64 {
	if x != nil {
		return x.Int
	}
	return 0
}

type ListNotesResponse_Bar struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Int           int32                  `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotesResponse_Bar) Reset() {
	*x = ListNotesResponse_Bar{}
	mi := &file_api_notes_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotesResponse_Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse_Bar) ProtoMessage() {}

func (x *ListNotesResponse_Bar) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse_Bar.ProtoReflect.Descriptor instead.
func (*ListNotesResponse_Bar) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5, 1}
}

func (x *ListNotesResponse_Bar) GetInt() int32 {
	if x != nil {
		return x.Int
	}
	return 0
}

var File_api_notes_messages_proto protoreflect.FileDescriptor

const file_api_notes_messages_proto_rawDesc = "" +
	"\n" +
	"\x18api/notes/messages.proto\x12\x1fgithub.com.structxz.grpc_simple\":\n" +
	"\bNoteInfo\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\"U\n" +
	"\x04Note\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12=\n" +
	"\x04info\x18\x02 \x01(\v2).github.com.structxz.grpc_simple.NoteInfoR\x04info\"P\n" +
	"\x0fSaveNoteRequest\x12=\n" +
	"\x04info\x18\x01 \x01(\v2).github.com.structxz.grpc_simple.NoteInfoR\x04info\"\"\n" +
	"\x10SaveNoteResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"\x12\n" +
	"\x10ListNotesRequest\"\xa2\x02\n" +
	"\x11ListNotesResponse\x12;\n" +
	"\x05notes\x18\x01 \x03(\v2%.github.com.structxz.grpc_simple.NoteR\x05notes\x12J\n" +
	"\x03foo\x18\x02 \x01(\v26.github.com.structxz.grpc_simple.ListNotesResponse.FooH\x00R\x03foo\x12J\n" +
	"\x03bar\x18\x03 \x01(\v26.github.com.structxz.grpc_simple.ListNotesResponse.BarH\x00R\x03bar\x1a\x17\n" +
	"\x03Foo\x12\x10\n" +
	"\x03int\x18\x01 \x01(\x03R\x03int\x1a\x17\n" +
	"\x03Bar\x12\x10\n" +
	"\x03int\x18\x01 \x01(\x05R\x03intB\x06\n" +
	"\x04metaB5Z3github.com/structxz/grpc_simple/pkg/api/notes;notesb\x06proto3"

var (
	file_api_notes_messages_proto_rawDescOnce sync.Once
	file_api_notes_messages_proto_rawDescData []byte
)

func file_api_notes_messages_proto_rawDescGZIP() []byte {
	file_api_notes_messages_proto_rawDescOnce.Do(func() {
		file_api_notes_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_notes_messages_proto_rawDesc), len(file_api_notes_messages_proto_rawDesc)))
	})
	return file_api_notes_messages_proto_rawDescData
}

var file_api_notes_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_notes_messages_proto_goTypes = []any{
	(*NoteInfo)(nil),              // 0: github.com.structxz.grpc_simple.NoteInfo
	(*Note)(nil),                  // 1: github.com.structxz.grpc_simple.Note
	(*SaveNoteRequest)(nil),       // 2: github.com.structxz.grpc_simple.SaveNoteRequest
	(*SaveNoteResponse)(nil),      // 3: github.com.structxz.grpc_simple.SaveNoteResponse
	(*ListNotesRequest)(nil),      // 4: github.com.structxz.grpc_simple.ListNotesRequest
	(*ListNotesResponse)(nil),     // 5: github.com.structxz.grpc_simple.ListNotesResponse
	(*ListNotesResponse_Foo)(nil), // 6: github.com.structxz.grpc_simple.ListNotesResponse.Foo
	(*ListNotesResponse_Bar)(nil), // 7: github.com.structxz.grpc_simple.ListNotesResponse.Bar
}
var file_api_notes_messages_proto_depIdxs = []int32{
	0, // 0: github.com.structxz.grpc_simple.Note.info:type_name -> github.com.structxz.grpc_simple.NoteInfo
	0, // 1: github.com.structxz.grpc_simple.SaveNoteRequest.info:type_name -> github.com.structxz.grpc_simple.NoteInfo
	1, // 2: github.com.structxz.grpc_simple.ListNotesResponse.notes:type_name -> github.com.structxz.grpc_simple.Note
	6, // 3: github.com.structxz.grpc_simple.ListNotesResponse.foo:type_name -> github.com.structxz.grpc_simple.ListNotesResponse.Foo
	7, // 4: github.com.structxz.grpc_simple.ListNotesResponse.bar:type_name -> github.com.structxz.grpc_simple.ListNotesResponse.Bar
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_notes_messages_proto_init() }
func file_api_notes_messages_proto_init() {
	if File_api_notes_messages_proto != nil {
		return
	}
	file_api_notes_messages_proto_msgTypes[5].OneofWrappers = []any{
		(*ListNotesResponse_Foo_)(nil),
		(*ListNotesResponse_Bar_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_notes_messages_proto_rawDesc), len(file_api_notes_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_notes_messages_proto_goTypes,
		DependencyIndexes: file_api_notes_messages_proto_depIdxs,
		MessageInfos:      file_api_notes_messages_proto_msgTypes,
	}.Build()
	File_api_notes_messages_proto = out.File
	file_api_notes_messages_proto_goTypes = nil
	file_api_notes_messages_proto_depIdxs = nil
}
