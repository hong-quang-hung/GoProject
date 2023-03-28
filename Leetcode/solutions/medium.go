package solutions

import (
	"math"
	"sort"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	var max int = -1

	n, max := len(grid)-1, -1
	queue := [][2]int{}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		land := queue[0]
		queue = queue[1:]
		point := grid[land[0]][land[1]]
		col, row := land[0], land[1]

		if col > 0 && grid[col-1][row] == 0 {
			grid[col-1][row] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col - 1, row})
		}

		if col < n && grid[col+1][row] == 0 {
			grid[col+1][row] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col + 1, row})
		}

		if row > 0 && grid[col][row-1] == 0 {
			grid[col][row-1] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col, row - 1})
		}

		if row < n && grid[col][row+1] == 0 {
			grid[col][row+1] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col, row + 1})
		}
	}
	return max
}

// Reference: https://leetcode.com/problems/sort-an-array/
func sortArray(nums []int) []int {
	utils.CountingSort(nums)
	return nums
}

// Reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	chars := [128]int{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		char := s[right]
		index := chars[char] - 1
		if (index != 0 || char == s[0]) && index >= left && index < right {
			left = index + 1
		}
		if res < (right - left + 1) {
			res = right - left + 1
		}
		chars[char] = right + 1
		right++
	}
	return res
}

// Reference: https://leetcode.com/problems/count-total-number-of-colored-cells/
func coloredCells(n int) int64 {
	return int64(2*n*n - 2*n + 1)
}

// Reference: https://leetcode.com/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) int {
	var left int = 1
	var right int = 0
	for _, p := range piles {
		if right < p {
			right = p
		}
	}
	var res int = right
	for left <= right {
		var hours int64 = 0
		var mid = left + (right-left)/2
		for _, p := range piles {
			hours += int64(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= int64(h) {
			if res > mid {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/rearrange-array-to-maximize-prefix-score/
func maxScore(nums []int) int {
	var res int = 0
	var prefix int64 = 0
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, num := range nums {
		prefix += int64(num)
		if prefix > 0 {
			res++
		} else {
			break
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/
func beautifulSubarrays(nums []int) int64 {
	var res int64 = 0
	var xor int64 = 0

	m := make(map[int64]int)
	m[0] = 1
	for _, num := range nums {
		xor ^= int64(num)
		x, v := m[xor]
		if v {
			res += int64(x)
		}
		m[xor] = x + 1
	}
	return res
}

// Reference: https://leetcode.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if root.Left != nil {
		root.Left.Val += 10 * root.Val
	}
	if root.Right != nil {
		root.Right.Val += 10 * root.Val
	}
	return sumNumbers(root.Left) + sumNumbers(root.Right)
}

// Reference: https://leetcode.com/problems/check-completeness-of-a-binary-tree/
func isCompleteTree(root *types.TreeNode) bool {
	queue := make([]*types.TreeNode, 0)
	queue = append(queue, root)
	check := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			check = true
		} else {
			if check {
				return false
			}
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}

// Reference: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *types.TreeNode {
	return buildTreeConstruct(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func buildTreeConstruct(inorder []int, postorder []int, i1 int, i2 int, p1 int, p2 int) *types.TreeNode {
	mid, rootVal := i1, postorder[p2]
	for mid <= i2 && inorder[mid] != rootVal {
		mid++
	}
	root := &types.TreeNode{Val: rootVal}
	if mid-i1 == 1 {
		root.Left = &types.TreeNode{Val: inorder[i1]}
	} else if mid-i1 > 1 {
		root.Left = buildTreeConstruct(inorder, postorder, i1, mid-1, p1, p1+mid-i1-1)
	}
	if i2-mid == 1 {
		root.Right = &types.TreeNode{Val: postorder[p2-1]}
	} else if i2-mid > 1 {
		root.Right = buildTreeConstruct(inorder, postorder, mid+1, i2, p1+mid-i1, p2-1)
	}
	return root
}

// Reference: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func maximizeGreatness(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res++
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/the-number-of-beautiful-subsets/
func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	return counter(0, map[int]int{}, nums, k) - 1
}

func counter(index int, m map[int]int, nums []int, k int) int {
	if index == len(nums) {
		return 1
	}
	count := counter(index+1, m, nums, k)
	if m[nums[index]-k] == 0 {
		m[nums[index]]++
		count += counter(index+1, m, nums, k)
		m[nums[index]]--
	}
	return count
}

// Reference: https://leetcode.com/problems/number-of-zero-filled-subarrays/
func zeroFilledSubarray(nums []int) int64 {
	var res int64 = 0
	var size int64 = 0
	for _, num := range nums {
		if num == 0 {
			size++
		} else {
			res += (size * (size + 1) / 2)
			size = 0
		}
	}
	if size == 0 {
		return res
	} else {
		return res + (size * (size + 1) / 2)
	}
}

// Reference: https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	union := types.NewUnionFind(n)
	var numberOfConnectedComponents int = n
	for _, connection := range connections {
		if union.Find(connection[0]) != union.Find(connection[1]) {
			numberOfConnectedComponents--
			union.UnionSet(connection[0], connection[1])
		}
	}
	return numberOfConnectedComponents - 1
}

// Reference: https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
func minScore(n int, roads [][]int) int {
	union := types.NewUnionFind(n + 1)
	var answer int = math.MaxInt16

	for _, road := range roads {
		union.UnionSet(road[0], road[1])
	}

	for _, road := range roads {
		if union.Find(1) == union.Find(road[0]) {
			if answer > road[2] {
				answer = road[2]
			}
		}
	}
	return answer
}

// Reference: https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
func minOperations(nums []int, queries []int) []int64 {
	var n int = len(nums)
	res := make([]int64, len(queries))
	sumIndex := make([]int64, n)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sumIndex[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		sumIndex[i] = sumIndex[i-1] + int64(nums[i])
	}

	for j, q := range queries {
		if q <= nums[0] || q > nums[n-1] {
			res[j] = int64(math.Abs(float64(int64(n)*int64(q) - int64(sumIndex[n-1]))))
		} else {
			index := findIndexMinOperations(nums, q)
			res[j] = int64(q)*int64(index) - 2*sumIndex[index-1] + sumIndex[n-1] - int64(q)*int64(n-index)
		}
	}

	return res
}

func findIndexMinOperations(nums []int, find int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < find {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Reference: https://leetcode.com/problems/prime-subtraction-operation/description/
func primeSubOperation(nums []int) bool {
	isPrime := utils.Make(1005, true)
	prime := make([]int, 0)
	seive(isPrime, 1005)

	for i := 0; i < 1005; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}

	var n int = len(nums)
	var flag bool = false
	var prev int = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < prev {
			prev = nums[i]
			continue
		}

		flag = true
		for sub := 0; sub < len(prime) && prime[sub] < nums[i]; sub++ {
			if nums[i]-prime[sub] < prev {
				prev = nums[i] - prime[sub]
				flag = false
				break
			}
		}
		if flag {
			break
		}
	}
	return !flag
}

func seive(v []bool, n int) {
	v[0] = false
	v[1] = false
	for i := 2; i*i < n; i++ {
		for j := 2 * i; j < n; j += i {
			v[j] = false
		}
	}
}

// Reference: https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	var start int = 0
	var end int = 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		var len int
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

// Reference: https://leetcode.com/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r)

	for i := range dp {
		dp[i] = make([]int, c)
	}

	dp[0][0] += grid[0][0]
	for i := 1; i < c; i++ {
		dp[0][i] += dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < r; i++ {
		dp[i][0] += dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[r-1][c-1]
}

// Reference: https://leetcode.com/problems/minimum-cost-for-tickets/
func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days))
	dur := []int{1, 7, 30}
	return mincostTicketsDP(days, costs, dur, dp, 0)
}

func mincostTicketsDP(days []int, costs []int, dur []int, dp []int, i int) int {
	if i >= len(dp) {
		return 0
	}

	if dp[i] != 0 {
		return dp[i]
	}

	var min int = math.MaxInt16
	var j int = i
	for c, cost := range costs {
		for (j < len(days)) && (days[j] < days[i]+dur[c]) {
			j++
		}
		dp := mincostTicketsDP(days, costs, dur, dp, j) + cost
		if min > dp {
			min = dp
		}
	}
	dp[i] = min
	return min
}
