// Package env allows to parse environment variable's to a specific type or specify a default value.
// This package provides some of the functionality of the go flag package for environmental variables.
package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var parserErrors = []error{}

// Bool reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Bool(name string, value bool) bool {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseBool(env)
	if err != nil {
		appendError("bool", name, env)
		return value
	}
	return val
}

// Duration reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Duration(name string, value time.Duration) time.Duration {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := time.ParseDuration(env)
	if err != nil {
		appendError("time.Duration", name, env)
		return value
	}
	return val
}

// Float64 reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Float64(name string, value float64) float64 {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseFloat(env, 64)
	if err != nil {
		appendError("float64", name, env)
		return value
	}
	return val
}

// Int64 reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Int64(name string, value int64) int64 {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		appendError("int64", name, env)
		return value
	}
	return val
}

// Int reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Int(name string, value int) int {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseInt(env, 10, 32)
	if err != nil {
		appendError("int", name, env)
		return value
	}
	return int(val)
}

// String reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func String(name, value string) string {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	return env
}

// Uint64 reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Uint64(name string, value uint64) uint64 {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseUint(env, 10, 64)
	if err != nil {
		appendError("uint64", name, env)
		return value
	}
	return val
}

// Uint reads and parses the specified environment variable.
// If the variable is not present the given value will be returned.
func Uint(name string, value uint) uint {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	val, err := strconv.ParseUint(env, 10, 32)
	if err != nil {
		appendError("uint", name, env)
		return value
	}
	return uint(val)
}

// ParserError returns the first parser error or nil if no errors are present.
func ParserError() error {
	if len(parserErrors) > 0 {
		return parserErrors[0]
	}
	return nil
}

// ParserErrors returns a slice of all parser errors.
func ParserErrors() []error {
	return parserErrors
}

func appendError(Type, name, value string) {
	err := fmt.Errorf("env variable, '%s' should be of type '%s' but is '%s'", name, Type, value)
	parserErrors = append(parserErrors, err)
}
