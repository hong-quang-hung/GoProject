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

	var j int = i
	var mincost int = math.MaxInt16
	for c, cost := range costs {
		for (j < len(days)) && (days[j] < days[i]+dur[c]) {
			j++
		}
		dp := mincostTicketsDP(days, costs, dur, dp, j) + cost
		mincost = min(mincost, dp)
	}
	dp[i] = mincost
	return mincost
}

// Reference: https://leetcode.com/problems/removing-stars-from-a-string/
func removeStars(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

// Reference: https://leetcode.com/problems/max-chunks-to-make-sorted/
func maxChunksToSorted(arr []int) int {
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

// Reference: https://leetcode.com/problems/count-the-hidden-sequences/
func numberOfArrays(differences []int, lower int, upper int) int {
	prefix := make([]int, len(differences))
	prefix[0] = differences[0]
	maxP, minP := prefix[0], prefix[0]
	for i := 1; i < len(differences); i++ {
		prefix[i] = prefix[i-1] + differences[i]
		maxP = max(maxP, prefix[i])
		minP = min(minP, prefix[i])
	}

	res := 0
	for i := lower; i <= upper; i++ {
		if maxP+i >= lower && maxP+i <= upper && minP+i >= lower && minP+i <= upper {
			res++
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *types.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	nums, isNodeLeft := 1, false
	queue := []*types.TreeNode{root}
	nodes := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodes = append(nodes, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if len(nodes) == nums {
			if isNodeLeft {
				for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
					nodes[i], nodes[j] = nodes[j], nodes[i]
				}
			}
			res = append(res, nodes)
			nodes = nil
			nums = len(queue)
			isNodeLeft = !isNodeLeft
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/find-the-substring-with-maximum-cost/
func maximumCostSubstring(s string, chars string, vals []int) int {
	m := make(map[rune]int)
	for i, ch := range chars {
		m[ch] = vals[i]
	}

	maxCost, cost := 0, 0
	for _, ch := range s {
		var val int
		if v, c := m[ch]; c {
			val = v
		} else {
			val = int(ch-'a') + 1
		}
		cost = max(0, cost+val)
		maxCost = max(maxCost, cost)
	}
	return maxCost
}

// Reference: https://leetcode.com/problems/find-the-substring-with-maximum-cost/
func findMatrix(nums []int) [][]int {
	res := make([][]int, 1)
	m := make(map[int]int)
	for _, num := range nums {
		if count, c := m[num]; c {
			if len(res) < count+1 {
				res = append(res, []int{})
			}
			res[count] = append(res[count], num)
			m[num] = count + 1
		} else {
			res[0] = append(res[0], num)
			m[num] = 1
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)

	m := len(potions)
	for i, spell := range spells {
		j := sort.SearchInts(potions, int((success+int64(spell)-1)/int64(spell)))
		res[i] = m - j
	}
	return res
}

// Reference: https://leetcode.com/problems/jump-game-ii/
func jump_ii(nums []int) int {
	res, last, next, n := 0, 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		next = max(next, i+nums[i])
		if next >= n-1 {
			res++
			break
		}
		if i == last {
			res++
			last = next
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/boats-to-save-people/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j, res := 0, len(people)-1, 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		res++
	}
	return res
}

// Reference: https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
func countPairs(n int, edges [][]int) int64 {
	union := types.NewUnionFind(n)
	for _, edge := range edges {
		union.UnionSet(edge[0], edge[1])
	}

	group := make(map[int]int)
	for i := 0; i < n; i++ {
		root := union.Find(i)
		group[root]++
	}

	remainingNodes, numberOfPairs := int64(n), int64(0)
	for _, v := range group {
		numberOfPairs += int64(v) * (remainingNodes - int64(v))
		remainingNodes -= int64(v)
	}
	return numberOfPairs
}

// Reference: https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
func detectCycle(head *types.ListNode) *types.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

// Reference: https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
func minimumPerimeter(neededApples int64) int64 {
	res, x := int64(0), int64(0)
	for res < neededApples {
		x++
		res += 12 * x * x
	}
	return x * 8
}

// Reference: https://leetcode.com/problems/min-cost-to-connect-all-points/
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	m := [][]int{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			m = append(m, []int{i, j, manhattanDistance(points[i], points[j])})
		}
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i][2] < m[j][2]
	})

	p, i, union, distance := 0, 0, types.NewUnionFind(n), make([]int, n)
	for p < n-1 {
		nextPoint := m[i]
		i++
		x := union.Find(nextPoint[0])
		y := union.Find(nextPoint[1])
		if x != y {
			distance[p] = nextPoint[2]
			union.UnionSet(x, y)
			p++
		}
	}

	minimumCost := 0
	for j := 0; j < p; j++ {
		minimumCost += distance[j]
	}
	return minimumCost
}

// Reference: https://leetcode.com/problems/optimal-partition-of-string/
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

// Reference: https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
func kthLargestLevelSum(root *types.TreeNode, k int) int64 {
	sumk := make(map[int]int64)
	sumNodeRecursive(root, sumk, 1)

	n := len(sumk)
	if n < k {
		return -1
	}

	res := int64(math.MinInt64)
	for _, v := range sumk {
		if res < v {
			res = v
		}
	}
	return res
}

func sumNodeRecursive(root *types.TreeNode, sumK map[int]int64, k int) {
	if root == nil {
		return
	}

	sumK[k] += int64(root.Val)
	if root.Left != nil {
		sumNodeRecursive(root.Left, sumK, k+1)
	}
	if root.Right != nil {
		sumNodeRecursive(root.Right, sumK, k+1)
	}
}
