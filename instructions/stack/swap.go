package stack

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type Swap struct {
	common.NoOperandsInstruction
}

func (swap *Swap) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	stack.Push(v1)
	stack.Push(v2)
}
