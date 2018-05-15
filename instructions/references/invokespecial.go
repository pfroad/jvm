package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type InvokeSpecial struct {
	common.Index16Instruction
}

// hack
func (i *InvokeSpecial) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopRef()
}