# gpt(golang-project-template)
- go 版本 > 1.16
> 本项目是通过CIL命令，一键创建一个 golang web 项目的结构。


## 使用方法

### 查看版本
```bash
    gpt version
```

### 创建项目

#### gin web 项目结构(默认结构)
```bash
    gpt create -p gin test
```

#### DDD 项目结构
```bash
    gpt create -p DDD tests
```

### 使用方法
#### go get 安装
```bash
    go get -u github.com/luenci/gpt
```

##### 验证版本
```shell
    gpt version
```

#### 源码安装
##### 下载仓库
```bash
    git clone https://github.com/luenci/gpt.git
```
##### 进入项目
```shell
➜ cd gpt

➜ go build -o gpt main.go

➜ ./gpt version
  💻 gpt version is 0.0.1

➜ > ./gpt create -p gin tests
   🚀 Creating project tests, please wait a moment.

   CREATED tests (512 bytes)
   CREATED tests/.gitignore (78 bytes)
   CREATED tests/.pre-commit-config.yaml (711 bytes)
   CREATED tests/CHANGELOG.md (0 bytes)
   CREATED tests/Dockerfile (460 bytes)
   CREATED tests/Makefile (366 bytes)
   CREATED tests/README.md (645 bytes)
   CREATED tests/config (512 bytes)
   CREATED tests/docs (512 bytes)
   CREATED tests/main.go (253 bytes)
   CREATED tests/middleware (512 bytes)
   CREATED tests/models (512 bytes)
   CREATED tests/pkg (512 bytes)
   CREATED tests/routers (512 bytes)
   CREATED tests/runtime (512 bytes)
   CREATED tests/service (512 bytes)
   CREATED tests/tests (512 bytes)
   CREATED tests/types (512 bytes)
   CREATED tests/types/request (512 bytes)
   CREATED tests/types/response (512 bytes)

   🍺 Project creation succeeded tests
   💻 Use the following command to start the project 👇:

   $ cd tests
                        🤝 Thanks for using gpt (golang-project-template)
        📚 Tutorial: https://github.com/luenci/gpt#readme

```

### 说明
- 默认的项目结构是为gin框架创建的，基于传统 MVC 架构而设计的。详情见：[gin-demo](https://github.com/Lucareful/gin-demo#readme)
- 我个人比较喜欢统一的代码风格和结构,所以将 `pre-commit` 这个配置文件设置为默认项。关于 `pre-commit` 的详细介绍，请参考：[pre-commit](https://pre-commit.com)

### TODO List
 - [x] 发布一个版本
 - [x] 增加更多的项目结构选择
   - [x] gin 项目
   - [x] DDD 项目
 - [ ] 补充单元测试
 - [ ] 增加更多的参数配置化

### 鸣谢
  > 感谢以下项目的给我的启发
  - [kratos](https://github.com/go-kratos/kratos)
  - [go-micro](https://github.com/asim/go-micro)

### 后记
  > 如果你有什么想法，欢迎在issue中提出。
