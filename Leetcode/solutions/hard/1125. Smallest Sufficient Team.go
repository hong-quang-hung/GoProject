package hard

import "fmt"

// Reference: https://leetcode.com/problems/put-marbles-in-bags/
func init() {
	Solutions[1125] = func() {
		fmt.Println("Input: req_skills = ['java','nodejs','reactjs'], people = [['java'],['nodejs'],['nodejs','reactjs']]")
		fmt.Println("Output:", smallestSufficientTeam([]string{"java", "nodejs", "reactjs"}, [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}))
		fmt.Println("Input: req_skills = ['algorithms','math','java','reactjs','csharp','aws'], people = [['algorithms','math','java'],['algorithms','math','reactjs'],['java','csharp','aws'],['reactjs','csharp'],['csharp','math'],['aws','java']]")
		fmt.Println("Output:", smallestSufficientTeam([]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}, [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}))
	}
}

type SufficientTeam struct {
	prevDP       uint16
	lastPersonId uint8
	peopleCount  uint8
}

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	skillNameToId := make(map[string]uint8, len(req_skills))
	for i, skill := range req_skills {
		skillNameToId[skill] = uint8(i)
	}

	dp := make([]SufficientTeam, 1<<len(req_skills))
	for i, skills := range people {
		if len(skills) == 0 {
			continue
		}

		var skillsBits uint16
		for _, skill := range skills {
			skillsBits |= 1 << skillNameToId[skill]
		}

		if dp[skillsBits].peopleCount == 1 {
			continue
		}

		dp[skillsBits] = SufficientTeam{
			lastPersonId: uint8(i),
			peopleCount:  1,
		}

		for j := range dp {
			if dp[j].peopleCount != 0 {
				or := uint16(j) | skillsBits
				if dp[or].peopleCount == 0 || dp[j].peopleCount+1 < dp[or].peopleCount {
					dp[or] = SufficientTeam{
						prevDP:       uint16(j),
						lastPersonId: uint8(i),
						peopleCount:  dp[j].peopleCount + 1,
					}
				}
			}
		}
	}

	res := []int{}
	for i := uint16(len(dp) - 1); i != 0; i = dp[i].prevDP {
		res = append(res, int(dp[i].lastPersonId))
	}
	return res
}
