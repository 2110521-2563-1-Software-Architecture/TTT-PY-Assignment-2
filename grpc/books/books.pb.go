package books

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Book struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BookList struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *BookList) Reset()                    { *m = BookList{} }
func (m *BookList) String() string            { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()               {}
func (*BookList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BookList) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type BookIdRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *BookIdRequest) Reset()                    { *m = BookIdRequest{} }
func (m *BookIdRequest) String() string            { return proto.CompactTextString(m) }
func (*BookIdRequest) ProtoMessage()               {}
func (*BookIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Empty)(nil), "books.Empty")
	proto.RegisterType((*Book)(nil), "books.Book")
	proto.RegisterType((*BookList)(nil), "books.BookList")
	proto.RegisterType((*BookIdRequest)(nil), "books.BookIdRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error)
	Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error)
	Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := grpc.Invoke(ctx, "/books.BookService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/books.BookService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BookService_serviceDesc.Streams[0], c.cc, "/books.BookService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_WatchClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceWatchClient struct {
	grpc.ClientStream
}

func (x *bookServiceWatchClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BookService service

type BookServiceServer interface {
	List(context.Context, *Empty) (*BookList, error)
	Insert(context.Context, *Book) (*Empty, error)
	Get(context.Context, *BookIdRequest) (*Book, error)
	Delete(context.Context, *BookIdRequest) (*Empty, error)
	Watch(*Empty, BookService_WatchServer) error
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Insert(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).Watch(m, &bookServiceWatchServer{stream})
}

type BookService_WatchServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceWatchServer struct {
	grpc.ServerStream
}

func (x *bookServiceWatchServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BookService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _BookService_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _BookService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("books.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x9b, 0xa6, 0x1b, 0x75, 0x62, 0x15, 0x1e, 0x1e, 0x42, 0xbc, 0xc8, 0x82, 0xb5, 0x78,
	0x08, 0x5a, 0xbf, 0x81, 0x28, 0x52, 0xf0, 0xa4, 0x07, 0xcf, 0xfd, 0xf3, 0xa0, 0xc1, 0xea, 0xd6,
	0xcd, 0xab, 0xe0, 0xc7, 0xf5, 0x9b, 0xb8, 0xbb, 0x15, 0xd9, 0x04, 0x3c, 0x85, 0x97, 0x99, 0xf9,
	0xcd, 0xb0, 0xc8, 0xe7, 0xc6, 0xbc, 0x36, 0xd5, 0xc6, 0x1a, 0x31, 0xa4, 0xc2, 0xa1, 0xf7, 0xa0,
	0xee, 0xdf, 0x36, 0xf2, 0xa5, 0xaf, 0x31, 0xb8, 0x75, 0x7f, 0x08, 0xe8, 0xd7, 0xcb, 0x22, 0x39,
	0x4b, 0xc6, 0x8a, 0x86, 0x50, 0x52, 0xcb, 0x9a, 0x8b, 0xbe, 0x3b, 0x0f, 0xe8, 0x08, 0xd9, 0x6c,
	0x2b, 0x2b, 0x63, 0x8b, 0xd4, 0xdf, 0x7a, 0x84, 0x7d, 0x1f, 0x79, 0xac, 0x1b, 0xa1, 0x12, 0x3b,
	0xa0, 0x4b, 0xa6, 0xe3, 0x7c, 0x92, 0x57, 0xbb, 0x2e, 0xaf, 0xeb, 0x53, 0x0c, 0xfd, 0x77, 0xba,
	0x7c, 0xe2, 0x8f, 0x2d, 0x3b, 0x73, 0xd4, 0x31, 0xf9, 0x4e, 0x90, 0x7b, 0xf5, 0x99, 0xed, 0x67,
	0xbd, 0x60, 0xba, 0xc0, 0x20, 0x00, 0x0f, 0x7f, 0x09, 0x61, 0x5d, 0x79, 0x1c, 0xf1, 0xbc, 0xac,
	0x7b, 0x74, 0x8e, 0x6c, 0xfa, 0xde, 0xb0, 0x15, 0x8a, 0xcb, 0xca, 0x56, 0xce, 0xd9, 0x2e, 0x91,
	0x3e, 0xb0, 0xd0, 0x49, 0xe4, 0xf9, 0x1b, 0x52, 0xb6, 0x66, 0xf6, 0xa8, 0x42, 0x76, 0xc7, 0x6b,
	0x16, 0xfe, 0xc7, 0xde, 0x65, 0x8f, 0xa0, 0x5e, 0x66, 0xb2, 0x58, 0x75, 0xc6, 0xb6, 0xa9, 0x57,
	0xc9, 0x3c, 0x0b, 0x4f, 0x7e, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x70, 0xef, 0x1d, 0x81,
	0x01, 0x00, 0x00,
}
