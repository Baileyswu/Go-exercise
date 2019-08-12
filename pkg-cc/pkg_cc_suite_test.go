package cc_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "local/Go-exercise/pkg-cc/package1"
	package2 "local/Go-exercise/pkg-cc/package2"
)

func TestPkgCc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PkgCc Suite")
}

func TestFizz(t *testing.T) {
	log.Println("TestFizz")
	if Fizz() != "Fizz" {
		t.Error("not a Fizz")
	}
}

func TestExample(t *testing.T) {
	log.Println("TestExample")
	package2.FizzBuzz(false)
}

/*

go test -run Example -coverprofile=cover.out -coverpkg=./package1,./package2 -covermode=count
go tool cover -html=cover.out

*/
