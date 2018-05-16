package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type InvokeStatic struct {
	common.Index16Instruction
}

func (i *InvokeStatic) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConst(i.Index).(*data.MethodRef)
	method := methodRef.ResolveMethod()

	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	InvokeMethod(method, frame)
}
