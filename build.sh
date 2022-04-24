#!/bin/zsh

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go build -o gin-gorm-oj main.go