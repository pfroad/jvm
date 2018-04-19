package classfile

// ConstantPool class file constant pool
type ConstantPool []ConstantPoolInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := reader.readUint16()
	for i := range cpCount - 1 {
		
	}
}
