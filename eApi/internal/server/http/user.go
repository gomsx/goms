package http

import (
	"context"
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
		log.Debug().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
	}
	return &em
}

// createUser create user.
func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc

	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to create user")

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		m := handValidataError(c, err)
		c.JSON(http.StatusBadRequest, m)
		return
	}
	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to create user")
	return
}

// readUser read user.
func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}
	uid := com.StrTo(uidstr).MustInt64()

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to read user")

	user := &m.User{}
	user.Uid = uid

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to create user data, uid = %v", user.Uid)

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		m := handValidataError(c, err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to create user data, uid = %v", uid)

	user, err := svc.ReadUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to read user")
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to read user")
	return
}

// updateUser update user.
func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc

	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}

	uid := com.StrTo(uidstr).MustInt64()
	name := com.StrTo(c.PostForm("name")).String()
	sex := com.StrTo(c.PostForm("sex")).MustInt64()

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to update user")

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		m := handValidataError(c, err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to create user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to update user")
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to update user")
	return
}

// deleteUser delete user.
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	uid := com.StrTo(uidstr).MustInt64()

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to delete user")

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		m := handValidataError(c, err)
		c.JSON(http.StatusBadRequest, m)
		return
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to create user data, uid = %v", uid)

	err := svc.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", uid).
			Msg("failed to delete user")
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", uid).
		Msg("succ to delete user")
	return
}
