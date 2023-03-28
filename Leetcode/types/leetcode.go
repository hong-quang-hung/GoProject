package types

import (
	"math/rand"
)

type Leetcode struct {
	total    int
	problems []bool
	solved   []int
}

func (L *Leetcode) SetTotal(size int) {
	L.total = size
	L.problems = make([]bool, size)
	L.solved = []int{}
}

func (L *Leetcode) SetSolved(s ...int) {
	for _, i := range s {
		if L.IsValid(i - 1) {
			L.problems[i-1] = true
			L.solved = append(L.solved, i)
		}
	}
}

func (L *Leetcode) GetRandom() int {
	random := rand.Intn(L.total)
	for L.IsSolved(random) {
		random = rand.Intn(L.total)
	}
	return random
}

func (L *Leetcode) Solved() int {
	return len(L.solved)
}

func (L *Leetcode) Total() int {
	return L.total
}

func (L *Leetcode) IsSolved(i int) bool {
	return i <= L.total && L.problems[i]
}

func (L *Leetcode) IsValid(i int) bool {
	return i <= L.total && !L.problems[i]
}
