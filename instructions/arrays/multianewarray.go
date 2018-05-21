package arrays

import (
	"jvm/runtime"
	"jvm/runtime/data"
)

type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (multi *MultiANewArray) FetchOperands(reader *runtime.ByteCodeReader) {
	multi.index = reader.ReadUint16()
	multi.dimensions = reader.ReadUint8()
}

func (multi *MultiANewArray) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConst(uint(multi.index)).(*data.ClassRef)
	resolveClass := classRef.ResolveClass()
	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, multi.dimensions)
	arr := newMultiDimensionalArray(counts, resolveClass)
	stack.PushRef(arr)
}

func newMultiDimensionalArray(counts []int32, resolveClass *data.Class) *data.Object {
	arr := resolveClass.NewArray(uint(counts[0]))

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], resolveClass.ComponentClass())
		}
	}

	return arr
}

func popAndCheckCounts(stack *data.OperandStack, dimensions uint8) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		count := stack.PopInt()
		if count < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
		counts[i] = count
	}
	return counts
}
