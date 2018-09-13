#!/bin/bash

GOOS=js GOARCH=wasm go build -o out/main.wasm
go run out/serve.go
