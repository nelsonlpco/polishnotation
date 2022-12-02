package domain

type Stack[T comparable] struct {
	data   []T
	length int
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s Stack[T]) Value() T {
	return s.data[s.length-1]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{length: 0}
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
	s.length++
}

func (s *Stack[T]) Pop() T {
	element := s.data[s.length-1]

	s.length--
	s.data = s.data[:s.length]

	return element
}
