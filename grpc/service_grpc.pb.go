// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: service.proto

package grpc

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	// Only small images since the connection is open during the whole process
	TransferImageBytes(ctx context.Context, opts ...grpc.CallOption) (ImageService_TransferImageBytesClient, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) TransferImageBytes(ctx context.Context, opts ...grpc.CallOption) (ImageService_TransferImageBytesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[0], "/grpc.ImageService/TransferImageBytes", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceTransferImageBytesClient{stream}
	return x, nil
}

type ImageService_TransferImageBytesClient interface {
	Send(*Image) error
	Recv() (*Image, error)
	grpc.ClientStream
}

type imageServiceTransferImageBytesClient struct {
	grpc.ClientStream
}

func (x *imageServiceTransferImageBytesClient) Send(m *Image) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceTransferImageBytesClient) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	// Only small images since the connection is open during the whole process
	TransferImageBytes(ImageService_TransferImageBytesServer) error
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) TransferImageBytes(ImageService_TransferImageBytesServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferImageBytes not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_TransferImageBytes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).TransferImageBytes(&imageServiceTransferImageBytesServer{stream})
}

type ImageService_TransferImageBytesServer interface {
	Send(*Image) error
	Recv() (*Image, error)
	grpc.ServerStream
}

type imageServiceTransferImageBytesServer struct {
	grpc.ServerStream
}

func (x *imageServiceTransferImageBytesServer) Send(m *Image) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceTransferImageBytesServer) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferImageBytes",
			Handler:       _ImageService_TransferImageBytes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
