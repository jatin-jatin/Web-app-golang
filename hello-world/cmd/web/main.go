package main

import (
	"fmt"
	"net/http"

	"github.com/jatin-jatin/go-course/pkg/handlers"
)

const portNumber = ":8010"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting application on port:", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Got error:", err.Error())
	}
}
