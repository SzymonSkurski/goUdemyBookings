package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SzymonSkursrki/goUdemyBookings/pkg/config"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/handlers"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":8081"

var app config.AppConfig //gloabal config
var session *scs.SessionManager

func main() {

	//should be env
	app.InProduction = false

	//init session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //persist after close browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //set true if https

	app.Session = session

	tc, err := render.CreateTemplatesCache()
	if err != nil {
		log.Fatal("cannot create template cach")
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("serve port", PORT_NUMBER)
	//simple router
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// _ = http.ListenAndServe(PORT_NUMBER, nil)

	//handle by routes
	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
