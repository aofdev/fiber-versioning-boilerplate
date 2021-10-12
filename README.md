# Fiber Versioning Boilerplate
[![Linter](https://github.com/aofdev/fiber-versioning-boilerplate/workflows/Lint/badge.svg)](https://github.com/aofdev/fiber-versioning-boilerplate/actions)

## Prerequisite

Make sure you have the following installed outside the current project directory and available in your `GOPATH`

- [golang](https://golang.org) 1.17+
- [docker](https://www.docker.com/get-started)
- [air](https://github.com/cosmtrek/air) for hot reloading
- [swagger](https://github.com/swaggo/swag) for automatically generate RESTful API

### Features ✨ 
- API Versioning
- Data Versioning
- Clean Architecture
- Postman Collections
- API Documentation with swagger

### Project structure 📁
```
api
  ├── adapters
  │  └── mongo.go
  ├── app.go
  ├── handlers
  │  ├── errors.go
  │  ├── parser_body.go
  │  └── response.go
  ├── utilities
  │  ├── configuration.go
  │  └── converter.go
  └── versions
    ├── v1
    │  ├── entities
    │  ├── factories
    │  ├── repositories
    │  ├── routes
    │  └── usecases
    └── v2
        ├── entities
        ├── factories
        ├── repositories
        ├── routes
        └── usecases
```

## Development environment setup

```sh
make setup
```

## Start the application ⚡️

```sh
make start
```
## Testing 🧪

```sh
make test
```

## Manual 🗒️

Run `make help` to list available commands:

```text
 Choose a command run in fiber-versioning-boilerplate:

  setup           Initialize project
  docker-start    Start docker-compose
  docker-stop     Stop docker-compose
  start           Start the application
  gendoc          Generate docs api with swagger
  test            Run tests coverage
  mongo-dump      Dump MongoDB data for testing
  mongo-restore   Restore MongoDB data for testingg

```