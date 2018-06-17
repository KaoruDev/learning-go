package main

import (
	"fmt"
	"github.com/KaoruDev/learning_go/Udemy_Training/03_scope/pscope"
)

func main() {
	fmt.Println(pscope.One())
	fmt.Println(pscope.Output())
	pscope.X = "convert"
	fmt.Println(pscope.X)
	fmt.Println(pscope.One())

	adHoc := "adhoc blocks:"
	fmt.Println(adHoc)

	{
		adHoc2 := "cannot be access outside of this block, weird"
		fmt.Println(adHoc2)
	}

	// println(adHoc2) // does not work because adHoc2 is defined within a block

	fmt.Println("======== Closure =========")

	sumOfFive := sumOf(5);
	fmt.Println("5 + 5: ", sumOfFive(5))
}

func sumOf(base int) func(int) int {
	return func(y int) int {
		return base + y;
	}
}
