package math

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// i++, iinc 1, 1 or i += 5 is iinc 1, 5
type IINC struct {
	Index uint
	Const int32
}

func (iinc *IINC) FetchOperands(reader *common.BytecodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadInt8())
}

func (iinc *IINC) Execute(frame *runtime.Frame) {
	vars := frame.LocalVars()
	val := vars.GetInt(iinc.Index)
	val += iinc.Const
	vars.SetInt(iinc.Index, val)
}
