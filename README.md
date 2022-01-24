# golang-project-template
> 本项目是通过CIL命令，一键创建一个golang项目的结构。

## 使用方法

### 查看版本
```bash
    golang-web version
```

### 创建项目
```bash
    golang-web create project
```

### 使用方法
#### 下载仓库
```bash
    git clone https://github.com/luenci/golang-project-template.git
```
#### 进入项目
```shell
➜ cd golang-project-template

➜ go build -o golang-web main.go

➜ ./golang-web version
  💻 golang-web version is 0.0.1

➜ ./golang-web create test
    🚀 Creating project test, please wait a moment.

    CREATED test (736 bytes)
    CREATED test/.gitignore (77 bytes)
    CREATED test/.pre-commit-config.yaml (712 bytes)
    CREATED test/CHANGELOG.md (0 bytes)
    CREATED test/Dockerfile (461 bytes)
    CREATED test/Makefile (365 bytes)
    CREATED test/README.md (0 bytes)
    CREATED test/config (64 bytes)
    CREATED test/docs (64 bytes)
    CREATED test/go.mod (62 bytes)
    CREATED test/go.sum (0 bytes)
    CREATED test/main.go (252 bytes)
    CREATED test/middleware (64 bytes)
    CREATED test/models (64 bytes)
    CREATED test/pkg (64 bytes)
    CREATED test/request (64 bytes)
    CREATED test/response (64 bytes)
    CREATED test/routers (64 bytes)
    CREATED test/runtime (64 bytes)
    CREATED test/service (64 bytes)
    CREATED test/tests (64 bytes)
    CREATED test/types (64 bytes)

    🍺 Project creation succeeded test
    💻 Use the following command to start the project 👇:

    $ cd test
                            🤝 Thanks for using golang-project-template
            📚 Tutorial: https://github.com/luenci/golang-project-template#readme

```

### 说明
- 默认的项目结构是为gin框架创建的，基于传统 MVC 架构而设计的。详情见：[gin-demo](https://github.com/Lucareful/gin-demo#readme)
- 我个人比较喜欢统一的代码风格和结构,所以将 `pre-commit` 这个配置文件设置为默认项。关于 `pre-commit` 的详细介绍，请参考：[pre-commit](https://pre-commit.com)

### TODO List
 - [ ] 发布一个版本
 - [ ] 增加单元测试s
 - [ ] 增加更多的项目结构选择
 - [ ] 增加更多的参数配置化
 - [ ] 目前代码写的比较丑陋，有时间优化下

### 鸣谢
  > 感谢以下项目的给我的启发
  - [kratos](https://github.com/go-kratos/kratos)
  - [go-micro](https://github.com/asim/go-micro)

### 后记
  > 如果你有什么想法，欢迎在issue中提出。
