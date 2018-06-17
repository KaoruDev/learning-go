package main

import "fmt"

func main() {
	varAssignments();
	declaringEmpties();
}

func varAssignments() {
	var role string // declaring
	role = "cowboy"

	var city  = "the city" // initializing variable

	var state, country string = "ca", "usa";

	a := 10 // Short hand variable assignment
	b:= "goland"

	fmt.Println("====== Variable Assignments =======")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", role)
	fmt.Printf("%v \n\n", city)
	fmt.Printf("state: %v, country: %v \n\n", state, country)
}

func declaringEmpties() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println("====== Declaring Empty Variables ==========")
	fmt.Printf("%v \n", a);
	fmt.Printf("%v \n", b);
	fmt.Printf("%v \n", c);
	fmt.Printf("%v \n", d);
}
