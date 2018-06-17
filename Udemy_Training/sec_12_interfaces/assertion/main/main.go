package main

import "fmt"

func main() {
	var name interface{} = "hello"
	str, ok := name.(string)

	// cannot do this
	// fail := "fail"
	// str, ok := fail.(string)

	if (ok) {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("value is not a string", str)
	}
}
