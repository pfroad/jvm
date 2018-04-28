package constants

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// aconst_null = 1 (0x01)
type AConstNull struct {
	common.NoOperandsInstruction
}

func (acn *AConstNull) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

// dconst_0 = 14 (0x0e)
type DConst0 struct {
	common.NoOperandsInstruction
}

func (dc0 *DConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// dconst_1 = 15 (0x0f)
type DConst1 struct {
	common.NoOperandsInstruction
}

func (dc1 *DConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// fconst_0 = 11 (0x0b)
type FConst0 struct {
	common.NoOperandsInstruction
}

func (fc0 *FConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// fconst_1 = 12 (0x0c)
type FConst1 struct {
	common.NoOperandsInstruction
}

func (fc1 *FConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// fconst_2 = 13 (0x0d)
type FConst2 struct {
	common.NoOperandsInstruction
}

func (fc2 *FConst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// iconst_m1 = 2 (0x02)
type IConstM1 struct {
	common.NoOperandsInstruction
}

func (icm1 *IConstM1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

// iconst_0 = 3 (0x03)
type IConst0 struct {
	common.NoOperandsInstruction
}

func (ic0 *IConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

// iconst_1 = 4 (0x04)
type IConst1 struct {
	common.NoOperandsInstruction
}

func (ic1 *IConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

// iconst_2 = 5 (0x05)
type IConst2 struct {
	common.NoOperandsInstruction
}

func (ic2 *IConst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

// iconst_3 = 6 (0x06)
type IConst3 struct {
	common.NoOperandsInstruction
}

func (ic3 *IConst3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

// iconst_4 = 7 (0x07)
type IConst4 struct {
	common.NoOperandsInstruction
}

func (ic4 *IConst4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

// iconst_5 = 8 (0x08)
type IConst5 struct {
	common.NoOperandsInstruction
}

func (ic5 *IConst5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

// lconst_0 = 9 (0x09)
type LConst0 struct {
	common.NoOperandsInstruction
}

func (lc0 *LConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0)
}

// lconst_1 = 10 (0x0a)
type LConst1 struct {
	common.NoOperandsInstruction
}

func (lc1 *LConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1)
}