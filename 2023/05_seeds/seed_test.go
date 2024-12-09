package seeds

import "testing"

func TestAssngMin(t *testing.T) {
	l := []int{13, 3, 4, 5, 2, 5, 6, 21}
	m := -1
	for i := 0; i < len(l); i++ {
		m = assingMin(m, l[i])
	}
	if m != 2 {
		t.Errorf("min is not 2")
	}
}
