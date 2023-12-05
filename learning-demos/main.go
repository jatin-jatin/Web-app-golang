package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German shephered",
	}
	PrintInfo(&dog)
	gori := Gorilla{
		Name:          "gauri",
		color:         "black",
		NumberOfTeeth: 2,
	}
	PrintInfo(&gori)
}

func (d *Dog) Says() string {
	return "wooh wooh"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
func (g *Gorilla) Says() string {
	return "moohhh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
