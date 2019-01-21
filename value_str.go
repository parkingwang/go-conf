package cfg

import (
	"fmt"
	"strconv"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func Value2String(val interface{}) string {
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
