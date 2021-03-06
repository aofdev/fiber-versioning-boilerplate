# Fiber Versioning Boilerplate
[![Linter](https://github.com/aofdev/fiber-versioning-boilerplate/workflows/Lint/badge.svg)](https://github.com/aofdev/fiber-versioning-boilerplate/actions)
[![Testing](https://github.com/aofdev/fiber-versioning-boilerplate/workflows/Testing/badge.svg)](https://github.com/aofdev/fiber-versioning-boilerplate/actions)

<p align="center">
    <img src="demo.gif" alt="fiber-versioning-boilerplate">
</p>

## Prerequisite

Make sure you have the following installed outside the current project directory and available in your `GOPATH`

- [golang](https://golang.org) 1.17+
- [docker](https://www.docker.com/get-started)
- [air](https://github.com/cosmtrek/air) for hot reloading
- [swagger](https://github.com/swaggo/swag) for automatically generate RESTful API

### Features โจ 
- API Versioning
- Data Versioning
- Dynamic Struct Type
- Clean Architecture
- Postman Collections
- API Documentation with swagger

### Project structure ๐
```
api
  โโโ adapters
  โ  โโโ mongo.go
  โโโ app.go
  โโโ handlers
  โ  โโโ errors.go
  โ  โโโ parser_body.go
  โ  โโโ response.go
  โโโ utilities
  โ  โโโ configuration.go
  โ  โโโ converter.go
  โโโ versions
    โโโ v1
    โ  โโโ entities
    โ  โโโ factories
    โ  โโโ repositories
    โ  โโโ routes
    โ  โโโ usecases
    โโโ v2
        โโโ entities
        โโโ factories
        โโโ repositories
        โโโ routes
        โโโ usecases
```

## Development environment setup

```sh
make setup
```

## Start the application โก๏ธ

```sh
make start
```
## Testing ๐งช

```sh
make test
```

## Manual ๐๏ธ

Run `make help` to list available commands:

```text
Choose a command run in fiber-versioning-boilerplate:

  setup           Initialize project
  docker-start    Start docker-compose
  docker-stop     Stop docker-compose
  start           Start the application
  copy-env        Copy environment file
  gendoc          Generate docs api with swagger
  test            Run tests coverage
  mongo-dump      Dump MongoDB data for testing
  mongo-restore   Restore MongoDB data for testing

```
