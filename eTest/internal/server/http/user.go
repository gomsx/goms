package http

import (
	"net/http"

	m "github.com/aivuca/goms/eTest/internal/model"
	e "github.com/aivuca/goms/eTest/internal/pkg/err"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"github.com/unknwon/com"
)

// handValidateError hand validate error.
func handValidateError(err error) *map[string]interface{} {
	em := make(map[string]interface{})
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		em["error"] = e.UserEcodeMap[field]
		em[field] = ev.Value()
		log.Debug().
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
	}
	return &em
}

// createUser create user.
func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc
	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	// 记录参数
	log.Ctx(c).Info().
		Msgf("start to create user, name:%v, sex:%v", name, sex)

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, handValidateError(err))
		// 记录异常
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return
	}
	// 记录中间结果
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to create user, data: %v, error: %v", *user, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	// 记录返回结果
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user, user = %v", *user)
	return
}

// readUser read user.
func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uid := com.StrTo(c.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(c.Query("uid")).MustInt64()
	}

	log.Ctx(c).Info().
		Msgf("start to read user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		c.JSON(http.StatusBadRequest, handValidateError(err))
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return
	}

	log.Ctx(c).Info().Msgf("succ to get user uid, uid = %v", user.Uid)

	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
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
func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc
	uid := com.StrTo(c.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(c.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	log.Ctx(c).Info().
		Msgf("start to update user, name: %v, sex %v", name, sex)

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, handValidateError(err))
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to update user, data: %v, error: %v", *user, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to update user, user = %v", *user)
	return
}

// deleteUser delete user.
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uid := com.StrTo(c.Param("uid")).MustInt64()

	log.Ctx(c).Info().
		Msgf("start to delete user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		c.JSON(http.StatusBadRequest, handValidateError(err))
		log.Ctx(c).Info().
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to read user, user = %v", *user)
	return
}
