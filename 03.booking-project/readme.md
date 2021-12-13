# go practice
This is the repository for my bookings and reservations project.

- Built in Go version 1.7
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards SCS session management](github.com/alexedwards/scs/v2)
- Uses [nosurf](github.com/justinas/nosurf)


## テスト

```go
go test -v
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```