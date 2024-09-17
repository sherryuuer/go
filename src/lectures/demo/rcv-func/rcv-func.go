package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(parkinglot *ParkingLot, occupyNum int) {
	parkinglot.spaces[occupyNum-1].occupied = true
}

func (parkinglot *ParkingLot) occupySpace(occupyNum int) {
	parkinglot.spaces[occupyNum-1].occupied = true
}

func (parkinglot *ParkingLot) vacateSpace(occupyNum int) {
	parkinglot.spaces[occupyNum-1].occupied = false
}

func main() {
	parkingLot := ParkingLot{
		spaces: make([]Space, 5),
	}
	fmt.Println(parkingLot)

	occupySpace(&parkingLot, 3)
	fmt.Println("Occupied the 3rd space,", parkingLot)

	parkingLot.occupySpace(4)
	fmt.Println("Occupied the 4th space,", parkingLot)

	parkingLot.vacateSpace(4)
	fmt.Println("Vacated the 4th space,", parkingLot)
}
