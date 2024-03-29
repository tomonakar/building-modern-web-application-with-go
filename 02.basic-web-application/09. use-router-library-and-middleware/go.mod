module cache-practice

go 1.17

replace cache-practice/pkg/render => ./pkg/render

replace cache-practice/pkg/handlers => ./pkg/handlers

replace cache-practice/pkg/config => ./pkg/config

replace cache-practice/pkg/models => ./pkg/models

require github.com/go-chi/chi v1.5.4

require (
	github.com/alexedwards/scs/v2 v2.5.0 // indirect
	github.com/justinas/nosurf v1.1.1 // indirect
)
