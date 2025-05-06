package internal

import "fmt"

type ParkingLot struct {
	slots []*Car
}

func NewParkingLot(size int) *ParkingLot {
	return &ParkingLot{
		slots: make([]*Car, size),
	}
}

func (pl *ParkingLot) Park(regNum string) {
	for i, slot := range pl.slots {
		if slot == nil {
			pl.slots[i] = &Car{Registration: regNum}
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

func (pl *ParkingLot) Leave(regNum string, hours int) {
	for i, car := range pl.slots {
		if car != nil && car.Registration == regNum {
			pl.slots[i] = nil
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", regNum, i+1, charge)
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", regNum)
}

func (pl *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i, car := range pl.slots {
		if car != nil {
			fmt.Printf("%d %s\n", i+1, car.Registration)
		}
	}
}
