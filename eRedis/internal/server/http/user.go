package http

import (
	"log"
	"strconv"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var err error
	user := model.User{}

	namestr := c.Query("name")
	sexstr := c.Query("sex")

	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("sex err:%v\n", user.Sex)
		return
	}
	user.Name = namestr

	err = svc.CreateUser(c, &user)
	if err != nil {
		c.JSON(404, gin.H{"error": "create failed!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http create user %v\n", user)
}

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
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http update user %v\n", user)
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
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http read user %v\n", user)
}

func deleteUser(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	err = svc.DeleteUser(c, uid)
	if err != nil {
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid": uid,
	})
	log.Printf("http delete user %v\n", uid)
}
