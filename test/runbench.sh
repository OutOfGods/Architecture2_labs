#!/usr/bin/env sh

[ $# -eq 0 ] && regex='.' || regex=$1

go test . -bench=$regex -benchtime 1s -benchmem
