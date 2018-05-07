package constants

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// Push byte constant as int to operand stack
type BIPush struct {
	val int8
} 

func (bip *BIPush) FetchOperands(reader *common.BytecodeReader, frame *runtime.Frame) {
	bip.val = reader.ReadInt8()
	frame.SetPC(reader.PC())
}

func (bip *BIPush) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(int32(bip.val))
}

// push short constant as int to operand stack
type SIPush struct {
	val int16
}

func (sip *SIPush) FetchOperands(reader *common.BytecodeReader, frame *runtime.Frame) {
	sip.val = reader.ReadInt16()
	frame.SetPC(reader.PC())
}

func (sip *SIPush) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(int32(sip.val))
}