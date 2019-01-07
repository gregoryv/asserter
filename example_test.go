package asserter_test

import (
	"testing"

	"github.com/gregoryv/asserter"
)

var t = &testing.T{}

func ExampleNew() {
	assert := asserter.New(t)
	assert(1 != 2).Errorf("...")
	assert(false).Log("...")
	assert(true).Fail()
}

func Test_something(t *testing.T) {
	assert := asserter.New(t)
	got, err := something()
	exp := 1
	assert(err == nil).Fatal(err)
	assert(got == exp).Errorf("%v, expected %v", got, exp)
}

func something() (int, error) {
	return 1, nil
}
