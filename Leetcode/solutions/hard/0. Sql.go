package hard

var sql_hard = []int{
	185, 601,
}

func init() {
	for _, i := range sql_hard {
		Solutions[i] = Leetcode_SQL
	}
}
