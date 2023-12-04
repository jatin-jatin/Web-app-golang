package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("mystring is set to", myString)
	changeUsingPointer(&myString)
	log.Println("mystring is set to", myString)

}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
