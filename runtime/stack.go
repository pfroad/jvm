package runtime

type Stack struct {
	maxSize  uint
	size     uint
	topFrame *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s.topFrame != nil {
		frame.previous = s.topFrame
	}

	s.topFrame = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s.topFrame == nil {
		panic("jvm stack is empty!")
	}

	frame := s.topFrame
	s.topFrame = frame.previous
	s.size--
	return frame
}

func (s *Stack) top() *Frame {
	if s.topFrame == nil {
		panic("jvm stack is empty!")
	}

	return s.topFrame
}

func (s *Stack) IsEmpty() bool {
	return s.topFrame == nil
}
