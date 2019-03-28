// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RestoreItemAction.proto

package generated

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

type RestoreExecuteRequest struct {
	Plugin         string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Item           []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Restore        []byte `protobuf:"bytes,3,opt,name=restore,proto3" json:"restore,omitempty"`
	ItemFromBackup []byte `protobuf:"bytes,4,opt,name=itemFromBackup,proto3" json:"itemFromBackup,omitempty"`
}

func (m *RestoreExecuteRequest) Reset()                    { *m = RestoreExecuteRequest{} }
func (m *RestoreExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreExecuteRequest) ProtoMessage()               {}
func (*RestoreExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RestoreExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *RestoreExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreExecuteRequest) GetRestore() []byte {
	if m != nil {
		return m.Restore
	}
	return nil
}

func (m *RestoreExecuteRequest) GetItemFromBackup() []byte {
	if m != nil {
		return m.ItemFromBackup
	}
	return nil
}

type RestoreExecuteResponse struct {
	Item    []byte `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Warning string `protobuf:"bytes,2,opt,name=warning" json:"warning,omitempty"`
}

func (m *RestoreExecuteResponse) Reset()                    { *m = RestoreExecuteResponse{} }
func (m *RestoreExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*RestoreExecuteResponse) ProtoMessage()               {}
func (*RestoreExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RestoreExecuteResponse) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreExecuteResponse) GetWarning() string {
	if m != nil {
		return m.Warning
	}
	return ""
}

func init() {
	proto.RegisterType((*RestoreExecuteRequest)(nil), "generated.RestoreExecuteRequest")
	proto.RegisterType((*RestoreExecuteResponse)(nil), "generated.RestoreExecuteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RestoreItemAction service

type RestoreItemActionClient interface {
	AppliesTo(ctx context.Context, in *AppliesToRequest, opts ...grpc.CallOption) (*AppliesToResponse, error)
	Execute(ctx context.Context, in *RestoreExecuteRequest, opts ...grpc.CallOption) (*RestoreExecuteResponse, error)
}

type restoreItemActionClient struct {
	cc *grpc.ClientConn
}

func NewRestoreItemActionClient(cc *grpc.ClientConn) RestoreItemActionClient {
	return &restoreItemActionClient{cc}
}

func (c *restoreItemActionClient) AppliesTo(ctx context.Context, in *AppliesToRequest, opts ...grpc.CallOption) (*AppliesToResponse, error) {
	out := new(AppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restoreItemActionClient) Execute(ctx context.Context, in *RestoreExecuteRequest, opts ...grpc.CallOption) (*RestoreExecuteResponse, error) {
	out := new(RestoreExecuteResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RestoreItemAction service

type RestoreItemActionServer interface {
	AppliesTo(context.Context, *AppliesToRequest) (*AppliesToResponse, error)
	Execute(context.Context, *RestoreExecuteRequest) (*RestoreExecuteResponse, error)
}

func RegisterRestoreItemActionServer(s *grpc.Server, srv RestoreItemActionServer) {
	s.RegisterService(&_RestoreItemAction_serviceDesc, srv)
}

func _RestoreItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, req.(*AppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestoreItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).Execute(ctx, req.(*RestoreExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestoreItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.RestoreItemAction",
	HandlerType: (*RestoreItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _RestoreItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _RestoreItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RestoreItemAction.proto",
}

func init() { proto.RegisterFile("RestoreItemAction.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4b, 0x4e, 0xc3, 0x30,
	0x14, 0x94, 0xa1, 0x6a, 0x95, 0xa7, 0x0a, 0x89, 0x27, 0x51, 0xac, 0xc0, 0x22, 0x74, 0x81, 0xba,
	0xca, 0x02, 0x4e, 0x50, 0x24, 0x22, 0xb1, 0x61, 0x61, 0xb8, 0x40, 0x48, 0x9f, 0x82, 0x45, 0x63,
	0x1b, 0x7f, 0x04, 0x17, 0xe0, 0x34, 0x5c, 0x12, 0xe1, 0xb8, 0x11, 0x94, 0x88, 0x9d, 0x67, 0xde,
	0xf8, 0xcd, 0x78, 0x0c, 0xa7, 0x82, 0x9c, 0xd7, 0x96, 0xee, 0x3c, 0x75, 0xeb, 0xc6, 0x4b, 0xad,
	0x4a, 0x63, 0xb5, 0xd7, 0x98, 0xb5, 0xa4, 0xc8, 0xd6, 0x9e, 0x36, 0xf9, 0xfc, 0xe1, 0xb9, 0xb6,
	0xb4, 0xe9, 0x07, 0xcb, 0x0f, 0x06, 0x27, 0xe9, 0xd2, 0xed, 0x3b, 0x35, 0xc1, 0x93, 0xa0, 0xd7,
	0x40, 0xce, 0xe3, 0x02, 0xa6, 0x66, 0x1b, 0x5a, 0xa9, 0x38, 0x2b, 0xd8, 0x2a, 0x13, 0x09, 0x21,
	0xc2, 0x44, 0x7a, 0xea, 0xf8, 0x41, 0xc1, 0x56, 0x73, 0x11, 0xcf, 0xc8, 0x61, 0x66, 0xfb, 0x25,
	0xfc, 0x30, 0xd2, 0x3b, 0x88, 0x97, 0x70, 0xf4, 0xad, 0xa8, 0xac, 0xee, 0x6e, 0xea, 0xe6, 0x25,
	0x18, 0x3e, 0x89, 0x82, 0x3d, 0x76, 0x59, 0xc1, 0x62, 0x3f, 0x86, 0x33, 0x5a, 0x39, 0x1a, 0xfc,
	0xd8, 0x6f, 0xbf, 0xb7, 0xda, 0x2a, 0xa9, 0xda, 0x18, 0x23, 0x13, 0x3b, 0x78, 0xf5, 0xc9, 0xe0,
	0xf8, 0x4f, 0x09, 0x58, 0x41, 0xb6, 0x36, 0x66, 0x2b, 0xc9, 0x3d, 0x6a, 0x3c, 0x2b, 0x87, 0x32,
	0xca, 0x81, 0x4d, 0xaf, 0xce, 0xcf, 0xc7, 0x87, 0x29, 0xcb, 0x3d, 0xcc, 0x52, 0x3c, 0x2c, 0x7e,
	0x08, 0x47, 0x0b, 0xcc, 0x2f, 0xfe, 0x51, 0xf4, 0xfb, 0x9e, 0xa6, 0xf1, 0x13, 0xae, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x19, 0xeb, 0x64, 0x57, 0xb8, 0x01, 0x00, 0x00,
}