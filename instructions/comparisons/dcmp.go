package comparisons

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// compare double
type DCMPG struct {
	common.NoOperandsInstruction
}

type DCMPL struct {
	common.NoOperandsInstruction
}

func dcomp(frame *runtime.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (dcmpg *DCMPG) Execute(frame *runtime.Frame)  {
	dcomp(frame, true)
}

func (dcmpl *DCMPL) Execute(frame *runtime.Frame) {
	dcomp(frame, false)
}