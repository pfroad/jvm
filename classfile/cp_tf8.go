package classfile

type ConstantUtf8 struct {
	str string
}

func (cUtf8 *ConstantUtf8) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	cUtf8.str = decodeMutf8(reader.readBytes(length))
}

func decodeMutf8(bytes []byte) string {
	return string(bytes)
}