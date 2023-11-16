package hard

import "fmt"

// Reference: https://leetcode.com/problems/bus-routes/
func init() {
	Solutions[815] = func() {
		fmt.Println("Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6")
		fmt.Println("Output:", numBusesToDestination(S2SoSliceInt("[[1,2,7],[3,6,7]]"), 1, 6))
		fmt.Println("Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12")
		fmt.Println("Output:", numBusesToDestination(S2SoSliceInt("[[7,12],[4,5,15],[6],[15,19],[9,12,13]]"), 15, 12))
	}
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	adjList := make(map[int][]int)
	for bus, route := range routes {
		for _, stop := range route {
			adjList[stop] = append(adjList[stop], bus)
		}
	}

	visitedBuses := make(map[int]bool)
	visitedStops := make(map[int]bool)

	var queue []int
	queue = append(queue, source)
	visitedStops[source] = true

	var busCount int

	for len(queue) > 0 {
		level := len(queue)

		for l := 0; l < level; l++ {
			currentStop := queue[0]
			queue = queue[1:]
			if currentStop == target {
				return busCount
			}

			for _, bus := range adjList[currentStop] {
				if !visitedBuses[bus] {
					visitedBuses[bus] = true

					for _, stop := range routes[bus] {
						if !visitedStops[stop] {
							visitedStops[stop] = true
							queue = append(queue, stop)
						}
					}
				}
			}
		}
		busCount++
	}
	return -1
}
