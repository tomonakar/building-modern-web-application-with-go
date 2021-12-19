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
# マイグレーション方法
1. database.yml追加
2. soda generate fizz <ファイル名>で、マイグレーション用ファイル作成
3. ./migrations/xxx_create_yyy_table.up.fizzを編集
4. soda migrateでup.fizz or down.fizzでマイグレーションを実施
```

- リセット（マイグレーションファイルの全てを実行）

```
1. DB接続を全て切る
2. soda reset を実行しマイグレーションファイルを全て実行
```

## PostgreSQLドライバー
- [pgx](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx)

### 途中で読んだものメモ
- [主キーとIndexについて整理する](https://qiita.com/pirorirori_n712/items/b47ade3fdaf8b4a109ba)

- [Parsing and formatting date/time in Go](https://www.pauladamsmith.com/blog/2011/05/go_time.html)

- [GoでMySQLにアクセスしてみる](https://kazuhira-r.hatenablog.com/entry/2021/03/16/223253)



### マイグレーションメモ

```
soda generate fizz AddNotNullReservationIDForRestrictions

soda generate sql SeedRoomsTable

soda generate sql SeedRestrictions

```