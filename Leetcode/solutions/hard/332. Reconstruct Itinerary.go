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
	g := make(map[string][]string)
	for _, ticket := range tickets {
		g[ticket[0]] = append(g[ticket[0]], ticket[1])
	}

	for k := range g {
		sort.Strings(g[k])
	}

	res := []string{}
	findItineraryDFS(g, &res, "JFK")
	return res
}

func findItineraryDFS(g map[string][]string, res *[]string, start string) {

}
