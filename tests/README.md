# gin-web

## 前言
本项目目录参考自 煎鱼 系列文章 《Gin搭建Blog API's》我重构了整体的目录结构和代码。

## 目录结构
```shell
config  	- 配置文件
docs    	- 接口文档
middleware 	- 中间件
models		- 模型 struct（model层）
routers		- 路由（view层）
service		- 服务处理层（controller 逻辑处理层）
runtime		- 运行时
pkg			- 包
tests		- 测试文件
types		- 定义接口数据结构struct
	request	- 请求体的定义
	response- 响应体的定义
go.sum
CHANGELOG.md
README.md
```

## 更多详情
- [gin-demo](https://github.com/Lucareful/gin-demo)
