// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/chains/ethereum/config/config.proto

package ethereum

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ChainConfig struct {
	ChainId    string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	EthChainId int64  `protobuf:"varint,2,opt,name=eth_chain_id,json=ethChainId,proto3" json:"eth_chain_id,omitempty"`
	RpcAddr    string `protobuf:"bytes,3,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	// use for relayer
	HdwMnemonic           string `protobuf:"bytes,4,opt,name=hdw_mnemonic,json=hdwMnemonic,proto3" json:"hdw_mnemonic,omitempty"`
	HdwPath               string `protobuf:"bytes,5,opt,name=hdw_path,json=hdwPath,proto3" json:"hdw_path,omitempty"`
	IbcAddress            string `protobuf:"bytes,6,opt,name=ibc_address,json=ibcAddress,proto3" json:"ibc_address,omitempty"`
	InitialSendCheckpoint uint64 `protobuf:"varint,7,opt,name=initial_send_checkpoint,json=initialSendCheckpoint,proto3" json:"initial_send_checkpoint,omitempty"`
	InitialRecvCheckpoint uint64 `protobuf:"varint,8,opt,name=initial_recv_checkpoint,json=initialRecvCheckpoint,proto3" json:"initial_recv_checkpoint,omitempty"`
	EnableDebugTrace      bool   `protobuf:"varint,9,opt,name=enable_debug_trace,json=enableDebugTrace,proto3" json:"enable_debug_trace,omitempty"`
	AverageBlockTimeMsec  uint64 `protobuf:"varint,10,opt,name=average_block_time_msec,json=averageBlockTimeMsec,proto3" json:"average_block_time_msec,omitempty"`
	MaxRetryForInclusion  uint64 `protobuf:"varint,11,opt,name=max_retry_for_inclusion,json=maxRetryForInclusion,proto3" json:"max_retry_for_inclusion,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a57ab2f9f14837, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.ethereum.config.ChainConfig")
}

func init() {
	proto.RegisterFile("relayer/chains/ethereum/config/config.proto", fileDescriptor_a8a57ab2f9f14837)
}

var fileDescriptor_a8a57ab2f9f14837 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x80, 0x1b, 0x3a, 0xb6, 0xce, 0xdd, 0x01, 0x45, 0x43, 0xcb, 0x38, 0x84, 0xc0, 0xa9, 0x12,
	0x34, 0x39, 0x20, 0xb8, 0xb3, 0x20, 0xa4, 0x1d, 0x26, 0xa1, 0xb0, 0x13, 0x17, 0xcb, 0xb1, 0xff,
	0xc5, 0x56, 0x63, 0x3b, 0xb2, 0xdd, 0x76, 0x7b, 0x0b, 0x1e, 0x6b, 0xc7, 0x1e, 0x39, 0x42, 0xfb,
	0x22, 0xc8, 0x4e, 0x56, 0x75, 0xa7, 0xc4, 0xfe, 0xbe, 0xef, 0xcf, 0x21, 0x3f, 0xfa, 0x60, 0xa0,
	0x25, 0x0f, 0x60, 0x0a, 0xca, 0x89, 0x50, 0xb6, 0x00, 0xc7, 0xc1, 0xc0, 0x52, 0x16, 0x54, 0xab,
	0x3b, 0xd1, 0x0c, 0x8f, 0xbc, 0x33, 0xda, 0xe9, 0x38, 0x1d, 0xe4, 0xbc, 0x97, 0xf3, 0x27, 0x39,
	0xef, 0xad, 0x37, 0xe7, 0x8d, 0x6e, 0x74, 0x50, 0x0b, 0xff, 0xd6, 0x57, 0xef, 0x37, 0x63, 0x34,
	0x2d, 0x7d, 0x50, 0x06, 0x2b, 0xbe, 0x44, 0x93, 0xd0, 0x63, 0xc1, 0x92, 0x28, 0x8b, 0x66, 0xa7,
	0xd5, 0x49, 0x38, 0x5f, 0xb3, 0x38, 0x43, 0x67, 0xe0, 0x38, 0xde, 0xe3, 0x17, 0x59, 0x34, 0x1b,
	0x57, 0x08, 0x1c, 0x2f, 0x07, 0xe3, 0x12, 0x4d, 0x4c, 0x47, 0x31, 0x61, 0xcc, 0x24, 0xe3, 0x3e,
	0x36, 0x1d, 0xfd, 0xca, 0x98, 0x89, 0xdf, 0xa1, 0x33, 0xce, 0xd6, 0x58, 0x2a, 0x90, 0x5a, 0x09,
	0x9a, 0x1c, 0x05, 0x3c, 0xe5, 0x6c, 0x7d, 0x33, 0x5c, 0xf9, 0xda, 0x2b, 0x1d, 0x71, 0x3c, 0x79,
	0xd9, 0xd7, 0x9c, 0xad, 0x7f, 0x10, 0xc7, 0xe3, 0xb7, 0x68, 0x2a, 0xea, 0x7e, 0x30, 0x58, 0x9b,
	0x1c, 0x07, 0x8a, 0x44, 0x1d, 0x66, 0x83, 0xb5, 0xf1, 0x17, 0x74, 0x21, 0x94, 0x70, 0x82, 0xb4,
	0xd8, 0x82, 0x62, 0x98, 0x72, 0xa0, 0x8b, 0x4e, 0x0b, 0xe5, 0x92, 0x93, 0x2c, 0x9a, 0x1d, 0x55,
	0xaf, 0x07, 0xfc, 0x13, 0x14, 0x2b, 0xf7, 0xf0, 0xb0, 0x33, 0x40, 0x57, 0x87, 0xdd, 0xe4, 0x59,
	0x57, 0x01, 0x5d, 0x1d, 0x74, 0x1f, 0x51, 0x0c, 0x8a, 0xd4, 0x2d, 0x60, 0x06, 0xf5, 0xb2, 0xc1,
	0xce, 0x10, 0x0a, 0xc9, 0x69, 0x16, 0xcd, 0x26, 0xd5, 0xab, 0x9e, 0x7c, 0xf3, 0xe0, 0xd6, 0xdf,
	0xc7, 0x9f, 0xd1, 0x05, 0x59, 0x81, 0x21, 0x0d, 0xe0, 0xba, 0xd5, 0x74, 0x81, 0x9d, 0x90, 0x80,
	0xa5, 0x05, 0x9a, 0xa0, 0xf0, 0x95, 0xf3, 0x01, 0x5f, 0x79, 0x7a, 0x2b, 0x24, 0xdc, 0x58, 0xa0,
	0x3e, 0x93, 0xe4, 0x1e, 0x1b, 0x70, 0xe6, 0x01, 0xdf, 0x69, 0x83, 0x85, 0xa2, 0xed, 0xd2, 0x0a,
	0xad, 0x92, 0x69, 0x9f, 0x49, 0x72, 0x5f, 0x79, 0xfa, 0x5d, 0x9b, 0xeb, 0x27, 0x76, 0x45, 0x1e,
	0xff, 0xa5, 0xa3, 0xc7, 0x6d, 0x1a, 0x6d, 0xb6, 0x69, 0xf4, 0x77, 0x9b, 0x46, 0xbf, 0x77, 0xe9,
	0x68, 0xb3, 0x4b, 0x47, 0x7f, 0x76, 0xe9, 0xe8, 0x57, 0xd9, 0x08, 0xc7, 0x97, 0x75, 0x4e, 0xb5,
	0x2c, 0x18, 0x71, 0x24, 0xfc, 0xce, 0x96, 0xd4, 0xfb, 0xe5, 0x9a, 0x8b, 0x9a, 0xce, 0xc3, 0x2e,
	0xcd, 0x03, 0x2b, 0xba, 0x45, 0x53, 0x84, 0xf3, 0x5e, 0xa9, 0x8f, 0xc3, 0xf2, 0x7c, 0xfa, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x4d, 0xc3, 0x80, 0xa1, 0x02, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxRetryForInclusion != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxRetryForInclusion))
		i--
		dAtA[i] = 0x58
	}
	if m.AverageBlockTimeMsec != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.AverageBlockTimeMsec))
		i--
		dAtA[i] = 0x50
	}
	if m.EnableDebugTrace {
		i--
		if m.EnableDebugTrace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.InitialRecvCheckpoint != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.InitialRecvCheckpoint))
		i--
		dAtA[i] = 0x40
	}
	if m.InitialSendCheckpoint != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.InitialSendCheckpoint))
		i--
		dAtA[i] = 0x38
	}
	if len(m.IbcAddress) > 0 {
		i -= len(m.IbcAddress)
		copy(dAtA[i:], m.IbcAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.HdwPath) > 0 {
		i -= len(m.HdwPath)
		copy(dAtA[i:], m.HdwPath)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.HdwPath)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.HdwMnemonic) > 0 {
		i -= len(m.HdwMnemonic)
		copy(dAtA[i:], m.HdwMnemonic)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.HdwMnemonic)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RpcAddr) > 0 {
		i -= len(m.RpcAddr)
		copy(dAtA[i:], m.RpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RpcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EthChainId != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.EthChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.EthChainId != 0 {
		n += 1 + sovConfig(uint64(m.EthChainId))
	}
	l = len(m.RpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.HdwMnemonic)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.HdwPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.InitialSendCheckpoint != 0 {
		n += 1 + sovConfig(uint64(m.InitialSendCheckpoint))
	}
	if m.InitialRecvCheckpoint != 0 {
		n += 1 + sovConfig(uint64(m.InitialRecvCheckpoint))
	}
	if m.EnableDebugTrace {
		n += 2
	}
	if m.AverageBlockTimeMsec != 0 {
		n += 1 + sovConfig(uint64(m.AverageBlockTimeMsec))
	}
	if m.MaxRetryForInclusion != 0 {
		n += 1 + sovConfig(uint64(m.MaxRetryForInclusion))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthChainId", wireType)
			}
			m.EthChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HdwMnemonic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HdwMnemonic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HdwPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HdwPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSendCheckpoint", wireType)
			}
			m.InitialSendCheckpoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSendCheckpoint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialRecvCheckpoint", wireType)
			}
			m.InitialRecvCheckpoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialRecvCheckpoint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableDebugTrace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableDebugTrace = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTimeMsec", wireType)
			}
			m.AverageBlockTimeMsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTimeMsec |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetryForInclusion", wireType)
			}
			m.MaxRetryForInclusion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRetryForInclusion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
