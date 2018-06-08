package byteutils

import "encoding/binary"

// Uint64 encodes []byte.
func Uint64(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

// FromUint64 decodes unit64 value.
func FromUint64(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// Uint32 encodes []byte.
func Uint32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// FromUint32 decodes uint32.
func FromUint32(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

// Uint16 encodes []byte.
func Uint16(data []byte) uint16 {
	return binary.BigEndian.Uint16(data)
}

// FromUint16 decodes uint16.
func FromUint16(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

// Int64 encodes []byte.
func Int64(data []byte) int64 {
	return int64(binary.BigEndian.Uint64(data))
}

// FromInt64 decodes int64 v.
func FromInt64(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Int32 encodes []byte.
func Int32(data []byte) int32 {
	return int32(binary.BigEndian.Uint32(data))
}

// FromInt32 decodes int32 v.
func FromInt32(v int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	return b
}

// Int16 encode []byte.
func Int16(data []byte) int16 {
	return int16(binary.BigEndian.Uint16(data))
}

// FromInt16 decodes int16 v.
func FromInt16(v int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))
	return b
}