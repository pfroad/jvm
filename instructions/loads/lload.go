package loads

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// load long type var from localVars and push to operand stack
type LLoad struct {
	common.Index8Instruction
}

func lLoad(frame *runtime.Frame, index uint) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(index))
}

func (ll *LLoad) Execute(frame *runtime.Frame) {
	lLoad(frame, ll.Index)
}

type LLoad0 struct {
	common.NoOperandsInstruction
}

func(ll0 *LLoad0) Execute(frame *runtime.Frame) {
	lLoad(frame, 0)
}

type LLoad1 struct {
	common.NoOperandsInstruction
}

func(ll1 *LLoad1) Execute(frame *runtime.Frame) {
	lLoad(frame, 1)
}

type LLoad2 struct {
	common.NoOperandsInstruction
}

func(ll2 *LLoad2) Execute(frame *runtime.Frame) {
	lLoad(frame, 2)
}

type LLoad3 struct {
	common.NoOperandsInstruction
}

func(ll3 *LLoad3) Execute(frame *runtime.Frame) {
	lLoad(frame, 3)
}