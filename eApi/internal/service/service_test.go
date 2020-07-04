package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/aivuca/goms/eApi/internal/dao"
	"github.com/aivuca/goms/eApi/internal/dao/mock"
)

var ctx = context.Background()

func Test_getConfig(t *testing.T) {
	type args struct {
		cfgpath string
	}
	tests := []struct {
		name    string
		args    args
		want    *config
		wantErr bool
	}{
		{
			name: "for succ",
			args: args{
				cfgpath: "./testdata",
			},
			want: &config{
				Name:    "user",
				Version: "v0.0.0",
			},
			wantErr: false,
		},
		{
			name: "for failed",
			args: args{
				cfgpath: "./testdata/xxx",
			},
			want:    nil,
			wantErr: true,
		},
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
	// daoa := &dao.Daot{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daoa := mock.NewMockDao(ctrl)

	s := &service{
		cfg: &config{
			Name:    "user",
			Version: "v0.0.0",
		},
		dao: daoa,
		// dao: &dao.Daot{},
	}
	type args struct {
		cfgpath string
		d       dao.Dao
	}
	tests := []struct {
		name    string
		args    args
		want    Svc
		want1   func()
		wantErr bool
	}{
		{
			name: "for succ",
			args: args{
				cfgpath: "./testdata",
				d:       daoa,
				// d: &dao.Daot{},
			},
			want:    s,
			want1:   s.Close,
			wantErr: false,
		},
		{
			name: "for failed",
			args: args{
				cfgpath: "./testdata/xxx",
				d:       daoa,
				// d: &dao.Daot{},
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := New(tt.args.cfgpath, tt.args.d)
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

func Test_service_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daot := mock.NewMockDao(ctrl)
	daot.EXPECT().Ping(gomock.Any()).Return(nil)
	svct := service{dao: daot}

	daof := mock.NewMockDao(ctrl)
	daof.EXPECT().Ping(gomock.Any()).Return(errors.New("x"))
	svcf := service{dao: daof}

	type args struct {
		c context.Context
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "for succ",
			s:    &svct,
			args: args{
				c: ctx,
			},
			wantErr: false,
		},
		{
			name: "for failed",
			s:    &svcf,
			args: args{
				c: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Ping(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("service.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Close(t *testing.T) {
	svct := service{}
	svcf := service{}
	tests := []struct {
		name string
		s    *service
	}{
		{
			name: "for succ",
			s:    &svct,
		},
		{
			name: "for failed",
			s:    &svcf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Close()
		})
	}
}

