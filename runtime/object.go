package runtime

import "jvm/runtime/heap"

type Object struct {
	class  *heap.Class
	fields Variables
}
