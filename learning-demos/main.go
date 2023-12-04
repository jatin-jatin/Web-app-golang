// practice to name the code file containing main function main.go
package main

import "fmt"

func main() {
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
