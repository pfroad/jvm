package classfile

type ConstantClass struct {
	ConstantPool
	classIndex uint16
}

func (cc *ConstantClass) readInfo(reader *ClassReader) {
	cc.classIndex = reader.readUint16()
}

func (cc *ConstantClass) String() string {
	return cc.ConstantPool.getUtf8(cc.classIndex)
}
