[![Build Status](https://travis-ci.org/gregoryv/asserter.svg?branch=master)](https://travis-ci.org/gregoryv/asserter)
[![codecov](https://codecov.io/gh/gregoryv/asserter/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/asserter)
[![Maintainability](https://api.codeclimate.com/v1/badges/b0001c5ba7cd098b183d/maintainability)](https://codeclimate.com/github/gregoryv/asserter/maintainability)

[asserter](https://godoc.org/github.com/gregoryv/asserter) - Package for oneline assertions

## Quick start

    go get github.com/gregoryv/asserter

In your tests

    func Test_something(t *testing.T) {
        got, err := something()

        assert := asserter.New(t)
        assert(err == nil).Fatal(err)
        assert(got == exp).Errorf("%v, expected %v", got, exp)
	    // same as
	    assert().Equals(got, exp)

        assert().Contains(got, "text")
	    assert().Contains(got, 1)

	    // Check readers content
	    resp, err := http.Get("http://example.com")
	    assert(err == nil).Fatal(err)
	    assert().Contains(resp.Body, "<title>")
    }


HTTP handler specific

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
