package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type ArrayLength struct {
	common.NoOperandsInstruction
}

func (arr *ArrayLength) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	stack.PushInt(arrRef.ArrayLength())
}