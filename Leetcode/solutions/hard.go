package solutions

import (
	"container/heap"
	"math"
	"sort"
)

// Reference: https://leetcode.com/problems/minimize-deviation-in-array/
type minimumDeviationHeap []int

func (h minimumDeviationHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h minimumDeviationHeap) Len() int           { return len(h) }
func (h minimumDeviationHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minimumDeviationHeap) Empty() bool        { return len(h) == 0 }
func (h *minimumDeviationHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *minimumDeviationHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func minimumDeviation(nums []int) int {
	set := new(minimumDeviationHeap)
	minDev := math.MaxInt32
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		minDev = min(minDev, num)
		heap.Push(set, num)
	}

	res := math.MaxInt32
	for !set.Empty() {
		pop := heap.Pop(set).(int)
		res = min(res, pop-minDev)
		if pop%2 == 0 {
			pop /= 2
			minDev = min(minDev, pop)
			heap.Push(set, pop)
		} else {
			break
		}
	}
	return res
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
		maxChunk := a
		for len(monotonic) > 0 && monotonic[len(monotonic)-1] > a {
			pop := monotonic[len(monotonic)-1]
			monotonic = monotonic[:len(monotonic)-1]
			maxChunk = max(maxChunk, pop)
		}
		monotonic = append(monotonic, maxChunk)
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
	n := len(nums)
	dp := make([]int, n+1)
	norms := minCostSolved(nums)
	for i := 1; i <= n; i++ {
		ones := 0
		count := make([]int, n)
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0; j-- {
			count[norms[j]]++
			if count[norms[j]] == 1 {
				ones++
			} else if count[norms[j]] == 2 {
				ones--
			}
			dp[i] = min(dp[i], dp[j]+k+i-j-ones)
		}
	}
	return dp[n]
}

func minCostSolved(nums []int) []int {
	norms := make([]int, len(nums))
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, c := m[nums[i]]; c {
			norms[i] = j
		} else {
			norms[i] = len(m)
			m[nums[i]] = norms[i]
		}
	}
	return norms
}

// Reference: https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/
func ways(pizza []string, k int) int {
	rows, cols := len(pizza), len(pizza[0])
	apples := make([][]int, rows+1)
	f := make([][]int, rows)

	for row := 0; row < rows; row++ {
		apples[row] = make([]int, cols+1)
		f[row] = make([]int, cols)
	}
	apples[rows] = make([]int, cols+1)

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if pizza[row][col] == 'A' {
				apples[row][col] = 1
			}
			apples[row][col] += apples[row+1][col] + apples[row][col+1] - apples[row+1][col+1]
			if apples[row][col] > 0 {
				f[row][col] = 1
			}
		}
	}

	const mod = int(1e9 + 7)
	for remain := 1; remain < k; remain++ {
		g := make([][]int, rows)
		for row := 0; row < rows; row++ {
			g[row] = make([]int, cols)
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				for next_row := row + 1; next_row < rows; next_row++ {
					if apples[row][col]-apples[next_row][col] > 0 {
						g[row][col] += f[next_row][col]
						g[row][col] %= mod
					}
				}

				for next_col := col + 1; next_col < cols; next_col++ {
					if apples[row][col]-apples[row][next_col] > 0 {
						g[row][col] += f[row][next_col]
						g[row][col] %= mod
					}
				}
			}
		}
		copy(f, g)
	}
	return f[0][0]
}

// Reference: https://leetcode.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	res := 0
	st := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = []int{i}
			} else {
				res = max(res, i-st[len(st)-1])
			}
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/naming-a-company/
func distinctNames(ideas []string) int64 {
	m := make([]map[string]bool, 26)
	for _, idea := range ideas {
		if m[idea[0]-'a'] == nil {
			m[idea[0]-'a'] = map[string]bool{}
		}
		m[idea[0]-'a'][idea[1:]] = true
	}

	res := int64(0)
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			if m[i] == nil || m[j] == nil {
				continue
			}

			countExist := 0
			for ideaI := range m[i] {
				if m[j][ideaI] {
					countExist++
				}
			}
			res += 2 * int64(len(m[i])-countExist) * int64(len(m[j])-countExist)
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adj := make(map[int][]int)
	indegree := make([]int, n)
	for _, edge := range edges {
		if adj[edge[0]] == nil {
			adj[edge[0]] = []int{}
		}
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	count, visit, inStack, res := make([][]int, n), make([]bool, n), make([]bool, n), 0
	for i := 0; i < n; i++ {
		count[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		res = max(res, largestPathValueDFS(i, colors, adj, count, visit, inStack))
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func largestPathValueDFS(node int, colors string, adj map[int][]int, count [][]int, visit []bool, inStack []bool) int {
	if inStack[node] {
		return math.MaxInt
	}

	if visit[node] {
		return count[node][colors[node]-'a']
	}

	visit[node] = true
	inStack[node] = true

	if adj[node] != nil {
		for _, neighbor := range adj[node] {
			if largestPathValueDFS(neighbor, colors, adj, count, visit, inStack) == math.MaxInt {
				return math.MaxInt
			}
			for i := 0; i < 26; i++ {
				count[node][i] = max(count[node][i], count[neighbor][i])
			}
		}
	}

	count[node][colors[node]-'a']++
	inStack[node] = false
	return count[node][colors[node]-'a']
}
