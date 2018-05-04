package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IShR struct {
	common.NoOperandsInstruction
}

type IShL struct {
	common.NoOperandsInstruction
}

type IUShR struct {
	common.NoOperandsInstruction
}

type LShR struct {
	common.NoOperandsInstruction
}

type LShL struct {
	common.NoOperandsInstruction
}

type LUShR struct {
	common.NoOperandsInstruction
}

func (sh *IShR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 5 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x1f
	val := stack.PopInt()
	stack.PushInt(val >> s)
}

func (sh *IShL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 5 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x1f
	val := stack.PopInt()
	stack.PushInt(val << s)
}

func (sh *IUShR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 7 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x1f
	val := uint32(stack.PopInt())
	stack.PushInt(int32(val >> s))
}

func (sh *LShR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 5 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x3f
	val := stack.PopLong()
	stack.PushLong(val >> s)
}

func (sh *LShL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 7 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x3f
	val := stack.PopLong()
	stack.PushLong(val << s)
}

func (sh *LUShR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	// int is 32 bits, 7 bits enough to shift it
	s := uint32(stack.PopInt()) & 0x3f
	val := uint64(stack.PopLong())
	stack.PushLong(int64(val >> s))
}
