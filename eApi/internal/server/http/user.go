package http

import (
	"context"
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"github.com/unknwon/com"
)

// handValidataError.
func handValidataError(c context.Context, err error) *map[string]interface{} {
	em := make(map[string]interface{})
	// for _, ev := range err.(validator.ValidationErrors){...} //todo
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		em["error"] = e.UserEcodeMap[field]
		em[field] = ev.Value()
		log.Ctx(c).Debug().
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
	}
	return &em
}

// createUser create user.
func (srv *Server) createUser(ctx *gin.Context) {
	c := ctx.MustGet("ctx").(context.Context)
	svc := srv.svc
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()

	log.Ctx(c).Info().
		Msg("start to create user")

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, handValidataError(c, err))
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to create user, data: %v, error: %v", *user, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user, user = %v", *user)
	return
}

// readUser read user.
func (srv *Server) readUser(ctx *gin.Context) {
	c := ctx.MustGet("ctx").(context.Context)
	svc := srv.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.Query("uid")).MustInt64()
	}

	log.Ctx(c).Info().
		Msgf("start to read user, arg: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, handValidataError(c, err))
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return
	}

	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to read user, user = %v", *user)
	return
}

// updateUser update user.
func (srv *Server) updateUser(ctx *gin.Context) {
	c := ctx.MustGet("ctx").(context.Context)
	svc := srv.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()

	log.Ctx(c).Info().
		Msgf("start to update user, arg: %v", uid)

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, handValidataError(c, err))
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to update user, data: %v, error: %v", *user, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to update user, user = %v", *user)
	return
}

// deleteUser delete user.
func (srv *Server) deleteUser(ctx *gin.Context) {
	c := ctx.MustGet("ctx").(context.Context)
	svc := srv.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()

	log.Ctx(c).Info().
		Msgf("start to delete user, arg: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, handValidataError(c, err))
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Ctx(c).Info().
		Int64("user_id", uid).
		Msgf("succ to read user, user = %v", *user)
	return
}
