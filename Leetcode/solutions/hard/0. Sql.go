package hard

var sql_hard = []int{
	601,
}

func init() {
	for _, i := range sql_hard {
		Solutions[i] = Leetcode_SQL
	}
}
