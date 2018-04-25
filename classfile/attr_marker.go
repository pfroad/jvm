package classfile

type DeprecatedAttribute struct {
	MakerAttribute
}

type SyntheticAttribute struct {
	MakerAttribute
}

type MakerAttribute struct {}

func (ma *MakerAttribute) readInfo(reader *ClassReader) {
}