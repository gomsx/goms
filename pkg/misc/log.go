package misc

import (
	"context"

	log "github.com/sirupsen/logrus"
)

//
func CarryCtxRequestId(ctx context.Context, requestid int64) context.Context {
	return CarryCtxId(ctx, "request_id", requestid)
}

//
func CarryCtxUserId(ctx context.Context, uid int64) context.Context {
	return CarryCtxId(ctx, "user_id", uid)
}

//
func CarryCtxId(ctx context.Context, key string, val int64) context.Context {
	return ctx
}

//
func GetLogLevel() string {
	level := log.GetLevel()
	return level.String()
}

//
func SetLogLevel(l string) {
	level, err := log.ParseLevel(l)
	if err != nil {
		level = log.Level(log.InfoLevel)
	}
	log.SetLevel(log.Level(level))
}
