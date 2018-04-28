package loads

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// load float type var from localVars and push to operand stack
type FLoad struct {
	common.Index8Instruction
}

func fLoad(frame *runtime.Frame, index uint) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(index))
}

func (fl *FLoad) Execute(frame *runtime.Frame) {
	fLoad(frame, fl.Index)
}

type FLoad0 struct {
	common.NoOperandsInstruction
}

func(fl0 *FLoad0) Execute(frame *runtime.Frame) {
	fLoad(frame, 0)
}

type FLoad1 struct {
	common.NoOperandsInstruction
}

func(fl1 *FLoad1) Execute(frame *runtime.Frame) {
	fLoad(frame, 1)
}

type FLoad2 struct {
	common.NoOperandsInstruction
}

func(fl2 *FLoad2) Execute(frame *runtime.Frame) {
	fLoad(frame, 2)
}

type FLoad3 struct {
	common.NoOperandsInstruction
}

func(fl3 *FLoad3) Execute(frame *runtime.Frame) {
	fLoad(frame, 3)
}