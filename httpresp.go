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
		message = fmt.Sprintf("%s %s %v", m, p, resp.StatusCode)
	}
	if resp.StatusCode != exp {
		t.Fatalf(
			"%s expected %v %s",
			message, exp, http.StatusText(exp),
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
		message = fmt.Sprintf("%s %s %v", m, p, resp.StatusCode)
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
		message = fmt.Sprintf("%s %s %v", m, p, resp.StatusCode)
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
		message = fmt.Sprintf("%s %s %v", m, p, resp.StatusCode)
	} else {
		message = fmt.Sprintf("%s %s %v %s", m, p, resp.StatusCode, message)
	}
	got := resp.Header.Get(k)
	if got != exp {
		t.Fatalf("%s\n%s: %q expected %q", message, k, got, exp)
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
