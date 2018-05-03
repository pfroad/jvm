package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// Div
type IDiv struct {
	common.NoOperandsInstruction
}

type LDiv struct {
	common.NoOperandsInstruction
}

type FDiv struct {
	common.NoOperandsInstruction
}

type DDiv struct {
	common.NoOperandsInstruction
}

func (div *IDiv) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

func (div *LDiv) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}

func (div *FDiv) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

func (div *DDiv) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}