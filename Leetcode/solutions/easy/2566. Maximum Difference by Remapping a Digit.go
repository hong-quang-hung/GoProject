package easy

import (
	"fmt"
	"strconv"
	"strings"
)

// Reference: https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
func init() {
	Solutions[2566] = func() {
		fmt.Println("Input: num = 11891")
		fmt.Println("Output:", minMaxDifference(11891))
		fmt.Println("Input: num = 90")
		fmt.Println("Output:", minMaxDifference(90))
	}
}

func minMaxDifference(num int) int {
	s, h := strconv.Itoa(num), 0
	for l1 := len(s) - 1; h < l1 && s[h] == '9'; {
		h++
	}

	a, _ := strconv.Atoi(strings.ReplaceAll(s, string(s[h]), "9"))
	b, _ := strconv.Atoi(strings.ReplaceAll(s, string(s[0]), "0"))
	return a - b
}
