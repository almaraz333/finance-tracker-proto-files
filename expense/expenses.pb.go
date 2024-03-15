// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: expense/expenses.proto

package finance_tracker_proto_files

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

type CreateExenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt string  `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Amount    float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Category  string  `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateExenseRequest) Reset() {
	*x = CreateExenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_expenses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExenseRequest) ProtoMessage() {}

func (x *CreateExenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expense_expenses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExenseRequest.ProtoReflect.Descriptor instead.
func (*CreateExenseRequest) Descriptor() ([]byte, []int) {
	return file_expense_expenses_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExenseRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateExenseRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateExenseRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type CreateExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateExpenseResponse) Reset() {
	*x = CreateExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expense_expenses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExpenseResponse) ProtoMessage() {}

func (x *CreateExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expense_expenses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExpenseResponse.ProtoReflect.Descriptor instead.
func (*CreateExpenseResponse) Descriptor() ([]byte, []int) {
	return file_expense_expenses_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExpenseResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_expense_expenses_proto protoreflect.FileDescriptor

var file_expense_expenses_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0x4a, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x6d, 0x61, 0x72, 0x61, 0x7a, 0x33, 0x33, 0x33, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_expense_expenses_proto_rawDescOnce sync.Once
	file_expense_expenses_proto_rawDescData = file_expense_expenses_proto_rawDesc
)

func file_expense_expenses_proto_rawDescGZIP() []byte {
	file_expense_expenses_proto_rawDescOnce.Do(func() {
		file_expense_expenses_proto_rawDescData = protoimpl.X.CompressGZIP(file_expense_expenses_proto_rawDescData)
	})
	return file_expense_expenses_proto_rawDescData
}

var file_expense_expenses_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_expense_expenses_proto_goTypes = []interface{}{
	(*CreateExenseRequest)(nil),   // 0: CreateExenseRequest
	(*CreateExpenseResponse)(nil), // 1: CreateExpenseResponse
}
var file_expense_expenses_proto_depIdxs = []int32{
	0, // 0: Expense.CreateExpense:input_type -> CreateExenseRequest
	1, // 1: Expense.CreateExpense:output_type -> CreateExpenseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_expense_expenses_proto_init() }
func file_expense_expenses_proto_init() {
	if File_expense_expenses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_expense_expenses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExenseRequest); i {
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
		file_expense_expenses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExpenseResponse); i {
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
			RawDescriptor: file_expense_expenses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_expense_expenses_proto_goTypes,
		DependencyIndexes: file_expense_expenses_proto_depIdxs,
		MessageInfos:      file_expense_expenses_proto_msgTypes,
	}.Build()
	File_expense_expenses_proto = out.File
	file_expense_expenses_proto_rawDesc = nil
	file_expense_expenses_proto_goTypes = nil
	file_expense_expenses_proto_depIdxs = nil
}
