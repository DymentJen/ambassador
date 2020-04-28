// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

package envoy_config_filter_network_mongo_proxy_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/datawire/ambassador/pkg/api/envoy/config/filter/fault/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

type MongoProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v2.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata  bool     `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoProxy) Reset()         { *m = MongoProxy{} }
func (m *MongoProxy) String() string { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()    {}
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d590dd12f767c61, []int{0}
}
func (m *MongoProxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MongoProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MongoProxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MongoProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoProxy.Merge(m, src)
}
func (m *MongoProxy) XXX_Size() int {
	return m.Size()
}
func (m *MongoProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MongoProxy proto.InternalMessageInfo

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *MongoProxy) GetEmitDynamicMetadata() bool {
	if m != nil {
		return m.EmitDynamicMetadata
	}
	return false
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v2.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto", fileDescriptor_4d590dd12f767c61)
}

var fileDescriptor_4d590dd12f767c61 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0xec, 0x40,
	0x14, 0x86, 0x99, 0xbd, 0xbb, 0xf7, 0xde, 0x9d, 0x2d, 0x84, 0x88, 0x18, 0x16, 0x0c, 0xc1, 0x2a,
	0x58, 0x64, 0x30, 0xdb, 0x58, 0x88, 0x45, 0x58, 0xac, 0x5c, 0x08, 0x79, 0x81, 0x30, 0x26, 0x93,
	0x30, 0x98, 0xcc, 0x84, 0xc9, 0xd9, 0xb8, 0x79, 0x0a, 0x2d, 0x7d, 0x24, 0xed, 0x7c, 0x04, 0xd9,
	0x47, 0xb0, 0xb4, 0x10, 0x99, 0x4c, 0xc4, 0x2d, 0xb6, 0xb0, 0x3b, 0x39, 0xdf, 0xf9, 0xcf, 0x9f,
	0xf3, 0x0f, 0xbe, 0x64, 0xa2, 0x95, 0x1d, 0x49, 0xa5, 0xc8, 0x79, 0x41, 0x72, 0x5e, 0x02, 0x53,
	0x44, 0x30, 0xb8, 0x97, 0xea, 0x8e, 0x54, 0x52, 0x14, 0x32, 0xa9, 0x95, 0xdc, 0x74, 0xa4, 0x0d,
	0x76, 0x3f, 0xfd, 0x5a, 0x49, 0x90, 0xd6, 0x59, 0xaf, 0xf6, 0x8d, 0xda, 0x37, 0x6a, 0x7f, 0x50,
	0xfb, 0xbb, 0xe3, 0x6d, 0x30, 0xf7, 0xf6, 0x39, 0xe5, 0x74, 0x5d, 0x82, 0xde, 0xdd, 0x17, 0x66,
	0xeb, 0xdc, 0x59, 0x67, 0x35, 0x25, 0x54, 0x08, 0x09, 0x14, 0xb8, 0x14, 0x0d, 0xa9, 0x78, 0xa1,
	0x28, 0xb0, 0x81, 0x1f, 0xb7, 0xb4, 0xe4, 0x19, 0x05, 0x46, 0xbe, 0x0b, 0x03, 0x4e, 0x5f, 0x10,
	0xc6, 0x2b, 0xed, 0x1a, 0x69, 0x53, 0xcb, 0xc3, 0xb3, 0x06, 0x28, 0x24, 0xb5, 0x62, 0x39, 0xdf,
	0xd8, 0xc8, 0x45, 0xde, 0x34, 0xfc, 0xf7, 0x11, 0x8e, 0xd5, 0xc8, 0x45, 0x31, 0xd6, 0x2c, 0xea,
	0x91, 0x75, 0x82, 0x31, 0x4d, 0x53, 0xd6, 0x34, 0x49, 0x29, 0x0b, 0x7b, 0xa4, 0x07, 0xe3, 0xa9,
	0xe9, 0xdc, 0xc8, 0xc2, 0xba, 0xc2, 0x93, 0x8c, 0x95, 0xb4, 0xb3, 0xff, 0xb8, 0xc8, 0x9b, 0x05,
	0x9e, 0xbf, 0xef, 0x6c, 0x73, 0x41, 0x1b, 0xf8, 0xd7, 0xba, 0x58, 0xea, 0xf9, 0xd8, 0xc8, 0xac,
	0x00, 0x1f, 0xb1, 0x8a, 0x43, 0x92, 0x75, 0x82, 0x56, 0x3c, 0x4d, 0x2a, 0x06, 0x34, 0xa3, 0x40,
	0xed, 0xb1, 0x8b, 0xbc, 0xff, 0xf1, 0xa1, 0x86, 0x4b, 0xc3, 0x56, 0x03, 0x0a, 0x1f, 0xd1, 0xf3,
	0xd6, 0x41, 0xaf, 0x5b, 0x07, 0xbd, 0x6d, 0x1d, 0xf4, 0xfe, 0xf4, 0xf9, 0x30, 0x39, 0xb7, 0x88,
	0x31, 0x66, 0x1b, 0x60, 0xa2, 0xd1, 0xc9, 0x0c, 0xe6, 0xcd, 0xfe, 0xd0, 0x17, 0xf8, 0x82, 0x4b,
	0xf3, 0xb3, 0xa6, 0xf3, 0xfb, 0xe7, 0x0a, 0x0f, 0x7e, 0x82, 0x8c, 0x74, 0xb8, 0x11, 0xba, 0xfd,
	0xdb, 0xa7, 0xbc, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x8c, 0xe9, 0xe6, 0x34, 0x02, 0x00,
	0x00,
}

func (m *MongoProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MongoProxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MongoProxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EmitDynamicMetadata {
		i--
		if m.EmitDynamicMetadata {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Delay != nil {
		{
			size, err := m.Delay.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMongoProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccessLog) > 0 {
		i -= len(m.AccessLog)
		copy(dAtA[i:], m.AccessLog)
		i = encodeVarintMongoProxy(dAtA, i, uint64(len(m.AccessLog)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintMongoProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMongoProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovMongoProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MongoProxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	l = len(m.AccessLog)
	if l > 0 {
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	if m.Delay != nil {
		l = m.Delay.Size()
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	if m.EmitDynamicMetadata {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMongoProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMongoProxy(x uint64) (n int) {
	return sovMongoProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MongoProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMongoProxy
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
			return fmt.Errorf("proto: MongoProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MongoProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLog = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delay == nil {
				m.Delay = &v2.FaultDelay{}
			}
			if err := m.Delay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmitDynamicMetadata", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
			m.EmitDynamicMetadata = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMongoProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMongoProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMongoProxy
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
					return 0, ErrIntOverflowMongoProxy
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
					return 0, ErrIntOverflowMongoProxy
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
				return 0, ErrInvalidLengthMongoProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMongoProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMongoProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMongoProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMongoProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMongoProxy = fmt.Errorf("proto: unexpected end of group")
)
