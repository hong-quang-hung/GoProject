package medium

import "fmt"

// Reference: https://leetcode.com/problems/palindrome-partitioning/
func init() {
	Solutions[131] = func() {
		fmt.Println(`Input: s = "aab"`)
		fmt.Println(`Output:`, partitionII(`aab`))
		fmt.Println(`Input: s = "a"`)
		fmt.Println(`Output:`, partitionII(`a`))
		fmt.Println(`Input: s = "aabb"`)
		fmt.Println(`Output:`, partitionII(`aabb`))
	}
}

func partitionII(s string) [][]string {
	res := [][]string{}
	partitionIISolved(s, &res, []string{})
	return res
}

func partitionIISolved(s string, res *[][]string, current []string) {
	n := len(s)
	if n == 0 {
		temp := make([]string, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}

	for i := 1; i <= n; i++ {
		curr := s[:i]
		if isPalindrome(curr) {
			current = append(current, curr)
			partitionIISolved(s[i:], res, current)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	if n == 1 {
		return true
	}

	for i := 0; i < n>>1; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
