package data

import "jvm/classfile"

type Field struct {
	ClassMember
	fieldId         uint
	constValueIndex uint
}

func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}

func (field *Field) ConstValueIndex() uint {
	return field.constValueIndex
}

func (field *Field) setValueAttribute(cField *classfile.MemberInfo) {
	if constValueAttr := cField.ConstValueAttribute(); constValueAttr != nil {
		field.constValueIndex = uint(constValueAttr.ConstantValueIndex())
	}
}

func newFields(class *Class, cFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cFields))

	for i, cField := range cFields {
		field := &Field{}
		field.copyFromMember(cField)
		field.SetClass(class)
		field.setValueAttribute(cField)
		fields[i] = field
	}

	return fields
}

func (field *Field) IsStatic() bool {
	return field.accessFlags.IsStatic()
}

func (field *Field) IsFinal() bool {
	return field.accessFlags.IsFinal()
}

func (field *Field) FieldId() uint {
	return field.fieldId
}