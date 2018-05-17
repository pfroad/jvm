package runtime

import "jvm/runtime/data"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) TopFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) NewFrame(method *data.Method) *Frame {
	return NewFrame(t, method)
}

func (t *Thread) StackIsEmpty() bool {
	return t.stack.IsEmpty()
}
