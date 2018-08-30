package conf

import (
	"errors"
	"strings"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// MustStringNotEmpty 获取指定Key的String值。如果不存在，返回Error。
func (m Map) MustStringNotEmpty(key string) (string, error) {
	str := m.GetStringOrDefault(key, "")
	if "" == str {
		return str, errors.New("value of key <" + key + "> is empty")
	} else {
		return str, nil
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

// MustStringArray 获取并返回String数组
func (m Map) MustStringArray(key string) ([]string, error) {
	out := make([]string, 0)
	if value, hit := m.GetOrDefault(key, out); hit {
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
