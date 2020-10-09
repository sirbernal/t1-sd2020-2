// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: proto/cliente_logistica.proto

package cliente_logistica

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Envio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Msg2 string `protobuf:"bytes,2,opt,name=msg2,proto3" json:"msg2,omitempty"`
}

func (x *Envio) Reset() {
	*x = Envio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cliente_logistica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envio) ProtoMessage() {}

func (x *Envio) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cliente_logistica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envio.ProtoReflect.Descriptor instead.
func (*Envio) Descriptor() ([]byte, []int) {
	return file_proto_cliente_logistica_proto_rawDescGZIP(), []int{0}
}

func (x *Envio) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Envio) GetMsg2() string {
	if x != nil {
		return x.Msg2
	}
	return ""
}

type SeguimientoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seg []*Envio `protobuf:"bytes,1,rep,name=seg,proto3" json:"seg,omitempty"`
}

func (x *SeguimientoRequest) Reset() {
	*x = SeguimientoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cliente_logistica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoRequest) ProtoMessage() {}

func (x *SeguimientoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cliente_logistica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoRequest.ProtoReflect.Descriptor instead.
func (*SeguimientoRequest) Descriptor() ([]byte, []int) {
	return file_proto_cliente_logistica_proto_rawDescGZIP(), []int{1}
}

func (x *SeguimientoRequest) GetSeg() []*Envio {
	if x != nil {
		return x.Seg
	}
	return nil
}

type SeguimientoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmation string `protobuf:"bytes,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *SeguimientoResponse) Reset() {
	*x = SeguimientoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cliente_logistica_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoResponse) ProtoMessage() {}

func (x *SeguimientoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cliente_logistica_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoResponse.ProtoReflect.Descriptor instead.
func (*SeguimientoResponse) Descriptor() ([]byte, []int) {
	return file_proto_cliente_logistica_proto_rawDescGZIP(), []int{2}
}

func (x *SeguimientoResponse) GetConfirmation() string {
	if x != nil {
		return x.Confirmation
	}
	return ""
}

var File_proto_cliente_logistica_proto protoreflect.FileDescriptor

var file_proto_cliente_logistica_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x61, 0x22, 0x2d, 0x0a, 0x05, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x73, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x73, 0x67,
	0x32, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x73, 0x65, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x52, 0x03,
	0x73, 0x65, 0x67, 0x22, 0x39, 0x0a, 0x13, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x74,
	0x0a, 0x12, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_cliente_logistica_proto_rawDescOnce sync.Once
	file_proto_cliente_logistica_proto_rawDescData = file_proto_cliente_logistica_proto_rawDesc
)

func file_proto_cliente_logistica_proto_rawDescGZIP() []byte {
	file_proto_cliente_logistica_proto_rawDescOnce.Do(func() {
		file_proto_cliente_logistica_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cliente_logistica_proto_rawDescData)
	})
	return file_proto_cliente_logistica_proto_rawDescData
}

var file_proto_cliente_logistica_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_cliente_logistica_proto_goTypes = []interface{}{
	(*Envio)(nil),               // 0: cliente_logistica.Envio
	(*SeguimientoRequest)(nil),  // 1: cliente_logistica.SeguimientoRequest
	(*SeguimientoResponse)(nil), // 2: cliente_logistica.SeguimientoResponse
}
var file_proto_cliente_logistica_proto_depIdxs = []int32{
	0, // 0: cliente_logistica.SeguimientoRequest.seg:type_name -> cliente_logistica.Envio
	1, // 1: cliente_logistica.SeguimientoService.Seguimiento:input_type -> cliente_logistica.SeguimientoRequest
	2, // 2: cliente_logistica.SeguimientoService.Seguimiento:output_type -> cliente_logistica.SeguimientoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cliente_logistica_proto_init() }
func file_proto_cliente_logistica_proto_init() {
	if File_proto_cliente_logistica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cliente_logistica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cliente_logistica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cliente_logistica_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cliente_logistica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cliente_logistica_proto_goTypes,
		DependencyIndexes: file_proto_cliente_logistica_proto_depIdxs,
		MessageInfos:      file_proto_cliente_logistica_proto_msgTypes,
	}.Build()
	File_proto_cliente_logistica_proto = out.File
	file_proto_cliente_logistica_proto_rawDesc = nil
	file_proto_cliente_logistica_proto_goTypes = nil
	file_proto_cliente_logistica_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SeguimientoServiceClient is the client API for SeguimientoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SeguimientoServiceClient interface {
	Seguimiento(ctx context.Context, in *SeguimientoRequest, opts ...grpc.CallOption) (*SeguimientoResponse, error)
}

type seguimientoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSeguimientoServiceClient(cc grpc.ClientConnInterface) SeguimientoServiceClient {
	return &seguimientoServiceClient{cc}
}

func (c *seguimientoServiceClient) Seguimiento(ctx context.Context, in *SeguimientoRequest, opts ...grpc.CallOption) (*SeguimientoResponse, error) {
	out := new(SeguimientoResponse)
	err := c.cc.Invoke(ctx, "/cliente_logistica.SeguimientoService/Seguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeguimientoServiceServer is the server API for SeguimientoService service.
type SeguimientoServiceServer interface {
	Seguimiento(context.Context, *SeguimientoRequest) (*SeguimientoResponse, error)
}

// UnimplementedSeguimientoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSeguimientoServiceServer struct {
}

func (*UnimplementedSeguimientoServiceServer) Seguimiento(context.Context, *SeguimientoRequest) (*SeguimientoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seguimiento not implemented")
}

func RegisterSeguimientoServiceServer(s *grpc.Server, srv SeguimientoServiceServer) {
	s.RegisterService(&_SeguimientoService_serviceDesc, srv)
}

func _SeguimientoService_Seguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeguimientoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeguimientoServiceServer).Seguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente_logistica.SeguimientoService/Seguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeguimientoServiceServer).Seguimiento(ctx, req.(*SeguimientoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeguimientoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cliente_logistica.SeguimientoService",
	HandlerType: (*SeguimientoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Seguimiento",
			Handler:    _SeguimientoService_Seguimiento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cliente_logistica.proto",
}
