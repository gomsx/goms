package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/aivuca/goms/eApi/internal/dao/mock"
	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	user := m.GetUser()
	//
	dao.EXPECT().
		CreateUser(ctxb, user).
		Return(nil)

	dao.EXPECT().
		CreateUser(ctxb, user).
		Return(errx)

	type args struct {
		ctxb context.Context
		user *m.User
	}

	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "Create user when dao do succeeded", svc: &svc, args: args{ctxb: ctxb, user: user}, wantErr: false},
		{name: "Create user when dao do failed", svc: &svc, args: args{ctxb: ctxb, user: user}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.CreateUser(tt.args.ctxb, tt.args.user); (err != nil) != tt.wantErr {
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
	user := m.GetUser()
	//
	dao.EXPECT().
		ReadUser(ctxb, user.Uid).
		Return(user, nil)

	dao.EXPECT().
		ReadUser(ctxb, user.Uid).
		Return(nil, errx)

	type args struct {
		ctxb context.Context
		uid  int64
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		want    *m.User
		wantErr bool
	}{
		{name: "Read user when dao do succeeded", svc: &svc, args: args{ctxb: ctxb, uid: user.Uid}, want: user, wantErr: false},
		{name: "Read user when dao do failed", svc: &svc, args: args{ctxb: ctxb, uid: user.Uid}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.ReadUser(tt.args.ctxb, tt.args.uid)
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
	user := m.GetUser()
	//
	dao.EXPECT().
		UpdateUser(ctxb, user).
		Return(nil)

	dao.EXPECT().
		UpdateUser(ctxb, user).
		Return(errx)

	type args struct {
		ctxb context.Context
		user *m.User
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "Update user when dao do succeeded", svc: &svc, args: args{ctxb: ctxb, user: user}, wantErr: false},
		{name: "Update user when dao do failed", svc: &svc, args: args{ctxb: ctxb, user: user}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.UpdateUser(tt.args.ctxb, tt.args.user); (err != nil) != tt.wantErr {
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
	user := m.GetUser()
	//
	dao.EXPECT().
		DeleteUser(ctxb, user.Uid).
		Return(nil)

	dao.EXPECT().
		DeleteUser(ctxb, user.Uid).
		Return(errx)

	type args struct {
		ctxb context.Context
		uid  int64
	}
	tests := []struct {
		name    string
		svc     *service
		args    args
		wantErr bool
	}{
		{name: "Delete user when dao do succeeded", svc: &svc, args: args{ctxb: ctxb, uid: user.Uid}, wantErr: false},
		{name: "Delete user when dao do failed", svc: &svc, args: args{ctxb: ctxb, uid: user.Uid}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.DeleteUser(tt.args.ctxb, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
