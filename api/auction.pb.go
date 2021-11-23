// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: auction.proto

package api

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

type BidReply_Outcome int32

const (
	BidReply_SUCCESS   BidReply_Outcome = 0
	BidReply_FAIL      BidReply_Outcome = 1
	BidReply_EXCEPTION BidReply_Outcome = 2
)

// Enum value maps for BidReply_Outcome.
var (
	BidReply_Outcome_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
		2: "EXCEPTION",
	}
	BidReply_Outcome_value = map[string]int32{
		"SUCCESS":   0,
		"FAIL":      1,
		"EXCEPTION": 2,
	}
)

func (x BidReply_Outcome) Enum() *BidReply_Outcome {
	p := new(BidReply_Outcome)
	*p = x
	return p
}

func (x BidReply_Outcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BidReply_Outcome) Descriptor() protoreflect.EnumDescriptor {
	return file_auction_proto_enumTypes[0].Descriptor()
}

func (BidReply_Outcome) Type() protoreflect.EnumType {
	return &file_auction_proto_enumTypes[0]
}

func (x BidReply_Outcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BidReply_Outcome.Descriptor instead.
func (BidReply_Outcome) EnumDescriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1, 0}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int32 `protobuf:"varint,1,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type BidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outcome BidReply_Outcome `protobuf:"varint,1,opt,name=outcome,proto3,enum=BidReply_Outcome" json:"outcome,omitempty"`
}

func (x *BidReply) Reset() {
	*x = BidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidReply) ProtoMessage() {}

func (x *BidReply) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidReply.ProtoReflect.Descriptor instead.
func (*BidReply) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BidReply) GetOutcome() BidReply_Outcome {
	if x != nil {
		return x.Outcome
	}
	return BidReply_SUCCESS
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{2}
}

type ResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultReply) Reset() {
	*x = ResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReply) ProtoMessage() {}

func (x *ResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReply.ProtoReflect.Descriptor instead.
func (*ResultReply) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{3}
}

func (x *ResultReply) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_auction_proto protoreflect.FileDescriptor

var file_auction_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22,
	0x68, 0x0a, 0x08, 0x42, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52,
	0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58,
	0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x57, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x03,
	0x42, 0x69, 0x64, 0x12, 0x0b, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x09, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x2f, 0x44, 0x4d, 0x50, 0x33,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_proto_rawDescOnce sync.Once
	file_auction_proto_rawDescData = file_auction_proto_rawDesc
)

func file_auction_proto_rawDescGZIP() []byte {
	file_auction_proto_rawDescOnce.Do(func() {
		file_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_proto_rawDescData)
	})
	return file_auction_proto_rawDescData
}

var file_auction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auction_proto_goTypes = []interface{}{
	(BidReply_Outcome)(0), // 0: BidReply.Outcome
	(*BidRequest)(nil),    // 1: BidRequest
	(*BidReply)(nil),      // 2: BidReply
	(*ResultRequest)(nil), // 3: ResultRequest
	(*ResultReply)(nil),   // 4: ResultReply
}
var file_auction_proto_depIdxs = []int32{
	0, // 0: BidReply.outcome:type_name -> BidReply.Outcome
	1, // 1: Auction.Bid:input_type -> BidRequest
	3, // 2: Auction.GetResult:input_type -> ResultRequest
	2, // 3: Auction.Bid:output_type -> BidReply
	4, // 4: Auction.GetResult:output_type -> ResultReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auction_proto_init() }
func file_auction_proto_init() {
	if File_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidReply); i {
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
		file_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRequest); i {
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
		file_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReply); i {
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
			RawDescriptor: file_auction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_proto_goTypes,
		DependencyIndexes: file_auction_proto_depIdxs,
		EnumInfos:         file_auction_proto_enumTypes,
		MessageInfos:      file_auction_proto_msgTypes,
	}.Build()
	File_auction_proto = out.File
	file_auction_proto_rawDesc = nil
	file_auction_proto_goTypes = nil
	file_auction_proto_depIdxs = nil
}