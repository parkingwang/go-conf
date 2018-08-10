package conf

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 增加类型转换扩展的Map
//

type Map map[string]interface{}

// MapToMap 将Map[String]Any 转换成Map对象
func MapToMap(m map[string]interface{}) Map {
	return m
}

// GetValue 获取Key的Value对象，返回对象保证非空。并返回Key是否存在的状态值。
func (m Map) GetValue(key string) (value Value, exist bool) {
	if rawValue, hit := m[key]; hit {
		if str, ok := rawValue.(string); ok {
			return Value(str), true
		} else {
			return Value(Value2String(rawValue)), true
		}
	} else {
		return Value("0"), false
	}
}

// GetOrDefault 获取指定Key的值。 如果不存在返回默认值。
func (m Map) GetOrDefault(key string, def interface{}) (interface{}, bool) {
	if val, ok := m[key]; ok {
		return val, ok
	} else {
		return def, ok
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
