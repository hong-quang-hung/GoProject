package medium

import "fmt"

// Reference: https://leetcode.com/problems/optimal-partition-of-string/
func Leetcode_Partition_String() {
	fmt.Println("Input: s = 'gizfdfri'")
	fmt.Println("Output:", partitionString("gizfdfri"))
	fmt.Println("Input: s = 'hdklqkcssgxlvehva'")
	fmt.Println("Output:", partitionString("hdklqkcssgxlvehva"))
}

func partitionString(s string) int {
	i, n, res := 0, len(s), 0
	for i < n {
		j, m := i, make([]int, 26)
		for ; j < n; j++ {
			m[s[j]-'a']++
			if m[s[j]-'a'] == 2 {
				break
			}
		}
		res++
		i = j
	}
	return res
}
