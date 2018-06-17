package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First                  string
	Last                   string `json:"-"`     // Tag means do not export
	Age                    int    `json:"Level"` // Change the property name
	lowerCasingNotExported int
}

func main() {
	marshall()
	unmarshall()
	encode()
	decode()
}

func encode() {
	fmt.Println("========== Encoding ==========")
	person := Person{
		First: "Bob",
		Last:  "Marley",
		Age:   20,
		lowerCasingNotExported: 1020,
	}

	json.NewEncoder(os.Stdout).Encode(person);
	fmt.Println("\n")
}

func decode() {
	fmt.Println("========== Decoding ==========")
	person := Person{}
	rawString := `{"First": "Timmy", "Last":"Walker", "Age":21, "level": 5, "extra": "ignored?"}`
	reader := strings.NewReader(rawString)

	json.NewDecoder(reader).Decode(&person)
	aboutPersons(person)
	fmt.Println("\n")
}

func marshall() {
	fmt.Println("========== Marshalling ==========")
	person := Person{
		First: "Bob",
		Last:  "Marley",
		Age:   20,
		lowerCasingNotExported: 1020,
	}

	bs, _ := json.Marshal(person)

	fmt.Println(bs)         // array of bytes
	fmt.Printf("%T \n", bs) // []unit8
	fmt.Println(string(bs)) // json string
	fmt.Println("\n")
}

func unmarshall() {
	fmt.Println("========== Unmarshaling ==========")
	person := Person{}

	rawBytes := []byte(`{"First": "Timmy", "Last":"Walker", "Age":21, "level": 5, "extra": "ignored?"}`)
	json.Unmarshal(rawBytes, &person)

	aboutPersons(person)
	fmt.Println("\n")
}

func aboutPersons (person Person) {
	fmt.Println("First Name: ", person.First)
	fmt.Println("Last Name: ", person.Last) // gets Ignored
	fmt.Println("Age: ", person.Age) // prints 5 because we're looking for "level"
}

