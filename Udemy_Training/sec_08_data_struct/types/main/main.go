package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) fullName() string {
	return p.First + " " + p.Last
}

func main() {
	jimmy := person{"Jimmy", "Jones", 33}
	timmy := person{"Timmy", "Timmah", 5}

	fmt.Println(jimmy)
	fmt.Println(jimmy.fullName(), " is ", jimmy.Age, " years old")
	fmt.Println(timmy.fullName(), " is ", timmy.Age, " years old")
}
