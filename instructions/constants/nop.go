package constants

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// operand code is 0X00
type NOP struct {
	common.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *runtime.Frame) {
	// nothing to do
}
