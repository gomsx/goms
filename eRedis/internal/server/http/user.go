package http

import (
	"log"
	"strconv"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/gin-gonic/gin"
)

func updateUser(c *gin.Context) {
	var err error
	user := model.User{}

	uidstr := c.Query("uid")
	namestr := c.Query("name")
	sexstr := c.Query("sex")

	user.Uid, err = strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", user.Uid)
		return
	}
	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("sex err:%v\n", user.Sex)
		return
	}
	user.Name = namestr

	err = svc.UpdateUser(c, &user)
	if err != nil {
		c.JSON(404, gin.H{"error": "xxx"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http updateuser %v\n", user)
}

func readUser(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	user, err := svc.ReadUser(c, uid)
	if err != nil {
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http readuser %v\n", user)
}

//
func updatename(c *gin.Context) {

	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	name := c.Query("name")
	if name == "" {
		log.Printf("name err:%v\n", uidstr)
		return
	}
	c.JSON(200, gin.H{
		"uid":  uidstr,
		"name": name,
	})
	log.Printf("http user updatename %v to %v\n", uidstr, name)
	svc.UpdateUserName(c, uid, name)
}

func readname(c *gin.Context) {

	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	name, err := svc.ReadUserName(c, uid)
	if err != nil {
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"uid":  uidstr,
		"name": name,
	})
	log.Printf("http user readname %v to %v\n", uidstr, name)
}
