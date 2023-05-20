package medium

import "fmt"

// Reference: https://leetcode.com/problems/evaluate-division/
func Leetcode_Calc_Equation() {
	fmt.Println("Input: equations = [['a','b'],['b','c']], values = [2.0,3.0], queries = [['a','c'],['b','a'],['a','e'],['a','a'],['x','x']]")
	fmt.Println("Output:", calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
	fmt.Println("Input: equations = [['a','b'],['b','c'],['bc','cd']], values = [1.5,2.5,5.0], queries = [['a','c'],['c','b'],['bc','cd'],['cd','bc']]")
	fmt.Println("Output:", calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
	fmt.Println("Input: equations = [['a','b']], values = [0.5], queries = [['a','b'],['b','a'],['a','c'],['x','y']]")
	fmt.Println("Output:", calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}))
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	return nil
}
