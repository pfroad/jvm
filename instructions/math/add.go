package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// add
type IAdd struct {
	common.NoOperandsInstruction
}

type LAdd struct {
	common.NoOperandsInstruction
}

type FAdd struct {
	common.NoOperandsInstruction
}

type DAdd struct {
	common.NoOperandsInstruction
}

func (add *IAdd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (add *LAdd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

func (add *FAdd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (add *DAdd) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}
