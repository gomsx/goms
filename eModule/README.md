# eModule

go module 管理依赖.它包含:

- go mod 命令
- go.mod 文件
- go.sum 文件

源码中特意引入特定的依赖,来全面覆盖各种情况,再进行分析,全面深入的理解 go module.

## go mod 命令

处理依赖

```
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```

> go init 新建模块等于新建依赖也算处理依赖

## go.mod 文件

描述依赖

```
//go.mod

module github.com/fuwensun/goms

go 1.13

require (
	github.com/gomsx/hello v0.0.2 // indirect
	github.com/gomsx/helloworld v1.0.1
	github.com/gomsx/world/v2 v2.0.2 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190926180325-855e68c8590b
	google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0
)

exclude (
	github.com/golang/mock v1.1.0
	github.com/golang/mock v1.1.1
	github.com/golang/mock v1.2.0
)
```

require:依赖的模块

replace:替换的模块

exclude:排除的模块

标准模块版本号:语义版本号,源码库tag.如:`v1.24.0`

伪模块版本号:源码库最近tag-commit日期时间-commit哈希值前缀.如:`v0.0.0-20190926180325-855e68c8590b`

`+incompatible` 表示兼容.

`//indirect` 表示间接引用.否则为直接引用.

## go.sum 文件

校验依赖

```
//go.sum
...
github.com/golang/mock v1.3.1 h1:qGJ6qTW+x6xX/my+8YUVl4WNpX9B7+/l2tRsHGZ7f2s=
github.com/golang/mock v1.3.1/go.mod h1:sBzyDLLjw3U8JLTeZvSv8jJB+tU5PVekmnlKIyFUx0Y=
...

```

引用 go/src/cmd/go/alldocs.go:
```
// The form of each line in go.sum is three fields:
//
// 	<module> <version>[/go.mod] <hash>
//
// Each known module version results in two lines in the go.sum file.
// The first line gives the hash of the module version's file tree.
// The second line appends "/go.mod" to the version and gives the hash
// of only the module version's (possibly synthesized) go.mod file.
// The go.mod-only hash allows downloading and authenticating a
// module version's go.mod file, which is needed to compute the
// dependency graph, without also downloading all the module's source code
```
总结:

- go.sun 文件行的格式: `<module> <version>[/go.mod] <hash>`.
- 每个模块占两行.第一行是模块文件树的 hash 值.第二行带 "/go.mod",仅是 go.mod 文件的 hash 值.
- go.mod 文件是用来计算模块的依赖图.
- 由于 go.mod 文件的哈希值的存在,只需下载并校验需要的版本模块的 go.mod 文件,而无需下载整个源码,来计算模块依赖图.

## 注意事项
1,有些包在国外，要设置代理 GOPROXY="https://goproxy.cn,direct"
2,replace 可以锁定版本

