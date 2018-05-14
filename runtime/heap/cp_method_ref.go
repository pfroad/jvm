package heap

import "jvm/classfile"

type MethodRef struct {
	SymbolRef
	method *Method
}

func newMethodRef(cp *ConstantPool, mRef *classfile.ConstantMethodRef) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	memberRef := mRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	//ref.ResolveClass()
	ref.resolveMethodRef(memberRef.NameAndDescriptor())
	return ref
}

func (ref *MethodRef) resolveMethodRef(name string, descriptor string) {

}

type InterfaceMethodRef struct {
	SymbolRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, mRef *classfile.ConstantMethodRef) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	memberRef := mRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	//ref.ResolveClass()
	ref.resolveMethodRef(memberRef.NameAndDescriptor())
	return ref
}