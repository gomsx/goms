package model

// const (
// 	HTTP PingType = "http"
// 	GRPC PingType = "grpc"
// )

type Ping struct {
	Type  string
	Count int64
}
