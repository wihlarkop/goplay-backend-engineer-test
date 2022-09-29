# GoPlay Backend Engineer Test

## How To Install

### Initial Project Setup

Clone Project

```bash
  git clone https://github.com/wihlarkop/goplay-backend-engineer-test
```

Go to the project directory

```bash
  cd goplay-backend-engineer-test
```

1. Create a file `.env` based on file `.env-example`
2. for SQLITE_DBNAME you need add extension .db

install golang library (you need install golang first)

```go
go mod download
```

for start server you can run

```go
go run main.go
```

you can open on http://localhost:3000

## Tech Stack

**Server:** Go with Gin Framework and SQLite for database

## Goals

- [x] Login API
- [x] Upload API
- [x] File List API
- [x] Download/Access API
- [ ] Unit Testing (Half Done)
