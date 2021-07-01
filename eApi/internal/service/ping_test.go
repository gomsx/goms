package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/gomsx/goms/eApi/internal/dao"
	"github.com/gomsx/goms/eApi/internal/dao/mock"
	m "github.com/gomsx/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
)

func TestHandPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daom := mock.NewMockDao(ctrl)

	ping := &m.Ping{Type: "http", Count: 2}
	want := &m.Ping{Type: "http", Count: 3}

	// succ
	daom.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, nil)
	daom.EXPECT().
		UpdatePing(ctxb, ping).
		Return(nil)

	// failed
	daom.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, errx)

	// failedx2
	daom.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, nil)
	daom.EXPECT().
		UpdatePing(ctxb, ping).
		Return(errx)

	type fields struct {
		cfg *config
		dao dao.Dao
	}
	type args struct {
		c context.Context
		p *m.Ping
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *m.Ping
		wantErr bool
	}{
		{name: "Hand ping when read succeeded and update succeeded", fields: fields{cfg: nil, dao: daom}, args: args{c: ctxb, p: ping}, want: want, wantErr: false},
		{name: "Hand ping when read failed", fields: fields{cfg: nil, dao: daom}, args: args{c: ctxb, p: ping}, want: nil, wantErr: true},
		{name: "Hand ping when read succeeded and update failed", fields: fields{cfg: nil, dao: daom}, args: args{c: ctxb, p: ping}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg: tt.fields.cfg,
				dao: tt.fields.dao,
			}
			got, err := s.HandPing(tt.args.c, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.HandPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.HandPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
