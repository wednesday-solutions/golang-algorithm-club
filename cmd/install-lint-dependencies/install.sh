#!/usr/bin/env bash

echo "--Installing lint dependencies---"

go get github.com/fzipp/gocyclo
go get golang.org/x/tools/cmd/goimports
GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic
go get -u golang.org/x/lint/golint

set -e

gofmt -d -w pkg/
go vet
golint -set_exit_status /...
goimports pkg/
go mod tidy
git add --all

if [ "$?" = "0" ]; then
    echo "---Dependencies Installed---"
else 
    echo "---Installation Failed---"
    exit 1
fi
