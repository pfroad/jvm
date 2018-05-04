package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IOR struct {
	common.NoOperandsInstruction
}

type LOR struct {
	common.NoOperandsInstruction
}

func (ior *IOR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 | v2)
}

func (lor *LOR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 | v2)
}
