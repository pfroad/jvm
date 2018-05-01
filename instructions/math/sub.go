package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type ISub struct {
	common.NoOperandsInstruction
}

type LSub struct {
	common.NoOperandsInstruction
}

type FSub struct {
	common.NoOperandsInstruction
}

type DSub struct {
	common.NoOperandsInstruction
}

func (sub *ISub) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 - v2)
}

func (sub *LSub) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 - v2)
}

func (sub *FSub) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	stack.PushFloat(v1 - v2)
}

func (sub *DSub) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	stack.PushDouble(v1 - v2)
}
