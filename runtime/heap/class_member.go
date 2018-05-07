package heap

type ClassMember struct {
	accessFlags AccessFlags
	name string
	descriptor string
	class *Class
}

func (cm *ClassMember) SetAccessFlags(flags AccessFlags) {
	cm.accessFlags = flags
}

func (cm *ClassMember) SetName(name string) {
	cm.name = name
}

func (cm *ClassMember) SetDescriptor(descriptor string) {
	cm.descriptor = descriptor
}

func (cm *ClassMember) SetClass(class *Class) {
	cm.class = class
}
