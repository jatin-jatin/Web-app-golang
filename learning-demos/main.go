package main

import "fmt"

func main() {
	count := 12
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	fmt.Println("************************************************")
	lst := []string{"golang", "python", "java", "cpp"}
	for i, language := range lst {
		fmt.Println(i, language)
	}
}
