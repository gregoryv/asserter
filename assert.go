/* package defines AssertFunc wrapper of testing.T

Online assertions are done by wrapping the T in a test

    func TestSomething(t *testing.T) {
        assert := asserter.New(t)
        got, err := Something()
        t.Logf("%v, %v := Something()", got, err)
        assert(err == nil).Fail()
        // Special case used very often is check equality
        assert().Equals(got, 1)
    }
*/
package asserter

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fail()
	FailNow()
	Log(...interface{})
	Logf(string, ...interface{})
}

type Asserter interface {
	T
	Equals(got, exp interface{}) T
	Contains(body, exp interface{}) T
	ResponseFrom(http.Handler) *HttpResponse
	Errors() (ok, bad AssertErrFunc)
	Mixed() (ok, bad MixedErrFunc)
}

// Assert returns an asserter for online assertions.
func New(t T) AssertFunc {
	return func(expr ...bool) Asserter {
		if len(expr) > 1 {
			t.Helper()
			t.Fatal("Only 0 or 1 bool expressions are allowed")
		}
		if len(expr) == 0 || !expr[0] {
			return &WrappedT{t}
		}
		return &noopT{}
	}
}

type AssertFunc func(expr ...bool) Asserter

type WrappedT struct {
	T
}

func (w *WrappedT) Helper() {
	/* Cannot use the asserter as helper */
}

func (w *WrappedT) Error(args ...interface{}) {
	w.T.Helper()
	w.T.Error(args...)
}

func (w *WrappedT) Errorf(format string, args ...interface{}) {
	w.T.Helper()
	w.T.Errorf(format, args...)
}

func (w *WrappedT) Fatal(args ...interface{}) {
	w.T.Helper()
	w.T.Fatal(args...)
}

func (w *WrappedT) Fatalf(format string, args ...interface{}) {
	w.T.Helper()
	w.T.Fatalf(format, args...)
}

func (w *WrappedT) Fail() {
	w.T.Helper()
	w.T.Fail()
}

func (w *WrappedT) FailNow() {
	w.T.Helper()
	w.T.FailNow()
}
func (w *WrappedT) Log(args ...interface{}) {
	w.T.Helper()
	w.T.Log(args...)
}

func (w *WrappedT) Logf(format string, args ...interface{}) {
	w.T.Helper()
	w.T.Logf(format, args...)
}

// Helpers

func (w *WrappedT) Equals(got, exp interface{}) T {
	w.T.Helper()
	if got != exp {
		w.Errorf("got %v, expected %v", got, exp)
	}
	return w.T
}

// Contains checks the body for the given expression.
// The body can be various types.
func (w *WrappedT) Contains(body, exp interface{}) T {
	w.T.Helper()
	b := toBytes(w.T, body, "body")
	e := toBytes(w.T, exp, "exp")

	if bytes.Index(b, e) == -1 {
		format := "%q does not contain %q"
		if bytes.Index(b, []byte("\n")) > -1 {
			format = "%s\ndoes not contain\n%s"
		}
		w.Errorf(format, string(b), string(e))
	}
	return w.T
}

// ----------------------------------------

func NewErrors(t T) (ok, bad AssertErrFunc) {
	t.Helper()
	return New(t)().Errors()
}

func (w *WrappedT) Errors() (ok, bad AssertErrFunc) {
	w.T.Helper()
	return assertOk(w.T), assertBad(w.T)
}

func assertOk(t T) AssertErrFunc {
	t.Helper()
	return func(err error) T {
		t.Helper()
		if err != nil {
			t.Error(err)
			return t
		}
		return &noopT{}
	}
}

func assertBad(t T) AssertErrFunc {
	t.Helper()
	return func(err error) T {
		t.Helper()
		if err == nil {
			t.Error("should fail")
			return t
		}
		return &noopT{}
	}
}

type AssertErrFunc func(error) T

// ----------------------------------------

func NewMixed(t T) (ok, bad MixedErrFunc) {
	t.Helper()
	return New(t)().Mixed()
}

func (w *WrappedT) Mixed() (ok, bad MixedErrFunc) {
	w.T.Helper()
	return mixOk(w.T), mixBad(w.T)
}

func mixOk(t T) MixedErrFunc {
	t.Helper()
	return func(any interface{}, err error) T {
		t.Helper()
		if err != nil {
			t.Error(err)
			return t
		}
		return &noopT{}
	}
}

func mixBad(t T) MixedErrFunc {
	t.Helper()
	return func(any interface{}, err error) T {
		t.Helper()
		if err == nil {
			t.Error("should fail")
			return t
		}
		return &noopT{}
	}
}

type MixedErrFunc func(interface{}, error) T

// ----------------------------------------

func (w *WrappedT) ResponseFrom(h http.Handler) *HttpResponse {
	return &HttpResponse{w.T, h}
}

func toBytes(t T, v interface{}, name string) (b []byte) {
	switch v := v.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	case int:
		return []byte(strconv.Itoa(v))
	case io.Reader:
		return bytesOrError(v)
	}
	t.Fatalf("%s must be io.Reader, []byte, string or int", name)
	return
}

func bytesOrError(r io.Reader) []byte {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte(err.Error())
	}
	return body
}

type noopT struct{}

func (t *noopT) Helper()                          {}
func (t *noopT) Error(...interface{})             {}
func (t *noopT) Errorf(string, ...interface{})    {}
func (t *noopT) Fatal(...interface{})             {}
func (t *noopT) Fatalf(string, ...interface{})    {}
func (t *noopT) Fail()                            {}
func (t *noopT) FailNow()                         {}
func (t *noopT) Log(...interface{})               {}
func (t *noopT) Logf(string, ...interface{})      {}
func (t *noopT) Equals(got, exp interface{}) T    { return t }
func (t *noopT) Contains(body, exp interface{}) T { return t }
func (t *noopT) ResponseFrom(h http.Handler) *HttpResponse {
	return &HttpResponse{t, h}
}
func (t *noopT) Errors() (ok, bad AssertErrFunc) {
	return func(error) T { return t },
		func(error) T { return t }
}

func (t *noopT) Mixed() (ok, bad MixedErrFunc) {
	return func(interface{}, error) T { return t },
		func(interface{}, error) T { return t }
}
