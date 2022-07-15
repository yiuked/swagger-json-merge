## 概述
该工具用于合并多个 `swagger json` 文件，工作原理，程序会读取 `base.json` 文件做为输出`json`文件的模板，所有在程序执行目录下的`json`文件都将合并到
`base.json` 文件中，合并的内容仅限于 `paths` 与 `definitions` 两个部分
## 安装
```
go install github.com/yiuked/swagger-json-merge
```