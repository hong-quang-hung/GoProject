package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/frequency-tracker/
func Leetcode_Frequency_Tracker() {
	f := FrequencyTrackerConstructor()
	f.Add(5)
	f.Add(5)
	fmt.Println(f.HasFrequency(1))
	fmt.Println(f.HasFrequency(2))
	f.Add(6)
	f.Add(5)
	f.Add(1)
}

type FrequencyTracker struct {
	m    map[int]int
	time map[int]int
}

func (f *FrequencyTracker) init() {
	f.m = make(map[int]int)
	f.time = make(map[int]int)
}

func FrequencyTrackerConstructor() FrequencyTracker {
	f := FrequencyTracker{}
	f.init()
	return f
}

func (f *FrequencyTracker) Add(number int) {
	if _, c := f.time[f.m[number]]; c {
		f.time[f.m[number]]--
		if f.time[f.m[number]] == 0 {
			delete(f.time, f.m[number])
		}
	}

	f.m[number]++
	f.time[f.m[number]]++
}

func (f *FrequencyTracker) DeleteOne(number int) {
	if _, c := f.time[f.m[number]]; c {
		f.time[f.m[number]]--
		if f.time[f.m[number]] == 0 {
			delete(f.time, f.m[number])
		}
	}

	if _, c := f.m[number]; c {
		f.m[number]--
		if f.m[number] == 0 {
			delete(f.m, number)
		}
	}

	if f.m[number] > 0 {
		f.time[f.m[number]]++
	}
}

func (f *FrequencyTracker) HasFrequency(frequency int) bool {
	if time, check := f.time[frequency]; check {
		return time != 0
	}
	return false
}
