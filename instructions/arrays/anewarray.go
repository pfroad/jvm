package arrays

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type ANewArray struct {
	common.Index16Instruction
}

func (arr *ANewArray) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()

	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	class := frame.Method().Class()
	cp := class.ConstantPool()
	classRef := cp.GetConst(arr.Index).(*data.ClassRef)
	resolveClass := classRef.ResolveClass()
	componentClass := resolveClass.ArrayClass()
	// classLoader := class.ClassLoader()
	stack.PushRef(componentClass.NewArray(uint(count)))
}
