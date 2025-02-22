// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.2
// source: proto/bybit.proto

package proto

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

type GetAllAssetWithNetworksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllAssetWithNetworksRequest) Reset() {
	*x = GetAllAssetWithNetworksRequest{}
	mi := &file_proto_bybit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllAssetWithNetworksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAssetWithNetworksRequest) ProtoMessage() {}

func (x *GetAllAssetWithNetworksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bybit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAssetWithNetworksRequest.ProtoReflect.Descriptor instead.
func (*GetAllAssetWithNetworksRequest) Descriptor() ([]byte, []int) {
	return file_proto_bybit_proto_rawDescGZIP(), []int{0}
}

type GetAllAssetWithNetworksRespose struct {
	state         protoimpl.MessageState                 `protogen:"open.v1"`
	Rows          []*GetAllAssetWithNetworksRespose_Rows `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllAssetWithNetworksRespose) Reset() {
	*x = GetAllAssetWithNetworksRespose{}
	mi := &file_proto_bybit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllAssetWithNetworksRespose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAssetWithNetworksRespose) ProtoMessage() {}

func (x *GetAllAssetWithNetworksRespose) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bybit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAssetWithNetworksRespose.ProtoReflect.Descriptor instead.
func (*GetAllAssetWithNetworksRespose) Descriptor() ([]byte, []int) {
	return file_proto_bybit_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAssetWithNetworksRespose) GetRows() []*GetAllAssetWithNetworksRespose_Rows {
	if x != nil {
		return x.Rows
	}
	return nil
}

type GetAllAssetWithNetworksRespose_Chains struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Chain                 string                 `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	ChainType             string                 `protobuf:"bytes,2,opt,name=chain_type,json=chainType,proto3" json:"chain_type,omitempty"`
	Confirmation          string                 `protobuf:"bytes,3,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
	WithdrawFee           string                 `protobuf:"bytes,4,opt,name=withdraw_fee,json=withdrawFee,proto3" json:"withdraw_fee,omitempty"`
	DepositMin            string                 `protobuf:"bytes,5,opt,name=deposit_min,json=depositMin,proto3" json:"deposit_min,omitempty"`
	WithdrawMin           string                 `protobuf:"bytes,6,opt,name=withdraw_min,json=withdrawMin,proto3" json:"withdraw_min,omitempty"`
	MinAccuracy           string                 `protobuf:"bytes,7,opt,name=min_accuracy,json=minAccuracy,proto3" json:"min_accuracy,omitempty"`
	ChainDeposit          string                 `protobuf:"bytes,8,opt,name=chain_deposit,json=chainDeposit,proto3" json:"chain_deposit,omitempty"`
	ChainWithdraw         string                 `protobuf:"bytes,9,opt,name=chain_withdraw,json=chainWithdraw,proto3" json:"chain_withdraw,omitempty"`
	WithdrawPercentageFee string                 `protobuf:"bytes,10,opt,name=withdraw_percentage_fee,json=withdrawPercentageFee,proto3" json:"withdraw_percentage_fee,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *GetAllAssetWithNetworksRespose_Chains) Reset() {
	*x = GetAllAssetWithNetworksRespose_Chains{}
	mi := &file_proto_bybit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllAssetWithNetworksRespose_Chains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAssetWithNetworksRespose_Chains) ProtoMessage() {}

func (x *GetAllAssetWithNetworksRespose_Chains) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bybit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAssetWithNetworksRespose_Chains.ProtoReflect.Descriptor instead.
func (*GetAllAssetWithNetworksRespose_Chains) Descriptor() ([]byte, []int) {
	return file_proto_bybit_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetConfirmation() string {
	if x != nil {
		return x.Confirmation
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetWithdrawFee() string {
	if x != nil {
		return x.WithdrawFee
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetDepositMin() string {
	if x != nil {
		return x.DepositMin
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetWithdrawMin() string {
	if x != nil {
		return x.WithdrawMin
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetMinAccuracy() string {
	if x != nil {
		return x.MinAccuracy
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetChainDeposit() string {
	if x != nil {
		return x.ChainDeposit
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetChainWithdraw() string {
	if x != nil {
		return x.ChainWithdraw
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Chains) GetWithdrawPercentageFee() string {
	if x != nil {
		return x.WithdrawPercentageFee
	}
	return ""
}

type GetAllAssetWithNetworksRespose_Rows struct {
	state         protoimpl.MessageState                   `protogen:"open.v1"`
	Name          string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Coin          string                                   `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
	RemainAmount  string                                   `protobuf:"bytes,3,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`
	Chains        []*GetAllAssetWithNetworksRespose_Chains `protobuf:"bytes,4,rep,name=chains,proto3" json:"chains,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllAssetWithNetworksRespose_Rows) Reset() {
	*x = GetAllAssetWithNetworksRespose_Rows{}
	mi := &file_proto_bybit_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllAssetWithNetworksRespose_Rows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAssetWithNetworksRespose_Rows) ProtoMessage() {}

func (x *GetAllAssetWithNetworksRespose_Rows) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bybit_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAssetWithNetworksRespose_Rows.ProtoReflect.Descriptor instead.
func (*GetAllAssetWithNetworksRespose_Rows) Descriptor() ([]byte, []int) {
	return file_proto_bybit_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetAllAssetWithNetworksRespose_Rows) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Rows) GetCoin() string {
	if x != nil {
		return x.Coin
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Rows) GetRemainAmount() string {
	if x != nil {
		return x.RemainAmount
	}
	return ""
}

func (x *GetAllAssetWithNetworksRespose_Rows) GetChains() []*GetAllAssetWithNetworksRespose_Chains {
	if x != nil {
		return x.Chains
	}
	return nil
}

var File_proto_bybit_proto protoreflect.FileDescriptor

var file_proto_bybit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x79, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x79, 0x62, 0x69, 0x74, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xee, 0x04, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x62, 0x79, 0x62, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x1a,
	0xef, 0x02, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x4d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x36, 0x0a, 0x17, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x46, 0x65,
	0x65, 0x1a, 0x99, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x79, 0x62, 0x69, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x32, 0x72, 0x0a,
	0x05, 0x42, 0x79, 0x62, 0x69, 0x74, 0x12, 0x69, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x12, 0x25, 0x2e, 0x62, 0x79, 0x62, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x79, 0x62, 0x69, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bybit_proto_rawDescOnce sync.Once
	file_proto_bybit_proto_rawDescData = file_proto_bybit_proto_rawDesc
)

func file_proto_bybit_proto_rawDescGZIP() []byte {
	file_proto_bybit_proto_rawDescOnce.Do(func() {
		file_proto_bybit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bybit_proto_rawDescData)
	})
	return file_proto_bybit_proto_rawDescData
}

var file_proto_bybit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_bybit_proto_goTypes = []any{
	(*GetAllAssetWithNetworksRequest)(nil),        // 0: bybit.GetAllAssetWithNetworksRequest
	(*GetAllAssetWithNetworksRespose)(nil),        // 1: bybit.GetAllAssetWithNetworksRespose
	(*GetAllAssetWithNetworksRespose_Chains)(nil), // 2: bybit.GetAllAssetWithNetworksRespose.Chains
	(*GetAllAssetWithNetworksRespose_Rows)(nil),   // 3: bybit.GetAllAssetWithNetworksRespose.Rows
}
var file_proto_bybit_proto_depIdxs = []int32{
	3, // 0: bybit.GetAllAssetWithNetworksRespose.rows:type_name -> bybit.GetAllAssetWithNetworksRespose.Rows
	2, // 1: bybit.GetAllAssetWithNetworksRespose.Rows.chains:type_name -> bybit.GetAllAssetWithNetworksRespose.Chains
	0, // 2: bybit.Bybit.GetAllAssetWithNetworks:input_type -> bybit.GetAllAssetWithNetworksRequest
	1, // 3: bybit.Bybit.GetAllAssetWithNetworks:output_type -> bybit.GetAllAssetWithNetworksRespose
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_bybit_proto_init() }
func file_proto_bybit_proto_init() {
	if File_proto_bybit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bybit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bybit_proto_goTypes,
		DependencyIndexes: file_proto_bybit_proto_depIdxs,
		MessageInfos:      file_proto_bybit_proto_msgTypes,
	}.Build()
	File_proto_bybit_proto = out.File
	file_proto_bybit_proto_rawDesc = nil
	file_proto_bybit_proto_goTypes = nil
	file_proto_bybit_proto_depIdxs = nil
}
