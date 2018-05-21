package arrays

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type AALoad struct {
	common.NoOperandsInstruction
}

func (aALoad *AALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), idx)
	stack.PushRef(refs[idx])
}

type BALoad struct {
	common.NoOperandsInstruction
}

func (bALoad *BALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), idx)
	stack.PushInt(int32(bytes[idx]))
}

type CALoad struct {
	common.NoOperandsInstruction
}

func (cALoad *CALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	chars := arrRef.Bytes()
	checkIndex(len(chars), idx)
	stack.PushInt(int32(chars[idx]))
}

type DALoad struct {
	common.NoOperandsInstruction
}

func (dALoad *DALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), idx)
	stack.PushDouble(doubles[idx])
}

type FALoad struct {
	common.NoOperandsInstruction
}

func (fALoad *FALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), idx)
	stack.PushFloat(floats[idx])
}

type IALoad struct {
	common.NoOperandsInstruction
}

func (iALoad *IALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), idx)
	stack.PushInt(ints[idx])
}

type LALoad struct {
	common.NoOperandsInstruction
}

func (lALoad *LALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), idx)
	stack.PushLong(longs[idx])
}

type SALoad struct {
	common.NoOperandsInstruction
}

func (sALoad *SALoad) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), idx)
	stack.PushInt(int32(shorts[idx]))
}
