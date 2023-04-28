package hard

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/similar-string-groups/
func Leetcode_Num_Similar_Groups() {
	fmt.Println("Input: strs = ['tars','rats','arts','star']")
	fmt.Println("Output:", numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
	fmt.Println("Input: strs = ['omv','ovm']")
	fmt.Println("Output:", numSimilarGroups([]string{"omv", "ovm"}))
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	union, res := types.NewUnionFind(n), n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if numSimilarCheck(strs[i], strs[j]) && union.Find(i) != union.Find(j) {
				res--
				union.UnionSet(i, j)
			}
		}
	}
	return res
}

func numSimilarCheck(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 2 || diff == 0
}
