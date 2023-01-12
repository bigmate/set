package set

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSet(t *testing.T) {
	s := New[float32](100)
	s.Add(12.54)
	s.Add(12.5412)

	if !s.Has(12.54) {
		t.Errorf("12.54 not found")
	}

	if !s.Has(12.5412) {
		t.Errorf("12.5412 not found")
	}

	els := s.Els()

	if len(els) != 2 {
		t.Errorf("invalid amount of elements returned")
		t.Logf("elements: %v", els)
	}

	s.Del(12.54)
	s.Del(12.5412)

	if s.Len() != 0 {
		t.Errorf("set is not empty")
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		a, b, c []int
	}{
		{
			a: []int{1, 2, 3, 4, 5, 6},
			b: []int{3, 4, 5, 6, 7, 8},
			c: []int{3, 4, 5, 6},
		},
		{
			a: []int{1, 2, 3, 4},
			b: []int{1, 3, 4},
			c: []int{1, 3, 4},
		},
		{
			a: []int{1, 2, 3},
			b: []int{1, 2, 3, 4},
			c: []int{1, 2, 3},
		},
		{
			a: []int{1, 2, 3},
			b: []int{1, 2, 2, 3},
			c: []int{1, 2, 3},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			els := Intersection(
				FromSlice(tt.a...),
				FromSlice(tt.b...),
			).Els()

			sort.Ints(els)

			if !reflect.DeepEqual(els, tt.c) {
				t.Errorf("unexpected result: %v", els)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		a, b, c []int
	}{
		{
			a: []int{1, 2, 3, 4, 5, 6},
			b: []int{3, 4, 5, 6, 7, 8},
			c: []int{1, 2, 7, 8},
		},
		{
			a: []int{1, 2, 3, 4},
			b: []int{1, 3, 4},
			c: []int{2},
		},
		{
			a: []int{1, 2, 3},
			b: []int{1, 2, 3, 4},
			c: []int{4},
		},
		{
			a: []int{1, 2, 3},
			b: []int{1, 2, 2, 3},
			c: []int{},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.FormatInt(int64(i+1), 10), func(t *testing.T) {
			els := Difference(
				FromSlice(tt.a...),
				FromSlice(tt.b...),
			).Els()

			sort.Ints(els)

			if !reflect.DeepEqual(els, tt.c) {
				t.Errorf("unexpected result: %v", els)
			}
		})
	}
}
