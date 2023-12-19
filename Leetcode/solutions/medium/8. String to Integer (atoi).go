package medium

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Reference: https://leetcode.com/problems/string-to-integer-atoi/
func init() {
	Solutions[8] = func() {
		fmt.Println(`Input: x = "42"`)
		fmt.Println(`Output:`, myAtoi(``))
		fmt.Println(`Input: x = "4193 with words"`)
		fmt.Println(`Output:`, myAtoi(`4193 with words`))
	}
}

func myAtoi(s string) int {
	tmp, res := strings.Trim(s, ` `), ``
	neg := false
	for i, v := range tmp {
		if i == 0 && (v == '-' || v == '+') {
			if v == '-' {
				neg = true
			}
		} else if v >= 48 && v <= 57 {
			res += string(v)
		} else {
			break
		}
	}

	intans, _ := strconv.Atoi(res)
	if neg {
		intans = -intans
	}

	if intans > math.MaxInt32 {
		intans = math.MaxInt32
	} else if intans < math.MinInt32 {
		intans = math.MinInt32
	}
	return intans
}
