package misc

import (
	"context"
	"math/rand"
	"time"

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

//
func init() {
	InitGenerator()
	SetRequestIdMax(0x0FFF_FFFF_FFFF_FFFF)
}

//
var requestIdMax int64 = 0x0FFF_FFFF_FFFF_FFFF

//
func InitGenerator() {
	rand.Seed(time.Now().UnixNano())
}

//
func SetRequestIdMax(max int64) {
	requestIdMax = max
}

//
func GetRequestId() int64 {
	return rand.Int63n(requestIdMax)
}

////
// MakePongMsg make pong msg.
func MakePongMsg(s string) string {
	if s == "" {
		s = "NONE!"
	}
	return "pong" + " " + s
}

/*
// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key string

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key = "request_id"

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, userKey, id)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userKey).(int64)
	return id, ok
}

//
func GetIdMust(ctx interface{}) int64 {
	switch c := ctx.(type) {
	case *gin.Context:
		rqkey := string(userKey)
		if id := c.GetInt64(rqkey); id != 0 {
			return id
		}
	case context.Context:
		if id, ok := FromContext(c); ok {
			return id
		}
	}
	return 0
}

//
func GetIdMustX(ctx interface{}) int64 {
	switch c := ctx.(type) {
	case *gin.Context:
		rqkey := string(userKey)
		if id := c.GetInt64(rqkey); id != 0 {
			return id
		}
		id := Get()
		c.Set(rqkey, id)
		return id
	case *context.Context:
		if id, ok := FromContext(*c); ok {
			// fmt.Println("get request id", id)
			return id
		}
		id := Get()
		nc := NewContext(*c, id)
		rv := reflect.ValueOf(ctx)
		re := rv.Elem()
		re.Set(reflect.ValueOf(nc))
		// fmt.Println("set request id", id)
		return id
	}
	return 0
}
*/
