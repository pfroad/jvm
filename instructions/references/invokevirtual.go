package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
	"fmt"
)

type InvokeVirtual struct {
	common.Index16Instruction
}

// hack
func (i *InvokeVirtual) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConst(i.Index).(*data.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V": fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V": fmt.Printf("%c\n", stack.PopInt())
		case "(B)V": fmt.Printf("%v\n", stack.PopInt())
		case "(S)V": fmt.Printf("%v\n", stack.PopInt())
		case "(I)V": fmt.Printf("%v\n", stack.PopInt())
		case "(J)V": fmt.Printf("%v\n", stack.PopLong())
		case "(F)V": fmt.Printf("%v\n", stack.PopFloat())
		case "(D)V": fmt.Printf("%v\n", stack.PopDouble())
		default: panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}