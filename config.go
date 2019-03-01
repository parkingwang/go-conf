package cfg

import (
	"errors"
	"strings"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 增加类型转换扩展的Config
//

type Config struct {
	data map[string]interface{}
}


// Wrap 将 map[String]interface{} 转换成 Config 对象
func Wrap(m map[string]interface{}) *Config {
	return &Config{data: m}
}

// 如果指定Key字段存在时，通过consumer函数处理返回值
func (cfg *Config) IfPresent(key string, consumer func(value interface{})) {
	if value, ok := cfg.data[key]; ok {
		consumer(value)
	}
}

// 获取Key的Value对象。
// 如果不存在，返回0值Value；
func (cfg *Config) MustValue(key string) Value {
	v, _ := cfg.GetValue(key)
	return v
}

// GetValue 获取Key的Value对象。
// 返回对象保证非空。并返回Key是否存在的状态值。
func (cfg *Config) GetValue(key string) (value Value, exist bool) {
	if rawValue, hit := cfg.data[key]; hit {
		if str, ok := rawValue.(string); ok {
			return Value(str), true
		} else {
			return Value(Value2String(rawValue)), true
		}
	} else {
		return Value("0"), false
	}
}

func (cfg *Config) IfPresentValue(key string, consumer func(value Value)) {
	if value, ok := cfg.GetValue(key); ok {
		consumer(value)
	}
}

// GetOrDefault 获取指定Key的值。
// 如果不存在返回默认值。
func (cfg *Config) GetOrDefault(key string, def interface{}) (interface{}, bool) {
	if val, ok := cfg.data[key]; ok {
		return val, ok
	} else {
		return def, ok
	}
}

// MustConfig 获取指定Key的值，类型为 Config。
// 如果不存在或者值类型不是map[string]interface{}时，返回空Config
func (cfg *Config) MustConfig(key string) *Config {
	return cfg.GetConfigOrDefault(key, &Config{})
}

// GetConfigOrDefault 获取指定Key的值，值类型为 Config。
// 如果不存在或者值类型不是map[string]interface{}时，返回指定的默认值Config
func (cfg *Config) GetConfigOrDefault(key string, def *Config) *Config {
	if ret, hit := cfg.GetOrDefault(key, def); hit {
		if mm, ok := ret.(map[string]interface{}); ok {
			return &Config{data: mm}
		} else {
			return def
		}
	} else {
		return def
	}
}

// MustConfigArray 获取指定Key的Config列表。
func (cfg *Config) MustConfigArray(key string) []*Config {
	return cfg.GetConfigArrayOrDefault(key, make([]*Config, 0))
}

// GetConfigArrayOrDefault 获取指定Key的值，值类型为 []Config。
// 如果存在以下情况时，返回指定的默认Config列表：
// 1. Key不存在；
// 2. Value不是数组；
// 3. Value数组条目不是map[string]interface{}
func (cfg *Config) GetConfigArrayOrDefault(key string, def []*Config) []*Config {
	if ret, hit := cfg.GetOrDefault(key, def); hit {
		if array, ok := ret.([]interface{}); ok {
			out := make([]*Config, 0)
			for _, item := range array {
				if mm, ok := item.(map[string]interface{}); ok {
					out = append(out, &Config{data: mm})
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
func (cfg *Config) MustString(key string) string {
	return cfg.GetStringOrDefault(key, "")
}

// GetStringOrDefault 获取指定Key的String值。
// 如果不存在，返回指定默认字符串。
func (cfg *Config) GetStringOrDefault(key string, def string) string {
	if val, hit := cfg.GetValue(key); !hit {
		return def
	} else {
		return val.String()
	}
}

// 返回指定Key的值是否与预期的相等
func (cfg *Config) IsFieldEqualToString(key string, except string) bool {
	return except == cfg.MustString(key)
}

// MustInt64 获取指定Key的 Int64 值。如果不存在，返回 0。
func (cfg *Config) MustInt64(key string) int64 {
	return cfg.GetInt64OrDefault(key, 0)
}

// GetInt64OrDefault 获取指定Key的 Int64 值。
// 如果不存在，返回指定的默认值。
func (cfg *Config) GetInt64OrDefault(key string, def int64) int64 {
	if val, hit := cfg.GetValue(key); !hit {
		return def
	} else {
		return val.Int64OrDefault(def)
	}
}

// 返回指定Key的值是否与预期的相等
func (cfg *Config) IsFieldEqualToInt64(key string, except int64) bool {
	return except == cfg.MustInt64(key)
}

// MustFloat64 获取指定Key的 Float64 值。
// 如果不存在，返回 0。
func (cfg *Config) MustFloat64(key string) float64 {
	return cfg.GetFloat64OrDefault(key, 0)
}

// GetFloat64OrDefault 获取指定Key的 Float64 值。
// 如果不存在，返回指定的默认值。
func (cfg *Config) GetFloat64OrDefault(key string, def float64) float64 {
	if val, hit := cfg.GetValue(key); !hit {
		return def
	} else {
		return val.Float64OrDefault(def)
	}
}

// MustBool 获取指定Key的 Bool 值。
// 如果不存在，返回 false。
func (cfg *Config) MustBool(key string) bool {
	return cfg.GetBoolOrDefault(key, false)
}

// GetBoolOrDefault 获取指定Key的 Bool 值。
// 如果不存在，返回指定的默认值。
func (cfg *Config) GetBoolOrDefault(key string, def bool) bool {
	if val, hit := cfg.GetValue(key); !hit {
		return def
	} else {
		return val.Bool()
	}
}

// MustDuration 获取指定Key的time.Duration值。
// 如果不存在，返回0。
func (cfg *Config) MustDuration(key string) time.Duration {
	return cfg.GetDurationOrDefault(key, 0)
}

// GetDurationOrDefault 获取指定Key的time.Duration值。
// 如果不存在或解析错误，返回指定默认值。
func (cfg *Config) GetDurationOrDefault(key string, def time.Duration) time.Duration {
	if val, hit := cfg.GetValue(key); !hit {
		return def
	} else {
		return val.DurationOrDefault(def)
	}
}

////

// EnsureString 获取指定Key的String值。如果不存在，抛出Panic错误。
func (cfg *Config) EnsureString(key string) string {
	str := cfg.GetStringOrDefault(key, "")
	if "" == str {
		panic("value of key <" + key + "> is empty")
	} else {
		return str
	}
}

// MustString2 获取并返回2个Key的值
func (cfg *Config) MustString2(a, b string) (x, y string) {
	x = cfg.MustString(a)
	y = cfg.MustString(b)
	return
}

// MustString3 获取并返回3个Key的值
func (cfg *Config) MustString3(a, b, c string) (x, y, z string) {
	x = cfg.MustString(a)
	y = cfg.MustString(b)
	z = cfg.MustString(c)
	return
}

// MustStringArray 获取并返回String数组
func (cfg *Config) MustStringArray(key string) ([]string, error) {
	out := make([]string, 0)
	if value, hit := cfg.GetOrDefault(key, out); hit {
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

// 转换成StringMap对象
func (cfg *Config) GetStringMapOrDefault(key string, def map[string]string) (map[string]string, error) {
	if v , hit := cfg.data[key]; !hit {
		return def, nil
	}else{
		switch v.(type) {
		case map[string]string:
			return v.(map[string]string), nil

		case map[string]interface{}:
			vmap := v.(map[string]interface{})
			omap := make(map[string]string, len(vmap))
			for k, v := range vmap {
				omap[k] = Value2String(v)
			}
			return omap, nil

		default:
			return nil, errors.New("value cannot convert to map[string]string: key=" + key)
		}
	}
}

//

func (cfg *Config) RefMap() map[string]interface{} {
	return cfg.data
}

// Contains 返回Config是否包含指定Key
func (cfg *Config) Contains(key string) bool {
	_, ok := cfg.data[key]
	return ok
}

// IsEmpty 返回当前Config是否为空
func (cfg *Config) IsEmpty() bool {
	return nil == cfg.data || 0 == len(cfg.data)
}

// IsNotEmpty 返回当前Config是否非空
func (cfg *Config) IsNotEmpty() bool {
	return !cfg.IsEmpty()
}

// ForEach KV
func (cfg *Config) ForEach(consumer func(name string, value interface{})) {
	for k, v := range cfg.data {
		consumer(k, v)
	}
}
