package data

type MemberRef struct {
	SymbolRef
	name       string
	descriptor string
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}
