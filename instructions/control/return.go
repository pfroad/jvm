package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type Return struct {
	common.NoOperandsInstruction
}

func (r *Return) Execute(frame *runtime.Frame) {
	frame.Thread().PopFrame()
}

type AReturn struct {
	common.NoOperandsInstruction
}

func (ar *AReturn) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(val)
}

type IReturn struct {
	common.NoOperandsInstruction
}

func (ir *IReturn) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(val)
}

type FReturn struct {
	common.NoOperandsInstruction
}

func (fr *FReturn) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(val)
}

type LReturn struct {
	common.NoOperandsInstruction
}

func (lr *LReturn) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(val)
}

type DReturn struct {
	common.NoOperandsInstruction
}

func (dr *DReturn) Execute(frame *runtime.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(val)
}
