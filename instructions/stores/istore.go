package stores

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type IStore struct {
	common.Index8Instruction
}

func iStore(frame *runtime.Frame, index uint) {
	frame.LocalVars().SetInt(index, frame.OperandStack().PopInt())
}

func (ls *IStore) Execute(frame *runtime.Frame) {
	iStore(frame, ls.Index)
}

type IStore0 struct {
	common.NoOperandsInstruction
}

func (ls0 *IStore0) Execute(frame *runtime.Frame) {
	iStore(frame, 0)
}

type IStore1 struct {
	common.NoOperandsInstruction
}

func (ls1 *IStore1) Execute(frame *runtime.Frame) {
	iStore(frame, 1)
}

type IStore2 struct {
	common.NoOperandsInstruction
}

func (ls2 *IStore2) Execute(frame *runtime.Frame) {
	iStore(frame, 2)
}

type IStore3 struct {
	common.NoOperandsInstruction
}

func (ls3 *IStore3) Execute(frame *runtime.Frame) {
	iStore(frame, 3)
}