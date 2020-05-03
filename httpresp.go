package asserter

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

// BodyIs the same as BodyEquals
func (t *HttpResponse) BodyIs(exp string, m, p string, opt ...interface{}) {
	t.Helper()
	t.BodyEquals([]byte(exp), m, p, opt...)
}

func (t *HttpResponse) BodyEquals(exp []byte, m, p string, opt ...interface{}) {
	t.Helper()
	body, headers, message := t.parse(opt...)
	resp := t.do(m, p, body, headers)
	if resp == nil {
		return
	}
	if message == "" {
		message = "Body"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(message, err)
	}
	if !bytes.Equal(data, exp) {
		t.Error(message, "\ngot:", string(data), "\nexp:", string(exp))
	}
}

// Contains fails if body does not contain exp
func (t *HttpResponse) Contains(exp string, m, p string, opt ...interface{}) {
	t.Helper()
	body, headers, message := t.parse(opt...)
	resp := t.do(m, p, body, headers)
	if resp == nil {
		return
	}
	if message == "" {
		message = "Contains"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(message, err)
	}
	if !bytes.Contains(data, []byte(exp)) {
		t.Error(message, "\ngot:", string(data), "\nexp:", string(exp))
	}
}

// Header checks if the k header matches exp value
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
