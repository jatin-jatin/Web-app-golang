// practice to name the code file containing main function main.go
package main

import "fmt"

func main() {
	/// multiple return values
	var x, y int = multiple_values()
	fmt.Println(x, y)
}

func multiple_values() (int, int) {
	return 10, 13
}
