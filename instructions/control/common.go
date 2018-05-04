package control

import "jvm/runtime"

const (
	EQ = 1
	NE = 2
	LT = 3
	GT = 4
	LE = 5
	GE = 6
)

// int comparison
func compare(v1, v2 int32, comparator int) bool {
	switch comparator {
	case EQ:
		return v1 == v2
	case NE:
		return v1 != v2
	case LT:
		return v1 < v2
	case GT:
		return v1 > v2
	case LE:
		return v1 <= v2
	case GE:
		return v1 >= v2
	default:
		panic("Unknown Comparator!")
	}
}

// int comparison condition
func condition(frame *runtime.Frame, comparator int) bool {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	return compare(v1, v2, comparator)
}

// ref comparison condition
func aCondition(frame *runtime.Frame, comparator int) bool {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()

	switch comparator {
	case EQ:
		return v1 == v2
	default:
		return v1 != v2
	}
}