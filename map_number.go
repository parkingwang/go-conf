package conf

import (
	"time"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// MustString 获取指定Key的String值。 如果不存在，返回空字符串。
func (m Map) MustString(key string) string {
	return m.GetStringOrDefault(key, "")
}

// GetStringOrDefault 获取指定Key的String值。如果不存在，返回指定默认字符串。
func (m Map) GetStringOrDefault(key string, def string) string {
	if val, hit := m.GetValue(key); !hit {
		return def
	} else {
		return val.String()
	}
}

func (m Map) IsFieldEqualToString(key string, except string) bool {
	return except == m.MustString(key)
}

// MustInt64 获取指定Key的 Int64 值。如果不存在，返回 0。
func (m Map) MustInt64(key string) int64 {
	return m.GetInt64OrDefault(key, 0)
}

// GetInt64OrDefault 获取指定Key的 Int64 值。 如果不存在，返回指定的默认值。
func (m Map) GetInt64OrDefault(key string, def int64) int64 {
	if val, hit := m.GetValue(key); !hit {
		return def
	} else {
		return val.Int64OrDefault(def)
	}
}

func (m Map) IsFieldEqualToInt64(key string, except int64) bool {
	return except == m.MustInt64(key)
}

// MustFloat64 获取指定Key的 Float64 值。如果不存在，返回 0。
func (m Map) MustFloat64(key string) float64 {
	return m.GetFloat64OrDefault(key, 0)
}

// GetFloat64OrDefault 获取指定Key的 Float64 值。 如果不存在，返回指定的默认值。
func (m Map) GetFloat64OrDefault(key string, def float64) float64 {
	if val, hit := m.GetValue(key); !hit {
		return def
	} else {
		return val.Float64OrDefault(def)
	}
}

// MustBool 获取指定Key的 Bool 值。如果不存在，返回 false。
func (m Map) MustBool(key string) bool {
	return m.GetBoolOrDefault(key, false)
}

// GetBoolOrDefault 获取指定Key的 Bool 值。 如果不存在，返回指定的默认值。
func (m Map) GetBoolOrDefault(key string, def bool) bool {
	if val, hit := m.GetValue(key); !hit {
		return def
	} else {
		return val.Bool()
	}
}

// MustDuration 获取指定Key的time.Duration值。如果不存在，返回0。
func (m Map) MustDuration(key string) time.Duration {
	return m.GetDurationOrDefault(key, 0)
}

// GetDurationOrDefault 获取指定Key的time.Duration值，如果不存在或解析错误，返回指定默认值。
func (m Map) GetDurationOrDefault(key string, def time.Duration) time.Duration {
	if val, hit := m.GetValue(key); !hit {
		return def
	} else {
		return val.DurationOrDefault(def)
	}
}
