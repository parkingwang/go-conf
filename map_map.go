package conf

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// MustMap 获取指定Key的值，类型为 Map。 如果不存在或者值类型不是map[string]interface{}时，返回空Map
func (m Map) MustMap(key string) Map {
	return m.GetMapOrDefault(key, Map{})
}

// GetMapOrDefault 获取指定Key的值，值类型为 Map。 如果不存在或者值类型不是map[string]interface{}时，返回指定的默认值Map
func (m Map) GetMapOrDefault(key string, def Map) Map {
	if ret, hit := m.GetOrDefault(key, def); hit {
		if mm, ok := ret.(map[string]interface{}); ok {
			return mm
		} else {
			return def
		}
	} else {
		return def
	}
}

// MustMapArray 获取指定Key的Map列表。
func (m Map) MustMapArray(key string) []Map {
	return m.GetMapArrayOrDefault(key, make([]Map, 0))
}

// GetMapArrayOrDefault 获取指定Key的值，值类型为 []Map。
// 如果存在以下情况时，返回指定的默认Map列表：
// 1. Key不存在；
// 2. Value不是数组；
// 3. Value数组条目不是map[string]interface{}
func (m Map) GetMapArrayOrDefault(key string, def []Map) []Map {
	if ret, hit := m.GetOrDefault(key, def); hit {
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
