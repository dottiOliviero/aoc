package lenslibrary

import "testing"

func TestHash(t *testing.T) {
	input := "HASH"
	result := Hash(input)

	if result != 52 {
		t.Error("result is not as expected")
	}
}

func TestCompute(t *testing.T) {
	result := Compute("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")
	if result != 1320 {
		t.Errorf("expected 1320, got %d", result)
	}
}

func TestComputePartTwo(t *testing.T) {
	result := ComputePartTwo("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")
	if result != 145 {
		t.Errorf("expected 145, got %d", result)
	}
}
