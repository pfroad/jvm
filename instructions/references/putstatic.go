package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type PutStatic struct {
	common.Index16Instruction
}

func (ps *PutStatic) Execute(frame *runtime.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConst(ps.Index).(*data.FieldRef)
	field := fieldRef.ResolveField()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := field.Class()
	if field.IsFinal() {
		if currentClass != class || currentMethod.ClassMember.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	staticVars := class.StaticVars()
	fieldId := field.FieldId()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': // boolean, byte, char, short, int
		staticVars.SetInt(fieldId, stack.PopInt())
	case 'J': // long
		staticVars.SetLong(fieldId, stack.PopLong())
	case 'F':
		staticVars.SetFloat(fieldId, stack.PopFloat())
	case 'D':
		staticVars.SetDouble(fieldId, stack.PopDouble())
	case 'L', '[':
		staticVars.SetRef(fieldId, stack.PopRef())
	default:
		// todo
	}
}
