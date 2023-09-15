package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/reconstruct-itinerary/
func init() {
	Solutions[332] = func() {
		fmt.Println("Input: tickets = [['MUC','LHR'],['JFK','MUC'],['SFO','SJC'],['LHR','SFO']]")
		fmt.Println("Output:", findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
		fmt.Println("Input: tickets = [['JFK','SFO'],['JFK','ATL'],['SFO','ATL'],['ATL','JFK'],['ATL','SFO']]")
		fmt.Println("Output:", findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
	}
}

func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	g := make(map[string][]string)
	for _, ticket := range tickets {
		g[ticket[0]] = append(g[ticket[0]], ticket[1])
	}

	for k := range g {
		sort.Strings(g[k])
	}
	return findItineraryDFS(g, []string{"JFK"}, n)
}

func findItineraryDFS(g map[string][]string, curr []string, n int) []string {
	if len(curr) == n+1 {
		return curr
	}

	next := (curr)[len(curr)-1]
	for i, v := range g[next] {
		if v == "" {
			continue
		}

		g[next][i] = ""
		d := append(curr, v)
		p := findItineraryDFS(g, d, n)
		if p != nil {
			return p
		}
		g[next][i] = v
	}
	return nil
}
