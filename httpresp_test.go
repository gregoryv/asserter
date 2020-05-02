package asserter

import (
	"net/http"
	"testing"
)

func TestHttpResponse(t *testing.T) {
	var code int
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	})
	assert := New(t)
	exp := assert().ResponseFrom(handler)

	code = 200
	exp.StatusCode(200, "GET", "/")

	code = 400
	exp.StatusCode(400, "GET", "/")

	assert = New(&noopT{})
	exp = assert().ResponseFrom(handler)
	exp.StatusCode(-1, "::;", "")
}
