package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/tweet-counts-per-frequency/
func init() {
	Solutions[1348] = func() {
		fmt.Println("Input:")
		fmt.Println("['TweetCounts','recordTweet','recordTweet','recordTweet','getTweetCountsPerFrequency','getTweetCountsPerFrequency','recordTweet','getTweetCountsPerFrequency']")
		fmt.Println("[[],['tweet3',0],['tweet3',60],['tweet3',10],['minute','tweet3',0,59],['minute','tweet3',0,60],['tweet3',120],['hour','tweet3',0,210]]")

		fmt.Println("Output:")
		tweetCounts := TweetCountsConstructor()
		tweetCounts.RecordTweet("tweet3", 0)
		tweetCounts.RecordTweet("tweet3", 60)
		tweetCounts.RecordTweet("tweet3", 10)
		tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)
		fmt.Println("tweetCounts.GetTweetCountsPerFrequency('minute', 'tweet3', 0, 59) =>", tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59))
		fmt.Println("tweetCounts.GetTweetCountsPerFrequency('minute', 'tweet3', 0, 60) =>", tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60))
		tweetCounts.RecordTweet("tweet3", 120)
		fmt.Println("tweetCounts.GetTweetCountsPerFrequency('hour', 'tweet3', 0, 210)  =>", tweetCounts.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210))
	}
}

type TweetCounts struct {
	times map[string][]int
	freq  map[string]int
}

func (tw *TweetCounts) init() {
	tw.times = make(map[string][]int)
	tw.freq = map[string]int{"minute": 60, "hour": 3600, "day": 86400}
}

func (tw *TweetCounts) bns(time int, store []int) int {
	low, high := 0, len(store)-1
	for low <= high {
		median := (low + high) / 2
		if store[median] < time {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return low
}

func TweetCountsConstructor() TweetCounts {
	tweetCounts := TweetCounts{}
	tweetCounts.init()
	return tweetCounts
}

func (tw *TweetCounts) RecordTweet(tweetName string, time int) {
	if tw.times[tweetName] == nil {
		tw.times[tweetName] = []int{}
	}

	mid := tw.bns(time, tw.times[tweetName])
	tw.times[tweetName] = insert(tw.times[tweetName], mid, time)
}

func (tw *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	times := tw.times[tweetName]
	timeChunk := tw.freq[freq]
	counts := make([]int, int(math.Ceil(float64(endTime-startTime+1)/float64(timeChunk))))
	for i := 0; i < len(times); i++ {
		if times[i] < startTime {
			continue
		}

		if times[i] > endTime {
			break
		}

		counts[(times[i]-startTime+1)/timeChunk]++
	}
	return counts
}
