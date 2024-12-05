package camelcard

import (
	"reflect"
	"testing"
)

func TestCountCards(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{
			input: "AATKJ",
			expected: []int{
				2, 1, 1, 1,
			},
		},
		{
			input: "6TT8T",
			expected: []int{
				1, 3, 1,
			},
		},
		{
			input: "27A83",
			expected: []int{
				1, 1, 1, 1, 1,
			},
		},
		{
			input: "4884J",
			expected: []int{
				2, 2, 1,
			},
		},
		{
			input: "AAAAA",
			expected: []int{
				5,
			},
		},
	}
	for _, tt := range testCases {
		result := countCards(tt.input)
		if !reflect.DeepEqual(tt.expected, result) {
			t.Errorf("expected %v, got %v", tt.expected, result)
		}

	}
}

func TestPower(t *testing.T) {
	testCases := []struct {
		input       Hand
		expectedPow int
	}{
		{
			input: Hand{
				cards: "4884J",
			},
			expectedPow: 3,
		},
		{
			input: Hand{
				cards: "255T5",
			},
			expectedPow: 4,
		},
		{
			input: Hand{
				cards: "4444K",
			},
			expectedPow: 6,
		},
		{
			input: Hand{
				cards: "48848",
			},
			expectedPow: 5,
		},
		{
			input: Hand{
				cards: "23456",
			},
			expectedPow: 1,
		},
		{
			input: Hand{
				cards: "AATKJ",
			},
			expectedPow: 2,
		},
	}
	for _, tt := range testCases {
		result := setPower(tt.input)
		if tt.expectedPow != result.power {
			t.Errorf("expected pow %d, got %d", tt.expectedPow, result.power)
		}
	}
}

func TestComputeRank(t *testing.T) {
	hands := []Hand{
		{
			cards: "AATKJ",
			power: 2,
			bid:   840,
		},
		{
			cards: "27A83",
			power: 1,
			bid:   251,
		},
		{
			cards: "6TT8T",
			power: 4,
			bid:   113,
		},
		{
			cards: "Q6T6T",
			power: 3,
			bid:   693,
		},
		{
			cards: "99K89",
			power: 4,
			bid:   553,
		},
	}

	computeRank(hands, LABELS)

	if !reflect.DeepEqual(hands[0], Hand{
		cards: "27A83",
		power: 1,
		bid:   251,
	}) {
		t.Errorf("error in compute ranking, got %v", hands[0])
	}
}

func Test5Rounds(t *testing.T) {
	INPUT := "AATKJ 840\n27A83 251\n6TT8T 113\nQ6T6T 693\n99K89 553"
	// 27A83 251
	// AATKJ 840
	// Q6T6T 693
	// 6TT8T 113
	// 99K89 553

	result := Compute(INPUT)
	expectedResult := (251 + 840*2 + 693*3 + 113*4 + 553*5)

	if result != expectedResult {
		t.Errorf("expected %d: got %d", expectedResult, result)
	}
}

func TestExample(t *testing.T) {
	INPUT := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	result := Compute(INPUT)
	expectedResult := 6440

	if result != expectedResult {
		t.Errorf("expected %d: got %d", expectedResult, result)
	}
}

func TestRunWithJoker(t *testing.T) {
	INPUT := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	result := RunWithJoker(INPUT)

	expectedResult := 5905

	if result != expectedResult {
		t.Errorf("expected %d: got %d", expectedResult, result)
	}
}

func TestSetPowerWithJoker(t *testing.T) {

	testCases := []struct {
		cards       string
		expectedPow int
	}{
		{
			cards:       "JJJJJ", // 5 all J
			expectedPow: 7,
		},
		{
			cards:       "QQQQQ", // 5 no J
			expectedPow: 7,
		},
		{
			cards:       "KKKK7", // poker no J
			expectedPow: 6,
		},
		{
			cards:       "KKKKJ", // poker + J
			expectedPow: 7,
		},
		{
			cards:       "23332", // full house no J
			expectedPow: 5,
		},
		{
			cards:       "2333J", // tris + J
			expectedPow: 6,
		},
		{
			cards:       "J333J", // tris + 2J
			expectedPow: 7,
		},
		{
			cards:       "3332K", // tris no J
			expectedPow: 4,
		},
		{
			cards:       "3332J", // tris + J
			expectedPow: 6,
		},
		{
			cards:       "333JJ", // tris + 2J
			expectedPow: 7,
		},
		{
			cards:       "33JJJ", // 3J + 2 equal
			expectedPow: 7,
		},
		{
			cards:       "3TJJJ", // 3J + 2 diff --> poker
			expectedPow: 6,
		},
	}

	for _, tt := range testCases {

		result := setPowerWithJoker(Hand{
			cards: tt.cards,
		})
		if v := result.power; v != tt.expectedPow {
			t.Errorf("expected %d, got %d", tt.expectedPow, v)
		}
	}
}
