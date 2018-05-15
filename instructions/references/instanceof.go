package references

import (
	"jvm/instructions/common"
	"jvm/runtime/data"
	"jvm/runtime"
)

/**
java code:
if (obj instanceof ClassYYY) {...}
 */
type InstanceOf struct {
	common.Index16Instruction
}

func (i *InstanceOf) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConst(i.Index).(*data.ClassRef)
	class := classRef.ResolveClass()

	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
