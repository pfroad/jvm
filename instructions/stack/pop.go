package stack

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// pop int, float, ref variable
// pop = 87 (0X57)
type Pop struct {
	common.NoOperandsInstruction
}

func (pop *Pop) Execute(frame *runtime.Frame) {
	frame.OperandStack().Pop()
}

// pop2 = 88 (0X57)
type Pop2 struct {
	common.NoOperandsInstruction
}

func (Pop2 *Pop2) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	stack.Pop()
	stack.Pop()
}