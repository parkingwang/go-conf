package conf

import (
	"strconv"
	"strings"
	"time"
)

type Value string

func (v Value) String() string { return string(v) }

func (v Value) Float64() (float64, error) {
	return strconv.ParseFloat(v.String(), 64)
}

func (v Value) Float64OrDefault(defaultValue float64) float64 {
	if i, e := v.Float64(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

func (v Value) Int64() (int64, error) {
	return strconv.ParseInt(v.String(), 10, 64)
}

func (v Value) Int64OrDefault(defaultValue int64) int64 {
	if i, e := v.Int64(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

func (v Value) Duration() (time.Duration, error) {
	return time.ParseDuration(v.String())
}

func (v Value) DurationOrDefault(defaultValue time.Duration) time.Duration {
	if i, e := v.Duration(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

func (v Value) Bool() bool {
	s := strings.ToLower(v.String())
	return "true" == s || "1" == s
}
