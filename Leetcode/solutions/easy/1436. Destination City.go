package easy

import "fmt"

// Reference: https://leetcode.com/problems/destination-city/
func init() {
	Solutions[1436] = func() {
		fmt.Println("Input: paths = [[\"London\",\"New York\"],[\"New York\",\"Lima\"],[\"Lima\",\"Sao Paulo\"]]")
		fmt.Println("Output:", destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
		fmt.Println("Input: paths = [[\"B\",\"C\"],[\"D\",\"B\"],[\"C\",\"A\"]]")
		fmt.Println("Output:", destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
		fmt.Println("Input: paths = [[\"A\",\"Z\"]]")
		fmt.Println("Output:", destCity([][]string{{"A", "Z"}}))
	}
}

func destCity(paths [][]string) string {
	m := make(map[string]interface{})
	for _, path := range paths {
		m[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := m[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
