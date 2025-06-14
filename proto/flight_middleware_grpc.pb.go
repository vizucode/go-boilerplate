// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/flight_middleware.proto

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FlightResponse_FlightSchedule_FullMethodName      = "/server.FlightResponse/FlightSchedule"
	FlightResponse_FlightPrice_FullMethodName         = "/server.FlightResponse/FlightPrice"
	FlightResponse_FlightBooking_FullMethodName       = "/server.FlightResponse/FlightBooking"
	FlightResponse_FlightCancelBooking_FullMethodName = "/server.FlightResponse/FlightCancelBooking"
	FlightResponse_FlightTicketing_FullMethodName     = "/server.FlightResponse/FlightTicketing"
	FlightResponse_FlightHistories_FullMethodName     = "/server.FlightResponse/FlightHistories"
	FlightResponse_CheckProvider_FullMethodName       = "/server.FlightResponse/CheckProvider"
	FlightResponse_GetInsurance_FullMethodName        = "/server.FlightResponse/GetInsurance"
)

// FlightResponseClient is the client API for FlightResponse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service Definitions
type FlightResponseClient interface {
	FlightSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	FlightPrice(ctx context.Context, in *FlightPriceRequest, opts ...grpc.CallOption) (*FlightPriceResponse, error)
	FlightBooking(ctx context.Context, in *FlightPriceRequest, opts ...grpc.CallOption) (*FlightBookResponse, error)
	FlightCancelBooking(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*RequestAfterBook, error)
	FlightTicketing(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*RequestAfterBook, error)
	FlightHistories(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*FlightHistoryResponse, error)
	CheckProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseCheckProvider, error)
	GetInsurance(ctx context.Context, in *RequestFlightInsurance, opts ...grpc.CallOption) (*ResponseFlightInsurance, error)
}

type flightResponseClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightResponseClient(cc grpc.ClientConnInterface) FlightResponseClient {
	return &flightResponseClient{cc}
}

func (c *flightResponseClient) FlightSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, FlightResponse_FlightSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) FlightPrice(ctx context.Context, in *FlightPriceRequest, opts ...grpc.CallOption) (*FlightPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlightPriceResponse)
	err := c.cc.Invoke(ctx, FlightResponse_FlightPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) FlightBooking(ctx context.Context, in *FlightPriceRequest, opts ...grpc.CallOption) (*FlightBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlightBookResponse)
	err := c.cc.Invoke(ctx, FlightResponse_FlightBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) FlightCancelBooking(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*RequestAfterBook, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestAfterBook)
	err := c.cc.Invoke(ctx, FlightResponse_FlightCancelBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) FlightTicketing(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*RequestAfterBook, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestAfterBook)
	err := c.cc.Invoke(ctx, FlightResponse_FlightTicketing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) FlightHistories(ctx context.Context, in *RequestAfterBook, opts ...grpc.CallOption) (*FlightHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlightHistoryResponse)
	err := c.cc.Invoke(ctx, FlightResponse_FlightHistories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) CheckProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseCheckProvider, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseCheckProvider)
	err := c.cc.Invoke(ctx, FlightResponse_CheckProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightResponseClient) GetInsurance(ctx context.Context, in *RequestFlightInsurance, opts ...grpc.CallOption) (*ResponseFlightInsurance, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseFlightInsurance)
	err := c.cc.Invoke(ctx, FlightResponse_GetInsurance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightResponseServer is the server API for FlightResponse service.
// All implementations must embed UnimplementedFlightResponseServer
// for forward compatibility.
//
// Service Definitions
type FlightResponseServer interface {
	FlightSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	FlightPrice(context.Context, *FlightPriceRequest) (*FlightPriceResponse, error)
	FlightBooking(context.Context, *FlightPriceRequest) (*FlightBookResponse, error)
	FlightCancelBooking(context.Context, *RequestAfterBook) (*RequestAfterBook, error)
	FlightTicketing(context.Context, *RequestAfterBook) (*RequestAfterBook, error)
	FlightHistories(context.Context, *RequestAfterBook) (*FlightHistoryResponse, error)
	CheckProvider(context.Context, *emptypb.Empty) (*ResponseCheckProvider, error)
	GetInsurance(context.Context, *RequestFlightInsurance) (*ResponseFlightInsurance, error)
	mustEmbedUnimplementedFlightResponseServer()
}

// UnimplementedFlightResponseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFlightResponseServer struct{}

func (UnimplementedFlightResponseServer) FlightSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightSchedule not implemented")
}
func (UnimplementedFlightResponseServer) FlightPrice(context.Context, *FlightPriceRequest) (*FlightPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightPrice not implemented")
}
func (UnimplementedFlightResponseServer) FlightBooking(context.Context, *FlightPriceRequest) (*FlightBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightBooking not implemented")
}
func (UnimplementedFlightResponseServer) FlightCancelBooking(context.Context, *RequestAfterBook) (*RequestAfterBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightCancelBooking not implemented")
}
func (UnimplementedFlightResponseServer) FlightTicketing(context.Context, *RequestAfterBook) (*RequestAfterBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightTicketing not implemented")
}
func (UnimplementedFlightResponseServer) FlightHistories(context.Context, *RequestAfterBook) (*FlightHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlightHistories not implemented")
}
func (UnimplementedFlightResponseServer) CheckProvider(context.Context, *emptypb.Empty) (*ResponseCheckProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProvider not implemented")
}
func (UnimplementedFlightResponseServer) GetInsurance(context.Context, *RequestFlightInsurance) (*ResponseFlightInsurance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInsurance not implemented")
}
func (UnimplementedFlightResponseServer) mustEmbedUnimplementedFlightResponseServer() {}
func (UnimplementedFlightResponseServer) testEmbeddedByValue()                        {}

// UnsafeFlightResponseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightResponseServer will
// result in compilation errors.
type UnsafeFlightResponseServer interface {
	mustEmbedUnimplementedFlightResponseServer()
}

func RegisterFlightResponseServer(s grpc.ServiceRegistrar, srv FlightResponseServer) {
	// If the following call pancis, it indicates UnimplementedFlightResponseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FlightResponse_ServiceDesc, srv)
}

func _FlightResponse_FlightSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightSchedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_FlightPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightPrice(ctx, req.(*FlightPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_FlightBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightBooking(ctx, req.(*FlightPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_FlightCancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAfterBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightCancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightCancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightCancelBooking(ctx, req.(*RequestAfterBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_FlightTicketing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAfterBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightTicketing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightTicketing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightTicketing(ctx, req.(*RequestAfterBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_FlightHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAfterBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).FlightHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_FlightHistories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).FlightHistories(ctx, req.(*RequestAfterBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_CheckProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).CheckProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_CheckProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).CheckProvider(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightResponse_GetInsurance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFlightInsurance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightResponseServer).GetInsurance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightResponse_GetInsurance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightResponseServer).GetInsurance(ctx, req.(*RequestFlightInsurance))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightResponse_ServiceDesc is the grpc.ServiceDesc for FlightResponse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightResponse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.FlightResponse",
	HandlerType: (*FlightResponseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FlightSchedule",
			Handler:    _FlightResponse_FlightSchedule_Handler,
		},
		{
			MethodName: "FlightPrice",
			Handler:    _FlightResponse_FlightPrice_Handler,
		},
		{
			MethodName: "FlightBooking",
			Handler:    _FlightResponse_FlightBooking_Handler,
		},
		{
			MethodName: "FlightCancelBooking",
			Handler:    _FlightResponse_FlightCancelBooking_Handler,
		},
		{
			MethodName: "FlightTicketing",
			Handler:    _FlightResponse_FlightTicketing_Handler,
		},
		{
			MethodName: "FlightHistories",
			Handler:    _FlightResponse_FlightHistories_Handler,
		},
		{
			MethodName: "CheckProvider",
			Handler:    _FlightResponse_CheckProvider_Handler,
		},
		{
			MethodName: "GetInsurance",
			Handler:    _FlightResponse_GetInsurance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/flight_middleware.proto",
}
