package common

type BytecodeReader struct {
	code []byte
	pc   int
}

func (reader *BytecodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.pc = pc
}

func (reader *BytecodeReader) ReadUint8() uint8 {
	val := reader.code[reader.pc]
	reader.pc++
	return val
}

func (reader *BytecodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *BytecodeReader) ReadUint16() uint16 {
	high := uint16(reader.ReadUint8())
	low := uint16(reader.ReadUint8())
	return (high << 8) | low
}

func (reader *BytecodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *BytecodeReader) ReadInt32() int32 {
	c1 := int32(reader.ReadUint8())
	c2 := int32(reader.ReadUint8())
	c3 := int32(reader.ReadUint8())
	c4 := int32(reader.ReadUint8())
	return (c1 << 24) | (c2 << 16) | (c3 << 8) | c4
}

func (reader *BytecodeReader) PC() int {
	return reader.pc
}
func (reader *BytecodeReader) ReadInt32s(len int32) []int32 {
	ints := make([]int32, len)
	for i := range ints {
		ints[i] = reader.ReadInt32()
	}
	return ints
}