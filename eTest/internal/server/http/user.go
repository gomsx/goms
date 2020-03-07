package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var err error
	user := model.User{}

	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err:%v\n", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "sex err!",
			"uid":   sexstr,
		})
		return
	}
	user.Name = namestr

	if err = svc.CreateUser(c, &user); err != nil {
		log.Printf("http create user, err: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "create failed!",
			"uid":   user.Uid,
			"name":  user.Name,
			"sex":   user.Sex,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http create user=%v\n", user)
}

func updateUser(c *gin.Context) {
	var err error
	user := model.User{}
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}
	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	user.Uid, err = strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid err!",
			"uid":   uidstr,
		})
		return
	}
	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err:%v\n", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "sex err!",
			"uid":   sexstr,
		})
		return
	}
	user.Name = namestr
	err = svc.UpdateUser(c, &user)
	if err != nil {
		log.Printf("http update user,err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   user.Uid,
			"name":  user.Name,
			"sex":   user.Sex,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http update user=%v\n", user)
}

func readUser(c *gin.Context) {
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid err!",
			"uid":   uidstr,
		})
		return
	}
	user, err := svc.ReadUser(c, uid)
	if err != nil {
		log.Printf("http read user,err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   uidstr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http read user=%v\n", user)
}

func deleteUser(c *gin.Context) {
	uidstr := c.Param("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid err!",
			"uid":   uidstr,
		})
		return
	}
	if err = svc.DeleteUser(c, uid); err != nil {
		log.Printf("http delete user,err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   uidstr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uid": uid})
	log.Printf("http delete user uid=%v\n", uid)
}
