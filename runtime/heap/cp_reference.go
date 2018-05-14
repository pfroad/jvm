package heap

type SymbolRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (sym *SymbolRef) ResolveClass() *Class {
	if sym.class == nil {
		d := sym.cp.class
		c := d.classLoader.LoadClass(sym.className)
		if !c.isAccessTo(d) {
			panic("java.lang.IllegalAccessError")
		}
		sym.class = c
	}
	return sym.class
}
