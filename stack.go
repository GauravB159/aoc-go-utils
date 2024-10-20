package aocutils

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, true
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, false
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, true
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element, false
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
