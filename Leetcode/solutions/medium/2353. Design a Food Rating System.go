package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/design-a-food-rating-system/
func init() {
	Solutions[2353] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]`)
		fmt.Println(`[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]`)
		fmt.Println(`Output:`)
		foodRatings := FoodRatingsConstructor([]string{`kimchi`, `miso`, `sushi`, `moussaka`, `ramen`, `bulgogi`}, []string{`korean`, `japanese`, `japanese`, `greek`, `japanese`, `korean`}, []int{9, 12, 8, 15, 14, 7})
		fmt.Println(`foodRatings.HighestRated("korean");`, `// return`, foodRatings.HighestRated(`korean`))
		fmt.Println(`foodRatings.HighestRated("japanese");`, `// return`, foodRatings.HighestRated(`japanese`))
		fmt.Println(`foodRatings.changeRating("sushi", 16);`)
		foodRatings.ChangeRating(`sushi`, 16)
		fmt.Println(`foodRatings.HighestRated("japanese");`, `// return`, foodRatings.HighestRated(`japanese`))
		fmt.Println(`foodRatings.changeRating("ramen", 16);`)
		foodRatings.ChangeRating(`ramen`, 16)
		fmt.Println(`foodRatings.HighestRated("japanese");`, `// return`, foodRatings.HighestRated(`japanese`))
	}
}

type Food struct {
	name    string
	cuisine string
	rating  int
	index   int
}

type FoodHeap []*Food

func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name
	}
	return h[i].rating > h[j].rating
}
func (h FoodHeap) Len() int { return len(h) }
func (h FoodHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h FoodHeap) Empty() bool { return len(h) == 0 }
func (h *FoodHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *FoodHeap) Peek() *Food {
	r := (*h)[0]
	return r
}
func (h *FoodHeap) Push(i interface{}) {
	i.(*Food).index = len(*h)
	*h = append(*h, i.(*Food))
}

type FoodRatings struct {
	hash  map[string]*Food
	heaps map[string]*FoodHeap
}

func FoodRatingsConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	heaps := make(map[string]*FoodHeap)
	hash := make(map[string]*Food)
	for i := 0; i < len(foods); i++ {
		if _, ok := heaps[cuisines[i]]; !ok {
			heaps[cuisines[i]] = new(FoodHeap)
		}
		f := &Food{
			name:    foods[i],
			cuisine: cuisines[i],
			rating:  ratings[i],
		}
		heap.Push(heaps[cuisines[i]], f)
		hash[foods[i]] = f
	}
	return FoodRatings{
		heaps: heaps,
		hash:  hash,
	}
}

func (f *FoodRatings) ChangeRating(food string, newRating int) {
	if item, ok := f.hash[food]; ok {
		f.hash[food].rating = newRating
		heap.Fix(f.heaps[item.cuisine], item.index)
	}
}

func (f *FoodRatings) HighestRated(cuisine string) string {
	if _, ok := f.heaps[cuisine]; ok {
		return f.heaps[cuisine].Peek().name
	}
	return ``
}
