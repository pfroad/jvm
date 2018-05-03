package control

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
	common.BranchInstruction
}

func (ts *TableSwitch) skipPadding(reader *common.BytecodeReader) {
	for reader.PC()%4 != 0 {
		reader.ReadUint8()
	}
}

func (ts *TableSwitch) FetchOperands(reader *common.BytecodeReader) {
	ts.skipPadding(reader)
	ts.defaultOffset = reader.ReadInt32()
	ts.low = reader.ReadInt32()
	ts.high = reader.ReadInt32()
	jumpOffsetsCount := ts.high - ts.low + 1
	ts.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ts *TableSwitch) Execute(frame runtime.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int32
	if index >= ts.low && index <= ts.high {
		offset = ts.jumpOffsets[index]
	} else {
		offset = ts.defaultOffset
	}
	ts.BranchByOffset(frame, offset)
}