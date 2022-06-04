# environ

[![codecov](https://codecov.io/gh/Akshit8/environ/branch/main/graph/badge.svg?token=DK9XSIA03T)](https://codecov.io/gh/Akshit8/environ)

a simple package to inject environment variables into golang structs

## Installation

```
go get -u github.com/Akshit8/environ
```

## Quickstart

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/Akshit8/environ"
)

type Config struct {
	Host  string `environ:"APP_HOST"`
	Port  int    `environ:"APP_PORT"`
	Debug bool   `environ:"APP_DEBUG"`
}

func IntParser(s string) (int, error) {
	return strconv.Atoi(s)
}

func BoolParser(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func main() {
	var config Config

	environ.UseParser(IntParser)
	environ.UseParser(BoolParser)

	err := environ.Inject(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", config)
}
```
