package jslice

import (
	"fmt"
	"sort"
)

// Slice functions

// mutation

func (s *Slice[T]) Append(newElements ...T) {
	s.Data = append(s.Data, newElements...)
}

func (s *Slice[T]) Prepend(newElements ...T) {
	s.Data = append(newElements, s.Data...)
}

func (s *Slice[T]) Push(newElement T) {
	s.Data = append(s.Data, newElement)
}

func (s *Slice[T]) Pop() *T {
  if len(s.Data) == 0 {
    return nil
  }
  removed := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
  return &removed
}

// NO mutation

func (s *Slice[T]) Unique() *Slice[T] {
  keys := make(map[T]bool)
	list := []T{}
	for _, entry := range s.Data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return &Slice[T]{
		Data: list,
	}
}

func (s *Slice[T]) Contains(e T) bool {
  for _, v := range s.Data {
		if v == e {
			return true
		}
	}
	return false
}
func (s *Slice[T]) Includes(e T) bool {
  return s.Contains(e)
}

func (s *Slice[T]) Take(n int) *Slice[T] {
  if n > len(s.Data) {
		return &Slice[T]{
      Data: s.Data,
    }
	}

	return &Slice[T]{
    Data: s.Data[:n],
  }
}

func (s *Slice[T]) Clone() *Slice[T] {
  b := make([]T, len(s.Data))
  copy(b, s.Data)

	return &Slice[T]{
		Data: b,
	}
}

func (s *Slice[T]) Concat(arr []T) *Slice[T] {
  return &Slice[T]{
		Data: append(s.Data, arr...),
	}
}

func (s *Slice[T]) At(n int) T {
  length := len(s.Data)
  if n < 0 {
    return s.Data[length + (n % length)]
  } else if n > length - 1 {
    return s.Data[n % length]
  }
  return s.Data[n]
}

func (s *Slice[T]) Every(f func (el T) bool) bool {
  for _, e := range s.Data {
    if !f(e) {
      return false
    }
  }
  return true
}

func (s *Slice[T]) Some(f func (el T) bool) bool {
  for _, e := range s.Data {
    if f(e) {
      return true
    }
  }
  return false
}

func (s *Slice[T]) Filter(f func (el T) bool) *Slice[T] {
  b := make([]T, 0, len(s.Data))
  
  for _, e := range s.Data {
    if f(e) {
      b = append(b, e)
    }
  }

  return &Slice[T]{
		Data: b,
	}
}

func (s *Slice[T]) Map(f func (el T) T) *Slice[T] {
  b := make([]T, len(s.Data))
  
  for i, e := range s.Data {
    b[i] = f(e)
  }

  return &Slice[T]{
		Data: b,
	}
}

func (s *Slice[T]) IndexOf(el T) int {
  for i, e := range s.Data {
    if e == el {
      return i
    }
  }
  return -1
}

func (s *Slice[T]) LastIndexOf(el T) int {
  for i:=len(s.Data)-1; i >= 0; i-- {
    if s.Data[i] == el {
      return i
    }
  }
  return -1
}

func (s *Slice[T]) Find(f func (el T) bool) *T {
  for _, e := range s.Data {
    if f(e) {
      return &e
    }
  }
  return nil
}

func (s *Slice[T]) FindIndex(f func (el T) bool) int {
  for i, e := range s.Data {
    if f(e) {
      return i
    }
  }
  return -1
}

func (s *Slice[T]) FindLast(f func (el T) bool) *T {
  for i:=len(s.Data)-1; i >= 0; i-- {
    if f(s.Data[i]) {
      return &s.Data[i]
    }
  }
  return nil
}

func (s *Slice[T]) FindLastIndex(f func (el T) bool) int {
  for i:=len(s.Data)-1; i >= 0; i-- {
    if f(s.Data[i]) {
      return i
    }
  }
  return -1
}

func (s *Slice[T]) ForEach(f func (el T))  {
  for _, e := range s.Data {
    f(e)
  }
}

func (s *Slice[T]) Join(sep string) string {
  var str string
  for _, e := range s.Data {
    str = str + fmt.Sprint(e) + sep
  }
  return str[:len(str)-len(sep)]
}

func (s *Slice[T]) Keys() []int {
  keys := make([]int, len(s.Data))
  for i := range s.Data {
    keys[i] = i
  }
  return keys
}

func (s *Slice[T]) Reverse() *Slice[T] {
  reversed := make([]T, len(s.Data))

  index := 0
  for i:=len(s.Data)-1; i >= 0; i-- {
    reversed[index] = s.Data[i]
    index++
  }

  return &Slice[T]{
		Data: reversed,
	}
}

func (s *Slice[T]) Shift() *Slice[T] {
  shifted := s.Data[1:len(s.Data)]

  return &Slice[T]{
		Data: shifted,
	}
}

func (s *Slice[T]) Slice(start, end int) *Slice[T] {
  newArr := make([]T, 0, len(s.Data))

  for i, e := range s.Data {
    if i >= start {
      if end > 0 && i < end {
        newArr = append(newArr, e)
      }
    }
  }

  return &Slice[T]{
		Data: newArr,
	}
}

func (s *Slice[T]) Sort(f func(a,b int) bool) *Slice[T] {
  sorted := s.Clone()
  sort.Slice(sorted.Data, f)
  return sorted
}

func (s *Slice[T]) ToString() string {
  return s.Join(",")
}

func (s *Slice[T]) Unshift(elements ...T) *Slice[T] {
  return &Slice[T]{
		Data: append(elements, s.Data...),
	}
}

func (s *Slice[T]) Values() []T {
  return s.Data
}