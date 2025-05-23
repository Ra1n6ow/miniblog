# 工具

```shell
go install github.com/nishanths/license/v5@latest

# 查看支持的代码协议
license -list

# 在项目根目录下执行，生成一个名为 LICENSE 的文件，该文件包含 MIT 开源协议声明
license -n 'ra1n6ow <jeffduuu@gmail.com>' -o LICENSE mit
```

# 源文件添加版本声明

将版权头信息保存在 scripts/boilerplate.txt，然后使用工具追加版权头信息

```shell
go install github.com/marmotedu/addlicense@latest

# 运行
addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,_output .
```
