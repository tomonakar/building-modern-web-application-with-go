package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tomonar/booking/internal/config"
	"github.com/tomonar/booking/internal/handlers"
)

// routes is the application router
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)

	// FileServerは、ルートからのファイルシステムのパスを指定して、そのパスに対応するファイルをHTTPリクエストを提供するハンドラを返す
	// Dirは、特定のディレクトリツリーに限定されたネイティブファイルシステムを提供する
	fileServer := http.FileServer(http.Dir("./static/"))
	// StripPrefixは、リクエストURLのパスから指定されたプレフィクスを削除したハンドラを返す
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
