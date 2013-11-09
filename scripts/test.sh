#!/bin/sh

PROJECTROOT="github.com/kir-dev/torpedo"

# test utils
go test "$PROJECTROOT/util"

# test engine
ENV=test go test "$PROJECTROOT/engine"

# test main
go test
