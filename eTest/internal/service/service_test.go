package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eTest/internal/dao"
	"github.com/fuwensun/goms/eTest/internal/dao/mock"

	"github.com/golang/mock/gomock"
)

var errx = errors.New("test error")
var ctxb = context.Background()

func TestGetConfig(t *testing.T) {
	type args struct {
		cfgpath string
	}
	argx := []args{
		{cfgpath: "./testdata"},
		{cfgpath: "./testxxxx"},
	}
	wantx := []*config{
		{Name: "user", Version: "v0.0.0"},
		nil,
	}
	tests := []struct {
		name    string
		args    args
		want    *config
		wantErr bool
	}{
		{name: "for succ", args: argx[0], want: wantx[0], wantErr: false},
		{name: "for failed", args: argx[1], want: wantx[1], wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfig(tt.args.cfgpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	adao := &dao.Daot{}
	svc := &service{
		cfg: &config{
			Name:    "user",
			Version: "v0.0.0",
		},
		dao: adao,
	}
	type args struct {
		cfgpath string
		dao     dao.Dao
	}
	tests := []struct {
		name    string
		args    args
		want    Svc
		want1   func()
		wantErr bool
	}{
		{name: "for succ", args: args{cfgpath: "./testdata", dao: adao}, want: svc, want1: svc.Close, wantErr: false},
		{name: "for failed", args: args{cfgpath: "./testxxxx", dao: adao}, want: nil, want1: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := New(tt.args.cfgpath, tt.args.dao)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			// if !reflect.DeepEqual(got1, tt.want1) {
			if !(fmt.Sprintf("%p", got1) == fmt.Sprintf("%p", tt.want1)) {
				t.Errorf("New() got1 = %p, want %p", got1, tt.want1)
			}
		})
	}
}

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	svc := service{dao: dao}
	//
	dao.EXPECT().
		Ping(ctxb).
		Return(nil)

	dao.EXPECT().
		Ping(ctxb).
		Return(errx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "for succ", svc: &svc, args: args{ctx: ctxb}, wantErr: false},
		{name: "for failed", svc: &svc, args: args{ctx: ctxb}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
