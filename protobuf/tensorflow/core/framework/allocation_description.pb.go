// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/tensorflow/core/framework/allocation_description.proto

/*
	Package tensorflow is a generated protocol buffer package.

	It is generated from these files:
		protobuf/tensorflow/core/framework/allocation_description.proto

	It has these top-level messages:
		AllocationDescription
*/
package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import fmt "fmt"
import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
}

func (m *AllocationDescription) Reset()      { *m = AllocationDescription{} }
func (*AllocationDescription) ProtoMessage() {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptorAllocationDescription, []int{0}
}

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}
func (this *AllocationDescription) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AllocationDescription)
	if !ok {
		that2, ok := that.(AllocationDescription)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.RequestedBytes != that1.RequestedBytes {
		return false
	}
	if this.AllocatedBytes != that1.AllocatedBytes {
		return false
	}
	if this.AllocatorName != that1.AllocatorName {
		return false
	}
	if this.AllocationId != that1.AllocationId {
		return false
	}
	if this.HasSingleReference != that1.HasSingleReference {
		return false
	}
	if this.Ptr != that1.Ptr {
		return false
	}
	return true
}
func (this *AllocationDescription) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&tensorflow.AllocationDescription{")
	s = append(s, "RequestedBytes: "+fmt.Sprintf("%#v", this.RequestedBytes)+",\n")
	s = append(s, "AllocatedBytes: "+fmt.Sprintf("%#v", this.AllocatedBytes)+",\n")
	s = append(s, "AllocatorName: "+fmt.Sprintf("%#v", this.AllocatorName)+",\n")
	s = append(s, "AllocationId: "+fmt.Sprintf("%#v", this.AllocationId)+",\n")
	s = append(s, "HasSingleReference: "+fmt.Sprintf("%#v", this.HasSingleReference)+",\n")
	s = append(s, "Ptr: "+fmt.Sprintf("%#v", this.Ptr)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAllocationDescription(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AllocationDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocationDescription) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestedBytes != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.RequestedBytes))
	}
	if m.AllocatedBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.AllocatedBytes))
	}
	if len(m.AllocatorName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(len(m.AllocatorName)))
		i += copy(dAtA[i:], m.AllocatorName)
	}
	if m.AllocationId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.AllocationId))
	}
	if m.HasSingleReference {
		dAtA[i] = 0x28
		i++
		if m.HasSingleReference {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Ptr != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.Ptr))
	}
	return i, nil
}

func encodeFixed64AllocationDescription(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32AllocationDescription(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAllocationDescription(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AllocationDescription) Size() (n int) {
	var l int
	_ = l
	if m.RequestedBytes != 0 {
		n += 1 + sovAllocationDescription(uint64(m.RequestedBytes))
	}
	if m.AllocatedBytes != 0 {
		n += 1 + sovAllocationDescription(uint64(m.AllocatedBytes))
	}
	l = len(m.AllocatorName)
	if l > 0 {
		n += 1 + l + sovAllocationDescription(uint64(l))
	}
	if m.AllocationId != 0 {
		n += 1 + sovAllocationDescription(uint64(m.AllocationId))
	}
	if m.HasSingleReference {
		n += 2
	}
	if m.Ptr != 0 {
		n += 1 + sovAllocationDescription(uint64(m.Ptr))
	}
	return n
}

func sovAllocationDescription(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAllocationDescription(x uint64) (n int) {
	return sovAllocationDescription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AllocationDescription) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AllocationDescription{`,
		`RequestedBytes:` + fmt.Sprintf("%v", this.RequestedBytes) + `,`,
		`AllocatedBytes:` + fmt.Sprintf("%v", this.AllocatedBytes) + `,`,
		`AllocatorName:` + fmt.Sprintf("%v", this.AllocatorName) + `,`,
		`AllocationId:` + fmt.Sprintf("%v", this.AllocationId) + `,`,
		`HasSingleReference:` + fmt.Sprintf("%v", this.HasSingleReference) + `,`,
		`Ptr:` + fmt.Sprintf("%v", this.Ptr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAllocationDescription(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AllocationDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocationDescription
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllocationDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocationDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedBytes", wireType)
			}
			m.RequestedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatedBytes", wireType)
			}
			m.AllocatedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocatedBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAllocationDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationId", wireType)
			}
			m.AllocationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocationId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasSingleReference", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasSingleReference = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ptr", wireType)
			}
			m.Ptr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ptr |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAllocationDescription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAllocationDescription
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
func skipAllocationDescription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllocationDescription
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
					return 0, ErrIntOverflowAllocationDescription
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
					return 0, ErrIntOverflowAllocationDescription
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAllocationDescription
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAllocationDescription
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
				next, err := skipAllocationDescription(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthAllocationDescription = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllocationDescription   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("protobuf/tensorflow/core/framework/allocation_description.proto", fileDescriptorAllocationDescription)
}

var fileDescriptorAllocationDescription = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0x33, 0xb6, 0x16, 0x5d, 0xac, 0x4a, 0x50, 0x08, 0x08, 0x4b, 0x50, 0xc4, 0x9c, 0x1a,
	0xc1, 0x0f, 0x10, 0x8b, 0x17, 0x2f, 0x52, 0xe2, 0x07, 0x84, 0x6d, 0x32, 0x69, 0x83, 0x69, 0xb6,
	0xce, 0xa6, 0x14, 0x6f, 0x7e, 0x82, 0x9f, 0xe1, 0xa7, 0x78, 0xec, 0xd1, 0xa3, 0x59, 0x2f, 0x1e,
	0x7b, 0xd4, 0x9b, 0x24, 0xb5, 0x1b, 0x0f, 0xde, 0x86, 0x37, 0x6f, 0x92, 0xe5, 0xb1, 0xcb, 0x29,
	0xc9, 0x42, 0x0e, 0x67, 0x89, 0x5f, 0x60, 0xae, 0x24, 0x25, 0x99, 0x9c, 0xfb, 0x91, 0x24, 0xf4,
	0x13, 0x12, 0x13, 0x9c, 0x4b, 0xba, 0xf7, 0x45, 0x96, 0xc9, 0x48, 0x14, 0xa9, 0xcc, 0xc3, 0x18,
	0x55, 0x44, 0xe9, 0xb4, 0x9a, 0x7b, 0xf5, 0xa5, 0xcd, 0x9a, 0xbb, 0xe3, 0x6f, 0x60, 0x87, 0x57,
	0x46, 0xbe, 0x6e, 0x5c, 0xfb, 0x8c, 0xed, 0x11, 0x3e, 0xcc, 0x50, 0x15, 0x18, 0x87, 0xc3, 0xc7,
	0x02, 0x95, 0x03, 0x2e, 0x78, 0xad, 0x60, 0xd7, 0xe0, 0x7e, 0x45, 0x2b, 0xf1, 0xf7, 0x77, 0x46,
	0xdc, 0x58, 0x89, 0x06, 0xaf, 0xc4, 0x53, 0xb6, 0x26, 0x92, 0xc2, 0x5c, 0x4c, 0xd0, 0x69, 0xb9,
	0xe0, 0x6d, 0x07, 0x5d, 0x43, 0x6f, 0xc5, 0x04, 0xed, 0x13, 0xd6, 0xfd, 0xf3, 0xfc, 0x34, 0x76,
	0xda, 0xf5, 0xd7, 0x76, 0x1a, 0x78, 0x13, 0xdb, 0xe7, 0xec, 0x60, 0x2c, 0x54, 0xa8, 0xd2, 0x7c,
	0x94, 0x61, 0x48, 0x98, 0x20, 0x61, 0x1e, 0xa1, 0xb3, 0xe9, 0x82, 0xb7, 0x15, 0xd8, 0x63, 0xa1,
	0xee, 0xea, 0x55, 0xb0, 0xde, 0xd8, 0xfb, 0xac, 0x35, 0x2d, 0xc8, 0xe9, 0xb8, 0xe0, 0xb5, 0x83,
	0x6a, 0xec, 0x8f, 0x17, 0x25, 0xb7, 0xde, 0x4a, 0x6e, 0x2d, 0x4b, 0x0e, 0x4f, 0x9a, 0xc3, 0x8b,
	0xe6, 0xf0, 0xaa, 0x39, 0x2c, 0x34, 0x87, 0x77, 0xcd, 0xe1, 0x53, 0x73, 0x6b, 0xa9, 0x39, 0x3c,
	0x7f, 0x70, 0x8b, 0x39, 0x92, 0x46, 0xbd, 0xa6, 0x5c, 0xcf, 0xc4, 0xee, 0x1f, 0xfd, 0x1b, 0x70,
	0x50, 0xb5, 0x56, 0x03, 0xf8, 0x02, 0x18, 0x76, 0xea, 0xf0, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x88, 0x72, 0x0b, 0xd3, 0xbb, 0x01, 0x00, 0x00,
}
