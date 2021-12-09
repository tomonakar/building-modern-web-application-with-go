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

	// FileServerは、ルートからのファイルシステムのパスを指定して、そのパスに対応するファイルをHTTPリクエストを提供するハンドラを返す
	// Dirは、特定のディレクトリツリーに限定されたネイティブファイルシステムを提供する
	fileServer := http.FileServer(http.Dir("./static/"))
	// StripPrefixは、リクエストURLのパスから指定されたプレフィクスを削除したハンドラを返す
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
