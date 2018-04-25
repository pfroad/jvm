package classfile

type ConstantString struct {
	cp       ConstantPool
	strIndex uint16
}

func (cs *ConstantString) readInfo(reader *ClassReader) {
	cs.strIndex = reader.readUint16()
}

func (cs *ConstantString) String() string {
	return cs.cp.getUtf8(cs.strIndex)
}
