package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	. "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Ping(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)
	router := gin.New()
	srv := Server{}
	router.GET("/ping", srv.ping)
	var pc PingCount = 2
	Convey("TestPing should respond http.StatusOK", t, func() {
		monkey.Patch(handping, func(c *gin.Context, svc service.Svc) (PingCount, error) {
			return pc, nil
		})
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
		monkey.Patch(handping, func(c *gin.Context, svc service.Svc) (PingCount, error) {
			return pc, nil
		})

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
		monkey.Patch(handping, func(c *gin.Context, svc service.Svc) (PingCount, error) {
			return pc, ErrNotFoundData
		})

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
