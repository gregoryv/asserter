[![Build Status](https://travis-ci.org/gregoryv/qual.svg?branch=master)](https://travis-ci.org/gregoryv/qual)
[![codecov](https://codecov.io/gh/gregoryv/qual/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/qual)
[asserter](https://godoc.org/github.com/gregoryv/asserter) - Go package oneline assertions

## Quick start

    go get github.com/gregoryv/asserter

In your tests

    func Test_something(t *testing.T) {
       assert := asserter.New(t)
       got, err := something()
       assert(err == nil).Fatal(err)
       assert(got == exp).Errorf("%v, expected %v", got, exp)
    }
