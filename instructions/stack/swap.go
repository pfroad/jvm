package stack
<<<<<<< HEAD

import (
	"ch02/instructions/common"
	"ch02/runtime"
)

type Swap struct {
	common.NoOperandsInstruction
}

func (swap *Swap) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	stack.Push(v1)
	stack.Push(v2)
}
=======
>>>>>>> ddd131059547a2782c3e8720eefd44c58f607b51
