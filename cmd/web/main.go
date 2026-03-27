package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/I-Maged/00-golang-first-server/pkg/config"
	"github.com/I-Maged/00-golang-first-server/pkg/handlers"
	"github.com/I-Maged/00-golang-first-server/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// should be true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //true in production

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHnadlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting server on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
