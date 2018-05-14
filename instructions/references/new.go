package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type New struct {
	common.Index16Instruction
}

func (n *New) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConst(n.Index).(*data.ClassRef)
	class := classRef.ResolveClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	obj := class.NewObject()
	frame.OperandStack().PushRef(obj)
}
