[![Build Status](https://img.shields.io/travis/walle/fizzr.svg?style=flat)](https://travis-ci.org/walle/fizzr)
[![Coverage](https://img.shields.io/codecov/c/github/walle/fizzr.svg?style=flat)](https://codecov.io/github/walle/fizzr)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/walle/fizzr)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/walle/fizzr/master/LICENSE)
[![Go Report Card](http://goreportcard.com/badge/walle/fizzr?t=3)](http:/goreportcard.com/report/walle/fizzr)

# fizzr

Extensible fizz buzz library.

## Installation

```shell
$ go get github.com/walle/fizzr
```

## Usage

Setup a mappings slice with either supplied implementations or your own.
Use the Fizzr method to generate the string.

```go
mappings := []fizzr.Mapping{fizzr.Fizz,fizzr.Buzz}
output := fizzr.Fizzr(16, mappings) 
// 1, 2 Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16
```

## Testing

Use the `go test` tool.

```shell
$ go test -cover
```

```shell
$ go test -cover -bench=. -benchmem
```

## License

The code is under the MIT license. See [LICENSE](LICENSE) for more
information.
