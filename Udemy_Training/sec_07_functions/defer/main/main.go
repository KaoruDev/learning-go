package main

import "fmt"

func main() {
	defer say(", bob");
	defer say("world");
	say("hello")
}

func say(word string) {
	fmt.Println(word)
}
