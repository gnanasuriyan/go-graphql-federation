#!/usr/bin/env bash

genGraphQL() {
    go run github.com/99designs/gqlgen
}


wire() {
    go run github.com/google/wire/cmd/wire ./internal/
}

if [[ $1 = 'gen-gql' ]]; then
    genGraphQL
elif [[ $1 = 'wire' ]]; then
    wire
else
    echo "Usage: ./do.sh (gen-gql|wire)"
    exit 2
fi