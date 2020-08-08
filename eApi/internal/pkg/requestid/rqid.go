package requestid

import (
	"context"
	"math/rand"
	"reflect"
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
func GetIdMust(c context.Context) int64 {
	if id, ok := FromContext(c); ok {
		return id
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
