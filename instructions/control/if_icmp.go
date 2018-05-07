package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IFICMPEQ struct {
	common.BranchInstruction
}

type IFICMPNE struct {
	common.BranchInstruction
}

type IFICMPLT struct {
	common.BranchInstruction
}

type IFICMPGT struct {
	common.BranchInstruction
}

type IFICMPLE struct {
	common.BranchInstruction
}

type IFICMPGE struct {
	common.BranchInstruction
}

func (eq *IFICMPEQ) Execute(frame *runtime.Frame) {
	if condition(frame, EQ) {
		eq.BranchInstruction.Branch(frame)
	} else {
		eq.NoBranch(frame)
	}
}

func (ne *IFICMPNE) Execute(frame *runtime.Frame) {
	if condition(frame, NE) {
		ne.BranchInstruction.Branch(frame)
	} else {
		ne.NoBranch(frame)
	}
}

func (lt *IFICMPLT) Execute(frame *runtime.Frame) {
	if condition(frame, LT) {
		lt.BranchInstruction.Branch(frame)
	} else {
		lt.NoBranch(frame)
	}
}

func (gt *IFICMPGT) Execute(frame *runtime.Frame) {
	if condition(frame, GT) {
		gt.BranchInstruction.Branch(frame)
	} else {
		gt.NoBranch(frame)
	}
}

func (le *IFICMPLE) Execute(frame *runtime.Frame) {
	if condition(frame, LE) {
		le.BranchInstruction.Branch(frame)
	} else {
		le.NoBranch(frame)
	}
}

func (ge *IFICMPGE) Execute(frame *runtime.Frame) {
	if condition(frame, GE) {
		ge.BranchInstruction.Branch(frame)
	} else {
		ge.NoBranch(frame)
	}
}
