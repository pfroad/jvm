package heap

type SymbolRef struct {
	cp *ConstantPool
	class *Class
}

func (sym *SymbolRef) ResolveClassRef(name string) {
	
}