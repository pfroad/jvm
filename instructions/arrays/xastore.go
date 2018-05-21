package arrays

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type AAStore struct {
	common.NoOperandsInstruction
}

func (aAstore *AAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), idx)
	refs[idx] = val
}

type BAStore struct {
	common.NoOperandsInstruction
}

func (bAstore *BAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), idx)
	bytes[idx] = int8(val)
}

type CAStore struct {
	common.NoOperandsInstruction
}

func (cAstore *CAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), idx)
	chars[idx] = int16(val)
}

type DAStore struct {
	common.NoOperandsInstruction
}

func (dAstore *DAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), idx)
	doubles[idx] = val
}

type FAStore struct {
	common.NoOperandsInstruction
}

func (fAstore *FAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), idx)
	floats[idx] = val
}

type IAStore struct {
	common.NoOperandsInstruction
}

func (iAstore *IAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), idx)
	ints[idx] = val
}

type LAStore struct {
	common.NoOperandsInstruction
}

func (lAstore *LAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), idx)
	longs[idx] = val
}

type SAStore struct {
	common.NoOperandsInstruction
}

func (sAstore *SAStore) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNull(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), idx)
	shorts[idx] = int16(val)
}
