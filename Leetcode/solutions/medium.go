package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"leetcode.com/Leetcode/types"
)

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

// Reference: https://leetcode.com/problems/find-the-divisibility-array-of-a-string/
func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	divisible := 0
	for i := 0; i < len(word); i++ {
		divisible = (divisible*10 + int(word[i]-'0')) % m
		if divisible == 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	res, n, m := 0, len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && !visited[i][0] {
			numEnclavesDFS(grid, visited, i, 0, n, m)
		}
		if grid[i][n-1] == 1 && !visited[i][m-1] {
			numEnclavesDFS(grid, visited, i, m-1, n, m)
		}
	}

	for i := 0; i < m; i++ {
		if grid[0][i] == 1 && !visited[0][i] {
			numEnclavesDFS(grid, visited, 0, i, n, m)
		}
		if grid[n-1][i] == 1 && !visited[n-1][i] {
			numEnclavesDFS(grid, visited, n-1, i, n, m)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				res++
			}
		}
	}
	return res
}

func numEnclavesDFS(grid [][]int, visited [][]bool, i int, j int, n int, m int) {
	if i < 0 || i >= n || j >= m || j < 0 || grid[i][j] == 0 || visited[i][j] {
		return
	}

	visited[i][j] = true
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		numEnclavesDFS(grid, visited, i+move[0], j+move[1], n, m)
	}
}

// Reference: https://leetcode.com/problems/check-knight-tour-configuration/
func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}

	m, n, i, j := 1, len(grid), 0, 0
	max := n*n - 1
	for m <= max {
		if i+2 < n && j+1 < n && checkValidGridMove(grid, &m, &i, &j, i+2, j+1) {
			continue
		}
		if i+2 < n && j-1 > -1 && checkValidGridMove(grid, &m, &i, &j, i+2, j-1) {
			continue
		}
		if i-2 > -1 && j+1 < n && checkValidGridMove(grid, &m, &i, &j, i-2, j+1) {
			continue
		}
		if i-2 > -1 && j-1 > -1 && checkValidGridMove(grid, &m, &i, &j, i-2, j-1) {
			continue
		}
		if i+1 < n && j-2 > -1 && checkValidGridMove(grid, &m, &i, &j, i+1, j-2) {
			continue
		}
		if i+1 < n && j+2 < n && checkValidGridMove(grid, &m, &i, &j, i+1, j+2) {
			continue
		}
		if i-1 > -1 && j-2 > -1 && checkValidGridMove(grid, &m, &i, &j, i-1, j-2) {
			continue
		}
		if i-1 > -1 && j+2 < n && checkValidGridMove(grid, &m, &i, &j, i-1, j+2) {
			continue
		}
		return false
	}
	return true
}

func checkValidGridMove(grid [][]int, m *int, i *int, j *int, x int, y int) bool {
	if grid[x][y] != *m {
		return false
	}

	*i = x
	*j = y
	*m++
	return true
}

// Reference: https://leetcode.com/problems/find-duplicate-subtrees/
func findDuplicateSubtrees(root *types.TreeNode) []*types.TreeNode {
	res := make([]*types.TreeNode, 0)
	tripletToID := make(map[string]int)
	cnt := make(map[int]int)
	findDuplicateSubtreesTraversal(root, tripletToID, cnt, &res)
	return res
}

func findDuplicateSubtreesTraversal(node *types.TreeNode, tripletToID map[string]int, cnt map[int]int, res *[]*types.TreeNode) int {
	if node == nil {
		return 0
	}

	triplet := strconv.Itoa(findDuplicateSubtreesTraversal(node.Left, tripletToID, cnt, res)) + "," + strconv.Itoa(node.Val) + "," + strconv.Itoa(findDuplicateSubtreesTraversal(node.Right, tripletToID, cnt, res))
	if _, c := tripletToID[triplet]; !c {
		tripletToID[triplet] = len(tripletToID) + 1
	}

	id := tripletToID[triplet]
	cnt[id]++
	if cnt[id] == 2 {
		*res = append(*res, node)
	}
	return id
}

// Reference: https://leetcode.com/problems/clone-graph/
func cloneGraph(node *types.Node) *types.Node {
	if node == nil {
		return nil
	}

	visited := make([]*types.Node, 101)
	return cloneGraphTraversal(node, visited)
}

func cloneGraphTraversal(node *types.Node, visited []*types.Node) *types.Node {
	if visited[node.Val] != nil {
		return visited[node.Val]
	}

	res := new(types.Node)
	res.Val = node.Val
	visited[node.Val] = res
	for _, neighbor := range node.Neighbors {
		res.Neighbors = append(res.Neighbors, cloneGraphTraversal(neighbor, visited))
	}
	return res
}

// Reference: https://leetcode.com/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	s1map, s2map := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < n2-n1; i++ {
		r, l := s2[i+n1]-'a', s2[i]-'a'
		if count == 26 {
			return true
		}

		s2map[r]++
		if s2map[r] == s1map[r] {
			count++
		} else if s2map[r] == s1map[r]+1 {
			count--
		}

		s2map[l]--
		if s2map[l] == s1map[l] {
			count++
		} else if s2map[l] == s1map[l]-1 {
			count--
		}
	}
	return count == 26
}

// Reference: https://leetcode.com/problems/find-the-maximum-number-of-marked-indices/
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i, n := 0, len(nums)
	for j := n / 2; i < n/2 && j < n; j++ {
		if 2*nums[i] <= nums[j] {
			i++
		}
	}
	return i * 2
}

// Reference: https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
func minReorder(n int, connections [][]int) int {
	count, adj := 0, make(map[int][][]int)
	for _, connection := range connections {
		if adj[connection[0]] == nil {
			adj[connection[0]] = make([][]int, 0)
		}
		adj[connection[0]] = append(adj[connection[0]], []int{connection[1], 1})

		if adj[connection[1]] == nil {
			adj[connection[1]] = make([][]int, 0)
		}
		adj[connection[1]] = append(adj[connection[1]], []int{connection[0], 0})
	}

	minReorderDFS(adj, 0, n, &count)
	return count
}

func minReorderDFS(adj map[int][][]int, node int, parent int, count *int) {
	if adj[node] == nil {
		return
	}

	for _, neib := range adj[node] {
		child, sign := neib[0], neib[1]
		if child != parent {
			*count += sign
			minReorderDFS(adj, child, node, count)
		}
	}
}

// Reference: https://leetcode.com/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			m[secret[i]]++
			m[guess[i]]--
			if m[secret[i]] <= 0 {
				cows++
			}
			if m[guess[i]] >= 0 {
				cows++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

// Reference: https://leetcode.com/problems/minimum-time-to-complete-trips/
func minimumTime(time []int, totalTrips int) int64 {
	left, right := int64(0), int64(math.MaxInt64)
	for left < right {
		mid := left + (right-left)/2
		trips := int64(0)
		for _, t := range time {
			trips += mid / int64(t)
			if trips >= int64(totalTrips) {
				break
			}
		}

		if trips < int64(totalTrips) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Reference: https://leetcode.com/problems/construct-quad-tree/
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return solveConstruct(grid, 0, 0, len(grid))
}

func solveConstruct(grid [][]int, x1, y1, length int) *Node {
	if length == 1 {
		return &Node{Val: grid[x1][y1] == 1, IsLeaf: true}
	}

	topLeft := solveConstruct(grid, x1, y1, length/2)
	topRight := solveConstruct(grid, x1, y1+length/2, length/2)
	bottomLeft := solveConstruct(grid, x1+length/2, y1, length/2)
	bottomRight := solveConstruct(grid, x1+length/2, y1+length/2, length/2)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{Val: topLeft.Val, IsLeaf: true}
	}

	return &Node{Val: false, IsLeaf: false, TopLeft: topLeft, TopRight: topRight, BottomLeft: bottomLeft, BottomRight: bottomRight}
}

// Reference: https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital
func minimumFuelCost(roads [][]int, seats int) int64 {
	m := make(map[int][]int)
	fuel := int64(0)
	for _, road := range roads {
		if m[road[0]] == nil {
			m[road[0]] = []int{}
		}
		m[road[0]] = append(m[road[0]], road[1])

		if m[road[1]] == nil {
			m[road[1]] = []int{}
		}
		m[road[1]] = append(m[road[1]], road[0])
	}
	minimumFuelCostDFS(m, 0, -1, seats, &fuel)
	return fuel
}

func minimumFuelCostDFS(m map[int][]int, from, to, seats int, fuel *int64) int64 {
	if len(m) == 0 {
		return int64(0)
	}

	representatives := int64(1)
	for _, node := range m[from] {
		if node != to {
			representatives += minimumFuelCostDFS(m, node, from, seats, fuel)
		}
	}

	if from != 0 {
		*fuel += int64(math.Ceil(float64(representatives) / float64(seats)))
	}
	return representatives
}

// Reference: https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/
func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	overlapping, end := 1, ranges[0][1]
	for i := 1; i < len(ranges); i++ {
		if end < ranges[i][0] {
			overlapping++
		}
		end = max(end, ranges[i][1])
	}

	res, mod := int64(1), int64(1_000_000_007)
	for i := overlapping; i > 0; i-- {
		res <<= 1
		res %= mod
	}
	return int(res)
}
