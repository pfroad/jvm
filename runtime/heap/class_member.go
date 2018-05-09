package heap

import "jvm/classfile"

type ClassMember struct {
	accessFlags AccessFlags
	name string
	descriptor string
	class *Class
}

func (cm *ClassMember) copyFromMember(cfMember *classfile.MemberInfo) {
	cm.SetAccessFlags(AccessFlags{cfMember.AccessFlag()})
	cm.SetName(cfMember.Name())
	cm.SetDescriptor(cfMember.Descriptor())
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
