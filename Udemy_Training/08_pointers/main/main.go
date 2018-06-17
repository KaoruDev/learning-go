package main

import "fmt"

func main() {
	a := 43

	var b *int = &a // *int means a pointer to an int, I'm guess *<type> means a pointer to some type

	fmt.Println("memory address of a: ", b)
	fmt.Println("dereference: ", *b) // this means dereference the pointer, i.e. get the value, not the memory address

	sumOf43 := sumOfClosure(a)
	sumOf43NonClosure := sumOfNonClosure(b)

	*b = 32 // we're reassigning a from underneathe it

	fmt.Println("a is: ", a)          // 32
	fmt.Println("b is: ", b)          // 32
	fmt.Println(sumOf43(1))           // 44 because we're passing the value to the function
	fmt.Println(sumOf43NonClosure(1)) // 33 because we changed the value of the memory address

	five := 5
	fmt.Println("five :", five)
	zero(&five)
	fmt.Println("five is no longer five: ", five)
}

func sumOfClosure(base int) func(int) int {
	fmt.Println("memory address of base in sumOfClosure: ", &base)
	return func(x int) int {
		return base + x
	}
}

func sumOfNonClosure(base *int) func(int) int {
	fmt.Println("memory address of arugment passed in sumOfNonClosure: ", base)

	// lol this is different cause go has to store the value elsewhere
	fmt.Println("memory address of base in sumOfNonClosure: ", &base)

	return func(x int) int {
		return *base + x
	}
}

// Changes the value of the memory address passed in to zero
func zero(og *int) {
	*og = 0
}
