package heap

import "jvm/classfile"

type Field struct {
	ClassMember
} 

func newFields(class *Class, cFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cFields))

	for i, cField := range cFields {
		field := &Field{}
		field.SetAccessFlags(AccessFlags{cField.AccessFlag()})
		field.SetName(cField.Name())
		field.SetDescriptor(cField.Descriptor())
		field.SetClass(class)
		fields[i] = field
	}

	return fields
}