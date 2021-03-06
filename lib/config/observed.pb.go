// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/observed.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"
	_ "github.com/syncthing/syncthing/proto/ext"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ObservedFolder struct {
	Time  time.Time `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time" xml:"time,attr"`
	ID    string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id" xml:"id,attr"`
	Label string    `protobuf:"bytes,3,opt,name=label,proto3" json:"label" xml:"label,attr"`
}

func (m *ObservedFolder) Reset()         { *m = ObservedFolder{} }
func (m *ObservedFolder) String() string { return proto.CompactTextString(m) }
func (*ObservedFolder) ProtoMessage()    {}
func (*ObservedFolder) Descriptor() ([]byte, []int) {
	return fileDescriptor_49f68ff7b178722f, []int{0}
}
func (m *ObservedFolder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObservedFolder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObservedFolder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObservedFolder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservedFolder.Merge(m, src)
}
func (m *ObservedFolder) XXX_Size() int {
	return m.ProtoSize()
}
func (m *ObservedFolder) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservedFolder.DiscardUnknown(m)
}

var xxx_messageInfo_ObservedFolder proto.InternalMessageInfo

type ObservedDevice struct {
	Time    time.Time                                            `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time" xml:"time,attr"`
	ID      github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,2,opt,name=id,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"deviceID" xml:"id,attr"`
	Name    string                                               `protobuf:"bytes,3,opt,name=name,proto3" json:"name" xml:"name,attr"`
	Address string                                               `protobuf:"bytes,4,opt,name=address,proto3" json:"address" xml:"address,attr"`
}

func (m *ObservedDevice) Reset()         { *m = ObservedDevice{} }
func (m *ObservedDevice) String() string { return proto.CompactTextString(m) }
func (*ObservedDevice) ProtoMessage()    {}
func (*ObservedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_49f68ff7b178722f, []int{1}
}
func (m *ObservedDevice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObservedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObservedDevice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObservedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservedDevice.Merge(m, src)
}
func (m *ObservedDevice) XXX_Size() int {
	return m.ProtoSize()
}
func (m *ObservedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ObservedDevice proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ObservedFolder)(nil), "config.ObservedFolder")
	proto.RegisterType((*ObservedDevice)(nil), "config.ObservedDevice")
}

func init() { proto.RegisterFile("lib/config/observed.proto", fileDescriptor_49f68ff7b178722f) }

var fileDescriptor_49f68ff7b178722f = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x3f, 0x6f, 0xd4, 0x30,
	0x00, 0xc5, 0xe3, 0xf4, 0x68, 0x39, 0x53, 0xfe, 0x28, 0xd3, 0x71, 0x83, 0x5d, 0x9d, 0x32, 0x1c,
	0x02, 0x25, 0xfc, 0x9b, 0x10, 0x42, 0xe2, 0x74, 0x42, 0x3a, 0x75, 0x40, 0x8a, 0x98, 0x98, 0x48,
	0x62, 0x37, 0xb5, 0x94, 0x9c, 0xab, 0xc4, 0xad, 0xca, 0xc6, 0xc8, 0xd8, 0xf2, 0x09, 0xf8, 0x38,
	0xb7, 0x5d, 0x46, 0xc4, 0x60, 0xd4, 0xcb, 0x96, 0x31, 0x12, 0x3b, 0x8a, 0x9d, 0xf8, 0x6e, 0x42,
	0x4c, 0x6c, 0x7e, 0x4f, 0xcf, 0x3f, 0xf9, 0xbd, 0x28, 0xf0, 0x61, 0xca, 0x22, 0x3f, 0xe6, 0xcb,
	0x13, 0x96, 0xf8, 0x3c, 0x2a, 0x68, 0x7e, 0x41, 0x89, 0x77, 0x96, 0x73, 0xc1, 0x9d, 0x7d, 0x6d,
	0x8f, 0x71, 0xc2, 0x79, 0x92, 0x52, 0x5f, 0xb9, 0xd1, 0xf9, 0x89, 0x2f, 0x58, 0x46, 0x0b, 0x11,
	0x66, 0x67, 0x3a, 0x38, 0x1e, 0xd2, 0x4b, 0xa1, 0x8f, 0x93, 0xdf, 0x00, 0xde, 0x7b, 0xdf, 0x61,
	0xde, 0xf1, 0x94, 0xd0, 0xdc, 0xf9, 0x04, 0x07, 0xed, 0x85, 0x11, 0x38, 0x02, 0xd3, 0x3b, 0xcf,
	0xc7, 0x9e, 0xa6, 0x79, 0x3d, 0xcd, 0xfb, 0xd0, 0xd3, 0x66, 0x4f, 0x57, 0x12, 0x5b, 0xb5, 0xc4,
	0x2a, 0xdf, 0x48, 0x7c, 0xff, 0x32, 0x4b, 0x5f, 0x4d, 0x5a, 0xf1, 0x24, 0x14, 0x22, 0x9f, 0x5c,
	0xfd, 0xc2, 0xa0, 0x5e, 0xbb, 0x43, 0xe3, 0x04, 0x2a, 0xe9, 0xbc, 0x81, 0x36, 0x23, 0x23, 0xfb,
	0x08, 0x4c, 0x87, 0x33, 0x6f, 0x23, 0xb1, 0xbd, 0x98, 0xd7, 0x12, 0xdb, 0x8c, 0x34, 0x12, 0xdf,
	0x55, 0x0c, 0x46, 0x34, 0xa1, 0x5e, 0xbb, 0x07, 0xdd, 0xf9, 0x5b, 0xe9, 0xda, 0x8b, 0x79, 0x60,
	0x33, 0xe2, 0xbc, 0x85, 0xb7, 0xd2, 0x30, 0xa2, 0xe9, 0x68, 0x4f, 0x21, 0x1e, 0xd7, 0x12, 0x6b,
	0xa3, 0x91, 0xf8, 0x81, 0xba, 0xaf, 0x94, 0x41, 0xc0, 0xad, 0x0c, 0x74, 0x70, 0x72, 0xbd, 0xb7,
	0xed, 0x3d, 0xa7, 0x17, 0x2c, 0xa6, 0xff, 0xa1, 0xf7, 0x35, 0x30, 0xc5, 0x0f, 0x67, 0x5f, 0x40,
	0x4b, 0xf9, 0x29, 0xf1, 0xcb, 0x84, 0x89, 0xd3, 0xf3, 0xc8, 0x8b, 0x79, 0xe6, 0x17, 0x9f, 0x97,
	0xb1, 0x38, 0x65, 0xcb, 0x64, 0xe7, 0xd4, 0x7e, 0x70, 0xf5, 0x88, 0x98, 0xa7, 0x9e, 0x7e, 0xeb,
	0x62, 0x6e, 0x56, 0xbb, 0x4d, 0x3a, 0xe7, 0x6f, 0xdb, 0x35, 0x6b, 0xd7, 0xe4, 0xbe, 0x96, 0x2e,
	0xd8, 0xd9, 0xf2, 0x35, 0x1c, 0x2c, 0xc3, 0x8c, 0x76, 0x53, 0x4e, 0xdb, 0x56, 0xad, 0x36, 0xad,
	0x5a, 0x61, 0x78, 0x43, 0xa3, 0x02, 0x95, 0x72, 0x8e, 0xe1, 0x41, 0x48, 0x48, 0x4e, 0x8b, 0x62,
	0x34, 0x50, 0x80, 0x67, 0xb5, 0xc4, 0xbd, 0xd5, 0x48, 0xec, 0x28, 0x46, 0xa7, 0x0d, 0xe6, 0x70,
	0xd7, 0x08, 0xfa, 0xf8, 0xec, 0x78, 0x75, 0x83, 0xac, 0xf2, 0x06, 0x59, 0xab, 0x0d, 0x02, 0xe5,
	0x06, 0x81, 0xab, 0x0a, 0x59, 0xdf, 0x2b, 0x04, 0xca, 0x0a, 0x59, 0x3f, 0x2a, 0x64, 0x7d, 0x7c,
	0xf4, 0x0f, 0x53, 0xe9, 0x9f, 0x20, 0xda, 0x57, 0x93, 0xbd, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0xf0, 0x82, 0x78, 0x30, 0x03, 0x00, 0x00,
}

func (m *ObservedFolder) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObservedFolder) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObservedFolder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintObserved(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintObserved(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintObserved(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ObservedDevice) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObservedDevice) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObservedDevice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintObserved(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintObserved(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.ID.ProtoSize()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintObserved(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintObserved(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintObserved(dAtA []byte, offset int, v uint64) int {
	offset -= sovObserved(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObservedFolder) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovObserved(uint64(l))
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovObserved(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovObserved(uint64(l))
	}
	return n
}

func (m *ObservedDevice) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovObserved(uint64(l))
	l = m.ID.ProtoSize()
	n += 1 + l + sovObserved(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovObserved(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovObserved(uint64(l))
	}
	return n
}

func sovObserved(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObserved(x uint64) (n int) {
	return sovObserved(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObservedFolder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserved
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
			return fmt.Errorf("proto: ObservedFolder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObservedFolder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObserved(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObserved
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObserved
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
func (m *ObservedDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserved
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
			return fmt.Errorf("proto: ObservedDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObservedDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserved
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
				return ErrInvalidLengthObserved
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserved
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObserved(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObserved
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObserved
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
func skipObserved(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObserved
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
					return 0, ErrIntOverflowObserved
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
					return 0, ErrIntOverflowObserved
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
				return 0, ErrInvalidLengthObserved
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObserved
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObserved
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObserved        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObserved          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObserved = fmt.Errorf("proto: unexpected end of group")
)
