package medium

import "fmt"

// Reference: https://leetcode.com/problems/group-anagrams/
func init() {
	Solutions[49] = func() {
		fmt.Println("Input: strs = ['eat','tea','tan','ate','nat','bat']")
		fmt.Println("Output:", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
		fmt.Println("Input: strs = ['']")
		fmt.Println("Output:", groupAnagrams([]string{""}))
		fmt.Println("Input: strs = ['a']")
		fmt.Println("Output:", groupAnagrams([]string{"a"}))
	}
}

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		k := [26]int{}
		for _, ch := range str {
			k[int(ch-'a')]++
		}
		m[k] = append(m[k], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
