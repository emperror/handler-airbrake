# Airbrake / Errbit handler

[![CircleCI](https://circleci.com/gh/emperror/handler-airbrake.svg?style=svg)](https://circleci.com/gh/emperror/handler-airbrake)
[![Go Report Card](https://goreportcard.com/badge/handler.emperror.dev/airbrake?style=flat-square)](https://goreportcard.com/report/handler.emperror.dev/airbrake)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-airbrake.svg)](https://golangci.com/r/github.com/emperror/handler-airbrake)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/handler.emperror.dev/airbrake)

**Error handler integration for [Airbrake](https://airbrake.com) / [Errbit](https://errbit.com/).**


## Installation

```bash
go get handler.emperror.dev/airbrake
```


## Usage

```go
package main

import (
	"github.com/airbrake/gobrake"

	"handler.emperror.dev/airbrake"
)

func main() {
    projectID := int64(1)
	projectKey := "key"

	handler := airbrake.New(projectID, projectKey)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
