package classfile

type ConstantClass struct {
	cp         ConstantPool
	classIndex uint16
}

func (cc *ConstantClass) readInfo(reader *ClassReader) {
	cc.classIndex = reader.readUint16()
}

func (cc *ConstantClass) Name() string {
	return cc.cp.getUtf8(cc.classIndex)
}
