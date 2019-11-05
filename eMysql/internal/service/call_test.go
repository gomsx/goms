package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/fuwensun/goms/eMysql/internal/dao/mock"
	"github.com/fuwensun/goms/eMysql/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/prashantv/gostub"
)

//http
func TestUpdateHttpPingCount(t *testing.T) {
	//new mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daom := mock.NewMockDao(ctrl)

	//new data
	svc := Service{}

	//stubs
	stubs := gostub.Stub(&svc.dao, daom)
	// defer stubs.Reset() //svc.dao == nil, panic!
	fmt.Println(stubs)

	var pc model.PingCount = 2
	daom.EXPECT().UpdatePingCount(gomock.Any(), model.HTTP, pc).Return(nil)

	svc.UpdateHttpPingCount(context.Background(), pc)
}

func TestReadHttpPingCount(t *testing.T) {
	//new mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daom := mock.NewMockDao(ctrl)

	//new data
	svc := Service{}

	//stubs
	gostub.Stub(&svc.dao, daom)

	var want model.PingCount = 2
	daom.EXPECT().ReadPingCount(gomock.Any(), model.HTTP).Return(want, nil)

	if got := svc.ReadHttpPingCount(context.Background()); got != want {
		t.Errorf("ReadHttpPingCount() get %v ,want %v", got, want)
	}
}

//grpc
func TestUpdateGrpcPingCount(t *testing.T) {
	//new mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daom := mock.NewMockDao(ctrl)

	//new data
	svc := Service{}

	//stubs
	gostub.Stub(&svc.dao, daom)

	var pc model.PingCount = 2
	daom.EXPECT().UpdatePingCount(gomock.Any(), model.GRPC, pc).Return(nil)

	svc.UpdateGrpcPingCount(context.Background(), pc)
}

func TestReadGrpcPingCount(t *testing.T) {
	//new mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	daom := mock.NewMockDao(ctrl)

	//new data
	svc := Service{}

	//stubs
	gostub.Stub(&svc.dao, daom)

	var want model.PingCount = 2
	daom.EXPECT().ReadPingCount(gomock.Any(), model.GRPC).Return(want, nil)

	if got := svc.ReadGrpcPingCount(context.Background()); got != want {
		t.Errorf("ReadGrpcPingCount() get %v ,want %v", got, want)
	}
}
