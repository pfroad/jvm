package classfile

import "encoding/binary"

// ClassReader class file reader
type ClassReader struct {
	data []byte
}

// ClassFile u1
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// ClassFile u2
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// ClassFile u4
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// method and fields: count(u2 type) and arrays[element type u2]
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	u2s := make([]uint16, n)
	for i := range u2s {
		u2s[i] = cr.readUint16()
	}
	return u2s
}

// read n bytes
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
