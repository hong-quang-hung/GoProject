package types

import (
	"math/rand"
	"time"
)

type Leetcode struct {
	total     int
	submitLen int
	problems  []bool
	solved    []int
	submit    []bool
}

func (L *Leetcode) SetTotal(size int) {
	L.total = size
	L.problems = make([]bool, size)
	L.submit = make([]bool, size)
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

func (L *Leetcode) SetSubmit(s ...int) {
	for _, i := range s {
		if L.IsSolved(i-1) && !L.submit[i-1] {
			L.submit[i-1] = true
			L.submitLen++
		}
	}
}

func (L *Leetcode) PickProblem() int {
	if len(L.solved) == L.total {
		return -1
	}

	rand.NewSource(time.Now().UnixNano())
	random := rand.Intn(L.total)
	for L.IsSolved(random) {
		random = rand.Intn(L.total)
	}
	return random + 1
}

func (L *Leetcode) GetSolved() int {
	if len(L.solved) == L.total {
		return -1
	}

	rand.NewSource(time.Now().UnixNano())
	random := rand.Intn(L.Solved())
	return L.solved[random]
}

func (L *Leetcode) PickSolution() int {
	if len(L.solved) == L.total {
		return -1
	}

	notSubmit := []int{}
	for _, s := range L.solved {
		if !L.submit[s-1] {
			notSubmit = append(notSubmit, s)
		}
	}

	if len(notSubmit) == 0 {
		return -1
	}

	rand.NewSource(time.Now().UnixNano())
	random := rand.Intn(len(notSubmit))
	return notSubmit[random]
}

func (L *Leetcode) Solved() int {
	return len(L.solved)
}

func (L *Leetcode) Submit() int {
	return L.submitLen
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
