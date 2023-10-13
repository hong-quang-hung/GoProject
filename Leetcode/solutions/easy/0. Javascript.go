package easy

var js_easy = []int{
	2619, 2620, 2621, 2626, 2629, 2634, 2635, 2637,
	2648, 2665, 2666, 2667, 2677, 2695, 2703, 2704,
	2715, 2723, 2724, 2725, 2726, 2727,
}

func init() {
	for _, i := range js_easy {
		Solutions[i] = Leetcode_Javascript
	}
}
