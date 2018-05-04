package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IFACMPEQ struct {
	common.BranchInstruction
}

type IFACMPNE struct {
	common.BranchInstruction
}

func (eq *IFACMPEQ) Execute(frame *runtime.Frame)  {
	if aCondition(frame, EQ) {
		eq.BranchInstruction.Branch(frame)
	}
}

func (ne *IFACMPNE) Execute(frame *runtime.Frame)  {
	if aCondition(frame, NE) {
		ne.BranchInstruction.Branch(frame)
	}
}