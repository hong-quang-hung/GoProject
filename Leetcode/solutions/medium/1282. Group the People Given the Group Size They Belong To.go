package medium

import "fmt"

// Reference: https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
func init() {
	Solutions[1282] = func() {
		fmt.Println(`Input: groupSizes = [3,3,3,3,3,1,3]`)
		fmt.Println(`Output:`, groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
		fmt.Println(`Input: groupSizes = [2,1,3,3,3,2]`)
		fmt.Println(`Output:`, groupThePeople([]int{2, 1, 3, 3, 3, 2}))
	}
}

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	for i, v := range groupSizes {
		m[v] = append(m[v], i)
	}

	res := [][]int{}
	for k, v := range m {
		for i := 0; i < len(v)/k; i++ {
			arr := []int{}
			for j := 0; j < k; j++ {
				arr = append(arr, v[i*k+j])
			}
			res = append(res, arr)
		}
	}
	return res
}
