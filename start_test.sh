#!/bin/bash
mkdir test
go test -coverprofile=./test/coverage.out ./...
go tool cover -html=./test/coverage.out