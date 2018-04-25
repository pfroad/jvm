package classfile

// constantvalue attribute length is always 2
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (cva *ConstantValueAttribute) readInfo(reader *ClassReader) {
	cva.constantValueIndex = reader.readUint16()
}

func (cva *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return cva.constantValueIndex
}
