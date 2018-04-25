package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPC   uint16
	endPC     uint16
	handlePC  uint16
	catchType uint16
}

func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	ca.maxStack = reader.readUint16()
	ca.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	ca.code = reader.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader, ca.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{reader.readUint16(),
			reader.readUint16(),
			reader.readUint16(),
			reader.readUint16()}
	}

	return exceptionTable
}
