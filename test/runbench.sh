#!/usr/bin/env sh

[ $# -eq 0 ] && regex='.' || regex=$1

# GOGC=OFF
go test . -cpuprofile cpu.prof -memprofile mem.prof -bench=$regex -benchtime 1s -benchmem

# go tool pprof test.test cpu.prof
