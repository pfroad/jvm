package references

import (
	"jvm/instructions/common"
	"jvm/runtime/data"
	"jvm/runtime"
)

/*
Invoke instance method; special handling for superclass, private,
and instance initialization method invocations

Operand Stack	..., objectref, [arg1, [arg2 ...]] → ...
*/
type InvokeSpecial struct {
	common.Index16Instruction
}

/*
 	1. current class is abstract, interface or real class
	2. objectref is object for a real class
	3. method.Class() is class that declare the method
*/
func (i *InvokeSpecial) Execute(frame *runtime.Frame) {
	//frame.OperandStack().PopRef()
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConst(i.Index).(*data.MethodRef)
	method := methodRef.ResolveMethod()
	resolvedClass := method.Class()

	if method.Name() == "<init>" && currentClass != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	
	ref := frame.OperandStack().GetTopRef(method.ArgCount() - 1)	// not static, except "this"
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if method.IsProtected() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		currentClass.GetPackageName() != resolvedClass.GetPackageName() &&
		!(ref.Class() == currentClass || ref.Class().IsSubClassOf(currentClass)) {
			panic("java.lang.IllegalAccessError")
	}

	invokeMethod := method
	// if resolved class is super class of current class and ACC_SUPER flag is set, method.Class() should be
	// direct super class of current class
	// ACC_SUPER means method has been override
	if method.Name() != "<init>" &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		currentClass.IsSuper() {
		invokeMethod = data.LookupMethodInClass(currentClass.SuperClass(), method.Name(), method.Descriptor())
	}

	InvokeMethod(invokeMethod, frame)
}