package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// Mul
type IMul struct {
	common.NoOperandsInstruction
}

type LMul struct {
	common.NoOperandsInstruction
}

type FMul struct {
	common.NoOperandsInstruction
}

type DMul struct {
	common.NoOperandsInstruction
}

func (mul *IMul) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

func (mul *LMul) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}

func (mul *FMul) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

func (mul *DMul) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}
