package easy

import "fmt"

// Reference: https://leetcode.com/problems/image-smoother/
func init() {
	Solutions[661] = func() {
		fmt.Println(`Input: img = [[1,1,1],[1,0,1],[1,1,1]]`)
		fmt.Println(`Output:`, imageSmoother(S2SoSliceInt("[[1,1,1],[1,0,1],[1,1,1]]")))
		fmt.Println(`Input: img = [[100,200,100],[200,50,200],[100,200,100]]`)
		fmt.Println(`Output:`, imageSmoother(S2SoSliceInt("[[100,200,100],[200,50,200],[100,200,100]]")))
	}
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if 0 <= x && x < m && 0 <= y && y < n {
						sum += img[x][y] & 255
						count += 1
					}
				}
			}
			img[i][j] |= (sum / count) << 8
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			img[i][j] >>= 8
		}
	}
	return img
}
