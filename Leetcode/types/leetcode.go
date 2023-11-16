package types

import (
	"math/rand"
	"sync"
	"time"
)

type Leetcode struct {
	total    int
	problems []bool
	solved   int
	count    []int
	mu       sync.Mutex
}

func (L *Leetcode) SetTotal(size int) {
	L.total = size
	L.problems = make([]bool, size)
	L.count = make([]int, size)
	L.solved = 0
}

func (L *Leetcode) SetSolved(wg *sync.WaitGroup, s ...int) {
	defer wg.Done()
	vl := 0
	for _, i := range s {
		if L.IsValid(i - 1) {
			L.problems[i-1] = true
			vl++
		}

		if i > 0 && i < L.total {
			L.mu.Lock()
			L.count[i-1]++
			L.mu.Unlock()
		}
	}

	L.mu.Lock()
	L.solved += vl
	L.mu.Unlock()
}

func (L *Leetcode) PickProblem() int {
	if L.solved == L.total {
		return -1
	}

	pr := []int{}
	for i, s := range L.problems {
		if !s {
			pr = append(pr, i)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rd := r.Intn(len(pr))
	return pr[rd]
}

func (L *Leetcode) Solved() int {
	return L.solved
}

func (L *Leetcode) Total() int {
	return L.total
}

func (L *Leetcode) Count(i int) int {
	return L.count[i]
}

func (L *Leetcode) IsSolved(i int) bool {
	return i >= 0 && i <= L.total && L.problems[i]
}

func (L *Leetcode) IsValid(i int) bool {
	return i >= 0 && i <= L.total && !L.problems[i]
}
