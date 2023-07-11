package hard

import "fmt"

// Reference: https://leetcode.com/problems/naming-a-company/
func init() {
	Solutions[2306] = func() {
		fmt.Println("Input: ideas = ['coffee','donuts','time','toffee']")
		fmt.Println("Output:", distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
		fmt.Println("Input: ideas = ['aaa','baa', 'caa', 'bbb', 'cbb', 'dbb']")
		fmt.Println("Output:", distinctNames([]string{"aaa", "baa", "caa", "bbb", "cbb", "dbb"}))
	}
}

func distinctNames(ideas []string) int64 {
	m := make([]map[string]bool, 26)
	for _, idea := range ideas {
		if m[idea[0]-'a'] == nil {
			m[idea[0]-'a'] = map[string]bool{}
		}
		m[idea[0]-'a'][idea[1:]] = true
	}

	res := int64(0)
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			if m[i] == nil || m[j] == nil {
				continue
			}

			countExist := 0
			for ideaI := range m[i] {
				if m[j][ideaI] {
					countExist++
				}
			}
			res += 2 * int64(len(m[i])-countExist) * int64(len(m[j])-countExist)
		}
	}
	return res
}
