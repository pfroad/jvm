package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPC    uint16
	lineNumber uint16
}

func (lnt *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberLength := reader.readUint16()
	lnt.lineNumberTable = readLineNumberTable(reader, lineNumberLength)
}

func readLineNumberTable(reader *ClassReader, lineNumberLength uint16) []*LineNumberTableEntry {
	lineNumberTable := make([]*LineNumberTableEntry, lineNumberLength)
	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{reader.readUint16(), reader.readUint16()}
	}
	return lineNumberTable
}
