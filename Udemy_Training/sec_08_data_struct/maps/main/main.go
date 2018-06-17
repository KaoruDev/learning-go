package main

import "fmt"

func main() {
	initNilMap()
	initMapWithMake()
	initMapWithLiterals()
	loopThroughMap()
}

func loopThroughMap() {
	fmt.Println("\n\n========== Loop Through Maps ==========")
	myMap := map[int]string {
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}

	for key, value := range myMap {
		fmt.Println("Key: ", key, " value: ", value)
	}
}

func initNilMap() {
	fmt.Println("\n\n========== Nil Map ==========")
	/*
	By not assigning something to the declaration, you're basically creating a nil map, which means you cannot
	interact with the map at all.
	 */

	var myMap map[string]string
	fmt.Println(myMap)

	if _, ok := myMap["foo"]; !ok {
		fmt.Println("There is no foo value")
	}

	// myMap["shit will break"] = "will break"
}

func initMapWithLiterals() {
	fmt.Println("\n\n========== Map Literals ==========")
	myMap := map[string]int {
		"foo": 1,
		"bar": 2,
	}

	fmt.Println(myMap)
}

func initMapWithMake() {
	fmt.Println("\n\n========== Creating Maps ==========")
	myMap := make(map[string]int)
	myMap["foo"] = 1
	myMap["bar"] = 1

	delete(myMap, "bar")

	val, ok := myMap["foo"]
	reportValue(val, ok)

	val, ok = myMap["bar"]
	reportValue(val, ok)

	// reportValue(myMap["foo"])
}

func reportValue(val int, ok bool) {
	if ok {
		fmt.Println(val, " is present")
	} else {
		fmt.Println("Not value found")
	}
}
