package math

impXORt (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IXOR struct {
	common.NoOperandsInstruction
}

type LXOR struct {
	common.NoOperandsInstruction
}

func (xor *IXOR) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 ^ v2)
}

func (xor *LXOR) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 ^ v2)
}
