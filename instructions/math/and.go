package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IAnd struct {
	common.NoOperandsInstruction
}

type LAnd struct {
	common.NoOperandsInstruction
}

func (ia *IAnd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 & v2)
}

func (la *LAnd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 & v2)
}
