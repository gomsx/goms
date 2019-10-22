
# go mod command
```
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```
# go.mod
```
module github.com/fuwensun/goms

go 1.13

require (
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	golang.org/x/sys v0.0.0-20190926180325-855e68c8590b // indirect
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


标准模块版本号:语义版本号,源码库标签的修订号.`v1.24.0`

伪模块版本号:`源码库最近标签修订号-提交日期时间-提交哈希值前缀`.`v0.0.0-20190926180325-855e68c8590b`

`+incompatible` 表示兼容.

`//indirect` 表示间接引用.否则为直接引用.




# go.sum
```

...
github.com/golang/mock v1.1.1 h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=
github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
github.com/golang/mock v1.3.0/go.mod h1:c8YoAQJ7+qIz9IQm9G72MJ4uDcrPeLjkrQ4yYIHdhyw=
github.com/golang/mock v1.3.1 h1:qGJ6qTW+x6xX/my+8YUVl4WNpX9B7+/l2tRsHGZ7f2s=
github.com/golang/mock v1.3.1/go.mod h1:sBzyDLLjw3U8JLTeZvSv8jJB+tU5PVekmnlKIyFUx0Y=
github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.1 h1:YF8+flBXS5eO826T4nzqPrxfhQThhXl0YzfuUPu4SBg=
github.com/golang/protobuf v1.3.1/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.2 h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=
github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
...

```