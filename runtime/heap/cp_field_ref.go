package heap

import "jvm/classfile"

type FieldRef struct {
	SymbolRef
	field *Field
}

func newFieldRef(cp *ConstantPool, fRef *classfile.ConstantFieldRef) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	memberRef := fRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	ref.resolveFieldRef(memberRef.NameAndDescriptor())
	return ref
}

func (ref *FieldRef) resolveFieldRef(name string, descriptor string) {
	ref.ResolveClass()
}
