package main

import "fmt"

func main() {
	demo()
	capacityInspection()
	creatingSlices()
}

func creatingSlices() {
	fmt.Println("\n\n---------- Creating Slices ----------")
	var usingVar []int
	fmt.Println(usingVar) // this will print [], even though it's nil
	fmt.Println(usingVar == nil)

	usingDeclaration := []int{}
	fmt.Println(usingDeclaration)        // this will print []
	fmt.Println(usingDeclaration == nil) // this is now false
}

func capacityInspection() {
	fmt.Println("\n\n---------- Capacity Inspection ----------")
	aSlice := make([]int, 0, 5)
	fmt.Println("----------")
	fmt.Println(aSlice)
	fmt.Println("length of slice: ", len(aSlice))
	fmt.Println("capacity of slice: ", cap(aSlice))
	fmt.Println("----------")

	for i := 0; i < 80; i++ {
		aSlice = append(aSlice, i)
		printSliceStats(aSlice)
	}

	//aSlice[100] = 1000 // this throws an error because array is not that big
	printSliceStats(aSlice)

	a := []int{0, 1, 2, 4, 5, 6}
	b := []int{7, 8, 9, 10, 11}

	fmt.Println(mergeSlices(a, b))
	fmt.Println(mergeSlices(a[:2], b[3:])) // [:2] means from 0-2, [3:] means 3 - end
}

func demo() {
	fmt.Println("\n\n---------- Demo ----------")
	aSlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(aSlice[3])
	fmt.Println(aSlice[2:4])   // prints c, d, i.e. range is exclusive
	fmt.Println("myString"[2]) // prints the rune representing S
}

func printSliceStats(target []int) {
	fmt.Printf("len now: %d, capacity: %d\n", len(target), cap(target))
}

func mergeSlices(a []int, b []int) []int {
	return append(a, b...)
}
