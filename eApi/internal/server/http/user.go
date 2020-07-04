package http

import (
	"net/http"

	. "github.com/aivuca/goms/eApi/internal/model"

	"github.com/gin-gonic/gin"
)

// createUser
func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc

	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	log.Debug().Msg("start to create user")

	ok := CheckName(namestr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrNameError.Error(),
			"name":  namestr,
		})
		log.Debug().Msgf("name err, name = %v", namestr)
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		log.Debug().Msgf("sex err, sex = %v", sexstr)
		return
	}

	user := &User{}
	user.Name = namestr
	user.Sex = sex

	log.Debug().Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().Int64("uid", user.Uid).Msg("failed to create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().Int64("uid", user.Uid).Msg("succ to create user")
	return
}

// readUser
func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}

	log.Debug().Msg("start to read user")

	uid, ok := CheckUidS(uidstr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		log.Debug().Msgf("uid err, uid = %v", uidstr)
		return
	}

	log.Debug().Msgf("succ to get user uid, uid = %v", uid)

	user, err := svc.ReadUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().Int64("uid", user.Uid).Msg("failed to read user")
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().Int64("uid", user.Uid).Msg("succ to read user")
	return
}

// updateUser
func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc

	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}
	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	log.Debug().Msg("start to update user")

	uid, ok := CheckUidS(uidstr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		log.Debug().Msgf("uid err, uid = %v", uidstr)
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		log.Debug().Msgf("sex err, sex = %v", sexstr)
		return
	}
	ok = CheckName(namestr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrNameError.Error(),
			"name":  namestr,
		})
		log.Debug().Msgf("name err, name = %v", namestr)
		return
	}

	user := &User{}
	user.Uid = uid
	user.Name = namestr
	user.Sex = sex

	log.Debug().Msgf("succ to get user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().Int64("uid", user.Uid).Msg("failed to update user")
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Info().Int64("uid", user.Uid).Msg("succ to update user")
	return
}

// deleteUser
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")

	log.Debug().Msg("start to delete user")

	uid, ok := CheckUidS(uidstr)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		log.Debug().Msgf("uid err, uid = %v", uidstr)
		return
	}

	log.Debug().Msgf("succ to get user uid, uid = %v", uid)

	err := svc.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		log.Info().Int64("uid", uid).Msg("failed to delete user")
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Info().Int64("uid", uid).Msg("succ to delete user")
	return
}

