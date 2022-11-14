#!/bin/bash
set -ex

GOOS=windows GOARCH=amd64 go build -o build
GOOS=linux GOARCH=amd64 go build -o build
GOOS=darwin GOARCH=amd64 go build -o build/mgoer
