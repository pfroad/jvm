package data

import (
	"jvm/classfile"
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
	staticVars    Variables
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

func (class *Class) NewObject() *Object {
	return &Object{class: class, fields: NewVariables(class.instanceCount)}
}

func (class *Class) StaticVars() Variables {
	return class.staticVars
}

func (class *Class) isAssignableFrom(other *Class) bool {
	if other == class {
		return true
	}

	if !class.IsInterface() {
		//if superClass := other.superClass; superClass != nil {
		//	return class.isAssignableFrom(superClass)
		//}
		if other.isExtendClass(class) {
			return true
		}
	} else {
		//for c := other; c != nil; c = c.superClass {
		//	for _, iface := range c.interfaces {
		//		return class.isAssignableFrom(iface)
		//	}
		//}
		if other.isImplements(class) {
			return true
		}
	}

	return false
}

func (class *Class) isExtendClass(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}

	return false
}

func (class *Class) isExtendInterface(other *Class) bool {
	for _, iface := range class.interfaces {
		if iface == other || iface.isExtendInterface(other) {
			return true
		}
	}

	return false
}

func (class *Class) isImplements(other *Class) bool {
	for c := class ; c != nil; c = c.superClass {
		for _, iface := range c.interfaces {
			if iface == other || iface.isExtendInterface(other) {
				return true
			}
		}
	}

	return false
}

func (class *Class) GetMainMethod() *Method {
	method := class.getStaticMethod("main", "([Ljava/lang/String;)V")

	if method != nil && method.IsPublic() {
		return method
	}

	return nil
}

func (class *Class) getStaticMethod(methodName string, descriptor string) *Method {
	for _, method := range class.methods {
		if method.accessFlags.IsStatic() &&
			method.Name() == methodName && method.descriptor == descriptor {
			return method
		}
	}

	return nil
}

