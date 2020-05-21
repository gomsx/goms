package http

import (
	"net/http"

	. "github.com/fuwensun/goms/eLog/internal/model"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"
)

// createUser
func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc
	var err error
	user := User{}
	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	ok := CheckName(namestr)
	if !ok {
		log.Info().Msgf("http name err: %v", namestr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrNameError.Error(),
			"name":  namestr,
		})
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		log.Info().Msgf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		return
	}

	user.Name = namestr
	user.Sex = sex

	if err = svc.CreateUser(c, &user); err != nil {
		log.Info().Msgf("http create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().Msgf("http create user=%v", user)
	return
}

// updateUser
func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc
	var err error
	user := User{}
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}
	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	uid, ok := CheckUidS(uidstr)
	if !ok {
		log.Info().Msgf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		log.Info().Msgf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		return
	}
	ok = CheckName(namestr)
	if !ok {
		log.Info().Msgf("http name err: %v", namestr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrNameError.Error(),
			"name":  namestr,
		})
		return
	}

	user.Uid = uid
	user.Name = namestr
	user.Sex = sex

	err = svc.UpdateUser(c, &user)
	if err == ErrNotFoundData {
		log.Info().Msgf("http update user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Info().Msgf("http update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Info().Msgf("http update user=%v", user)
	return
}

// readUser
func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}
	uid, ok := CheckUidS(uidstr)
	if !ok {
		log.Info().Msgf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}

	user, err := svc.ReadUser(c, uid)
	if err == ErrNotFoundData {
		log.Info().Msgf("http read user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Info().Msgf("http read user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Info().Msgf("http read user=%v", user)
	return
}

// deleteUser
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	uid, ok := CheckUidS(uidstr)
	if !ok {
		log.Info().Msgf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}

	err := svc.DeleteUser(c, uid)
	if err == ErrNotFoundData {
		log.Info().Msgf("http delete user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Info().Msgf("http delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Info().Msgf("http delete user uid=%v", uid)
	return
}
