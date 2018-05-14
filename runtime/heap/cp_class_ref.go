package heap

import "jvm/classfile"

type ClassRef struct {
	SymbolRef
}

func newClassRef(cp *ConstantPool, cfcc *classfile.ConstantClass) *ClassRef {
	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.className = cfcc.Name()
	//classRef.ResolveClass()
	return classRef
}
