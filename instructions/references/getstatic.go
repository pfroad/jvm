package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type GetStatic struct {
	common.Index16Instruction
}

func (getStatic *GetStatic) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	ref := cp.GetConst(getStatic.Index).(*data.FieldRef)
	field := ref.ResolveField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertPC()
		InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	fieldId := field.FieldId()
	staticVars := class.StaticVars()
	descriptor := field.Descriptor()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': // boolean, byte, char, short, int
		stack.PushInt(staticVars.GetInt(fieldId))
	case 'J': // long
		stack.PushLong(staticVars.GetLong(fieldId))
	case 'F':
		stack.PushFloat(staticVars.GetFloat(fieldId))
	case 'D':
		stack.PushDouble(staticVars.GetDouble(fieldId))
	case 'L', '[':
		stack.PushRef(staticVars.GetRef(fieldId))
	default:
		// todo
	}
}
