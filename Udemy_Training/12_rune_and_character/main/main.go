package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello World"))

	for i := 50000; i < 50140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo) // prints 97, the decimal value of the "a"
	fmt.Printf("%T \n", foo) // prints the type
	fmt.Printf("%T \n", rune(foo)) // rune is the alias of int32, thus this prints a int32
	fmt.Println(rune(foo))
	fmt.Println(string(97))
}
