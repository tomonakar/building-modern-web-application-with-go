package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cache-practice/pkg/config"
	"cache-practice/pkg/handlers"
	"cache-practice/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	// クッキーの有効期限を設定
	session.Lifetime = 24 * time.Hour
	// クッキーの永続化設定
	session.Cookie.Persist = true
	// クッキーのSameSite属性
	session.Cookie.SameSite = http.SameSiteLaxMode
	// クッキーの暗号化設定
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
