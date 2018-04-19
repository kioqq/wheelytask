// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TripInformation.proto

/*
Package wheelypb is a generated protocol buffer package.

It is generated from these files:
	TripInformation.proto

It has these top-level messages:
	Point
	Trip
	TripInformationResponse
*/
package wheelypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	Lat  float64 `protobuf:"fixed64,1,opt,name=lat" json:"lat,omitempty"`
	Long float64 `protobuf:"fixed64,2,opt,name=long" json:"long,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Point) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

type Trip struct {
	Start *Point `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *Point `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *Trip) Reset()                    { *m = Trip{} }
func (m *Trip) String() string            { return proto.CompactTextString(m) }
func (*Trip) ProtoMessage()               {}
func (*Trip) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Trip) GetStart() *Point {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Trip) GetEnd() *Point {
	if m != nil {
		return m.End
	}
	return nil
}

type TripInformationResponse struct {
	Distance int64 `protobuf:"varint,1,opt,name=distance" json:"distance,omitempty"`
	Time     int64 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
}

func (m *TripInformationResponse) Reset()                    { *m = TripInformationResponse{} }
func (m *TripInformationResponse) String() string            { return proto.CompactTextString(m) }
func (*TripInformationResponse) ProtoMessage()               {}
func (*TripInformationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TripInformationResponse) GetDistance() int64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *TripInformationResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "wheelypb.Point")
	proto.RegisterType((*Trip)(nil), "wheelypb.Trip")
	proto.RegisterType((*TripInformationResponse)(nil), "wheelypb.TripInformationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TripInformation service

type TripInformationClient interface {
	GetTripInformation(ctx context.Context, in *Trip, opts ...grpc.CallOption) (*TripInformationResponse, error)
}

type tripInformationClient struct {
	cc *grpc.ClientConn
}

func NewTripInformationClient(cc *grpc.ClientConn) TripInformationClient {
	return &tripInformationClient{cc}
}

func (c *tripInformationClient) GetTripInformation(ctx context.Context, in *Trip, opts ...grpc.CallOption) (*TripInformationResponse, error) {
	out := new(TripInformationResponse)
	err := grpc.Invoke(ctx, "/wheelypb.TripInformation/GetTripInformation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TripInformation service

type TripInformationServer interface {
	GetTripInformation(context.Context, *Trip) (*TripInformationResponse, error)
}

func RegisterTripInformationServer(s *grpc.Server, srv TripInformationServer) {
	s.RegisterService(&_TripInformation_serviceDesc, srv)
}

func _TripInformation_GetTripInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripInformationServer).GetTripInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wheelypb.TripInformation/GetTripInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripInformationServer).GetTripInformation(ctx, req.(*Trip))
	}
	return interceptor(ctx, in, info, handler)
}

var _TripInformation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wheelypb.TripInformation",
	HandlerType: (*TripInformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTripInformation",
			Handler:    _TripInformation_GetTripInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TripInformation.proto",
}

func init() { proto.RegisterFile("TripInformation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0xd9, 0x95, 0x65, 0x04, 0x57, 0x06, 0xc4, 0x65, 0x4f, 0x6e, 0x40, 0xf0, 0x62,
	0x0f, 0xeb, 0x43, 0x48, 0x6f, 0x25, 0x78, 0xf4, 0x92, 0xda, 0x51, 0x03, 0x6d, 0x12, 0x92, 0x01,
	0xf1, 0xed, 0x25, 0x53, 0x54, 0xb6, 0xf4, 0xf6, 0xcf, 0xfc, 0xc3, 0x97, 0x8f, 0xc0, 0xcd, 0x4b,
	0x72, 0xb1, 0xf1, 0xef, 0x21, 0x8d, 0x96, 0x5d, 0xf0, 0x75, 0x4c, 0x81, 0x03, 0x6e, 0xbe, 0x3e,
	0x89, 0x86, 0xef, 0xd8, 0xe9, 0x47, 0x58, 0xb7, 0xc1, 0x79, 0xc6, 0x6b, 0x50, 0x83, 0xe5, 0x5d,
	0x75, 0x57, 0x3d, 0x54, 0xa6, 0x44, 0x44, 0x58, 0x0d, 0xc1, 0x7f, 0xec, 0xce, 0x65, 0x25, 0x59,
	0xb7, 0xb0, 0x2a, 0x44, 0xbc, 0x87, 0x75, 0x66, 0x9b, 0xa6, 0xfb, 0xcb, 0xe3, 0xb6, 0xfe, 0x05,
	0xd6, 0x42, 0x33, 0x53, 0x8b, 0x07, 0x50, 0xe4, 0x7b, 0x21, 0x2c, 0x1c, 0x95, 0x4e, 0x37, 0x70,
	0x3b, 0x73, 0x34, 0x94, 0x63, 0xf0, 0x99, 0x70, 0x0f, 0x9b, 0xde, 0x65, 0xb6, 0xfe, 0x8d, 0xe4,
	0x1d, 0x65, 0xfe, 0xe6, 0x22, 0xc7, 0x6e, 0x24, 0x41, 0x2b, 0x23, 0xf9, 0xf8, 0x0a, 0xdb, 0x19,
	0x0a, 0x1b, 0xc0, 0x67, 0xe2, 0xf9, 0xf6, 0xea, 0xdf, 0xa4, 0x54, 0xfb, 0xc3, 0xe9, 0xbc, 0xe0,
	0xa2, 0xcf, 0xba, 0x0b, 0xf9, 0xba, 0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xcf, 0x87,
	0xa0, 0x53, 0x01, 0x00, 0x00,
}