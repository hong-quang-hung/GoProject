package easy

var sql_easy = []int{
	197, 577, 584, 595, 620, 1068, 619, 1251,
	1280, 1378, 1581, 1661, 1683, 1757, 1965,
	1075, 1633, 1211, 2356, 1141, 1517, 1148,
	1729, 1731, 1789, 1527, 1978, 1667, 1484,
	610, 196, 596, 1327,
}

func init() {
	for _, i := range sql_easy {
		Solutions[i] = Leetcode_SQL
	}
}
