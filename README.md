# jimi-cli

This is the all-in-one command line, made by Jimi Family and friends, for their personal projects. The name Jimi means [Jingwen](https://github.com/jingwen-z) and [Mincong](https://github.com/mincong-h).

## Installation

Install Golang:

```sh
brew install go

go version
# go version go1.23.3 darwin/arm64
```

Then build the CLI:

```sh
go mod tidy
go build -o dist/jimi cmd/main.go
```

## Data Sources

Real estate data can be retrieved from <https://www.immo-data.fr/>.
