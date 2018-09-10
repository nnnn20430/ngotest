#!/bin/bash

# constants
declare -r SCRIPT_ROOT="$(cd "${0%/*}" && echo "$PWD")"

# exports
export GOPATH="${SCRIPT_ROOT}"

# build
cd "${SCRIPT_ROOT}"
mkdir -p bin
go build -o bin/ngotest main.go
