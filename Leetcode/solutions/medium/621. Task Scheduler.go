package medium

import "fmt"

// Reference: https://leetcode.com/problems/task-scheduler/
func init() {
	Solutions[621] = func() {
		fmt.Println(`Input: tasks = ["A","A","A","B","B","B"], n = 2`)
		fmt.Println(`Output:`, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
		fmt.Println(`Input: tasks = ["A","C","A","B","D","B"], n = 1`)
		fmt.Println(`Output:`, leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
		fmt.Println(`Input: tasks = ["A","A","A", "B","B","B"], n = 3`)
		fmt.Println(`Output:`, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 1))
	}
}

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	maxCount := 0
	for _, ch := range tasks {
		k := int(ch - 'A')
		m[k]++
		maxCount = max(maxCount, m[k])
	}

	time := (maxCount - 1) * (n + 1)
	for _, f := range m {
		if f == maxCount {
			time++
		}
	}

	return max(len(tasks), time)
}
