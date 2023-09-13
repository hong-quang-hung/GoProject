package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-window-substring/
func init() {
	Solutions[76] = func() {
		fmt.Println("Input: s = 'ADOBECODEBANC', t = 'ABC'")
		fmt.Println("Output:", minWindow("ADOBECODEBANC", "ABC"))
		fmt.Println("Input: s = 'a', t = 'a'")
		fmt.Println("Output:", minWindow("a", "a"))
		fmt.Println("Input: s = 'a', t = 'aa'")
		fmt.Println("Output:", minWindow("a", "aa"))
	}
}

func minWindow(s string, t string) string {
	check, ms, mt := [52]bool{}, [52]int{}, [52]int{}
	for _, ch := range t {
		check[int(ch-'A')] = true
		mt[int(ch-'A')]++
	}

	var res string
	m, n := len(s), len(t)
	for i := 0; i < m; i++ {
		if check[int(s[i]-'A')] {
			ms[s[i]-'A']++
			if i-n >= 0 {
				ms[s[i-n]-'A']--
			}
		}
	}
	return res
}
