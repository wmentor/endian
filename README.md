# endian

[![test](https://github.com/wmentor/endian/actions/workflows/ci.yml/badge.svg)](https://github.com/wmentor/endian/actions/workflows/ci.yml)
[![https://goreportcard.com/report/github.com/wmentor/endian](https://goreportcard.com/badge/github.com/wmentor/endian)](https://goreportcard.com/report/github.com/wmentor/endian)
[![https://pkg.go.dev/github.com/wmentor/endian](https://pkg.go.dev/badge/github.com/wmentor/endian.svg)](https://pkg.go.dev/github.com/wmentor/endian)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Detect byte order library (BigEndian or LittleEndian)

# Get bytes order

```go
package main

import (
	"fmt"

	"github.com/wmentor/endian"
)

func main() {
    order := endian.GetOrder()

    if order == endian.BigEndian {
        fmt.Println("Big Endian")
    } else {
        fmt.Println("Little Endian")
    }
}
```

Insted of *endian.BigEndian* you can use *endian.LittleEndian* constant.
