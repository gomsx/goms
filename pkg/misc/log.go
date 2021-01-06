package misc

import (
	"context"

	"github.com/rs/zerolog/log"
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
	l := log.Ctx(ctx).With().Int64(key, val).Logger() // 保存 ctx 中的 (key，val)
	return l.WithContext(ctx)
}