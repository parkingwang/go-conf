package conf

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 数值类型转换
//

func ConvertToString(val interface{}) string {
	switch val.(type) {
	case json.Number:
		return val.(json.Number).String()
	default:
		return AnyToStr(val)
	}
}

func ConvertToInt(val interface{}) int {
	switch val.(type) {
	case json.Number:
		v, _ := val.(json.Number).Int64()
		return int(v)

	case int:
		return val.(int)

	case int32:
		return int(val.(int32))

	case int64:
		return int(val.(int64))

	case float32:
		return int(val.(float32))

	case float64:
		return int(val.(float64))

	default:
		s := NonEmptyValueString(val)
		return int(StringMustInt64(s))
	}
}

func ConvertToInt32(val interface{}) int32 {
	switch val.(type) {
	case json.Number:
		i, _ := val.(json.Number).Int64()
		return int32(i)

	case int:
		return int32(val.(int))

	case int32:
		return val.(int32)

	case int64:
		return int32(val.(int64))

	case float32:
		return int32(val.(float32))

	case float64:
		return int32(val.(float64))

	default:
		s := NonEmptyValueString(val)
		return int32(StringMustInt64(s))
	}
}

func ConvertToInt64(val interface{}) int64 {
	switch val.(type) {
	case json.Number:
		i, _ := val.(json.Number).Int64()
		return i
	case int:
		return int64(val.(int))

	case int32:
		return int64(val.(int32))

	case int64:
		return val.(int64)

	case float32:
		return int64(val.(float32))

	case float64:
		return int64(val.(float64))

	default:
		s := NonEmptyValueString(val)
		return StringMustInt64(s)
	}
}

func ConvertToFloat32(val interface{}) float32 {
	switch val.(type) {
	case json.Number:
		i, _ := val.(json.Number).Float64()
		return float32(i)

	case float32:
		return val.(float32)

	case float64:
		return float32(val.(float64))

	case int:
		return float32(val.(int))

	case int32:
		return float32(val.(int32))

	case int64:
		return float32(val.(int64))

	default:
		s := NonEmptyValueString(val)
		return float32(StringMustFloat64(s))
	}
}

func ConvertToFloat64(val interface{}) float64 {
	switch val.(type) {
	case json.Number:
		i, _ := val.(json.Number).Float64()
		return i

	case float32:
		return float64(val.(float32))

	case float64:
		return val.(float64)

	case int:
		return float64(val.(int))

	case int32:
		return float64(val.(int32))

	case int64:
		return float64(val.(int64))

	default:
		s := NonEmptyValueString(val)
		return StringMustFloat64(s)
	}
}

func ConvertToBool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return val.(bool)

	case json.Number:
		i, _ := val.(json.Number).Int64()
		return 1 == i

	default:
		str := AnyToStr(val)
		return "true" == strings.ToLower(str) || "1" == str
	}
}

func AnyToStr(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func NonEmptyValueString(val interface{}) string {
	if s := AnyToStr(val); "" == s {
		return "0"
	} else {
		return s
	}
}

func StringMustInt64(s string) int64 {
	if v, e := strconv.ParseInt(s, 10, 64); nil != e {
		return 0
	} else {
		return v
	}
}

func StringMustFloat64(s string) float64 {
	if v, e := strconv.ParseFloat(s, 64); nil != e {
		return 0
	} else {
		return v
	}
}
