package historianhysteria

import "testing"

func TestRun(t *testing.T) {
	l := []int{3, 4, 2, 1, 3, 3}
	r := []int{4, 3, 5, 3, 9, 3}
	result := Run(l, r)
	if result != 11 {
		t.Errorf("expected 11, got %d", result)
	}
}
