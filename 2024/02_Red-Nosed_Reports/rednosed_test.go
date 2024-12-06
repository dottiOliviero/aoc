package rednosedreports

import "testing"

func Test_HowManyAreSafe(t *testing.T) {
	var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := HowManyAreSafe(input)
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

func Test_HowManyAreSafeWithDamper(t *testing.T) {
	var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := HowManyAreSafeWithDamper(input)
	if result != 4 {
		t.Errorf("expected 4, got %d", result)
	}
}
