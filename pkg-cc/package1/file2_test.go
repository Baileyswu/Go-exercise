package package1_test

import (
	"testing"

	. "local/Go-exercise/pkg-cc/package1"
	"local/Go-exercise/pkg-cc/package2"
)

func TestBuzz(t *testing.T) {
	if Buzz() != "Buzz" {
		t.Error("not a Buzz")
	}
}

func TestFoo(t *testing.T) {
	if package2.Foo() != "Foo" {
		t.Error("not a Foo")
	}
}
