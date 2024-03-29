package easy

import "fmt"

// Reference: https://leetcode.com/problems/verifying-an-alien-dictionary/
func init() {
	Solutions[953] = func() {
		fmt.Println(`Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"`)
		fmt.Println(`Output:`, isAlienSorted([]string{`hello`, `leetcode`}, `hlabcdefgijkmnopqrstuvwxyz`))
	}
}

func isAlienSorted(words []string, order string) bool {
	hash := make(map[byte]int)
	for i := 0; i < 26; i++ {
		hash[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				if hash[words[i][j]] > hash[words[i+1][j]] {
					return false
				} else {
					break
				}
			}
		}
	}
	return true
}
