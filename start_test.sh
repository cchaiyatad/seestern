#!/bin/bash
mkdir test
mkdir test/cover
go test -coverprofile=./test/cover/coverage.out ./...
go tool cover -html=./test/cover/coverage.out