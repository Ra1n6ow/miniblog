# 安装

```shell
go install github.com/air-verse/air@latest
```

# 配置

官方示例配置 https://github.com/air-verse/air/blob/master/air_example.toml
基于官方示例配置修改了如下：

```ini
# air在运行时存储临时文件的目录
tmp_dir = "/tmp/air"

[build]
# 普通的shell命令。你也可以使用`make`。
# cmd = "make build"
cmd = "go build -o _output/mb-apiserver -v cmd/mb-apiserver/main.go"
# 从`cmd`得到的二进制文件。
bin = "_output/mb-apiserver"
# 运行二进制文件（bin/full_bin）时添加额外参数，这里设置为空。
args_bin = []
```

# 启动

```shell
air -c .air.toml
```
