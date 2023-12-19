package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-window-substring/
func init() {
	Solutions[76] = func() {
		fmt.Println(`Input: s = "ADOBECODEBANC", t = "ABC"`)
		fmt.Println(`Output:`, minWindow(`ADOBECODEBANC`, `ABC`))
		fmt.Println(`Input: s = "cabefgecdaecf", t = "cae"`)
		fmt.Println(`Output:`, minWindow(`cabefgecdaecf`, `cae`))
	}
}

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ``
	}

	mt := [58]int{}
	for _, ch := range t {
		mt[int(ch-'A')]++
	}

	res := ``
	left, right := 0, -1
	count := 0
	ms := [58]int{}
	for left < m {
		for left < m-1 && mt[int(s[left]-'A')] == 0 {
			left++
		}

		flag := count != n
		if flag {
			right++
		}

		for right < m {
			cright := int(s[right] - 'A')
			if mt[cright] > 0 {
				if flag {
					ms[cright]++
					if mt[cright] >= ms[cright] {
						count++
					}
				}

				if count == n {
					if len(res) == 0 || len(res) > right-left+1 {
						res = s[left : right+1]
					}
					break
				}
			}
			right++
		}

		cleft := int(s[left] - 'A')
		if mt[cleft] > 0 {
			ms[cleft]--
			if mt[cleft] > ms[cleft] {
				count--
			}
		}
		left++
	}
	return res
}
