package easy

import "fmt"

// Reference: https://leetcode.com/problems/isomorphic-strings/
func init() {
	Solutions[205] = func() {
		fmt.Println(`Input: s = 'egg', t = 'add'`)
		fmt.Println(`Output:`, isIsomorphic(`egg`, `add`))
		fmt.Println(`Input: s = 'foo', t = 'bar'`)
		fmt.Println(`Output:`, isIsomorphic(`foo`, `bar`))
		fmt.Println(`Input: s = 'paper', t = 'title'`)
		fmt.Println(`Output:`, isIsomorphic(`paper`, `title`))
	}
}

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := range s {
		if val, ok := m1[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			if _, ok := m2[t[i]]; ok {
				return false
			}
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		}
	}
	return true
}
