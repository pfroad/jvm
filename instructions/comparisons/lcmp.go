package comparisons

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// compare long
type LCMP struct {
	common.NoOperandsInstruction
}

func (lcmp *LCMP) Execute(frame *runtime.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
