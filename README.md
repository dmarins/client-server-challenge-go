# Client-Server-API - FullCycle Go Expert Challenge
https://goexpert.fullcycle.com.br/pos-goexpert/

[![Go](https://img.shields.io/badge/go-1.21.6-informational?logo=go)](https://go.dev)

This repository contains the first challenge of Postgraduate Go Expert.

## Clone the project

```
$ git clone https://github.com/dmarins/client-server-challenge-go.git
$ cd client-server-challenge-go
```

## Download dependencies

```
$ go mod tidy
```

## Start server

```
$ go run ./server/server.go
```

The server exposes the endpoint `http://localhost:8080/cotacao` which consumes an external API to obtain the current dollar exchange rate. It then persists in a SQLITE database.

## Start client

```
$ go run ./client/client.go
```

The client consumes the endpoint exposed by the server, interprets the response in JSON and generates a file `cotacao.txt` with the current dollar exchange rate.