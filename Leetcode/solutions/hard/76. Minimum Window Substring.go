package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-window-substring/
func init() {
	Solutions[76] = func() {
		fmt.Println("Input: s = 'ADOBECODEBANC', t = 'ABC'")
		fmt.Println("Output:", minWindow("ADOBECODEBANC", "ABC"))
		fmt.Println("Input: s = 'a', t = 'a'")
		fmt.Println("Output:", minWindow("a", "a"))
		fmt.Println("Input: s = 'a', t = 'b'")
		fmt.Println("Output:", minWindow("a", "b"))
	}
}

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	mt := [58]int{}
	for _, ch := range t {
		mt[int(ch-'A')]++
	}

	res := ""
	for left := 0; left < m; left++ {
		for left < m && mt[int(s[left]-'A')] == 0 {
			left++
		}

		counter := 0
		ms := [58]int{}
		for right := left; right < m; right++ {
			ch := int(s[right] - 'A')
			if mt[ch] > 0 && mt[ch] > ms[ch] {
				ms[ch]++
				counter++
			}

			if counter == n {
				if len(res) == 0 || len(res) > right-left+1 {
					res = s[left : right+1]
				}
			}
		}
	}
	return res
}
