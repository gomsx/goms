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

	"github.com/fuwensun/goms/eApi/internal/model"
	. "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/eApi/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var errx = errors.New("error")

func TestCreateUser(t *testing.T) {
	//设置gin测试模式
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	router := gin.New()
	router.POST("/user", srv.createUser)

	var uid int64 = 2
	Patch(model.GetUid, func() int64 {
		return uid
	})

	Convey("createUser should respond http.StatusCreated", t, func() {
		user := &User{
			Uid:  uid,
			Name: "xxx",
			Sex:  1,
		}

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

	Convey("createUser should respond http.StatusBadRequest", t, func() {

		user := &User{
			Uid:  uid,
			Name: "xxx",
			Sex:  99,
		}

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

	Convey("createUser should respond http.StatusInternalServerError", t, func() {

		user := &User{
			Uid:  uid,
			Name: "xxx",
			Sex:  1,
		}

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

		//解析 resp 到 map
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &m)
		if err != nil {
			panic(err)
		}
		// fmt.Println(" ==>", m)

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

		//构建请求
		w := httptest.NewRecorder()
		uidstr := strconv.FormatInt(user.Uid, 10)
		r, _ := http.NewRequest("GET", "/user/"+uidstr, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		//解析 resp 到 map
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &m)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", m)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
		So(m["Uid"], ShouldEqual, float64(user.Uid))
	})

	Convey("readUser should respond http.StatusInternalServerError", t, func() {

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

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
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
			Uid:  GetUid(),
			Name: "xxx",
			Sex:  1,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(nil)

		//构建请求数据
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

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("updateUser should respond http.StatusBadRequest", t, func() {

		user := &User{
			Uid:  -1 * GetUid(),
			Name: "xxx",
			Sex:  1,
		}

		//构建请求数据
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

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("updateUser should respond http.StatusInternalServerError", t, func() {

		user := &User{
			Uid:  GetUid(),
			Name: "xxx",
			Sex:  1,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(errx)

		//构建请求数据
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

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
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
		uid := GetUid()
		uidstr := strconv.FormatInt(uid, 10)
		svcm.EXPECT().
			DeleteUser(gomock.Any(), uid).
			Return(nil)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+uidstr, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("deleteUser should respond http.StatusBadRequest", t, func() {
		uid := -1 * GetUid()
		uidstr := strconv.FormatInt(uid, 10)
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+uidstr, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("deleteUser should respond http.StatusInternalServerError", t, func() {
		uid := GetUid()
		uidstr := strconv.FormatInt(uid, 10)
		svcm.EXPECT().
			DeleteUser(gomock.Any(), uid).
			Return(errx)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+uidstr, nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
