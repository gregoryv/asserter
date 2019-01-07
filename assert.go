/* package defines AssertFunc wrapper of testing.T

Online assertions are done by wrapping the T in a test

    func TestSomething(t *testing.T) {
        assert := asserter.New(t)
        got, err := Something()
        t.Logf("%v, %v := Something()", got, err)
        assert(err == nil).Fail()
        assert(got == exp).Fail()
    }
*/
package asserter

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Log(...interface{})
	Logf(string, ...interface{})
	Fail()
	FailNow()
	Skip(...interface{})
	SkipNow()
	Skipf(string, ...interface{})
}

type AssertFunc func(bool) T

type noopT struct{}

func (t *noopT) Helper()                       {}
func (t *noopT) Error(...interface{})          {}
func (t *noopT) Errorf(string, ...interface{}) {}
func (t *noopT) Fatal(...interface{})          {}
func (t *noopT) Fatalf(string, ...interface{}) {}
func (t *noopT) Log(...interface{})            {}
func (t *noopT) Logf(string, ...interface{})   {}
func (t *noopT) Fail()                         {}
func (t *noopT) FailNow()                      {}
func (t *noopT) Skip(...interface{})           {}
func (t *noopT) SkipNow()                      {}
func (t *noopT) Skipf(string, ...interface{})  {}

var ok *noopT = &noopT{}

// Assert returns an asserter for online assertions.
func New(t T) AssertFunc {
	return func(expr bool) T {
		if !expr {
			return t
		}
		return ok
	}
}
