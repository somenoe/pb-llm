# https://just.systems

set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default:
    just build

build:
    go build ./cmd/main.go

run:
    go run ./cmd/main.go

lint:
    golangci-lint run
