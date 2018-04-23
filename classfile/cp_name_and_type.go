package classfile

type ConstantNameAndType struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cnt *ConstantNameAndType) readInfo(reader *ClassReader) {
	cnt.nameIndex = reader.readUint16()
	cnt.descriptorIndex = reader.readUint16()
}
