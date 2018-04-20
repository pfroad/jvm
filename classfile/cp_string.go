package classfile

type ConstantString struct {
	ConstantPool
	strIndex uint16
}

func (cs *ConstantString) readInfo(reader *ClassReader) {
	cs.strIndex = reader.readUint16()
}

func (cs *ConstantString) String() string {
	return cs.getUtf8(cs.strIndex)
}