package medium

import "fmt"

// Reference: https://leetcode.com/problems/partition-labels/
func init() {
	Solutions[763] = func() {
		fmt.Println(`Input: s = "ababcbacadefegdehijhklij"`)
		fmt.Println(`Output:`, partitionLabels(`ababcbacadefegdehijhklij`))
		fmt.Println(`Input: s = "eccbbbbdec"`)
		fmt.Println(`Output:`, partitionLabels(`eccbbbbdec`))
		fmt.Println(`Input: s = "caedbdedda"`)
		fmt.Println(`Output:`, partitionLabels(`caedbdedda`))
	}
}

func partitionLabels(s string) []int {
	m := make([]int, 26)
	for i, ch := range s {
		m[int(ch-'a')] = i
	}

	res := []int{}
	for i := 0; i < len(s); {
		next := m[int(s[i]-'a')]
		for j := i + 1; j <= next; j++ {
			next = max(next, m[int(s[j]-'a')])
		}

		res = append(res, next-i+1)
		i = next + 1
	}
	return res
}
