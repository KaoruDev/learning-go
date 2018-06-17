package main

import (
	"fmt"
	"strconv"
)

const metersToYards float64 = 1.09361

func main() {
	var meters string
	fmt.Println("Enter meters swam: ")

	fmt.Scan(&meters)

	for {
		fmt.Scan(&meters)
		result, err := strconv.Atoi(meters)

		if err == nil {
			yards := float64(result) * metersToYards

			fmt.Println(meters, " meters is ", yards, " yards")
			break
		}
		fmt.Println("Please enter a number: ")
	}

}
