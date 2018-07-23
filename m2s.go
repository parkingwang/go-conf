package conf

import "github.com/mitchellh/mapstructure"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// 将Map对象，转换成Struct结构体对象。
// 字段须声明`toml:"xxx_name"`
func Map2Struct(mapIn interface{}, structOut interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   structOut,
		TagName:  "toml",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(mapIn)
}
