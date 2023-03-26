package types

import (
	"math/rand"
	"time"
)

type Leetcode struct {
	total  int
	solved []int
}

func (L *Leetcode) SetTotalProblem(n int) {
	L.total = n
}

func (L *Leetcode) SetSolvedProblems(s ...int) {
	if L.solved == nil {
		L.solved = []int{}
	}
	L.solved = append(L.solved, s...)
}

func (L *Leetcode) GetRandomProblem() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(L.total)
}

func (L *Leetcode) Solved() int {
	return len(L.solved)
}
