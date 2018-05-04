package runtime

// A new frame is created each time a method is invoked. A frame is destroyed when
// its method invocation completes, whether that completion is normal or abrupt (it
// throws an uncaught exception)
// Only one frame, the frame for the executing method, is active at any point in a given
// thread of control
type Frame struct {
	previous     *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) PC() int {
	return f.nextPC
}

func (f *Frame) SetPC(pc int) {
	f.nextPC = pc
}
