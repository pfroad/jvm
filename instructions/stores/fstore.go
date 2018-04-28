package stores

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type FStore struct {
	common.Index8Instruction
}

func fStore(frame *runtime.Frame, index uint) {
	frame.LocalVars().SetFloat(index, frame.OperandStack().PopFloat())
}

func (is *FStore) Execute(frame *runtime.Frame) {
	fStore(frame, is.Index)
}

type FStore0 struct {
	common.NoOperandsInstruction
}

func (is0 *FStore0) Execute(frame *runtime.Frame) {
	fStore(frame, 0)
}

type FStore1 struct {
	common.NoOperandsInstruction
}

func (is1 *FStore1) Execute(frame *runtime.Frame) {
	fStore(frame, 1)
}

type FStore2 struct {
	common.NoOperandsInstruction
}

func (is2 *FStore2) Execute(frame *runtime.Frame) {
	fStore(frame, 2)
}

type FStore3 struct {
	common.NoOperandsInstruction
}

func (is3 *FStore3) Execute(frame *runtime.Frame) {
	fStore(frame, 3)
}