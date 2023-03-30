package solutions

import (
	"sort"
)

// Reference: https://leetcode.com/problems/minimize-deviation-in-array/
func minimumDeviation(nums []int) int {
	return 0
}

// Reference: https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	min, max, left := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			left = i
		}
		if nums[i] == maxK {
			max = i
		}
		if nums[i] == minK {
			min = i
		}
		if min == -1 || max == -1 {
			continue
		}
		var minLeft int
		if min > max {
			minLeft = max
		} else {
			minLeft = min
		}
		if minLeft > left {
			res += int64(minLeft - left)
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	res := make([]int, (n+m)/2+1)
	even := (n+m)&1 == 0

	l, r := 0, 0
	for i := 0; i < len(res); i++ {
		if l == len(nums1) && r < len(nums2) {
			res[i] = nums2[r]
			r++
			continue
		} else if r == len(nums2) && l < len(nums1) {
			res[i] = nums1[l]
			l++
			continue
		}

		if nums1[l] <= nums2[r] {
			res[i] = nums1[l]
			l++
		} else if r < len(nums2) {
			res[i] = nums2[r]
			r++
		}

	}

	if even && n+m > 1 {
		return float64(res[len(res)-1]+res[len(res)-2]) / 2
	} else {
		return float64(res[len(res)-1])
	}
}

// Reference: https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
func longestCycle(edges []int) int {
	n := len(edges)
	visit := make([]bool, n)
	indegree := make([]int, n)

	for _, edge := range edges {
		if edge != -1 {
			indegree[edge]++
		}
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		var node int = q[0]
		q = q[1:]
		visit[node] = true
		var neighbor int = edges[node]
		if neighbor != -1 {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	var answer int = -1
	for i := 0; i < n; i++ {
		if !visit[i] {
			var neighbor int = edges[i]
			var count int = 1
			visit[i] = true
			for neighbor != i {
				visit[neighbor] = true
				count++
				neighbor = edges[neighbor]
			}
			if answer < count {
				answer = count
			}
		}
	}
	return answer
}

// Reference: https://leetcode.com/problems/reducing-dishes/
func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool { return satisfaction[i] < satisfaction[j] })
	n, max, total := len(satisfaction), 0, 0

	for i := n - 1; i >= 0 && satisfaction[i]+total > 0; i-- {
		total += satisfaction[i]
		max += total
	}
	return max
}

// Reference: https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
func maxChunksToSorted_ii(arr []int) int {
	monotonic := []int{}
	for _, a := range arr {
		max := a
		for len(monotonic) > 0 && monotonic[len(monotonic)-1] > a {
			pop := monotonic[len(monotonic)-1]
			monotonic = monotonic[:len(monotonic)-1]
			if pop > max {
				max = pop
			}
		}
		monotonic = append(monotonic, max)
	}
	return len(monotonic)
}

// Reference: https://leetcode.com/problems/number-of-visible-people-in-a-queue/
func canSeePersonsCount(heights []int) []int {
	monotonic := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		count, height := 0, heights[i]
		for len(monotonic) > 0 && monotonic[len(monotonic)-1] < height {
			monotonic = monotonic[:len(monotonic)-1]
			count++
		}
		if len(monotonic) > 0 {
			count++
		}
		heights[i] = count
		monotonic = append(monotonic, height)
	}
	return heights
}

// Reference: https://leetcode.com/problems/scramble-string/
func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]bool, n+1)
	dp[1] = make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[1][i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}
	for l := 2; l <= n; l++ {
		if dp[l] == nil {
			dp[l] = make([][]bool, n)
		}
		for i := 0; i < n+1-l; i++ {
			if dp[l][i] == nil {
				dp[l][i] = make([]bool, n)
			}
			for j := 0; j < n+1-l; j++ {
				for k := 1; k < l; k++ {
					dp1 := dp[k][i]
					dp2 := dp[l-k][i+k]

					dp[l][i][j] = dp[l][i][j] || (dp1[j] && dp2[j+k])
					dp[l][i][j] = dp[l][i][j] || (dp1[j+l-k] && dp2[j])
				}
			}
		}
	}
	return dp[n][0][0]
}

// Reference: https://leetcode.com/problems/minimum-cost-to-split-an-array/
func minCost(nums []int, k int) int {
	return 0
}
