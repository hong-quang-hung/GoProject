package medium

import (
	"fmt"
	"math/rand"
	"time"
)

// Reference: https://leetcode.com/problems/linked-list-random-node/
func init() {
	Solutions[382] = func() {
		fmt.Println("['Solution', 'getRandom', 'getRandom', 'getRandom', 'getRandom', 'getRandom']")
		fmt.Println("[[[1, 2, 3]], [], [], [], [], []]")

		solutions := SolutionConstructor(S2ListNode("[1, 2, 3]"))
		fmt.Println(solutions.GetRandom())
		fmt.Println(solutions.GetRandom())
		fmt.Println(solutions.GetRandom())
		fmt.Println(solutions.GetRandom())
		fmt.Println(solutions.GetRandom())
	}
}

type Random struct {
	list []int
}

func SolutionConstructor(head *ListNode) Random {
	rand.NewSource(time.Now().UnixNano())
	solutions := Random{}
	for head != nil {
		solutions.Append(head.Val)
		head = head.Next
	}
	return solutions
}

func (s *Random) GetRandom() int {
	return s.list[rand.Intn(len(s.list))]
}

func (s *Random) Append(val int) {
	s.list = append(s.list, val)
}
