package http

import (
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// createUser create user.
func (s *Server) createUser(ctx *gin.Context) {
	// 获取参数
	svc := s.svc
	c := ms.GetCtxVal(ctx)
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()

	// 创建数据
	log.Info("start to create user")
	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = name
	user.Sex = sex

	// 检验数据
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return
	}
	log.Infof("succ to create data, user: %v", *user)

	// 使用数据
	c = ms.CarryCtxUserId(c, user.Uid)
	if err := svc.CreateUser(c, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to create user, error: %v", err)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Infof("succ to create user, user: %v", *user)
	return
}

// readUser read user.
func (s *Server) readUser(ctx *gin.Context) {
	svc := s.svc
	c := ms.GetCtxVal(ctx)
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.Query("uid")).MustInt64()
	}
	log.Infof("start to read user, arg: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	user, err := svc.ReadUser(c, user.Uid)
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
	c := ms.GetCtxVal(ctx)
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	log.Infof("start to update user, arg: %v", uid)

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

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.UpdateUser(c, user)
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
	c := ms.GetCtxVal(ctx)
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	log.Infof("start to delete user, arg: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to delete user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Infof("succ to delete user, user: %v", *user)
	return
}
