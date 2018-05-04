package conversions

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type D2I struct {
	common.NoOperandsInstruction
}

type D2L struct {
	common.NoOperandsInstruction
}

type D2F struct {
	common.NoOperandsInstruction
}

func (d2i *D2I) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushInt(int32(val))
}

func (d2l *D2L) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushLong(int64(val))
}

func (d2f *D2F) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushFloat(float32(val))
}
