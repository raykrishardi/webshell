package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/raykrishardi/webshell-front/internal/handler"
	"github.com/raykrishardi/webshell-front/internal/pkg/config"
	"github.com/raykrishardi/webshell-front/internal/pkg/render"
)

var portNumber = fmt.Sprintf(":%s", os.Getenv("FRONT_PORT"))

var app config.AppConfig

func main() {
	fmt.Printf("Starting application on port %s\n", portNumber)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewRenderer(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
