package asserter

import (
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
	w := httptest.NewRecorder()
	r, err := http.NewRequest(m, p, body)
	if err != nil {
		t.Fatal(err)
		return
	}
	r.Header = headers
	t.ServeHTTP(w, r)
	resp := w.Result()
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
