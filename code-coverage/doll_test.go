package size

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
}

func TestSize(t *testing.T) {
	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

/*

https://blog.golang.org/cover
go test -cover
go test -coverprofile=coverage.out
go test -covermode=count -coverprofile=count.out
        -covermode=set/count/atomic

go tool cover -func=coverage.out
go tool cover -html=coverage.out

*/
