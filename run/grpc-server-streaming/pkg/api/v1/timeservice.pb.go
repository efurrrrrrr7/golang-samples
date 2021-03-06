// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timeservice.proto

package timeservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	DurationSecs         uint32   `protobuf:"varint,2,opt,name=duration_secs,json=durationSecs,proto3" json:"duration_secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d50486e4ed4c3, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDurationSecs() uint32 {
	if m != nil {
		return m.DurationSecs
	}
	return 0
}

type TimeResponse struct {
	CurrentTime          *timestamp.Timestamp `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeResponse) Reset()         { *m = TimeResponse{} }
func (m *TimeResponse) String() string { return proto.CompactTextString(m) }
func (*TimeResponse) ProtoMessage()    {}
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d50486e4ed4c3, []int{1}
}

func (m *TimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeResponse.Unmarshal(m, b)
}
func (m *TimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeResponse.Marshal(b, m, deterministic)
}
func (m *TimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeResponse.Merge(m, src)
}
func (m *TimeResponse) XXX_Size() int {
	return xxx_messageInfo_TimeResponse.Size(m)
}
func (m *TimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimeResponse proto.InternalMessageInfo

func (m *TimeResponse) GetCurrentTime() *timestamp.Timestamp {
	if m != nil {
		return m.CurrentTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "timeservice.Request")
	proto.RegisterType((*TimeResponse)(nil), "timeservice.TimeResponse")
}

func init() { proto.RegisterFile("timeservice.proto", fileDescriptor_c24d50486e4ed4c3) }

var fileDescriptor_c24d50486e4ed4c3 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3d, 0x4f, 0x86, 0x30,
	0x14, 0x85, 0xad, 0x83, 0x26, 0xb7, 0x30, 0xd8, 0x38, 0x20, 0x8b, 0x04, 0x17, 0xa6, 0x62, 0x70,
	0x76, 0xf0, 0x07, 0x98, 0x98, 0xc2, 0x4e, 0xa0, 0x5e, 0x49, 0x13, 0x4b, 0xb1, 0x1f, 0xfe, 0xfe,
	0x37, 0xe5, 0x23, 0xe1, 0x9d, 0x9a, 0x3c, 0xe7, 0xa4, 0xcf, 0xb9, 0xf0, 0xe0, 0x95, 0x46, 0x87,
	0xf6, 0x5f, 0x49, 0xe4, 0x8b, 0x35, 0xde, 0x30, 0x7a, 0x42, 0xf9, 0xf3, 0x64, 0xcc, 0xf4, 0x8b,
	0xf5, 0x1a, 0x8d, 0xe1, 0xa7, 0x5e, 0x43, 0x3f, 0xe8, 0x65, 0x6b, 0x97, 0x1c, 0xee, 0x05, 0xfe,
	0x05, 0x74, 0x9e, 0xbd, 0x40, 0xfa, 0x1d, 0xec, 0xe0, 0x95, 0x99, 0x7b, 0x87, 0xd2, 0x65, 0xb7,
	0x05, 0xa9, 0x52, 0x91, 0x1c, 0xb0, 0x45, 0xe9, 0xca, 0x4f, 0x48, 0x3a, 0xa5, 0x51, 0xa0, 0x5b,
	0xcc, 0xec, 0x90, 0xbd, 0x43, 0x22, 0x83, 0xb5, 0x38, 0xfb, 0x3e, 0x7e, 0x9d, 0x91, 0x82, 0x54,
	0xb4, 0xc9, 0xf9, 0xe6, 0xe5, 0x87, 0x97, 0x77, 0x87, 0x57, 0xd0, 0xbd, 0x1f, 0x49, 0xf3, 0x05,
	0x34, 0xbe, 0xed, 0x36, 0x97, 0x7d, 0x00, 0xb4, 0xde, 0xe2, 0xa0, 0x23, 0x64, 0x8f, 0xfc, 0x7c,
	0xdd, 0x3e, 0x33, 0x7f, 0xba, 0xa2, 0xe7, 0x31, 0xe5, 0xcd, 0x2b, 0x19, 0xef, 0x56, 0xe5, 0xdb,
	0x25, 0x00, 0x00, 0xff, 0xff, 0x54, 0x1e, 0x75, 0xcb, 0x1a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeServiceClient interface {
	StreamTime(ctx context.Context, in *Request, opts ...grpc.CallOption) (TimeService_StreamTimeClient, error)
}

type timeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTimeServiceClient(cc *grpc.ClientConn) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) StreamTime(ctx context.Context, in *Request, opts ...grpc.CallOption) (TimeService_StreamTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TimeService_serviceDesc.Streams[0], "/timeservice.TimeService/StreamTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeServiceStreamTimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeService_StreamTimeClient interface {
	Recv() (*TimeResponse, error)
	grpc.ClientStream
}

type timeServiceStreamTimeClient struct {
	grpc.ClientStream
}

func (x *timeServiceStreamTimeClient) Recv() (*TimeResponse, error) {
	m := new(TimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeServiceServer is the server API for TimeService service.
type TimeServiceServer interface {
	StreamTime(*Request, TimeService_StreamTimeServer) error
}

// UnimplementedTimeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTimeServiceServer struct {
}

func (*UnimplementedTimeServiceServer) StreamTime(req *Request, srv TimeService_StreamTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTime not implemented")
}

func RegisterTimeServiceServer(s *grpc.Server, srv TimeServiceServer) {
	s.RegisterService(&_TimeService_serviceDesc, srv)
}

func _TimeService_StreamTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeServiceServer).StreamTime(m, &timeServiceStreamTimeServer{stream})
}

type TimeService_StreamTimeServer interface {
	Send(*TimeResponse) error
	grpc.ServerStream
}

type timeServiceStreamTimeServer struct {
	grpc.ServerStream
}

func (x *timeServiceStreamTimeServer) Send(m *TimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timeservice.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTime",
			Handler:       _TimeService_StreamTime_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timeservice.proto",
}
