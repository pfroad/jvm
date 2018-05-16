package data

import (
	"math"
)

type OperandStack struct {
	size uint
	arr  []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{arr: make([]Slot, maxStack)}
	}
	return nil
}

func (oStack *OperandStack) Pop() Slot {
	oStack.size--
	return oStack.arr[oStack.size]
}

func (oStack *OperandStack) Push(val Slot) {
	oStack.arr[oStack.size] = val
	oStack.size++
}

func (oStack *OperandStack) PopInt() int32 {
	oStack.size--
	return oStack.arr[oStack.size].(int32)
}

func (oStack *OperandStack) PushInt(val int32) {
	oStack.arr[oStack.size] = val
	oStack.size++
}

func (oStack *OperandStack) PopFloat() float32 {
	oStack.size--
	return oStack.arr[oStack.size].(float32)
}

func (oStack *OperandStack) PushFloat(val float32) {
	oStack.arr[oStack.size] = val
	oStack.size++
}

func (oStack *OperandStack) PopLong() int64 {
	oStack.size--
	high := uint32(oStack.arr[oStack.size].(int32))
	oStack.size--
	low := uint32(oStack.arr[oStack.size].(int32))

	return int64(high)<<32 | int64(low)
}

func (oStack *OperandStack) PushLong(val int64) {
	oStack.arr[oStack.size] = int32(val)
	oStack.size++
	oStack.arr[oStack.size] = int32(val >> 32)
	oStack.size++
}

func (oStack *OperandStack) PopDouble() float64 {
	bits := oStack.PopLong()
	return math.Float64frombits(uint64(bits))
}

func (oStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	oStack.PushLong(int64(bits))
}

func (oStack *OperandStack) PushRef(obj *Object) {
	oStack.arr[oStack.size] = obj // Pop and release reference. int and float is not reference type, needn't release
	oStack.size++
}

func (oStack *OperandStack) PopRef() *Object {
	oStack.size--
	ref := oStack.arr[oStack.size].(*Object)
	oStack.arr[oStack.size] = nil // Pop and release reference. int and float is not reference type, needn't release
	return ref
}

func (oStack *OperandStack) GetTopRef(n uint) *Object {
	return oStack.arr[oStack.size - n - 1].(*Object)
}
