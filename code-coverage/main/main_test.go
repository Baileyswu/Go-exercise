package main

import (
	"testing"
)

// TestMain - test to drive external testing coverage
func TestMain(t *testing.T) {
	runMain()
}

/*
go test -coverprofile=coverage.out -run=TestMain

go test -c -covermode=count -coverpkg=./...
./main.test -test.coverprofile=coverage.out -test.run=TestMain

curl -XPOST http://127.0.0.1:1234/test
curl -XPOST http://127.0.0.1:1234/deathblow
*/
