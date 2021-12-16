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

## Postgres.app
- [Postgres.app](https://postgresapp.com/documentation/)
- [Dbeaver](https://dbeaver.io/download/)

## ORM
- [Soda](https://gobuffalo.io/en/docs/db/getting-started/)
使い方

```
1. database.yml追加
2. soda generate fizz <Table_name>で、マイグレーション用ファイル作成
3. ./migrations/xxx_create_yyy_table.up.fizzを編集
4. soda migrateでup.fizz or down.fizzでマイグレーションを実施
```