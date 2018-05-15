package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

type PutField struct {
	common.Index16Instruction
}

func (p *PutField) Execute(frame *runtime.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConst(p.Index).(*data.FieldRef)
	field := fieldRef.ResolveField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := field.Class()
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	stack := frame.OperandStack()
	fieldId := field.FieldId()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': // boolean, byte, char, short, int
		val := stack.PopInt()
		ref := stack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}

		ref.Fields().SetInt(fieldId, val)
		//stack.PushInt(staticVars.GetInt(fieldId))
	case 'J': // long
		//stack.PushLong(staticVars.GetLong(fieldId))
		val := stack.PopLong()
		ref := stack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}

		ref.Fields().SetLong(fieldId, val)
	case 'F':
		//stack.PushFloat(staticVars.GetFloat(fieldId))
		val := stack.PopFloat()
		ref := stack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}

		ref.Fields().SetFloat(fieldId, val)
	case 'D':
		//stack.PushDouble(staticVars.GetDouble(fieldId))
		val := stack.PopDouble()
		ref := stack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}

		ref.Fields().SetDouble(fieldId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}

		ref.Fields().SetRef(fieldId, val)
	default:
		// todo
	}
}
