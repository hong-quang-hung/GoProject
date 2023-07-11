package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
func init() {
	Solutions[2300] = func() {
		fmt.Println("Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7")
		fmt.Println("Output:", successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
		fmt.Println("Input: spells = [3,1,2], potions = [8,5,8], success = 16")
		fmt.Println("Output:", successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	}
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)

	m := len(potions)
	for i, spell := range spells {
		j := sort.SearchInts(potions, int((success+int64(spell)-1)/int64(spell)))
		res[i] = m - j
	}
	return res
}
