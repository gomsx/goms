package http

import (
	"net/http"

	. "github.com/aivuca/goms/eRedis/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/unknwon/com"
)

func handValidateError(err error) *map[string]interface{} {
	m := make(map[string]interface{})
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		m["error"] = UserEcodeMap[field]
		m[field] = ev.Value()
	}
	return &m
}

// createUser
func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc

	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	user := &User{}
	user.Uid = GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		m := handValidateError(err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	if err := svc.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	return
}

// readUser
func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}
	uid := com.StrTo(uidstr).MustInt64()

	user := &User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		m := handValidateError(err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	user, err := svc.ReadUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})

	return
}

// updateUser
func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc

	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}

	uid := com.StrTo(uidstr).MustInt64()
	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	user := &User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		m := handValidateError(err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	err := svc.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok

	return
}

// deleteUser
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	uid := com.StrTo(uidstr).MustInt64()

	user := &User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		m := handValidateError(err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	err := svc.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok

	return
}
