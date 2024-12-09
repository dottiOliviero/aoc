package diskfragmenter

import "testing"

func Test_expandDiskMap(t *testing.T) {
	input := "2333133121414131402"
	result := expandDiskMap(input)

	if result.String() != "00...111...2...333.44.5555.6666.777.888899" {
		t.Errorf("got something not expected: %s", result)
	}

}

func Test_areThereGaps(t *testing.T) {
	testCases := []struct {
		name     string
		input    Elements
		expected bool
	}{
		{
			name:     "compact elements",
			input:    Elements{Element{value: 1}, Element{value: 1}, Element{value: 1}, Empty},
			expected: false,
		},
		{
			name:     "some gaps",
			input:    Elements{Element{value: 1}, Empty, Empty, Element{value: 1}, Element{value: 1}, Empty},
			expected: true,
		},
	}

	for _, tt := range testCases {
		result := areThereGaps(tt.input)
		if result != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result)
		}
	}
}

func Test_moveFileBlock(t *testing.T) {
	input := "2333133121414131402"
	result := expandDiskMap(input)

	compacted := moveFileBlocks(result)

	if compacted.String() != "0099811188827773336446555566.............." {
		t.Errorf("got %s, which is not as expected", compacted)
	}
}

func Test_checksum(t *testing.T) {
	input := "2333133121414131402"
	result := expandDiskMap(input)

	compacted := moveFileBlocks(result)

	checksum := computeCheckSum(compacted)
	if checksum != 1928 {
		t.Errorf("expected 1928, got %d", checksum)
	}
}
