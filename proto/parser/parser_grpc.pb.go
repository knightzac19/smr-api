// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.4
// source: parser/parser.proto

package parser

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

const (
	Parser_Parse_FullMethodName = "/Parser/Parse"
)

// ParserClient is the client API for Parser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserClient interface {
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (Parser_ParseClient, error)
}

type parserClient struct {
	cc grpc.ClientConnInterface
}

func NewParserClient(cc grpc.ClientConnInterface) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (Parser_ParseClient, error) {
	stream, err := c.cc.NewStream(ctx, &Parser_ServiceDesc.Streams[0], Parser_Parse_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &parserParseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Parser_ParseClient interface {
	Recv() (*AssetResponse, error)
	grpc.ClientStream
}

type parserParseClient struct {
	grpc.ClientStream
}

func (x *parserParseClient) Recv() (*AssetResponse, error) {
	m := new(AssetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ParserServer is the server API for Parser service.
// All implementations must embed UnimplementedParserServer
// for forward compatibility
type ParserServer interface {
	Parse(*ParseRequest, Parser_ParseServer) error
	mustEmbedUnimplementedParserServer()
}

// UnimplementedParserServer must be embedded to have forward compatible implementations.
type UnimplementedParserServer struct {
}

func (UnimplementedParserServer) Parse(*ParseRequest, Parser_ParseServer) error {
	return status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedParserServer) mustEmbedUnimplementedParserServer() {}

// UnsafeParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServer will
// result in compilation errors.
type UnsafeParserServer interface {
	mustEmbedUnimplementedParserServer()
}

func RegisterParserServer(s grpc.ServiceRegistrar, srv ParserServer) {
	s.RegisterService(&Parser_ServiceDesc, srv)
}

func _Parser_Parse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ParseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ParserServer).Parse(m, &parserParseServer{stream})
}

type Parser_ParseServer interface {
	Send(*AssetResponse) error
	grpc.ServerStream
}

type parserParseServer struct {
	grpc.ServerStream
}

func (x *parserParseServer) Send(m *AssetResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Parser_ServiceDesc is the grpc.ServiceDesc for Parser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Parser",
	HandlerType: (*ParserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Parse",
			Handler:       _Parser_Parse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "parser/parser.proto",
}
