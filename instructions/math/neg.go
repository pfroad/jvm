package math

import "jvm/instructions/common"
import "jvm/runtime"

// negate
type INeg struct {
	common.NoOperandsInstruction
}

type LNeg struct {
	common.NoOperandsInstruction
}

type FNeg struct {
	common.NoOperandsInstruction
}

type DNeg struct {
	common.NoOperandsInstruction
}

func (neg *INeg) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (neg *LNeg) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

func (neg *FNeg) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (neg *DNeg) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}