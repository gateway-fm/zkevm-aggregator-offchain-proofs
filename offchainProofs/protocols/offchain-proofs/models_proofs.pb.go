// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: models_proofs.proto

package proofs_pb

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

type VerifyBatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RollupId         uint32   `protobuf:"varint,1,opt,name=rollup_id,json=rollupId,proto3" json:"rollup_id,omitempty"`
	PendingStateNum  uint64   `protobuf:"varint,2,opt,name=pending_state_num,json=pendingStateNum,proto3" json:"pending_state_num,omitempty"`
	InitNumBatch     uint64   `protobuf:"varint,3,opt,name=init_num_batch,json=initNumBatch,proto3" json:"init_num_batch,omitempty"`
	FinalNewBatch    uint64   `protobuf:"varint,4,opt,name=final_new_batch,json=finalNewBatch,proto3" json:"final_new_batch,omitempty"`
	NewLocalExitRoot []byte   `protobuf:"bytes,5,opt,name=new_local_exit_root,json=newLocalExitRoot,proto3" json:"new_local_exit_root,omitempty"`
	NewStateRoot     []byte   `protobuf:"bytes,6,opt,name=new_state_root,json=newStateRoot,proto3" json:"new_state_root,omitempty"`
	Beneficiary      string   `protobuf:"bytes,7,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Proof            [][]byte `protobuf:"bytes,8,rep,name=proof,proto3" json:"proof,omitempty"`
	Signer           string   `protobuf:"bytes,9,opt,name=signer,proto3" json:"signer,omitempty"`
	TxHash           string   `protobuf:"bytes,10,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *VerifyBatchesRequest) Reset() {
	*x = VerifyBatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proofs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBatchesRequest) ProtoMessage() {}

func (x *VerifyBatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proofs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBatchesRequest.ProtoReflect.Descriptor instead.
func (*VerifyBatchesRequest) Descriptor() ([]byte, []int) {
	return file_models_proofs_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyBatchesRequest) GetRollupId() uint32 {
	if x != nil {
		return x.RollupId
	}
	return 0
}

func (x *VerifyBatchesRequest) GetPendingStateNum() uint64 {
	if x != nil {
		return x.PendingStateNum
	}
	return 0
}

func (x *VerifyBatchesRequest) GetInitNumBatch() uint64 {
	if x != nil {
		return x.InitNumBatch
	}
	return 0
}

func (x *VerifyBatchesRequest) GetFinalNewBatch() uint64 {
	if x != nil {
		return x.FinalNewBatch
	}
	return 0
}

func (x *VerifyBatchesRequest) GetNewLocalExitRoot() []byte {
	if x != nil {
		return x.NewLocalExitRoot
	}
	return nil
}

func (x *VerifyBatchesRequest) GetNewStateRoot() []byte {
	if x != nil {
		return x.NewStateRoot
	}
	return nil
}

func (x *VerifyBatchesRequest) GetBeneficiary() string {
	if x != nil {
		return x.Beneficiary
	}
	return ""
}

func (x *VerifyBatchesRequest) GetProof() [][]byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *VerifyBatchesRequest) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *VerifyBatchesRequest) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type VerifyBatchesTrustedAggregatorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalNewBatch uint64   `protobuf:"varint,1,opt,name=final_new_batch,json=finalNewBatch,proto3" json:"final_new_batch,omitempty"`
	NewStateRoot  []byte   `protobuf:"bytes,2,opt,name=new_state_root,json=newStateRoot,proto3" json:"new_state_root,omitempty"`
	Sender        string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	BlockNumber   uint64   `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	TxHash        string   `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Topics        [][]byte `protobuf:"bytes,6,rep,name=topics,proto3" json:"topics,omitempty"`
	Data          []byte   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Address       string   `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *VerifyBatchesTrustedAggregatorEvent) Reset() {
	*x = VerifyBatchesTrustedAggregatorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proofs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBatchesTrustedAggregatorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBatchesTrustedAggregatorEvent) ProtoMessage() {}

func (x *VerifyBatchesTrustedAggregatorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_models_proofs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBatchesTrustedAggregatorEvent.ProtoReflect.Descriptor instead.
func (*VerifyBatchesTrustedAggregatorEvent) Descriptor() ([]byte, []int) {
	return file_models_proofs_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetFinalNewBatch() uint64 {
	if x != nil {
		return x.FinalNewBatch
	}
	return 0
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetNewStateRoot() []byte {
	if x != nil {
		return x.NewStateRoot
	}
	return nil
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetTopics() [][]byte {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *VerifyBatchesTrustedAggregatorEvent) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type VerifyBatchesTrustedAggregatorEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*VerifyBatchesTrustedAggregatorEvent `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *VerifyBatchesTrustedAggregatorEvents) Reset() {
	*x = VerifyBatchesTrustedAggregatorEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proofs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyBatchesTrustedAggregatorEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBatchesTrustedAggregatorEvents) ProtoMessage() {}

func (x *VerifyBatchesTrustedAggregatorEvents) ProtoReflect() protoreflect.Message {
	mi := &file_models_proofs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBatchesTrustedAggregatorEvents.ProtoReflect.Descriptor instead.
func (*VerifyBatchesTrustedAggregatorEvents) Descriptor() ([]byte, []int) {
	return file_models_proofs_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyBatchesTrustedAggregatorEvents) GetValues() []*VerifyBatchesTrustedAggregatorEvent {
	if x != nil {
		return x.Values
	}
	return nil
}

type Blank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Blank) Reset() {
	*x = Blank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proofs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_models_proofs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_models_proofs_proto_rawDescGZIP(), []int{3}
}

var File_models_proofs_proto protoreflect.FileDescriptor

var file_models_proofs_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0xeb, 0x02, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0f,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x65, 0x77, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x2d, 0x0a, 0x13, 0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x69, 0x74, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6e, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x8d, 0x02, 0x0a, 0x23, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x65, 0x77, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x72, 0x0a, 0x24, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x42,
	0x1b, 0x5a, 0x19, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proofs_proto_rawDescOnce sync.Once
	file_models_proofs_proto_rawDescData = file_models_proofs_proto_rawDesc
)

func file_models_proofs_proto_rawDescGZIP() []byte {
	file_models_proofs_proto_rawDescOnce.Do(func() {
		file_models_proofs_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proofs_proto_rawDescData)
	})
	return file_models_proofs_proto_rawDescData
}

var file_models_proofs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_proofs_proto_goTypes = []interface{}{
	(*VerifyBatchesRequest)(nil),                 // 0: proofs.models.VerifyBatchesRequest
	(*VerifyBatchesTrustedAggregatorEvent)(nil),  // 1: proofs.models.VerifyBatchesTrustedAggregatorEvent
	(*VerifyBatchesTrustedAggregatorEvents)(nil), // 2: proofs.models.VerifyBatchesTrustedAggregatorEvents
	(*Blank)(nil), // 3: proofs.models.Blank
}
var file_models_proofs_proto_depIdxs = []int32{
	1, // 0: proofs.models.VerifyBatchesTrustedAggregatorEvents.values:type_name -> proofs.models.VerifyBatchesTrustedAggregatorEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_proofs_proto_init() }
func file_models_proofs_proto_init() {
	if File_models_proofs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proofs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBatchesRequest); i {
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
		file_models_proofs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBatchesTrustedAggregatorEvent); i {
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
		file_models_proofs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyBatchesTrustedAggregatorEvents); i {
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
		file_models_proofs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blank); i {
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
			RawDescriptor: file_models_proofs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proofs_proto_goTypes,
		DependencyIndexes: file_models_proofs_proto_depIdxs,
		MessageInfos:      file_models_proofs_proto_msgTypes,
	}.Build()
	File_models_proofs_proto = out.File
	file_models_proofs_proto_rawDesc = nil
	file_models_proofs_proto_goTypes = nil
	file_models_proofs_proto_depIdxs = nil
}
