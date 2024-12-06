package mullitover

import "testing"

func Test_Run(t *testing.T) {
	var input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := Run(input)

	if result != 161 {
		t.Errorf("expected 161, got %d", result)
	}
}


func Test_RunWithOnOff(t *testing.T) {
	var input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := RunWithOnOff(input)

	if result != 48 {
		t.Errorf("expected 48, got %d", result)
	}
}
