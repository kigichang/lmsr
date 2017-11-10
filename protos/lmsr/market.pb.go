// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lmsr/market.proto

package lmsr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Share struct {
	Id      string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Market  string  `protobuf:"bytes,2,opt,name=market" json:"market,omitempty"`
	Outcome string  `protobuf:"bytes,3,opt,name=outcome" json:"outcome,omitempty"`
	Volume  float64 `protobuf:"fixed64,4,opt,name=volume" json:"volume,omitempty"`
}

func (m *Share) Reset()                    { *m = Share{} }
func (m *Share) String() string            { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()               {}
func (*Share) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Share) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Share) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *Share) GetOutcome() string {
	if m != nil {
		return m.Outcome
	}
	return ""
}

func (m *Share) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type Price struct {
	Share string  `protobuf:"bytes,1,opt,name=share" json:"share,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Price) GetShare() string {
	if m != nil {
		return m.Share
	}
	return ""
}

func (m *Price) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Market struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User      string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Event     string   `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
	Liquidity float64  `protobuf:"fixed64,4,opt,name=liquidity" json:"liquidity,omitempty"`
	Fund      float64  `protobuf:"fixed64,5,opt,name=fund" json:"fund,omitempty"`
	Cost      float64  `protobuf:"fixed64,6,opt,name=cost" json:"cost,omitempty"`
	Shares    []*Share `protobuf:"bytes,7,rep,name=shares" json:"shares,omitempty"`
	Prices    []*Price `protobuf:"bytes,8,rep,name=prices" json:"prices,omitempty"`
	Settled   bool     `protobuf:"varint,9,opt,name=settled" json:"settled,omitempty"`
}

func (m *Market) Reset()                    { *m = Market{} }
func (m *Market) String() string            { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()               {}
func (*Market) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Market) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Market) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Market) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Market) GetLiquidity() float64 {
	if m != nil {
		return m.Liquidity
	}
	return 0
}

func (m *Market) GetFund() float64 {
	if m != nil {
		return m.Fund
	}
	return 0
}

func (m *Market) GetCost() float64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *Market) GetShares() []*Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *Market) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *Market) GetSettled() bool {
	if m != nil {
		return m.Settled
	}
	return false
}

func init() {
	proto.RegisterType((*Share)(nil), "lmsr.Share")
	proto.RegisterType((*Price)(nil), "lmsr.Price")
	proto.RegisterType((*Market)(nil), "lmsr.Market")
}

func init() { proto.RegisterFile("lmsr/market.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x49, 0xb7, 0xed, 0x6e, 0x67, 0x41, 0x30, 0x88, 0xe4, 0xe0, 0xa1, 0xd4, 0x4b, 0x4f,
	0x15, 0xdc, 0xe7, 0x10, 0x24, 0x3e, 0x41, 0x6d, 0x47, 0x0c, 0xb6, 0x9b, 0x35, 0x7f, 0x16, 0x7c,
	0x67, 0x1f, 0x42, 0x66, 0x92, 0x45, 0xf1, 0x36, 0xbf, 0xef, 0x9b, 0xce, 0x7c, 0x9d, 0xc0, 0xf5,
	0xb2, 0x7a, 0xf7, 0xb0, 0x8e, 0xee, 0x03, 0xc3, 0x70, 0x72, 0x36, 0x58, 0x59, 0x92, 0xd4, 0x8d,
	0x50, 0xbd, 0xbc, 0x8f, 0x0e, 0xe5, 0x15, 0x14, 0x66, 0x56, 0xa2, 0x15, 0x7d, 0xa3, 0x0b, 0x33,
	0xcb, 0x5b, 0xa8, 0x53, 0xbb, 0x2a, 0x58, 0xcb, 0x24, 0x15, 0x6c, 0x6d, 0x0c, 0x93, 0x5d, 0x51,
	0x6d, 0xd8, 0xb8, 0x20, 0x7d, 0x71, 0xb6, 0x4b, 0x5c, 0x51, 0x95, 0xad, 0xe8, 0x85, 0xce, 0xd4,
	0x1d, 0xa0, 0x7a, 0x76, 0x66, 0x42, 0x79, 0x03, 0x95, 0xa7, 0x5d, 0x79, 0x4b, 0x02, 0x52, 0x4f,
	0x64, 0xf3, 0x1e, 0xa1, 0x13, 0x74, 0xdf, 0x02, 0xea, 0xa7, 0xb4, 0xf1, 0x7f, 0x32, 0x09, 0x65,
	0xf4, 0xe8, 0x72, 0x2e, 0xae, 0x69, 0x08, 0x9e, 0xf1, 0x18, 0x72, 0xa6, 0x04, 0xf2, 0x0e, 0x9a,
	0xc5, 0x7c, 0x46, 0x33, 0x9b, 0xf0, 0x95, 0x43, 0xfd, 0x0a, 0x34, 0xe7, 0x2d, 0x1e, 0x67, 0x55,
	0xb1, 0xc1, 0x35, 0x69, 0x93, 0xf5, 0x41, 0xd5, 0x49, 0xa3, 0x5a, 0xde, 0x43, 0xcd, 0x49, 0xbd,
	0xda, 0xb6, 0x9b, 0x7e, 0xff, 0xb8, 0x1f, 0xe8, 0x72, 0x03, 0x9f, 0x4d, 0x67, 0x8b, 0x9a, 0x38,
	0xb8, 0x57, 0xbb, 0xbf, 0x4d, 0xfc, 0xe3, 0x3a, 0x5b, 0x74, 0x3b, 0x8f, 0x21, 0x2c, 0x38, 0xab,
	0xa6, 0x15, 0xfd, 0x4e, 0x5f, 0xf0, 0xb5, 0xe6, 0x37, 0x39, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x29, 0x07, 0x7b, 0xa8, 0x01, 0x00, 0x00,
}