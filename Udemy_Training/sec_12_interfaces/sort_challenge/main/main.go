package main

import (
	"sort"
	"fmt"
)

type mySorted []string

func (s mySorted) foobar() string {
	return "foobar"
}

type Animal struct {
	name string
}

type sortByName []Animal

func (s sortByName) Len() int {
	return len(s)
}

func (s sortByName) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s sortByName) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

type Tiger Animal

func (t Tiger) roar() {
	fmt.Println("ROOARRARA")
}

func main() {
	//sortStrings()
	//sortAnimals()
	sortInt()
}

func sortStrings() {
	fmt.Println("========== Sorted Strings ============")
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println("studyGroup pointer: ", &studyGroup)

	pStudyGroup := &studyGroup
	fmt.Println("pStudyGroup pointer: ", &pStudyGroup)


	// Make copy of original slice
	dup := make([]string, len(studyGroup))
	fmt.Println("new dup: ", dup)
	copy(dup, studyGroup)
	pDup := &dup
	fmt.Println("dup pointer:", &pDup)

	// Add custom functions to string slice
	//sortedDup := mySorted(dup)
	//fmt.Println(sortedDup.foobar())
	//fmt.Println("sortedDup pointer: ", &sortedDup, " dup pointer: ", &dup)

	// Add lib functions to string slice
	sorted := sort.StringSlice(dup)
	sorted.Sort()
	pSorted := &sorted
	fmt.Printf("pSorted type: %T, pDup Type: %T\n", pSorted, pDup)

	fmt.Println("sorted pointer: ", &pSorted, " pDup pointer: ", &pDup)
	fmt.Println("DUP: ", dup)
	fmt.Println("SORTED: ", sorted)

	sort.Sort(sort.Reverse(sorted))

	fmt.Println(dup)
	fmt.Println("OG: ", studyGroup)

	fmt.Println("\n\n")
}

func sortAnimals() {
	fmt.Println("========== Sorted Animals ============")
	animals := []Animal{
		Animal{"zebra"},
		Animal{"monkey"},
		Animal{"aligator"},
		Animal{"croc"},
	}

	tiger := Tiger(Animal{"tiger"})
	tiger.roar()

	sortedAnimals := sortByName(animals)
	sort.Sort(sortedAnimals)

	fmt.Println(animals)
	fmt.Println("\n\n")
}

func sortInt() {
	nums := []int {6, 2, 4, 1, 3, 8}
	sort.Ints(nums)
	fmt.Println("Sorted: ", nums)
	reverseInts(nums)
	fmt.Println("Reversed: ", nums)
}

func reverseInts(og []int) []int {
	for i := len(og) / 2 - 1; i >=0 ; i-- {
		ops := len(og) - 1 - i
		og[i], og[ops] = og[ops], og[i]
	}

	return og
}

