package ds

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(t T) {
	s.data = append(s.data, t)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var t T
		return t, false
	}
	t := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return t, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var t T
		return t, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
