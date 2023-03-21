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

// Reference: https://leetcode.com/problems/maximize-greatness-of-an-array
func beautifulSubsets(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	count := make([](map[int]int), k)
	for i := 0; i < k; i++ {
		count[i] = map[int]int{}
	}
	for _, num := range nums {
		if v, c := count[num%k][num]; c {
			count[num%k][num] = v + 1
		} else {
			count[num%k][num] = 1
		}
	}

	var res int = 1
	for j := 0; j < k; j++ {
		prev, dp0, dp1 := 0, 1, 0
		s := count[j]
		for s1, s2 := range s {
			var temp int = dp0
			var pow int = int(math.Pow(2, float64(s2)))
			dp0 += dp1
			if prev+k == s1 {
				dp1 = temp * (pow - 1)
			} else {
				dp1 = (temp + dp1) * (pow - 1)
			}
			prev = s1
		}
		res = res * (dp0 + dp1)
	}
	return res - 1
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
