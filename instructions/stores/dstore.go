package stores

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type DStore struct {
	common.Index8Instruction
}

func dStore(frame *runtime.Frame, index uint) {
	frame.LocalVars().SetDouble(index, frame.OperandStack().PopDouble())
}

func (ds *DStore) Execute(frame *runtime.Frame) {
	dStore(frame, ds.Index)
}

type DStore0 struct {
	common.NoOperandsInstruction
}

func (ds0 *DStore0) Execute(frame *runtime.Frame) {
	dStore(frame, 0)
}

type DStore1 struct {
	common.NoOperandsInstruction
}

func (ds1 *DStore1) Execute(frame *runtime.Frame) {
	dStore(frame, 1)
}

type DStore2 struct {
	common.NoOperandsInstruction
}

func (ds2 *DStore2) Execute(frame *runtime.Frame) {
	dStore(frame, 2)
}

type DStore3 struct {
	common.NoOperandsInstruction
}

func (ds3 *DStore3) Execute(frame *runtime.Frame) {
	dStore(frame, 3)
}
