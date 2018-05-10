package heap

import "jvm/classfile"

type Field struct {
	ClassMember
	fieldId uint
}

func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}

func newFields(class *Class, cFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cFields))

	for i, cField := range cFields {
		field := &Field{}
		field.copyFromMember(cField)
		field.SetClass(class)
		fields[i] = field
	}

	return fields
}
