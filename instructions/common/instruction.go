package common

import "jvm/runtime"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtime.Frame)
}

type NoOperandsInstruction struct {
}

func (noop *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// offset is uint 16
type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

func (bi *BranchInstruction) Branch(frame runtime.Frame) {
	currentPC := frame.PC()
	frame.SetPC(currentPC + bi.Offset)
}

func (bi *BranchInstruction) BranchByOffset(frame runtime.Frame, offset int32) {
	currentPC := frame.PC()
	frame.SetPC(currentPC + int(offset))
}

// localVars index is uint8
type Index8Instruction struct {
	Index uint
}

func (i8 *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8.Index = uint(reader.ReadInt8())
}

type Index16Instruction struct {
	Index uint
}

func (i16 *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i16.Index = uint(reader.ReadInt16())
}

