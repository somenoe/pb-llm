# https://just.systems

set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default:
    just build

build *FLAGS:
    go build ./cmd/main.go {{ FLAGS }}

run *FLAGS:
    go run ./cmd/main.go {{ FLAGS }}

lint *FLAGS:
    golangci-lint run {{ FLAGS }}
