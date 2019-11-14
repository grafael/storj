// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: referralmanager.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetTokensRequest struct {
	UserId               []byte   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokensRequest) Reset()         { *m = GetTokensRequest{} }
func (m *GetTokensRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokensRequest) ProtoMessage()    {}
func (*GetTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{0}
}
func (m *GetTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensRequest.Unmarshal(m, b)
}
func (m *GetTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensRequest.Marshal(b, m, deterministic)
}
func (m *GetTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensRequest.Merge(m, src)
}
func (m *GetTokensRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokensRequest.Size(m)
}
func (m *GetTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensRequest proto.InternalMessageInfo

func (m *GetTokensRequest) GetUserId() []byte {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetTokensResponse struct {
	Token                [][]byte `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokensResponse) Reset()         { *m = GetTokensResponse{} }
func (m *GetTokensResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokensResponse) ProtoMessage()    {}
func (*GetTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{1}
}
func (m *GetTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensResponse.Unmarshal(m, b)
}
func (m *GetTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensResponse.Marshal(b, m, deterministic)
}
func (m *GetTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensResponse.Merge(m, src)
}
func (m *GetTokensResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokensResponse.Size(m)
}
func (m *GetTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensResponse proto.InternalMessageInfo

func (m *GetTokensResponse) GetToken() [][]byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type ReserveTokenRequest struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RedeemingSatelliteId NodeID   `protobuf:"bytes,2,opt,name=redeeming_satellite_id,json=redeemingSatelliteId,proto3,customtype=NodeID" json:"redeeming_satellite_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveTokenRequest) Reset()         { *m = ReserveTokenRequest{} }
func (m *ReserveTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveTokenRequest) ProtoMessage()    {}
func (*ReserveTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{2}
}
func (m *ReserveTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveTokenRequest.Unmarshal(m, b)
}
func (m *ReserveTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveTokenRequest.Marshal(b, m, deterministic)
}
func (m *ReserveTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveTokenRequest.Merge(m, src)
}
func (m *ReserveTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveTokenRequest.Size(m)
}
func (m *ReserveTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveTokenRequest proto.InternalMessageInfo

func (m *ReserveTokenRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type ReserveTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveTokenResponse) Reset()         { *m = ReserveTokenResponse{} }
func (m *ReserveTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ReserveTokenResponse) ProtoMessage()    {}
func (*ReserveTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{3}
}
func (m *ReserveTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveTokenResponse.Unmarshal(m, b)
}
func (m *ReserveTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveTokenResponse.Marshal(b, m, deterministic)
}
func (m *ReserveTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveTokenResponse.Merge(m, src)
}
func (m *ReserveTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ReserveTokenResponse.Size(m)
}
func (m *ReserveTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveTokenResponse proto.InternalMessageInfo

type RedeemTokenRequest struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               []byte   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemTokenRequest) Reset()         { *m = RedeemTokenRequest{} }
func (m *RedeemTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RedeemTokenRequest) ProtoMessage()    {}
func (*RedeemTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{4}
}
func (m *RedeemTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemTokenRequest.Unmarshal(m, b)
}
func (m *RedeemTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemTokenRequest.Marshal(b, m, deterministic)
}
func (m *RedeemTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemTokenRequest.Merge(m, src)
}
func (m *RedeemTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RedeemTokenRequest.Size(m)
}
func (m *RedeemTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemTokenRequest proto.InternalMessageInfo

func (m *RedeemTokenRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RedeemTokenRequest) GetUserId() []byte {
	if m != nil {
		return m.UserId
	}
	return nil
}

type RedeemTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemTokenResponse) Reset()         { *m = RedeemTokenResponse{} }
func (m *RedeemTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RedeemTokenResponse) ProtoMessage()    {}
func (*RedeemTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{5}
}
func (m *RedeemTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemTokenResponse.Unmarshal(m, b)
}
func (m *RedeemTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemTokenResponse.Marshal(b, m, deterministic)
}
func (m *RedeemTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemTokenResponse.Merge(m, src)
}
func (m *RedeemTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RedeemTokenResponse.Size(m)
}
func (m *RedeemTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTokensRequest)(nil), "referralmanager.GetTokensRequest")
	proto.RegisterType((*GetTokensResponse)(nil), "referralmanager.GetTokensResponse")
	proto.RegisterType((*ReserveTokenRequest)(nil), "referralmanager.ReserveTokenRequest")
	proto.RegisterType((*ReserveTokenResponse)(nil), "referralmanager.ReserveTokenResponse")
	proto.RegisterType((*RedeemTokenRequest)(nil), "referralmanager.RedeemTokenRequest")
	proto.RegisterType((*RedeemTokenResponse)(nil), "referralmanager.RedeemTokenResponse")
}

func init() { proto.RegisterFile("referralmanager.proto", fileDescriptor_45d96ad24f1e021c) }

var fileDescriptor_45d96ad24f1e021c = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x6d, 0x55, 0x8c, 0xcf, 0x06, 0x74, 0x04, 0x24, 0xdd, 0x80, 0x15, 0x13, 0x8c, 0x09,
	0x24, 0x7a, 0x03, 0x24, 0x31, 0x2c, 0x74, 0x31, 0xba, 0x30, 0xba, 0x20, 0x25, 0x7d, 0x34, 0x8d,
	0x6d, 0xa7, 0xcc, 0x0c, 0x9e, 0xc6, 0x03, 0x79, 0x06, 0x17, 0x9c, 0xc5, 0xb4, 0x53, 0x9a, 0xa1,
	0x12, 0x74, 0xf9, 0x4f, 0xbf, 0xf7, 0xcf, 0xfc, 0xff, 0x2b, 0x34, 0x38, 0xce, 0x90, 0x73, 0x37,
	0x8c, 0xdc, 0xd8, 0xf5, 0x91, 0xf7, 0x13, 0xce, 0x24, 0x23, 0xb5, 0xd2, 0xb1, 0x0d, 0x3e, 0xf3,
	0x99, 0xfa, 0x68, 0xb7, 0x7d, 0xc6, 0xfc, 0x10, 0x07, 0x99, 0x9a, 0x2e, 0x66, 0x03, 0x19, 0x44,
	0x28, 0xa4, 0x1b, 0x25, 0x39, 0x50, 0x8d, 0x50, 0xba, 0x41, 0x3c, 0x5b, 0x0d, 0x58, 0x8c, 0x7b,
	0xc8, 0x85, 0x52, 0xce, 0x35, 0x1c, 0xdf, 0xa3, 0x7c, 0x66, 0xef, 0x18, 0x0b, 0x8a, 0xf3, 0x05,
	0x0a, 0x49, 0xce, 0xe0, 0x60, 0x21, 0x90, 0x4f, 0x02, 0xaf, 0x65, 0x74, 0x8c, 0x9e, 0x45, 0x2b,
	0xa9, 0x1c, 0x7b, 0xce, 0x15, 0x9c, 0x68, 0xb0, 0x48, 0x58, 0x2c, 0x90, 0xd4, 0x61, 0x5f, 0xa6,
	0x27, 0x2d, 0xa3, 0xb3, 0xdb, 0xb3, 0xa8, 0x12, 0xce, 0x1c, 0x4e, 0x29, 0x0a, 0xe4, 0x1f, 0x98,
	0xe1, 0x2b, 0x6b, 0x0d, 0x36, 0x0a, 0x98, 0x8c, 0xa0, 0xc9, 0xd1, 0x43, 0x8c, 0x82, 0xd8, 0x9f,
	0x08, 0x57, 0x62, 0x18, 0x06, 0x12, 0xd3, 0xfb, 0xcd, 0x14, 0x1b, 0x56, 0xbf, 0x96, 0xed, 0x9d,
	0xef, 0x65, 0xbb, 0xf2, 0xc8, 0x3c, 0x1c, 0x8f, 0x68, 0xbd, 0xa0, 0x9f, 0x56, 0xf0, 0xd8, 0x73,
	0x9a, 0x50, 0x5f, 0xbf, 0x52, 0x3d, 0xd0, 0xb9, 0x03, 0x42, 0x33, 0xfe, 0x1f, 0x2f, 0xd1, 0xa2,
	0x9b, 0x6b, 0xd1, 0x1b, 0x69, 0x1e, 0xcd, 0x44, 0x79, 0xdf, 0x7c, 0x9a, 0x50, 0xa3, 0xf9, 0x76,
	0x1e, 0xd4, 0x76, 0x08, 0x85, 0xc3, 0xa2, 0x25, 0x72, 0xde, 0x2f, 0xef, 0xb4, 0x5c, 0xb7, 0xed,
	0x6c, 0x43, 0xf2, 0x92, 0xdf, 0xc0, 0xd2, 0xb3, 0x91, 0xee, 0xaf, 0x99, 0x0d, 0x6d, 0xdb, 0x97,
	0x7f, 0x50, 0xb9, 0xf9, 0x0b, 0x1c, 0x69, 0xd9, 0xc8, 0xc5, 0x86, 0xa9, 0x72, 0x7d, 0x76, 0x77,
	0x3b, 0xa4, 0x9c, 0x87, 0x7b, 0xaf, 0x66, 0x32, 0x9d, 0x56, 0xb2, 0x5f, 0xed, 0xf6, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x8c, 0x84, 0x1b, 0xdf, 0x02, 0x00, 0x00,
}

type DRPCReferralManagerClient interface {
	DRPCConn() drpc.Conn

	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(ctx context.Context, in *GetTokensRequest) (*GetTokensResponse, error)
	// ReserveToken validates a referral token from referral manager
	ReserveToken(ctx context.Context, in *ReserveTokenRequest) (*ReserveTokenResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(ctx context.Context, in *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

type drpcReferralManagerClient struct {
	cc drpc.Conn
}

func NewDRPCReferralManagerClient(cc drpc.Conn) DRPCReferralManagerClient {
	return &drpcReferralManagerClient{cc}
}

func (c *drpcReferralManagerClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcReferralManagerClient) GetTokens(ctx context.Context, in *GetTokensRequest) (*GetTokensResponse, error) {
	out := new(GetTokensResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/GetTokens", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcReferralManagerClient) ReserveToken(ctx context.Context, in *ReserveTokenRequest) (*ReserveTokenResponse, error) {
	out := new(ReserveTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/ReserveToken", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcReferralManagerClient) RedeemToken(ctx context.Context, in *RedeemTokenRequest) (*RedeemTokenResponse, error) {
	out := new(RedeemTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/RedeemToken", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCReferralManagerServer interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error)
	// ReserveToken validates a referral token from referral manager
	ReserveToken(context.Context, *ReserveTokenRequest) (*ReserveTokenResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(context.Context, *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

type DRPCReferralManagerDescription struct{}

func (DRPCReferralManagerDescription) NumMethods() int { return 3 }

func (DRPCReferralManagerDescription) Method(n int) (string, drpc.Handler, interface{}, bool) {
	switch n {
	case 0:
		return "/referralmanager.ReferralManager/GetTokens",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCReferralManagerServer).
					GetTokens(
						ctx,
						in1.(*GetTokensRequest),
					)
			}, DRPCReferralManagerServer.GetTokens, true
	case 1:
		return "/referralmanager.ReferralManager/ReserveToken",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCReferralManagerServer).
					ReserveToken(
						ctx,
						in1.(*ReserveTokenRequest),
					)
			}, DRPCReferralManagerServer.ReserveToken, true
	case 2:
		return "/referralmanager.ReferralManager/RedeemToken",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCReferralManagerServer).
					RedeemToken(
						ctx,
						in1.(*RedeemTokenRequest),
					)
			}, DRPCReferralManagerServer.RedeemToken, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterReferralManager(srv drpc.Server, impl DRPCReferralManagerServer) {
	srv.Register(impl, DRPCReferralManagerDescription{})
}

type DRPCReferralManager_GetTokensStream interface {
	drpc.Stream
	SendAndClose(*GetTokensResponse) error
}

type drpcReferralManagerGetTokensStream struct {
	drpc.Stream
}

func (x *drpcReferralManagerGetTokensStream) SendAndClose(m *GetTokensResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCReferralManager_ReserveTokenStream interface {
	drpc.Stream
	SendAndClose(*ReserveTokenResponse) error
}

type drpcReferralManagerReserveTokenStream struct {
	drpc.Stream
}

func (x *drpcReferralManagerReserveTokenStream) SendAndClose(m *ReserveTokenResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCReferralManager_RedeemTokenStream interface {
	drpc.Stream
	SendAndClose(*RedeemTokenResponse) error
}

type drpcReferralManagerRedeemTokenStream struct {
	drpc.Stream
}

func (x *drpcReferralManagerRedeemTokenStream) SendAndClose(m *RedeemTokenResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReferralManagerClient is the client API for ReferralManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReferralManagerClient interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error)
	// ReserveToken validates a referral token from referral manager
	ReserveToken(ctx context.Context, in *ReserveTokenRequest, opts ...grpc.CallOption) (*ReserveTokenResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(ctx context.Context, in *RedeemTokenRequest, opts ...grpc.CallOption) (*RedeemTokenResponse, error)
}

type referralManagerClient struct {
	cc *grpc.ClientConn
}

func NewReferralManagerClient(cc *grpc.ClientConn) ReferralManagerClient {
	return &referralManagerClient{cc}
}

func (c *referralManagerClient) GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error) {
	out := new(GetTokensResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referralManagerClient) ReserveToken(ctx context.Context, in *ReserveTokenRequest, opts ...grpc.CallOption) (*ReserveTokenResponse, error) {
	out := new(ReserveTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/ReserveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referralManagerClient) RedeemToken(ctx context.Context, in *RedeemTokenRequest, opts ...grpc.CallOption) (*RedeemTokenResponse, error) {
	out := new(RedeemTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/RedeemToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReferralManagerServer is the server API for ReferralManager service.
type ReferralManagerServer interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error)
	// ReserveToken validates a referral token from referral manager
	ReserveToken(context.Context, *ReserveTokenRequest) (*ReserveTokenResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(context.Context, *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

func RegisterReferralManagerServer(s *grpc.Server, srv ReferralManagerServer) {
	s.RegisterService(&_ReferralManager_serviceDesc, srv)
}

func _ReferralManager_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferralManagerServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/referralmanager.ReferralManager/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferralManagerServer).GetTokens(ctx, req.(*GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferralManager_ReserveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferralManagerServer).ReserveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/referralmanager.ReferralManager/ReserveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferralManagerServer).ReserveToken(ctx, req.(*ReserveTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferralManager_RedeemToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferralManagerServer).RedeemToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/referralmanager.ReferralManager/RedeemToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferralManagerServer).RedeemToken(ctx, req.(*RedeemTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReferralManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "referralmanager.ReferralManager",
	HandlerType: (*ReferralManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokens",
			Handler:    _ReferralManager_GetTokens_Handler,
		},
		{
			MethodName: "ReserveToken",
			Handler:    _ReferralManager_ReserveToken_Handler,
		},
		{
			MethodName: "RedeemToken",
			Handler:    _ReferralManager_RedeemToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "referralmanager.proto",
}
