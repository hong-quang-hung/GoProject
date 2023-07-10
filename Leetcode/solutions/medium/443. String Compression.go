package medium

import (
	"fmt"
	"strconv"
)

func init() {
	Solutions[443] = Leetcode_Compress
}

// Reference: https://leetcode.com/problems/string-compression/
func Leetcode_Compress() {
	fmt.Println("Input: chars = ['a','a','b','b','c','c','c']")
	fmt.Println("Output:", compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println("Input: chars = ['a','b','b','b','b','b','b','b','b','b','b','b','b']")
	fmt.Println("Output:", compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}

func compress(chars []byte) int {
	i, j, n := 0, 0, len(chars)
	for i < n {
		count := 1
		chars[j] = chars[i]
		for i < n-1 && chars[i] == chars[i+1] {
			i++
			count++
		}

		if count > 1 {
			countStr := strconv.Itoa(count)
			for item := range countStr {
				j++
				chars[j] = countStr[item]
			}
		}
		j++
		i++
	}
	return j
}
