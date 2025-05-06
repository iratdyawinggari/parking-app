package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	Registration string
}

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file path")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var parkingLot *ParkingLot
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "create_parking_lot":
			if len(parts) != 2 {
				fmt.Println("Invalid create_parking_lot command")
				continue
			}
			size, _ := strconv.Atoi(parts[1])
			parkingLot = NewParkingLot(size)
		case "park":
			if parkingLot == nil || len(parts) != 2 {
				fmt.Println("Invalid park command or parking lot not created")
				continue
			}
			parkingLot.Park(parts[1])
		case "leave":
			if parkingLot == nil || len(parts) != 3 {
				fmt.Println("Invalid leave command or parking lot not created")
				continue
			}
			hours, _ := strconv.Atoi(parts[2])
			parkingLot.Leave(parts[1], hours)
		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			parkingLot.Status()
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}