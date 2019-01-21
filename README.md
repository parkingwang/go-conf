# Util for config

用于读取Toml配置文件，将其转换成 Config 对象。Config 对象提供一些类型转换函数。

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

cfg.Config 包装`map[string]interface{}`，提供一个不可变访问接口，通过 GetXXX, MustXXX 等函数来读取内部数据。

详细见 **MustXX** 函数和 **GetXXOrDefault** 函数。

![Config](Config.png)

### LoadConfig 加载TOML配置文件

- *LoadConfig* 加载指定配置文件夹名称或者TOML文件路径，返回全部配置文件的Map对象；
- *LoadDirConfigText* 加载指定TOML配置文件目录，返回所有配置文件的合并Text文本； 
