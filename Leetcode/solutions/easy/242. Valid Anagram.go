package easy

import "fmt"

// Reference: https://leetcode.com/problems/valid-anagram/
func init() {
	Solutions[242] = func() {
		fmt.Println(`Input: s = "anagram", t = "nagaram"`)
		fmt.Println(`Output:`, isAnagram(`anagram`, `nagaram`))
		fmt.Println(`Input: s = "rat", t = "car"`)
		fmt.Println(`Output:`, isAnagram(`rat`, `car`))
	}
}

func isAnagram(s string, t string) bool {
	ms, mt := [26]int{}, [26]int{}
	for _, ch := range s {
		ms[int(ch-'a')]++
	}
	for _, ch := range t {
		mt[int(ch-'a')]++
	}
	return ms == mt
}
