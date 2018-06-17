package main

import "fmt"

func main() {
	fmt.Println("avg: ", average(10, 20, 30, 40, 50))

	prebuilt := []float64{43, 56, 87, 12, 57, 45}
	fmt.Println("vardiatic argument ==========")
	fmt.Println(average(prebuilt...)) // vardiatic argument
}

func average(numbers ...float64) float64 {
	fmt.Println("numbers: ", numbers) // numbers is a Slice (I think this is the same as an array
	fmt.Printf("numbers type: %T \n", numbers)

	var total float64

	for _, val := range numbers { // range returns idx, value
		total += val
	}

	return total / float64(len(numbers))
}
