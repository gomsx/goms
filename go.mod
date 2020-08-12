module github.com/aivuca/goms

go 1.14

require (
	bou.ke/monkey v1.0.2
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/alicebob/miniredis/v2 v2.13.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gomsx/hello v0.0.2 // indirect
	github.com/gomsx/helloworld v1.0.1
	github.com/gomsx/world/v2 v2.0.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/kr/pretty v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/prashantv/gostub v1.0.0
	github.com/rs/zerolog v1.18.0
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/unknwon/com v1.0.1
	golang.org/x/exp/errors v0.0.0-20200224162631-6cc2880d07d6
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sys v0.0.0-20200808120158-1030fc2bf1d9 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200808173500-a06252235341
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190926180325-855e68c8590b

exclude (
	github.com/golang/mock v1.1.0
	github.com/golang/mock v1.1.1
	github.com/golang/mock v1.2.0
)
