// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.12
// source: strategy.proto

package statistico

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Strategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Market        MarketEnum             `protobuf:"varint,2,opt,name=market,proto3,enum=statistico.MarketEnum" json:"market,omitempty"`
	MinOdds       *wrapperspb.FloatValue `protobuf:"bytes,3,opt,name=min_odds,json=minOdds,proto3" json:"min_odds,omitempty"`
	MaxOdds       *wrapperspb.FloatValue `protobuf:"bytes,4,opt,name=max_odds,json=maxOdds,proto3" json:"max_odds,omitempty"`
	CompetitionId uint64                 `protobuf:"varint,5,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	SeasonId      uint64                 `protobuf:"varint,6,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	Model         string                 `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
	StakingPlan   *StakingPlan           `protobuf:"bytes,8,opt,name=staking_plan,json=stakingPlan,proto3" json:"staking_plan,omitempty"`
	Status        StrategyStatusEnum     `protobuf:"varint,9,opt,name=status,proto3,enum=statistico.StrategyStatusEnum" json:"status,omitempty"`
	Type          StrategyTypeEnum       `protobuf:"varint,10,opt,name=type,proto3,enum=statistico.StrategyTypeEnum" json:"type,omitempty"`
	StartingFund  float32                `protobuf:"fixed32,11,opt,name=starting_fund,json=startingFund,proto3" json:"starting_fund,omitempty"`
	CreatedAt     uint64                 `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     uint64                 `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Strategy) Reset() {
	*x = Strategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strategy) ProtoMessage() {}

func (x *Strategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strategy.ProtoReflect.Descriptor instead.
func (*Strategy) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *Strategy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Strategy) GetMarket() MarketEnum {
	if x != nil {
		return x.Market
	}
	return MarketEnum_OVER_UNDER_05
}

func (x *Strategy) GetMinOdds() *wrapperspb.FloatValue {
	if x != nil {
		return x.MinOdds
	}
	return nil
}

func (x *Strategy) GetMaxOdds() *wrapperspb.FloatValue {
	if x != nil {
		return x.MaxOdds
	}
	return nil
}

func (x *Strategy) GetCompetitionId() uint64 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

func (x *Strategy) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *Strategy) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Strategy) GetStakingPlan() *StakingPlan {
	if x != nil {
		return x.StakingPlan
	}
	return nil
}

func (x *Strategy) GetStatus() StrategyStatusEnum {
	if x != nil {
		return x.Status
	}
	return StrategyStatusEnum_ACTIVE
}

func (x *Strategy) GetType() StrategyTypeEnum {
	if x != nil {
		return x.Type
	}
	return StrategyTypeEnum_LIVE
}

func (x *Strategy) GetStartingFund() float32 {
	if x != nil {
		return x.StartingFund
	}
	return 0
}

func (x *Strategy) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Strategy) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_strategy_proto protoreflect.FileDescriptor

var file_strategy_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x1a, 0x0a, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x6d, 0x69, 0x6e, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x4f, 0x64, 0x64,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75,
	0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x32, 0xfc, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strategy_proto_rawDescOnce sync.Once
	file_strategy_proto_rawDescData = file_strategy_proto_rawDesc
)

func file_strategy_proto_rawDescGZIP() []byte {
	file_strategy_proto_rawDescOnce.Do(func() {
		file_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_proto_rawDescData)
	})
	return file_strategy_proto_rawDescData
}

var file_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_strategy_proto_goTypes = []interface{}{
	(*Strategy)(nil),              // 0: statistico.Strategy
	(MarketEnum)(0),               // 1: statistico.MarketEnum
	(*wrapperspb.FloatValue)(nil), // 2: google.protobuf.FloatValue
	(*StakingPlan)(nil),           // 3: statistico.StakingPlan
	(StrategyStatusEnum)(0),       // 4: statistico.StrategyStatusEnum
	(StrategyTypeEnum)(0),         // 5: statistico.StrategyTypeEnum
	(*CreateStrategyRequest)(nil), // 6: statistico.CreateStrategyRequest
	(*ListStrategiesRequest)(nil), // 7: statistico.ListStrategiesRequest
	(*UpdateStrategyRequest)(nil), // 8: statistico.UpdateStrategyRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_strategy_proto_depIdxs = []int32{
	1, // 0: statistico.Strategy.market:type_name -> statistico.MarketEnum
	2, // 1: statistico.Strategy.min_odds:type_name -> google.protobuf.FloatValue
	2, // 2: statistico.Strategy.max_odds:type_name -> google.protobuf.FloatValue
	3, // 3: statistico.Strategy.staking_plan:type_name -> statistico.StakingPlan
	4, // 4: statistico.Strategy.status:type_name -> statistico.StrategyStatusEnum
	5, // 5: statistico.Strategy.type:type_name -> statistico.StrategyTypeEnum
	6, // 6: statistico.StrategyService.CreateStrategy:input_type -> statistico.CreateStrategyRequest
	7, // 7: statistico.StrategyService.ListStrategies:input_type -> statistico.ListStrategiesRequest
	8, // 8: statistico.StrategyService.UpdateStrategy:input_type -> statistico.UpdateStrategyRequest
	0, // 9: statistico.StrategyService.CreateStrategy:output_type -> statistico.Strategy
	0, // 10: statistico.StrategyService.ListStrategies:output_type -> statistico.Strategy
	9, // 11: statistico.StrategyService.UpdateStrategy:output_type -> google.protobuf.Empty
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_strategy_proto_init() }
func file_strategy_proto_init() {
	if File_strategy_proto != nil {
		return
	}
	file_enum_proto_init()
	file_requests_proto_init()
	file_utility_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strategy); i {
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
			RawDescriptor: file_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strategy_proto_goTypes,
		DependencyIndexes: file_strategy_proto_depIdxs,
		MessageInfos:      file_strategy_proto_msgTypes,
	}.Build()
	File_strategy_proto = out.File
	file_strategy_proto_rawDesc = nil
	file_strategy_proto_goTypes = nil
	file_strategy_proto_depIdxs = nil
}