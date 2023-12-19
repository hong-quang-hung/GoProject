package hard

import "fmt"

// Reference: https://leetcode.com/problems/shortest-path-to-get-all-keys/
func init() {
	Solutions[864] = func() {
		fmt.Println(`Input: grid = ["@.a..","###.#","b.A.B"]`)
		fmt.Println(`Output:`, shortestPathAllKeys([]string{`@.a..`, `###.#`, `b.A.B`}))
		fmt.Println(`Input: grid = ["@..aA","..B#.","....b"]`)
		fmt.Println(`Output:`, shortestPathAllKeys([]string{`@..aA`, `..B#.`, `....b`}))
	}
}

type Pair struct {
	r int
	c int
}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	queue := make([][4]int, 0)
	moves := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	seen := make(map[int]map[Pair]bool)
	keySet := make(map[byte]bool)
	lockSet := make(map[byte]bool)
	allKeys := 0
	startR, startC := -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cell := grid[i][j]
			if cell >= 'a' && cell <= 'f' {
				allKeys += (1 << (cell - 'a'))
				keySet[cell] = true
			}

			if cell >= 'A' && cell <= 'F' {
				lockSet[cell] = true
			}

			if cell == '@' {
				startR = i
				startC = j
			}
		}
	}

	queue = append(queue, [4]int{startR, startC, 0, 0})
	seen[0] = make(map[Pair]bool)
	seen[0][Pair{startR, startC}] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curR, curC := cur[0], cur[1]
		keys, dist := cur[2], cur[3]

		for _, move := range moves {
			newR := curR + move[0]
			newC := curC + move[1]
			if newR >= 0 && newR < m && newC >= 0 && newC < n && grid[newR][newC] != '#' {
				cell := grid[newR][newC]
				if keySet[cell] {

					if ((1 << (cell - 'a')) & keys) != 0 {
						continue
					}
					newKeys := (keys | (1 << (cell - 'a')))
					if newKeys == allKeys {
						return dist + 1
					}

					if _, ok := seen[newKeys]; !ok {
						seen[newKeys] = make(map[Pair]bool)
					}
					seen[newKeys][Pair{newR, newC}] = true
					queue = append(queue, [4]int{newR, newC, newKeys, dist + 1})
				}

				if lockSet[cell] && ((keys & (1 << (cell - 'A'))) == 0) {
					continue
				}

				if !seen[keys][Pair{newR, newC}] {
					seen[keys][Pair{newR, newC}] = true
					queue = append(queue, [4]int{newR, newC, keys, dist + 1})

				}
			}
		}
	}
	return -1
}
