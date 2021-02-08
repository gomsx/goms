package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	m "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"
	ms "github.com/fuwensun/goms/pkg/misc"

	. "bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var errx = errors.New("test error")
var ctxa = gomock.Any()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	router := gin.New()
	router.POST("/user", srv.createUser)

	Convey("createUser should respond http.StatusCreated", t, func() {
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxa, user).
			Return(nil)
		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
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
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusCreated)
		So(rm["name"], ShouldEqual, user.Name)
		So(rm["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("createUser should respond http.StatusBadRequest", t, func() {
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = ms.GetSexBad()

		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
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
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("createUser should respond http.StatusInternalServerError", t, func() {
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxa, user).
			Return(errx)

		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
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

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	router := gin.New()
	router.GET("/user/:uid", srv.readUser)

	Convey("readUser should respond http.StatusOK", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
			Return(user, nil)
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+ms.StrInt(user.Uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(rm["uid"], ShouldEqual, float64(user.Uid))
		So(rm["name"], ShouldEqual, user.Name)
		So(rm["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("readUser should respond http.StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+ms.StrInt(user.Uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
		So(rm["Uid"], ShouldEqual, float64(user.Uid))
	})

	Convey("readUser should respond http.StatusInternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
			Return(user, errx)
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+ms.StrInt(user.Uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	router := gin.New()
	router.PUT("/user/:uid", srv.updateUser)

	Convey("updateUser should respond http.StatusNoContent", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctxa, user).
			Return(nil)
		//构建请UidUid
		v := url.Values{}
		v.Set("uid", ms.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+ms.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("updateUser should respond http.StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建请求数据
		v := url.Values{}
		v.Set("uid", ms.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+ms.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("updateUser should respond http.StatusInternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctxa, user).
			Return(errx)
		//构建请求数据
		v := url.Values{}
		v.Set("uid", ms.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", ms.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+ms.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	router := gin.New()
	router.DELETE("/user/:uid", srv.deleteUser)

	Convey("deleteUser should respond http.StatusNoContent", t, func() {
		uid := ms.GetUid()
		//mock
		svcm.EXPECT().
			DeleteUser(ctxa, uid).
			Return(nil)
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+ms.StrInt(uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("deleteUser should respond http.StatusBadRequest", t, func() {
		uid := ms.GetUidBad()
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+ms.StrInt(uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("deleteUser should respond http.StatusInternalServerError", t, func() {
		uid := ms.GetUid()
		//mock
		svcm.EXPECT().
			DeleteUser(ctxa, uid).
			Return(errx)
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+ms.StrInt(uid), nil)
		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
