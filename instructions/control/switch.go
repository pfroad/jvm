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

func skipPadding(reader *runtime.ByteCodeReader) {
	for reader.PC()%4 != 0 {
		reader.ReadUint8()
	}
}

func (ts *TableSwitch) FetchOperands(reader *runtime.ByteCodeReader) {
	skipPadding(reader)
	ts.defaultOffset = reader.ReadInt32()
	ts.low = reader.ReadInt32()
	ts.high = reader.ReadInt32()
	jumpOffsetsCount := ts.high - ts.low + 1
	ts.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ts *TableSwitch) Execute(frame *runtime.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int32
	if index >= ts.low && index <= ts.high {
		offset = ts.jumpOffsets[index]
	} else {
		offset = ts.defaultOffset
	}
	ts.BranchByOffset(frame, offset)
}

type LookupSwitch struct {
	defaultOffset    int32
	npairs           int32
	matchOffsetPairs []int32	// key, offset, key, offset
	common.BranchInstruction
}

func (ls *LookupSwitch) FetchOperands(reader *runtime.ByteCodeReader) {
	skipPadding(reader)
	ls.defaultOffset = reader.ReadInt32()
	ls.npairs = reader.ReadInt32()
	ls.matchOffsetPairs = reader.ReadInt32s(ls.npairs)
}

func (ls *LookupSwitch) Execute(frame *runtime.Frame) {
	key := frame.OperandStack().PopInt()

	len := ls.npairs * 2
	for i := int32(0); i < len; i += 2 {
		if key == ls.matchOffsetPairs[i] {
			ls.BranchByOffset(frame, ls.matchOffsetPairs[i + 1])
			return
		}
	}
	ls.BranchByOffset(frame, ls.defaultOffset)
}