package runtime

import (
	"jvm/runtime/data"
)

// A new frame is created each time a method is invoked. A frame is destroyed when
// its method invocation completes, whether that completion is normal or abrupt (it
// throws an uncaught exception)
// Only one frame, the frame for the executing method, is active at any point in a given
// thread of control
type Frame struct {
	previous     *Frame
	localVars    data.Variables
	operandStack *data.OperandStack
	thread       *Thread
	method       *data.Method
	//nextPC       int
}

func NewFrame(thread *Thread, method *data.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    data.NewVariables(method.MaxLocals()),
		operandStack: data.NewOperandStack(method.MaxStack()),
	}
}

func (f *Frame) LocalVars() data.Variables {
	return f.localVars
}

func (f *Frame) OperandStack() *data.OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) PC() int {
	return f.thread.pc
}

func (f *Frame) SetPC(pc int) {
	f.thread.SetPC(pc)
}

func (f *Frame) Method() *data.Method {
	return f.method
}
