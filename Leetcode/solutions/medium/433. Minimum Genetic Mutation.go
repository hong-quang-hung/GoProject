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
	b := make(map[string]interface{})
	for _, gen := range bank {
		b[gen] = struct{}{}
	}

	res := 0
	choices := []byte{'A', 'C', 'G', 'T'}
	queue := []string{startGene}
	visited := make(map[string]interface{})
	visited[startGene] = struct{}{}
	for len(queue) > 0 {
		nextGens := []string{}
		for _, q := range queue {
			if q == endGene {
				return res
			}

			for _, choice := range choices {
				for i := 0; i < 8; i++ {
					nextGen := q[:i] + string(choice) + q[i+1:]
					if _, isVisited := visited[nextGen]; !isVisited {
						if _, ok := b[nextGen]; ok {
							visited[nextGen] = struct{}{}
							nextGens = append(nextGens, nextGen)
						}
					}
				}
			}
		}
		queue = nextGens
		nextGens = nil
		res++
	}
	return -1
}
