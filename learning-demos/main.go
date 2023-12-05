package main

import (
	"fmt"
	"sort"
)

type user struct {
	FirstName string
	LastName  string
}

func map_example() {
	myMap := make(map[string]user)
	myMap["dog"] = user{
		FirstName: "jatin",
		LastName:  "Lachhwani",
	}
	fmt.Println(myMap["dog"])
}

func slice_example() {
	var slc []string
	slc = append(slc, "Trevor")
	slc = append(slc, "Green")
	fmt.Println(slc)
	sort.Strings(slc)
	fmt.Println(slc)
}

func main() {
	map_example()
	slice_example()
}
