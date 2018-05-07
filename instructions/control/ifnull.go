package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IfNull struct {
	common.BranchInstruction
}

func (ifNull *IfNull) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()

	if ref == nil {
		ifNull.Branch(frame)
	} else {
		ifNull.NoBranch(frame)
	}
}

type IfNonNull struct {
	common.BranchInstruction
}

func (nonNull *IfNonNull) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()

	if ref != nil {
		nonNull.Branch(frame)
	} else {
		nonNull.NoBranch(frame)
	}
}