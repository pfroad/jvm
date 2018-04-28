package loads

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// load Int type var from localVars and push to operand stack
type ILoad struct {
	common.Index8Instruction
}

func iLoad(frame *runtime.Frame, index uint) {
	frame.OperandStack().PushInt(frame.LocalVars().GetInt(index))
}

func (il *ILoad) Execute(frame *runtime.Frame) {
	iLoad(frame, il.Index)
}

type ILoad0 struct {
	common.NoOperandsInstruction
}

func(il0 *ILoad0) Execute(frame *runtime.Frame) {
	iLoad(frame, 0)
}

type ILoad1 struct {
	common.NoOperandsInstruction
}

func(il1 *ILoad1) Execute(frame *runtime.Frame) {
	iLoad(frame, 1)
}

type ILoad2 struct {
	common.NoOperandsInstruction
}

func(il2 *ILoad2) Execute(frame *runtime.Frame) {
	iLoad(frame, 2)
}

type ILoad3 struct {
	common.NoOperandsInstruction
}

func(il3 *ILoad3) Execute(frame *runtime.Frame) {
	iLoad(frame, 3)
}