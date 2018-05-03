package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type GOTO struct {
	common.BranchInstruction
}

func (g2 *GOTO) Execute(frame runtime.Frame) {
	g2.BranchInstruction.Branch(frame)
}