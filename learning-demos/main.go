// practice to name the code file containing main function main.go
package main

import "fmt"

// you can have global variables that you do not use
// but not local variables

var globalvar string

func unused_local() {
	// var locaval string
	// the above line will cause error
}

func main() {
	unused_local()
	// variables in golang
	var greeting string
	greeting = "Goodbye Moonmen"
	fmt.Println(greeting)
	words := saysomething()
	fmt.Println(words)
}

func saysomething() string {
	return "something"
}
