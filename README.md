# miniblog 项目（参照 github.com/marmotedu）

# 添加 LICENSE 文件
使用 license 工具来生成
```shell
go install github.com/nishanths/license/v5@latest
license -list # 查看支持的代码协议
license -n 'ra1n6ow <jeffduuu@gmail.com>' -o LICENSE mit # 在 miniblog 项目根目录下执行
ls LICENSE 
LICENSE
```

# 给源文件添加版本声明
miniblog 的版权头信息保存在 scripts/boilerplate.txt 文件中。
```shell
go install github.com/marmotedu/addlicense@latest
# 给当前目录下的所有文件添加头信息
addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,vendor,_output .
```

# air 动态构建运行
```shell
go install github.com/cosmtrek/air@latest

# 运行
air [-c .air.toml]
```

# 根据数据表生成结构体
```shell
# 安装db2struct 工具
go install github.com/Shelnutt2/db2struct/cmd/db2struct

# 进入生成文件的目录 cd miniblog/internal/pkg/model
db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t user --package model --struct UserM -u root -p '123456' --target=user.go
db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t post --package model --struct PostM -u root -p '123456' --target=post.go
```