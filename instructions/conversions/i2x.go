package conversions

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

type I2F struct {
	common.NoOperandsInstruction
}

type I2L struct {
	common.NoOperandsInstruction
}

type I2D struct {
	common.NoOperandsInstruction
}

// int to byte
type I2B struct {
	common.NoOperandsInstruction
}

// int to char
type I2C struct {
	common.NoOperandsInstruction
}

// int to short
type I2S struct {
	common.NoOperandsInstruction
}

func (i2l *I2L) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushLong(int64(val))
}

func (i2f *I2F) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushFloat(float32(val))
}

func (i2d *I2D) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushDouble(float64(val))
}

func (i2l *I2B) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int8(val)))
}

func (i2c *I2C) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(uint16(val)))
}

func (i2s *I2S) Execute(frame runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int16(val)))
}