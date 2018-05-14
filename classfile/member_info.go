package classfile

// MemberInfo fields or method
/*field_info {
u2 access_flags;
u2 name_index;
u2 descriptor_index;
u2 attributes_count;
attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	constantPool    ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	// memberInfo := &MemberInfo{}
	// memberInfo.constantPool = cp
	// memberInfo.accessFlag = reader.readUint16()
	// memberInfo.nameIndex = reader.readUint16()
	// memberInfo.descriptorIndex = reader.readUint16()
	// memberInfo.attributes = readAttributes(reader, cp)
	return &MemberInfo{cp,
		reader.readUint16(),
		reader.readUint16(),
		reader.readUint16(),
		readAttributes(reader, cp)}
}

func (mi *MemberInfo) Name() string {
	return mi.constantPool.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.constantPool.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) AccessFlag() uint16 {
	return mi.accessFlag
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
