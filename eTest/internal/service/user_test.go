package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/aivuca/goms/eTest/internal/dao/mock"
	m "github.com/aivuca/goms/eTest/internal/model"
	e "github.com/aivuca/goms/eTest/internal/pkg/err"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	user := m.GetUser()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		CreateUser(gomock.Any(), user).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		CreateUser(gomock.Any(), user).
		Return(e.ErrFailedCreateData)

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
		{name: "for sucee", svc: &svct, args: args{ctx: ctx, user: user}, wantErr: false},
		{name: "for failed", svc: &svcf, args: args{ctx: ctx, user: user}, wantErr: true},
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

	ctx := context.Background()
	user := m.GetUser()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		ReadUser(gomock.Any(), user.Uid).
		Return(user, nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		ReadUser(gomock.Any(), user.Uid).
		Return(nil, e.ErrNotFoundData)

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
		{name: "for sucee", svc: &svct, args: args{ctx: ctx, uid: user.Uid}, want: user, wantErr: false},
		{name: "for failed", svc: &svcf, args: args{ctx: ctx, uid: user.Uid}, want: nil, wantErr: true},
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

	ctx := context.Background()
	user := m.GetUser()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		UpdateUser(gomock.Any(), user).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		UpdateUser(gomock.Any(), user).
		Return(e.ErrNotFoundData)

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
		{name: "for sucee", svc: &svct, args: args{ctx: ctx, user: user}, wantErr: false},
		{name: "for failed", svc: &svcf, args: args{ctx: ctx, user: user}, wantErr: true},
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

	ctx := context.Background()
	user := m.GetUser()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		DeleteUser(gomock.Any(), user.Uid).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		DeleteUser(gomock.Any(), user.Uid).
		Return(e.ErrNotFoundData)

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
		{name: "for sucee", svc: &svct, args: args{ctx: ctx, uid: user.Uid}, wantErr: false},
		{name: "for failed", svc: &svcf, args: args{ctx: ctx, uid: user.Uid}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.DeleteUser(tt.args.ctx, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
