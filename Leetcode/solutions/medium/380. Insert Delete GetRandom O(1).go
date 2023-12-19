package medium

import (
	"fmt"
	"math/rand"
)

// Reference: https://leetcode.com/problems/insert-delete-getrandom-o1/
func init() {
	Solutions[380] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]`)
		fmt.Println(`[[], [1], [2], [2], [], [1], [2], []]`)
		fmt.Println(`Output:`)

		randomizedSet := RandomizedSetConstructor()
		fmt.Println(`randomizedSet.insert(1)`, `//return`, randomizedSet.Insert(1))
		fmt.Println(`randomizedSet.remove(2) `, `//return`, randomizedSet.Remove(2))
		fmt.Println(`randomizedSet.insert(2)`, `//return`, randomizedSet.Insert(2))
		fmt.Println(`randomizedSet.getRandom()`, `//return`, randomizedSet.GetRandom())
		fmt.Println(`randomizedSet.remove(1) `, `//return`, randomizedSet.Remove(1))
		fmt.Println(`randomizedSet.insert(2)`, `//return`, randomizedSet.Insert(2))
		fmt.Println(`randomizedSet.getRandom()`, `//return`, randomizedSet.GetRandom())
	}
}

type RandomizedSet struct {
	m map[int]interface{}
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]interface{})}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.m[val]; ok {
		return false
	}
	r.m[val] = val
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.m[val]; ok {
		delete(r.m, val)
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(r.m))
	counter := 0
	for k := range r.m {
		if counter == idx {
			return k
		}
		counter++
	}
	return -1
}
