package misc

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCarryCtxId(t *testing.T) {

	ctxb := context.Background()
	ctxfoo := CarryCtxId(ctxb, "foo", 123)
	ctxbar := CarryCtxId(ctxfoo, "bar", 456)

	fmt.Printf("%p\n", ctxb)
	fmt.Printf("%p\n", ctxfoo)
	fmt.Printf("%p\n", ctxbar)

	fmt.Printf("%#v\n", ctxb)
	fmt.Printf("%#v\n", ctxfoo)
	fmt.Printf("%#v\n", ctxbar)

	fmt.Printf("%p\n", log.Ctx(ctxb))
	fmt.Printf("%p\n", log.Ctx(ctxfoo))
	fmt.Printf("%p\n", log.Ctx(ctxbar))

	fmt.Printf("%#v\n", log.Ctx(ctxb))
	fmt.Printf("%#v\n", log.Ctx(ctxfoo))
	fmt.Printf("%#v\n", log.Ctx(ctxbar))

	fmt.Printf("%#v\n", log.Ctx(ctxb).Info())
	fmt.Printf("%#v\n", log.Ctx(ctxfoo).Info())
	fmt.Printf("%#v\n", log.Ctx(ctxbar).Info())

	type args struct {
		ctx context.Context
		key string
		val int64
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{name: "for succ", args: args{ctx: ctxb, key: "foo", val: 123}, want: ctxfoo},
		{name: "for succx2", args: args{ctx: ctxfoo, key: "bar", val: 456}, want: ctxbar},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CarryCtxId(tt.args.ctx, tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarryCtxId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenUid(t *testing.T) {
	uidx := GenUid()
	uidy := GenUid()
	uidz := GenUid()
	tests := []struct {
		name    string
		compute bool
		want    bool
	}{
		{"succx1", uidx-0 > 0, true},
		{"succx2", uidy-uidx > 0, true},
		{"succx3", uidz-uidy > 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.compute; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRequestId(t *testing.T) {
	ridx := GenRequestId()
	ridy := GenRequestId()
	ridz := GenRequestId()
	tests := []struct {
		name    string
		compute bool
		want    bool
	}{
		{"succx1", ridx-0 > 0, true},
		{"succx2", ridy-ridx > 0, true},
		{"succx3", ridz-ridy > 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.compute; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
