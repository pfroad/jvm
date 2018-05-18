package references

import (
	"jvm/runtime"
	"jvm/runtime/data"
)

// Invoke interface method
type InvokeInterface struct {
	Index uint
	count uint8
	zero  uint8
}

func (i *InvokeInterface) FetchOperands(reader *runtime.ByteCodeReader) {
	i.Index = uint(reader.ReadInt16())
	i.count = reader.ReadUint8() // same as argCount
	i.zero = reader.ReadUint8()  // must be zero, in some oracle jvm
}

func (i *InvokeInterface) Execute(frame *runtime.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	interfaceMethodRef := cp.GetConst(i.Index).(*data.InterfaceMethodRef)
	method := interfaceMethodRef.ResolveInterfaceMethod()

	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	objectref := frame.OperandStack().GetTopRef(method.ArgCount() - 1)
	if objectref == nil {
		panic("java.lang.NullPointerException")
	}

	if !objectref.IsImplements(method.Class()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	invokeMethod := data.LookupMethodInClass(objectref.Class(), method.Name(), method.Descriptor())
	if invokeMethod == nil || invokeMethod.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !invokeMethod.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	InvokeMethod(invokeMethod, frame)
}
