package sorter_test

import (
	"advent_of_code/sorter"
	"testing"
)

func TestRegularSorter(t *testing.T) {
	coins := []sorter.Coin{
		{
			Value:    12,
			Currency: "EUR",
		},
		{
			Value:    44,
			Currency: "EUR",
		},
		{
			Value:    4,
			Currency: "USD",
		},
	}
	sorter.SortMeWithSorter(coins)
	if v := coins[0].Value; v != 4 {
		t.Errorf("sort not correctly, got %d", v)
	}
}

func TestWithSliceSorter(t *testing.T) {
	coins := []sorter.Coin{
		{
			Value:    12,
			Currency: "EUR",
		},
		{
			Value:    44,
			Currency: "EUR",
		},
		{
			Value:    4,
			Currency: "USD",
		},
	}
	sorter.SortMe(coins)
	if v := coins[0].Value; v != 4 {
		t.Errorf("sort with slice not correctly, got %d", v)
	}
}
