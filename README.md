# password

[![go](https://github.com/135yshr/password/actions/workflows/go.yml/badge.svg)](https://github.com/135yshr/password/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/135yshr/password)](https://goreportcard.com/report/github.com/135yshr/password)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/135yshr/password)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

_This is a simple password generator written in Go._

## Getting Started

```bash
brew install 135yshr/tap/mkpass
```

### Run

```bash
$ mkpass -h
Usage of mkpass:
  -a	use all characters (default true)
  -all
    	use all characters (default true)
  -custom string
    	use custom characters
  -h	show help
  -help
    	show help
  -l	use lowercase
  -length int
    	length of password (default 12)
  -lower
    	use lowercase
  -n	use numbers
  -number
    	use numbers
  -s	use symbols
  -symbol
    	use symbols
  -u	use uppercase
  -upper
    	use uppercase
  -v	show version
  -version
    	show version
```


## Prerequisites

- [Go](https://go.dev/): 1.22+

## Development

```bash
$ go get -u github.com/135yshr/password
```

## Lean more

### How to generate

```go
import (
	"fmt"

	"github.com/135yshr/password"
)

func main() {
	gen := password.New()
	fmt.Printf("Generated password: %s\n", gen.Generate(12))
}
```

## Contributing

This project is an open-source endeavor that thrives on your active participation. We're always on the lookout for individuals interested in contributing to the project's growth. If you have any ideas or improvements, no matter how small, they are welcome. Feel free to submit a [pull request](https://github.com/135yshr/password/pulls) at any time. We're eagerly awaiting your collaboration!

## License

This project is released under the MIT license. See the [LICENSE](LICENSE) file for details.
