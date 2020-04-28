# Airbrake / Errbit handler

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/handler-airbrake/CI?style=flat-square)](https://github.com/emperror/handler-airbrake/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/handler-airbrake?style=flat-square)](https://codecov.io/gh/emperror/handler-airbrake)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/airbrake?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/airbrake)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/emperror.dev/handler/airbrake)


**Error handler integration for [Airbrake](https://airbrake.com) / [Errbit](https://errbit.com/).**


## Installation

```bash
go get emperror.dev/handler/airbrake
```


## Usage

```go
package main

import (
	"github.com/airbrake/gobrake"

	"emperror.dev/handler/airbrake"
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


### Running integration tests

```bash
# Set up the environment
cp docker-compose.override.yml.dist docker-compose.override.yml
cp .env.test.dist .env.test
docker-compose up -d
docker-compose run --rm errbit rake db:seed
source .env.test

# Run tests
make test

# Cleanup
docker-compose down
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
