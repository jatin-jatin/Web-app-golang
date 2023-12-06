package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jatin-jatin/go-course/pkg/config"
	"github.com/jatin-jatin/go-course/pkg/handlers"
	"github.com/jatin-jatin/go-course/pkg/render"
)

const portNumber = ":8010"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Got error:", err.Error())
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting application on port:", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Got error:", err.Error())
	}
}
