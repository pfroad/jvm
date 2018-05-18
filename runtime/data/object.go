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
