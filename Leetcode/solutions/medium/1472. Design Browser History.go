package medium

import "fmt"

// Reference: https://leetcode.com/problems/design-browser-history/
func init() {
	Solutions[1472] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]`)
		fmt.Println(`[["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]`)

		fmt.Println(`Output:`)
		browserHistory := BrowserHistoryConstructor(`leetcode.com`)
		browserHistory.Visit(`google.com`)
		browserHistory.Visit(`facebook.com`)
		browserHistory.Visit(`youtube.com`)
		fmt.Println(`browserHistory.Back(1)`, "// return", browserHistory.Back(1))
		fmt.Println(`browserHistory.Back(1)`, "// return", browserHistory.Back(1))
		fmt.Println(`browserHistory.Forward(1)`, "// return", browserHistory.Forward(1))
		browserHistory.Visit(`linkedin.com`)
		fmt.Println(`browserHistory.Forward(2)`, "// return", browserHistory.Forward(2))
		fmt.Println(`browserHistory.Back(2)`, "// return", browserHistory.Back(2))
		fmt.Println(`browserHistory.Back(7)`, "// return", browserHistory.Back(7))
	}
}

type BrowserHistory struct {
	current int
	history map[int]string
}

func (br *BrowserHistory) add(current int, url string) {
	br.history[current] = url
}

func (br *BrowserHistory) init(homepage string) {
	br.current = 0
	br.history = make(map[int]string)
	br.history[br.current] = homepage
}

func (br *BrowserHistory) len() int {
	return len(br.history)
}

func BrowserHistoryConstructor(homepage string) BrowserHistory {
	br := BrowserHistory{}
	br.init(homepage)
	return br
}

func (br *BrowserHistory) Visit(url string) {
	n := br.len()
	for i := br.current + 1; i < n; i++ {
		delete(br.history, i)
	}

	br.current = br.len()
	br.add(br.current, url)
}

func (br *BrowserHistory) Back(steps int) string {
	if br.current-steps > 0 {
		br.current = br.current - steps
	} else {
		br.current = 0
	}
	return br.history[br.current]
}

func (br *BrowserHistory) Forward(steps int) string {
	if br.current+steps < br.len() {
		br.current = br.current + steps
	} else {
		br.current = br.len() - 1
	}
	return br.history[br.current]
}
