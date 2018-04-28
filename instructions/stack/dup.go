package stack

/*
Table 2.11.1-B. Actual and Computational types in the Java Virtual Machine
Actual type 	Computational type 	Category
boolean 	int 	1
byte 	int 	1
char 	int 	1
short 	int 	1
int 	int 	1
float 	float 	1
reference 	reference 	1
returnAddress 	returnAddress 	1
long 	long 	2
double 	double 	2
*/

import (
	"jvm/instructions/common"
	"jvm/runtime"
)

// dup = 89 (0x59)
type Dup struct {
	common.NoOperandsInstruction
}

// dup_x1 = 90 (0x5a)
type DupX1 struct {
	common.NoOperandsInstruction
}

// dup_x2 = 91 (0x5b)
type DupX2 struct {
	common.NoOperandsInstruction
}

// dup2 = 92 (0x5c)
type Dup2 struct {
	common.NoOperandsInstruction
}

// dup2_x1 = 93 (0x5d)
type Dup2X1 struct {
	common.NoOperandsInstruction
}

// dup2_x2 = 94 (0x5e)
type Dup2X2 struct {
	common.NoOperandsInstruction
}

/* Duplicate the top operand stack value
bottom -> top
..., value →
..., value, value
The dup instruction must not be used unless value is a value of a category 1 computational type (§2.11.1).
*/
func (dup *Dup) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	top := stack.Pop()
	stack.Push(top)
	stack.Push(top)
}

/* Duplicate the top operand stack value and insert two values down
bottom -> top
..., value2, value1 →
..., value1, value2, value1
The dup_x1 instruction must not be used unless both value1 and value2 are values of a category 1 computational type (§2.11.1).
*/
func (dup *DupX1) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	stack.Push(v1)
	stack.Push(v2)
	stack.Push(v1)
}

/* Duplicate the top operand stack value and insert two or three values down
bottom -> top
Form 1:
..., value3, value2, value1 →
..., value1, value3, value2, value1
where value1, value2, and value3 are all values of a category 1 computational type (§2.11.1).
Form 2:
..., value2, value1 →
..., value1, value2, value1
where value1 is a value of a category 1 computational type and value2 is a value of a category 2 computational type (§2.11.1).
*/
func (dup *DupX2) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	v3 := stack.Pop()
	stack.Push(v1)
	stack.Push(v3)
	stack.Push(v2)
	stack.Push(v1)
}

/* Duplicate the top one or two operand stack values
bottom -> top
Form 1:
..., value2, value1 →
..., value2, value1, value2, value1
where both value1 and value2 are values of a category 1 computational type (§2.11.1).
Form 2:
..., value →
..., value, value
where value is a value of a category 2 computational type (§2.11.1).
*/
func (dup *Dup2) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	stack.Push(v2)
	stack.Push(v1)
	stack.Push(v2)
	stack.Push(v1)
}

/* duplicate top one (long or double) or two operand stack value and insert 2 or 3 values down
bottom -> top
Form 1:
..., value3, value2, value1 →
..., value2, value1, value3, value2, value1
where value1, value2, and value3 are all values of a category 1 computational type (§2.11.1).
Form 2:
..., value2, value1 →
..., value1, value2, value1
where value1 is a value of a category 2 computational type and value2 is a value of a category 1 computational type (§2.11.1).
*/
func (dup *Dup2X1) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	v3 := stack.Pop()
	stack.Push(v2)
	stack.Push(v1)
	stack.Push(v3)
	stack.Push(v2)
	stack.Push(v1)
}

/* duplicate top one (long or double) or two operand stack value and insert 2, 3 or 4 values down
bottom -> top
Form 1:
..., value4, value3, value2, value1 →
..., value2, value1, value4, value3, value2, value1
where value1, value2, value3, and value4 are all values of a category 1 computational type.
Form 2:
..., value3, value2, value1 →
..., value1, value3, value2, value1
where value1 is a value of a category 2 computational type and value2 and value3 are both values of a category 1 computational type.
Form 3:
..., value3, value2, value1 →
..., value2, value1, value3, value2, value1
where value1 and value2 are both values of a category 1 computational type and value3 is a value of a category 2 computational type.
Form 4:
..., value2, value1 →
..., value1, value2, value1
where value1 and value2 are both values of a category 2 computational type.
*/
func (dup *Dup2X2) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.Pop()
	v2 := stack.Pop()
	v3 := stack.Pop()
	v4 := stack.Pop()
	stack.Push(v2)
	stack.Push(v1)
	stack.Push(v4)
	stack.Push(v3)
	stack.Push(v2)
	stack.Push(v1)
}