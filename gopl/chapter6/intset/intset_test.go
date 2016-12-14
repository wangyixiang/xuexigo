package main

import "testing"

func TestIntSet(t *testing.T) {
	intset := new(intSet)
	// testing boundary
	intset.set = []uint{(1 << uint(uintBits - 1)) + 1, (1 << uint(uintBits - 1)) + 1, (1 << uint(uintBits - 1)) + 1}
	if !(intset.Has(0) && intset.Has(64) && intset.Has(128)) {
		t.Failed()
	}

	testintset := new(intSet)
	testintset.Add(0)
	testintset.Add(63)
	testintset.Add(64)
	testintset.Add(127)
	testintset.Add(128)
	testintset.Add(191)

	if !compareIntSet(intset, testintset) {
		t.Failed()
	}

	if intset.Has(200) {
		t.Failed()
	}

	testintset.Add(-1)
	if testintset.Has(-1) {
		t.Failed()
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	s := new(intSet)
	a := new(intSet)
	st := new(intSet)

	initset := []uint{1, 1, 1}

	ttestcase := [][]uint{
		[]uint{2, 2},
		[]uint{2, 2, 2},
		[]uint{2, 2, 2, 2, 2},
	}

	stresult := [][]uint{
		[]uint{1, 3, 3},
		[]uint{3, 3, 3},
		[]uint{2, 2, 3, 3, 3},
	}

	for i := range ttestcase {
		s.set = initset
		a.set = ttestcase[i]
		st.set = stresult[i]
		s.UnionWith(a)
		if !(compareIntSet(s, st)) {
			t.Failed()
		}
	}
}

func compareIntSet(s1, s2 *intSet) bool {
	if len(s1.set) != len(s2.set) {
		return false
	}
	for i := range s1.set {
		if s1.set[i] != s2.set[i] {
			return false
		}
	}
	return true
}