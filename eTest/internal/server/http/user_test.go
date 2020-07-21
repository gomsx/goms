package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	. "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.POST("/user", srv.createUser)

	Convey("TestPing should respond http.StatusCreated", t, func() {

		user := &User{
			Name: "xxx",
			Sex:  1,
		}
		// 问题：CreateUser 方法中 user 参数是指针类型，
		// Name,Sex 或者整个 User 是要输入的参数
		// Uid 或者整个 User 是要输出的参数
		// 这里他们共用了一个参数，没法对输出参数进行 mock，
		// CreateUser 被调用时 Uid 字段没被赋值，默认值是 0.
		// 这里 Uid 字段也必须是 0，不然会报错，没法 mock.
		// 只要分离了输入参数和输出参数，就能给他们 mock 任何合法的值,
		// 这样的代码具有可测试性.
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(nil)

		sexstr := strconv.FormatInt(user.Sex, 10)

		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

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
		So(resp.StatusCode, ShouldEqual, http.StatusCreated)
		So(m["name"], ShouldEqual, user.Name)
		So(m["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("TestPing should respond http.StatusBadRequest", t, func() {

		user := &User{
			Name: "xxx",
			Sex:  99,
		}
		// svcm.EXPECT().
		// 	CreateUser(gomock.Any(), user).
		// 	Return(nil)

		sexstr := strconv.FormatInt(user.Sex, 10)

		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

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
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
		// So(m["name"], ShouldEqual, user.Name)
		// So(m["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("TestPing should respond http.StatusInternalServerError", t, func() {

		user := &User{
			Name: "xxx",
			Sex:  1,
		}

		errx := errors.New("error!")
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(errx)

		sexstr := strconv.FormatInt(user.Sex, 10)

		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

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
	})
}

func TestReadUser(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.GET("/user/:uid", srv.readUser)

	Convey("readUser should respond http.StatusOK", t, func() {

		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  1,
		}

		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, nil)

		//构建请求
		w := httptest.NewRecorder()
		uidstr := strconv.FormatInt(user.Uid, 10)
		r, _ := http.NewRequest("GET", "/user/"+uidstr, nil)

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
		So(m["uid"], ShouldEqual, float64(user.Uid))
		So(m["name"], ShouldEqual, user.Name)
		So(m["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("readUser should respond http.StatusBadRequest", t, func() {

		user := &User{
			Uid:  -123,
			Name: "xxx",
			Sex:  1,
		}

		// mock 的必须调用到,否则报错
		// missing call(s) to *mock.MockSvc.ReadUser(is anything, is equal to -123)
		// svcm.EXPECT().
		// 	ReadUser(gomock.Any(), user.Uid).
		// 	Return(user, nil)

		//构建请求
		w := httptest.NewRecorder()
		uidstr := strconv.FormatInt(user.Uid, 10)
		r, _ := http.NewRequest("GET", "/user/"+uidstr, nil)
		// r, _ := http.NewRequest("GET", "/user/-123", nil)

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
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
		So(m["uid"], ShouldEqual, uidstr)
	})

	Convey("readUser should respond http.StatusNotFound", t, func() {

		user := &User{Uid: 789}

		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, ErrNotFoundData)

		//构建请求
		w := httptest.NewRecorder()
		uidstr := strconv.FormatInt(user.Uid, 10)
		r, _ := http.NewRequest("GET", "/user/"+uidstr, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
	})
}

func TestUpdateUser(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.PUT("/user/:uid", srv.updateUser)

	Convey("updateUser should respond http.StatusNoContent", t, func() {

		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  1,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(nil)

		uidstr := strconv.FormatInt(user.Uid, 10)
		sexstr := strconv.FormatInt(user.Sex, 10)
		v := url.Values{}
		v.Set("uid", uidstr)
		v.Set("name", "xxx")
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+uidstr, reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("updateUser should respond http.StatusBadRequest", t, func() {

		user := &User{
			Uid:  -123,
			Name: "xxx",
			Sex:  1,
		}
		// svcm.EXPECT().
		// 	UpdateUser(gomock.Any(), user).
		// 	Return(nil)

		uidstr := strconv.FormatInt(user.Uid, 10)
		sexstr := strconv.FormatInt(user.Sex, 10)
		v := url.Values{}
		v.Set("uid", uidstr)
		v.Set("name", "xxx")
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+uidstr, reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("updateUser should respond http.StatusNotFound", t, func() {

		user := &User{
			Uid:  789,
			Name: "xxx",
			Sex:  1,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(ErrNotFoundData)

		uidstr := strconv.FormatInt(user.Uid, 10)
		sexstr := strconv.FormatInt(user.Sex, 10)
		v := url.Values{}
		v.Set("uid", uidstr)
		v.Set("name", "xxx")
		v.Set("sex", sexstr)
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+uidstr, reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
	})
}

func TestDeleteUser(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.DELETE("/user/:uid", srv.deleteUser)

	Convey("deleteUser should respond http.StatusNoContent", t, func() {

		var uid int64 = 123
		svcm.EXPECT().
			DeleteUser(gomock.Any(), uid).
			Return(nil)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/123", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("deleteUser should respond http.StatusBadRequest", t, func() {

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/-123", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("deleteUser should respond http.StatusNotFound", t, func() {

		var uid int64 = 789
		svcm.EXPECT().
			DeleteUser(gomock.Any(), uid).
			Return(ErrNotFoundData)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/789", nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
	})
}
