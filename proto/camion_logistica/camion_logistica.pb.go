// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: proto/camion_logistica.proto

package camion_logistica

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

type CamionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CamionRequest) Reset() {
	*x = CamionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_camion_logistica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CamionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CamionRequest) ProtoMessage() {}

func (x *CamionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_camion_logistica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CamionRequest.ProtoReflect.Descriptor instead.
func (*CamionRequest) Descriptor() ([]byte, []int) {
	return file_proto_camion_logistica_proto_rawDescGZIP(), []int{0}
}

func (x *CamionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CamionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algo string `protobuf:"bytes,1,opt,name=algo,proto3" json:"algo,omitempty"`
}

func (x *CamionResponse) Reset() {
	*x = CamionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_camion_logistica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CamionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CamionResponse) ProtoMessage() {}

func (x *CamionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_camion_logistica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CamionResponse.ProtoReflect.Descriptor instead.
func (*CamionResponse) Descriptor() ([]byte, []int) {
	return file_proto_camion_logistica_proto_rawDescGZIP(), []int{1}
}

func (x *CamionResponse) GetAlgo() string {
	if x != nil {
		return x.Algo
	}
	return ""
}

var File_proto_camion_logistica_proto protoreflect.FileDescriptor

var file_proto_camion_logistica_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x61,
	0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x6c, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f,
	0x32, 0x64, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e,
	0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x63, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_camion_logistica_proto_rawDescOnce sync.Once
	file_proto_camion_logistica_proto_rawDescData = file_proto_camion_logistica_proto_rawDesc
)

func file_proto_camion_logistica_proto_rawDescGZIP() []byte {
	file_proto_camion_logistica_proto_rawDescOnce.Do(func() {
		file_proto_camion_logistica_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_camion_logistica_proto_rawDescData)
	})
	return file_proto_camion_logistica_proto_rawDescData
}

var file_proto_camion_logistica_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_camion_logistica_proto_goTypes = []interface{}{
	(*CamionRequest)(nil),  // 0: cliente_logistica.CamionRequest
	(*CamionResponse)(nil), // 1: cliente_logistica.CamionResponse
}
var file_proto_camion_logistica_proto_depIdxs = []int32{
	0, // 0: cliente_logistica.CamionService.Camion:input_type -> cliente_logistica.CamionRequest
	1, // 1: cliente_logistica.CamionService.Camion:output_type -> cliente_logistica.CamionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_camion_logistica_proto_init() }
func file_proto_camion_logistica_proto_init() {
	if File_proto_camion_logistica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_camion_logistica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CamionRequest); i {
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
		file_proto_camion_logistica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CamionResponse); i {
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
			RawDescriptor: file_proto_camion_logistica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_camion_logistica_proto_goTypes,
		DependencyIndexes: file_proto_camion_logistica_proto_depIdxs,
		MessageInfos:      file_proto_camion_logistica_proto_msgTypes,
	}.Build()
	File_proto_camion_logistica_proto = out.File
	file_proto_camion_logistica_proto_rawDesc = nil
	file_proto_camion_logistica_proto_goTypes = nil
	file_proto_camion_logistica_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CamionServiceClient is the client API for CamionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CamionServiceClient interface {
	Camion(ctx context.Context, opts ...grpc.CallOption) (CamionService_CamionClient, error)
}

type camionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCamionServiceClient(cc grpc.ClientConnInterface) CamionServiceClient {
	return &camionServiceClient{cc}
}

func (c *camionServiceClient) Camion(ctx context.Context, opts ...grpc.CallOption) (CamionService_CamionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CamionService_serviceDesc.Streams[0], "/cliente_logistica.CamionService/Camion", opts...)
	if err != nil {
		return nil, err
	}
	x := &camionServiceCamionClient{stream}
	return x, nil
}

type CamionService_CamionClient interface {
	Send(*CamionRequest) error
	Recv() (*CamionResponse, error)
	grpc.ClientStream
}

type camionServiceCamionClient struct {
	grpc.ClientStream
}

func (x *camionServiceCamionClient) Send(m *CamionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *camionServiceCamionClient) Recv() (*CamionResponse, error) {
	m := new(CamionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CamionServiceServer is the server API for CamionService service.
type CamionServiceServer interface {
	Camion(CamionService_CamionServer) error
}

// UnimplementedCamionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCamionServiceServer struct {
}

func (*UnimplementedCamionServiceServer) Camion(CamionService_CamionServer) error {
	return status.Errorf(codes.Unimplemented, "method Camion not implemented")
}

func RegisterCamionServiceServer(s *grpc.Server, srv CamionServiceServer) {
	s.RegisterService(&_CamionService_serviceDesc, srv)
}

func _CamionService_Camion_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CamionServiceServer).Camion(&camionServiceCamionServer{stream})
}

type CamionService_CamionServer interface {
	Send(*CamionResponse) error
	Recv() (*CamionRequest, error)
	grpc.ServerStream
}

type camionServiceCamionServer struct {
	grpc.ServerStream
}

func (x *camionServiceCamionServer) Send(m *CamionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *camionServiceCamionServer) Recv() (*CamionRequest, error) {
	m := new(CamionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CamionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cliente_logistica.CamionService",
	HandlerType: (*CamionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Camion",
			Handler:       _CamionService_Camion_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/camion_logistica.proto",
}
