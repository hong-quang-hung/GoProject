package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-smallest-letter-greater-than-target/
func init() {
	Solutions[744] = func() {
		fmt.Println(`Input: letters = ["c","f","j"], target = "a"`)
		fmt.Println(`Output:`, nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
		fmt.Println(`Input: letters = ["c","f","j"], target = "c"`)
		fmt.Println(`Output:`, nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
		fmt.Println(`Input: letters = ["x","x","y","y"], target = "z"`)
		fmt.Println(`Output:`, nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))
	}
}

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	index := sort.Search(n, func(i int) bool { return letters[i] > target })
	if n == index {
		return letters[0]
	}
	return letters[index]
}
