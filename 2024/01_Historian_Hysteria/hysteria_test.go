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

func TestRunSimilarity(t *testing.T) {
	l := []int{3, 4, 2, 1, 3, 3}
	r := []int{4, 3, 5, 3, 9, 3}
	result := RunSimilarityScore(l, r)
	if result != 31 {
		t.Errorf("expected 31, got %d", result)
	}
}
