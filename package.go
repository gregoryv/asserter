/*
Package asserter defines helpers for asserting various outputs.

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
