package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eTest/internal/dao"
	"github.com/fuwensun/goms/eTest/internal/dao/mock"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
)

var errx = errors.New("test error")
var ctxb = context.Background()

func TestGetConfig(t *testing.T) {
	type args struct {
		cfgpath string
	}
	argx := []args{
		{cfgpath: "./testdata/configs"},
		{cfgpath: "./testdata/configsx"},
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
	acfg := &config{
		Name:    "user",
		Version: "v0.0.0",
	}
	asvc := &service{
		cfg: acfg,
		dao: adao,
	}

	getCfgSucc := Patch(getConfig, func(cfgpath string) (*config, error) {
		return acfg, nil
	})
	getCfgFail := Patch(getConfig, func(cfgpath string) (*config, error) {
		return nil, errx
	})

	type args struct {
		cfgpath string
		dao     dao.Dao
	}
	argsv := args{
		cfgpath: "",
		dao:     adao,
	}

	tests := []struct {
		name    string
		args    args
		patch   *PatchGuard
		want    Svc
		want1   func()
		wantErr bool
	}{
		{name: "New service when getConfig succeeded", args: argsv, patch: getCfgSucc, want: asvc, want1: asvc.Close, wantErr: false},
		{name: "New service when getConfig failed", args: argsv, patch: getCfgFail, want: nil, want1: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.patch.Restore()
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
		{name: "Ping when dao ping succeeded", svc: &svc, args: args{ctx: ctxb}, wantErr: false},
		{name: "Ping when dao ping failed", svc: &svc, args: args{ctx: ctxb}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
