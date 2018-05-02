package conversions

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type F2I struct {
	common.NoOperandsInstruction
}

type F2L struct {
	common.NoOperandsInstruction
}

type F2D struct {
	common.NoOperandsInstruction
}

func (f2i *F2I) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushInt(int32(val))
}

func (f2l *F2L) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushLong(int64(val))
}

func (f2D *F2D) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushDouble(float64(val))
}