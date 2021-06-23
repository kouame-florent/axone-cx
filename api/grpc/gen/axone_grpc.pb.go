// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AxoneClient is the client API for Axone service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AxoneClient interface {
	SendNewTicket(ctx context.Context, in *NewTicketRequest, opts ...grpc.CallOption) (*NewTicketResponse, error)
	SendAttachment(ctx context.Context, opts ...grpc.CallOption) (Axone_SendAttachmentClient, error)
	ListRequesterTickets(ctx context.Context, in *ListRequesterTicketsRequest, opts ...grpc.CallOption) (*ListRequesterTicketsResponse, error)
	ListAgentTickets(ctx context.Context, in *AgentTicketsListRequest, opts ...grpc.CallOption) (*AgentTicketsListResponse, error)
	Subscribe(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Axone_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type axoneClient struct {
	cc grpc.ClientConnInterface
}

func NewAxoneClient(cc grpc.ClientConnInterface) AxoneClient {
	return &axoneClient{cc}
}

func (c *axoneClient) SendNewTicket(ctx context.Context, in *NewTicketRequest, opts ...grpc.CallOption) (*NewTicketResponse, error) {
	out := new(NewTicketResponse)
	err := c.cc.Invoke(ctx, "/api.Axone/SendNewTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *axoneClient) SendAttachment(ctx context.Context, opts ...grpc.CallOption) (Axone_SendAttachmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Axone_ServiceDesc.Streams[0], "/api.Axone/SendAttachment", opts...)
	if err != nil {
		return nil, err
	}
	x := &axoneSendAttachmentClient{stream}
	return x, nil
}

type Axone_SendAttachmentClient interface {
	Send(*AttachmentRequest) error
	CloseAndRecv() (*AttachmentResponse, error)
	grpc.ClientStream
}

type axoneSendAttachmentClient struct {
	grpc.ClientStream
}

func (x *axoneSendAttachmentClient) Send(m *AttachmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *axoneSendAttachmentClient) CloseAndRecv() (*AttachmentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AttachmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *axoneClient) ListRequesterTickets(ctx context.Context, in *ListRequesterTicketsRequest, opts ...grpc.CallOption) (*ListRequesterTicketsResponse, error) {
	out := new(ListRequesterTicketsResponse)
	err := c.cc.Invoke(ctx, "/api.Axone/ListRequesterTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *axoneClient) ListAgentTickets(ctx context.Context, in *AgentTicketsListRequest, opts ...grpc.CallOption) (*AgentTicketsListResponse, error) {
	out := new(AgentTicketsListResponse)
	err := c.cc.Invoke(ctx, "/api.Axone/ListAgentTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *axoneClient) Subscribe(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Axone_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Axone_ServiceDesc.Streams[1], "/api.Axone/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &axoneSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Axone_SubscribeClient interface {
	Recv() (*NotificationResponse, error)
	grpc.ClientStream
}

type axoneSubscribeClient struct {
	grpc.ClientStream
}

func (x *axoneSubscribeClient) Recv() (*NotificationResponse, error) {
	m := new(NotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *axoneClient) Unsubscribe(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/api.Axone/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *axoneClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.Axone/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AxoneServer is the server API for Axone service.
// All implementations must embed UnimplementedAxoneServer
// for forward compatibility
type AxoneServer interface {
	SendNewTicket(context.Context, *NewTicketRequest) (*NewTicketResponse, error)
	SendAttachment(Axone_SendAttachmentServer) error
	ListRequesterTickets(context.Context, *ListRequesterTicketsRequest) (*ListRequesterTicketsResponse, error)
	ListAgentTickets(context.Context, *AgentTicketsListRequest) (*AgentTicketsListResponse, error)
	Subscribe(*NotificationRequest, Axone_SubscribeServer) error
	Unsubscribe(context.Context, *NotificationRequest) (*NotificationResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAxoneServer()
}

// UnimplementedAxoneServer must be embedded to have forward compatible implementations.
type UnimplementedAxoneServer struct {
}

func (UnimplementedAxoneServer) SendNewTicket(context.Context, *NewTicketRequest) (*NewTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNewTicket not implemented")
}
func (UnimplementedAxoneServer) SendAttachment(Axone_SendAttachmentServer) error {
	return status.Errorf(codes.Unimplemented, "method SendAttachment not implemented")
}
func (UnimplementedAxoneServer) ListRequesterTickets(context.Context, *ListRequesterTicketsRequest) (*ListRequesterTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequesterTickets not implemented")
}
func (UnimplementedAxoneServer) ListAgentTickets(context.Context, *AgentTicketsListRequest) (*AgentTicketsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentTickets not implemented")
}
func (UnimplementedAxoneServer) Subscribe(*NotificationRequest, Axone_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedAxoneServer) Unsubscribe(context.Context, *NotificationRequest) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedAxoneServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAxoneServer) mustEmbedUnimplementedAxoneServer() {}

// UnsafeAxoneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AxoneServer will
// result in compilation errors.
type UnsafeAxoneServer interface {
	mustEmbedUnimplementedAxoneServer()
}

func RegisterAxoneServer(s grpc.ServiceRegistrar, srv AxoneServer) {
	s.RegisterService(&Axone_ServiceDesc, srv)
}

func _Axone_SendNewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxoneServer).SendNewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Axone/SendNewTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxoneServer).SendNewTicket(ctx, req.(*NewTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Axone_SendAttachment_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AxoneServer).SendAttachment(&axoneSendAttachmentServer{stream})
}

type Axone_SendAttachmentServer interface {
	SendAndClose(*AttachmentResponse) error
	Recv() (*AttachmentRequest, error)
	grpc.ServerStream
}

type axoneSendAttachmentServer struct {
	grpc.ServerStream
}

func (x *axoneSendAttachmentServer) SendAndClose(m *AttachmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *axoneSendAttachmentServer) Recv() (*AttachmentRequest, error) {
	m := new(AttachmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Axone_ListRequesterTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequesterTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxoneServer).ListRequesterTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Axone/ListRequesterTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxoneServer).ListRequesterTickets(ctx, req.(*ListRequesterTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Axone_ListAgentTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentTicketsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxoneServer).ListAgentTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Axone/ListAgentTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxoneServer).ListAgentTickets(ctx, req.(*AgentTicketsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Axone_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AxoneServer).Subscribe(m, &axoneSubscribeServer{stream})
}

type Axone_SubscribeServer interface {
	Send(*NotificationResponse) error
	grpc.ServerStream
}

type axoneSubscribeServer struct {
	grpc.ServerStream
}

func (x *axoneSubscribeServer) Send(m *NotificationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Axone_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxoneServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Axone/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxoneServer).Unsubscribe(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Axone_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxoneServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Axone/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxoneServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Axone_ServiceDesc is the grpc.ServiceDesc for Axone service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Axone_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Axone",
	HandlerType: (*AxoneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNewTicket",
			Handler:    _Axone_SendNewTicket_Handler,
		},
		{
			MethodName: "ListRequesterTickets",
			Handler:    _Axone_ListRequesterTickets_Handler,
		},
		{
			MethodName: "ListAgentTickets",
			Handler:    _Axone_ListAgentTickets_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Axone_Unsubscribe_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Axone_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAttachment",
			Handler:       _Axone_SendAttachment_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Axone_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "axone.proto",
}
