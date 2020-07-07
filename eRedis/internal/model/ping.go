package model

type PingCount int64

type PingType string

const (
	HTTP PingType = "http"
	GRPC PingType = "grpc"
)
