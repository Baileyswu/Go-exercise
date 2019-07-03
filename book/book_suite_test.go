package book_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Suite")
}

// You are only allowed to define BeforeSuite and AfterSuite once in a test suite 
var _ = BeforeSuite(func() {
	fmt.Println("---Before Suite---")
})

var _ = AfterSuite(func() {
	fmt.Println("---After Suite---")
})