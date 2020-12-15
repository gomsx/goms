module github.com/fuwensun/goms

go 1.15

require (
	bou.ke/monkey v1.0.2
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/alicebob/miniredis/v2 v2.13.1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gomsx/hello v0.0.2 // indirect
	github.com/gomsx/helloworld v1.0.1
	github.com/gomsx/world/v2 v2.0.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/unknwon/com v1.0.1
	golang.org/x/exp/errors v0.0.0-20200224162631-6cc2880d07d6
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/genproto v0.0.0-20201019141844-1ed22bb0c154
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190926180325-855e68c8590b

exclude (
	github.com/golang/mock v1.1.0
	github.com/golang/mock v1.1.1
	github.com/golang/mock v1.2.0
)
