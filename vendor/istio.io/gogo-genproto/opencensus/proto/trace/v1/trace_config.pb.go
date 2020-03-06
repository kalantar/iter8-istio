// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opencensus/proto/trace/v1/trace_config.proto

package opencensus_proto_trace_v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfAnnotations int64 `protobuf:"varint,5,opt,name=max_number_of_annotations,json=maxNumberOfAnnotations,proto3" json:"max_number_of_annotations,omitempty"`
	// The global default max number of message events per span.
	MaxNumberOfMessageEvents int64 `protobuf:"varint,6,opt,name=max_number_of_message_events,json=maxNumberOfMessageEvents,proto3" json:"max_number_of_message_events,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,1,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof"`
}
type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,2,opt,name=constant_sampler,json=constantSampler,proto3,oneof"`
}
type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof"`
}

func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler()  {}
func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler()     {}
func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAnnotations() int64 {
	if m != nil {
		return m.MaxNumberOfAnnotations
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfMessageEvents() int64 {
	if m != nil {
		return m.MaxNumberOfMessageEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability float64 `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{1}
}
func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(m, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return m.Size()
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opencensus.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return m.Size()
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps int64 `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opencensus.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opencensus.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ProbabilitySampler)(nil), "opencensus.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*ConstantSampler)(nil), "opencensus.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opencensus.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opencensus/proto/trace/v1/trace_config.proto", fileDescriptor_5359209b41ff50c5)
}

var fileDescriptor_5359209b41ff50c5 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6e, 0xd3, 0x40,
	0x18, 0x85, 0x33, 0x49, 0x69, 0xe9, 0x5f, 0xb5, 0x35, 0x13, 0x15, 0x39, 0x52, 0x65, 0x55, 0xd9,
	0x50, 0x21, 0xea, 0x90, 0xb2, 0x40, 0x08, 0x09, 0x29, 0x49, 0x1b, 0xb1, 0x08, 0x49, 0xe4, 0x56,
	0x44, 0xac, 0xcc, 0xd8, 0x9d, 0x44, 0x23, 0xe2, 0x19, 0xe3, 0x99, 0x44, 0xe5, 0x16, 0x1c, 0x80,
	0x43, 0x70, 0x0c, 0x96, 0x5d, 0xb2, 0x44, 0xc9, 0x8e, 0x53, 0xa0, 0x8c, 0x8d, 0xed, 0x24, 0x6d,
	0xd5, 0xdd, 0xfc, 0xef, 0xbd, 0xef, 0x25, 0xf6, 0xfc, 0x86, 0x17, 0x22, 0xa4, 0xdc, 0xa7, 0x5c,
	0x4e, 0x64, 0x2d, 0x8c, 0x84, 0x12, 0x35, 0x15, 0x11, 0x9f, 0xd6, 0xa6, 0xf5, 0xf8, 0xe0, 0xfa,
	0x82, 0x0f, 0xd9, 0xc8, 0xd6, 0x1e, 0xae, 0x64, 0xe9, 0x58, 0xb1, 0x75, 0xc8, 0x9e, 0xd6, 0xab,
	0x3f, 0x36, 0x60, 0xe7, 0x72, 0x31, 0xb4, 0x34, 0x80, 0x3f, 0x43, 0x39, 0x8c, 0x84, 0x47, 0x3c,
	0x36, 0x66, 0xea, 0x9b, 0x2b, 0x49, 0x10, 0x8e, 0x69, 0x64, 0xa2, 0x23, 0x74, 0xbc, 0x73, 0x7a,
	0x62, 0xdf, 0x59, 0x64, 0xf7, 0x33, 0xea, 0x22, 0x86, 0xde, 0x17, 0x1c, 0x1c, 0xae, 0xa9, 0x78,
	0x00, 0x86, 0x2f, 0xb8, 0x54, 0x84, 0xab, 0xb4, 0xbe, 0xa8, 0xeb, 0x9f, 0xdf, 0x53, 0xdf, 0x4a,
	0x90, 0xac, 0x7b, 0xdf, 0x5f, 0x96, 0xf0, 0x15, 0x1c, 0x44, 0x44, 0x51, 0x77, 0xcc, 0x02, 0xa6,
	0x18, 0x1f, 0xa5, 0xed, 0x25, 0xdd, 0x6e, 0xdf, 0xd3, 0xee, 0x10, 0x45, 0x3b, 0x09, 0x96, 0xfd,
	0x42, 0x39, 0x5a, 0x97, 0xf1, 0x6b, 0x30, 0x03, 0x72, 0xed, 0xf2, 0x49, 0xe0, 0xd1, 0xc8, 0x15,
	0x43, 0x97, 0x28, 0x15, 0x31, 0x6f, 0xa2, 0xa8, 0x34, 0x37, 0x8e, 0xd0, 0x71, 0xc9, 0x39, 0x08,
	0xc8, 0x75, 0x57, 0xdb, 0xbd, 0x61, 0x23, 0x35, 0xf1, 0x1b, 0xa8, 0xac, 0x80, 0x9c, 0x0b, 0x45,
	0x14, 0x13, 0x5c, 0x9a, 0x8f, 0x34, 0xf9, 0x34, 0x4f, 0x66, 0x2e, 0x7e, 0x07, 0x87, 0xcb, 0x68,
	0x40, 0xa5, 0x24, 0x23, 0xea, 0xd2, 0x29, 0xe5, 0x4a, 0x9a, 0x9b, 0x9a, 0x36, 0x73, 0xf4, 0x87,
	0x38, 0x70, 0xae, 0x7d, 0x7c, 0x02, 0xe5, 0x65, 0x7e, 0xcc, 0xf8, 0x17, 0x69, 0x6e, 0x69, 0xcc,
	0xc8, 0x61, 0x9d, 0x85, 0xde, 0xdc, 0x86, 0xad, 0xe4, 0xd5, 0x55, 0xdb, 0x80, 0xd7, 0x2f, 0x16,
	0xbf, 0x84, 0xb2, 0x0e, 0x30, 0x3e, 0xca, 0xb9, 0x7a, 0x49, 0x90, 0x73, 0x9b, 0x55, 0xfd, 0x89,
	0x60, 0x7f, 0xe5, 0x0a, 0xf1, 0x00, 0x1e, 0x5f, 0x51, 0x9f, 0x49, 0x26, 0xb8, 0x46, 0xf7, 0x4e,
	0xdf, 0x3e, 0x7c, 0x01, 0xd2, 0xf9, 0x2c, 0xa9, 0x70, 0xd2, 0xb2, 0xea, 0x19, 0x18, 0xab, 0x2e,
	0xde, 0x03, 0x68, 0x74, 0x06, 0x8d, 0x4f, 0x17, 0x6e, 0xaf, 0xdd, 0x36, 0x0a, 0x78, 0x17, 0xb6,
	0xff, 0xcf, 0x5d, 0x03, 0xe1, 0x27, 0xb0, 0x9b, 0x8c, 0xfd, 0x86, 0x73, 0xde, 0xbd, 0x34, 0x8a,
	0xd5, 0x67, 0x50, 0xbe, 0x65, 0x2d, 0xb0, 0x01, 0xa5, 0xaf, 0xa1, 0xd4, 0x7f, 0xb8, 0xe4, 0x2c,
	0x8e, 0xcd, 0xe1, 0xaf, 0x99, 0x85, 0x6e, 0x66, 0x16, 0xfa, 0x33, 0xb3, 0xd0, 0xf7, 0xb9, 0x55,
	0xb8, 0x99, 0x5b, 0x85, 0xdf, 0x73, 0xab, 0x00, 0x87, 0x4c, 0xdc, 0xfd, 0x44, 0x4d, 0x23, 0xf7,
	0xdd, 0xf5, 0x17, 0x56, 0x1f, 0xfd, 0x2d, 0x56, 0x7a, 0x21, 0xe5, 0xad, 0x38, 0xaf, 0x45, 0x5b,
	0xa7, 0xec, 0x8f, 0x75, 0x6f, 0x53, 0xf3, 0xaf, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xeb,
	0x18, 0xb2, 0xfc, 0x03, 0x00, 0x00,
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxNumberOfLinks != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfLinks))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxNumberOfMessageEvents != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfMessageEvents))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxNumberOfAnnotations != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAnnotations))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxNumberOfAttributes != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributes))
		i--
		dAtA[i] = 0x20
	}
	if m.Sampler != nil {
		{
			size := m.Sampler.Size()
			i -= size
			if _, err := m.Sampler.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *TraceConfig_ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *TraceConfig_ProbabilitySampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProbabilitySampler != nil {
		{
			size, err := m.ProbabilitySampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *TraceConfig_ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ConstantSampler != nil {
		{
			size, err := m.ConstantSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *TraceConfig_RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RateLimitingSampler != nil {
		{
			size, err := m.RateLimitingSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ProbabilitySampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProbabilitySampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SamplingProbability))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *ConstantSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decision != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Decision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitingSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Qps != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Qps))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTraceConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraceConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sampler != nil {
		n += m.Sampler.Size()
	}
	if m.MaxNumberOfAttributes != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributes))
	}
	if m.MaxNumberOfAnnotations != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAnnotations))
	}
	if m.MaxNumberOfMessageEvents != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfMessageEvents))
	}
	if m.MaxNumberOfLinks != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfLinks))
	}
	return n
}

func (m *TraceConfig_ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProbabilitySampler != nil {
		l = m.ProbabilitySampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConstantSampler != nil {
		l = m.ConstantSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RateLimitingSampler != nil {
		l = m.RateLimitingSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		n += 9
	}
	return n
}

func (m *ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Decision != 0 {
		n += 1 + sovTraceConfig(uint64(m.Decision))
	}
	return n
}

func (m *RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Qps != 0 {
		n += 1 + sovTraceConfig(uint64(m.Qps))
	}
	return n
}

func sovTraceConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraceConfig(x uint64) (n int) {
	return sovTraceConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProbabilitySampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ProbabilitySampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ProbabilitySampler{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConstantSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ConstantSampler{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitingSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RateLimitingSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_RateLimitingSampler{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributes", wireType)
			}
			m.MaxNumberOfAttributes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAnnotations", wireType)
			}
			m.MaxNumberOfAnnotations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAnnotations |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfMessageEvents", wireType)
			}
			m.MaxNumberOfMessageEvents = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfMessageEvents |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfLinks", wireType)
			}
			m.MaxNumberOfLinks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfLinks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *ProbabilitySampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: ProbabilitySampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProbabilitySampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingProbability", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SamplingProbability = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *ConstantSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: ConstantSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConstantSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decision", wireType)
			}
			m.Decision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decision |= ConstantSampler_ConstantDecision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *RateLimitingSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: RateLimitingSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitingSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qps", wireType)
			}
			m.Qps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Qps |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func skipTraceConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceConfig
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
					return 0, ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceConfig
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
				return 0, ErrInvalidLengthTraceConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTraceConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTraceConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTraceConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTraceConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTraceConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceConfig   = fmt.Errorf("proto: integer overflow")
)