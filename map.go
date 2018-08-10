package conf

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 增加类型转换扩展的Map
//

type Map map[string]interface{}

// MapToMap 将Map[String]Any 转换成Map对象
func MapToMap(m map[string]interface{}) Map {
	return m
}

// GetValueOrDefault 获取指定Key的值。 如果不存在返回默认值。
func (m Map) GetValueOrDefault(key string, def interface{}) (interface{}, bool) {
	if val, ok := m[key]; ok {
		return val, ok
	} else {
		return def, ok
	}
}

// MustStrMap 确保返回一个Str-Map对象，或者默认空对象。
func (m Map) MustStrMap(key string) (map[string]string, error) {
	if value, ok := m[key]; ok {
		if found, ok := value.(map[string]string); ok {
			return found, nil
		} else {
			out := make(map[string]string)
			if sm, ok := value.(map[string]interface{}); ok {
				for k, v := range sm {
					out[k] = ConvertToString(v)
				}
				return out, nil
			} else {
				msg := fmt.Sprintf("value of %s is not map[string]string/interface{}, was: %s", key, reflect.TypeOf(value))
				return out, errors.New(msg)
			}
		}
	} else {
		return make(map[string]string), nil
	}
}

// MustMap 获取指定Key的值，确保返回值类型为 Map。 如果不存在，返回空Map
func (m Map) MustMap(key string) Map {
	return m.GetMapOrDefault(key, Map{})
}

// GetMapOrDefault 获取指定Key的值，值类型为 Map。
// 如果不存在，返回指定的默认值Map
func (m Map) GetMapOrDefault(key string, def Map) Map {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		if found, ok := ret.(map[string]interface{}); ok {
			return found
		} else {
			out := Map{}
			for k, v := range ret.(map[interface{}]interface{}) {
				out[AnyToStr(k)] = v
			}
			return out
		}
	} else {
		return def
	}
}

func (m Map) MustArrayMap(key string) []Map {
	return m.GetArrayMapOrDefault(key, make([]Map, 0))
}

// GetArrayMapValue 获取指定Key的值，值类型为 []Map。
// 如果不存在，返回空Map列表
func (m Map) GetArrayMapValue(key string) []Map {
	return m.MustArrayMap(key)
}

// GetArrayMapOrDefault 获取指定Key的值，值类型为 []Map。
// 如果不存在，返回指定的默认Map列表
func (m Map) GetArrayMapOrDefault(key string, def []Map) []Map {
	if array, ok := m.GetValueOrDefault(key, def); ok {
		out := make([]Map, 0)
		for _, item := range array.([]interface{}) {
			out = append(out, item.(map[string]interface{}))
		}
		return out
	} else {
		return def
	}
}

// MustString 获取指定Key的String值。 如果不存在，返回空字符串。
func (m Map) MustString(key string) string {
	return m.GetStringOrDefault(key, "")
}

// MustStringNotEmpty 获取指定Key的String值。如果不存在，返回Error。
func (m Map) MustStringNotEmpty(key string) (string, error) {
	s := m.GetStringOrDefault(key, "")
	if "" == s {
		return s, errors.New("value of key <" + key + "> is empty")
	} else {
		return s, nil
	}
}

// MustString2 获取并返回2个Key的值
func (m Map) MustString2(a, b string) (x, y string) {
	x = m.MustString(a)
	y = m.MustString(b)
	return
}

// MustString3 获取并返回3个Key的值
func (m Map) MustString3(a, b, c string) (x, y, z string) {
	x = m.MustString(a)
	y = m.MustString(b)
	z = m.MustString(c)
	return
}

// GetStringOrDefault 获取指定Key的String值。
// 如果不存在，返回指定默认字符串。
func (m Map) GetStringOrDefault(key string, def string) string {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToString(ret)
	} else {
		return def
	}
}

// MustInt 获取指定Key的 Int 值。 如果不存在，返回 0。
func (m Map) MustInt(key string) int {
	return m.GetIntOrDefault(key, 0)
}

// GetIntOrDefault 获取指定Key的 Int 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetIntOrDefault(key string, def int) int {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToInt(ret)
	} else {
		return def
	}
}

// MustInt32 // 获取指定Key的 Int32
func (m Map) MustInt32(key string) int32 {
	return m.GetInt32OrDefault(key, 0)
}

// GetInt32OrDefault 获取指定Key的 Int32 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetInt32OrDefault(key string, def int32) int32 {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToInt32(ret)
	} else {
		return def
	}
}

// MustInt64 获取指定Key的 Int64 值。如果不存在，返回 0。
func (m Map) MustInt64(key string) int64 {
	return m.GetInt64OrDefault(key, 0)
}

// GetInt64OrDefault 获取指定Key的 Int64 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetInt64OrDefault(key string, def int64) int64 {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToInt64(ret)
	} else {
		return def
	}
}

// MustFloat32 获取指定Key的 Float32 值。如果不存在，返回 0。
func (m Map) MustFloat32(key string) float32 {
	return m.GetFloat32OrDefault(key, 0)
}

// GetFloat32OrDefault 获取指定Key的 Float32 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetFloat32OrDefault(key string, def float32) float32 {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToFloat32(ret)
	} else {
		return def
	}
}

// MustFloat64 获取指定Key的 Float64 值。如果不存在，返回 0。
func (m Map) MustFloat64(key string) float64 {
	return m.GetFloat64OrDefault(key, 0)
}

// GetFloat64OrDefault 获取指定Key的 Float64 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetFloat64OrDefault(key string, def float64) float64 {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToFloat64(ret)
	} else {
		return def
	}
}

// MustBool 获取指定Key的 Bool 值。如果不存在，返回 false。
func (m Map) MustBool(key string) bool {
	return m.GetBoolOrDefault(key, false)
}

// GetBoolOrDefault 获取指定Key的 Bool 值。
// 如果不存在，返回指定的默认值。
func (m Map) GetBoolOrDefault(key string, def bool) bool {
	if ret, ok := m.GetValueOrDefault(key, def); ok {
		return ConvertToBool(ret)
	} else {
		return def
	}
}

// MustDuration 获取指定Key的time.Duration值。如果不存在，返回0。
// 如果Key所指向的值无法被解析成Duration，将引发Panic。
func (m Map) MustDuration(key string) time.Duration {
	return m.GetDurationOrDefault(key, 0)
}

// GetDurationOrDefault 获取指定Key的time.Duration值，如果不存在或解析错误，返回指定默认值。
func (m Map) GetDurationOrDefault(key string, defaultT time.Duration) time.Duration {
	str := m.MustString(key)
	if "" == str {
		return defaultT
	}
	if duration, err := time.ParseDuration(str); nil != err {
		return defaultT
	} else {
		return duration
	}
}

// Contains 返回Map是否包含指定Key
func (m Map) Contains(key string) bool {
	_, ok := m[key]
	return ok
}

// IsEmpty 返回当前Map是否为空
func (m Map) IsEmpty() bool {
	return 0 == len(m)
}

// IsNotEmpty 返回当前Map是否非空
func (m Map) IsNotEmpty() bool {
	return !m.IsEmpty()
}
