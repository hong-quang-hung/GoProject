package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
func init() {
	Solutions[1647] = func() {
		fmt.Println("Input: s = 'aaabbbcc'")
		fmt.Println("Output:", minDeletions("aaabbbcc"))
		fmt.Println("Input: s = 'ceabaacb'")
		fmt.Println("Output:", minDeletions("ceabaacb"))
		fmt.Println("Input: s = 'bbcebab'")
		fmt.Println("Output:", minDeletions("bbcebab"))
	}
}

func minDeletions(s string) int {
	ch := make([]int, 26)
	for _, r := range s {
		ch[int(r-'a')]++
	}

	m := make(map[int]bool)
	h := new(MaxHeap)
	for i := 0; i < 26; i++ {
		if ch[i] != 0 {
			heap.Push(h, ch[i])
			m[ch[i]] = true
		}
	}

	res := 0
	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := h.Peek()
		if first == second {
			i := second - 1
			for i > 0 && m[i] {
				i--
			}

			res += second - i
			m[i] = true
		}
	}
	return res
}
