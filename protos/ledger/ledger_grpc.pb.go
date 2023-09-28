// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: protos/ledger/ledger.proto

package __

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

// LedgerServiceClient is the client API for LedgerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerServiceClient interface {
	AddExpense(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error)
	ResetExpenses(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error)
	GetAllExpenses(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error)
}

type ledgerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerServiceClient(cc grpc.ClientConnInterface) LedgerServiceClient {
	return &ledgerServiceClient{cc}
}

func (c *ledgerServiceClient) AddExpense(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error) {
	out := new(LedgerResponse)
	err := c.cc.Invoke(ctx, "/ledger.LedgerService/AddExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ResetExpenses(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error) {
	out := new(LedgerResponse)
	err := c.cc.Invoke(ctx, "/ledger.LedgerService/ResetExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetAllExpenses(ctx context.Context, in *LedgerRequest, opts ...grpc.CallOption) (*LedgerResponse, error) {
	out := new(LedgerResponse)
	err := c.cc.Invoke(ctx, "/ledger.LedgerService/GetAllExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerServiceServer is the server API for LedgerService service.
// All implementations must embed UnimplementedLedgerServiceServer
// for forward compatibility
type LedgerServiceServer interface {
	AddExpense(context.Context, *LedgerRequest) (*LedgerResponse, error)
	ResetExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error)
	GetAllExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error)
	mustEmbedUnimplementedLedgerServiceServer()
}

// UnimplementedLedgerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLedgerServiceServer struct {
}

func (UnimplementedLedgerServiceServer) AddExpense(context.Context, *LedgerRequest) (*LedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExpense not implemented")
}
func (UnimplementedLedgerServiceServer) ResetExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetExpenses not implemented")
}
func (UnimplementedLedgerServiceServer) GetAllExpenses(context.Context, *LedgerRequest) (*LedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllExpenses not implemented")
}
func (UnimplementedLedgerServiceServer) mustEmbedUnimplementedLedgerServiceServer() {}

// UnsafeLedgerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerServiceServer will
// result in compilation errors.
type UnsafeLedgerServiceServer interface {
	mustEmbedUnimplementedLedgerServiceServer()
}

func RegisterLedgerServiceServer(s grpc.ServiceRegistrar, srv LedgerServiceServer) {
	s.RegisterService(&LedgerService_ServiceDesc, srv)
}

func _LedgerService_AddExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).AddExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.LedgerService/AddExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).AddExpense(ctx, req.(*LedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ResetExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ResetExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.LedgerService/ResetExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ResetExpenses(ctx, req.(*LedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetAllExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetAllExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.LedgerService/GetAllExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetAllExpenses(ctx, req.(*LedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LedgerService_ServiceDesc is the grpc.ServiceDesc for LedgerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LedgerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.LedgerService",
	HandlerType: (*LedgerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddExpense",
			Handler:    _LedgerService_AddExpense_Handler,
		},
		{
			MethodName: "ResetExpenses",
			Handler:    _LedgerService_ResetExpenses_Handler,
		},
		{
			MethodName: "GetAllExpenses",
			Handler:    _LedgerService_GetAllExpenses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/ledger/ledger.proto",
}
