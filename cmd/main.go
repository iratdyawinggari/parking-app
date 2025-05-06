package main


import (
	"fmt"
	"os"
	"parking-app/internal" 
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file path")
		return
	}

	err := internal.ExecuteCommands(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
}