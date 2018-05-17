package data

import "jvm/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, mRef *classfile.ConstantMethodRef) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	memberRef := mRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	//ref.ResolveClass()
	ref.name, ref.descriptor = memberRef.NameAndDescriptor()
	return ref
}

func (ref *MethodRef) ResolveMethod() *Method {
	if ref.method == nil {
		d := ref.cp.class
		c := ref.ResolveClass()

		if c.IsInterface() {
			panic("java.lang.IncompatibleClassChangeError")
		}

		method := lookupMethod(c, ref.name, ref.descriptor)

		if method == nil {
			panic("java.lang.NoSuchMethodError")
		}

		if !method.isAccessibleTo(d) {
			panic("java.lang.IllegalAccessError")
		}

		ref.method = method
	}

	return ref.method
}

func lookupMethod(class *Class, name , descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)

	if method == nil {
		method = lookupMethodInInterface(class.interfaces, name, descriptor)
	}

	return method
}

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, imRef *classfile.ConstantInterfaceMethodRef) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	memberRef := imRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	//ref.ResolveClass()
	ref.name, ref.descriptor = memberRef.NameAndDescriptor()
	return ref
}

func (ref *InterfaceMethodRef) ResolveInterfaceMethod() *Method {
	if ref.method == nil {
		d := ref.cp.class
		c := ref.ResolveClass()

		if !c.IsInterface() {
			panic("java.lang.IncompatibleClassChangeError")
		}

		method := lookupInterfaceMethod(c, ref.name, ref.descriptor)

		if method == nil {
			panic("java.lang.NoSuchMethodError")
		}

		if !method.isAccessibleTo(d) {
			panic("java.lang.IllegalAccessError")
		}

		ref.method = method
	}

	return ref.method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterface(iface.interfaces, name, descriptor)
}
