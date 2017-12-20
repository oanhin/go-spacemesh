// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// data common to all messages - Top level msg format
type CommonMessageData struct {
	SessionId []byte `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload   []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *CommonMessageData) Reset()                    { *m = CommonMessageData{} }
func (m *CommonMessageData) String() string            { return proto.CompactTextString(m) }
func (*CommonMessageData) ProtoMessage()               {}
func (*CommonMessageData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommonMessageData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *CommonMessageData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Handshake protocol data used for both request and response - sent unencrypted on the wire
type HandshakeData struct {
	SessionId  []byte `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload    []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Protocol   string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	NodePubKey []byte `protobuf:"bytes,4,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Iv         []byte `protobuf:"bytes,5,opt,name=iv,proto3" json:"iv,omitempty"`
	PubKey     []byte `protobuf:"bytes,6,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Hmac       []byte `protobuf:"bytes,7,opt,name=hmac,proto3" json:"hmac,omitempty"`
	TcpAddress string `protobuf:"bytes,8,opt,name=tcpAddress" json:"tcpAddress,omitempty"`
	Sign       string `protobuf:"bytes,9,opt,name=sign" json:"sign,omitempty"`
}

func (m *HandshakeData) Reset()                    { *m = HandshakeData{} }
func (m *HandshakeData) String() string            { return proto.CompactTextString(m) }
func (*HandshakeData) ProtoMessage()               {}
func (*HandshakeData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *HandshakeData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *HandshakeData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HandshakeData) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *HandshakeData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *HandshakeData) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *HandshakeData) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *HandshakeData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

func (m *HandshakeData) GetTcpAddress() string {
	if m != nil {
		return m.TcpAddress
	}
	return ""
}

func (m *HandshakeData) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

// used for protocol messages (non-handshake) - this is the decrypted CommonMessageData.payload
// it allows multiplexing back to higher level protocol
// data is here and not in CommonMessageData to limit leaked data on unencrypted connections
type ProtocolMessage struct {
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ProtocolMessage) Reset()                    { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string            { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()               {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ProtocolMessage) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	Protocol      string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	ReqId         []byte `protobuf:"bytes,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
	ClientVersion string `protobuf:"bytes,3,opt,name=clientVersion" json:"clientVersion,omitempty"`
	Timestamp     int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Gossip        bool   `protobuf:"varint,5,opt,name=gossip" json:"gossip,omitempty"`
	AuthPubKey    []byte `protobuf:"bytes,6,opt,name=authPubKey,proto3" json:"authPubKey,omitempty"`
	AuthorSign    string `protobuf:"bytes,7,opt,name=authorSign" json:"authorSign,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Metadata) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Metadata) GetReqId() []byte {
	if m != nil {
		return m.ReqId
	}
	return nil
}

func (m *Metadata) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *Metadata) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metadata) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *Metadata) GetAuthPubKey() []byte {
	if m != nil {
		return m.AuthPubKey
	}
	return nil
}

func (m *Metadata) GetAuthorSign() string {
	if m != nil {
		return m.AuthorSign
	}
	return ""
}

// example protocol
type PingReqData struct {
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Ping     string    `protobuf:"bytes,2,opt,name=ping" json:"ping,omitempty"`
}

func (m *PingReqData) Reset()                    { *m = PingReqData{} }
func (m *PingReqData) String() string            { return proto.CompactTextString(m) }
func (*PingReqData) ProtoMessage()               {}
func (*PingReqData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PingReqData) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PingReqData) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PingRespData struct {
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Pong     string    `protobuf:"bytes,2,opt,name=pong" json:"pong,omitempty"`
}

func (m *PingRespData) Reset()                    { *m = PingRespData{} }
func (m *PingRespData) String() string            { return proto.CompactTextString(m) }
func (*PingRespData) ProtoMessage()               {}
func (*PingRespData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *PingRespData) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PingRespData) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonMessageData)(nil), "pb.CommonMessageData")
	proto.RegisterType((*HandshakeData)(nil), "pb.HandshakeData")
	proto.RegisterType((*ProtocolMessage)(nil), "pb.ProtocolMessage")
	proto.RegisterType((*Metadata)(nil), "pb.Metadata")
	proto.RegisterType((*PingReqData)(nil), "pb.PingReqData")
	proto.RegisterType((*PingRespData)(nil), "pb.PingRespData")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0x87, 0x91, 0xfc, 0x4f, 0x1a, 0xdb, 0x2d, 0x5d, 0x4a, 0x59, 0x4a, 0x29, 0x46, 0xf4, 0xe0,
	0x93, 0x0f, 0xed, 0xb1, 0xa7, 0x36, 0x39, 0xc4, 0x38, 0x06, 0xa1, 0x40, 0x0e, 0xb9, 0xad, 0xa4,
	0x45, 0x5e, 0x62, 0xed, 0xae, 0xb5, 0x6b, 0x83, 0xdf, 0x34, 0xef, 0x91, 0x17, 0x08, 0x3b, 0x92,
	0x6c, 0x39, 0xa7, 0x90, 0xdc, 0x66, 0xbe, 0xd1, 0xfe, 0xb4, 0xf3, 0x49, 0x30, 0x2d, 0xb9, 0x31,
	0xac, 0xe0, 0x0b, 0x5d, 0x29, 0xab, 0x88, 0xaf, 0xd3, 0x68, 0x05, 0x5f, 0xae, 0x54, 0x59, 0x2a,
	0xb9, 0xae, 0x47, 0xd7, 0xcc, 0x32, 0xf2, 0x03, 0x42, 0xc3, 0x8d, 0x11, 0x4a, 0x2e, 0x73, 0xea,
	0xcd, 0xbc, 0xf9, 0x24, 0x39, 0x03, 0x42, 0x61, 0xa4, 0xd9, 0x71, 0xab, 0x58, 0x4e, 0x7d, 0x9c,
	0xb5, 0x6d, 0xf4, 0xec, 0xc1, 0xf4, 0x86, 0xc9, 0xdc, 0x6c, 0xd8, 0xe3, 0x87, 0x92, 0xc8, 0x77,
	0x08, 0xf0, 0x8e, 0x99, 0xda, 0xd2, 0xde, 0xcc, 0x9b, 0x87, 0xc9, 0xa9, 0x27, 0x3f, 0x01, 0xa4,
	0xca, 0x79, 0xbc, 0x4f, 0x57, 0xfc, 0x48, 0xfb, 0x78, 0xb0, 0x43, 0xc8, 0x27, 0xf0, 0xc5, 0x81,
	0x0e, 0x90, 0xfb, 0xe2, 0x40, 0xbe, 0xc1, 0x50, 0xd7, 0xcf, 0x0e, 0x91, 0x35, 0x1d, 0x21, 0xd0,
	0xdf, 0x94, 0x2c, 0xa3, 0x23, 0xa4, 0x58, 0xbb, 0x6c, 0x9b, 0xe9, 0x7f, 0x79, 0x5e, 0x71, 0x63,
	0x68, 0x80, 0x6f, 0xee, 0x10, 0x77, 0xc6, 0x88, 0x42, 0xd2, 0x10, 0x27, 0x58, 0x47, 0x7f, 0xe1,
	0x73, 0xdc, 0xdc, 0xad, 0x91, 0x48, 0xe6, 0x10, 0x94, 0xdc, 0xb2, 0x9c, 0x59, 0x86, 0x5b, 0x8f,
	0x7f, 0x4f, 0x16, 0x3a, 0x5d, 0xac, 0x1b, 0x96, 0x9c, 0xa6, 0xd1, 0x93, 0x07, 0x41, 0x8b, 0x2f,
	0xb6, 0xf6, 0x5e, 0x6d, 0xfd, 0x15, 0x06, 0x15, 0xdf, 0x2d, 0x5b, 0x53, 0x75, 0x43, 0x7e, 0xc1,
	0x34, 0xdb, 0x0a, 0x2e, 0xed, 0x3d, 0xaf, 0x9c, 0xd4, 0x46, 0xd6, 0x25, 0x74, 0x5f, 0xc1, 0x8a,
	0x92, 0x1b, 0xcb, 0x4a, 0x8d, 0xc2, 0x7a, 0xc9, 0x19, 0x38, 0x3f, 0x85, 0x32, 0x46, 0x68, 0x74,
	0x16, 0x24, 0x4d, 0xe7, 0x5c, 0xb0, 0xbd, 0xdd, 0xc4, 0x5d, 0x77, 0x1d, 0xd2, 0xce, 0x55, 0x75,
	0xe7, 0x8c, 0x8c, 0x6a, 0x57, 0x67, 0x12, 0xad, 0x60, 0x1c, 0x0b, 0x59, 0x24, 0x7c, 0x87, 0xbf,
	0xc2, 0x9b, 0x9d, 0x38, 0xc9, 0x5a, 0xc8, 0x02, 0x37, 0x0d, 0x13, 0xac, 0xa3, 0x5b, 0x98, 0xd4,
	0x61, 0x46, 0xbf, 0x23, 0x4d, 0x75, 0xd2, 0x94, 0x2c, 0xfe, 0xf7, 0x1f, 0x7c, 0x9d, 0xa6, 0x43,
	0x94, 0xfb, 0xe7, 0x25, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xf3, 0xcd, 0xf8, 0x17, 0x03, 0x00, 0x00,
}
