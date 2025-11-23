#!/bin/bash

set -e

echo "===> Starting Go build..."

export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

OUTPUT="bin/api"
MAIN_FILE="cmd/main.go"

echo "Building $MAIN_FILE -> $OUTPUT"

go build \
    -ldflags="-s -w" \
    -trimpath \
    -o "$OUTPUT" \
    "$MAIN_FILE"

echo "===> Build success: $OUTPUT"
