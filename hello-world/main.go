package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8010"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	result := fmt.Sprintf("This is the about page and 2+2 is %d", sum)
	_, _ = fmt.Fprint(w, result)

}

func addValues(a, b int) (int, error) {
	sum := a + b
	return sum, nil
}
func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 2.0, 9.0
	res, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}
	result := fmt.Sprintf("%f divided by %f gives %f", x, y, res)
	fmt.Fprint(w, result)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by number less than or equal to zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println("Starting application on port:", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Got error:", err.Error())
	}
}
