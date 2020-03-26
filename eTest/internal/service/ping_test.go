package service

//http
// func TestUpdateHttpPingCount(t *testing.T) {
// 	//new mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	daom := mock.NewMockDao(ctrl)

// 	//new data
// 	svc := service{}
// 	// svc, _, _ := New("", daom)
// 	// svc.dao = daom
// 	//stubs
// 	stubs := gostub.Stub(&svc.dao, daom)
// 	// defer stubs.Reset() //svc.dao == nil, panic!
// 	fmt.Println(stubs)

// 	var pc model.PingCount = 2
// 	daom.EXPECT().UpdatePingCount(gomock.Any(), model.HTTP, pc).Return(nil)

// 	svc.UpdateHttpPingCount(context.Background(), pc)
// }

// func TestReadHttpPingCount(t *testing.T) {
// 	//new mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	daom := mock.NewMockDao(ctrl)

// 	//new data
// 	svc := Service{}

// 	//stubs
// 	gostub.Stub(&svc.dao, daom)

// 	var want model.PingCount = 2
// 	daom.EXPECT().ReadPingCount(gomock.Any(), model.HTTP).Return(want, nil)

// 	if got := svc.ReadHttpPingCount(context.Background()); got != want {
// 		t.Errorf("ReadHttpPingCount() get %v ,want %v", got, want)
// 	}
// }

// //grpc
// func TestUpdateGrpcPingCount(t *testing.T) {
// 	//new mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	daom := mock.NewMockDao(ctrl)

// 	//new data
// 	svc := Service{}

// 	//stubs
// 	gostub.Stub(&svc.dao, daom)

// 	var pc model.PingCount = 2
// 	daom.EXPECT().UpdatePingCount(gomock.Any(), model.GRPC, pc).Return(nil)

// 	svc.UpdateGrpcPingCount(context.Background(), pc)
// }

// func TestReadGrpcPingCount(t *testing.T) {
// 	//new mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	daom := mock.NewMockDao(ctrl)

// 	//new data
// 	svc := Service{}

// 	//stubs
// 	gostub.Stub(&svc.dao, daom)

// 	var want model.PingCount = 2
// 	daom.EXPECT().ReadPingCount(gomock.Any(), model.GRPC).Return(want, nil)

// 	if got := svc.ReadGrpcPingCount(context.Background()); got != want {
// 		t.Errorf("ReadGrpcPingCount() get %v ,want %v", got, want)
// 	}
// }
