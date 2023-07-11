package medium

import "fmt"

// Reference: https://leetcode.com/problems/evaluate-division/
func init() {
	Solutions[399] = func() {
		fmt.Println("Input: equations = [['a','b'],['b','c']], values = [2.0,3.0], queries = [['a','c'],['b','a'],['a','e'],['a','a'],['x','x']]")
		fmt.Println("Output:", calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
		fmt.Println("Input: equations = [['a','b'],['b','c'],['bc','cd']], values = [1.5,2.5,5.0], queries = [['a','c'],['c','b'],['bc','cd'],['cd','bc']]")
		fmt.Println("Output:", calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make(map[string]map[string]float64)
	for i, v := range equations {
		if g[v[0]] == nil {
			g[v[0]] = make(map[string]float64)
		}
		if g[v[1]] == nil {
			g[v[1]] = make(map[string]float64)
		}
		g[v[0]][v[1]] = values[i]
		g[v[1]][v[0]] = 1 / values[i]
	}

	results := make([]float64, 0, len(queries))
	for _, v := range queries {
		x := v[0]
		y := v[1]
		v := make(map[string]bool)
		result := calcEquationDFS(x, y, v, g)
		results = append(results, result)
	}
	return results
}

func calcEquationDFS(x, y string, visited map[string]bool, g map[string]map[string]float64) float64 {
	_, okX := g[x]
	_, okY := g[y]
	if !okX || !okY {
		return -1.0
	}

	if len(g[x]) == 0 {
		return -1.0
	}

	val, okResult := g[x][y]
	if okResult {
		return val
	}

	for k := range g[x] {
		if !visited[k] {
			visited[k] = true
			tmp := calcEquationDFS(k, y, visited, g)
			if tmp == -1.0 {
				continue
			} else {
				return tmp * g[x][k]
			}
		}
	}
	return -1.0
}
