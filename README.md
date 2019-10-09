# Env
[![Go Report Card](https://goreportcard.com/badge/github.com/jksch/env)](https://goreportcard.com/report/github.com/jksch/env)
[![GoDoc](https://godoc.org/github.com/jksch/env?status.svg)](https://godoc.org/github.com/jksch/env)
[![License](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/jksch/env/blob/master/LICENSE)

Is a simple environment variable parser inspired by the go flag package.

Env allows you to simply parse an environment variable or specify a default.

## Basic usage

Assume this environmental variables are set:
```
DB_USER=admin
DB_PASS=password
```
You can do the following:
```
package main

import (
	"fmt"
	"time"

	"github.com/jksch/env"
)

var dbHost = env.String("DB_HOST", "localhost")
var dbPort = env.Int("DB_PORT", 27017)
var dbUser = env.String("DB_USER", "")
var dbPass = env.String("DB_PASS", "")
var dbTimeout = env.Duration("DB_TIMEOUT", 15*time.Second)

func main() {
	if err := env.ParserError(); err != nil {
		fmt.Printf("Could not parse env variables, %v", err)
	}
	fmt.Printf("dbHost: %v\n", dbHost)
	fmt.Printf("dbPort: %v\n", dbPort)
	fmt.Printf("dbUser: %v\n", dbUser)
	fmt.Printf("dbPass: %v\n", dbPass)
	fmt.Printf("dbTimeout: %v\n", dbTimeout)
}
```
This will parse two environmental variables and load three defaults:
```
dbHost: localhost
dbPort: 27017
dbUser: admin
dbPass: password
dbTimeout: 15s
```
