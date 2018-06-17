package main

import "fmt"

func main() {
	var x [58]string

	fmt.Println(x) // string's default value is an empty string, which is why you see [       ] on ouput
	fmt.Println("length is: ", len(x))
	fmt.Println(x[3])

	for i := 65; i < (65+26); i++ {
		x[i - 65] = string(i)
	}

	x[46] = "foo"

	foo := "foo"
	fmt.Println("bytes of foo: ",[]byte(foo))
	fmt.Printf("back to string: %s \n",[]byte(foo))


	fmt.Println(x)
	fmt.Println(x[3])
	fmt.Println(x[3:5]) // slicing the array, the return type is a slice and not another array
}
