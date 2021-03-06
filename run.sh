#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "kittycash binary dir:" "$DIR"
pushd "$DIR" >/dev/null
go run cmd/kittycash/kittycash.go --gui-dir="${DIR}/src/gui/static/" $@
popd >/dev/null
