package data

type Object struct {
	class *Class
	//fields Variables
	data interface{}
}

func (obj *Object) Fields() Variables {
	return obj.data.(Variables)
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) IsImplements(class *Class) bool {
	return obj.class.isImplements(class)
}

func (obj *Object) SetRefVar(fieldName string, descriptor string, val *Object) {
	field := obj.class.getField(fieldName, descriptor, false)
	slots := obj.data.(Variables)
	slots.SetRef(field.fieldId, val)
}

func (obj *Object) GetRefVar(fieldName string, descriptor string) *Object {
	field := obj.class.getField(fieldName, descriptor, false)
	slots := obj.data.(Variables)
	return slots.GetRef(field.fieldId)
}
