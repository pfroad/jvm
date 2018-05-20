package references

import (
	"jvm/runtime"
	"jvm/runtime/data"
)

type NewArray struct {
	atype uint8
}

func (arr *NewArray) FetchOperands(reader *runtime.ByteCodeReader) {
	arr.atype = reader.ReadUint8()
}

func (newArray *NewArray) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()

	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().ClassLoader()
	arrClass := getPrimitiveArrayClass(classLoader, newArray.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(classLoader *data.ClassLoader, atype uint8) *data.Class {
	switch atype {
	case AT_BOOLEAN:
		return classLoader.LoadClass("[Z")
	case AT_BYTE:
		return classLoader.LoadClass("[B")
	case AT_CHAR:
		return classLoader.LoadClass("[C")
	case AT_FLOAT:
		return classLoader.LoadClass("[F")
	case AT_DOUBLE:
		return classLoader.LoadClass("[D")
	case AT_INT:
		return classLoader.LoadClass("[I")
	case AT_SHORT:
		return classLoader.LoadClass("[S")
	case AT_LONG:
		return classLoader.LoadClass("[J")
	default:
		panic("Invalid atype!")
	}
}
