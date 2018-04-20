package classfile

// ConstantPool class file constant pool
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := reader.readUint16()
	cp := make([]ConstantInfo, cpCount)
	for i := range cp {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

func (cp ConstantPool) getUtf8(index uint16) string {

}
