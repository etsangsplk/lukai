// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/tensorflow/core/framework/resource_handle.proto

/*
	Package tensorflow is a generated protocol buffer package.

	It is generated from these files:
		protobuf/tensorflow/core/framework/resource_handle.proto

	It has these top-level messages:
		ResourceHandleProto
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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandleProto) Reset()      { *m = ResourceHandleProto{} }
func (*ResourceHandleProto) ProtoMessage() {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptorResourceHandle, []int{0}
}

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}
func (this *ResourceHandleProto) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ResourceHandleProto)
	if !ok {
		that2, ok := that.(ResourceHandleProto)
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
	if this.Device != that1.Device {
		return false
	}
	if this.Container != that1.Container {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.HashCode != that1.HashCode {
		return false
	}
	if this.MaybeTypeName != that1.MaybeTypeName {
		return false
	}
	return true
}
func (this *ResourceHandleProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&tensorflow.ResourceHandleProto{")
	s = append(s, "Device: "+fmt.Sprintf("%#v", this.Device)+",\n")
	s = append(s, "Container: "+fmt.Sprintf("%#v", this.Container)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "HashCode: "+fmt.Sprintf("%#v", this.HashCode)+",\n")
	s = append(s, "MaybeTypeName: "+fmt.Sprintf("%#v", this.MaybeTypeName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringResourceHandle(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ResourceHandleProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHandleProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	if len(m.Container) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Container)))
		i += copy(dAtA[i:], m.Container)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.HashCode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(m.HashCode))
	}
	if len(m.MaybeTypeName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.MaybeTypeName)))
		i += copy(dAtA[i:], m.MaybeTypeName)
	}
	return i, nil
}

func encodeFixed64ResourceHandle(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32ResourceHandle(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintResourceHandle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceHandleProto) Size() (n int) {
	var l int
	_ = l
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Container)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	if m.HashCode != 0 {
		n += 1 + sovResourceHandle(uint64(m.HashCode))
	}
	l = len(m.MaybeTypeName)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	return n
}

func sovResourceHandle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceHandle(x uint64) (n int) {
	return sovResourceHandle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ResourceHandleProto) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResourceHandleProto{`,
		`Device:` + fmt.Sprintf("%v", this.Device) + `,`,
		`Container:` + fmt.Sprintf("%v", this.Container) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`HashCode:` + fmt.Sprintf("%v", this.HashCode) + `,`,
		`MaybeTypeName:` + fmt.Sprintf("%v", this.MaybeTypeName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringResourceHandle(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ResourceHandleProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceHandle
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
			return fmt.Errorf("proto: ResourceHandleProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHandleProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashCode", wireType)
			}
			m.HashCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HashCode |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaybeTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaybeTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceHandle
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
func skipResourceHandle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
				return 0, ErrInvalidLengthResourceHandle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceHandle
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
				next, err := skipResourceHandle(dAtA[start:])
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
	ErrInvalidLengthResourceHandle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceHandle   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("protobuf/tensorflow/core/framework/resource_handle.proto", fileDescriptorResourceHandle)
}

var fileDescriptorResourceHandle = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x33, 0x1a, 0x0f, 0xb3, 0xa0, 0xc2, 0x0a, 0x12, 0x50, 0x86, 0xc3, 0x42, 0xae, 0x4a,
	0x0a, 0x1b, 0xeb, 0xb3, 0xb1, 0x12, 0x09, 0x76, 0x16, 0x61, 0x93, 0x4c, 0xcc, 0xe1, 0x25, 0x13,
	0x36, 0x39, 0x8f, 0x74, 0xfe, 0x04, 0x7f, 0x84, 0x85, 0x3f, 0xc5, 0xf2, 0x4a, 0x4b, 0xb3, 0x36,
	0x96, 0x57, 0x5a, 0x8a, 0xcb, 0x61, 0xc0, 0x6e, 0xe6, 0xbd, 0xf7, 0x15, 0xef, 0x89, 0x8b, 0x5a,
	0x73, 0xcb, 0xc9, 0x22, 0x0f, 0x5b, 0xaa, 0x1a, 0xd6, 0xf9, 0x9c, 0x97, 0x61, 0xca, 0x9a, 0xc2,
	0x5c, 0xab, 0x92, 0x96, 0xac, 0x1f, 0x42, 0x4d, 0x0d, 0x2f, 0x74, 0x4a, 0x71, 0xa1, 0xaa, 0x6c,
	0x4e, 0x81, 0x45, 0xa4, 0x18, 0x80, 0xd3, 0x17, 0x10, 0x87, 0xd1, 0x26, 0x75, 0x65, 0x43, 0x37,
	0x36, 0x73, 0x24, 0x46, 0x19, 0x3d, 0xce, 0x52, 0xf2, 0x61, 0x0c, 0x13, 0x2f, 0xda, 0x7c, 0xf2,
	0x44, 0x78, 0x29, 0x57, 0xad, 0x9a, 0x55, 0xa4, 0xfd, 0x2d, 0x6b, 0x0d, 0x82, 0x94, 0xc2, 0xad,
	0x54, 0x49, 0xfe, 0xb6, 0x35, 0xec, 0x2d, 0x8f, 0x85, 0x57, 0xa8, 0xa6, 0x88, 0x53, 0xce, 0xc8,
	0x77, 0xc7, 0x30, 0x71, 0xa3, 0xdd, 0x5f, 0xe1, 0x92, 0x33, 0x92, 0x67, 0xe2, 0xa0, 0x54, 0x5d,
	0x42, 0x71, 0xdb, 0xd5, 0x14, 0x5b, 0x76, 0xc7, 0xb2, 0x7b, 0x56, 0xbe, 0xed, 0x6a, 0xba, 0x56,
	0x25, 0x4d, 0xef, 0x56, 0x3d, 0x3a, 0xef, 0x3d, 0x3a, 0xeb, 0x1e, 0xe1, 0xc9, 0x20, 0xbc, 0x1a,
	0x84, 0x37, 0x83, 0xb0, 0x32, 0x08, 0x1f, 0x06, 0xe1, 0xcb, 0xa0, 0xb3, 0x36, 0x08, 0xcf, 0x9f,
	0xe8, 0x08, 0x9f, 0xf5, 0x7d, 0x30, 0x94, 0x0c, 0xfe, 0x06, 0x99, 0xee, 0xff, 0xeb, 0x0a, 0xdf,
	0x00, 0xc9, 0xc8, 0xce, 0x72, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x16, 0x9c, 0x6c, 0xa6, 0x52,
	0x01, 0x00, 0x00,
}
