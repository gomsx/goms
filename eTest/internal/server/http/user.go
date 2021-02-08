package http

import (
	"net/http"

	m "github.com/fuwensun/goms/eTest/internal/model"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// createUser create user.
func (s *Server) createUser(ctx *gin.Context) {
	svc := s.svc
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	// 记录参数
	log.Infof("start to create user, name:%v, sex:%v", name, sex)

	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		// 记录异常
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return
	}
	// 记录中间结果
	log.Infof("succ to create data, user: %v", *user)

	if err := svc.CreateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to create user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	// 记录返回结果
	log.Infof("succ to create user, user: %v", *user)
	return
}

// readUser read user.
func (s *Server) readUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.Query("uid")).MustInt64()
	}
	log.Infof("start to read user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	user, err := svc.ReadUser(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to read user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Infof("succ to read user, user: %v", *user)
	return
}

// updateUser update user.
func (s *Server) updateUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	log.Infof("start to update user, name: %v, sex %v", name, sex)

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return
	}
	log.Infof("succ to create data, user: %v", *user)

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to update user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Infof("succ to update user, user: %v", *user)
	return
}

// deleteUser delete user.
func (s *Server) deleteUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	log.Infof("start to delete user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	err := svc.DeleteUser(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to delete user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Infof("succ to delete user, user: %v", *user)
	return
}
