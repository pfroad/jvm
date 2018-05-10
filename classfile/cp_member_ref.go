package classfile

type ConstantMemberRef struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldRef struct{ ConstantMemberRef }
type ConstantMethodRef struct{ ConstantMemberRef }
type ConstantInterfaceMethodRef struct{ ConstantMemberRef }

func (cmr *ConstantMemberRef) readInfo(reader *ClassReader) {
	cmr.classIndex = reader.readUint16()
	cmr.nameAndTypeIndex = reader.readUint16()
}

func (cmr *ConstantMemberRef) ClassName() string {
	return cmr.cp.getClassName(cmr.classIndex)
}

func (cmr *ConstantMemberRef) NameAndDescriptor() (string, string) {
	return cmr.cp.getNameAndType(cmr.nameAndTypeIndex)
}
