package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eApi/internal/dao"
	"github.com/fuwensun/goms/eApi/internal/dao/mock"

	"github.com/golang/mock/gomock"
)

var errx = errors.New("error xxx")
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
		{name: "Get config with correct config path", args: argx[0], want: wantx[0], wantErr: false},
		{name: "Get config with incorrect config path", args: argx[1], want: wantx[1], wantErr: true},
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
	asvc := &service{
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
		{name: "New service with correct config path", args: args{cfgpath: "./testdata", dao: adao}, want: asvc, want1: asvc.Close, wantErr: false},
		{name: "New service with incorrect config path", args: args{cfgpath: "./testxxxx", dao: adao}, want: nil, want1: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := New(tt.args.cfgpath, tt.args.dao)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	adao := mock.NewMockDao(ctrl)
	asvc := service{dao: adao}
	//
	adao.EXPECT().
		Ping(ctxb).Return(nil)
	adao.EXPECT().
		Ping(ctxb).Return(errx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "Ping when dao ping succeeded", svc: &asvc, args: args{ctx: ctxb}, wantErr: false},
		{name: "Ping when dao ping failed", svc: &asvc, args: args{ctx: ctxb}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
