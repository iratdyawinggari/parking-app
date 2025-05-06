package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lot *ParkingLot

func ExecuteCommands(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

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
			lot = NewParkingLot(size)
		case "park":
			if lot == nil || len(parts) != 2 {
				fmt.Println("Invalid park command or parking lot not created")
				continue
			}
			lot.Park(parts[1])
		case "leave":
			if lot == nil || len(parts) != 3 {
				fmt.Println("Invalid leave command or parking lot not created")
				continue
			}
			hours, _ := strconv.Atoi(parts[2])
			lot.Leave(parts[1], hours)
		case "status":
			if lot == nil {
				fmt.Println("Parking lot not created")
				continue
			}
			lot.Status()
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}

	return scanner.Err()
}
