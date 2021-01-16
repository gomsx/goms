# swagger

## go-swagger

- Swagger2.0 又称 OpenAPI2.0
- Swagger 是 RESTful API 的简单但功能强大的表示形式
- go-swagger 包是 Swagger2.0 的 go 实现,主要功能是序列化和反序列化 Swagger 规范
- go-swagger 是主要用于生成或分析源代码的工具

### 特性

go-swagger 包含完整的，功能齐全的高性能 API 组件，兼容 Swagger API：服务器，客户端和数据模型。

- 根据规范生成服务器
- 根据规范生成客户端
- 支持 jsonschema 和 swagger 提供的大多数功能，包括多态
- 从带注释的 go 代码生成规范
- 配合规范使用的其他工具
- 强大的定制功能，带有供应商扩展和可定制的模板

生成惯用的，快速运行的代码,并且兼容 golint, go vet 等.

### 安装

```
# from docker
docker pull quay.io/goswagger/swagger
alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
swagger version

# from source
go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger version
```

>https://github.com/go-swagger/go-swagger  
https://goswagger.io  
https://goswagger.io/install.html  
https://goswagger.io/generate/spec.html  
https://goswagger.io/use/spec.html  
https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9  

## swaggo

>https://github.com/swaggo/swag  
https://github.com/swaggo/gin-swagger 

## swagger ui

```
# 安装
docker pull swaggerapi/swagger-ui 

# 使用
function swaggerui(){ docker run -p 80:8080 -e SWAGGER_JSON=/$HOME/$1 -v /$PWD:/$HOME swaggerapi/swagger-ui;}

# 访问
curl http://127.0.0.1:80/
http://localhost:80/

```
>docker run -p 80:8080 -e SWAGGER_JSON=/foo/api.swagger.json -v /$PWD:/foo swaggerapi/swagger-ui &  
docker run -p 80:8080 -e SWAGGER_JSON=/$HOME/api.swagger.json -v /$PWD:/$HOME swaggerapi/swagger-ui &  
docker run -p 80:8080 -e SWAGGER_JSON=/$HOME/api.swagger.json -v /$PWD swaggerapi/swagger-ui &  