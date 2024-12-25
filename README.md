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

Now you can use it:

```sh
dist/jimi -h
# jimi - A CLI tool for all personal projects of Jimi.

# Usage:
#   jimi [command]

# Available Commands:
#   completion  Generate the autocompletion script for the specified shell
#   help        Help about any command
#   immo        Commands for a real-estate  project.

# Flags:
#   -h, --help   help for jimi

# Use "jimi [command] --help" for more information about a command.
```

## Data Sources

Real estate data can be retrieved from <https://www.immo-data.fr/>.
