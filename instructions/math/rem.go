package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"math"
)

// remainder
type IRem struct {
	common.NoOperandsInstruction
}

type LRem struct {
	common.NoOperandsInstruction
}

type FRem struct {
	common.NoOperandsInstruction
}

type DRem struct {
	common.NoOperandsInstruction
}

func (rem *IRem) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

func (rem *LRem) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}

func (rem *FRem) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}

func (rem *DRem) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
