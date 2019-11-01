package model

type PingCount int

type PingType string

const (
	HTTP PingType = "http"
	GRPC PingType = "grpc"
)

