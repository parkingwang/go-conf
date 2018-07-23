package conf

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestConvertToBool(t *testing.T) {
	b0 := ConvertToBool(true)
	if !b0 {
		t.Error("value not matched")
	}
	b1 := ConvertToBool(false)
	if b1 {
		t.Error("value not matched")
	}
	b2 := ConvertToBool("True")
	if !b2 {
		t.Error("value not matched")
	}
	b3 := ConvertToBool("false")
	if b3 {
		t.Error("value not matched")
	}
	b4 := ConvertToBool("1")
	if !b4 {
		t.Error("value not matched")
	}
	b5 := ConvertToBool("0")
	if b5 {
		t.Error("value not matched")
	}
}

func TestConvertToInt(t *testing.T) {
	checkIfEquals := func(val int) {
		if 2018 != val {
			t.Errorf("value(%d) not matched [to int]", val)
		}
	}

	checkIfEquals(ConvertToInt(int(2018)))
	checkIfEquals(ConvertToInt(int32(2018)))
	checkIfEquals(ConvertToInt(int64(2018)))
	checkIfEquals(ConvertToInt(float32(2018)))
	checkIfEquals(ConvertToInt(float64(2018)))
	checkIfEquals(ConvertToInt("2018"))
}

func TestConvertToInt32(t *testing.T) {
	checkIfEquals := func(val int64) {
		if 2018 != val {
			t.Errorf("value(%d) not matched [to int64]", val)
		}
	}

	checkIfEquals(ConvertToInt64(int(2018)))
	checkIfEquals(ConvertToInt64(int32(2018)))
	checkIfEquals(ConvertToInt64(int64(2018)))
	checkIfEquals(ConvertToInt64(float32(2018)))
	checkIfEquals(ConvertToInt64(float64(2018)))
	checkIfEquals(ConvertToInt64("2018"))
}

func TestConvertToFloat32(t *testing.T) {
	checkIfEquals := func(val float32) {
		if 2018 != val {
			t.Errorf("value(%f) not matched [to float32]", val)
		}
	}

	checkIfEquals(ConvertToFloat32(int(2018)))
	checkIfEquals(ConvertToFloat32(int32(2018)))
	checkIfEquals(ConvertToFloat32(int64(2018)))
	checkIfEquals(ConvertToFloat32(float32(2018)))
	checkIfEquals(ConvertToFloat32(float64(2018)))
	checkIfEquals(ConvertToFloat32("2018"))
}

func TestConvertToFloat64(t *testing.T) {
	checkIfEquals := func(val float64) {
		if 2018 != val {
			t.Errorf("value(%f) not matched [to float64]", val)
		}
	}

	checkIfEquals(ConvertToFloat64(int(2018)))
	checkIfEquals(ConvertToFloat64(int32(2018)))
	checkIfEquals(ConvertToFloat64(int64(2018)))
	checkIfEquals(ConvertToFloat64(float32(2018)))
	checkIfEquals(ConvertToFloat64(float64(2018)))
	checkIfEquals(ConvertToFloat64("2018"))
}

func TestAnyToStr(t *testing.T) {
	checkIfEquals := func(txt string) {
		if "2018" != txt {
			t.Errorf("text(%s) not matched [any to string]", txt)
		}
	}
	checkIfEquals(AnyToStr(2018))
	checkIfEquals(AnyToStr(2018.0000))
	checkIfEquals(AnyToStr(int32(2018)))
	checkIfEquals(AnyToStr(int64(2018)))
	checkIfEquals(AnyToStr(float32(2018)))
	checkIfEquals(AnyToStr(float64(2018)))
	checkIfEquals(AnyToStr("2018"))
}
