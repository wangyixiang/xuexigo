package popcount

import "testing"

func TestPopCount(t *testing.T) {
	if PopCount(0xff) != 8 {
		t.Fail()
	}
}
