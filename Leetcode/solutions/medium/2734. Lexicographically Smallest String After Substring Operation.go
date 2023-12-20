package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/lexicographically-smallest-string-after-substring-operation/
func init() {
	Solutions[2734] = func() {
		fmt.Println(`Input: s = "cbabc"`)
		fmt.Println(`Output:`, smallestString(`cbabc`))
		fmt.Println(`Input: s = "aacbbc"`)
		fmt.Println(`Output:`, smallestString(`acbbc`))
	}
}

func smallestString(s string) string {
	spl := strings.Split(s, "a")
	index := 0
	for index < len(spl) && len(spl[index]) == 0 {
		index++
	}

	if index == len(spl) {
		spl[index-1] = "z"
		spl = spl[1:]
		return strings.Join(spl, `a`)
	}

	res := make([]byte, len(spl[index]))
	for i := range spl[index] {
		res[i] = spl[index][i] - 1
	}
	spl[index] = string(res)
	return strings.Join(spl, "a")
}
