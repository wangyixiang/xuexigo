package main

import (
	"reflect"
)

var uintBits = int(reflect.TypeOf(uint(1)).Size() * 8)

type intSet struct {
	set []uint
}

func (s *intSet) Has(x int) bool {
	if x >= 0 {
		wPos, bPos := x / uintBits, uint(x % uintBits)
		if wPos < len(s.set) && (s.set[wPos] & (1 << bPos) != 0) {
			return true
		}
	}
	return false
}

func (s *intSet) Has2(x int) bool {
	word, bit := x / 64, uint(x % 64)
	return word < len(s.set) && s.set[word] & (1 << bit) != 0
}

func (s *intSet) Add(x int) {
	if x < 0 {
		return
	}
	wPos, bPos := x / uintBits, uint(x % uintBits)
	if wPos >= len(s.set) {
		for i := len(s.set); i <= wPos; i++ {
			s.set = append(s.set, 0)
		}
	}
	s.set[wPos] = s.set[wPos] | 1 << bPos
}

func (s *intSet) Add2(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(s.set) {
		s.set = append(s.set, 0)
	}
	s.set[word] |= 1 << bit
}

func (s *intSet) UnionWith(t *intSet) {
	for i := range t.set {
		if i < len(s.set) {
			s.set[i] |= t.set[i]
		} else {
			s.set = append(s.set, t.set[i])
		}
	}
}

func (s *intSet) UnionWith2(t *intSet) {
	for i, tword := range t.set {
		if i < len(s.set) {
			s.set[i] |= tword
		} else {
			s.set = append(s.set, tword)
		}
	}
}

//func (s *intSet) String() string  {
//
//}

//Len
//Remove
//Clear
//Copy
//AddAll variadic
//intersectWith
//differenceWith
//symmetricDifference

func main() {
	intset := new(intSet)
	intset.Add(0x2ffffffff)
}
