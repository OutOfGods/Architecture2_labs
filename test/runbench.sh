#!/usr/bin/env bash

[ $# -eq 0 ] && regex='.' || regex=$1

# this dir
DIR=$(cd $(dirname ${BASH_SOURCE[0]}) && pwd)

# GOGC=OFF
go test $DIR/tree_test.go -cpuprofile cpu.prof -memprofile mem.prof -bench=$regex -benchtime 1s -benchmem

# go tool pprof test.test cpu.prof
