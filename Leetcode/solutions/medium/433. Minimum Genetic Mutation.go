package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-genetic-mutation/
func init() {
	Solutions[433] = func() {
		fmt.Println(`Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]`)
		fmt.Println(`Output:`, minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
		fmt.Println(`Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]`)
		fmt.Println(`Output:`, minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	}
}

func minMutation(startGene string, endGene string, bank []string) int {
	return 0
}
