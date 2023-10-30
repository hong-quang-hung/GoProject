package medium

var sql_medium = []int{
	570, 585, 1934, 1193, 1174, 550,
	1070, 1045, 180,
}

func init() {
	for _, i := range sql_medium {
		Solutions[i] = Leetcode_SQL
	}
}
