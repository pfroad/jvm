package references

import (
	"fmt"
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	common.Index16Instruction
}

// hack
func (i *InvokeVirtual) Execute(frame *runtime.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConst(i.Index).(*data.MethodRef)
	method := methodRef.ResolveMethod()

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetTopRef(method.ArgCount() - 1)
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
	}

	resolvedClass := method.Class()
	c := ref.Class()
	if method.IsProtected() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		currentClass.GetPackageName() != resolvedClass.GetPackageName() &&
		!(c == currentClass || c.IsSubClassOf(currentClass)) {
		panic("java.lang.IllegalAccessError")
	}

	invokeMethod := data.LookupMethodInClass(c, method.Name(), method.Descriptor())
	if invokeMethod == nil || invokeMethod.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	InvokeMethod(invokeMethod, frame)
}

func _println(stack *data.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
