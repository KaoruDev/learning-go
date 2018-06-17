package main

import (
	"fmt"
)

func main() {
	sayHello("Billy Bob")
	sayHello("sally")
	sayHello("stranger")
	detectType(5)
	detectType("string")
	detectType([]string { "foo", "bar" })
}

func sayHello(friendName string) {
	switch {
	case len(friendName) == 2:
		fmt.Println("my friend has 2 letters in their name")
	case friendName == "Sally":
		fmt.Println("Hey Sally")
	case friendName == "billy bob":
		fmt.Println("yo Billy bob")
	default:
		fmt.Println("hello stranger named, ", friendName)
	}
}

func detectType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("this is an int")
	case string:
		fmt.Println("this is a string")
	case complex64:
		fmt.Println("thiis is a complex64")
	default:
		fmt.Println("Unexpected type")
	}
}
