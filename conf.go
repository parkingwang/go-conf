package conf

import (
	"os"
	"bytes"
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 加载配置文件
//

func LoadConfig(dirOrFile string) (Map, error) {
	if "" == dirOrFile {
		return nil, errors.New("Dir or file path is required")
	}

	fi, err := os.Stat(dirOrFile)
	if nil != err {
		return nil, errors.WithMessage(err, "Failed to get file/dir info")
	}

	var confBytes []byte
	if fi.IsDir() {
		if bs, err := LoadDirConfigText(dirOrFile); nil != err {
			return nil, err
		} else {
			confBytes = bs
		}
	} else {
		if bs, err := ioutil.ReadFile(dirOrFile); nil != err {
			return nil, errors.WithMessage(err, "Failed to read .toml config file")
		} else {
			confBytes = bs
		}
	}

	if tree, err := toml.LoadBytes(confBytes); nil != err {
		return nil, errors.WithMessage(err, "Failed to decode toml config file")
	} else {
		return tree.ToMap(), nil
	}

}

func LoadDirConfigText(dirName string) ([]byte, error) {
	out := new(bytes.Buffer)
	if files, err := ioutil.ReadDir(dirName); nil != err {
		return nil, errors.New("Failed to list file in dir: " + dirName)
	} else {
		if 0 == len(files) {
			return nil, errors.New("Config file NOT FOUND in dir: " + dirName)
		}
		for _, f := range files {
			name := f.Name()
			if !strings.HasSuffix(name, ".toml") {
				continue
			}
			path := fmt.Sprintf("%s%s%s", dirName, "/", f.Name())
			if bs, err := ioutil.ReadFile(path); nil != err {
				return nil, errors.New("Failed to load file: %s" + path)
			} else {
				out.Write(bs)
			}
		}
	}
	return out.Bytes(), nil
}
