package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
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
		i, _ := val.(json.Number).Int64()
		return int(i)

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
		i, e := strconv.Atoi(AnyToStr(val))
		if nil != e {
			panic(e)
		}
		return i
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
		i, e := strconv.ParseInt(AnyToStr(val), 10, 64)
		if nil != e {
			panic(e)
		}
		return int32(i)
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
		i, e := strconv.ParseInt(AnyToStr(val), 10, 64)
		if nil != e {
			panic(e)
		}
		return i
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
		i, e := strconv.ParseFloat(AnyToStr(val), 32)
		if nil != e {
			panic(e)
		}
		return float32(i)
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
		i, e := strconv.ParseFloat(AnyToStr(val), 64)
		if nil != e {
			panic(e)
		}
		return i
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

func ConvertToBufferedReader(source interface{}) (*bufio.Reader, error) {
	var bufReader *bufio.Reader
	switch source.(type) {
	case *os.File:
		bufReader = bufio.NewReader(source.(*os.File))

	case string:
		bufReader = bufio.NewReader(strings.NewReader(source.(string)))

	case io.Reader:
		bufReader = bufio.NewReader(source.(io.Reader))

	case *bufio.Reader:
		bufReader = source.(*bufio.Reader)

	default:
		return nil, errors.Errorf("unsupported source type, was: %t", source)
	}

	return bufReader, nil
}
