package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type GoTo struct {
	common.BranchInstruction
}

func (g2 *GoTo) Execute(frame *runtime.Frame) {
	g2.BranchInstruction.Branch(frame)
}

type GoToW struct {
	common.BranchInstruction
}

func (g2w *GoToW) FetchOperands(reader *runtime.ByteCodeReader) {
	g2w.Offset = int(reader.ReadInt32())
}

func (g2w *GoToW) Execute(frame *runtime.Frame) {
	g2w.BranchInstruction.Branch(frame)
}