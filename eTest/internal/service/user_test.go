package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eTest/internal/dao/mock"
	. "github.com/fuwensun/goms/eTest/internal/model"

	"github.com/golang/mock/gomock"
)

var user = User{
	Uid:  GetUid(),
	Name: "foo",
	Sex:  0,
}

func Test_service_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		CreateUser(gomock.Any(), &user).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		CreateUser(gomock.Any(), &user).
		Return(ErrFailedCreateData)

	type args struct {
		c    context.Context
		user *User
	}

	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "for sucee",
			s:    &svct,
			args: args{
				c:    ctx,
				user: &user,
			},
			wantErr: false,
		},
		{
			name: "for failed",
			s:    &svcf,
			args: args{
				c:    ctx,
				user: &user,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CreateUser(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		UpdateUser(gomock.Any(), &user).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		UpdateUser(gomock.Any(), &user).
		Return(ErrNotFoundData)

	type args struct {
		c    context.Context
		user *User
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "for sucee",
			s:    &svct,
			args: args{
				c:    ctx,
				user: &user,
			},
			wantErr: false,
		},
		{
			name: "for failed",
			s:    &svcf,
			args: args{
				c:    ctx,
				user: &user,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdateUser(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		ReadUser(gomock.Any(), user.Uid).
		Return(user, nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		ReadUser(gomock.Any(), user.Uid).
		Return(user, ErrNotFoundData)

	type args struct {
		c   context.Context
		uid int64
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "for sucee",
			s:    &svct,
			args: args{
				c:   ctx,
				uid: user.Uid,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "for failed",
			s:    &svcf,
			args: args{
				c:   ctx,
				uid: user.Uid,
			},
			want:    user,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ReadUser(tt.args.c, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ReadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ReadUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daot := mock.NewMockDao(ctrl)
	svct := service{dao: daot}
	daot.EXPECT().
		DeleteUser(gomock.Any(), user.Uid).
		Return(nil)

	daof := mock.NewMockDao(ctrl)
	svcf := service{dao: daof}
	daof.EXPECT().
		DeleteUser(gomock.Any(), user.Uid).
		Return(ErrNotFoundData)

	type args struct {
		c   context.Context
		uid int64
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "for sucee",
			s:    &svct,
			args: args{
				c:   ctx,
				uid: user.Uid,
			},
			wantErr: false,
		},
		{
			name: "for failed",
			s:    &svcf,
			args: args{
				c:   ctx,
				uid: user.Uid,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteUser(tt.args.c, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
