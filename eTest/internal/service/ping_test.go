package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/aivuca/goms/eTest/internal/dao/mock"
	m "github.com/aivuca/goms/eTest/internal/model"

	"github.com/golang/mock/gomock"
)

func TestHandPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ctx := context.Background()
	errt := errors.New("error")
	ping := &m.Ping{Type: "http", Count: 2}
	want := &m.Ping{Type: "http", Count: 3}
	//1
	dao.EXPECT().
		ReadPing(ctx, ping.Type).
		Return(ping, nil)

	dao.EXPECT().
		UpdatePing(ctx, ping).
		Return(nil)
	//2
	dao.EXPECT().
		ReadPing(ctx, ping.Type).
		Return(ping, errt)
	//3
	dao.EXPECT().
		ReadPing(ctx, ping.Type).
		Return(ping, nil)

	dao.EXPECT().
		UpdatePing(ctx, ping).
		Return(errt)
	//
	type args struct {
		c context.Context
		p *m.Ping
	}
	tests := []struct {
		name    string
		args    args
		want    *m.Ping
		wantErr bool
	}{
		{name: "for succ", args: args{ctx, ping}, want: want, wantErr: false},
		{name: "for failed", args: args{ctx, ping}, want: nil, wantErr: true},
		{name: "for failedx2", args: args{ctx, ping}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.HandPing(tt.args.c, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.HandPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("====>", got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.HandPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
