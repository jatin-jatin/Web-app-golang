package main

import (
	"errors"
	"log"
)

func main() {
	res, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", res)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}

// go test -cover // to see the percentage of functions covered for testing
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
