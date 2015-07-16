// Code generated by protoc-gen-go.
// source: dc-smartMerge.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	dc-smartMerge.proto

It has these top-level messages:
	State
	Blueprint
	ReadReply
	WriteReply
	ReadRequest
	WriteRequest
	WriteNRequest
	ReadNRequest
	WriteNReply
	ReadNReply
	NewCur
	NewCurReply
	AdvRead
	AdvReadReply
	AdvWriteS
	AdvWriteSReply
	AdvWriteN
	AdvWriteNReply
	LAProposal
	LAReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type State struct {
	Value     []byte `protobuf:"bytes,1,opt,proto3" json:"Value,omitempty"`
	Timestamp int32  `protobuf:"varint,2,opt" json:"Timestamp,omitempty"`
	Writer    uint32 `protobuf:"varint,3,opt" json:"Writer,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto1.CompactTextString(m) }
func (*State) ProtoMessage()    {}

type Blueprint struct {
	Add []uint32 `protobuf:"varint,1,rep" json:"Add,omitempty"`
	Rem []uint32 `protobuf:"varint,2,rep" json:"Rem,omitempty"`
}

func (m *Blueprint) Reset()         { *m = Blueprint{} }
func (m *Blueprint) String() string { return proto1.CompactTextString(m) }
func (*Blueprint) ProtoMessage()    {}

type ReadReply struct {
	State *State     `protobuf:"bytes,1,opt" json:"State,omitempty"`
	Cur   *Blueprint `protobuf:"bytes,2,opt" json:"Cur,omitempty"`
}

func (m *ReadReply) Reset()         { *m = ReadReply{} }
func (m *ReadReply) String() string { return proto1.CompactTextString(m) }
func (*ReadReply) ProtoMessage()    {}

func (m *ReadReply) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ReadReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

type WriteReply struct {
	Cur *Blueprint `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
}

func (m *WriteReply) Reset()         { *m = WriteReply{} }
func (m *WriteReply) String() string { return proto1.CompactTextString(m) }
func (*WriteReply) ProtoMessage()    {}

func (m *WriteReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

type ReadRequest struct {
	CurC uint32 `protobuf:"varint,1,opt" json:"CurC,omitempty"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto1.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}

type WriteRequest struct {
	State *State `protobuf:"bytes,1,opt" json:"State,omitempty"`
	CurC  uint32 `protobuf:"varint,2,opt" json:"CurC,omitempty"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto1.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

func (m *WriteRequest) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type WriteNRequest struct {
	CurC uint32     `protobuf:"varint,1,opt" json:"CurC,omitempty"`
	Next *Blueprint `protobuf:"bytes,2,opt" json:"Next,omitempty"`
}

func (m *WriteNRequest) Reset()         { *m = WriteNRequest{} }
func (m *WriteNRequest) String() string { return proto1.CompactTextString(m) }
func (*WriteNRequest) ProtoMessage()    {}

func (m *WriteNRequest) GetNext() *Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

type ReadNRequest struct {
	CurC uint32 `protobuf:"varint,1,opt" json:"CurC,omitempty"`
}

func (m *ReadNRequest) Reset()         { *m = ReadNRequest{} }
func (m *ReadNRequest) String() string { return proto1.CompactTextString(m) }
func (*ReadNRequest) ProtoMessage()    {}

type WriteNReply struct {
	Cur *Blueprint `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
}

func (m *WriteNReply) Reset()         { *m = WriteNReply{} }
func (m *WriteNReply) String() string { return proto1.CompactTextString(m) }
func (*WriteNReply) ProtoMessage()    {}

func (m *WriteNReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

type ReadNReply struct {
	Cur  *Blueprint   `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
	Next []*Blueprint `protobuf:"bytes,2,rep" json:"Next,omitempty"`
}

func (m *ReadNReply) Reset()         { *m = ReadNReply{} }
func (m *ReadNReply) String() string { return proto1.CompactTextString(m) }
func (*ReadNReply) ProtoMessage()    {}

func (m *ReadNReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

func (m *ReadNReply) GetNext() []*Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

type NewCur struct {
	Cur  *Blueprint `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
	CurC uint32     `protobuf:"varint,2,opt" json:"CurC,omitempty"`
}

func (m *NewCur) Reset()         { *m = NewCur{} }
func (m *NewCur) String() string { return proto1.CompactTextString(m) }
func (*NewCur) ProtoMessage()    {}

func (m *NewCur) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

type NewCurReply struct {
	New bool `protobuf:"varint,1,opt" json:"New,omitempty"`
}

func (m *NewCurReply) Reset()         { *m = NewCurReply{} }
func (m *NewCurReply) String() string { return proto1.CompactTextString(m) }
func (*NewCurReply) ProtoMessage()    {}

type AdvRead struct {
	CurC uint32 `protobuf:"varint,1,opt" json:"CurC,omitempty"`
}

func (m *AdvRead) Reset()         { *m = AdvRead{} }
func (m *AdvRead) String() string { return proto1.CompactTextString(m) }
func (*AdvRead) ProtoMessage()    {}

type AdvReadReply struct {
	State *State       `protobuf:"bytes,1,opt" json:"State,omitempty"`
	Cur   *Blueprint   `protobuf:"bytes,2,opt" json:"Cur,omitempty"`
	Next  []*Blueprint `protobuf:"bytes,3,rep" json:"Next,omitempty"`
}

func (m *AdvReadReply) Reset()         { *m = AdvReadReply{} }
func (m *AdvReadReply) String() string { return proto1.CompactTextString(m) }
func (*AdvReadReply) ProtoMessage()    {}

func (m *AdvReadReply) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *AdvReadReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

func (m *AdvReadReply) GetNext() []*Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

type AdvWriteS struct {
	State *State `protobuf:"bytes,1,opt" json:"State,omitempty"`
	CurC  uint32 `protobuf:"varint,2,opt" json:"CurC,omitempty"`
}

func (m *AdvWriteS) Reset()         { *m = AdvWriteS{} }
func (m *AdvWriteS) String() string { return proto1.CompactTextString(m) }
func (*AdvWriteS) ProtoMessage()    {}

func (m *AdvWriteS) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type AdvWriteSReply struct {
	Cur  *Blueprint   `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
	Next []*Blueprint `protobuf:"bytes,2,rep" json:"Next,omitempty"`
}

func (m *AdvWriteSReply) Reset()         { *m = AdvWriteSReply{} }
func (m *AdvWriteSReply) String() string { return proto1.CompactTextString(m) }
func (*AdvWriteSReply) ProtoMessage()    {}

func (m *AdvWriteSReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

func (m *AdvWriteSReply) GetNext() []*Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

type AdvWriteN struct {
	CurC uint32     `protobuf:"varint,1,opt" json:"CurC,omitempty"`
	Next *Blueprint `protobuf:"bytes,2,opt" json:"Next,omitempty"`
}

func (m *AdvWriteN) Reset()         { *m = AdvWriteN{} }
func (m *AdvWriteN) String() string { return proto1.CompactTextString(m) }
func (*AdvWriteN) ProtoMessage()    {}

func (m *AdvWriteN) GetNext() *Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

type AdvWriteNReply struct {
	Cur     *Blueprint   `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
	State   *State       `protobuf:"bytes,2,opt" json:"State,omitempty"`
	Next    []*Blueprint `protobuf:"bytes,3,rep" json:"Next,omitempty"`
	LAState *Blueprint   `protobuf:"bytes,4,opt" json:"LAState,omitempty"`
}

func (m *AdvWriteNReply) Reset()         { *m = AdvWriteNReply{} }
func (m *AdvWriteNReply) String() string { return proto1.CompactTextString(m) }
func (*AdvWriteNReply) ProtoMessage()    {}

func (m *AdvWriteNReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

func (m *AdvWriteNReply) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *AdvWriteNReply) GetNext() []*Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *AdvWriteNReply) GetLAState() *Blueprint {
	if m != nil {
		return m.LAState
	}
	return nil
}

type LAProposal struct {
	CurC uint32     `protobuf:"varint,1,opt" json:"CurC,omitempty"`
	Prop *Blueprint `protobuf:"bytes,2,opt" json:"Prop,omitempty"`
}

func (m *LAProposal) Reset()         { *m = LAProposal{} }
func (m *LAProposal) String() string { return proto1.CompactTextString(m) }
func (*LAProposal) ProtoMessage()    {}

func (m *LAProposal) GetProp() *Blueprint {
	if m != nil {
		return m.Prop
	}
	return nil
}

type LAReply struct {
	Cur     *Blueprint   `protobuf:"bytes,1,opt" json:"Cur,omitempty"`
	LAState *Blueprint   `protobuf:"bytes,2,opt" json:"LAState,omitempty"`
	Next    []*Blueprint `protobuf:"bytes,3,rep" json:"Next,omitempty"`
}

func (m *LAReply) Reset()         { *m = LAReply{} }
func (m *LAReply) String() string { return proto1.CompactTextString(m) }
func (*LAReply) ProtoMessage()    {}

func (m *LAReply) GetCur() *Blueprint {
	if m != nil {
		return m.Cur
	}
	return nil
}

func (m *LAReply) GetLAState() *Blueprint {
	if m != nil {
		return m.LAState
	}
	return nil
}

func (m *LAReply) GetNext() []*Blueprint {
	if m != nil {
		return m.Next
	}
	return nil
}

func init() {
}

// Client API for Register service

type RegisterClient interface {
	ReadS(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
	WriteS(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error)
	ReadN(ctx context.Context, in *ReadNRequest, opts ...grpc.CallOption) (*ReadNReply, error)
	WriteN(ctx context.Context, in *WriteNRequest, opts ...grpc.CallOption) (*WriteNReply, error)
	SetCur(ctx context.Context, in *NewCur, opts ...grpc.CallOption) (*NewCurReply, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) ReadS(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := grpc.Invoke(ctx, "/proto.Register/ReadS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) WriteS(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error) {
	out := new(WriteReply)
	err := grpc.Invoke(ctx, "/proto.Register/WriteS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) ReadN(ctx context.Context, in *ReadNRequest, opts ...grpc.CallOption) (*ReadNReply, error) {
	out := new(ReadNReply)
	err := grpc.Invoke(ctx, "/proto.Register/ReadN", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) WriteN(ctx context.Context, in *WriteNRequest, opts ...grpc.CallOption) (*WriteNReply, error) {
	out := new(WriteNReply)
	err := grpc.Invoke(ctx, "/proto.Register/WriteN", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) SetCur(ctx context.Context, in *NewCur, opts ...grpc.CallOption) (*NewCurReply, error) {
	out := new(NewCurReply)
	err := grpc.Invoke(ctx, "/proto.Register/SetCur", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterServer interface {
	ReadS(context.Context, *ReadRequest) (*ReadReply, error)
	WriteS(context.Context, *WriteRequest) (*WriteReply, error)
	ReadN(context.Context, *ReadNRequest) (*ReadNReply, error)
	WriteN(context.Context, *WriteNRequest) (*WriteNReply, error)
	SetCur(context.Context, *NewCur) (*NewCurReply, error)
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_ReadS_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ReadRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RegisterServer).ReadS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Register_WriteS_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(WriteRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RegisterServer).WriteS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Register_ReadN_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ReadNRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RegisterServer).ReadN(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Register_WriteN_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(WriteNRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RegisterServer).WriteN(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Register_SetCur_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NewCur)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RegisterServer).SetCur(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadS",
			Handler:    _Register_ReadS_Handler,
		},
		{
			MethodName: "WriteS",
			Handler:    _Register_WriteS_Handler,
		},
		{
			MethodName: "ReadN",
			Handler:    _Register_ReadN_Handler,
		},
		{
			MethodName: "WriteN",
			Handler:    _Register_WriteN_Handler,
		},
		{
			MethodName: "SetCur",
			Handler:    _Register_SetCur_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for AdvRegister service

type AdvRegisterClient interface {
	AReadS(ctx context.Context, in *AdvRead, opts ...grpc.CallOption) (*AdvReadReply, error)
	AWriteS(ctx context.Context, in *AdvWriteS, opts ...grpc.CallOption) (*AdvWriteSReply, error)
	AWriteN(ctx context.Context, in *AdvWriteN, opts ...grpc.CallOption) (*AdvWriteNReply, error)
	SetCur(ctx context.Context, in *NewCur, opts ...grpc.CallOption) (*NewCurReply, error)
	LAProp(ctx context.Context, in *LAProposal, opts ...grpc.CallOption) (*LAReply, error)
}

type advRegisterClient struct {
	cc *grpc.ClientConn
}

func NewAdvRegisterClient(cc *grpc.ClientConn) AdvRegisterClient {
	return &advRegisterClient{cc}
}

func (c *advRegisterClient) AReadS(ctx context.Context, in *AdvRead, opts ...grpc.CallOption) (*AdvReadReply, error) {
	out := new(AdvReadReply)
	err := grpc.Invoke(ctx, "/proto.AdvRegister/AReadS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advRegisterClient) AWriteS(ctx context.Context, in *AdvWriteS, opts ...grpc.CallOption) (*AdvWriteSReply, error) {
	out := new(AdvWriteSReply)
	err := grpc.Invoke(ctx, "/proto.AdvRegister/AWriteS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advRegisterClient) AWriteN(ctx context.Context, in *AdvWriteN, opts ...grpc.CallOption) (*AdvWriteNReply, error) {
	out := new(AdvWriteNReply)
	err := grpc.Invoke(ctx, "/proto.AdvRegister/AWriteN", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advRegisterClient) SetCur(ctx context.Context, in *NewCur, opts ...grpc.CallOption) (*NewCurReply, error) {
	out := new(NewCurReply)
	err := grpc.Invoke(ctx, "/proto.AdvRegister/SetCur", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advRegisterClient) LAProp(ctx context.Context, in *LAProposal, opts ...grpc.CallOption) (*LAReply, error) {
	out := new(LAReply)
	err := grpc.Invoke(ctx, "/proto.AdvRegister/LAProp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdvRegister service

type AdvRegisterServer interface {
	AReadS(context.Context, *AdvRead) (*AdvReadReply, error)
	AWriteS(context.Context, *AdvWriteS) (*AdvWriteSReply, error)
	AWriteN(context.Context, *AdvWriteN) (*AdvWriteNReply, error)
	SetCur(context.Context, *NewCur) (*NewCurReply, error)
	LAProp(context.Context, *LAProposal) (*LAReply, error)
}

func RegisterAdvRegisterServer(s *grpc.Server, srv AdvRegisterServer) {
	s.RegisterService(&_AdvRegister_serviceDesc, srv)
}

func _AdvRegister_AReadS_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(AdvRead)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AdvRegisterServer).AReadS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AdvRegister_AWriteS_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(AdvWriteS)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AdvRegisterServer).AWriteS(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AdvRegister_AWriteN_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(AdvWriteN)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AdvRegisterServer).AWriteN(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AdvRegister_SetCur_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(NewCur)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AdvRegisterServer).SetCur(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AdvRegister_LAProp_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(LAProposal)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AdvRegisterServer).LAProp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AdvRegister_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdvRegister",
	HandlerType: (*AdvRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AReadS",
			Handler:    _AdvRegister_AReadS_Handler,
		},
		{
			MethodName: "AWriteS",
			Handler:    _AdvRegister_AWriteS_Handler,
		},
		{
			MethodName: "AWriteN",
			Handler:    _AdvRegister_AWriteN_Handler,
		},
		{
			MethodName: "SetCur",
			Handler:    _AdvRegister_SetCur_Handler,
		},
		{
			MethodName: "LAProp",
			Handler:    _AdvRegister_LAProp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
