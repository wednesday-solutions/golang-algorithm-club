#!/usr/bin/env bash

echo "Installing pre-commit dependencies"

go get github.com/fzipp/gocyclo
go get golang.org/x/tools/cmd/goimports
GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic
go get -u golang.org/x/lint/golint

if [ "$?" = "0" ]; then
    echo "Dependencies Installed."
else 
    echo "Installation Failed."
    exit 1
fi