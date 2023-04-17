package types

import (
	"math/rand"
	"time"
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

func (L *Leetcode) PickProblem() int {
	if len(L.solved) == L.total {
		return -1
	}

	prepare := []int{}
	for i, s := range L.problems {
		if !s {
			prepare = append(prepare, i)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(len(prepare))
	return prepare[random]
}

func (L *Leetcode) GetSolved() int {
	if len(L.solved) == L.total {
		return -1
	}

	rand.NewSource(time.Now().UnixNano())
	random := rand.Intn(L.Solved())
	return L.solved[random]
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
