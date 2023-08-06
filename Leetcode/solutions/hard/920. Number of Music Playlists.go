package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-music-playlists/
func init() {
	Solutions[920] = func() {
		fmt.Println("Input: n = 3, goal = 3, k = 1")
		fmt.Println("Output:", numMusicPlaylists(3, 1, 1))
		fmt.Println("Input: n = 2, goal = 3, k = 0")
		fmt.Println("Output:", numMusicPlaylists(2, 3, 0))
		fmt.Println("Input: n = 2, goal = 3, k = 1")
		fmt.Println("Output:", numMusicPlaylists(2, 3, 1))
	}
}

func numMusicPlaylists(n int, goal int, k int) int {
	return 1
}
