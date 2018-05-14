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


}
