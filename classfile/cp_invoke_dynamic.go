package classfile

/*CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}*/
type ConstantMethodHandle struct {
	refKind  uint8
	refIndex uint16
}

func (mh *ConstantMethodHandle) readInfo(reader *ClassReader) {
	mh.refKind = reader.readUint8()
	mh.refIndex = reader.readUint16()
}

/*CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}*/
type ConstantMethodType struct {
	descriptorIndex uint16
}

func (mt *ConstantMethodType) readInfo(reader *ClassReader) {
	mt.descriptorIndex = reader.readUint16()
}

/*CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}*/
type ConstantInvokeDynamic struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cid *ConstantInvokeDynamic) readInfo(reader *ClassReader) {
	cid.bootstrapMethodAttrIndex = reader.readUint16()
	cid.nameAndTypeIndex = reader.readUint16()
}
