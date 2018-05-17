package common

import "jvm/runtime"

type Instruction interface {
	FetchOperands(reader *runtime.ByteCodeReader)
	Execute(frame *runtime.Frame)
}

type NoOperandsInstruction struct {
}

func (noop *NoOperandsInstruction) FetchOperands(reader *runtime.ByteCodeReader) {
	// nothing to do
	//frame.SetPC(reader.pc)
}

// offset is uint 16
type BranchInstruction struct {
	Offset    int
	currentPC int
}

func (bi *BranchInstruction) FetchOperands(reader *runtime.ByteCodeReader) {
	bi.Offset = int(reader.ReadInt16())
	bi.currentPC = reader.PC()
}

//func (bi *BranchInstruction) Execute(frame *runtime.Frame) {
//	// goto
//}

func (bi *BranchInstruction) Branch(frame *runtime.Frame) {
	cPC := frame.Thread().PC()
	frame.SetNextPC(cPC + bi.Offset)
}

func (bi *BranchInstruction) BranchByOffset(frame *runtime.Frame, offset int32) {
	cPC := frame.Thread().PC()
	frame.SetNextPC(cPC + int(offset))
}

func (bi *BranchInstruction) NoBranch(frame *runtime.Frame) {
	frame.SetNextPC(bi.currentPC)
}

// localVars index is uint8
type Index8Instruction struct {
	Index uint
}

func (i8 *Index8Instruction) FetchOperands(reader *runtime.ByteCodeReader) {
	i8.Index = uint(reader.ReadInt8())
	//frame.SetPC(reader.PC())
}

type Index16Instruction struct {
	Index uint
}

func (i16 *Index16Instruction) FetchOperands(reader *runtime.ByteCodeReader) {
	i16.Index = uint(reader.ReadInt16())
	//frame.SetPC(reader.PC())
}
