package loads

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// load reference var from localVars and push to operand stack
type ALoad struct {
	common.Index8Instruction
}

func aLoad(frame *runtime.Frame, index uint) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(index))
}

func (al *ALoad) Execute(frame *runtime.Frame) {
	aLoad(frame, al.Index)
}

type ALoad0 struct {
	common.NoOperandsInstruction
}

func (al0 *ALoad0) Execute(frame *runtime.Frame) {
	aLoad(frame, 0)
}

type ALoad1 struct {
	common.NoOperandsInstruction
}

func (al1 *ALoad1) Execute(frame *runtime.Frame) {
	aLoad(frame, 1)
}

type ALoad2 struct {
	common.NoOperandsInstruction
}

func (al2 *ALoad2) Execute(frame *runtime.Frame) {
	aLoad(frame, 2)
}

type ALoad3 struct {
	common.NoOperandsInstruction
}

func (al3 *ALoad3) Execute(frame *runtime.Frame) {
	aLoad(frame, 3)
}