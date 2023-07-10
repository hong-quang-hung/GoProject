package easy

import "fmt"

func init() {
	Solutions[1603] = Leetcode_Design_Parking_System
}

// Reference: https://leetcode.com/problems/design-parking-system/
func Leetcode_Design_Parking_System() {
	fmt.Println("Input:")
	fmt.Println("['ParkingSystem', 'addCar', 'addCar', 'addCar', 'addCar']")
	fmt.Println("[[1, 1, 0], [1], [2], [3], [1]]")
	fmt.Println("Output:")
	parkingSystem := ParkingSystemConstructor(1, 1, 0)
	fmt.Println("parkingSystem.AddCar(1)", "-->", parkingSystem.AddCar(1))
	fmt.Println("parkingSystem.AddCar(2)", "-->", parkingSystem.AddCar(2))
	fmt.Println("parkingSystem.AddCar(3)", "-->", parkingSystem.AddCar(3))
	fmt.Println("parkingSystem.AddCar(1)", "-->", parkingSystem.AddCar(1))
}

type ParkingSystem struct {
	carTypes map[int]int
}

func ParkingSystemConstructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		carTypes: map[int]int{1: big, 2: medium, 3: small},
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.carTypes[carType] > 0 {
		p.carTypes[carType] -= 1
		return true
	}
	return false
}
