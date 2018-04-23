package classfile

type AttributeInfo interface {
	readerInfo(reader *ClassReader)
}

func readerAttributes(reader *ClassReader) []AttributeInfo {

}

func readerAttribute(reader *ClassReader) AttributeInfo {

}

func newAttributeInfo(reader *ClassReader) AttributeInfo {

}
