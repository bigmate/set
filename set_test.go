package slice

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet[float32](100)
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
