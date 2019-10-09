package env

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value bool

		exp bool
		err string
	}{
		{
			name: "Get 'true' from 't'",
			key:  "BOOL_KEY",
			env:  "TRUE",
			exp:  true,
		},
		{
			name: "Get 'true' from 'true'",
			key:  "BOOL_KEY",
			env:  "TRUE",
			exp:  true,
		},
		{
			name: "Get 'true' from 'TRUE'",
			key:  "BOOL_KEY",
			env:  "TRUE",
			exp:  true,
		},
		{
			name:  "Get default",
			key:   "BOOL_KEY",
			env:   "",
			value: true,
			exp:   true,
		},
		{
			name:  "Error on invalid",
			key:   "BOOL_KEY",
			env:   "invalid",
			value: true,
			exp:   true,
			err:   "env variable, 'BOOL_KEY' should be of type 'bool' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Bool(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp bool: '%t' got: '%t'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value time.Duration

		exp time.Duration
		err string
	}{
		{
			name: "Get 3 hours",
			key:  "DURATION_KEY",
			env:  "3h",
			exp:  3 * time.Hour,
		},
		{
			name:  "Get default",
			key:   "DURATION_KEY",
			env:   "",
			value: 4 * time.Minute,
			exp:   4 * time.Minute,
		},
		{
			name:  "Error on invalid",
			key:   "DURATION_KEY",
			env:   "invalid",
			value: 2 * time.Second,
			exp:   2 * time.Second,
			err:   "env variable, 'DURATION_KEY' should be of type 'time.Duration' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Duration(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp duration: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value float64

		exp float64
		err string
	}{
		{
			name: "Get 3",
			key:  "FLOAT64_KEY",
			env:  "3",
			exp:  3,
		},
		{
			name:  "Get default",
			key:   "FLOAT64_KEY",
			env:   "",
			value: 4,
			exp:   4,
		},
		{
			name:  "Error on invalid",
			key:   "FLOAT64_KEY",
			env:   "invalid",
			value: 2,
			exp:   2,
			err:   "env variable, 'FLOAT64_KEY' should be of type 'float64' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Float64(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp float64: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value int64

		exp int64
		err string
	}{
		{
			name: "Get 3",
			key:  "INT64_KEY",
			env:  "3",
			exp:  3,
		},
		{
			name: "Parse max int64",
			key:  "INT64_KEY",
			env:  "9223372036854775807",
			exp:  9223372036854775807,
		},
		{
			name:  "Get default",
			key:   "INT64_KEY",
			env:   "",
			value: 4,
			exp:   4,
		},
		{
			name:  "Error on invalid",
			key:   "INT64_KEY",
			env:   "invalid",
			value: 2,
			exp:   2,
			err:   "env variable, 'INT64_KEY' should be of type 'int64' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Int64(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp int64: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestInt(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value int

		exp int
		err string
	}{
		{
			name: "Get 3",
			key:  "INT_KEY",
			env:  "3",
			exp:  3,
		},
		{
			name: "Parse max int",
			key:  "INT_KEY",
			env:  "2147483647",
			exp:  2147483647,
		},
		{
			name:  "Get default",
			key:   "INT_KEY",
			env:   "",
			value: 4,
			exp:   4,
		},
		{
			name:  "Error on invalid",
			key:   "INT_KEY",
			env:   "invalid",
			value: 2,
			exp:   2,
			err:   "env variable, 'INT_KEY' should be of type 'int' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Int(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp int: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value string

		exp string
	}{
		{
			name: "Get 'str'",
			key:  "STRING_KEY",
			env:  "str",
			exp:  "str",
		},
		{
			name:  "Get default",
			key:   "STRING_KEY",
			env:   "",
			value: "default",
			exp:   "default",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			logErr(t, os.Setenv(tt.key, tt.env))

			got := String(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp string: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value uint64

		exp uint64
		err string
	}{
		{
			name: "Get 3",
			key:  "UINT64_KEY",
			env:  "3",
			exp:  3,
		},
		{
			name: "Parse max uint64",
			key:  "UINT64_KEY",
			env:  "18446744073709551615",
			exp:  18446744073709551615,
		},
		{
			name:  "Get default",
			key:   "UINT64_KEY",
			env:   "",
			value: 4,
			exp:   4,
		},
		{
			name:  "Error on invalid",
			key:   "UINT64_KEY",
			env:   "invalid",
			value: 2,
			exp:   2,
			err:   "env variable, 'UINT64_KEY' should be of type 'uint64' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Uint64(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp uint64: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestUint(t *testing.T) {
	var tests = []struct {
		name  string
		key   string
		env   string
		value uint

		exp uint
		err string
	}{
		{
			name: "Get 3",
			key:  "UINT_KEY",
			env:  "3",
			exp:  3,
		},
		{
			name: "Parse max uint",
			key:  "UINT_KEY",
			env:  "4294967295",
			exp:  4294967295,
		},
		{
			name:  "Get default",
			key:   "UINT_KEY",
			env:   "",
			value: 4,
			exp:   4,
		},
		{
			name:  "Error on invalid",
			key:   "UINT_KEY",
			env:   "invalid",
			value: 2,
			exp:   2,
			err:   "env variable, 'UINT_KEY' should be of type 'uint' but is 'invalid'",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			defer resetParserErrors()
			logErr(t, os.Setenv(tt.key, tt.env))

			got := Uint(tt.key, tt.value)
			if tt.exp != got {
				t.Errorf("%d. %s: Exp uint: '%v' got: '%v'", ti, tt.name, tt.exp, got)
			}

			err := errStr(ParserError())
			if tt.err != err {
				t.Errorf("%d. %s: Exp err: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
		})
	}
}

func TestMultipleErrors(t *testing.T) {
	defer resetParserErrors()

	exp1 := "env variable, 'FLOAT64_INVALID' should be of type 'float64' but is 'invalid_float'"
	exp2 := "env variable, 'INT64_INVALID' should be of type 'int64' but is 'invalid_int'"

	logErr(t, os.Setenv("FLOAT64_INVALID", "invalid_float"))
	logErr(t, os.Setenv("INT64_INVALID", "invalid_int"))

	_ = Float64("FLOAT64_INVALID", 0)
	_ = Int64("INT64_INVALID", 0)

	errors := ParserErrors()
	length := len(errors)
	if length != 2 {
		t.Fatalf("Exp errors length: '2' got: '%d'", length)
	}

	err1 := errStr(errors[0])
	if exp1 != err1 {
		t.Errorf("Exp err1: '%s' got: '%s'", exp1, err1)
	}

	err2 := errStr(errors[1])
	if exp2 != err2 {
		t.Errorf("Exp err2: '%s' got: '%s'", exp2, err2)
	}
}

func resetParserErrors() {
	parserErrors = []error{}
}

func errStr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func logErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
