package data

import "jvm/classfile"

type MethodRef struct {
	MemberRef
	method     *Method
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

//func (ref *MethodRef) resolveMethodRef(name string, descriptor string) {
//
//}

type InterfaceMethodRef struct {
	MemberRef
	method     *Method
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
