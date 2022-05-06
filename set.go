package slice

import "golang.org/x/exp/constraints"

type Set[T constraints.Ordered] interface {
	Add(el T)
	Del(el T)
	Has(el T) bool
	Len() int
	Els() []T
}

func NewSet[T constraints.Ordered](size int) Set[T] {
	return make(set[T], size)
}

func FromSlice[T constraints.Ordered](els ...T) Set[T] {
	s := make(set[T], len(els))

	for _, el := range els {
		s.Add(el)
	}

	return s
}

type set[T constraints.Ordered] map[T]struct{}

func (s set[T]) Add(el T) {
	s[el] = struct{}{}
}

func (s set[T]) Del(el T) {
	delete(s, el)
}

func (s set[T]) Has(el T) bool {
	_, has := s[el]
	return has
}

func (s set[T]) Len() int {
	return len(s)
}

func (s set[T]) Els() []T {
	els := make([]T, 0, len(s))

	for k := range s {
		els = append(els, k)
	}

	return els
}
