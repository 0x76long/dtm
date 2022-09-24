// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: workflow/wfpb/wf.proto

package wfpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowClient interface {
	Execute(ctx context.Context, in *WorkflowData, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type workflowClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowClient(cc grpc.ClientConnInterface) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) Execute(ctx context.Context, in *WorkflowData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/workflow.Workflow/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
// All implementations must embed UnimplementedWorkflowServer
// for forward compatibility
type WorkflowServer interface {
	Execute(context.Context, *WorkflowData) (*emptypb.Empty, error)
	mustEmbedUnimplementedWorkflowServer()
}

// UnimplementedWorkflowServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (UnimplementedWorkflowServer) Execute(context.Context, *WorkflowData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedWorkflowServer) mustEmbedUnimplementedWorkflowServer() {}

// UnsafeWorkflowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServer will
// result in compilation errors.
type UnsafeWorkflowServer interface {
	mustEmbedUnimplementedWorkflowServer()
}

func RegisterWorkflowServer(s grpc.ServiceRegistrar, srv WorkflowServer) {
	s.RegisterService(&Workflow_ServiceDesc, srv)
}

func _Workflow_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.Workflow/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).Execute(ctx, req.(*WorkflowData))
	}
	return interceptor(ctx, in, info, handler)
}

// Workflow_ServiceDesc is the grpc.ServiceDesc for Workflow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workflow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Workflow_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow/wfpb/wf.proto",
}