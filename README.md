# gin-restful-best-practice
该项目以最简单易懂的框架向大家展示了我在开发gin web服务的过程中的一些最佳实践。

数据实体：用户（user）、深度学习模型（model）
场景：用户可以上传深度学习模型、下载深度学习模型。

涉及到的技术有：
- [x] GIN
- [x] GORM & PostgreSQL
- [x] JWT
- [x] gRPC （与python机器学习服务通过rpc通信）
- [ ] Travis-ci
- [ ] and more…

## server.go
## conf
### config.go
存放所有服务器配置，环境分为`dev, test, prod`，根据环境变量`ENV`的不同会调用不同的函数来初始化配置，部分配置也可以从环境变量中读取。

由loadConfig函数载入配置，配置由`conf.Conf()`读取，使用了`sync.Once`保证配置只被初始化一次。

## controllers
存放各接口的逻辑以及单元测试（注意，该路径下的单元测试会绕过中间件，这其实是正确的，中间件的测试就应该放在所属文件夹下测试）。
### common.go
一些controller中通用的函数，例如对错误的返回`ErrorResponse`、方便单元测试发送JSON网络请求的`testRequest`（参数通过`gin.H`类型传入，该函数会自动把参数转换为Body或Query）。
### models.go

## middleware
## models
## photos
## routes
## services
## utils

## Build Setup

```shell script
# clone the project
git clone https://github.com/Bingmang/gin-restful-best-practice.git

# enter the project directory
cd gin-restful-best-practice

# install dependency
go run server.go
```

This will automatically open http://localhost:8000

## Build

```shell script
go build server.go
```

## gRPC

如果修改了proto文件需要重新编译pb文件，输入以下命令（要先安装protoc）：
https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.3

```shell script
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=plugins=grpc:. -I./protos ./protos/ml_service.proto
```

## PostgreSQL initialization

```shell script
docker run -d --name isp_test -p 5432:5432 postgres
docker exec -it isp_test /bin/bash
su postgres
create user dev_user with password 'dev_password';
create database dev_database owner dev_user;
grant all on database dev_database to dev_user;
```

## Deploy

```shell script
go build server.go
ENV=dev ./server
ENV=test ./server
ENV=prod ./server
...
```
