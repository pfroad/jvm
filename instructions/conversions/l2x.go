package conversions

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type L2F struct {
	common.NoOperandsInstruction
}

type L2I struct {
	common.NoOperandsInstruction
}

type L2D struct {
	common.NoOperandsInstruction
}

func (l2i *L2I) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushInt(int32(val))
}

func (l2f *L2F) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushFloat(float32(val))
}

func (l2d *L2D) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushDouble(float64(val))
}