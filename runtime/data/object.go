package data

type Object struct {
	class  *Class
	fields Variables
}

func (obj *Object) Fields() Variables {
	return obj.fields
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}