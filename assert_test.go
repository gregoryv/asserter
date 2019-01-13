package asserter

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	assert := New(t)
	assert(true).Fail()
	assert(true).Equals(1, 1)
	assert(true).Contains(1, 1)

	assert().Contains([]byte("hello"), "h")
	assert().Contains("hello", "h")
	assert().Contains("100", 1)
	assert().Contains("100", []byte("1"))

	assert = New(&noopT{})
	assert(false).Helper()
	assert(false).Error()
	assert(false).Errorf("%s", "yes")
	assert(false).Fatal()
	assert(false).Fatalf("%s", "yes")
	assert(false).Fail()
	assert(false).FailNow()
	assert(false).Log()
	assert(false).Logf("%s", "yes")
	assert(false).Equals(true, false)
	assert(false).Equals(true, false).Log("case 1")
	assert(false).Contains("hello", "h")
	assert(false).Contains([]byte("hello"), "h")
	assert(false).Contains([]byte("hello"), 1)
	assert(false).Contains([]byte("hello"), 1.0).Log("a float")
	assert(true, false) // More than one is disallowed
}

var t *noopT = &noopT{} // mock for *testing.T
var err error = fmt.Errorf("")
