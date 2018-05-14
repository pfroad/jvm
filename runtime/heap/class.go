package heap

import (
	"jvm/classfile"
	"jvm/runtime"
	"strings"
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
	staticVars    runtime.Variables
}

func NewClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = AccessFlags{cf.AccessFlags()}
	class.className = cf.ClassName()
	class.cp = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) isAccessTo(other *Class) bool {
	return class.accessFlags.IsPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.className, "/"); i >= 0 {
		return class.className[:i]
	}

	return ""
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.cp
}

func (class *Class) IsInterface() bool {
	return class.accessFlags.IsInterface()
}

func (class *Class) IsAbstract() bool {
	return class.accessFlags.IsAbstract()
}

func (class *Class) NewObject() *runtime.Object {
	return &runtime.Object{class: class, fields: runtime.NewVariables(class.instanceCount)}
}

