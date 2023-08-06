package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-music-playlists/
func init() {
	Solutions[920] = func() {
		fmt.Println("Input: n = 3, goal = 3, k = 1")
		fmt.Println("Output:", numMusicPlaylists(3, 3, 1))
		fmt.Println("Input: n = 2, goal = 3, k = 0")
		fmt.Println("Output:", numMusicPlaylists(2, 3, 0))
		fmt.Println("Input: n = 2, goal = 3, k = 1")
		fmt.Println("Output:", numMusicPlaylists(2, 3, 1))
	}
}

func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, goal+1)
		for j := 0; j <= goal; j++ {
			dp[i][j] = -1
		}
	}
	return numMusicPlaylistsSolved(n, k, dp, n, goal)
}

func numMusicPlaylistsSolved(n int, k int, dp [][]int, i int, j int) int {
	if i == 0 && j == 0 {
		return 1
	}

	if i == 0 || j == 0 {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	dp[i][j] = (numMusicPlaylistsSolved(n, k, dp, i-1, j-1) * (n - i + 1)) % mod
	if i > k {
		dp[i][j] += (numMusicPlaylistsSolved(n, k, dp, i, j-1) * (i - k)) % mod
		dp[i][j] %= mod
	}
	return dp[i][j]
}
