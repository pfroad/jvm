package stores

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type LStore struct {
	common.Index8Instruction
}

func lStore(frame *runtime.Frame, index uint) {
	frame.LocalVars().SetLong(index, frame.OperandStack().PopLong())
}

func (ls *LStore) Execute(frame *runtime.Frame) {
	lStore(frame, ls.Index)
}

type LStore0 struct {
	common.NoOperandsInstruction
}

func (ls0 *LStore0) Execute(frame *runtime.Frame) {
	lStore(frame, 0)
}

type LStore1 struct {
	common.NoOperandsInstruction
}

func (ls1 *LStore1) Execute(frame *runtime.Frame) {
	lStore(frame, 1)
}

type LStore2 struct {
	common.NoOperandsInstruction
}

func (ls2 *LStore2) Execute(frame *runtime.Frame) {
	lStore(frame, 2)
}

type LStore3 struct {
	common.NoOperandsInstruction
}

func (ls3 *LStore3) Execute(frame *runtime.Frame) {
	lStore(frame, 3)
}