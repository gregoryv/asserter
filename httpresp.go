package asserter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type HttpResponse struct {
	T
	http.Handler
}

// StatusCode checks the expected status code.  Options can be
// io.Reader for body, http.Header for headers or string for an error
// message.
func (t *HttpResponse) StatusCode(exp int, m, p string, opt ...interface{}) {
	t.Helper()
	body, headers, message := t.parse(opt...)
	resp := t.do(m, p, body, headers)
	if resp == nil {
		return
	}
	if message == "" {
		message = "StatusCode"
	}
	if resp.StatusCode != exp {
		t.Fatalf(
			"%s: %s, expected %v %s",
			message, resp.Status, exp, http.StatusText(exp),
		)
	}
}

func (t *HttpResponse) Header(k, exp string, m, p string, opt ...interface{}) {
	t.Helper()
	body, headers, message := t.parse(opt...)
	resp := t.do(m, p, body, headers)
	if message == "" {
		message = k
	} else {
		message = fmt.Sprintf("%s %s", k, message)
	}
	got := resp.Header.Get(k)
	if got != exp {
		t.Fatalf("%s: %q expected %q", message, got, exp)
	}
}

func (t *HttpResponse) do(
	m, p string, body io.Reader, headers http.Header,
) *http.Response {
	t.Helper()
	w := httptest.NewRecorder()
	r, err := http.NewRequest(m, p, body)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	r.Header = headers
	t.ServeHTTP(w, r)
	return w.Result()
}

func (t *HttpResponse) parse(options ...interface{}) (
	body io.Reader, headers http.Header, message string,
) {
	for _, opt := range options {
		switch v := opt.(type) {
		case io.Reader:
			body = v
		case http.Header:
			headers = v
		case string:
			message = v
		}
	}
	return
}
