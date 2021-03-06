// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/route/v3/route.proto

package envoy_config_route_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#next-free-field: 11]
type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_extensions.filters.network.http_connection_manager.v3.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_extensions.filters.network.http_connection_manager.v3.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// An array of virtual hosts will be dynamically loaded via the VHDS API.
	// Both *virtual_hosts* and *vhds* fields will be used when present. *virtual_hosts* can be used
	// for a base routing table or for infrequently changing virtual hosts. *vhds* is used for
	// on-demand discovery of virtual hosts. The contents of these two fields will be merged to
	// generate a routing table for a given RouteConfiguration, with *vhds* derived configuration
	// taking precedence.
	Vhds *Vhds `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_config.route.v3.VirtualHost` or
	// :ref:`envoy_api_msg_config.route.v3.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_config.route.v3.VirtualHost` or
	// :ref:`envoy_api_msg_config.route.v3.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// routed by the HTTP connection manager.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// By default, headers that should be added/removed are evaluated from most to least specific:
	//
	// * route level
	// * virtual host level
	// * connection manager level
	//
	// To allow setting overrides at the route or virtual host level, this order can be reversed
	// by setting this option to true. Defaults to false.
	//
	// [#next-major-version: In the v3 API, this will default to true.]
	MostSpecificHeaderMutationsWins bool `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.rds>`
	// option. Users may wish to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *types.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb402428dbbe6a7, []int{0}
}
func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*v3.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*v3.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if m != nil {
		return m.MostSpecificHeaderMutationsWins
	}
	return false
}

func (m *RouteConfiguration) GetValidateClusters() *types.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	// Configuration source specifier for VHDS.
	ConfigSource         *v3.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb402428dbbe6a7, []int{1}
}
func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return m.Size()
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *v3.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.route.v3.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.config.route.v3.Vhds")
}

func init() { proto.RegisterFile("envoy/config/route/v3/route.proto", fileDescriptor_bdb402428dbbe6a7) }

var fileDescriptor_bdb402428dbbe6a7 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0x9b, 0xf4, 0x27, 0xd3, 0x56, 0x6a, 0xa7, 0x5f, 0x5b, 0x7f, 0x41, 0x4a, 0xd3, 0x56,
	0x02, 0x2f, 0x90, 0x2d, 0xa5, 0x2b, 0xca, 0x02, 0xe1, 0x2e, 0xda, 0x05, 0xa8, 0xc5, 0x45, 0x65,
	0x69, 0x4d, 0xec, 0x49, 0x32, 0x92, 0x33, 0xd7, 0xcc, 0x8c, 0x5d, 0xf2, 0x06, 0x88, 0x25, 0x4b,
	0xde, 0x07, 0x89, 0x25, 0x8f, 0x00, 0x59, 0xf1, 0x0c, 0x5d, 0x21, 0xcf, 0x8c, 0x05, 0x69, 0x02,
	0x0b, 0x76, 0x73, 0x7d, 0xcf, 0xb9, 0xe7, 0xf8, 0xce, 0x19, 0x74, 0x48, 0x79, 0x09, 0x93, 0x20,
	0x01, 0x3e, 0x60, 0xc3, 0x40, 0x40, 0xa1, 0x68, 0x50, 0x9e, 0x98, 0x83, 0x9f, 0x0b, 0x50, 0x80,
	0x77, 0x35, 0xc4, 0x37, 0x10, 0xdf, 0x74, 0xca, 0x93, 0xf6, 0xc1, 0x0c, 0x33, 0x01, 0xa1, 0x89,
	0x7d, 0x22, 0x2d, 0xaf, 0xed, 0x2d, 0x04, 0x98, 0x32, 0x96, 0x50, 0x88, 0xa4, 0x46, 0x3e, 0xfe,
	0x8b, 0x89, 0x38, 0x81, 0x71, 0x0e, 0x9c, 0x72, 0x25, 0x2d, 0xba, 0x33, 0x04, 0x18, 0x66, 0x34,
	0xd0, 0x55, 0xbf, 0x18, 0x04, 0xb7, 0x82, 0xe4, 0x39, 0x15, 0x75, 0xff, 0xb0, 0x48, 0x73, 0x12,
	0x10, 0xce, 0x41, 0x11, 0xc5, 0x80, 0xcb, 0xa0, 0xa4, 0x42, 0x32, 0xe0, 0x8c, 0x0f, 0x2d, 0x64,
	0xbf, 0x24, 0x19, 0x4b, 0x49, 0x25, 0x63, 0x0f, 0xa6, 0x71, 0xf4, 0x7d, 0x19, 0xe1, 0xa8, 0x92,
	0x3d, 0xd3, 0x5e, 0x0a, 0xa1, 0x27, 0x60, 0x8c, 0x9a, 0x9c, 0x8c, 0xa9, 0xeb, 0x74, 0x1d, 0xaf,
	0x15, 0xe9, 0x33, 0x3e, 0x47, 0x9b, 0x25, 0x13, 0xaa, 0x20, 0x59, 0x3c, 0x02, 0xa9, 0xa4, 0xbb,
	0xd4, 0x6d, 0x78, 0xeb, 0xbd, 0x23, 0x7f, 0xe1, 0xba, 0xfc, 0x1b, 0x83, 0xbd, 0x00, 0xa9, 0xa2,
	0x8d, 0xf2, 0x57, 0x21, 0x71, 0x80, 0x9a, 0xe5, 0x28, 0x95, 0x6e, 0xab, 0xeb, 0x78, 0xeb, 0xbd,
	0x07, 0x7f, 0xe2, 0x8f, 0x52, 0x19, 0x69, 0x20, 0xee, 0xa1, 0x5d, 0xc6, 0x15, 0x15, 0x9c, 0x64,
	0x31, 0xf0, 0x6c, 0x12, 0x8f, 0x28, 0x49, 0xa9, 0x90, 0x6e, 0xa3, 0xdb, 0xf0, 0x5a, 0xd1, 0x4e,
	0xdd, 0xbc, 0xe4, 0xd9, 0xe4, 0xc2, 0xb4, 0x30, 0x43, 0xfb, 0x82, 0xca, 0x1c, 0xb8, 0xa4, 0x35,
	0x3c, 0x56, 0x10, 0x93, 0x34, 0x75, 0x9b, 0xda, 0xf7, 0xa3, 0x59, 0xdd, 0xea, 0xba, 0x2a, 0x59,
	0xc3, 0xbf, 0x21, 0x59, 0x41, 0x2f, 0xf3, 0x6a, 0x17, 0x61, 0xeb, 0x2e, 0x5c, 0xf9, 0xe8, 0x34,
	0xb6, 0x7e, 0xac, 0x46, 0xff, 0xd5, 0x23, 0xad, 0xca, 0x6b, 0x78, 0x9e, 0xa6, 0xf8, 0x29, 0x6a,
	0x2f, 0x92, 0x12, 0x74, 0x0c, 0x25, 0x75, 0x97, 0xb5, 0xc7, 0xfd, 0x39, 0x66, 0xa4, 0xdb, 0x78,
	0x88, 0xf6, 0x04, 0x7d, 0x5b, 0x50, 0xa9, 0xee, 0xdb, 0x5c, 0xf9, 0x67, 0x9b, 0x3b, 0x76, 0xe2,
	0x8c, 0xcb, 0x27, 0xe8, 0xff, 0x05, 0x42, 0xd6, 0xe4, 0x9a, 0x36, 0xb9, 0x77, 0x9f, 0x67, 0x3d,
	0xbe, 0x40, 0xc7, 0x63, 0x90, 0x2a, 0x96, 0x39, 0x4d, 0xd8, 0x80, 0x25, 0x76, 0x40, 0x3c, 0x2e,
	0x6c, 0xe0, 0xe2, 0x5b, 0xc6, 0xa5, 0x8b, 0xba, 0x8e, 0xb7, 0x16, 0x1d, 0x54, 0xd0, 0x6b, 0x8b,
	0x34, 0x93, 0x5e, 0xd6, 0xb8, 0x37, 0x8c, 0x4b, 0x7c, 0x8e, 0xb6, 0xeb, 0x10, 0xc6, 0x49, 0x56,
	0x48, 0x55, 0xdd, 0xe4, 0xaa, 0xce, 0x42, 0xdb, 0x37, 0x51, 0xf7, 0xeb, 0xa8, 0xfb, 0x21, 0x40,
	0xa6, 0xff, 0x32, 0xda, 0xaa, 0x49, 0x67, 0x96, 0x73, 0xfa, 0xf0, 0xd3, 0xe7, 0xf7, 0x9d, 0x43,
	0x64, 0xde, 0xa5, 0x4f, 0x72, 0xe6, 0x97, 0x3d, 0x7f, 0x3e, 0xcc, 0x47, 0x12, 0x35, 0xab, 0x30,
	0xe1, 0x57, 0x68, 0x73, 0xe6, 0x31, 0xea, 0x74, 0xcf, 0x05, 0xb8, 0xde, 0xb0, 0x99, 0x71, 0xad,
	0x91, 0xe1, 0xda, 0x5d, 0xb8, 0xfc, 0xc1, 0x59, 0xda, 0x72, 0xa2, 0x8d, 0xe4, 0xb7, 0xef, 0xa7,
	0x6e, 0x65, 0x61, 0x07, 0x6d, 0xcf, 0x58, 0xa8, 0xc4, 0xc2, 0x67, 0x5f, 0xa6, 0x1d, 0xe7, 0xeb,
	0xb4, 0xe3, 0x7c, 0x9b, 0x76, 0x1c, 0x74, 0xcc, 0xc0, 0xa8, 0xe4, 0x02, 0xde, 0x4d, 0x16, 0x27,
	0x3e, 0x44, 0xda, 0xfb, 0x55, 0xf5, 0xeb, 0x57, 0x4e, 0x7f, 0x45, 0xef, 0xe0, 0xe4, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0x70, 0x66, 0x8e, 0xb1, 0x04, 0x00, 0x00,
}

func (m *RouteConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RouteConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MostSpecificHeaderMutationsWins {
		i--
		if m.MostSpecificHeaderMutationsWins {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.Vhds != nil {
		{
			size, err := m.Vhds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.RequestHeadersToRemove) > 0 {
		for iNdEx := len(m.RequestHeadersToRemove) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RequestHeadersToRemove[iNdEx])
			copy(dAtA[i:], m.RequestHeadersToRemove[iNdEx])
			i = encodeVarintRoute(dAtA, i, uint64(len(m.RequestHeadersToRemove[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ValidateClusters != nil {
		{
			size, err := m.ValidateClusters.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RequestHeadersToAdd) > 0 {
		for iNdEx := len(m.RequestHeadersToAdd) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestHeadersToAdd[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ResponseHeadersToRemove) > 0 {
		for iNdEx := len(m.ResponseHeadersToRemove) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResponseHeadersToRemove[iNdEx])
			copy(dAtA[i:], m.ResponseHeadersToRemove[iNdEx])
			i = encodeVarintRoute(dAtA, i, uint64(len(m.ResponseHeadersToRemove[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for iNdEx := len(m.ResponseHeadersToAdd) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResponseHeadersToAdd[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InternalOnlyHeaders) > 0 {
		for iNdEx := len(m.InternalOnlyHeaders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InternalOnlyHeaders[iNdEx])
			copy(dAtA[i:], m.InternalOnlyHeaders[iNdEx])
			i = encodeVarintRoute(dAtA, i, uint64(len(m.InternalOnlyHeaders[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VirtualHosts) > 0 {
		for iNdEx := len(m.VirtualHosts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VirtualHosts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Vhds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vhds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vhds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConfigSource != nil {
		{
			size, err := m.ConfigSource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RouteConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.VirtualHosts) > 0 {
		for _, e := range m.VirtualHosts {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if len(m.InternalOnlyHeaders) > 0 {
		for _, s := range m.InternalOnlyHeaders {
			l = len(s)
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for _, e := range m.ResponseHeadersToAdd {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if len(m.ResponseHeadersToRemove) > 0 {
		for _, s := range m.ResponseHeadersToRemove {
			l = len(s)
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if len(m.RequestHeadersToAdd) > 0 {
		for _, e := range m.RequestHeadersToAdd {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if m.ValidateClusters != nil {
		l = m.ValidateClusters.Size()
		n += 1 + l + sovRoute(uint64(l))
	}
	if len(m.RequestHeadersToRemove) > 0 {
		for _, s := range m.RequestHeadersToRemove {
			l = len(s)
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if m.Vhds != nil {
		l = m.Vhds.Size()
		n += 1 + l + sovRoute(uint64(l))
	}
	if m.MostSpecificHeaderMutationsWins {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Vhds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigSource != nil {
		l = m.ConfigSource.Size()
		n += 1 + l + sovRoute(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoute(x uint64) (n int) {
	return sovRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RouteConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: RouteConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualHosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualHosts = append(m.VirtualHosts, &VirtualHost{})
			if err := m.VirtualHosts[len(m.VirtualHosts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalOnlyHeaders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalOnlyHeaders = append(m.InternalOnlyHeaders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeadersToAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseHeadersToAdd = append(m.ResponseHeadersToAdd, &v3.HeaderValueOption{})
			if err := m.ResponseHeadersToAdd[len(m.ResponseHeadersToAdd)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeadersToRemove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseHeadersToRemove = append(m.ResponseHeadersToRemove, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHeadersToAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestHeadersToAdd = append(m.RequestHeadersToAdd, &v3.HeaderValueOption{})
			if err := m.RequestHeadersToAdd[len(m.RequestHeadersToAdd)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidateClusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidateClusters == nil {
				m.ValidateClusters = &types.BoolValue{}
			}
			if err := m.ValidateClusters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHeadersToRemove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestHeadersToRemove = append(m.RequestHeadersToRemove, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vhds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Vhds == nil {
				m.Vhds = &Vhds{}
			}
			if err := m.Vhds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MostSpecificHeaderMutationsWins", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
			m.MostSpecificHeaderMutationsWins = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRoute
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
func (m *Vhds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: Vhds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vhds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigSource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConfigSource == nil {
				m.ConfigSource = &v3.ConfigSource{}
			}
			if err := m.ConfigSource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoute
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRoute
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
func skipRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoute
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
					return 0, ErrIntOverflowRoute
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
					return 0, ErrIntOverflowRoute
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
				return 0, ErrInvalidLengthRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoute = fmt.Errorf("proto: unexpected end of group")
)
