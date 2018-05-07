package heap

import (
	"jvm/classfile"
	"jvm/runtime"
)

type Class struct {
	accessFlags   AccessFlags
	className     string
	cp            *ConstantPool
	fields        []*Field
	methods       []*Method
	classLoader   *ClassLoader
	superClass    *Class
	interfaces    []*Class
	instanceCount uint
	staticCount   uint
	staticVars    *runtime.Variables
}

func NewClass(cf classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = AccessFlags{cf.AccessFlags()}
	class.className = cf.ClassName()
	class.cp = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(cf.Fields())
	class.methods = newMethods(cf.Methods())
	return class
}
