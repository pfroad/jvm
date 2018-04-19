package classfile

// ConstantPool class file constant pool
type ConstantPool []ConstantPoolInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := reader.readUint16()
	cp := make([]ConstantPoolInfo, cpCount)
	for i := range cp {
		cp[i] =
	}
}
