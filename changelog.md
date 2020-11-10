# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]

- Add NotEqual
- Equals shows diff for values containing line breaks

## [0.4.0] - 2020-07-22
### Changed

- Equals uses reflect.DeepEqual
- New Testar interface combines T and Asserter
- Renamed A interface to Asserter
- Expose WrappedT struct using Wrap constructor
- AssertErrFunc signature returns T and doesn't take a message

## [0.3.0] - 2020-06-16
### Added

- MixedErrFunc via assert().Mixed() for mixed return checks
- AssertErrFunc for error only assertions
- BodyIs,  BodyEquals and Contains for body asserts
- HttpResponse.Header for checking single value headers

## [0.2.2] - 2020-05-02
## [0.2.1] - 2020-05-02
## [0.2.0] - 2020-05-02

- assert().ResponseFrom(handler) for quick http response assertions


## [0.1.0] - 2019-03-06
### Added

- Contains supports io.Reader
- Test wrapper with Contains and Equals
- Asserter type
