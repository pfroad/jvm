package classfile

// ConstantPool class file constant pool
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		//fmt.Printf("cp index: %d\n", i)
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLong, *ConstantDouble:
			i++
		}
	}

	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getUtf8(index uint16) string {
	return cp.getConstantInfo(index).(*ConstantUtf8).str
}

func (cp ConstantPool) getClassName(classIndex uint16) string {
	return cp.getUtf8(cp.getConstantInfo(classIndex).(*ConstantClass).classIndex)
}

func (cp ConstantPool) getNameAndType(nameAndTypeIndex uint16) (string, string) {
	nameAndType := cp.getConstantInfo(nameAndTypeIndex).(*ConstantNameAndType)
	return cp.getUtf8(nameAndType.nameIndex), cp.getUtf8(nameAndType.descriptorIndex)
}
