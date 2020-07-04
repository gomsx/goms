package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/aivuca/goms/eRedis/internal/model"
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
		log.Printf("http name err: %v", namestr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrNameError.Error(),
			"name":  namestr,
		})
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		log.Printf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		return
	}

	user.Name = namestr
	user.Sex = sex

	if err = svc.CreateUser(c, &user); err != nil {
		log.Printf("http create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http create user=%v", user)
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
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}

	user, err := svc.ReadUser(c, uid)
	if err == ErrNotFoundData {
		log.Printf("http read user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Printf("http read user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http read user=%v", user)
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
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}
	sex, ok := CheckSexS(sexstr)
	if !ok {
		log.Printf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrSexError.Error(),
			"sex":   sexstr,
		})
		return
	}
	ok = CheckName(namestr)
	if !ok {
		log.Printf("http name err: %v", namestr)
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
		// 问题：当资源不存在时，是返回 404;
		// 还是返回 200 再在报文体里说明资源不存在，
		// 即 http 协议返回码和用户状态码分离
		log.Printf("http update user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Printf("http update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Printf("http update user=%v", user)
	return
}

// deleteUser
func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	uid, ok := CheckUidS(uidstr)
	if !ok {
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrUidError.Error(),
			"uid":   uidstr,
		})
		return
	}

	err := svc.DeleteUser(c, uid)
	if err == ErrNotFoundData {
		log.Printf("http delete user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		log.Printf("http delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Printf("http delete user uid=%v", uid)
	return
}

