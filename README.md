# gator-rss

A multi-user RSS aggregator CLI written in Go and backed by PostgreSQL. **Work in progress.**

## Status

- [x] Config file handling (`~/.gatorconfig.json`)
- [x] Command registry and `login` command
- [x] Users table and type-safe queries generated with [sqlc](https://sqlc.dev)
- [ ] Register and reset users
- [ ] Add, follow, and list feeds
- [ ] Background worker that scrapes posts on an interval
- [ ] Browse aggregated posts

## Tech

Go · PostgreSQL · sqlc · goose migrations

## Setup

Requires Go and a running PostgreSQL instance.

1. Create `~/.gatorconfig.json`:

   ```json
   { "db_url": "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable" }
   ```

2. Run the migrations in `sql/schema` (they use [goose](https://github.com/pressly/goose)):

   ```sh
   goose -dir sql/schema postgres "<your db_url>" up
   ```

3. Build and run:

   ```sh
   go build -o gator .
   ./gator login <username>
   ```

## Project layout

```
main.go              entry point and command wiring
handler_*.go         command handlers
internal/config      reads/writes ~/.gatorconfig.json
internal/database    sqlc-generated query code
sql/schema           goose migrations
sql/queries          SQL queries used by sqlc
```
