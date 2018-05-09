package classfile

import "math"

type ConstantInteger struct {
	val int32
}

func (cInt *ConstantInteger) readInfo(reader *ClassReader) {
	cInt.val = int32(reader.readUint32())
}

func (cInt *ConstantInteger) Value() int32 {
	return cInt.val
}

type ConstantFloat struct {
	val float32
}

func (cFloat *ConstantFloat) readInfo(reader *ClassReader) {
	cFloat.val = math.Float32frombits(reader.readUint32())
}

func (cFloat *ConstantFloat) Value() float32 {
	return cFloat.val
}

type ConstantLong struct {
	val int64
}

func (cLong *ConstantLong) readInfo(reader *ClassReader) {
	cLong.val = int64(reader.readUint64())
}

func (cLong *ConstantLong) Value() int64 {
	return cLong.val
}

type ConstantDouble struct {
	val float64
}

func (cDouble *ConstantDouble) readInfo(reader *ClassReader) {
	cDouble.val = math.Float64frombits(reader.readUint64())
}

func (cDouble *ConstantDouble) Value() float64 {
	return cDouble.val
}