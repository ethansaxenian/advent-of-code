package stack

type Stack[T comparable] []T

func (s Stack[T]) Push(item T) Stack[T] {
	return append(s, item)
}

func (s Stack[T]) Pop() (Stack[T], T) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}

func NewStack[T comparable](items ...T) Stack[T] {
	s := Stack[T]{}
	for _, i := range items {
		s = s.Push(i)
	}
	return s
}
