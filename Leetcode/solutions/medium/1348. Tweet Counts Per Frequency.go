package medium

import "fmt"

// Reference: https://leetcode.com/problems/tweet-counts-per-frequency/
func Leetcode_Tweet_Counts_Per_Frequency() {
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

type TweetCounts struct {
}

func TweetCountsConstructor() TweetCounts {
	tweetCounts := TweetCounts{}
	return tweetCounts
}

func (tw *TweetCounts) RecordTweet(tweetName string, time int) {

}

func (tw *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	return nil
}
