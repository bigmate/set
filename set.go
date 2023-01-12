package set

type Set[T comparable] interface {
	Add(el T)
	Del(el T)
	Has(el T) bool
	Len() int
	Els() []T
}

func New[T comparable](size int) Set[T] {
	return make(set[T], size)
}

func FromSlice[T comparable](els ...T) Set[T] {
	s := make(set[T], len(els))

	for _, el := range els {
		s.Add(el)
	}

	return s
}

type set[T comparable] map[T]struct{}

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

func extract[T comparable](a, b, c Set[T], hasFlip bool) {
	els := b.Els()

	for i := range els {
		k := els[i]

		if a.Has(k) == hasFlip {
			c.Add(k)
		}
	}
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	res := New[T](0)

	extract(a, b, res, true)

	return res
}

func Difference[T comparable](a, b Set[T]) Set[T] {
	res := New[T](0)

	extract(a, b, res, false)
	extract(b, a, res, false)

	return res
}
