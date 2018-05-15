package data

import "jvm/classfile"

type FieldRef struct {
	MemberRef
	field      *Field
}

func newFieldRef(cp *ConstantPool, fRef *classfile.ConstantFieldRef) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	memberRef := fRef.ConstantMemberRef
	ref.className = memberRef.ClassName()
	ref.name, ref.descriptor = memberRef.NameAndDescriptor()
	return ref
}

func (ref *FieldRef) ResolveField() *Field {
	if ref.field == nil {
		d := ref.cp.class
		c := ref.ResolveClass()
		field := lookupField(ref.name, ref.descriptor, c)

		if field == nil {
			panic("java.lang.NoSuchFieldError")
		}

		if !c.isAccessTo(d) {
			panic("java.lang.IllegalAccessError")
		}

		ref.field = field
	}
	return ref.field
}

func lookupField(name string, descriptor string, class *Class) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, it := range class.interfaces {
		if field := lookupField(name, descriptor, it); field != nil {
			return field
		}
	}

	if superClass := class.superClass; superClass != nil {
		if field := lookupField(name, descriptor, superClass); field != nil {
			return field
		}
	}

	return nil
}
