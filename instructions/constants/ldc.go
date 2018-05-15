package constants

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

func _ldc(frame *runtime.Frame, index uint) {
	cp := frame.Method().Class().ConstantPool()
	val := cp.GetConst(index)

	switch val.(type) {
	case int32:
		frame.OperandStack().PushInt(val.(int32))
	case float32:
		frame.OperandStack().PushFloat(val.(float32))
	case string:
		// implement ch8
	case *data.ClassRef:
		// implement ch9
	default:
		panic("to do: ldc!")
	}
}

type LDC struct {
	common.Index8Instruction
}

func (l *LDC) Execute(frame *runtime.Frame) {
	_ldc(frame, l.Index)
}

type LDCW struct {
	common.Index16Instruction
}

func (l *LDCW) Execute(frame *runtime.Frame) {
	_ldc(frame, l.Index)
}

type LDC2W struct {
	common.Index16Instruction
}

func (l *LDC2W) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	val := cp.GetConst(l.Index)

	switch val.(type) {
	case int64:
		frame.OperandStack().PushLong(val.(int64))
	case float64:
		frame.OperandStack().PushDouble(val.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
