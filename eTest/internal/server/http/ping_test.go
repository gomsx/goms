package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Ping(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.GET("/ping", srv.ping)

	Convey("TestPing should respond http.StatusOK", t, func() {

		var pc PingCount = 2
		svcm.EXPECT().
			HandPingHttp(gomock.Any()).
			Return(pc, nil)

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
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &m)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", m)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(m["message"], ShouldEqual, "pong NONE!")
		So(m["count"], ShouldEqual, float64(pc))
	})

	Convey("TestPing should respond http.StatusOK", t, func() {

		var pc PingCount = 2
		svcm.EXPECT().
			HandPingHttp(gomock.Any()).
			Return(pc, nil)

		//构建req
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/ping?message=xxx", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &m)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", m)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(m["message"], ShouldEqual, "pong xxx")
		So(m["count"], ShouldEqual, float64(pc))
	})

	Convey("TestPing should respond http.StatusInternalServerError", t, func() {

		var pc PingCount = 2
		svcm.EXPECT().
			HandPingHttp(gomock.Any()).
			Return(pc, ErrNotFoundData)

		//构建req
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/ping?message=xxx", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &m)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", m)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
		// So(m["error"], ShouldEqual, "internal error!")
	})
}
