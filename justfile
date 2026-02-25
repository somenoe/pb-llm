# https://just.systems

set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default:
    just build

build *FLAGS:
    go build . {{ FLAGS }}

run *FLAGS:
    go run . {{ FLAGS }}

lint *FLAGS:
    golangci-lint run {{ FLAGS }}
