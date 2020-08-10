package rqid

import (
	"context"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	InitGenerator()
	SetMax(0x0FFF_FFFF_FFFF_FFFF)
}

//
var rqidmax int64 = 0x0FFF_FFFF_FFFF_FFFF

func InitGenerator() {
	rand.Seed(time.Now().UnixNano())
}
func SetMax(max int64) {
	rqidmax = max
}
func Get() int64 {
	return rand.Int63n(rqidmax)
}

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
	switch v := ctx.(type) {
	case *gin.Context:
		return v.GetInt64(string(userKey))
	case context.Context:
		id, ok := FromContext(v)
		if !ok {
			return 0
		}
		return id
	}
	return 0
}
