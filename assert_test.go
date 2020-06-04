package asserter

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	assert := New(t)
	assert(true).Fail()
	assert(true).Equals(1, 1)
	assert(true).Contains(1, 1)

	// cover the case when Contains fails and contains new lines
	New(&noopT{})().Contains("hello\nworld", "q")

	assert().Contains([]byte("hello"), "h")
	assert().Contains("hello", "h")
	assert().Contains("hello\nworld", "wo")
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

	reader := bytes.NewBufferString("hello")
	assert(false).Contains(reader, "hello")

	broken := brokenReader("break")
	assert(false).Contains(broken, "break")
	assert(true, false) // More than one is disallowed

	ok, bad := assert().Errors()
	ok(nil)
	ok(nil, "message")
	ok(errors.New(""))
	ok(errors.New(""), "message")
	ok(nil, "message")
	bad(nil)
	bad(nil, "message")
}

type brokenReader string

func (br brokenReader) Read(buf []byte) (int, error) {
	return 0, fmt.Errorf(string(br))
}
