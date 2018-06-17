package main

import "fmt"

const bar = "BAR"

const (
	pi   = 3.14
	lang = "Go"
	A    = iota
	B    // there's no need to assign these as it will all equal iota
	C
	cainine = "canine"
	dog
	wolf
)

const (
	D = iota // this iota will reset back to 0
	E
	F
)

const (
	_  = iota             // throw away zero
	KB = 1 << (iota * 10) // shift bit 10 times, == 1024, or 1kb
	MB = 1 << (iota * 10) // shift bit 20 times, == 1MG
	GB = 1 << (iota * 10) // shift bit 30 times == 1GB
	TB = 1 << (iota * 10) // shift bit 30 times == 1TB
)

func main() {
	const foo = "FOO"
	CAR := "world"

	CAR = "hello"

	fmt.Println(bar)
	fmt.Println(foo)
	fmt.Println(CAR)
	fmt.Printf("a wolf is a %s and a dog is a %s \n", wolf, dog)

	fmt.Println("\n========== memory size ==========")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
