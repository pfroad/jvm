package stores

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type AStore struct {
	common.Index8Instruction
}

func aStore(frame *runtime.Frame, index uint) {
	frame.LocalVars().SetRef(index, frame.OperandStack().PopRef())
}

func (as *AStore) Execute(frame *runtime.Frame) {
	aStore(frame, as.Index)
}

type AStore0 struct {
	common.NoOperandsInstruction
}

func (as0 *AStore0) Execute(frame *runtime.Frame) {
	aStore(frame, 0)
}

type AStore1 struct {
	common.NoOperandsInstruction
}

func (as1 *AStore1) Execute(frame *runtime.Frame) {
	aStore(frame, 1)
}

type AStore2 struct {
	common.NoOperandsInstruction
}

func (as2 *AStore2) Execute(frame *runtime.Frame) {
	aStore(frame, 2)
}

type AStore3 struct {
	common.NoOperandsInstruction
}

func (as3 *AStore3) Execute(frame *runtime.Frame) {
	aStore(frame, 3)
}