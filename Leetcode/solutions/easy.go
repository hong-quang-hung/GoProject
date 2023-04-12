package solutions

import (
	"math"
	"sort"
	"strconv"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/split-with-minimum-sum/
func splitNum(num int) int {
	arr := make([]int, len(strconv.Itoa(num)))
	i, numA, numB := 0, 0, 0
	for num > 0 {
		arr[i] = num % 10
		num = num / 10
		i++
	}

	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			numA = numA*10 + arr[i]
		} else {
			numB = numB*10 + arr[i]
		}
	}
	return numA + numB
}

// Reference: https://leetcode.com/problems/left-and-right-sum-differences/
func leftRigthDifference(nums []int) []int {
	leftSum, rightSum := 0, 0
	for _, num := range nums {
		rightSum += num
	}

	ans := make([]int, len(nums))
	for i, num := range nums {
		rightSum -= num
		ans[i] = int(math.Abs(float64(leftSum - rightSum)))
		leftSum += num
	}
	return ans
}

// Reference: https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *types.TreeNode) *types.TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

// Reference: https://leetcode.com/problems/minimum-distance-between-bst-nodes
func minDiffInBST(root *types.TreeNode) int {
	minDiff, preValue := math.MaxInt, -1
	inorderTraversal(root, &preValue, &minDiff)
	return minDiff
}

func inorderTraversal(root *types.TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, preValue, minDiff)

	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val
	inorderTraversal(root.Right, preValue, minDiff)
}

// Reference: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeAdjacentDuplicates(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

// Reference: https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == size-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			i++
		}
	}
	return n <= 0
}

// Reference: https://leetcode.com/problems/add-to-array-form-of-integer/
func addToArrayForm(num []int, k int) []int {
	var ext int
	n := len(num)
	i := n - 1
	for k > 0 {
		if i >= 0 {
			ext = (num[i] + k%10) / 10
			num[i] = (num[i] + k%10) - ext*10
		} else {
			ext = 0
			num = append([]int{k % 10}, num...)
		}
		k = k/10 + ext
		i--
	}
	return num
}

// Reference: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int][]int)
	for _, num := range nums1 {
		m[num[0]] = num
	}

	for _, num := range nums2 {
		if m[num[0]] == nil {
			m[num[0]] = []int{num[0], num[1]}
		} else {
			m[num[0]][1] += num[1]
		}
	}

	res := make([][]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	return res
}
