package hard

import "fmt"

// Reference: https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
func Leetcode_Max_Num_Edges_To_Removet() {
	fmt.Println("Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]")
	fmt.Println("Output:", maxNumEdgesToRemove(4, S2SoSliceInt("[[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]")))
	fmt.Println("Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]")
	fmt.Println("Output:", maxNumEdgesToRemove(4, S2SoSliceInt("[[3,1,2],[3,2,3],[1,1,4],[2,1,4]]")))
	fmt.Println("Input: n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]")
	fmt.Println("Output:", maxNumEdgesToRemove(4, S2SoSliceInt("[[3,2,3],[1,1,2],[2,3,4]]")))
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	res, unionA, unionB := len(edges), NewUnionFind(n+1), NewUnionFind(n+1)
	componentsA, componentsB := n, n
	for _, edge := range edges {
		if edge[0] == 3 {
			setA := unionA.UnionSet(edge[1], edge[2])
			setB := unionB.UnionSet(edge[1], edge[2])
			if setA {
				componentsA--
			}
			if setB {
				componentsB--
			}
			if setA || setB {
				res--
			}
		}
	}

	for _, edge := range edges {
		if edge[0] == 1 {
			if unionA.UnionSet(edge[1], edge[2]) {
				componentsA--
				res--
			}
		} else if edge[0] == 2 {
			if unionB.UnionSet(edge[1], edge[2]) {
				componentsB--
				res--
			}
		}
	}

	if componentsA == 1 && componentsB == 1 {
		return res
	}
	return -1
}
