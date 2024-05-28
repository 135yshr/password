# password

[![go](https://github.com/135yshr/passowrd/actions/workflows/go.yml/badge.svg)](https://github.com/135yshr/passowrd/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/135yshr/passowrd)](https://goreportcard.com/report/github.com/135yshr/passowrd)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/135yshr/passowrd)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

_This is a simple password generator written in Go._

## Prerequisites

- [Go](https://go.dev/): 1.22+

## Installation

```bash
go get -u github.com/135yshr/passowrd
```

## Lean more

### How to generate


```go
import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewPasswordGenerator()
	fmt.Printf("Generated password: %s\n", gen.Generate(12))
}
```

## Contributing

This project is an open-source endeavor that thrives on your active participation. We're always on the lookout for individuals interested in contributing to the project's growth. If you have any ideas or improvements, no matter how small, they are welcome. Feel free to submit a [pull request](https://github.com/135yshr/password/pulls) at any time. We're eagerly awaiting your collaboration!

## License

This project is released under the MIT license. See the [LICENSE](LICENSE) file for details.
