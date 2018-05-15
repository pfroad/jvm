package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type GetField struct {
	common.Index16Instruction
}

func (g *GetField) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConst(g.Index).(*data.FieldRef)
	field := fieldRef.ResolveField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	fieldId := field.FieldId()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': // boolean, byte, char, short, int
		stack.PushInt(ref.Fields().GetInt(fieldId))
	case 'J': // long
		stack.PushLong(ref.Fields().GetLong(fieldId))
	case 'F':
		stack.PushFloat(ref.Fields().GetFloat(fieldId))
	case 'D':
		stack.PushDouble(ref.Fields().GetDouble(fieldId))
	case 'L', '[':
		stack.PushRef(ref.Fields().GetRef(fieldId))
	default:
		// todo
	}
}
