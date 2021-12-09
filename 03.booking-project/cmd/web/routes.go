package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tomonar/booking/pkg/config"
	"github.com/tomonar/booking/pkg/handlers"
)

// routes is the application router
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// FileServerは、ルートからのファイルシステムのパスを指定して、そのパスに対応するファイルをHTTPリクエストを提供するハンドラを返す
	// Dirは、特定のディレクトリツリーに限定されたネイティブファイルシステムを提供する
	fileServer := http.FileServer(http.Dir("./static/"))
	// StripPrefixは、リクエストURLのパスから指定されたプレフィクスを削除したハンドラを返す
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
