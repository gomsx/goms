package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eApi/internal/dao/mock"
	m "github.com/fuwensun/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ctx := context.Background()
	user := m.GetUser()
	errt := errors.New("error")
	//
	dao.EXPECT().
		CreateUser(ctx, user).
		Return(nil)

	dao.EXPECT().
		CreateUser(ctx, user).
		Return(errt)

	type args struct {
		ctx  context.Context
		user *m.User
	}

	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "for succ", svc: &svc, args: args{ctx: ctx, user: user}, wantErr: false},
		{name: "for failed", svc: &svc, args: args{ctx: ctx, user: user}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ctx := context.Background()
	user := m.GetUser()
	errt := errors.New("error")
	//
	dao.EXPECT().
		ReadUser(ctx, user.Uid).
		Return(user, nil)

	dao.EXPECT().
		ReadUser(ctx, user.Uid).
		Return(nil, errt)

	type args struct {
		ctx context.Context
		uid int64
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		want    *m.User
		wantErr bool
	}{
		{name: "for succ", svc: &svc, args: args{ctx: ctx, uid: user.Uid}, want: user, wantErr: false},
		{name: "for failed", svc: &svc, args: args{ctx: ctx, uid: user.Uid}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.ReadUser(tt.args.ctx, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ReadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("service.ReadUser() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ctx := context.Background()
	user := m.GetUser()
	errt := errors.New("error")

	dao.EXPECT().
		UpdateUser(ctx, user).
		Return(nil)

	dao.EXPECT().
		UpdateUser(ctx, user).
		Return(errt)

	type args struct {
		ctx  context.Context
		user *m.User
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "for succ", svc: &svc, args: args{ctx: ctx, user: user}, wantErr: false},
		{name: "for failed", svc: &svc, args: args{ctx: ctx, user: user}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ctx := context.Background()
	user := m.GetUser()
	errt := errors.New("error")

	dao.EXPECT().
		DeleteUser(ctx, user.Uid).
		Return(nil)

	dao.EXPECT().
		DeleteUser(ctx, user.Uid).
		Return(errt)

	type args struct {
		ctx context.Context
		uid int64
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "for succ", svc: &svc, args: args{ctx: ctx, uid: user.Uid}, wantErr: false},
		{name: "for failed", svc: &svc, args: args{ctx: ctx, uid: user.Uid}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.DeleteUser(tt.args.ctx, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
