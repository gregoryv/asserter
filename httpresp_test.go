package asserter

import (
	"net/http"
	"strings"
	"testing"
)

func TestHttpResponse(t *testing.T) {
	var code int
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
	})
	assert := New(t)
	exp := assert().ResponseFrom(handler)
	code = 200
	exp.StatusCode(200, "GET", "/")
	exp.Header("Content-Type", "application/json", "GET", "/")
	exp.Header("Content-Type", "application/json", "GET", "/", "checking")
	code = 400
	exp.StatusCode(400, "GET", "/")

	assert = New(&noopT{})
	exp = assert().ResponseFrom(handler)
	exp.StatusCode(-1, "::;", "")
	exp.StatusCode(500, "GET", "/", "noopT should fail") // covers only
	exp.Header("Content-Type", "text/html", "GET", "/")  // covers only
}

func TestHttpResponse_uses_header(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x") == "" {
			t.Fatal("expected to find x header")
		}
	})
	assert := New(t)
	exp := assert().ResponseFrom(handler)
	header := make(http.Header)
	header.Set("x", "yes")
	exp.StatusCode(200, "GET", "/", header)
}

func TestHttpResponse_parse(t *testing.T) {
	r := &HttpResponse{t, nil}

	header := make(http.Header)
	header.Set("x", "yes")

	_, h, _ := r.parse("msg", header, strings.NewReader("ljklj"))
	if len(h) == 0 {
		t.Fatal(h)
	}
}
