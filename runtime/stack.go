package runtime

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s._top != nil {
		frame.previous = s._top
	}

	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	frame := s._top
	s._top = frame.previous
	s.size--
	return frame
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	return s._top
}
