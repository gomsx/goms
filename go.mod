module github.com/aivuca/goms

go 1.14

require (
	bou.ke/monkey v1.0.2
	cloud.google.com/go v0.58.0 // indirect
	github.com/fullstorydev/grpcurl v1.6.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gomsx/hello v0.0.2 // indirect
	github.com/gomsx/helloworld v1.0.1
	github.com/gomsx/world/v2 v2.0.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/nbio/st v0.0.0-20140626010706-e9e8d9816f32
	github.com/prashantv/gostub v1.0.0
	github.com/qiniu/x v1.11.5
	github.com/rs/zerolog v1.18.0
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/unknwon/com v1.0.1
	golang.org/x/exp/errors v0.0.0-20200224162631-6cc2880d07d6
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200618031402-d15173dcc7e4 // indirect
	google.golang.org/genproto v0.0.0-20200618031413-b414f8b61790
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v2 v2.3.0
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
