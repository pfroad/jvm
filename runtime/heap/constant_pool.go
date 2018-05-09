package heap

import (
	"jvm/classfile"
	"fmt"
)

type Constant interface {
}

type ConstantPool struct {
	class *Class
	consts []Constant
}

func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {
	cp := &ConstantPool{class:class}
	cpCount := len(cfcp)
	consts := make([]Constant, cpCount)

	for i := 0; i < cpCount; i++ {
		cInfo := cfcp[i]
		switch cInfo.(type) {
		case *classfile.ConstantInteger:
			consts[i] = cInfo.(*classfile.ConstantInteger).Value()
		case *classfile.ConstantFloat:
			consts[i] = cInfo.(*classfile.ConstantFloat).Value()
		case *classfile.ConstantLong:
			consts[i] = cInfo.(*classfile.ConstantLong).Value()
			i++
		case *classfile.ConstantDouble:
			consts[i] = cInfo.(*classfile.ConstantDouble).Value()
			i++
		case *classfile.ConstantString:
			consts[i] = cInfo.(*classfile.ConstantString).String()
		case *classfile.ConstantClass:
			consts[i] = newClassRef(cp, cInfo.(*classfile.ConstantClass))
		case *classfile.ConstantFieldRef:
			consts[i] = newFieldRef(cp, cInfo.(*classfile.ConstantFieldRef))
		case *classfile.ConstantMethodRef:
			consts[i] = newMethodRef(cp, cInfo.(*classfile.ConstantMethodRef))
		case *classfile.ConstantInterfaceMethodRef:
			consts[i] = newInterfaceMethodRef(cp, cInfo.(*classfile.ConstantInterfaceMethodRef))
		default:
		}
	}
}

func (cp *ConstantPool) GetConst(index uint) Constant {
	if cst := cp.consts[index]; cst != nil {
		return cst
	}
	panic(fmt.Sprintf("No constant at index %d\n", index))
}
