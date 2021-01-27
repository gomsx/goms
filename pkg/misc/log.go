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
	// l := log.With().Int64(key, val).Logger() // 丢失 ctx 中的 (key，val)
	// l := log.Ctx(ctx).With().Int64(key, val).Logger() // 保存 ctx 中的 (key，val)
	// return l.WithContext(ctx)
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
		level = log.InfoLevel
	}
	log.SetLevel(log.Level(level))
}
