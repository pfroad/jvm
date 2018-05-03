package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IFEQ struct {
	common.BranchInstruction
}

type IFNE struct {
	common.BranchInstruction
}

type IFLT struct {
	common.BranchInstruction
}

type IFGT struct {
	common.BranchInstruction
}

type IFLE struct {
	common.BranchInstruction
}

type IFGE struct {
	common.BranchInstruction
}

func (ifeq *IFEQ) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, EQ) {
		ifeq.BranchInstruction.Branch(frame)
	}
}

func (ifne *IFNE) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, NE) {
		ifne.BranchInstruction.Branch(frame)
	}
}

func (iflt *IFLT) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, LT) {
		iflt.BranchInstruction.Branch(frame)
	}
}

func (ifgt *IFGT) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, GT) {
		ifgt.BranchInstruction.Branch(frame)
	}
}

func (ifle *IFLE) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, LE) {
		ifle.BranchInstruction.Branch(frame)
	}
}

func (ifge *IFGE) Execute(frame runtime.Frame) {
	val := frame.OperandStack().PopInt()

	if compare(val, 0, GE) {
		ifge.BranchInstruction.Branch(frame)
	}
}