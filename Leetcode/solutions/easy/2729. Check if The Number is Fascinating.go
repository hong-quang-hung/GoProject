package easy

import "fmt"

// Reference: https://leetcode.com/problems/check-if-the-number-is-fascinating/
func init() {
	Solutions[2729] = func() {
		fmt.Println(`Input: n = 192`)
		fmt.Println(`Output:`, isFascinating(192))
		fmt.Println(`Input: n = 783`)
		fmt.Println(`Output:`, isFascinating(783))
		fmt.Println(`Input: n = 267`)
		fmt.Println(`Output:`, isFascinating(267))
	}
}

func isFascinating(n int) bool {
	m := make(map[int]bool)
	n2 := n * 2
	n3 := n * 3

	fmt.Println(n, n2, n3)
	for n > 0 {
		v := n % 10
		if v != 0 && m[v] {
			return false
		}
		m[v] = true
		n /= 10
	}

	for n2 > 0 {
		v := n2 % 10
		if v != 0 && m[v] {
			return false
		}
		m[v] = true
		n2 /= 10
	}

	for n3 > 0 {
		v := n3 % 10
		if v != 0 && m[v] {
			return false
		}
		m[v] = true
		n3 /= 10
	}

	delete(m, 0)
	return len(m) == 9
}
