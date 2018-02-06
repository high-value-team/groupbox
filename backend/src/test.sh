#!/bin/sh
go test -v `(go list ./... | grep -v vendor | grep -v playground)`
