package jslice

type Slice[T comparable] struct {
	Data []T
}

func New[T comparable]() *Slice[T] {
	return &Slice[T]{
		Data: []T{},
	}
}

func Of[T comparable](elements ...T) *Slice[T] {
	return &Slice[T]{
		Data: elements,
	}
}

func From[T comparable](elements []T) *Slice[T] {
	return &Slice[T]{
		Data: elements,
	}
}

func (s *Slice[T]) Length() int {
  return len(s.Data)
}

func (s *Slice[T]) String() string {
  return "[" + s.ToString() + "]"
}