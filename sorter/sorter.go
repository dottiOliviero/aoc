package sorter

import (
	"slices"
	"sort"
)

type Coin struct {
	Value    int
	Currency string
}

type lessFunc func(p1, p2 *Coin) bool

type sorter struct {
	coins    []Coin
	lessFunc lessFunc
}

// Len implements sort.Interface.
func (s *sorter) Len() int {
	return len(s.coins)
}

// Less implements sort.Interface.
func (s *sorter) Less(i int, j int) bool {
	return s.lessFunc(&s.coins[i], &s.coins[j])
}

func amount(c1, c2 *Coin) bool {
	comparison := compareCoins(*c1, *c2)
	switch {
	case comparison >= 0:
		return false
	case comparison < 0:
		return true
	default:
		return false
	}
}

// Swap implements sort.Interface.
func (s *sorter) Swap(i int, j int) {
	s.coins[i], s.coins[j] = s.coins[j], s.coins[i]
}

func (s *sorter) Sort(coins []Coin) {
	s.coins = coins
	sort.Sort(s)
}

func OrderedBy(less lessFunc) *sorter {
	return &sorter{
		lessFunc: less,
	}
}

func compareCoins(c1, c2 Coin) int {
	switch {
	case c1.Value < c2.Value:
		return -1
	case c1.Value == c2.Value:
		return 0
	default:
		return 1
	}
}

func SortMeWithSorter(coins []Coin) {
	OrderedBy(amount).Sort(coins)

}

func SortMe(coins []Coin) {
	slices.SortFunc(coins, compareCoins)
}
