package main

import "fmt"

func main() {
	sayStatement(true)
	sayStatement(false)
}

func sayStatement(result bool) {
	if foo:= "foobar"; result {
		fmt.Println("foo is: ", foo)
	} else {
		fmt.Println("foo is not in scope!")
	}
}
