package data

import "jvm/classfile"

type ClassMember struct {
	accessFlags AccessFlags
	name        string
	descriptor  string
	class       *Class
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

func (cm *ClassMember) Descriptor() string {
	return cm.descriptor
}

func (cm *ClassMember) Class() *Class {
	return cm.class
}

func (cm *ClassMember) Name() string {
	return cm.name
}

func (cm *ClassMember) IsPublic() bool {
	return cm.accessFlags.IsPublic()
}

func (cm *ClassMember) IsPrivate() bool {
	return cm.accessFlags.IsPrivate()
}

func (cm *ClassMember) IsProtected() bool {
	return cm.accessFlags.IsProtected()
}

func (cm *ClassMember) isAccessibleTo(other *Class) bool {
	if cm.IsPublic() {
		return true
	}

	class := cm.class
	if cm.IsPrivate() {
		return class == other
	} else {
		return other.GetPackageName() == class.GetPackageName() ||
			(cm.IsProtected() && (other.isExtendClass(class)))
	}
}