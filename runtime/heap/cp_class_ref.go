package heap

import "jvm/classfile"

type ClassRef struct {
	SymbolRef
}

func newClassRef(cp *ConstantPool, cfcc *classfile.ConstantClass) *ClassRef {
	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.ResolveClassRef(cfcc.Name())
	return classRef
}
