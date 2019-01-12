package conf

import (
	"errors"
	"strings"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 增加类型转换扩展的Map
//

type Map map[string]interface{}

type ImmutableMap struct {
	data map[string]interface{}
}

// MapToMap 将Map[String]Any 转换成Map对象
func MapToMap(m map[string]interface{}) *ImmutableMap {
	return WrapImmutableMap(m)
}

// MapToMap 将Map[String]Any 转换成ImmutableMap对象
func WrapImmutableMap(m map[string]interface{}) *ImmutableMap {
	return &ImmutableMap{data: m}
}

// 获取Key的Value对象。
// 如果不存在，返回0值Value；
func (im *ImmutableMap) MustValue(key string) Value {
	v, _ := im.GetValue(key)
	return v
}

// GetValue 获取Key的Value对象。
// 返回对象保证非空。并返回Key是否存在的状态值。
func (im *ImmutableMap) GetValue(key string) (value Value, exist bool) {
	if rawValue, hit := im.data[key]; hit {
		if str, ok := rawValue.(string); ok {
			return Value(str), true
		} else {
			return Value(Value2String(rawValue)), true
		}
	} else {
		return Value("0"), false
	}
}

// GetOrDefault 获取指定Key的值。
// 如果不存在返回默认值。
func (im *ImmutableMap) GetOrDefault(key string, def interface{}) (interface{}, bool) {
	if val, ok := im.data[key]; ok {
		return val, ok
	} else {
		return def, ok
	}
}

// MustMap 获取指定Key的值，类型为 Map。
// 如果不存在或者值类型不是map[string]interface{}时，返回空Map
func (im *ImmutableMap) MustMap(key string) Map {
	return im.GetMapOrDefault(key, Map{})
}

// GetMapOrDefault 获取指定Key的值，值类型为 Map。
// 如果不存在或者值类型不是map[string]interface{}时，返回指定的默认值Map
func (im *ImmutableMap) GetMapOrDefault(key string, def Map) Map {
	if ret, hit := im.GetOrDefault(key, def); hit {
		if mm, ok := ret.(map[string]interface{}); ok {
			return mm
		} else {
			return def
		}
	} else {
		return def
	}
}

// MustMap 获取指定Key的值，类型为 Map。
// 如果不存在或者值类型不是map[string]interface{}时，返回空Map
func (im *ImmutableMap) MustImmutableMap(key string) *ImmutableMap {
	return WrapImmutableMap(im.GetMapOrDefault(key, Map{}))
}

// MustMapArray 获取指定Key的Map列表。
func (im *ImmutableMap) MustMapArray(key string) []Map {
	return im.GetMapArrayOrDefault(key, make([]Map, 0))
}

// GetMapArrayOrDefault 获取指定Key的值，值类型为 []Map。
// 如果存在以下情况时，返回指定的默认Map列表：
// 1. Key不存在；
// 2. Value不是数组；
// 3. Value数组条目不是map[string]interface{}
func (im *ImmutableMap) GetMapArrayOrDefault(key string, def []Map) []Map {
	if ret, hit := im.GetOrDefault(key, def); hit {
		if array, ok := ret.([]interface{}); ok {
			out := make([]Map, 0)
			for _, item := range array {
				if mm, ok := item.(map[string]interface{}); ok {
					out = append(out, mm)
				}
			}
			return out
		} else {
			return def
		}
	} else {
		return def
	}
}

// MustString 获取指定Key的String值。
// 如果不存在，返回空字符串。
func (im *ImmutableMap) MustString(key string) string {
	return im.GetStringOrDefault(key, "")
}

// GetStringOrDefault 获取指定Key的String值。
// 如果不存在，返回指定默认字符串。
func (im *ImmutableMap) GetStringOrDefault(key string, def string) string {
	if val, hit := im.GetValue(key); !hit {
		return def
	} else {
		return val.String()
	}
}

func (im *ImmutableMap) IsFieldEqualToString(key string, except string) bool {
	return except == im.MustString(key)
}

// MustInt64 获取指定Key的 Int64 值。如果不存在，返回 0。
func (im *ImmutableMap) MustInt64(key string) int64 {
	return im.GetInt64OrDefault(key, 0)
}

// GetInt64OrDefault 获取指定Key的 Int64 值。
// 如果不存在，返回指定的默认值。
func (im *ImmutableMap) GetInt64OrDefault(key string, def int64) int64 {
	if val, hit := im.GetValue(key); !hit {
		return def
	} else {
		return val.Int64OrDefault(def)
	}
}

func (im *ImmutableMap) IsFieldEqualToInt64(key string, except int64) bool {
	return except == im.MustInt64(key)
}

// MustFloat64 获取指定Key的 Float64 值。
// 如果不存在，返回 0。
func (im *ImmutableMap) MustFloat64(key string) float64 {
	return im.GetFloat64OrDefault(key, 0)
}

// GetFloat64OrDefault 获取指定Key的 Float64 值。
// 如果不存在，返回指定的默认值。
func (im *ImmutableMap) GetFloat64OrDefault(key string, def float64) float64 {
	if val, hit := im.GetValue(key); !hit {
		return def
	} else {
		return val.Float64OrDefault(def)
	}
}

// MustBool 获取指定Key的 Bool 值。
// 如果不存在，返回 false。
func (im *ImmutableMap) MustBool(key string) bool {
	return im.GetBoolOrDefault(key, false)
}

// GetBoolOrDefault 获取指定Key的 Bool 值。
// 如果不存在，返回指定的默认值。
func (im *ImmutableMap) GetBoolOrDefault(key string, def bool) bool {
	if val, hit := im.GetValue(key); !hit {
		return def
	} else {
		return val.Bool()
	}
}

// MustDuration 获取指定Key的time.Duration值。
// 如果不存在，返回0。
func (im *ImmutableMap) MustDuration(key string) time.Duration {
	return im.GetDurationOrDefault(key, 0)
}

// GetDurationOrDefault 获取指定Key的time.Duration值。
// 如果不存在或解析错误，返回指定默认值。
func (im *ImmutableMap) GetDurationOrDefault(key string, def time.Duration) time.Duration {
	if val, hit := im.GetValue(key); !hit {
		return def
	} else {
		return val.DurationOrDefault(def)
	}
}

// MustStringNotEmpty 获取指定Key的String值。
// 如果不存在，返回Error。
func (im *ImmutableMap) MustStringNotEmpty(key string) (string, error) {
	str := im.GetStringOrDefault(key, "")
	if "" == str {
		return str, errors.New("value of key <" + key + "> is empty")
	} else {
		return str, nil
	}
}

// MustString2 获取并返回2个Key的值
func (im *ImmutableMap) MustString2(a, b string) (x, y string) {
	x = im.MustString(a)
	y = im.MustString(b)
	return
}

// MustString3 获取并返回3个Key的值
func (im *ImmutableMap) MustString3(a, b, c string) (x, y, z string) {
	x = im.MustString(a)
	y = im.MustString(b)
	z = im.MustString(c)
	return
}

// MustStringArray 获取并返回String数组
func (im *ImmutableMap) MustStringArray(key string) ([]string, error) {
	out := make([]string, 0)
	if value, hit := im.GetOrDefault(key, out); hit {
		switch value.(type) {

		case []interface{}:
			array := value.([]interface{})
			out = make([]string, len(array))
			for i, v := range array {
				out[i] = Value2String(v)
			}

		case []string:
			out = value.([]string)

		case string:
			out = strings.Split(value.(string), ",")

		default:
			return out, errors.New("value of key cannot convert to string array: " + key)
		}
	}
	return out, nil
}

// Contains 返回Map是否包含指定Key
func (im *ImmutableMap) Contains(key string) bool {
	_, ok := im.data[key]
	return ok
}

// IsEmpty 返回当前Map是否为空
func (im *ImmutableMap) IsEmpty() bool {
	return nil == im.data || 0 == len(im.data)
}

// IsNotEmpty 返回当前Map是否非空
func (im *ImmutableMap) IsNotEmpty() bool {
	return !im.IsEmpty()
}
