package asserter_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
)

var (
	t       = &testing.T{}
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
)

func ExampleNew() {
	assert := asserter.New(t)
	assert(1 != 2).Errorf("...")
	got, exp := 1, 1
	assert(got == exp).Fail()
	assert().Equals(got, exp)
}

func ExampleHttpResponse_StatusCode() {
	assert := asserter.New(t)
	exp := assert().ResponseFrom(handler)
	// io.Reader option means body
	exp.StatusCode(200, "POST", "/", strings.NewReader("the body"))

	// string option means error message
	exp.StatusCode(200, "GET", "/", "should be ok")

	// http.Header additional headers
	exp.StatusCode(200, "GET", "/", http.Header{
		"Content-Type": []string{"text/plain"},
	})
}

func ExampleWrap() {
	w := asserter.Wrap(t)
	w.MixOk(something())

	got := 2
	exp := 1
	w.Assert().Equals(got, exp)
	// Same as
	w.Assert(got == exp).Errorf("got %v, expected %v", got, exp)

	resp, _ := http.Get("http://example.com")
	w.Assert().Contains(resp.Body, "</html>")
}

func something() (int, error) {
	return 1, nil
}
