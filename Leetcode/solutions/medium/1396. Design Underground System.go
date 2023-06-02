package medium

import "fmt"

// Reference: https://leetcode.com/problems/design-underground-system/
func Leetcode_Design_Underground_System() {
	fmt.Println("Input:")
	fmt.Println("['UndergroundSystem','checkIn','checkIn','checkIn','checkOut','checkOut','checkOut','getAverageTime','getAverageTime','checkIn','getAverageTime','checkOut','getAverageTime']")
	fmt.Println("[[],[45,'Leyton',3],[32,'Paradise',8],[27,'Leyton',10],[45,'Waterloo',15],[27,'Waterloo',20],[32,'Cambridge',22],['Paradise','Cambridge'],['Leyton','Waterloo'],[10,'Leyton',24],['Leyton','Waterloo'],[10,'Waterloo',38],['Leyton','Waterloo']]")
	fmt.Println("Output:")
	undergroundSystem := UndergroundSystemConstructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	fmt.Println("undergroundSystem.GetAverageTime('Paradise', 'Cambridge')", "-->", undergroundSystem.GetAverageTime("Paradise", "Cambridge"))
	fmt.Println("undergroundSystem.GetAverageTime('Leyton', 'Waterloo')", "-->", undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckIn(10, "Leyton", 24)
	fmt.Println("undergroundSystem.GetAverageTime('Leyton', 'Waterloo')", "-->", undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	undergroundSystem.GetAverageTime("Leyton", "Waterloo")
	fmt.Println("undergroundSystem.GetAverageTime('Leyton', 'Waterloo')", "-->", undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
}

type UndergroundSystem struct {
	idTime     map[int]int
	idStart    map[int]string
	travelTime map[string]map[string]int
	travelFreq map[string]map[string]int
}

func UndergroundSystemConstructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]int),
		make(map[int]string),
		make(map[string]map[string]int),
		make(map[string]map[string]int)}
}

func (u *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	u.idTime[id] = t
	u.idStart[id] = stationName
	if _, ok := u.travelTime[stationName]; !ok {
		u.travelTime[stationName] = make(map[string]int)
		u.travelFreq[stationName] = make(map[string]int)
	}
}

func (u *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	time := t - u.idTime[id]
	start := u.idStart[id]
	end := stationName
	u.travelTime[start][end] += time
	u.travelFreq[start][end]++
}

func (u *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return float64(u.travelTime[startStation][endStation]) / float64(u.travelFreq[startStation][endStation])
}
