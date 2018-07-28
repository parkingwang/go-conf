# Util for Toml config

用于读取Toml配置文件，将其转换成Map对象。Map对象提供一些类型转换函数。

## Install

```bash
go get -u github.com/parkingwang/go-conf
```

OR

```bash
dep ensure -add github.com/parkingwang/go-conf
```

## Usage

### Map对象

conf.Map 是 `map[string]interface{}` 类型的别名。它主要是提供了一系统Get函数，返回特定类型。

详细见 **MustXX** 函数和 **GetXXOrDefault** 函数。

### LoadConfig 加载TOML配置文件

- *LoadConfig* 加载指定配置文件夹名称或者TOML文件路径，返回全部配置文件的Map对象；
- *LoadDirConfigText* 加载指定TOML配置文件目录，返回所有配置文件的合并Text文本； 

### Map2Struct Map转Struct

将Map对象，转换成Struct结构体对象。字段须声明`toml:"xxx_name"`。