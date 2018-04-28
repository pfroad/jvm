package constants

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type AConstNull struct {
	common.NoOperandsInstruction
}

func (acn *AConstNull) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DConst0 struct {
	common.NoOperandsInstruction
}

func (dc0 *DConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DConst1 struct {
	common.NoOperandsInstruction
}

func (dc1 *DConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FConst0 struct {
	common.NoOperandsInstruction
}

func (fc0 *FConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FConst1 struct {
	common.NoOperandsInstruction
}

func (fc1 *FConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FConst2 struct {
	common.NoOperandsInstruction
}

func (fc2 *FConst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type IConstM1 struct {
	common.NoOperandsInstruction
}

func (icm1 *IConstM1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

type IConst0 struct {
	common.NoOperandsInstruction
}

func (ic0 *IConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

type IConst1 struct {
	common.NoOperandsInstruction
}

func (ic1 *IConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

type IConst2 struct {
	common.NoOperandsInstruction
}

func (ic2 *IConst2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

type IConst3 struct {
	common.NoOperandsInstruction
}

func (ic3 *IConst3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

type IConst4 struct {
	common.NoOperandsInstruction
}

func (ic4 *IConst4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

type IConst5 struct {
	common.NoOperandsInstruction
}

func (ic5 *IConst5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

type LConst0 struct {
	common.NoOperandsInstruction
}

func (lc0 *LConst0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0)
}

type LConst1 struct {
	common.NoOperandsInstruction
}

func (lc1 *LConst1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1)
}