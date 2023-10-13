package easy

var sql_easy = []int{
	197, 577, 584, 595, 620, 1068, 1148, 1251,
	1280, 1378, 1581, 1661, 1683, 1757, 1965,
	1075, 1633, 1211,
}

func init() {
	for _, i := range sql_easy {
		Solutions[i] = Leetcode_SQL
	}
}
