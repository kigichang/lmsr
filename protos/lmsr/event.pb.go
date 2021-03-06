// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lmsr/event.proto

package lmsr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Outcome struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Outcome) Reset()                    { *m = Outcome{} }
func (m *Outcome) String() string            { return proto.CompactTextString(m) }
func (*Outcome) ProtoMessage()               {}
func (*Outcome) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Outcome) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Outcome) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Event struct {
	Id       string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User     string     `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Title    string     `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Outcomes []*Outcome `protobuf:"bytes,4,rep,name=outcomes" json:"outcomes,omitempty"`
	Result   string     `protobuf:"bytes,5,opt,name=result" json:"result,omitempty"`
	Approved bool       `protobuf:"varint,6,opt,name=approved" json:"approved,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetOutcomes() []*Outcome {
	if m != nil {
		return m.Outcomes
	}
	return nil
}

func (m *Event) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Event) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

func init() {
	proto.RegisterType((*Outcome)(nil), "lmsr.Outcome")
	proto.RegisterType((*Event)(nil), "lmsr.Event")
}

func init() { proto.RegisterFile("lmsr/event.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x2d, 0x2e,
	0xd2, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89,
	0x28, 0xe9, 0x73, 0xb1, 0xfb, 0x97, 0x96, 0x24, 0xe7, 0xe7, 0xa6, 0x0a, 0xf1, 0x71, 0x31, 0x65,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x81, 0x85, 0x20, 0x1c, 0xa5, 0x05, 0x8c, 0x5c, 0xac, 0xae,
	0x20, 0x63, 0x30, 0xd4, 0x0b, 0x71, 0xb1, 0x94, 0x16, 0xa7, 0x16, 0x41, 0x95, 0x83, 0xd9, 0x08,
	0x33, 0x98, 0x91, 0xcc, 0x10, 0xd2, 0xe4, 0xe2, 0xc8, 0x87, 0x58, 0x5a, 0x2c, 0xc1, 0xa2, 0xc0,
	0xac, 0xc1, 0x6d, 0xc4, 0xab, 0x07, 0x72, 0x8d, 0x1e, 0xd4, 0x29, 0x41, 0x70, 0x69, 0x21, 0x31,
	0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x09, 0x56, 0xb0, 0x09, 0x50, 0x9e, 0x90, 0x14,
	0x17, 0x47, 0x62, 0x41, 0x41, 0x51, 0x7e, 0x59, 0x6a, 0x8a, 0x04, 0x9b, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x9c, 0x9f, 0xc4, 0x06, 0xf6, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x7d, 0x30,
	0xfe, 0xf4, 0x00, 0x00, 0x00,
}
