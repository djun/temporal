// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/persistence/v1/queues.proto

package persistence

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
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

type QueueState struct {
	ReaderStates                 map[int64]*QueueReaderState `protobuf:"bytes,1,rep,name=reader_states,json=readerStates,proto3" json:"reader_states,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExclusiveReaderHighWatermark *TaskKey                    `protobuf:"bytes,2,opt,name=exclusive_reader_high_watermark,json=exclusiveReaderHighWatermark,proto3" json:"exclusive_reader_high_watermark,omitempty"`
}

func (m *QueueState) Reset()      { *m = QueueState{} }
func (*QueueState) ProtoMessage() {}
func (*QueueState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7fa5f143ac80378, []int{0}
}
func (m *QueueState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueState.Merge(m, src)
}
func (m *QueueState) XXX_Size() int {
	return m.Size()
}
func (m *QueueState) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueState.DiscardUnknown(m)
}

var xxx_messageInfo_QueueState proto.InternalMessageInfo

func (m *QueueState) GetReaderStates() map[int64]*QueueReaderState {
	if m != nil {
		return m.ReaderStates
	}
	return nil
}

func (m *QueueState) GetExclusiveReaderHighWatermark() *TaskKey {
	if m != nil {
		return m.ExclusiveReaderHighWatermark
	}
	return nil
}

type QueueReaderState struct {
	Scopes []*QueueSliceScope `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (m *QueueReaderState) Reset()      { *m = QueueReaderState{} }
func (*QueueReaderState) ProtoMessage() {}
func (*QueueReaderState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7fa5f143ac80378, []int{1}
}
func (m *QueueReaderState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueReaderState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueReaderState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueReaderState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueReaderState.Merge(m, src)
}
func (m *QueueReaderState) XXX_Size() int {
	return m.Size()
}
func (m *QueueReaderState) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueReaderState.DiscardUnknown(m)
}

var xxx_messageInfo_QueueReaderState proto.InternalMessageInfo

func (m *QueueReaderState) GetScopes() []*QueueSliceScope {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type QueueSliceScope struct {
	Range     *QueueSliceRange `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	Predicate *Predicate       `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (m *QueueSliceScope) Reset()      { *m = QueueSliceScope{} }
func (*QueueSliceScope) ProtoMessage() {}
func (*QueueSliceScope) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7fa5f143ac80378, []int{2}
}
func (m *QueueSliceScope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueSliceScope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueSliceScope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueSliceScope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueSliceScope.Merge(m, src)
}
func (m *QueueSliceScope) XXX_Size() int {
	return m.Size()
}
func (m *QueueSliceScope) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueSliceScope.DiscardUnknown(m)
}

var xxx_messageInfo_QueueSliceScope proto.InternalMessageInfo

func (m *QueueSliceScope) GetRange() *QueueSliceRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *QueueSliceScope) GetPredicate() *Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

type QueueSliceRange struct {
	InclusiveMin *TaskKey `protobuf:"bytes,1,opt,name=inclusive_min,json=inclusiveMin,proto3" json:"inclusive_min,omitempty"`
	ExclusiveMax *TaskKey `protobuf:"bytes,2,opt,name=exclusive_max,json=exclusiveMax,proto3" json:"exclusive_max,omitempty"`
}

func (m *QueueSliceRange) Reset()      { *m = QueueSliceRange{} }
func (*QueueSliceRange) ProtoMessage() {}
func (*QueueSliceRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7fa5f143ac80378, []int{3}
}
func (m *QueueSliceRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueSliceRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueSliceRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueSliceRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueSliceRange.Merge(m, src)
}
func (m *QueueSliceRange) XXX_Size() int {
	return m.Size()
}
func (m *QueueSliceRange) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueSliceRange.DiscardUnknown(m)
}

var xxx_messageInfo_QueueSliceRange proto.InternalMessageInfo

func (m *QueueSliceRange) GetInclusiveMin() *TaskKey {
	if m != nil {
		return m.InclusiveMin
	}
	return nil
}

func (m *QueueSliceRange) GetExclusiveMax() *TaskKey {
	if m != nil {
		return m.ExclusiveMax
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueState)(nil), "temporal.server.api.persistence.v1.QueueState")
	proto.RegisterMapType((map[int64]*QueueReaderState)(nil), "temporal.server.api.persistence.v1.QueueState.ReaderStatesEntry")
	proto.RegisterType((*QueueReaderState)(nil), "temporal.server.api.persistence.v1.QueueReaderState")
	proto.RegisterType((*QueueSliceScope)(nil), "temporal.server.api.persistence.v1.QueueSliceScope")
	proto.RegisterType((*QueueSliceRange)(nil), "temporal.server.api.persistence.v1.QueueSliceRange")
}

func init() {
	proto.RegisterFile("temporal/server/api/persistence/v1/queues.proto", fileDescriptor_b7fa5f143ac80378)
}

var fileDescriptor_b7fa5f143ac80378 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x77, 0x12, 0x5a, 0x70, 0x9a, 0x62, 0x9d, 0x53, 0x08, 0x32, 0x86, 0x3d, 0x05, 0xc4,
	0x59, 0xda, 0xf4, 0x20, 0x7a, 0x11, 0x41, 0x50, 0x43, 0xa1, 0x4e, 0x05, 0xc1, 0x4b, 0x18, 0xb7,
	0x2f, 0xc9, 0x98, 0x64, 0x77, 0x9d, 0x99, 0x5d, 0x93, 0x9b, 0x1f, 0xc1, 0x8f, 0xa1, 0x1f, 0xc0,
	0xef, 0xe0, 0x31, 0xc7, 0x9e, 0xc4, 0x6c, 0x2e, 0x1e, 0xfb, 0x11, 0x64, 0xff, 0xaf, 0x15, 0x71,
	0xdb, 0xdb, 0xcc, 0xce, 0x3e, 0xbf, 0xe7, 0x99, 0x77, 0xde, 0x17, 0x3b, 0x06, 0x16, 0x81, 0xaf,
	0xc4, 0xdc, 0xd1, 0xa0, 0x22, 0x50, 0x8e, 0x08, 0xa4, 0x13, 0x80, 0xd2, 0x52, 0x1b, 0xf0, 0x5c,
	0x70, 0xa2, 0x43, 0xe7, 0x43, 0x08, 0x21, 0x68, 0x16, 0x28, 0xdf, 0xf8, 0xc4, 0x2e, 0x04, 0x2c,
	0x13, 0x30, 0x11, 0x48, 0x56, 0x13, 0xb0, 0xe8, 0xb0, 0x37, 0x6c, 0x00, 0x0d, 0x14, 0x9c, 0x4b,
	0x57, 0x98, 0x02, 0xdc, 0x63, 0x0d, 0x44, 0x46, 0xe8, 0x59, 0xfe, 0xbf, 0xfd, 0xa3, 0x85, 0xf1,
	0xab, 0x24, 0xd9, 0x99, 0x11, 0x06, 0x08, 0xe0, 0x7d, 0x05, 0xe2, 0x1c, 0xd4, 0x58, 0x27, 0x7b,
	0xdd, 0x45, 0xfd, 0xf6, 0x60, 0xef, 0xe8, 0x09, 0xfb, 0x7f, 0x5e, 0x56, 0x61, 0x18, 0x4f, 0x19,
	0xe9, 0x5a, 0x3f, 0xf3, 0x8c, 0x5a, 0xf1, 0x8e, 0xaa, 0x7d, 0x22, 0x0a, 0xdf, 0x83, 0xa5, 0x3b,
	0x0f, 0xb5, 0x8c, 0x60, 0x9c, 0x1b, 0x4e, 0xe5, 0x64, 0x3a, 0xfe, 0x28, 0x0c, 0xa8, 0x85, 0x50,
	0xb3, 0x6e, 0xab, 0x8f, 0x06, 0x7b, 0x47, 0xf7, 0x9b, 0x18, 0xbf, 0x16, 0x7a, 0x36, 0x82, 0x15,
	0xbf, 0x5b, 0x32, 0x33, 0xff, 0xe7, 0x72, 0x32, 0x7d, 0x53, 0x00, 0x7b, 0x21, 0xbe, 0xf3, 0x57,
	0x2c, 0x72, 0x80, 0xdb, 0x33, 0x58, 0x75, 0x51, 0x1f, 0x0d, 0xda, 0x3c, 0x59, 0x92, 0x97, 0x78,
	0x27, 0x12, 0xf3, 0x10, 0xf2, 0x00, 0xc7, 0x8d, 0x6f, 0x5e, 0x83, 0xf3, 0x0c, 0xf1, 0xa8, 0xf5,
	0x10, 0xd9, 0x63, 0x7c, 0x70, 0xf5, 0x98, 0x8c, 0xf0, 0xae, 0x76, 0xfd, 0xa0, 0x2c, 0xef, 0xb0,
	0x79, 0x79, 0xe7, 0xd2, 0x85, 0xb3, 0x44, 0xcb, 0x73, 0x84, 0xfd, 0x15, 0xe1, 0xdb, 0x57, 0xce,
	0xc8, 0x0b, 0xbc, 0xa3, 0x84, 0x37, 0x81, 0xf4, 0x62, 0xd7, 0xe6, 0xf3, 0x44, 0xca, 0x33, 0x02,
	0x19, 0xe1, 0x5b, 0x65, 0x93, 0xe5, 0x35, 0x79, 0xd0, 0x04, 0x77, 0x5a, 0x88, 0x78, 0xa5, 0xb7,
	0xbf, 0xfd, 0x91, 0x35, 0xf5, 0x21, 0xa7, 0x78, 0x5f, 0x7a, 0x45, 0x2f, 0x2c, 0xa4, 0x97, 0x67,
	0xbe, 0xd6, 0xcb, 0x77, 0x4a, 0xc2, 0x89, 0xf4, 0x12, 0x62, 0xd5, 0x5d, 0x0b, 0xb1, 0xbc, 0x49,
	0x2f, 0x75, 0x4a, 0xc2, 0x89, 0x58, 0x3e, 0x7d, 0xbf, 0xde, 0x50, 0xeb, 0x62, 0x43, 0xad, 0xcb,
	0x0d, 0x45, 0x9f, 0x62, 0x8a, 0xbe, 0xc4, 0x14, 0x7d, 0x8f, 0x29, 0x5a, 0xc7, 0x14, 0xfd, 0x8c,
	0x29, 0xfa, 0x15, 0x53, 0xeb, 0x32, 0xa6, 0xe8, 0xf3, 0x96, 0x5a, 0xeb, 0x2d, 0xb5, 0x2e, 0xb6,
	0xd4, 0x7a, 0x7b, 0x3c, 0xf1, 0x2b, 0x4b, 0xe9, 0xff, 0x7b, 0x22, 0x1f, 0xd7, 0xb6, 0xef, 0x76,
	0xd3, 0xc1, 0x1c, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xfe, 0x92, 0xa2, 0x54, 0x04, 0x00,
	0x00,
}

func (this *QueueState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueState)
	if !ok {
		that2, ok := that.(QueueState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ReaderStates) != len(that1.ReaderStates) {
		return false
	}
	for i := range this.ReaderStates {
		if !this.ReaderStates[i].Equal(that1.ReaderStates[i]) {
			return false
		}
	}
	if !this.ExclusiveReaderHighWatermark.Equal(that1.ExclusiveReaderHighWatermark) {
		return false
	}
	return true
}
func (this *QueueReaderState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueReaderState)
	if !ok {
		that2, ok := that.(QueueReaderState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Scopes) != len(that1.Scopes) {
		return false
	}
	for i := range this.Scopes {
		if !this.Scopes[i].Equal(that1.Scopes[i]) {
			return false
		}
	}
	return true
}
func (this *QueueSliceScope) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueSliceScope)
	if !ok {
		that2, ok := that.(QueueSliceScope)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Range.Equal(that1.Range) {
		return false
	}
	if !this.Predicate.Equal(that1.Predicate) {
		return false
	}
	return true
}
func (this *QueueSliceRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueSliceRange)
	if !ok {
		that2, ok := that.(QueueSliceRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.InclusiveMin.Equal(that1.InclusiveMin) {
		return false
	}
	if !this.ExclusiveMax.Equal(that1.ExclusiveMax) {
		return false
	}
	return true
}
func (this *QueueState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&persistence.QueueState{")
	keysForReaderStates := make([]int64, 0, len(this.ReaderStates))
	for k, _ := range this.ReaderStates {
		keysForReaderStates = append(keysForReaderStates, k)
	}
	github_com_gogo_protobuf_sortkeys.Int64s(keysForReaderStates)
	mapStringForReaderStates := "map[int64]*QueueReaderState{"
	for _, k := range keysForReaderStates {
		mapStringForReaderStates += fmt.Sprintf("%#v: %#v,", k, this.ReaderStates[k])
	}
	mapStringForReaderStates += "}"
	if this.ReaderStates != nil {
		s = append(s, "ReaderStates: "+mapStringForReaderStates+",\n")
	}
	if this.ExclusiveReaderHighWatermark != nil {
		s = append(s, "ExclusiveReaderHighWatermark: "+fmt.Sprintf("%#v", this.ExclusiveReaderHighWatermark)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueueReaderState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&persistence.QueueReaderState{")
	if this.Scopes != nil {
		s = append(s, "Scopes: "+fmt.Sprintf("%#v", this.Scopes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueueSliceScope) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&persistence.QueueSliceScope{")
	if this.Range != nil {
		s = append(s, "Range: "+fmt.Sprintf("%#v", this.Range)+",\n")
	}
	if this.Predicate != nil {
		s = append(s, "Predicate: "+fmt.Sprintf("%#v", this.Predicate)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueueSliceRange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&persistence.QueueSliceRange{")
	if this.InclusiveMin != nil {
		s = append(s, "InclusiveMin: "+fmt.Sprintf("%#v", this.InclusiveMin)+",\n")
	}
	if this.ExclusiveMax != nil {
		s = append(s, "ExclusiveMax: "+fmt.Sprintf("%#v", this.ExclusiveMax)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQueues(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QueueState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExclusiveReaderHighWatermark != nil {
		{
			size, err := m.ExclusiveReaderHighWatermark.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueues(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ReaderStates) > 0 {
		for k := range m.ReaderStates {
			v := m.ReaderStates[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintQueues(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintQueues(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintQueues(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueueReaderState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueReaderState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueReaderState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scopes) > 0 {
		for iNdEx := len(m.Scopes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Scopes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueues(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueueSliceScope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueSliceScope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueSliceScope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Predicate != nil {
		{
			size, err := m.Predicate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueues(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Range != nil {
		{
			size, err := m.Range.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueues(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueueSliceRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueSliceRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueSliceRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExclusiveMax != nil {
		{
			size, err := m.ExclusiveMax.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueues(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.InclusiveMin != nil {
		{
			size, err := m.InclusiveMin.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueues(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueues(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueues(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueueState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReaderStates) > 0 {
		for k, v := range m.ReaderStates {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovQueues(uint64(l))
			}
			mapEntrySize := 1 + sovQueues(uint64(k)) + l
			n += mapEntrySize + 1 + sovQueues(uint64(mapEntrySize))
		}
	}
	if m.ExclusiveReaderHighWatermark != nil {
		l = m.ExclusiveReaderHighWatermark.Size()
		n += 1 + l + sovQueues(uint64(l))
	}
	return n
}

func (m *QueueReaderState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Scopes) > 0 {
		for _, e := range m.Scopes {
			l = e.Size()
			n += 1 + l + sovQueues(uint64(l))
		}
	}
	return n
}

func (m *QueueSliceScope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovQueues(uint64(l))
	}
	if m.Predicate != nil {
		l = m.Predicate.Size()
		n += 1 + l + sovQueues(uint64(l))
	}
	return n
}

func (m *QueueSliceRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InclusiveMin != nil {
		l = m.InclusiveMin.Size()
		n += 1 + l + sovQueues(uint64(l))
	}
	if m.ExclusiveMax != nil {
		l = m.ExclusiveMax.Size()
		n += 1 + l + sovQueues(uint64(l))
	}
	return n
}

func sovQueues(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueues(x uint64) (n int) {
	return sovQueues(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueueState) String() string {
	if this == nil {
		return "nil"
	}
	keysForReaderStates := make([]int64, 0, len(this.ReaderStates))
	for k, _ := range this.ReaderStates {
		keysForReaderStates = append(keysForReaderStates, k)
	}
	github_com_gogo_protobuf_sortkeys.Int64s(keysForReaderStates)
	mapStringForReaderStates := "map[int64]*QueueReaderState{"
	for _, k := range keysForReaderStates {
		mapStringForReaderStates += fmt.Sprintf("%v: %v,", k, this.ReaderStates[k])
	}
	mapStringForReaderStates += "}"
	s := strings.Join([]string{`&QueueState{`,
		`ReaderStates:` + mapStringForReaderStates + `,`,
		`ExclusiveReaderHighWatermark:` + strings.Replace(fmt.Sprintf("%v", this.ExclusiveReaderHighWatermark), "TaskKey", "TaskKey", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueueReaderState) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForScopes := "[]*QueueSliceScope{"
	for _, f := range this.Scopes {
		repeatedStringForScopes += strings.Replace(f.String(), "QueueSliceScope", "QueueSliceScope", 1) + ","
	}
	repeatedStringForScopes += "}"
	s := strings.Join([]string{`&QueueReaderState{`,
		`Scopes:` + repeatedStringForScopes + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueueSliceScope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueueSliceScope{`,
		`Range:` + strings.Replace(this.Range.String(), "QueueSliceRange", "QueueSliceRange", 1) + `,`,
		`Predicate:` + strings.Replace(fmt.Sprintf("%v", this.Predicate), "Predicate", "Predicate", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueueSliceRange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueueSliceRange{`,
		`InclusiveMin:` + strings.Replace(fmt.Sprintf("%v", this.InclusiveMin), "TaskKey", "TaskKey", 1) + `,`,
		`ExclusiveMax:` + strings.Replace(fmt.Sprintf("%v", this.ExclusiveMax), "TaskKey", "TaskKey", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQueues(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueueState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueues
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
			return fmt.Errorf("proto: QueueState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReaderStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReaderStates == nil {
				m.ReaderStates = make(map[int64]*QueueReaderState)
			}
			var mapkey int64
			var mapvalue *QueueReaderState
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQueues
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueues
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueues
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthQueues
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthQueues
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &QueueReaderState{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQueues(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQueues
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ReaderStates[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExclusiveReaderHighWatermark", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExclusiveReaderHighWatermark == nil {
				m.ExclusiveReaderHighWatermark = &TaskKey{}
			}
			if err := m.ExclusiveReaderHighWatermark.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueues(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueues
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueues
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
func (m *QueueReaderState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueues
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
			return fmt.Errorf("proto: QueueReaderState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueReaderState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scopes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scopes = append(m.Scopes, &QueueSliceScope{})
			if err := m.Scopes[len(m.Scopes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueues(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueues
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueues
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
func (m *QueueSliceScope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueues
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
			return fmt.Errorf("proto: QueueSliceScope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueSliceScope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &QueueSliceRange{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Predicate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Predicate == nil {
				m.Predicate = &Predicate{}
			}
			if err := m.Predicate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueues(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueues
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueues
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
func (m *QueueSliceRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueues
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
			return fmt.Errorf("proto: QueueSliceRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueSliceRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InclusiveMin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InclusiveMin == nil {
				m.InclusiveMin = &TaskKey{}
			}
			if err := m.InclusiveMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExclusiveMax", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueues
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
				return ErrInvalidLengthQueues
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueues
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExclusiveMax == nil {
				m.ExclusiveMax = &TaskKey{}
			}
			if err := m.ExclusiveMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueues(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueues
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueues
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
func skipQueues(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueues
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
					return 0, ErrIntOverflowQueues
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
					return 0, ErrIntOverflowQueues
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
				return 0, ErrInvalidLengthQueues
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueues
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueues
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueues        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueues          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueues = fmt.Errorf("proto: unexpected end of group")
)
