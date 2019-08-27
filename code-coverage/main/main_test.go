package main

import (
	goflag "flag"
	"fmt"
	"os"
	"strings"
	"testing"

	flag "github.com/spf13/pflag"
)

var (
	pg *bool
)

func init() {
	pg = flag.Bool("print", false, "whether to print weather")
	flag.Parse()
}

func filterGoFlags(args []string, prefixes map[string]bool) ([]string, []string) {
	fmt.Println("filterGoFlags")
	var goFlags []string
	for i := 0; 0 < len(args) && i < len(args); i++ {
		for prefix, hasValue := range prefixes {
			if strings.HasPrefix(args[i], "-"+prefix) {
				// Add the flag to our goflags slice so we can parse it seperatly
				goFlags = append(goFlags, args[i])
				skip := 1
				// If the flag has a corresponding value, add that too
				if hasValue && i+1 < len(args) {
					goFlags = append(goFlags, args[i+1])
					skip = 2
				}
				// Remove the go flags from our pflags slice
				if i+skip <= len(args) {
					args = append(args[:i], args[i+skip:]...)
				}
				// go back one as we just deleted an element
				i--
				break
			}
		}
	}

	return args, goFlags
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	// We need to split up the test flags and the regular app pflags.
	// Then hand them off the the std flags and pflags parsers respectively.
	args, testFlags := filterGoFlags(os.Args, map[string]bool{"test.": true})
	os.Args = args

	// Parse the testing flags
	if err := goflag.CommandLine.Parse(testFlags); err != nil {
		fmt.Println("Error parsing regular test flags:", err)
	}

	// Parse our pflags
	flag.Parse()

	os.Exit(m.Run())
}

// TestMain - test to drive external testing coverage
func TestCurl(t *testing.T) {
	fmt.Println("TestCurl")
	flagset := flagSet{
		prFlag: *pg,
	}
	fmt.Println(flagset)
	runMain(flagset)
}

/*
go test -coverprofile=coverage.out -run=TestMain

go test -c -covermode=count -coverpkg=./...
./main.test -test.coverprofile=coverage.out -test.run=TestMain

curl -XPOST http://127.0.0.1:1234/test
curl -XPOST http://127.0.0.1:1234/deathblow
*/
