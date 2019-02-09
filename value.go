package cfg

import (
	"strconv"
	"time"
	"strings"
	"fmt"
)

// Value是字符串数据的别名
type Value string

// 返回String值
func (v Value) String() string { return string(v) }

// 返回解析Float64数值
func (v Value) Float64() (float64, error) {
	return strconv.ParseFloat(v.String(), 64)
}

// 返回解析Float64数值，或指定默认值
func (v Value) Float64OrDefault(defaultValue float64) float64 {
	if i, e := v.Float64(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

// 返回解析Int64数值
func (v Value) Int64() (int64, error) {
	return strconv.ParseInt(v.String(), 10, 64)
}

// 返回解析Int64数值，或指定默认值
func (v Value) Int64OrDefault(defaultValue int64) int64 {
	if i, e := v.Int64(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

// 返回时差Duration值
func (v Value) Duration() (time.Duration, error) {
	return time.ParseDuration(v.String())
}

// 返回时差Duration值，或指定默认值
func (v Value) DurationOrDefault(defaultValue time.Duration) time.Duration {
	if i, e := v.Duration(); nil != e {
		return defaultValue
	} else {
		return i
	}
}

// 返回Boolean值
func (v Value) Bool() bool {
	s := strings.ToLower(v.String())
	return "true" == s || "1" == s
}

// 将Any类型，转换成String
func Value2String(value interface{}) string {
	return ToString(value)
}

// 将Any类型，转换成String
func ToString(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)

	case int:
		i := int64(val.(int))
		return strconv.FormatInt(i, 10)

	case int32:
		i := int64(val.(int32))
		return strconv.FormatInt(i, 10)

	case int64:
		return strconv.FormatInt(val.(int64), 10)

	case float32:
		f := float64(val.(float32))
		return strconv.FormatFloat(f, 'E', -1, 32)

	case float64:
		return strconv.FormatFloat(val.(float64), 'E', -1, 64)

	default:
		return fmt.Sprintf("%v", val)
	}
}
