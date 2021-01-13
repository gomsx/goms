package http

import (
	"net/http"

	m "github.com/fuwensun/goms/eLog/internal/model"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"github.com/unknwon/com"
)

// createUser create user.
func (s *Server) createUser(ctx *gin.Context) {
	svc := s.svc
	c := ms.GetCtxVal(ctx)
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	// 记录参数
	log.Ctx(c).Info().
		Msgf("start to create user, name:%v, sex:%v", name, sex)

	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		// 记录异常
		log.Ctx(c).Info().
			Msgf("failed to validate data, data: %v, error: %v", *user, err)
		return
	}
	// 记录中间结果
	log.Ctx(c).Info().
		Msgf("succ to create data, user: %v", *user)

	c = ms.CarryCtxUserId(c, user.Uid)
	if err := svc.CreateUser(c, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Msgf("failed to create user, data: %v, error: %v", *user, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	// 记录返回结果
	log.Ctx(c).Info().
		Msgf("succ to create user, user: %v", *user)
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
	log.Ctx(c).Info().
		Msgf("start to read user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Ctx(c).Info().
			Msgf("failed to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	log.Ctx(c).Info().
		Msgf("succ to get user uid, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Msgf("failed to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Ctx(c).Info().
		Msgf("succ to read user, user: %v", *user)
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
	log.Ctx(c).Info().
		Msgf("start to update user, name: %v, sex %v", name, sex)

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Ctx(c).Info().
			Msgf("failed to validate data, data: %v, error: %v", *user, err)
		return
	}
	log.Ctx(c).Info().
		Msgf("succ to create data, user: %v", *user)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.UpdateUser(c, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Msgf("failed to update user, data: %v, error: %v", *user, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Ctx(c).Info().
		Msgf("succ to update user, user: %v", *user)
	return
}

// deleteUser delete user.
func (s *Server) deleteUser(ctx *gin.Context) {
	svc := s.svc
	c := ms.GetCtxVal(ctx)
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	log.Ctx(c).Info().
		Msgf("start to delete user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, ms.GetValidateError(err))
		log.Ctx(c).Info().
			Msgf("failed to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	log.Ctx(c).Info().
		Msgf("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Msgf("failed to read user, data: %v, error: %v", user.Uid, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Ctx(c).Info().
		Msgf("succ to read user, user: %v", *user)
	return
}
