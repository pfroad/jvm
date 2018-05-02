package comparisons

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// compare float
type FCMPG struct {
	common.NoOperandsInstruction
}

type FCMPL struct {
	common.NoOperandsInstruction
}

func fcomp(frame runtime.Frame, gFlag bool) {
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

func (fcmpg *FCMPG) Execute(frame runtime.Frame)  {
	fcomp(frame, true)
}

func (fcmpl *FCMPL) Execute(frame runtime.Frame) {
	fcomp(frame, false)
}