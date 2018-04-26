package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPC         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (lvt *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	lvtLength := reader.readUint16()
	lvt.localVariableTable = readLocalVariableTable(reader, lvtLength)
}

func readLocalVariableTable(reader *ClassReader, lvtLength uint16) []*LocalVariableTableEntry {
	localVariableTable := make([]*LocalVariableTableEntry, lvtLength)
	for i := range localVariableTable {
		localVariableTable[i] = &LocalVariableTableEntry{reader.readUint16(),
		reader.readUint16(),
		reader.readUint16(),
		reader.readUint16(),
		reader.readUint16()}
	}
	return localVariableTable
}
