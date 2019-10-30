module github.com/fuwensun/goms

go 1.13

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271 // indirect
	golang.org/x/sys v0.0.0-20191029155521-f43be2a4598c // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
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
