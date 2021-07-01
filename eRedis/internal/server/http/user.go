package http

import (
	"net/http"

	m "github.com/gomsx/goms/eRedis/internal/model"
	e "github.com/gomsx/goms/pkg/err"
	"github.com/gomsx/goms/pkg/id"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/unknwon/com"
)

// CreateUser create user.
func (s *Server) createUser(ctx *gin.Context) {
	svc := s.svc
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()

	user := &m.User{}
	user.Uid = id.GenUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		return
	}

	if err := svc.CreateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	return
}

// ReadUser read user.
func (s *Server) readUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.Query("uid")).MustInt64()
	}

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		return
	}

	user, err := svc.ReadUser(ctx, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	return
}

// UpdateUser update user.
func (s *Server) updateUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		return
	}

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //update ok
	return
}

// DeleteUser delete user.
func (s *Server) deleteUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		return
	}

	err := svc.DeleteUser(ctx, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	return
}

// GetValidateError get validate error.
func GetValidateError(err error) *map[string]interface{} {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	value := ev.Value()

	em := make(map[string]interface{})
	em["error"] = e.UserEcodeMap[field]
	em[field] = value
	return &em
}
