package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i % 2 == 0 {
			continue // skips this iteration
		}

		fmt.Println(i)

		if i >= 50 {
			break // exits for loop
		}
	}
}
