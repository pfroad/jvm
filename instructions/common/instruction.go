package common

import "jvm/runtime"

type Instruction interface {
	FetchOperands(reader *BytecodeReader, frame *runtime.Frame)
	Execute(frame *runtime.Frame)
}

type NoOperandsInstruction struct {
}

func (noop *NoOperandsInstruction) FetchOperands(reader *BytecodeReader, frame *runtime.Frame) {
	// nothing to do
	frame.SetPC(reader.pc)
}

// offset is uint 16
type BranchInstruction struct {
	Offset    int
	currentPC int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader, frame *runtime.Frame) {
	bi.Offset = int(reader.ReadInt16())
	bi.currentPC = reader.PC()
}

//func (bi *BranchInstruction) Execute(frame *runtime.Frame) {
//	// goto
//}

func (bi *BranchInstruction) Branch(frame *runtime.Frame) {
	cPC := frame.Thread().PC()
	frame.SetPC(cPC + bi.Offset)
}

func (bi *BranchInstruction) BranchByOffset(frame *runtime.Frame, offset int32) {
	cPC := frame.Thread().PC()
	frame.SetPC(cPC + int(offset))
}

func (bi *BranchInstruction) NoBranch(frame *runtime.Frame) {
	frame.Thread().SetPC(bi.currentPC)
}

// localVars index is uint8
type Index8Instruction struct {
	Index uint
}

func (i8 *Index8Instruction) FetchOperands(reader *BytecodeReader, frame *runtime.Frame) {
	i8.Index = uint(reader.ReadInt8())
	frame.SetPC(reader.PC())
}

type Index16Instruction struct {
	Index uint
}

func (i16 *Index16Instruction) FetchOperands(reader *BytecodeReader, frame *runtime.Frame) {
	i16.Index = uint(reader.ReadInt16())
	frame.SetPC(reader.PC())
}

