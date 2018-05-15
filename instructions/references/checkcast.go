package references

import (
	"jvm/instructions/common"
	"jvm/runtime"
	"jvm/runtime/data"
)

/*
If objectref is null, then the operand stack is unchanged.

Otherwise, the named class, array, or interface type is resolved
(§5.4.3.1). If objectref can be cast to the resolved class, array,
or interface type, the operand stack is unchanged; otherwise, the
checkcast instruction throws a ClassCastException

Java code:
if (xxx instanceof ClassYYY) {
	yyy = (ClassYYY) xxx;
	// use yyy
}

*/
type CheckCast struct {
	common.Index16Instruction
}


func (c *CheckCast) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConst(c.Index).(*data.ClassRef)
	class := classRef.ResolveClass()

	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
