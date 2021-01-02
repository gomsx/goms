package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/aivuca/goms/eApi/internal/model"
	"github.com/aivuca/goms/eApi/internal/service/mock"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	ping := &m.Ping{Type: "http"}
	want := &m.Ping{Type: "http", Count: 5}
	//
	router := gin.New()
	router.Use(setRequestId())
	router.GET("/ping", srv.ping)

	Convey("TestPing should respond http.StatusOK", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, nil)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/ping", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(rm["message"], ShouldEqual, ms.MakePongMsg(""))
		So(rm["count"], ShouldEqual, want.Count)
	})

	Convey("TestPing should respond http.StatusOK", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, nil)

		//构建req
		msg := "xxx"
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/ping?message="+msg, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(rm["message"], ShouldEqual, ms.MakePongMsg(msg))
		So(rm["count"], ShouldEqual, want.Count)
	})

	Convey("TestPing should respond http.StatusInternalServerError", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, errx)

		//构建req
		msg := "xxx"
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/ping?message="+msg, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
