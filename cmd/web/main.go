package main

import (

	"fmt"
	"net/http"
	"github.com/BozaCraft/go-course/pkg/handlers"
	"github.com/BozaCraft/go-course/pkg/config"
	"github.com/BozaCraft/go-course/pkg/render"
	"log"
)

const portNumber = ":8080"

// main is the main aplication function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}