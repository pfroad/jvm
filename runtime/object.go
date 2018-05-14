package runtime

import "jvm/runtime/heap"

type Object struct {
	class  *heap.Class
	fields Variables
}

func NewObject(class *Class, fields Variables) *Object {
	return
}
