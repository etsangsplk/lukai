// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/tensorflow/core/protobuf/saver.proto

/*
	Package tensorflow is a generated protocol buffer package.

	It is generated from these files:
		protobuf/tensorflow/core/protobuf/saver.proto

	It has these top-level messages:
		SaverDef
*/
package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import fmt "fmt"
import strings "strings"
import reflect "reflect"

import math "math"

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

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	LEGACY SaverDef_CheckpointFormatVersion = 0
	// Deprecated format: tf.Saver() which works with tensorflow::table::Table.
	V1 SaverDef_CheckpointFormatVersion = 1
	// Current format: more efficient.
	V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}
var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorSaver, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName,proto3" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName,proto3" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName,proto3" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep,proto3" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours,proto3" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,proto3,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
}

func (m *SaverDef) Reset()                    { *m = SaverDef{} }
func (*SaverDef) ProtoMessage()               {}
func (*SaverDef) Descriptor() ([]byte, []int) { return fileDescriptorSaver, []int{0} }

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return LEGACY
}

func init() {
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
}
func (x SaverDef_CheckpointFormatVersion) String() string {
	s, ok := SaverDef_CheckpointFormatVersion_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *SaverDef) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SaverDef)
	if !ok {
		that2, ok := that.(SaverDef)
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
	if this.FilenameTensorName != that1.FilenameTensorName {
		return false
	}
	if this.SaveTensorName != that1.SaveTensorName {
		return false
	}
	if this.RestoreOpName != that1.RestoreOpName {
		return false
	}
	if this.MaxToKeep != that1.MaxToKeep {
		return false
	}
	if this.Sharded != that1.Sharded {
		return false
	}
	if this.KeepCheckpointEveryNHours != that1.KeepCheckpointEveryNHours {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *SaverDef) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&tensorflow.SaverDef{")
	s = append(s, "FilenameTensorName: "+fmt.Sprintf("%#v", this.FilenameTensorName)+",\n")
	s = append(s, "SaveTensorName: "+fmt.Sprintf("%#v", this.SaveTensorName)+",\n")
	s = append(s, "RestoreOpName: "+fmt.Sprintf("%#v", this.RestoreOpName)+",\n")
	s = append(s, "MaxToKeep: "+fmt.Sprintf("%#v", this.MaxToKeep)+",\n")
	s = append(s, "Sharded: "+fmt.Sprintf("%#v", this.Sharded)+",\n")
	s = append(s, "KeepCheckpointEveryNHours: "+fmt.Sprintf("%#v", this.KeepCheckpointEveryNHours)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSaver(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SaverDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SaverDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FilenameTensorName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.FilenameTensorName)))
		i += copy(dAtA[i:], m.FilenameTensorName)
	}
	if len(m.SaveTensorName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.SaveTensorName)))
		i += copy(dAtA[i:], m.SaveTensorName)
	}
	if len(m.RestoreOpName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSaver(dAtA, i, uint64(len(m.RestoreOpName)))
		i += copy(dAtA[i:], m.RestoreOpName)
	}
	if m.MaxToKeep != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSaver(dAtA, i, uint64(m.MaxToKeep))
	}
	if m.Sharded {
		dAtA[i] = 0x28
		i++
		if m.Sharded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.KeepCheckpointEveryNHours != 0 {
		dAtA[i] = 0x35
		i++
		i = encodeFixed32Saver(dAtA, i, uint32(math.Float32bits(float32(m.KeepCheckpointEveryNHours))))
	}
	if m.Version != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSaver(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func encodeFixed64Saver(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Saver(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSaver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SaverDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.FilenameTensorName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	l = len(m.SaveTensorName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	l = len(m.RestoreOpName)
	if l > 0 {
		n += 1 + l + sovSaver(uint64(l))
	}
	if m.MaxToKeep != 0 {
		n += 1 + sovSaver(uint64(m.MaxToKeep))
	}
	if m.Sharded {
		n += 2
	}
	if m.KeepCheckpointEveryNHours != 0 {
		n += 5
	}
	if m.Version != 0 {
		n += 1 + sovSaver(uint64(m.Version))
	}
	return n
}

func sovSaver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSaver(x uint64) (n int) {
	return sovSaver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SaverDef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SaverDef{`,
		`FilenameTensorName:` + fmt.Sprintf("%v", this.FilenameTensorName) + `,`,
		`SaveTensorName:` + fmt.Sprintf("%v", this.SaveTensorName) + `,`,
		`RestoreOpName:` + fmt.Sprintf("%v", this.RestoreOpName) + `,`,
		`MaxToKeep:` + fmt.Sprintf("%v", this.MaxToKeep) + `,`,
		`Sharded:` + fmt.Sprintf("%v", this.Sharded) + `,`,
		`KeepCheckpointEveryNHours:` + fmt.Sprintf("%v", this.KeepCheckpointEveryNHours) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSaver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SaverDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSaver
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
			return fmt.Errorf("proto: SaverDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SaverDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilenameTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilenameTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaveTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SaveTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestoreOpName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
				return ErrInvalidLengthSaver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RestoreOpName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxToKeep", wireType)
			}
			m.MaxToKeep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxToKeep |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sharded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
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
			m.Sharded = bool(v != 0)
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepCheckpointEveryNHours", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.KeepCheckpointEveryNHours = float32(math.Float32frombits(v))
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSaver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (SaverDef_CheckpointFormatVersion(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSaver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSaver
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
func skipSaver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSaver
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
					return 0, ErrIntOverflowSaver
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
					return 0, ErrIntOverflowSaver
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
				return 0, ErrInvalidLengthSaver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSaver
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
				next, err := skipSaver(dAtA[start:])
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
	ErrInvalidLengthSaver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSaver   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protobuf/tensorflow/core/protobuf/saver.proto", fileDescriptorSaver) }

var fileDescriptorSaver = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0x07, 0xf0, 0xbc, 0x1c, 0x97, 0xde, 0xf9, 0xc4, 0x11, 0x19, 0x24, 0xc2, 0x80, 0x15, 0xdd,
	0x80, 0x32, 0x40, 0x0a, 0x45, 0xec, 0xd0, 0xd2, 0x82, 0x04, 0x2a, 0x55, 0x5a, 0x55, 0x62, 0xb2,
	0xd2, 0xf6, 0x85, 0x46, 0x6d, 0xe2, 0xc8, 0x49, 0x4b, 0xd9, 0x58, 0xd9, 0xf8, 0x18, 0x7c, 0x14,
	0xc6, 0x8e, 0x8c, 0x34, 0x2c, 0x8c, 0x1d, 0x19, 0x91, 0x5d, 0xd2, 0x96, 0xe1, 0xa6, 0xc4, 0xfe,
	0xff, 0xfe, 0x96, 0xe5, 0x47, 0x1e, 0x65, 0x52, 0x14, 0x62, 0xb4, 0x88, 0xea, 0x05, 0xa6, 0xb9,
	0x90, 0xd1, 0x5c, 0x7c, 0xac, 0x8f, 0x85, 0xc4, 0xfa, 0x3e, 0xc8, 0xc3, 0x25, 0x4a, 0x5f, 0x2f,
	0x29, 0x39, 0xa8, 0xab, 0x2f, 0x27, 0xe4, 0xac, 0xaf, 0xb2, 0x97, 0x18, 0xd1, 0xc7, 0xe4, 0x4e,
	0x14, 0xcf, 0x31, 0x0d, 0x13, 0xe4, 0x3b, 0xc3, 0xd5, 0xbf, 0x03, 0x2e, 0x78, 0xe7, 0x01, 0xad,
	0xb2, 0x81, 0x8e, 0xba, 0x61, 0x82, 0xd4, 0x23, 0xb6, 0x3a, 0xf9, 0x3f, 0x6d, 0x6a, 0x7d, 0xa9,
	0xf6, 0x8f, 0xe4, 0x03, 0x72, 0x4b, 0x62, 0x5e, 0x08, 0x89, 0x5c, 0x64, 0x3b, 0x78, 0xa2, 0xe1,
	0xcd, 0x7f, 0xdb, 0xef, 0x32, 0xed, 0x18, 0xb9, 0x48, 0xc2, 0x15, 0x2f, 0x04, 0x9f, 0x21, 0x66,
	0xce, 0x0d, 0x17, 0xbc, 0xd3, 0xe0, 0x3c, 0x09, 0x57, 0x03, 0xf1, 0x06, 0x31, 0xa3, 0x0e, 0xa9,
	0xe5, 0xd3, 0x50, 0x4e, 0x70, 0xe2, 0x9c, 0xba, 0xe0, 0x9d, 0x05, 0xd5, 0x92, 0x3e, 0x27, 0xf7,
	0x55, 0x85, 0x8f, 0xa7, 0x38, 0x9e, 0x65, 0x22, 0x4e, 0x0b, 0x8e, 0x4b, 0x94, 0x9f, 0x78, 0xca,
	0xa7, 0x62, 0x21, 0x73, 0xc7, 0x72, 0xc1, 0x33, 0x83, 0x7b, 0x0a, 0xb5, 0xf6, 0xa6, 0xad, 0x48,
	0xf7, 0xb5, 0x02, 0xb4, 0x43, 0x6a, 0x4b, 0x94, 0x79, 0x2c, 0x52, 0xa7, 0xe6, 0x82, 0x77, 0xd9,
	0x78, 0xe8, 0x1f, 0x9e, 0xca, 0xaf, 0x9e, 0xc9, 0x3f, 0x94, 0x3b, 0x42, 0x26, 0x61, 0x31, 0xdc,
	0x75, 0x82, 0xaa, 0x7c, 0xf5, 0x8c, 0xdc, 0xbd, 0xc6, 0x50, 0x42, 0xac, 0xb7, 0xed, 0x57, 0x2f,
	0x5a, 0xef, 0x6d, 0x83, 0x5a, 0xc4, 0x1c, 0x3e, 0xb1, 0x41, 0x7f, 0x1b, 0xb6, 0xd9, 0xec, 0xaf,
	0x37, 0xcc, 0xf8, 0xb1, 0x61, 0xc6, 0x76, 0xc3, 0xe0, 0x73, 0xc9, 0xe0, 0x5b, 0xc9, 0xe0, 0x7b,
	0xc9, 0x60, 0x5d, 0x32, 0xf8, 0x59, 0x32, 0xf8, 0x5d, 0x32, 0x63, 0x5b, 0x32, 0xf8, 0xfa, 0x8b,
	0x19, 0xe4, 0xb6, 0x90, 0x1f, 0x8e, 0xaf, 0xb7, 0x28, 0xe2, 0x79, 0xf3, 0x42, 0x5f, 0xb2, 0xa7,
	0xc6, 0x9c, 0xf7, 0xe0, 0x0f, 0xc0, 0xc8, 0xd2, 0x33, 0x7f, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0x11, 0x60, 0x54, 0x24, 0x02, 0x00, 0x00,
}
