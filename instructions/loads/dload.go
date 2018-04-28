package loads

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// load double type var from localVars and push to operand stack
type DLoad struct {
	common.Index8Instruction
}

func dLoad(frame *runtime.Frame, index uint) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(index))
}

func (dl *DLoad) Execute(frame *runtime.Frame) {
	dLoad(frame, dl.Index)
}

type DLoad0 struct {
	common.NoOperandsInstruction
}

func(dl0 *DLoad0) Execute(frame *runtime.Frame) {
	dLoad(frame, 0)
}

type DLoad1 struct {
	common.NoOperandsInstruction
}

func(dl1 *DLoad1) Execute(frame *runtime.Frame) {
	dLoad(frame, 1)
}

type DLoad2 struct {
	common.NoOperandsInstruction
}

func(dl2 *DLoad2) Execute(frame *runtime.Frame) {
	dLoad(frame, 2)
}

type DLoad3 struct {
	common.NoOperandsInstruction
}

func(dl3 *DLoad3) Execute(frame *runtime.Frame) {
	dLoad(frame, 3)
}